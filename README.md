# hahuTesting

## About the project

This is a project to practice test automation using Go, Gherkin and Selenium.

I started this project (and learning Go) on the 17th of October, 2019.
System under test is hasznaltauto.hu, hahu in short.
I hope this will help me to get an offer at the AC of Adevinta on Friday.

I tried out Travis CI but yet I couldn't set chromeOptions in driver
to be able to run the tests,
so check another simple project at https://github.com/zoltanNemeth/travis/
that runs fine at Travis.

Dockerfile is coming soon...

## Running tests in hahuTesting

### Preconditions:
- Go 1.11 or newer version should be added to PATH environmental variable in your operating system
- The feature files should be in the features directory
- Java should be installed
- Chrome browser should be installed
- Chromedriver should be downloaded (proper version for your Chrome browser)
- .env text file should be created in the root directory of the project
  and it should contain webdriverPath=\<path to your webdriver\>

#### Using command: `go test`

Precondition:
- run go mod vendor before the first go test execution.

## Adding new tests

Add new features, scenarios then run `go test`.
Godog will print to the console the instructions.
You should call *(s *godog.Suite) {...} method(s)
in godog.RunWithOptions("godogs", func(s *godog.Suite) {...} in godog_test.go






