package cluster_sanity_check

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jianzhangbjz/openshift-tests/utils"
)

func TestGinkoDemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cluster Sanity Check")
}

func getNodes() {
	utils.CheckResult("oc get nodes -n default", "NAME")
}

func getProjects() {
	utils.CheckResult("oc get projects", "default")
}

var _ = Describe("Cluster sanity check", func() {

	Context("Scenario: I would like to see if oc can connect to the cluster and extract basic informtion", func() {
		It("I would like to verify the functionality of oc get nodes", func() {
			getNodes()
		})
		It("I would like to see if oc get projects works", func() {
			getProjects()
		})

	})

})

