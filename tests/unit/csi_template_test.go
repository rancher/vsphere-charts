package unit

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

const csiChart = "../../charts/rancher-vsphere-csi"

var csi29_30NodeImages = []string{
	"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
	"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
}

var csi29_30ControllerImages = []string{
	"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
	"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
	"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
}

var csiDefaultNodeImages = []string{
	"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.13.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.7.3",
	"rancher/mirrored-sig-storage-livenessprobe:v2.15.0",
}

var csiDefaultControllerImages = []string{
	"rancher/mirrored-sig-storage-csi-attacher:v4.9.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.7.3",
	"rancher/mirrored-sig-storage-livenessprobe:v2.15.0",
	"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.7.3",
	"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
}

func TestCSITemplateRenderedNodeDaemonset(t *testing.T) {
	tests := []struct {
		name           string
		kubeVersion    string
		expectedImages []string
	}{
		{name: "Kubernetes 1.27", kubeVersion: "1.27", expectedImages: []string{
			"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.10.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
			"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
		}},
		{name: "Kubernetes 1.28", kubeVersion: "1.28", expectedImages: []string{
			"rancher/mirrored-sig-storage-csi-node-driver-registrar:v2.12.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
			"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
		}},
		{name: "Kubernetes 1.29", kubeVersion: "1.29", expectedImages: csi29_30NodeImages},
		{name: "Kubernetes 1.30", kubeVersion: "1.30", expectedImages: csi29_30NodeImages},
		{name: "Kubernetes 1.31", kubeVersion: "1.31", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.32", kubeVersion: "1.32", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.33", kubeVersion: "1.33", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.34", kubeVersion: "1.34", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.35", kubeVersion: "1.35", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.36", kubeVersion: "1.36", expectedImages: csiDefaultNodeImages},
		{name: "Kubernetes 1.37", kubeVersion: "1.37", expectedImages: csiDefaultNodeImages},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), test.kubeVersion, map[string]string{"vCenter.clusterId": random.UniqueId()}, "templates/node/daemonset.yaml")
			var daemonSet appsv1.DaemonSet
			unmarshalYAML(t, output, &daemonSet)
			require.Equal(t, test.expectedImages, containerImages(daemonSet.Spec.Template.Spec.Containers))
		})
	}
}

func TestCSITemplateRenderedNodeDaemonsetPrimeAndWindows(t *testing.T) {
	tests := []struct {
		name        string
		kubeVersion string
		template    string
	}{
		{name: "Kubernetes 1.36 Linux", kubeVersion: "1.36", template: "templates/node/daemonset.yaml"},
		{name: "Kubernetes 1.36 Windows", kubeVersion: "1.36", template: "templates/node/windows-daemonset.yaml"},
		{name: "Kubernetes 1.37 Linux", kubeVersion: "1.37", template: "templates/node/daemonset.yaml"},
		{name: "Kubernetes 1.37 Windows", kubeVersion: "1.37", template: "templates/node/windows-daemonset.yaml"},
	}
	expectedImages := []string{
		"registry.rancher.com/rancher/hardened-csi-node-driver-registrar:v2.13.0-build20260909",
		"registry.rancher.com/rancher/hardened-vsphere-csi-driver:v3.7.3-build20260909",
		"registry.rancher.com/rancher/hardened-livenessprobe:v2.15.0-build20260909",
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			values := map[string]string{
				"vCenter.clusterId":                   random.UniqueId(),
				"global.prime.enabled":                "true",
				"global.cattle.systemDefaultRegistry": "registry.rancher.com",
				"csiWindowsSupport.enabled":           "true",
			}
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), test.kubeVersion, values, test.template)
			var daemonSet appsv1.DaemonSet
			unmarshalYAML(t, output, &daemonSet)
			require.Equal(t, expectedImages, containerImages(daemonSet.Spec.Template.Spec.Containers))
		})
	}
}

func TestCSITemplateRenderedControllerDeployment(t *testing.T) {
	tests := []struct {
		name           string
		kubeVersion    string
		expectedImages []string
	}{
		{name: "Kubernetes 1.27", kubeVersion: "1.27", expectedImages: []string{
			"rancher/mirrored-sig-storage-csi-attacher:v4.5.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.2.0",
			"rancher/mirrored-sig-storage-livenessprobe:v2.12.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.2.0",
			"rancher/mirrored-sig-storage-csi-provisioner:v4.0.0",
		}},
		{name: "Kubernetes 1.28", kubeVersion: "1.28", expectedImages: []string{
			"rancher/mirrored-sig-storage-csi-attacher:v4.7.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-driver:v3.3.1",
			"rancher/mirrored-sig-storage-livenessprobe:v2.14.0",
			"rancher/mirrored-cloud-provider-vsphere-csi-release-syncer:v3.3.1",
			"rancher/mirrored-sig-storage-csi-provisioner:v4.0.1",
		}},
		{name: "Kubernetes 1.29", kubeVersion: "1.29", expectedImages: csi29_30ControllerImages},
		{name: "Kubernetes 1.30", kubeVersion: "1.30", expectedImages: csi29_30ControllerImages},
		{name: "Kubernetes 1.31", kubeVersion: "1.31", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.32", kubeVersion: "1.32", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.33", kubeVersion: "1.33", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.34", kubeVersion: "1.34", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.35", kubeVersion: "1.35", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.36", kubeVersion: "1.36", expectedImages: csiDefaultControllerImages},
		{name: "Kubernetes 1.37", kubeVersion: "1.37", expectedImages: csiDefaultControllerImages},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), test.kubeVersion, map[string]string{"vCenter.clusterId": random.UniqueId()}, "templates/controller/deployment.yaml")
			var deployment appsv1.Deployment
			unmarshalYAML(t, output, &deployment)
			require.Equal(t, test.expectedImages, containerImages(deployment.Spec.Template.Spec.Containers))
		})
	}
}

