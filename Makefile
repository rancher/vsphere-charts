default: ci

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
	### removing chart packages ###
	rm *.tgz

.PHONY: setup
setup:
	### setup ###
	go mod tidy

.PHONY: unit-tests
unit-tests: package setup
	### running unit tests ###
	go test -v -tags helm ./tests/unit

.PHONY: ci
ci: unit-tests