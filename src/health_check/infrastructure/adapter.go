package infrastructure

import (
	"context"
	"github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip14"
	"github.com/google/uuid"
)

type KMIPAdapter struct {
	client KMIPClient
}

func (a *KMIPAdapter) Inject(c KMIPClient) {
	a.client = c
}

func (a *KMIPAdapter) Create(ctx context.Context) (string, error) {
	resp, err := a.client.Create(ctx, createReq())
	if err != nil {
		return "", err
	}

	return resp.UniqueIdentifier, nil
}

func (a *KMIPAdapter) Destroy(ctx context.Context, id string) error {
	_, err := a.client.Destroy(ctx, destroyReq(id))
	if err != nil {
		return err
	}

	return nil
}

func createReq() kmip.RequestMessage {
	biID := uuid.New()

	payload := kmip.CreateRequestPayload{
		ObjectType: kmip14.ObjectTypeSymmetricKey,
	}

	payload.TemplateAttribute.Append(kmip14.TagCryptographicAlgorithm, kmip14.CryptographicAlgorithmAES)
	payload.TemplateAttribute.Append(kmip14.TagCryptographicLength, 256)
	payload.TemplateAttribute.Append(kmip14.TagCryptographicUsageMask, kmip14.CryptographicUsageMaskEncrypt|kmip14.CryptographicUsageMaskDecrypt)
	payload.TemplateAttribute.Append(kmip14.TagName, kmip.Name{
		NameValue: biID.String(),
		NameType:  kmip14.NameTypeUninterpretedTextString,
	})

	return kmip.RequestMessage{
		RequestHeader: kmip.RequestHeader{
			ProtocolVersion: kmip.ProtocolVersion{
				ProtocolVersionMajor: 1,
				ProtocolVersionMinor: 4,
			},
			BatchCount: 1,
		},
		BatchItem: []kmip.RequestBatchItem{
			{
				UniqueBatchItemID: biID[:],
				Operation:         kmip14.OperationCreate,
				RequestPayload:    &payload,
			},
		},
	}
}

func destroyReq(id string) kmip.RequestMessage {
	biID := uuid.New()

	payload := kmip.DestroyRequestPayload{
		UniqueIdentifier: id,
	}

	return kmip.RequestMessage{
		RequestHeader: kmip.RequestHeader{
			ProtocolVersion: kmip.ProtocolVersion{
				ProtocolVersionMajor: 1,
				ProtocolVersionMinor: 4,
			},
			BatchCount: 1,
		},
		BatchItem: []kmip.RequestBatchItem{
			{
				UniqueBatchItemID: biID[:],
				Operation:         kmip14.OperationDestroy,
				RequestPayload:    &payload,
			},
		},
	}
}
