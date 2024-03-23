package http

import (
	"errors"
	"gocrudsample/domain/tutorial/mocks"
	"gocrudsample/schema/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTutorialType(t *testing.T) {
	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name string

		expectedOutput output
		configureMock  func(
			mockTutorial *mocks.TutorialUsecaseInterface,
		)
	}{
		{
			name:           "#1 success get data",
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {
				tResponse := []response.TutorialTypes{}

				mockTutorial.
					On("GetTutorialTypes", mock.Anything).
					Return(tResponse, nil)
			},
		},
		{
			name:           "#2 internal server error",
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {
				tResponse := []response.TutorialTypes{}

				mockTutorial.
					On("GetTutorialTypes", mock.Anything).
					Return(tResponse, errors.New("internal server error"))
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockTutorial := new(mocks.TutorialUsecaseInterface)

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/tutorials/types", nil)

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			testCase.configureMock(
				mockTutorial,
			)

			handler := TutorialHandler{
				tutorialUsecase: mockTutorial,
			}

			err = handler.GetTutorialTypes(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
