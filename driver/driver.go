package driver

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
	"net"
	"os"
	"strconv"
)

type driver struct {
	service   selenium.Service
	webdriver selenium.WebDriver
}

//type simpleWebDriver struct {
//	webdriver selenium.WebDriver
//}

var (
	d *driver
	//webDriver *simpleWebDriver
)

//func SimpleDriver() selenium.WebDriver  {
//	caps := selenium.Capabilities{"browserName": "chrome"}
//	driver, _ := selenium.NewRemote(caps, "")
//	webDriver = &simpleWebDriver{
//		webdriver: driver,
//	}
//	return webDriver.webdriver
//}

func Driver() selenium.WebDriver {
	return d.webdriver
}

func Service() selenium.Service {
	return d.service
}

func init() {
	if d == nil {
		port, err := pickUnusedPort()

		_ = godotenv.Load()
		var opts []selenium.ServiceOption
		s, err := selenium.NewChromeDriverService(os.Getenv("webdriverPath"), port, opts...)

		if err != nil {
			fmt.Printf("Error starting the ChromeDriver server: %v", err)
		}

		caps := selenium.Capabilities{
			"browserName": "chrome",
			"args":        "--no-sandbox",
		}

		wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")
		if err != nil {
			panic(err)
		}

		d = &driver{
			service:   *s,
			webdriver: wd,
		}
	}
}

func pickUnusedPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0, err
	}
	return port, nil
}
