package pages

import (
	"github.com/zoltanNemeth/hahuTesting/driver"
	"strconv"
)

type homePage struct {
	Page Page
}

var (
	hp                      *homePage
	searchButton            = "//button[@name='submitKereses']"
	resultsLimitingSelectId = "hirdetesszemelyautosearch-results"
)

func HomePage() *homePage {
	if hp == nil {
		hp = &homePage{
			Page{
				driver.Driver(),
				"https://www.hasznaltauto.hu/",
			},
		}
	}
	return hp
}

func (h *homePage) LimitResultsTo(limit int) {
	_ = h.Page.FindElementById(resultsLimitingSelectId).Click()
	stringLimit := strconv.Itoa(limit)
	desiredLimitOptionXpath := "//select[@id='" + resultsLimitingSelectId + "']" + "//option[@value='" + stringLimit + "']"
	_ = h.Page.FindElementByXpath(desiredLimitOptionXpath).Click()
}

func (h *homePage) ClickSearchButton() ResultsPage {
	_ = h.Page.FindElementByXpath(searchButton).Click()
	currentURL, _ := driver.Driver().CurrentURL()
	return NewResultsPage(currentURL)
}
