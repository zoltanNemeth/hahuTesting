Running tests in hahuTesting

Preconditions:
    The feature files should be in the features directory, 
    Java should be installed, 
    Chrome browser should be installed, 
    Chromedriver should be downloaded (proper version for your Chrome browser)
    .env text file should be created in the root directory of the project
        and it should contain webdriverPath=<path to your webdriver>

Use: go run $GOPATH/src/github.com/DATA-DOG/godog/cmd/godog in the root directory of the project.
Preconditions: run go get github.com/DATA-DOG/godog and delete the go.mod and go.sum files. 
In this case, you should have the func *(s *godog.Suite) {...} function(s) in  *_test.go files
    in the root directory of your project.

Use: go test to run godog tests and other tests together.
Precondition: run go mod vendor before the first go test execution. 
In this case you should call func *(s *godog.Suite) {...} function(s)
    in godog.RunWithOptions("godogs", func(s *godog.Suite) {...} function in godog_test.go






