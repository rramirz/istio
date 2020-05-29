// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vm

import (
	"testing"
)

func TestInitializeConfig(t *testing.T) {
	config := NewConfig()

	if config.CAClientConfig.CSRInitialRetrialInterval != defaultCSRInitialRetrialInterval {
		t.Errorf("Unexpected config.CSRInitialRetrialInterval: %v", config.CAClientConfig.CSRInitialRetrialInterval)
	}

	if config.CAClientConfig.CSRMaxRetries != defaultCSRMaxRetries {
		t.Errorf("Unexpected config.CSRMaxRetries: %v", config.CAClientConfig.CSRMaxRetries)
	}

	if config.CAClientConfig.CSRGracePeriodPercentage != defaultCSRGracePeriodPercentage {
		t.Errorf("Unexpected config.CSRGracePeriodPercentage: %v", config.CAClientConfig.CSRGracePeriodPercentage)
	}

}
