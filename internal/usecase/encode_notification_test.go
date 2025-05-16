package usecase_test

import (
	"fmt"
	"testing"

	"github.com/z6wdc/go-avro/internal/entity"
	"github.com/z6wdc/go-avro/internal/mocks"
	"github.com/z6wdc/go-avro/internal/usecase"
	"go.uber.org/mock/gomock"
)

func TestEncodeNotificationUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEncoder := mocks.NewMockNotificationEncoder(ctrl)
	uc := usecase.NewEncodeNotificationUseCase(mockEncoder)

	input := &entity.Notification{
		ID:      "001",
		UserID:  10,
		Message: "hello",
	}

	expected := []byte{0x01, 0x02, 0x03}

	mockEncoder.EXPECT().
		Encode(input).
		Return(expected, nil)

	result, err := uc.Execute(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(result) != string(expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestEncodeNotificationUseCase_Execute_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEncoder := mocks.NewMockNotificationEncoder(ctrl)
    uc := usecase.NewEncodeNotificationUseCase(mockEncoder)

    input := &entity.Notification{
        ID:      "fail",
        UserID:  99,
        Message: "fail",
    }

    mockEncoder.EXPECT().
        Encode(input).
        Return(nil, fmt.Errorf("mock encode error"))

    _, err := uc.Execute(input)
    if err == nil {
        t.Fatal("expected error, got nil")
    }
    if err.Error() != "encode error: mock encode error" {
        t.Errorf("unexpected error message: %v", err)
    }
}
