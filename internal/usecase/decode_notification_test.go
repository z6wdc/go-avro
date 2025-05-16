package usecase_test

import (
	"fmt"
	"testing"

	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/mocks"
	"github.com/z6wdc/go-avro/internal/usecase"
	"go.uber.org/mock/gomock"
)

func TestDecodeNotificationUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDecoder := mocks.NewMockNotificationDecoder(ctrl)
	uc := usecase.NewDecodeNotificationUseCase(mockDecoder)

	input := []byte{0x01, 0x02, 0x03}
	expected := &entity.Notification{
		ID:      "001",
		UserID:  10,
		Message: "decoded",
	}

	mockDecoder.EXPECT().
		Decode(input).
		Return(expected, nil)

	result, err := uc.Execute(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if *result != *expected {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}

func TestDecodeNotificationUseCase_Execute_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockDecoder := mocks.NewMockNotificationDecoder(ctrl)
    uc := usecase.NewDecodeNotificationUseCase(mockDecoder)

    input := []byte{0x00}

    mockDecoder.EXPECT().
        Decode(input).
        Return(nil, fmt.Errorf("mock decode error"))

    _, err := uc.Execute(input)
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != "decode error: mock decode error" {
        t.Errorf("unexpected error message: %v", err)
    }
}
