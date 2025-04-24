# scale testing the ngrok operator

This repo uses Github Actions to deploy the ngrok operator in a single node Kubernetes cluster.

On push to this repo, the `kubernetes.yml` file in `.github/workflows` runs. It:

1. starts a k8s cluster
1. installs the Kubernetes Gateway API
1. installs Helm
1. installs ngrok Helm chart
1. deploys ngrok Helm chart
1. installs the app defined in `charts/app`
1. spins down and cleans up

## ngrok account info

This repo uses features that are only available on a Pay-As-You-Go plan.
