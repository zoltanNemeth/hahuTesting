package step_definitions

import (
	"github.com/DATA-DOG/godog"
)

func theUserChoosesTheLoginMenu() error {
	return godog.ErrPending
}

func ensureToProvideALoginFormWithUsernameAndPasswordField() error {
	return godog.ErrPending
}

func theUserSubmitsTheFormWithValidCredentials(arg1, arg2 string) error {
	return godog.ErrPending
}

func aLogoutOptionAppearsAndIsShown(arg1 string) error {
	return godog.ErrPending
}

func theUserSubmitsTheFormWithInvalidCredentials(arg1, arg2 string) error {
	return godog.ErrPending
}

func anErrorMessageAppears() error {
	return godog.ErrPending
}

func AuthenticationFeature(s *godog.Suite) {
	s.Step(`^the user chooses the Login menu$`, theUserChoosesTheLoginMenu)
	s.Step(`^ensure to provide a login form with username and password field$`, ensureToProvideALoginFormWithUsernameAndPasswordField)
	s.Step(`^the user submits the form with valid credentials "([^"]*)", "([^"]*)"$`, theUserSubmitsTheFormWithValidCredentials)
	s.Step(`^a Logout option appears and "([^"]*)" is shown\.$`, aLogoutOptionAppearsAndIsShown)
	s.Step(`^the user submits the form with invalid credentials \("([^"]*)", "([^"]*)"\)$`, theUserSubmitsTheFormWithInvalidCredentials)
	s.Step(`^an error message appears\.$`, anErrorMessageAppears)
}
