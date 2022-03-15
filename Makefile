default: clean

.PHONY: lint
lint:
	echo "### linting charts ###"
	helm lint ./charts/rancher-vsphere-cpi/
	helm lint ./charts/rancher-vsphere-csi/

.PHONY: package
package: lint
	echo "### packaging charts for testing purposes ###"
	helm package ./charts/rancher-vsphere-cpi/
	helm package ./charts/rancher-vsphere-csi/

.PHONY: clean
clean: package
	echo "### removing chart packages ###"
	rm *.tgz