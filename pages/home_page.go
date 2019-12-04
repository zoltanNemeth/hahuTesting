package pages

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/zoltanNemeth/hahuTesting/driver"
	"regexp"
	"strconv"
	"time"
)

type homePage struct {
	Page Page
}

var (
	hp                                *homePage
	brand                             selenium.WebElement
	privacySettingsOkButtonId         = "CybotCookiebotDialogBodyButtonAccept"
	searchButton                      = "//button[@name='submitKereses']"
	resultsLimitingSelectId           = "hirdetesszemelyautosearch-results"
	brandsWithPotentionalResultsXpath = "//select[@id='hirdetesszemelyautosearch-marka_id']/option"
	menuToggleXpath                   = "//nav[contains(@class, 'navbar')]//button[contains(@class, 'navbar-toggle collapsed')]"
	menuLoginButtonLinkText           = "Belépés"
	loginFormLoginButtonXpath         = "//button[contains(text(), 'Belépés')]"
	usernameInputFieldId              = "loginform-felhasznalo"
	passwordInputFieldId              = "loginform-jelszo"
	errorMessageXpath                 = "//p[contains(text(), 'Hibás felhasználónév vagy jelszó')]"
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
	condition := selenium.Condition(func(wd selenium.WebDriver) (bool, error) {
		if h.Page.FindElementByXpath(menuToggleXpath) == nil {
			return false, nil
		}
		return true, nil
	})
	driver.Driver().WaitWithTimeout(condition, time.Duration(time.Second*10))
	err := h.Page.FindElementByXpath(menuToggleXpath).Click()
	if err != nil {
		fmt.Errorf("error during clicking on menuToggle")
	}
	h.Page.FindElementByLinkText(menuLoginButtonLinkText).Click()
}

func (h *homePage) LoginWith(username, password string) {
	h.Page.FindElementById(usernameInputFieldId).SendKeys(username)
	h.Page.FindElementById(passwordInputFieldId).SendKeys(password)
	h.Page.FindElementByXpath(loginFormLoginButtonXpath).Click()
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

func (h *homePage) getBrandWithResults(from, to int) selenium.WebElement {
	brandsWithPotentionalResults := h.Page.FindElementsByXpath(brandsWithPotentionalResultsXpath)
	for _, brand := range brandsWithPotentionalResults {
		SumOfPotentialResults := h.getBrandPotentialResult(brand)
		if (SumOfPotentialResults > from) && (SumOfPotentialResults < to) {
			return brand
		}
	}
	return nil
}

func (h *homePage) getBrandPotentialResult(brand selenium.WebElement) int {
	b, _ := brand.Text()
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	matches := re.FindAllString(b, 2)
	if matches != nil {
		potentialResult, _ := strconv.Atoi(matches[0])
		if len(matches) >= 2 {
			part1 := matches[0]
			part2 := matches[1]
			num := part1 + part2
			potentialResult, _ = strconv.Atoi(num)
		}
		return potentialResult
	}
	return 0
}

func (h *homePage) SelectBrandWithResults(from, to int) {
	brand = h.getBrandWithResults(from, to)
	brand.Click()
}

func (h *homePage) GetSumOfResultsBesideBrand() int {
	return h.getBrandPotentialResult(brand)
}

func (h *homePage) GetSumOfResultsInSearchButton() int {
	return h.getBrandPotentialResult(h.Page.FindElementByXpath(searchButton))
}
