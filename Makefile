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

.PHONY: validate-csi-prime-images
validate-csi-prime-images:
	@test -n "$(PRIME_REGISTRY)" || (echo "PRIME_REGISTRY must be set"; exit 1)
	@for image in $$(yq -r '.. | select(type == "!!map" and has("primeRepository") and has("primeTag")) | "\(.primeRepository):\(.primeTag)"' charts/rancher-vsphere-csi/values.yaml | sort -u); do \
		full_image="$${PRIME_REGISTRY%/}/$$image"; \
		echo "Validating $$image"; \
		docker manifest inspect "$$full_image" > /dev/null || exit 1; \
	done

.PHONY: ci
ci: unit-tests