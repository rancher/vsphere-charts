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
			name: "Kubernetes 1.32 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.32",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.32 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.32",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.31 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.31",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.31 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.31",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.30 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.30",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.30 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.30",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.29 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.29",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.29 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.29",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.28 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.28",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.28 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.28",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
				},
			},
		},
		{
			name: "Kubernetes 1.27 Linux Only",
			args: args{
				values:         map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:    "1.27",
				namespace:      "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:    "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:   csiChart,
				windowsEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.10.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
					"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
				},
			},
		},
		{
			name: "Kubernetes 1.27 Linux and Windows",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":         random.UniqueId(),
					"csiWindowsSupport:enabled": "true",
				},
				kubeVersion:  "1.27",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.10.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
					"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
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
			name: "Kubernetes 1.32",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.32",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.31",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.31",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.30 with Block Snapshotter Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":           random.UniqueId(),
					"blockVolumeSnapshot.enabled": "true",
				},
				kubeVersion:       "1.30",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-snapshotter:v7.0.2",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.30 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.30",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.10.1",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.30",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.30",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		//
		{
			name: "Kubernetes 1.29 with Block Snapshotter Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":           random.UniqueId(),
					"blockVolumeSnapshot.enabled": "true",
				},
				kubeVersion:       "1.29",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-snapshotter:v7.0.2",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.29 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.29",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.10.1",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.29",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.29",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.28 with Block Snapshotter Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":           random.UniqueId(),
					"blockVolumeSnapshot.enabled": "true",
				},
				kubeVersion:       "1.28",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-snapshotter:v7.0.2",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.28 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.28",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.10.1",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.28",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.28",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
					"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
				},
			},
		},
		{
			name: "Kubernetes 1.27 with CSI Resizer Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":                random.UniqueId(),
					"csiController.csiResizer.enabled": "true",
				},
				kubeVersion:       "1.27",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.5.0",
					"rancher/mirrored-sig-storage-csi-resizer:v1.10.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
					"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.2.0",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.0",
				},
			},
		},
		{
			name: "Kubernetes 1.27 with Block Snapshotter Enabled",
			args: args{
				values: map[string]string{
					"vCenter.clusterId":           random.UniqueId(),
					"blockVolumeSnapshot.enabled": "true",
				},
				kubeVersion:       "1.27",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.5.0",
					"rancher/mirrored-sig-storage-csi-snapshotter:v7.0.1",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
					"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.2.0",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.0",
				},
			},
		},
		{
			name: "Kubernetes 1.27",
			args: args{
				values:            map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:       "1.27",
				namespace:         "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:       "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath:      csiChart,
				csiResizerEnabled: false,
				expectedImages: []string{
					"rancher/mirrored-sig-storage-csi-attacher:v4.5.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
					"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
					"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.2.0",
					"rancher/mirrored-sig-storage-csi-provisioner:v4.0.0",
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

func TestCSITemplateRenderedControllerDeploymentArgs(t *testing.T) {
	type args struct {
		values       map[string]string
		kubeVersion  string
		namespace    string
		releaseName  string
		chartRelPath string
		expectedArgs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Kubernetes 1.32",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.32",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.31",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.31",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.30",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.30",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.29",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.29",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.28",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.28",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.27",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.27",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
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
			var args []string
			for _, container := range deployment.Spec.Template.Spec.Containers {
				if container.Name == "vsphere-csi-controller" {
					args = container.Args
				}
			}

			// assert
			require.Equal(t, tt.args.namespace, deployment.Namespace)
			require.Equal(t, len(tt.args.expectedArgs), len(args))
			for i := range tt.args.expectedArgs {
				require.Equal(t, tt.args.expectedArgs[i], args[i])
			}
		})
	}
}

func TestCSITemplateRenderedNodeDaemonSetArgs(t *testing.T) {
	type args struct {
		values       map[string]string
		kubeVersion  string
		namespace    string
		releaseName  string
		chartRelPath string
		expectedArgs []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Kubernetes 1.32",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.32",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.31",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.31",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.30",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.30",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.29",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.29",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.28",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.28",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
				},
			},
		},
		{
			name: "Kubernetes 1.27",
			args: args{
				values:       map[string]string{"vCenter.clusterId": random.UniqueId()},
				kubeVersion:  "1.27",
				namespace:    "csitest-" + strings.ToLower(random.UniqueId()),
				releaseName:  "csitest-" + strings.ToLower(random.UniqueId()),
				chartRelPath: csiChart,
				expectedArgs: []string{
					"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
					"--fss-namespace=$(CSI_NAMESPACE)",
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

			var deployment appsv1.Deployment
			helm.UnmarshalK8SYaml(t, output, &deployment)
			var args []string
			for _, container := range deployment.Spec.Template.Spec.Containers {
				if container.Name == "vsphere-csi-node" {
					args = container.Args
				}
			}

			// assert
			require.Equal(t, tt.args.namespace, deployment.Namespace)
			require.Equal(t, len(tt.args.expectedArgs), len(args))
			for i := range tt.args.expectedArgs {
				require.Equal(t, tt.args.expectedArgs[i], args[i])
			}
		})
	}
}
