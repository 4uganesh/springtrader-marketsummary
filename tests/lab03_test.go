package main

import (
	"fmt"
	"log"

	. "github.com/liatrio/springtrader/tests/validate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lab 3", func() {
	var failMessage string

	BeforeEach(func() {
		failMessage = ""
	})

	Context("Step 3", func() {
		It("should have a virtualService.yaml file", func() {
			failMessage = "virtualService.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/virtualService.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 4", func() {
		It("should have a hpa.yaml file", func() {
			failMessage = "hpa.yaml doesn't exist or is in the wrong location\n"
			Expect("../charts/marketsummary/templates/hpa.yaml").To(BeAnExistingFile(), failMessage)
		})
	})

	Context("Step 5", func() {
		It("should have a skaffold.yaml file", func() {
			failMessage = "skaffold.yaml doesn't exist or is in the wrong location\n"
			Expect("../skaffold.yaml").To(BeAnExistingFile(), failMessage)
		})
		It("should have a valid skaffold.yaml", func() {
			skaffoldExpected := ExpectYamlToParse("../skaffold.yaml")
			skaffoldActual := ExpectYamlToParse("./validate/solution-data/lab03/step05-skaffold.yaml")
			_, err := ValidateYamlObject(skaffoldExpected, &failMessage).Match(skaffoldActual)
			if err != nil {
				failMessage = fmt.Sprintf("skaffold.yaml has incorrect configuration; %s\n", err.Error())
			}
			Expect(err).To(BeNil())
		})
	})

	AfterEach(func() {
		log.Printf("%v\n", CurrentGinkgoTestDescription())
		if CurrentGinkgoTestDescription().Failed {
			ConcatenatedMessage += failMessage
		}
	})
})
