package unit

import (
	"fmt"
	"os/exec"
	"sort"
	"sync/atomic"
	"testing"

	"sigs.k8s.io/yaml"
)

func renderTemplate(t *testing.T, chartPath, releaseName, namespace, kubeVersion string, values map[string]string, template string) string {
	t.Helper()

	args := []string{
		"template", releaseName, chartPath,
		"--namespace", namespace,
		"--kube-version", kubeVersion,
		"--show-only", template,
	}

	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		args = append(args, "--set", fmt.Sprintf("%s=%s", key, values[key]))
	}

	output, err := exec.Command("helm", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("helm template failed: %v\n%s", err, output)
	}

	return string(output)
}

func unmarshalYAML(t *testing.T, contents string, destination any) {
	t.Helper()
	if err := yaml.Unmarshal([]byte(contents), destination); err != nil {
		t.Fatalf("unmarshal rendered YAML: %v", err)
	}
}

var testIDCounter atomic.Uint64

type testIDGenerator struct{}

var random testIDGenerator

func uniqueID() string {
	return fmt.Sprintf("test-%d", testIDCounter.Add(1))
}

func (testIDGenerator) UniqueId() string {
	return uniqueID()
}
