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

func TestDeleteTutorial(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error invalid id", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id: "1",
		}

		err := u.DeleteTutorial(context.Background(), req)

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error delete tutorial", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
		}

		mockRepo.On("DeleteTutorial", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		err := u.DeleteTutorial(context.Background(), req)

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
		}

		mockRepo.On("DeleteTutorial", mock.Anything, mock.Anything).Return(nil).Once()

		err := u.DeleteTutorial(context.Background(), req)

		assert.NilError(t, err)
	})

}
