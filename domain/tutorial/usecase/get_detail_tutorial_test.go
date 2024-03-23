package usecase

import (
	"context"
	"errors"
	"gocrudssample/domain/tutorial/mocks"
	"gocrudssample/domain/tutorial/model"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"gotest.tools/assert"
)

func TestGetDetailTutorial(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error invalid id format", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)
		mockRepo.On("GetDetailTutorial", mock.Anything, mock.Anything).Return(model.Tutorials{
			Id:             "1",
			TutorialTypeId: "2",
			Keywords:       "3",
		}, errors.New("failed")).Once()

		_, err := u.GetDetailTutorial(context.Background(), "abc")

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error get data", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)
		mockRepo.On("GetDetailTutorial", mock.Anything, mock.Anything).Return(model.Tutorials{
			Id:             "1",
			TutorialTypeId: "2",
			Keywords:       "3",
		}, errors.New("failed")).Once()

		_, err := u.GetDetailTutorial(context.Background(), "89bc2029-1ed4-461c-8c0f-79c9489e04a2")

		assert.Error(t, err, "failed")
	})
}

func TestGetDetailTutorialSuccess(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Success", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)
		mockRepo.On("GetDetailTutorial", mock.Anything, mock.Anything).Return(model.Tutorials{
			Id:             "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "ada",
			Title:          "title tutorial",
		}, nil).Once()

		_, err := u.GetDetailTutorial(context.Background(), "89bc2029-1ed4-461c-8c0f-79c9489e04a2")

		assert.NilError(t, err)
	})

	t.Run("Success updated at", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		updatedAt := time.Now().UTC()
		updatedBy := "Admin"

		mockRepo.On("GetDetailTutorial", mock.Anything, mock.Anything).Return(model.Tutorials{
			Id:             "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "ada",
			Title:          "title tutorial",
			UpdatedAt:      &updatedAt,
			UpdatedBy:      &updatedBy,
		}, nil).Once()

		_, err := u.GetDetailTutorial(context.Background(), "89bc2029-1ed4-461c-8c0f-79c9489e04a2")

		assert.NilError(t, err)
	})
}
