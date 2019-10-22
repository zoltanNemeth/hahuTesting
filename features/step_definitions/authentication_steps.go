package step_definitions

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/pages"
)

func theUserAcceptsThePrivacyInformation() error {
	pages.HomePage().AcceptPrivacyInformation()
	return nil
}

func theUserChoosesTheLoginMenu() error {
	pages.HomePage().ClickLoginButton()
	return nil
}

func ensureToProvideALoginFormWithUsernameAndPasswordField() error {
	return nil
}

func theUserSubmitsTheFormWithInvalidCredentials(username, password string) error {
	pages.HomePage().LoginWith(username, password)
	return nil
}

func anErrorMessageAppears() error {
	if pages.HomePage().IsThereAnErrorMessage() == true {
		return nil
	}
	return fmt.Errorf("there is not an error message")
}

func AuthenticationFeature(s *godog.Suite) {
	s.Step(`^the user accepts the privacy information$`, theUserAcceptsThePrivacyInformation)
	s.Step(`^the user chooses the Login menu$`, theUserChoosesTheLoginMenu)
	s.Step(`^ensure to provide a login form with username and password field$`, ensureToProvideALoginFormWithUsernameAndPasswordField)
	s.Step(`^the user submits the form with invalid credentials \("([^"]*)", "([^"]*)"\)$`, theUserSubmitsTheFormWithInvalidCredentials)
	s.Step(`^an error message appears\.$`, anErrorMessageAppears)
}