func TestCSITemplateRenderedControllerDeploymentPrime(t *testing.T) {
	values := map[string]string{
		"vCenter.clusterId":                   random.UniqueId(),
		"global.prime.enabled":                "true",
		"global.cattle.systemDefaultRegistry": "registry.rancher.com",
	}
	expectedImages := []string{
		"registry.rancher.com/rancher/hardened-csi-attacher:v4.9.0-build20260909",
		"registry.rancher.com/rancher/hardened-vsphere-csi-driver:v3.7.3-build20260909",
		"registry.rancher.com/rancher/hardened-livenessprobe:v2.15.0-build20260909",
		"registry.rancher.com/rancher/hardened-vsphere-csi-syncer:v3.7.3-build20260909",
		"registry.rancher.com/rancher/hardened-csi-provisioner:v4.0.1-build20260909",
	}
	chartPath, err := filepath.Abs(csiChart)
	require.NoError(t, err)
	output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), "1.37", values, "templates/controller/deployment.yaml")
	var deployment appsv1.Deployment
	unmarshalYAML(t, output, &deployment)
	require.Equal(t, expectedImages, containerImages(deployment.Spec.Template.Spec.Containers))
}

func TestCSITemplateRenderedControllerDeploymentOptionalImages(t *testing.T) {
	tests := []struct {
		name           string
		values         map[string]string
		kubeVersion    string
		expectedImages []string
	}{
		{
			name:        "Kubernetes 1.30 block snapshotter",
			kubeVersion: "1.30",
			values: map[string]string{
				"vCenter.clusterId":           random.UniqueId(),
				"blockVolumeSnapshot.enabled": "true",
			},
			expectedImages: append([]string{"rancher/mirrored-sig-storage-csi-attacher:v4.7.0", "rancher/mirrored-sig-storage-csi-snapshotter:v7.0.2"}, csi29_30ControllerImages[1:]...),
		},
		{
			name:        "Kubernetes 1.30 resizer",
			kubeVersion: "1.30",
			values: map[string]string{
				"vCenter.clusterId":                random.UniqueId(),
				"csiController.csiResizer.enabled": "true",
			},
			expectedImages: append([]string{"rancher/mirrored-sig-storage-csi-attacher:v4.7.0", "rancher/mirrored-sig-storage-csi-resizer:v1.10.1"}, csi29_30ControllerImages[1:]...),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), test.kubeVersion, test.values, "templates/controller/deployment.yaml")
			var deployment appsv1.Deployment
			unmarshalYAML(t, output, &deployment)
			require.Equal(t, test.expectedImages, containerImages(deployment.Spec.Template.Spec.Containers))
		})
	}
}

func TestCSITemplateRenderedControllerDeploymentArgs(t *testing.T) {
	expectedArgs := []string{
		"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
		"--fss-namespace=$(CSI_NAMESPACE)",
	}
	for _, kubeVersion := range []string{"1.27", "1.28", "1.29", "1.30", "1.31", "1.32", "1.33", "1.34", "1.35"} {
		t.Run("Kubernetes "+kubeVersion, func(t *testing.T) {
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), kubeVersion, map[string]string{"vCenter.clusterId": random.UniqueId()}, "templates/controller/deployment.yaml")
			var deployment appsv1.Deployment
			unmarshalYAML(t, output, &deployment)
			var args []string
			for _, container := range deployment.Spec.Template.Spec.Containers {
				if container.Name == "vsphere-csi-controller" {
					args = container.Args
				}
			}
			require.Equal(t, expectedArgs, args)
		})
	}
}

func TestCSITemplateRenderedNodeDaemonSetArgs(t *testing.T) {
	expectedArgs := []string{
		"--fss-name=internal-feature-states.csi.vsphere.vmware.com",
		"--fss-namespace=$(CSI_NAMESPACE)",
	}
	for _, kubeVersion := range []string{"1.27", "1.28", "1.29", "1.30", "1.31", "1.32", "1.33", "1.34", "1.35"} {
		t.Run("Kubernetes "+kubeVersion, func(t *testing.T) {
			chartPath, err := filepath.Abs(csiChart)
			require.NoError(t, err)
			output := renderTemplate(t, chartPath, "csitest-"+strings.ToLower(random.UniqueId()), "csitest-"+strings.ToLower(random.UniqueId()), kubeVersion, map[string]string{"vCenter.clusterId": random.UniqueId()}, "templates/node/daemonset.yaml")
			var daemonSet appsv1.DaemonSet
			unmarshalYAML(t, output, &daemonSet)
			var args []string
			for _, container := range daemonSet.Spec.Template.Spec.Containers {
				if container.Name == "vsphere-csi-node" {
					args = container.Args
				}
			}
			require.Equal(t, expectedArgs, args)
		})
	}
}

func containerImages(containers []corev1.Container) []string {
	images := make([]string, 0, len(containers))
	for _, container := range containers {
		images = append(images, container.Image)
	}
	return images
}
