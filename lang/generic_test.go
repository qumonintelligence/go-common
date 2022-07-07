package lang_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/qumonintelligence/go-common/v2/lang"
	"golang.org/x/exp/maps"
)

var _ = Describe("Generic function test", func() {
	var (
		intSlice    []int
		int64Slice  []int64
		stringSlice []string
	)

	BeforeEach(func() {
		intSlice = []int{1, 2, 3, 4, 5}
		int64Slice = []int64{1, 2, 3, 4, 5}
		stringSlice = []string{"hello", "world"}
	})

	Describe("Is Slice Contain", func() {
		Context("Provide value in the int slice", func() {
			It("Should be found the element in the slice and return true", func() {
				Expect(lang.IsSliceContain(intSlice, 1)).To(Equal(true))
			})
		})

		Context("Provide value not in the int slice", func() {
			It("Should not be found the element in the slice and return false", func() {
				Expect(lang.IsSliceContain(intSlice, 6)).To(Equal(false))
			})
		})

		Context("Provide value in the int64 slice", func() {
			It("Should be found the element in the slice and return true", func() {
				Expect(lang.IsSliceContain(int64Slice, 1)).To(Equal(true))
			})
		})

		Context("Provide value not in the int64 slice", func() {
			It("Should not be found the element in the slice and return false", func() {
				Expect(lang.IsSliceContain(int64Slice, 6)).To(Equal(false))
			})
		})

		Context("Provide value in the string slice", func() {
			It("Should be found the element in the slice and return true", func() {
				Expect(lang.IsSliceContain(stringSlice, "hello")).To(Equal(true))
			})
		})

		Context("Provide value not in the string slice", func() {
			It("Should not be found the element in the slice and return false", func() {
				Expect(lang.IsSliceContain(stringSlice, "notfound")).To(Equal(false))
			})
		})
	})

	Describe("Convert Slice To Map", func() {
		Context("Provide int slice and a boolean value", func() {
			It("Should be return a map[int]bool{}", func() {
				result := lang.ConvertSliceToMap(intSlice, true)
				expected := map[int]bool{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				}
				Expect(maps.Equal(result, expected)).To(Equal(true))
			})
		})

		Context("Provide int64 slice and a boolean value", func() {
			It("Should be return a map[int64]bool{}", func() {
				result := lang.ConvertSliceToMap(int64Slice, true)
				expected := map[int64]bool{
					1: true,
					2: true,
					3: true,
					4: true,
					5: true,
				}
				Expect(maps.Equal(result, expected)).To(Equal(true))
			})
		})

		Context("Provide string slice and a boolean value", func() {
			It("Should be return a map[string]bool{}", func() {
				result := lang.ConvertSliceToMap(stringSlice, true)
				expected := map[string]bool{
					"hello": true,
					"world": true,
				}
				Expect(maps.Equal(result, expected)).To(Equal(true))
			})
		})
	})

})
