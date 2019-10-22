package pages

import (
	"github.com/tebeka/selenium"
	"github.com/zoltanNemeth/hahuTesting/driver"
	"time"
)

type ResultsPage struct {
	page Page
}

var (
	resultsXpath = "//div[contains(@class, 'talalatisor-tartalom')]"
)

func NewResultsPage(URL string) ResultsPage {
	return ResultsPage{
		Page{
			driver.Driver(),
			URL,
		},
	}
}

func (r *ResultsPage) GetNumOfResults() int {
	condition := selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
		if len(r.page.FindElementsByXpath(resultsXpath)) < 1 {
			return false, nil
		}
		return true, nil
	})
	err := driver.Driver().WaitWithTimeoutAndInterval(condition, time.Duration(time.Second*100), time.Duration(time.Second*3))
	if err != nil {
		return 0
	}
	return len(r.page.FindElementsByXpath(resultsXpath))
}
