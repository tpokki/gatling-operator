# gatling-operator

## Usage

### Deploy operator

```
# Setup Service Account
$ kubectl create -f deploy/service_account.yaml
# Setup RBAC
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
# Setup the CRD
$ kubectl create -f deploy/crds/tpokki.github.com_gatlingtasks_crd.yaml
# Deploy the app-operator
$ kubectl create -f deploy/operator.yaml
```

### Running the simulations

Currently the only supported way is to define the simulation scripts inline. See [tpokki.github.com_v1alpha1_gatlingtask_cr.yaml](deploy/crds/tpokki.github.com_v1alpha1_gatlingtask_cr.yaml) for an example.

```
kubectl create -f deploy/crds/tpokki.github.com_v1alpha1_gatlingtask_cr.yaml
```

The simulation will loop forever. To stop the simulation, remove the cr
```
kubectl delete -f deploy/crds/tpokki.github.com_v1alpha1_gatlingtask_cr.yaml
```

### Prepare dashboard

Install prometheus to collect metrics from gatling pods. The example below installs it without persistent volumes and alertmanager. See [stable/prometheus](https://github.com/helm/charts/tree/master/stable/prometheus) chart documentation for details.
```
helm install prometheus stable/prometheus --set alertmanager.enabled=false --set server.persistentVolume.enabled=false
```

Next you need something to view the metrics from prometheus. Grafana is good candidate for that. See [stable/grafana](https://github.com/helm/charts/tree/master/stable/grafana) chart documentation for details.
```
helm install grafana stable/grafana --set ingress.enabled=true --set ingress.hosts[0]=grafana.example.com
```

Finally you need to prepare some dashboard with prompql queries to visualise your data. The collected data looks like this:
```
total_started_users{simulation="example",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.005",} 0.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.01",} 0.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.025",} 0.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.05",} 0.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.075",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.1",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.25",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.5",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="0.75",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="1.0",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="2.5",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="5.0",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="7.5",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="10.0",} 1.0
requests_latency_secondsHistogram_bucket{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",le="+Inf",} 1.0
requests_latency_secondsHistogram_count{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",} 1.0
requests_latency_secondsHistogram_sum{simulation="example",metric="request_1 Redirect 1",error="",responseCode="200",oK="OK",} 0.057
```
