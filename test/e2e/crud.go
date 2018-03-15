/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"fmt"

	"github.com/marun/fnord/pkg/federatedtypes"
	"github.com/marun/fnord/test/common"
	"github.com/marun/fnord/test/e2e/framework"

	. "github.com/onsi/ginkgo"
)

func addE2ECrudTests() {
	Describe("Federated types", func() {
		f := framework.NewFederationFramework("federated-types")

		fedTypeConfigs := federatedtypes.FederatedTypeConfigs()
		for kind, _ := range fedTypeConfigs {
			// TODO (font): e2e tests for FederatedNamespace using a test managed
			// federation does not work until namespace controller is added.
			if framework.TestContext.TestManagedFederation &&
				kind == federatedtypes.FederatedNamespaceKind {
				continue
			}

			// Bind the type config inside the loop to ensure the ginkgo
			// closure gets a different value for every loop iteration.
			//
			// Reference: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
			fedTypeConfig := fedTypeConfigs[kind]

			Describe(fmt.Sprintf("%q resources", kind), func() {
				It("should be created, read, updated and deleted successfully", func() {
					// Initialize an in-memory controller if configuration requires
					f.SetUpControllerFixture(kind, fedTypeConfig.AdapterFactory)

					userAgent := fmt.Sprintf("crud-test-%s", kind)
					adapter := fedTypeConfig.AdapterFactory(f.FedClient(userAgent))
					clusterClients := f.ClusterClients(userAgent)
					crudTester := common.NewFederatedTypeCrudTester(framework.NewE2ELogger(), adapter, clusterClients, framework.PollInterval, framework.SingleCallTimeout)
					clusterNames := []string{}
					for name, _ := range clusterClients {
						clusterNames = append(clusterNames, name)
					}
					template, placement, override := federatedtypes.NewTestObjects(fedTypeConfig.Kind, f.TestNamespaceName(), clusterNames)

					crudTester.CheckLifecycle(template, placement, override)
				})
			})
		}
	})
}
