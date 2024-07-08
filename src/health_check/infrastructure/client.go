package infrastructure

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip14"
	"github.com/gemalto/kmip-go/ttlv"
	"time"
)

const (
	tcpNet = "tcp"
)

type (
	KMIPClient interface {
		Create(ctx context.Context, msg kmip.RequestMessage) (*kmip.CreateResponsePayload, error)
		Destroy(ctx context.Context, msg kmip.RequestMessage) (*kmip.DestroyResponsePayload, error)
	}
	KMIPClientImpl struct {
		addr       string
		timeout    time.Duration
		cert       tls.Certificate
		cipherType uint16
	}
	config struct {
		Host       string `inject:"config:app.kmipServer.host"`
		Port       string `inject:"config:app.kmipServer.port"`
		Timeout    string `inject:"config:app.kmipServer.timeout"`
		CertFile   string `inject:"config:app.kmipServer.certFile"`
		KeyFile    string `inject:"config:app.kmipServer.keyFile"`
		CipherType int    `inject:"config:app.kmipServer.cipherType"`
	}
)

func (c *KMIPClientImpl) Inject(cfg *config) {
	duration, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		panic("failed to parse timeout")
	}

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		panic("failed to load TLS certificate")
	}

	c.cert = cert
	c.timeout = duration
	c.addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	c.cipherType = uint16(cfg.CipherType)
}

func (c *KMIPClientImpl) Create(ctx context.Context, msg kmip.RequestMessage) (*kmip.CreateResponsePayload, error) {
	bi, err := c.execute(ctx, msg)
	if err != nil {
		return nil, err
	}

	decoder := ttlv.Decoder{}
	var respPayload kmip.CreateResponsePayload
	err = decoder.DecodeValue(&respPayload, bi.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, err
	}

	return &respPayload, nil
}

func (c *KMIPClientImpl) Destroy(ctx context.Context, msg kmip.RequestMessage) (*kmip.DestroyResponsePayload, error) {
	bi, err := c.execute(ctx, msg)
	if err != nil {
		return nil, err
	}

	decoder := ttlv.Decoder{}
	var respPayload kmip.DestroyResponsePayload
	err = decoder.DecodeValue(&respPayload, bi.ResponsePayload.(ttlv.TTLV))
	if err != nil {
		return nil, err
	}

	return &respPayload, nil
}

func (c *KMIPClientImpl) execute(ctx context.Context, msg kmip.RequestMessage) (*kmip.ResponseBatchItem, error) {
	conn, err := c.tlsConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	req, err := ttlv.Marshal(msg)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(req)
	if err != nil {
		return nil, err
	}

	decoder := ttlv.NewDecoder(conn)
	resp, err := decoder.NextTTLV()

	var respMsg kmip.ResponseMessage
	err = decoder.DecodeValue(&respMsg, resp)
	if err != nil {
		return nil, err
	}

	// we only expect one item
	if respMsg.ResponseHeader.BatchCount != 1 {
		return nil, fmt.Errorf("strange batch count response from KMIP server: %d", respMsg.ResponseHeader.BatchCount)
	}

	bi := respMsg.BatchItem[0]
	if bi.ResultStatus != kmip14.ResultStatusSuccess {
		return nil, fmt.Errorf("error response from KMIP server: %s", bi.ResultMessage)
	}

	return &bi, nil
}

func (c *KMIPClientImpl) tlsConn(ctx context.Context) (*tls.Conn, error) {
	dialer := tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{c.cert},
			CipherSuites:       []uint16{c.cipherType},
		},
	}

	conn, err := dialer.DialContext(ctx, tcpNet, c.addr)
	if err != nil {
		return nil, err
	}

	return conn.(*tls.Conn), nil
}
