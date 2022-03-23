default: clean

.PHONY: lint
lint:
	### linting charts ###
	helm lint ./charts/rancher-vsphere-cpi/
	helm lint ./charts/rancher-vsphere-csi/

.PHONY: package
package: lint
	### packaging charts for testing purposes ###
	helm package ./charts/rancher-vsphere-cpi/
	helm package ./charts/rancher-vsphere-csi/

.PHONY: clean
clean: package
	### removing chart packages ###
	rm *.tgz