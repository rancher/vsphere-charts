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
			name: "Kubernetes 1.27",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.27",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.26.0",
			},
		},
		{
			name: "Kubernetes 1.26",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.26",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.26.0",
			},
		},
		{
			name: "Kubernetes 1.25",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.25",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.25.0",
			},
		},
		{
			name: "Kubernetes 1.24",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.24",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.24.3",
			},
		},
		{
			name: "Kubernetes 1.23",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.23",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.23.3",
			},
		},
		{
			name: "Kubernetes 1.22",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.22",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.22.7",
			},
		},
		{
			name: "Kubernetes 1.21",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.21",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.21.3",
			},
		},
		{
			name: "Kubernetes 1.20",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.20",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.20.1",
			},
		},
		{
			name: "Kubernetes 1.19",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.19",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.19.0",
			},
		},
		{
			name: "Kubernetes 1.18",
			args: args{
				values:        map[string]string{},
				kubeVersion:   "1.18",
				namespace:     "cpitest-" + strings.ToLower(random.UniqueId()),
				releaseName:   "cpitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:  cpiChart,
				expectedImage: "rancher/mirrored-cloud-provider-vsphere-cpi-release-manager:v1.18.0",
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
