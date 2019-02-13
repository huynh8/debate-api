package opinion

import (
	"encoding/json"
	"net/http"
	"strings"
)

const ENDPOINT = "/opinion"

type Opinion struct {
	Name string `json:"name"`
	PercentageOfYes int `json:"percentage_of_yes"`
	PercentageOfNo int `json:"percentage_of_no"`
	Arguments []Argument `json:"arguments"`
}

type Argument struct {
	Text string `json:"text"`
	Author string `json:"author"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	urlParams, ok := r.URL.Query()["url"]
	if !ok  ||
		!strings.HasPrefix(urlParams[0], "https://www.debate.org/opinions/") {
		http.Error(w, "Debate url is missing or not correct", 400)
		return
	}

	o, oerr := FindOpinion(urlParams[0])

	if oerr != nil {
		http.Error(w, oerr.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(o)

	if err != nil {
		http.Error(w, "Response was unable to be parsed into json", 500)
		return
	}
}
