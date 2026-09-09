package unit

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

const cpiChart = "../../charts/rancher-vsphere-cpi"

func TestCPITemplateRenderedDaemonset(t *testing.T) {
	type args struct {
		values        map[string]string
		kubeVersion   string
		namespace     string
		releaseName   string
		chartRelPath  string
		expectedImage string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Kubernetes 1.37",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.37",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.37.0",
			},
		},
		{
			name: "Kubernetes 1.37 Prime",
			args: args{
				values: map[string]string{
					"global.prime.enabled": "true",
				},
				kubeVersion:   "1.37",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/hardened-cloud-provider-vsphere:v1.37.0-build20260901",
			},
		},
		{
			name: "Kubernetes 1.36",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.36",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.36.0",
			},
		},
		{
			name: "Kubernetes 1.36 Prime",
			args: args{
				values: map[string]string{
					"global.prime.enabled": "true",
				},
				kubeVersion:   "1.36",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/hardened-cloud-provider-vsphere:v1.36.0-build20260722",
			},
		},
		{
			name: "Kubernetes 1.36 Prime with system default registry",
			args: args{
				values: map[string]string{
					"global.prime.enabled":                "true",
					"global.cattle.systemDefaultRegistry": "registry.rancher.com",
				},
				kubeVersion:   "1.36",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "registry.rancher.com/rancher/hardened-cloud-provider-vsphere:v1.36.0-build20260722",
			},
		},
		{
			name: "Kubernetes 1.35",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.35",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.35.1",
			},
		},
		{
			name: "Kubernetes 1.34",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.34",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.34.0",
			},
		},
		{
			name: "Kubernetes 1.33",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.33",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.33.1",
			},
		},
		{
			name: "Kubernetes 1.32",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.32",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.32.2",
			},
		},
		{
			name: "Kubernetes 1.31",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.31",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.31.1",
			},
		},
		{
			name: "Kubernetes 1.30",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.30",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.30.1",
			},
		},
		{
			name: "Kubernetes 1.29",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.29",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.29.0",
			},
		},
		{
			name: "Kubernetes 1.28",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.28",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.28.0",
			},
		},
		{
			name: "Kubernetes 1.27",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.27",
				namespace:     "cpitest",
				releaseName:   "cpitest",
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.27.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			chartPath, err := filepath.Abs(tt.args.chartRelPath)
			require.NoError(t, err)

			// act
			output := renderTemplate(t, chartPath, tt.args.releaseName, tt.args.namespace, tt.args.kubeVersion, tt.args.values, "templates/daemonset.yaml")

			var daemonSet appsv1.DaemonSet
			unmarshalYAML(t, output, &daemonSet)

			// assert
			require.Equal(t, tt.args.namespace, daemonSet.Namespace)
			daemonSetContainers := daemonSet.Spec.Template.Spec.Containers
			require.Equal(t, 1, len(daemonSetContainers))
			require.Equal(t, tt.args.expectedImage, daemonSetContainers[0].Image)
		})
	}
}

func TestCPITemplateRenderedSecret(t *testing.T) {
	// arrange
	chartPath, err := filepath.Abs(cpiChart)
	require.NoError(t, err)

	namespace := "cpitest"
	releaseName := "cpitest"
	values := map[string]string{"vCenter.host": "test", "vCenter.username": "test", "vCenter.password": "test"}

	// act
	output := renderTemplate(t, chartPath, releaseName, namespace, "1.23", values, "templates/secret.yaml")

	var secret v1.Secret
	unmarshalYAML(t, output, &secret)

	// assert
	require.Equal(t, namespace, secret.Namespace)
	require.NotEmpty(t, secret.Data)
	require.Contains(t, secret.Data, "test.username")
	require.Contains(t, secret.Data, "test.password")
}

func TestCPITemplateRenderedServiceAccount(t *testing.T) {
	// arrange
	chartPath, err := filepath.Abs(cpiChart)
	require.NoError(t, err)

	namespace := "cpitest"
	releaseName := "cpitest"

	// act
	output := renderTemplate(t, chartPath, releaseName, namespace, "1.23", nil, "templates/service-account.yaml")
	var sa v1.ServiceAccount
	unmarshalYAML(t, output, &sa)

	// assert
	require.Equal(t, namespace, sa.Namespace)
}
