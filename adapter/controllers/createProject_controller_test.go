package controller

import (
	"bytes"
	"context"
	"fmt"
	"marketplace/usecases"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockProjectCreateProject struct {
	result usecases.ProjectOutput
	err    error
}

func (m mockProjectCreateProject) Execute(c context.Context, uc usecases.ProjectInput) (usecases.ProjectOutput, error) {
	return m.result, m.err
}

func TestCreateProjectController_Execute(t *testing.T) {
	t.Parallel()

	type args struct {
		payload []byte
	}

	tests := []struct {
		name               string
		args               args
		ucMock             usecases.CreateProjectUseCase
		expectedBody       string
		expectedStatusCode int
	}{
		{
			name: "CreateProjectController success",
			args: args{
				payload: []byte(
					`{ 
						"name": "project1",
				 		"budget": 500000,
						"driverImputation": [{
							"ceco": "203100010",
							"cia": "20",
							"percentage": 100 }]
					 }`,
				),
			},
			ucMock: mockProjectCreateProject{
				result: usecases.ProjectOutput{
					Id:     "786546776",
					Name:   "project1",
					Budget: 500000,
					DriverImputation: []usecases.ImputationUnityInOutput{
						usecases.ImputationUnityInOutput{
							Ceco:       "203100010",
							Cia:        "20",
							Percentage: 100,
						},
					},
				},
				err: nil,
			},
			expectedBody:       `{"id":"786546776","name":"project1","budget":500000,"driverImputation":[{"ceco":"203100010","cia":"20","percentage":100}]}`,
			expectedStatusCode: http.StatusOK,
		},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				"/projects",
				bytes.NewReader(tt.args.payload))
			var (
				w          = httptest.NewRecorder()
				controller = NewCreateProjectAction(tt.ucMock)
			)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			controller.Execute(c)

			if w.Code != tt.expectedStatusCode {
				t.Errorf(
					"[TestCase '%s'] invalid response: returned '%v' expected value '%v'",
					tt.name,
					w.Code,
					tt.expectedStatusCode,
				)
			}

			var result = strings.TrimSpace(w.Body.String())
			fmt.Println(result)
			fmt.Println(tt.expectedBody)
			if !strings.EqualFold(result, tt.expectedBody) {
				t.Errorf(
					"[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					result,
					tt.expectedBody,
				)
			}
		})
	}
}
