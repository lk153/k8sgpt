/*
Copyright 2023 The K8sGPT Authors.
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

package integration

import (
	"github.com/spf13/cobra"
)

var (
	namespace string
)

// IntegrationCmd represents the integrate command
var IntegrationCmd = &cobra.Command{
	Use:     "integration",
	Aliases: []string{"integrations"},
	Short:   "Intergrate another tool into K8sGPT",
	Long: `Intergrate another tool into K8sGPT. For example:
	
	k8sgpt integration activate trivy
	
	This would allow you to deploy trivy into your cluster and use a K8sGPT analyzer to parse trivy results.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	IntegrationCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "The namespace to use for the integration")
}
