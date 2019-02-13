package opinion

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func FindOpinion(url string) (Opinion, error) {
	resp, httpErr := http.Get(url)

	if httpErr != nil {
		return Opinion{}, httpErr
	}

	document, gqErr := goquery.NewDocumentFromReader(io.Reader(resp.Body))

	if gqErr != nil {
		return Opinion{}, gqErr
	}

	yesPercentage, err := strconv.Atoi(strings.Replace(strings.Fields(document.Find(".yes-text").
		Text())[0], "%", "", -1))
	noPercentage, err := strconv.Atoi(strings.Replace(strings.Fields(document.Find(".no-text").
		Text())[0], "%", "", -1))

	name := document.Find(".q-title").Text()

	var arguments []Argument

	document.Find(".arguments").Each(func(_ int, argument *goquery.Selection) {
		argument.Find("ul").Each(func(_ int, post *goquery.Selection) {
			post.Find("li").Each(func(_ int, li *goquery.Selection) {
				text := li.Find("p").Text()
				author := li.Find("cite").Find("a").Text()
				if len(text) > 0 && len(author) > 0 {
					arguments = append(arguments, Argument{
						Text:   text,
						Author: author,
					})
				}
			})
		})
	})

	if err != nil {
		return Opinion{}, err
	}

	return Opinion{
		Name:            name,
		PercentageOfYes: yesPercentage,
		PercentageOfNo:  noPercentage,
		Arguments:       arguments,
	}, nil

}
