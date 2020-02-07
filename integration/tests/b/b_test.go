// +build integration

package b_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("BEFORE: b")
	fmt.Println(time.Now())
})

var _ = Describe("b", func() {
	It("3=3", func() {
		fmt.Fprintln(GinkgoWriter, ">>> B 3")
		fmt.Println("b: 3=3")
		time.Sleep(5 * time.Second)
		Expect(3).To(Equal(3))
	})

	It("4=4", func() {
		fmt.Fprintln(GinkgoWriter, ">>> B 4")
		fmt.Println("b: 4=4")
		time.Sleep(10 * time.Second)
		Expect(4).To(Equal(4))
	})
})

var _ = AfterSuite(func() {
	fmt.Println("AFTER: b")
})
