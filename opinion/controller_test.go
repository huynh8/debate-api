package opinion

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/opinion?url=https://www.debate.org/opinions/should-drug-users-be-put-in-prison", nil)

	recorder := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(Handler)

	handlerFunc.ServeHTTP(recorder, request)

	assert.Equal(t, recorder.Code, http.StatusOK, "HTTP request responds with OK")

	var opinion Opinion
	_ = json.Unmarshal(recorder.Body.Bytes(), &opinion)

	assert.Equal(t, opinion.Name, "Should drug users be put in prison?", "Title matches")
	assert.Equal(t, opinion.PercentageOfYes, 50, "Percentage of Yes matches")
	assert.Equal(t, opinion.PercentageOfNo, 50, "Percentage of No matches")

	assert.Equal(t, len(opinion.Arguments), 7, "Number of arguments match")

	assert.Equal(t, opinion.Arguments[0].Author, "Adalman", "Authors match")
	assert.NotEmpty(t, opinion.Arguments[0].Text, "Post is not empty")

}
