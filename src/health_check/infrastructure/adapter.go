package infrastructure

import (
	"context"

	"github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip14"
	"github.com/google/uuid"
)

// KMIPAdapter is an implementation of the KMIPRepository interface
type KMIPAdapter struct {
	client KMIPClient
}

// Inject dependencies
func (a *KMIPAdapter) Inject(c KMIPClient) {
	a.client = c
}

// Create passes a create-request object to the client and returns its UniqueIdentifier after creation
func (a *KMIPAdapter) Create(ctx context.Context) (string, error) {
	resp, err := a.client.Create(ctx, createReq())
	if err != nil {
		return "", err
	}

	return resp.UniqueIdentifier, nil
}

// Destroy passes a destroy-request object to the client
func (a *KMIPAdapter) Destroy(ctx context.Context, id string) error {
	_, err := a.client.Destroy(ctx, destroyReq(id))
	if err != nil {
		return err
	}

	return nil
}

// createReq creates a request object for creating a key
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

// createReq creates a request object for destroying a key
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
