package pages

import (
	"github.com/zoltanNemeth/hahuTesting/driver"
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
			Page: Page{
				driver: driver.Driver(),
				URL:    "https://www.hasznaltauto.hu/",
			},
		}
	}
	return hp
}

func (h *homePage) LimitResultsTo(limit int) {
	h.Page.FindElementById(resultsLimitingSelectId).Click()
}

func (h *homePage) ClickSearchButton() {
	h.Page.FindElementByXpath(searchButton).Click()
}
