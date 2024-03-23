package usecase

import (
	"context"
	"errors"
	"gocrudssample/domain/tutorial/mocks"
	"gocrudssample/domain/tutorial/model"
	"testing"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestGetTutorialsError(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error invalid id", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		_, err := u.GetTutorials(context.Background(), "abc")

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error get data", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		mockRepo.On("GetTutorials", mock.Anything, mock.Anything).Return([]model.Tutorials{
			{
				Id:             "1",
				TutorialTypeId: "2",
				Keywords:       "keywords",
				Sequence:       0,
				Title:          "title tutorial",
			},
		}, errors.New("failed"))
		_, err := u.GetTutorials(context.Background(), "89bc2029-1ed4-461c-8c0f-79c9489e04a2")

		assert.Error(t, err, "failed")
	})
}

func TestGetTutorialsSuccess(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Success", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		mockRepo.On("GetTutorials", mock.Anything, mock.Anything).Return([]model.Tutorials{
			{
				Id:             "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
				TutorialTypeId: "2",
				Keywords:       "keywords",
				Sequence:       0,
				Title:          "title tutorial",
			},
		}, nil)
		_, err := u.GetTutorials(context.Background(), "89bc2029-1ed4-461c-8c0f-79c9489e04a2")

		assert.NilError(t, err)
	})
}
