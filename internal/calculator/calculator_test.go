package calculator

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/iamnande/ff8-api/internal/datastore"
	mockDatastore "github.com/iamnande/ff8-api/internal/datastore/mocks"
)

var (
	calculate Calculator
	ctx       context.Context
	ctrl      *gomock.Controller
	ds        *mockDatastore.MockDatastore
)

// testSetup handles common test setup actions for the calculator package
func testSetup(t *testing.T) {
	ctx = context.Background()
	ctrl = gomock.NewController(t)
	ds = mockDatastore.NewMockDatastore(ctrl)
	calculate = NewCalculator(ds)
}

func TestNewCalculator(t *testing.T) {
	testSetup(t)

	// test: validations
	if calculate == nil {
		t.Errorf("failed to initialize new calculator instance")
	}

}

func TestCardMagicRatio_Successful(t *testing.T) {
	testSetup(t)

	// test: case setup (mock expectations)
	expectedCount := 15
	expectedName := "Firaga"
	expectedCard := "Hexadragon"

	ds.EXPECT().DescribeRecord(gomock.Any(), gomock.Any()).Return(&datastore.Record{
		Name:           expectedName,
		Type:           datastore.Magic,
		CardEquivalent: expectedCard,
		CardMagicRatio: 20.0,
	}, nil).Times(1)

	// test: execution
	actual, err := calculate.CardMagicRatio(ctx, expectedName, datastore.Magic, 300.0)

	// test: error checking
	if err != nil {
		t.Errorf(" did not expect error: %s", err)
	}

	// test: value checking
	if expectedCard != actual.Card || expectedCount != actual.Quantity {
		t.Errorf("expected-card=%s actual-card=%s expected-count=%v actual-count=%d",
			expectedCard, actual.Card, expectedCount, actual.Quantity)
	}

}

func TestCardMagicRatio_ErrItemNotFound(t *testing.T) {
	testSetup(t)

	// test: case setup (mock expectations)
	ds.EXPECT().DescribeRecord(gomock.Any(), gomock.Any()).Return(nil, datastore.ErrItemNotFound).Times(1)

	// test: execution
	actual, err := calculate.CardMagicRatio(ctx, "Aero", datastore.Magic, 100.0)

	// test: validate we received an error
	if err == nil || actual != nil {
		t.Errorf("expected error, received value: %+v\n", actual)
	}

	if !errors.Is(err, datastore.ErrItemNotFound) {
		t.Errorf("expected ErrItemNotFound, received %+v", err)
	}

}
