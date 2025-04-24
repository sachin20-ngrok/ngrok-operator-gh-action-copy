# Plans

These files define different test scenarios we might want to run. These are specified as the first argument to the main program and determine which plan will be run. These are also able to
be specified to the GitHub action, so that you can validate changes against a small number of endpoints, before ramping it up. It also allows you to test different scenarios like using ingress or just using the plain CRDs.

- [agent-endpoints-with-traffic-policy.yaml](./agent-endpoints-with-traffic-policy.yaml): This plan spins up 2 copies of the same 2048 application using agent endpoint CRDs and applies a traffic policy with an `oauth` action. The expected HTTP status of both tests is 302.
- [gateway.yaml](./gateway.yaml): This plan spins up 2 copies of the same 2048 application using the Kubernetes Gateway API. The expected HTTP status of both tests is 200.
