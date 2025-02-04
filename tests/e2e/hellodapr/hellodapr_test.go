//go:build e2e
// +build e2e

/*
Copyright 2021 The Dapr Authors
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

package hellodapr_e2e

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dapr/dapr/pkg/json"
	"github.com/dapr/dapr/tests/e2e/utils"
	kube "github.com/dapr/dapr/tests/platforms/kubernetes"
	"github.com/dapr/dapr/tests/runner"
)

type testCommandRequest struct {
	Message string `json:"message,omitempty"`
}

type appResponse struct {
	Message   string `json:"message,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
}

const numHealthChecks = 60 // Number of times to check for endpoint health per app.

var tr *runner.TestRunner

func TestMain(m *testing.M) {
	// This test shows how to deploy the multiple test apps, validate the side-car injection
	// and validate the response by using test app's service endpoint

	// These apps will be deployed for hellodapr test before starting actual test
	// and will be cleaned up after all tests are finished automatically
	testApps := []kube.AppDescription{
		{
			AppName:           "hellobluedapr",
			DaprEnabled:       true,
			ImageName:         "e2e-hellodapr",
			Replicas:          1,
			IngressEnabled:    true,
			MetricsEnabled:    true,
			DaprMemoryLimit:   "200Mi",
			DaprMemoryRequest: "100Mi",
			AppMemoryLimit:    "200Mi",
			AppMemoryRequest:  "100Mi",
		},
		{
			AppName:           "hellogreendapr",
			DaprEnabled:       true,
			ImageName:         "e2e-hellodapr",
			Replicas:          1,
			IngressEnabled:    true,
			MetricsEnabled:    true,
			DaprMemoryLimit:   "200Mi",
			DaprMemoryRequest: "100Mi",
			AppMemoryLimit:    "200Mi",
			AppMemoryRequest:  "100Mi",
		},
		{
			AppName:           "helloenvtestdapr",
			DaprEnabled:       true,
			ImageName:         "e2e-hellodapr",
			Replicas:          1,
			IngressEnabled:    true,
			MetricsEnabled:    true,
			DaprMemoryLimit:   "200Mi",
			DaprMemoryRequest: "100Mi",
			AppMemoryLimit:    "200Mi",
			AppMemoryRequest:  "100Mi",
		},
	}

	tr = runner.NewTestRunner("hellodapr", testApps, nil, nil)
	os.Exit(tr.Start(m))
}

var helloAppTests = []struct {
	in               string
	app              string
	testCommand      string
	expectedResponse string
}{
	{
		"green dapr",
		"hellogreendapr",
		"green",
		"Hello green dapr!",
	},
	{
		"blue dapr",
		"hellobluedapr",
		"blue",
		"Hello blue dapr!",
	},
	{
		"envTest dapr",
		"helloenvtestdapr",
		"envTest",
		"3500 50001",
	},
}

func TestHelloDapr(t *testing.T) {
	for _, tt := range helloAppTests {
		t.Run(tt.in, func(t *testing.T) {
			// Get the ingress external url of test app
			externalURL := tr.Platform.AcquireAppExternalURL(tt.app)
			require.NotEmpty(t, externalURL, "external URL must not be empty")

			// Check if test app endpoint is available
			_, err := utils.HTTPGetNTimes(externalURL, numHealthChecks)
			require.NoError(t, err)

			// Trigger test
			body, err := json.Marshal(testCommandRequest{
				Message: "Hello Dapr.",
			})
			require.NoError(t, err)

			resp, err := utils.HTTPPost(fmt.Sprintf("%s/tests/%s", externalURL, tt.testCommand), body)
			require.NoError(t, err)

			var appResp appResponse
			err = json.Unmarshal(resp, &appResp)
			require.NoError(t, err)
			require.Equal(t, tt.expectedResponse, appResp.Message)
		})
	}
}

func TestScaleReplicas(t *testing.T) {
	err := tr.Platform.Scale("hellobluedapr", 3)
	require.NoError(t, err, "fails to scale hellobluedapr app to 3 replicas")
}

func TestScaleAndRestartInstances(t *testing.T) {
	err := tr.Platform.Scale("hellobluedapr", 3)
	require.NoError(t, err, "fails to scale hellobluedapr app to 3 replicas")

	err = tr.Platform.Restart("hellobluedapr")
	require.NoError(t, err, "fails to restart hellobluedapr pods")
}
