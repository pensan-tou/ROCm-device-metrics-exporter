/**
# Copyright (c) Advanced Micro Devices, Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the \"License\");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an \"AS IS\" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package k8e2e

import (
	"github.com/ROCm/device-metrics-exporter/test/k8s-e2e/clients"
	restclient "k8s.io/client-go/rest"
)

// E2ESuite e2e config
type E2ESuite struct {
	k8sclient  *clients.K8sClient
	helmClient *clients.HelmClient
	restConfig *restclient.Config
	registry   string
	helmChart  string
	imageTag   string
	ns         string
	kubeconfig string
	platform   string
}
