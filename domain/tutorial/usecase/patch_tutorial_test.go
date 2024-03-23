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

func TestPatchTutorial(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error invalid id", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id: "1",
		}

		err := u.PatchTutorial(context.Background(), req)

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error title required", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
		}

		err := u.PatchTutorial(context.Background(), req)

		assert.Error(t, err, "Title required")
	})

	t.Run("Error patch", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:    "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Title: "Tutorial Title",
		}

		mockRepo.On("PatchTutorial", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		err := u.PatchTutorial(context.Background(), req)

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:    "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Title: "Tutorial Title",
		}

		mockRepo.On("PatchTutorial", mock.Anything, mock.Anything).Return(nil).Once()

		err := u.PatchTutorial(context.Background(), req)

		assert.NilError(t, err)
	})

}
