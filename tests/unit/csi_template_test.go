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
)

const csiChart = "../../charts/rancher-vsphere-csi"

func TestCSITemplateRenderedNodeDaemonset(t *testing.T) {
	type args struct {
		values         map[string]string
		kubeVersion    string
		namespace      string
		releaseName    string
		chartRelPath   string
		windowsEnabled bool
		expectedImages []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Kubernetes 1.23 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.23",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.23 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.23",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.22 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.22",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.22 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.22",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.21 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.21",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.21 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.21",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
			},
		},
		{
			name: "Kubernetes 1.20 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.20",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.3.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.4.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.4.0",
				},
			},
		},
		{
			name: "Kubernetes 1.21 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.21",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
				},
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
			output := helm.RenderTemplate(t, options, chartPath, tt.args.releaseName, []string{"templates/node/daemonset.yaml"}, "--kube-version", tt.args.kubeVersion)

			var daemonSet appsv1.DaemonSet
			helm.UnmarshalK8SYaml(t, output, &daemonSet)

			// assert
			require.Equal(t, tt.args.namespace, daemonSet.Namespace)
			daemonSetContainers := daemonSet.Spec.Template.Spec.Containers
			require.Equal(t, len(tt.args.expectedImages), len(daemonSetContainers))
			for i := range tt.args.expectedImages {
				require.Equal(t, tt.args.expectedImages[i], daemonSetContainers[i].Image)
			}

			if tt.args.windowsEnabled {
				// act
				windowsOutput := helm.RenderTemplate(t, options, chartPath, tt.args.releaseName, []string{"templates/node/windows-daemonset.yaml"}, "--kube-version", tt.args.kubeVersion)

				var windowsDaemonSet appsv1.DaemonSet
				helm.UnmarshalK8SYaml(t, windowsOutput, &daemonSet)

				// assert
				require.Equal(t, tt.args.namespace, windowsDaemonSet.Namespace)
				windowsDaemonSetSetContainers := windowsDaemonSet.Spec.Template.Spec.Containers
				require.Equal(t, len(tt.args.expectedImages), len(windowsDaemonSetSetContainers))
				for i := range tt.args.expectedImages {
					require.Equal(t, tt.args.expectedImages[i], windowsDaemonSetSetContainers[i].Image)
				}
			}
		})
	}
}

func TestCSITemplateRenderedControllerDeployment(t *testing.T) {
	type args struct {
		values            map[string]string
		kubeVersion       string
		namespace         string
		releaseName       string
		chartRelPath      string
		csiResizerEnabled bool
		expectedImages    []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Kubernetes 1.23",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.23",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.5.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.1.0",
				},
			},
		},
		{
			name: "Kubernetes 1.22",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.22",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.5.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.1.0",
				},
			},
		},
		{
			name: "Kubernetes 1.21",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.21",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.5.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.1.0",
				},
			},
		},
		{
			name: "Kubernetes 1.20",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.20",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.3.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.4.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.4.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.0.0",
				},
			},
		},
		{
			name: "Kubernetes 1.23 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.23",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.4.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.5.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.6.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.5.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.1.0",
				},
			},
		},
		{
			name: "Kubernetes 1.20 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.20",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v3.3.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.3.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v2.4.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.4.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v2.4.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v3.0.0",
				},
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
			output := helm.RenderTemplate(t, options, chartPath, tt.args.releaseName, []string{"templates/controller/deployment.yaml"}, "--kube-version", tt.args.kubeVersion)

			var deployment appsv1.Deployment
			helm.UnmarshalK8SYaml(t, output, &deployment)

			// assert
			require.Equal(t, tt.args.namespace, deployment.Namespace)
			deploymentSetContainers := deployment.Spec.Template.Spec.Containers

			require.Equal(t, len(tt.args.expectedImages), len(deploymentSetContainers))
			for i := range tt.args.expectedImages {
				require.Equal(t, tt.args.expectedImages[i], deploymentSetContainers[i].Image)
			}
		})
	}
}
