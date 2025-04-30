package unit

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
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
			name: "Kubernetes 1.33",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.33",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.33.0",
			},
		},
		{
			name: "Kubernetes 1.32",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.32",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.32.2",
			},
		},
		{
			name: "Kubernetes 1.31",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.31",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere:v1.31.1",
			},
		},
		{
			name: "Kubernetes 1.30",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.30",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.30.1",
			},
		},
		{
			name: "Kubernetes 1.29",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.29",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.29.0",
			},
		},
		{
			name: "Kubernetes 1.28",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.28",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.28.0",
			},
		},
		{
			name: "Kubernetes 1.27",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.27",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
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

			options := &helm.Options{
				SetValues:      tt.args.values,
				KubectlOptions: k8s.NewKubectlOptions("", "", tt.args.namespace),
			}

			// act
			output := helm.RenderTemplate(t, options, chartPath, tt.args.releaseName, []string{"templates/daemonset.yaml"}, "--kube-version", tt.args.kubeVersion)

			var daemonSet appsv1.DaemonSet
			helm.UnmarshalK8SYaml(t, output, &daemonSet)

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

	namespace := "cpitest-" + strings.ToLower(random.UniqueId())
	releaseName := "cpitest-" + strings.ToLower(random.UniqueId())
	options := &helm.Options{
		SetValues:      map[string]string{"vCenter.host": "test", "vCenter.username": "test", "vCenter.password": "test"},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespace),
	}

	// act
	output := helm.RenderTemplate(t, options, chartPath, releaseName, []string{"templates/secret.yaml"}, "--kube-version", "1.23")

	var secret v1.Secret
	helm.UnmarshalK8SYaml(t, output, &secret)

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

	namespace := "cpitest-" + strings.ToLower(random.UniqueId())
	releaseName := "cpitest-" + strings.ToLower(random.UniqueId())
	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", namespace),
	}

	// act
	output := helm.RenderTemplate(t, options, chartPath, releaseName, []string{"templates/service-account.yaml"}, "--kube-version", "1.23")
	var sa v1.ServiceAccount
	helm.UnmarshalK8SYaml(t, output, &sa)

	// assert
	require.Equal(t, namespace, sa.Namespace)
}
