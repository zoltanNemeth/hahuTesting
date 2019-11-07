package driver

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"net"
	"os"
	"strconv"
)

type driver struct {
	service   selenium.Service
	webdriver selenium.WebDriver
}

var d *driver

func Driver() selenium.WebDriver {
	return d.webdriver
}

func Service() selenium.Service {
	return d.service
}

func init() {
	if d == nil {
		port, err := pickUnusedPort()

		err = godotenv.Load()
		if err != nil {
			fmt.Printf("Error loading the .env file: %v\n", err)
		}

		var opts []selenium.ServiceOption
		s, err := selenium.NewChromeDriverService(os.Getenv("WEBDRIVER_PATH"), port, opts...)

		if err != nil {
			fmt.Printf("Error starting the ChromeDriver server: %v\n", err)
		}

		caps := selenium.Capabilities{
			"browserName": "chrome",
		}

		args := []string{
			"start-maximized",
			"no-sandbox",
			"no-default-browser-check",
			"no-first-run",
			"disable-default-apps",
		}

		caps.AddChrome(chrome.Capabilities{Args: args})

		wd, err := selenium.NewRemote(caps, "http://127.0.0.1:"+strconv.Itoa(port)+"/wd/hub")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
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
