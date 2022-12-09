package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/pkg"
)

var _ = Describe("Config", func() {
	It("returns the blank config", func() {
		Run("config")
		Eventually(CommandSession).Should(Exit(0))
		output := CommandSession.Out.Contents()
		var blankConfig pkg.Config
		Expect(yaml.Unmarshal(output, &blankConfig)).To(Succeed())
		Expect(blankConfig.APIKey).To(BeEmpty())
		Expect(blankConfig.Date).To(BeEmpty())
	})
})
