#### Pull Request Checklist ####

- [ ] Any new images or tags consumed by charts has been added [here](https://github.com/rancher/image-mirror)
- [ ] Chart version has been incremented (if necessary)
- [ ] That helm lint and pack run successfully on the chart.
- [ ] Deployment of the chart has been tested and verified that it functions as expected.
- [ ] Changes to scripting or CI config have been tested to the best of your ability

#### Types of Change ####

<!-- New image, version bump. script update, etc etc -->

#### Linked Issues ####

<!-- Link any related issues, pull-requests, or commit hashes that are relevant to this pull request.  -->

#### Additional Notes ####

<!-- Any additional details / test results / etc -->

#### After the PR is merged ####

Once the PR is merged, typically upon a new release, the necessary teams will be notified Slack hook to perform the rancher/rke2-charts and RKE2 changes. Any developer working on this issue is not responsible for updating rancher/rke2-charts or RKE2.