package form

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/y3kawaguchi/quizen/internal/domains"
	"github.com/y3kawaguchi/quizen/test/testutils"
)

func TestQuiz(t *testing.T) {

	cases := map[string]struct {
		body       []byte
		wantErr    string
		wantDomain *domains.Quiz
	}{
		"validation check ok": {
			body:       testutils.GetBytesFromFile("testdata/quiz_req.json"),
			wantErr:    "",
			wantDomain: func() *domains.Quiz { return buildDefaultQuiz() }(),
		},
	}

	gin.SetMode(gin.ReleaseMode)
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			gc, _ := gin.CreateTestContext(httptest.NewRecorder())
			gc.Request, _ = http.NewRequest("POST", "/quizzes", bytes.NewBuffer(tc.body))
			quiz := Quiz{}
			err := gc.ShouldBindJSON(&quiz)
			if err == nil {
				if diff := cmp.Diff(tc.wantErr, ""); diff != "" {
					t.Errorf("%s: failed (-want +got):\n%s", k, diff)
				}
				if tc.wantDomain != nil {
					if diff := cmp.Diff(tc.wantDomain, quiz.BuildDomain()); diff != "" {
						t.Errorf("%s: failed (-want +got):\n%s", k, diff)
					}
				}
			} else {
				if diff := cmp.Diff(tc.wantErr, fmt.Sprintf("%v", err.Error())); diff != "" {
					t.Errorf("%s: failed (-want +got):\n%s", k, diff)
				}
			}
		})
	}
}

func buildDefaultQuiz() *domains.Quiz {
	return &domains.Quiz{
		Title:    "TestTitle",
		Question: "TestQuestion",
		Answer:   "TestAnswer",
	}
}
