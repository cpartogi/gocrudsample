package usecase

import (
	"context"
	"errors"
	"gocrudsample/domain/tutorial/mocks"
	"gocrudsample/domain/tutorial/model"
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestTutorialTypesError(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)
		mockRepo.On("GetTutorialTypes", mock.Anything, mock.Anything).Return(nil, errors.New("failed"))

		_, err := u.GetTutorialTypes(context.Background())

		assert.Error(t, err, "failed")
	})
}

func TestTutorialTypesSuccess(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)
		mockRepo.On("GetTutorialTypes", mock.Anything, mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "abc",
				TypeName: "type name",
			},
		}, nil)

		_, err := u.GetTutorialTypes(context.Background())

		assert.NilError(t, err)
	})
}
