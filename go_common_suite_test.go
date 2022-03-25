package go_common_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoCommon Suite")
}
