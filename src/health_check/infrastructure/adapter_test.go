package infrastructure

import (
	"context"
	"errors"
	"github.com/gemalto/kmip-go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type KMIPAdapterTestSuite struct {
	suite.Suite

	context        context.Context
	kmipClientMock *MockKMIPClient
	kmipAdapter    *KMIPAdapter
}

func (t *KMIPAdapterTestSuite) SetupTest() {
	t.context = context.Background()
	t.kmipClientMock = &MockKMIPClient{}
	t.kmipAdapter = &KMIPAdapter{}

	t.kmipAdapter.Inject(t.kmipClientMock)
}

func TestHealthCheckServiceTestSuite(t *testing.T) {
	suite.Run(t, new(KMIPAdapterTestSuite))
}

func (t *KMIPAdapterTestSuite) TestCreate() {
	t.Run("success", func() {
		createResp := kmip.CreateResponsePayload{
			UniqueIdentifier: "123",
		}

		t.kmipClientMock.On("Create", t.context, mock.IsType(kmip.RequestMessage{})).Return(&createResp, nil).Once()

		id, err := t.kmipAdapter.Create(t.context)

		t.NoError(err)
		t.Equal(createResp.UniqueIdentifier, id)
	})

	t.Run("creating returns error", func() {
		t.kmipClientMock.On("Create", t.context, mock.IsType(kmip.RequestMessage{})).Return(nil, errors.New("error")).Once()

		id, err := t.kmipAdapter.Create(t.context)

		t.Error(err)
		t.Empty(id)
	})
}

func (t *KMIPAdapterTestSuite) TestDestroy() {
	t.Run("success", func() {
		destroyResp := kmip.DestroyResponsePayload{
			UniqueIdentifier: "123",
		}

		t.kmipClientMock.On("Destroy", t.context, mock.IsType(kmip.RequestMessage{})).Return(&destroyResp, nil).Once()

		err := t.kmipAdapter.Destroy(t.context, "1")

		t.NoError(err)
	})

	t.Run("creating returns error", func() {
		t.kmipClientMock.On("Destroy", t.context, mock.IsType(kmip.RequestMessage{})).Return(nil, errors.New("error")).Once()

		err := t.kmipAdapter.Destroy(t.context, "1")

		t.Error(err)
	})
}
