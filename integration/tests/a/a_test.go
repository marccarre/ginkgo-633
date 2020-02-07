// +build integration

package a_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("BEFORE: a")
	fmt.Println(time.Now())
})

var _ = Describe("a", func() {
	It("1=1", func() {
		if GinkgoWriter == nil {
			fmt.Println("it is fucking nil")
		}
		fmt.Fprintln(GinkgoWriter, ">>> A 1")
		fmt.Println("a: 1=1")
		time.Sleep(5 * time.Second)
		Expect(1).To(Equal(1))
	})

	It("2=2", func() {
		fmt.Fprintln(GinkgoWriter, ">>> A 2")
		fmt.Println("a: 2=2")
		time.Sleep(10 * time.Second)
		Expect(2).To(Equal(2))
	})
})

var _ = AfterSuite(func() {
	fmt.Println("AFTER: a")
})
