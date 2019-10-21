package pages

import (
	"github.com/tebeka/selenium"
)

type Page struct {
	driver selenium.WebDriver
	URL    string
}

func (s *Page) GoTo() {
	s.driver.Get(s.URL)
}

func (s *Page) IsAt() bool {
	current, _ := s.driver.CurrentURL()
	return current == s.URL
}

func (s *Page) FindElementById(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByID, locator)
	return element
}

func (s *Page) FindElementByXpath(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByXPATH, locator)
	return element
}

func (s *Page) FindElementByLinkText(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByLinkText, locator)
	return element
}

func (s *Page) FindElementByPartialLink(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByPartialLinkText, locator)
	return element
}

func (s *Page) FindElementByName(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByName, locator)
	return element
}

func (s *Page) FindElementByTag(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByTagName, locator)
	return element
}

func (s *Page) FindElementByClass(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByClassName, locator)
	return element
}

func (s *Page) FindElementByCss(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByCSSSelector, locator)
	return element
}

func (s *Page) FindElementsById(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByID, locator)
	return element
}

func (s *Page) FindElementsByXpath(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByXPATH, locator)
	return element
}

func (s *Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByLinkText, locator)
	return element
}

func (s *Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByPartialLinkText, locator)
	return element
}

func (s *Page) FindElementsByName(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByName, locator)
	return element
}

func (s *Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByTagName, locator)
	return element
}

func (s *Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByClassName, locator)
	return element
}

func (s *Page) FindElementsByCss(locator string) []selenium.WebElement {
	element, _ := s.driver.FindElements(selenium.ByCSSSelector, locator)
	return element
}

func (s *Page) MouseHoverToElement(locator string) selenium.WebElement {
	element, _ := s.driver.FindElement(selenium.ByCSSSelector, locator)
	element.MoveTo(0, 0)
	return element
}
