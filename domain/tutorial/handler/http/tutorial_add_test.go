package http

import (
	"encoding/json"
	"errors"
	"gocrudssample/domain/tutorial/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddTutorial(t *testing.T) {
	type input struct {
		req map[string]interface{}
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
			name: "#1 success add data",
			expectedInput: input{
				req: map[string]interface{}{
					"tutorialTypeId": "93d899ef-b918-4a94-b7fb-c51df7c7e144",
					"title":          "title tutorial",
					"sequence":       1,
					"keywords":       "php",
					"description":    "desc",
				},
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {

				mockTutorial.
					On("AddTutorial", mock.Anything, mock.Anything).
					Return(nil)
			},
		},
		{
			name: "#2 internal server error",
			expectedInput: input{
				req: map[string]interface{}{
					"tutorialTypeId": "93d899ef-b918-4a94-b7fb-c51df7c7e144",
					"title":          "title tutorial",
					"sequence":       1,
					"keywords":       "php",
					"description":    "desc",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockTutorial *mocks.TutorialUsecaseInterface,
			) {

				mockTutorial.
					On("AddTutorial", mock.Anything, mock.Anything).
					Return(errors.New("internal server error"))
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockTutorial := new(mocks.TutorialUsecaseInterface)

			payload, err := json.Marshal(testCase.expectedInput.req)

			assert.NoError(t, err)

			e := echo.New()

			req, err := http.NewRequest(echo.POST, "/tutorials",
				strings.NewReader(string(payload)))

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

			err = handler.Addtutorial(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}