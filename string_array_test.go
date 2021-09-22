package go_common_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/qumonintelligence/go-common/v2/lang"
)

var _ = Describe("StringArray", func() {
	var (
		strList []string
		strVar1 string
		strVar2 string
	)

	BeforeEach(func() {
		strList = []string{
			"test1",
			"test2",
			"test3",
		}

		strVar1 = "test2"
		strVar2 = "test4"
	})

	Describe("Calling StringSliceContains() function", func() {
		Context("While passing value contained in string slice", func() {
			It("Should return true", func() {
				Expect(lang.StringSliceContains(strList, strVar1)).To(Equal(true))
			})
		})

		Context("While passing value not contained in string slice", func() {
			It("Should return false", func() {
				Expect(lang.StringSliceContains(strList, strVar2)).To(Equal(false))
			})
		})
	})
})
