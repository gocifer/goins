<!-- This file was autogenerated via cilium-operator --cmdref, do not edit manually-->

## cilium-operator

Run cilium-operator

```
cilium-operator [flags]
```

### Options

```
      --alibaba-cloud-vpc-id string               Specific VPC ID for AlibabaCloud ENI. If not set use same VPC as operator
      --aws-enable-prefix-delegation              Allows operator to allocate prefixes to ENIs instead of individual IP addresses
      --aws-instance-limit-mapping map            Add or overwrite mappings of AWS instance limit in the form of {"AWS instance type": "Maximum Network Interfaces","IPv4 Addresses per Interface","IPv6 Addresses per Interface"}. cli example: --aws-instance-limit-mapping=a1.medium=2,4,4 --aws-instance-limit-mapping=a2.somecustomflavor=4,5,6 configmap example: {"a1.medium": "2,4,4", "a2.somecustomflavor": "4,5,6"}
      --aws-release-excess-ips                    Enable releasing excess free IP addresses from AWS ENI.
      --azure-resource-group string               Resource group to use for Azure IPAM
      --azure-subscription-id string              Subscription ID to access Azure API
      --azure-use-primary-address                 Use Azure IP address from interface's primary IPConfigurations (default true)
      --azure-user-assigned-identity-id string    ID of the user assigned identity used to auth with the Azure API
      --bgp-announce-lb-ip                        Announces service IPs of type LoadBalancer via BGP
      --bgp-config-path string                    Path to file containing the BGP configuration (default "/var/lib/cilium/bgp/config.yaml")
      --cilium-endpoint-gc-interval duration      GC interval for cilium endpoints (default 5m0s)
      --cilium-pod-labels string                  Cilium Pod's labels. Used to detect if a Cilium pod is running to remove the node taints where its running and set NetworkUnavailable to false (default "k8s-app=cilium")
      --cilium-pod-namespace string               Name of the Kubernetes namespace in which Cilium is deployed in. Defaults to the same namespace defined in k8s-namespace
      --cluster-id int                            Unique identifier of the cluster
      --cluster-name string                       Name of the cluster (default "default")
      --cluster-pool-ipv4-cidr strings            IPv4 CIDR Range for Pods in cluster. Requires 'ipam=cluster-pool|cluster-pool-v2beta' and 'enable-ipv4=true'
      --cluster-pool-ipv4-mask-size int           Mask size for each IPv4 podCIDR per node. Requires 'ipam=cluster-pool|cluster-pool-v2beta' and 'enable-ipv4=true' (default 24)
      --cluster-pool-ipv6-cidr strings            IPv6 CIDR Range for Pods in cluster. Requires 'ipam=cluster-pool|cluster-pool-v2beta' and 'enable-ipv6=true'
      --cluster-pool-ipv6-mask-size int           Mask size for each IPv6 podCIDR per node. Requires 'ipam=cluster-pool|cluster-pool-v2beta' and 'enable-ipv6=true' (default 112)
      --cnp-node-status-gc-interval duration      GC interval for nodes which have been removed from the cluster in CiliumNetworkPolicy Status (default 2m0s)
      --cnp-status-update-interval duration       Interval between CNP status updates sent to the k8s-apiserver per-CNP (default 1s)
      --config string                             Configuration file (default "$HOME/ciliumd.yaml")
      --config-dir string                         Configuration directory that contains a file for each option
  -D, --debug                                     Enable debugging mode
      --ec2-api-endpoint string                   AWS API endpoint for the EC2 service
      --enable-cilium-endpoint-slice              If set to true, the CiliumEndpointSlice feature is enabled. If any CiliumEndpoints resources are created, updated, or deleted in the cluster, all those changes are broadcast as CiliumEndpointSlice updates to all of the Cilium agents.
      --enable-ipv4                               Enable IPv4 support (default true)
      --enable-ipv6                               Enable IPv6 support (default true)
      --enable-k8s-api-discovery                  Enable discovery of Kubernetes API groups and resources with the discovery API
      --enable-k8s-endpoint-slice                 Enables k8s EndpointSlice feature into Cilium-Operator if the k8s cluster supports it (default true)
      --enable-k8s-event-handover                 Enable k8s event handover to kvstore for improved scalability
      --enable-metrics                            Enable Prometheus metrics
      --eni-tags map                              ENI tags in the form of k1=v1 (multiple k/v pairs can be passed by repeating the CLI flag)
      --excess-ip-release-delay int               Number of seconds operator would wait before it releases an IP previously marked as excess (default 180)
      --gops-port int                             Port for gops server to listen on (default 9891)
  -h, --help                                      help for cilium-operator
      --identity-allocation-mode string           Method to use for identity allocation (default "kvstore")
      --identity-gc-interval duration             GC interval for security identities (default 15m0s)
      --identity-gc-rate-interval duration        Interval used for rate limiting the GC of security identities (default 1m0s)
      --identity-gc-rate-limit int                Maximum number of security identities that will be deleted within the identity-gc-rate-interval (default 2500)
      --identity-heartbeat-timeout duration       Timeout after which identity expires on lack of heartbeat (default 30m0s)
      --instance-tags-filter stringToString       EC2 Instance tags in the form of k1=v1,k2=v2 (multiple k/v pairs can also be passed by repeating the CLI flag (default [])
      --ipam string                               Backend to use for IPAM (default "cluster-pool")
      --k8s-api-server string                     Kubernetes API server URL
      --k8s-client-burst int                      Burst value allowed for the K8s client
      --k8s-client-qps float32                    Queries per second limit for the K8s client
      --k8s-heartbeat-timeout duration            Configures the timeout for api-server heartbeat, set to 0 to disable (default 30s)
      --k8s-kubeconfig-path string                Absolute path of the kubernetes kubeconfig file
      --k8s-namespace string                      Name of the Kubernetes namespace in which Cilium Operator is deployed in
      --k8s-service-proxy-name string             Value of K8s service-proxy-name label for which Cilium handles the services (empty = all services without service.kubernetes.io/service-proxy-name label)
      --kvstore string                            Key-value store type
      --kvstore-opt map                           Key-value store options e.g. etcd.address=127.0.0.1:4001
      --leader-election-lease-duration duration   Duration that non-leader operator candidates will wait before forcing to acquire leadership (default 15s)
      --leader-election-renew-deadline duration   Duration that current acting master will retry refreshing leadership in before giving up the lock (default 10s)
      --leader-election-retry-period duration     Duration that LeaderElector clients should wait between retries of the actions (default 2s)
      --limit-ipam-api-burst int                  Upper burst limit when accessing external APIs (default 4)
      --limit-ipam-api-qps float                  Queries per second limit when accessing external IPAM APIs (default 20)
      --log-driver strings                        Logging endpoints to use for example syslog
      --log-opt map                               Log driver options for cilium-operator, configmap example for syslog driver: {"syslog.level":"info","syslog.facility":"local4"}
      --nodes-gc-interval duration                GC interval for CiliumNodes
      --operator-api-serve-addr string            Address to serve API requests (default "localhost:9234")
      --operator-prometheus-serve-addr string     Address to serve Prometheus metrics (default ":6942")
      --parallel-alloc-workers int                Maximum number of parallel IPAM workers (default 50)
      --pprof                                     Enable pprof debugging endpoint
      --pprof-port int                            Port that the pprof listens on (default 6061)
      --remove-cilium-node-taints                 Remove node taint "node.cilium.io/agent-not-ready" from Kubernetes nodes once Cilium is up and running (default true)
      --set-cilium-is-up-condition                Set CiliumIsUp Node condition to mark a Kubernetes Node that a Cilium pod is up and running in that node (default true)
      --skip-crd-creation                         When true, Kubernetes Custom Resource Definitions will not be created
      --subnet-ids-filter strings                 Subnets IDs (separated by commas)
      --subnet-tags-filter stringToString         Subnets tags in the form of k1=v1,k2=v2 (multiple k/v pairs can also be passed by repeating the CLI flag (default [])
      --synchronize-k8s-nodes                     Synchronize Kubernetes nodes to kvstore and perform CNP GC (default true)
      --synchronize-k8s-services                  Synchronize Kubernetes services to kvstore (default true)
      --unmanaged-pod-watcher-interval int        Interval to check for unmanaged kube-dns pods (0 to disable) (default 15)
      --update-ec2-adapter-limit-via-api          Use the EC2 API to update the instance type to adapter limits
      --version                                   Print version information
```

### SEE ALSO

* [cilium-operator completion](cilium-operator_completion.md)	 - Generate the autocompletion script for the specified shell
* [cilium-operator metrics](cilium-operator_metrics.md)	 - Access metric status of the operator
