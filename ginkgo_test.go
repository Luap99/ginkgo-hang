package ginkgo_hang_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Test", func() {
	It("test hang", func() {
		cmd := exec.Command("./daemon")
		session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		session.Wait(5)
		Eventually(session, 5).Should(Exit(0))
	})
})
