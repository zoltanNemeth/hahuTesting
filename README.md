# Running tests in hahuTesting

This is a project to practice using Go, Gherkin and Selenium together.

I started this project (and learning Go) on the 17th of October, 2019.
System under test is hasznaltauto.hu, hahu in short.
I hope this will help me to get an offer at the AC of Adevinta on Friday.

## Preconditions:
- The feature files should be in the features directory
- Java should be installed
- Chrome browser should be installed
- Chromedriver should be downloaded (proper version for your Chrome browser)
- .env text file should be created in the root directory of the project
  and it should contain webdriverPath=\<path to your webdriver\>

#### Using command: go run $GOPATH/src/github.com/DATA-DOG/godog/cmd/godog
Preconditions:
- run go get github.com/DATA-DOG/godog and delete the go.mod and go.sum files.

Run command in the root directory of the project.

In this case, you should have the func *(s *godog.Suite) {...} function(s) in  *_test.go files
    in the root directory of your project.

#### Using command: go test
Precondition:
- run go mod vendor before the first go test execution.

In this case you should call func *(s *godog.Suite) {...} function(s)
    in godog.RunWithOptions("godogs", func(s *godog.Suite) {...} function in godog_test.go






