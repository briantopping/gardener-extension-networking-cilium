# Using the Networking Cilium extension with Gardener as end-user

The [`core.gardener.cloud/v1beta1.Shoot` resource](https://github.com/gardener/gardener/blob/master/example/90-shoot.yaml) declares a `networking` field that is meant to contain network-specific configuration.

In this document we are describing how this configuration looks like for Cilium and provide an example `Shoot` manifest with minimal configuration that you can use to create a cluster.

## Cilium Hubble

Hubble is a fully distributed networking and security observability platform build on top of Cilium and BPF. It is optional and is deployed to the cluster when enabled in the `NetworkConfig`.
If the dashboard is not externally exposed
```
kubectl port-forward -n kube-system deployment/hubble-ui 8081
```
can be used to acess it locally.

## Example `NetworkingConfig` manifest

An example `NetworkingConfig` for the Cilium extension looks as follows:

```yaml
apiVersion: cilium.networking.extensions.gardener.cloud/v1alpha1
kind: NetworkConfig
hubble:
  enabled: true
#debug: false
#tunnel: vxlan
#store: kubernetes
```

## `NetworkingConfig` options

The `hubble.enabled` field describes whether hubble should be deployed into the cluster or not (default).

The `debug` field describes whether you want to run cilium in debug mode or not (default), change this value to `true` to use debug mode.

The `tunnel` field describes the encapsulation mode for communication between nodes. Possible values are `vxlan` (default), `geneve` or `disabled`.

The `bpfSocketLBHostnsOnly.enabled` field describes whether socket LB will be skipped for services when inside a pod namespace (default), in favor of service LB at the pod interface. Socket LB is still used when in the host namespace. This feature is required when using cilium with a service mesh like istio or linkerd.

Setting the field `cni.exclusive` to `false` might be useful when additional plugins, such as Istio or Linkerd, wish to chain after Cilium. This action disables the default behavior of Cilium, which is to overwrite changes to the CNI configuration file.

The `egressGateway.enabled` field describes whether egress gateways are enabled or not (default). To use this feature kube-proxy must be disabled. This can be done with the following configuration in the Shoot:
```yaml
spec:
  kubernetes:
    kubeProxy:
      enabled: false
```
The egress gateway feature is only supported in gardener with an overlay network (shoot.spec.networking.providerConfig.overlay.enabled: true) at the moment. This is due to the reason that bpf masquerading is required for the egress gateway feature. Once the overlay network is enabled `bpf.masquerade` is set to `true` in the cilium configmap.

The `snatToUpstreamDNS.enabled` field describes whether the traffic to the upstream dns server should be masqueraded or not (default). This is needed on some infrastructures where traffic to the dns server with the pod CIDR range is blocked.

## Example `Shoot` manifest

Please find below an example `Shoot` manifest with cilium networking configuration:

```yaml
apiVersion: core.gardener.cloud/v1beta1
kind: Shoot
metadata:
  name: aws-cilium
  namespace: garden-dev
spec:
  networking:
    type: cilium
    providerConfig:
      apiVersion: cilium.networking.extensions.gardener.cloud/v1alpha1
      kind: NetworkConfig
      hubble:
        enabled: true
    pods: 100.96.0.0/11
    nodes: 10.250.0.0/16
    services: 100.64.0.0/13
  ...
```

If you would like to see a provider specific shoot example, please check out the documentation of the well-known extensions. A list of them can be found [here](https://github.com/gardener/gardener/tree/master/extensions#infrastructure-provider).

## High Availabilty

The cilium-operator operates in high availability (HA) mode when the worker group in the shoot specification is configured with a minimum of at least two nodes as shown in the following example:

```
spec:
  provider:
    workers:
      - ...
        maximum: 5
        minimum: 2
...
```

 This setup prevents unnecessary node spin-ups and reduces the compute costs in single-node clusters.