package hahuTesting

import (
	"flag"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/zoltanNemeth/hahuTesting/features/step_definitions"
	"os"
	"testing"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		step_definitions.CommonSteps(s)
		step_definitions.AuthenticationFeature(s)
		step_definitions.SearchingFeature(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
