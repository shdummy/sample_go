package sample

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
	"testing"

	"github.com/onsi/ginkgo/reporters"
)

func TestSample(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
	RunSpecsWithDefaultAndCustomReporters(t, "Sample Suite", []Reporter{junitReporter})
}
