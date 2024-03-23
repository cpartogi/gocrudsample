package http

import (
	"errors"
	"gocrudssample/domain/tutorial/mocks"
	"gocrudssample/schema/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTutorials(t *testing.T) {
	type input struct {
		tutorialTypeId string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockTutorial *mocks.TutorialUsecaseInterface,
		)
	}{
		{
			name: "#1 success get data",
			expectedInput: input{
				tutorialTypeId: "abc",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {
				tResponse := []response.TutorialList{}

				mockTutorial.
					On("GetTutorials", mock.Anything, mock.Anything).
					Return(tResponse, nil)
			},
		},
		{
			name:           "#2 internal server error",
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {
				tResponse := []response.TutorialList{}

				mockTutorial.
					On("GetTutorials", mock.Anything, mock.Anything).
					Return(tResponse, errors.New("internal server error"))
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockTutorial := new(mocks.TutorialUsecaseInterface)

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/tutorials?tutorialTypeId=abc", nil)

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			testCase.configureMock(
				testCase.expectedInput,
				mockTutorial,
			)

			handler := TutorialHandler{
				tutorialUsecase: mockTutorial,
			}

			err = handler.GetTutorials(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
