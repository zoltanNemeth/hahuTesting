package pages

import (
	"github.com/tebeka/selenium"
	"github.com/zoltanNemeth/hahuTesting/driver"
	"strconv"
	"time"
)

type homePage struct {
	Page Page
}

var (
	hp                        *homePage
	privacySettingsOkButtonId = "CybotCookiebotDialogBodyButtonAccept"
	searchButton              = "//button[@name='submitKereses']"
	resultsLimitingSelectId   = "hirdetesszemelyautosearch-results"
	menuToggleXpath           = "//nav[contains(@class, 'navbar')]//button[contains(@class, 'navbar-toggle collapsed')]"
	loginButtonLinkText       = "Belépés"
	usernameInputFieldId      = "loginform-felhasznalo"
	passwordInputFieldId      = "loginform-jelszo"
	errorMessageXpath         = "//p[contains(text(), 'Hibás felhasználónév vagy jelszó')]"
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
	h.Page.FindElementById(resultsLimitingSelectId).Click()
	stringLimit := strconv.Itoa(limit)
	desiredLimitOptionXpath := "//select[@id='" + resultsLimitingSelectId + "']" + "//option[@value='" + stringLimit + "']"
	h.Page.FindElementByXpath(desiredLimitOptionXpath).Click()
}

func (h *homePage) AcceptPrivacyInformation() {
	h.Page.FindElementById(privacySettingsOkButtonId).Click()
}

func (h *homePage) ClickSearchButton() ResultsPage {
	h.Page.FindElementByXpath(searchButton).Click()
	currentURL, _ := driver.Driver().CurrentURL()
	return NewResultsPage(currentURL)
}

func (h *homePage) ClickLoginButton() {
	h.Page.FindElementByXpath(menuToggleXpath).Click()
	h.Page.FindElementByLinkText(loginButtonLinkText).Click()
}

func (h *homePage) LoginWith(username, password string) {
	h.Page.FindElementById(usernameInputFieldId).SendKeys(username)
	h.Page.FindElementById(passwordInputFieldId).SendKeys(password)
	h.Page.FindElementByLinkText(loginButtonLinkText)
}

func (h *homePage) IsThereAnErrorMessage() bool {
	condition := selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
		if h.Page.FindElementByXpath(errorMessageXpath) == nil {
			return false, nil
		}
		return true, nil
	})
	err := driver.Driver().WaitWithTimeout(condition, time.Duration(time.Second*10))
	if err != nil {
		return false
	}
	if h.Page.FindElementByXpath(errorMessageXpath) == nil {
		return false
	}
	return true
}
