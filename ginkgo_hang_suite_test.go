package ginkgo_hang_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoHang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoHang Suite")
}
