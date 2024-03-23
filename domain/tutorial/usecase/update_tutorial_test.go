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

func TestUpdateTutorial(t *testing.T) {
	mockRepo := new(mocks.TutorialRepoInterface)

	t.Run("Error invalid id", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			TutorialTypeId: "abc",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error invalid tutorial type id", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			TutorialTypeId: "abc",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "Invalid Id format (uuid required)")
	})

	t.Run("Error get tutorial types", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "5dca5b91-46cf-49c7-b827-e825f7de04ba",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		mockRepo.On("GetTutorialTypes", mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "2",
				TypeName: "4",
			},
		}, errors.New("failed")).Once()
		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "failed")
	})

	t.Run("Error tutorial types not found", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "5dca5b91-46cf-49c7-b827-e825f7de04ba",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		mockRepo.On("GetTutorialTypes", mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "5dca5b91-46cf-49c7-b827-e825f7de04ba",
				TypeName: "golang",
			},
		}, nil).Once()
		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "Tutorial Type not found")
	})

	t.Run("Error title empty", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "5dca5b91-46cf-49c7-b827-e825f7de04ba",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "",
			Description:    "description",
		}

		mockRepo.On("GetTutorialTypes", mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
				TypeName: "golang",
			},
		}, nil).Once()
		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "Title required")
	})

	t.Run("Error add tutorial", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "5dca5b91-46cf-49c7-b827-e825f7de04ba",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		mockRepo.On("GetTutorialTypes", mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
				TypeName: "golang",
			},
		}, nil).Once()

		mockRepo.On("UpdateTutorial", mock.Anything, mock.Anything).Return(errors.New("failed")).Once()

		err := u.UpdateTutorial(context.Background(), req)

		assert.Error(t, err, "failed")
	})

	t.Run("Success", func(t *testing.T) {
		u := NewTutorialUsecase(mockRepo, 1000)

		req := model.Tutorials{
			Id:             "5dca5b91-46cf-49c7-b827-e825f7de04ba",
			TutorialTypeId: "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
			Keywords:       "keywords",
			Sequence:       1,
			Title:          "title",
			Description:    "description",
		}

		mockRepo.On("GetTutorialTypes", mock.Anything).Return([]model.TutorialTypes{
			{
				Id:       "89bc2029-1ed4-461c-8c0f-79c9489e04a2",
				TypeName: "golang",
			},
		}, nil).Once()

		mockRepo.On("UpdateTutorial", mock.Anything, mock.Anything).Return(nil).Once()

		err := u.UpdateTutorial(context.Background(), req)

		assert.NilError(t, err)
	})
}
