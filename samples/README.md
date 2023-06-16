

## Test Environment 

### Nats Deploy 

```bash
```
helm repo add nats https://nats-io.github.io/k8s/helm/charts/

helm install nats nats/nats \
--namespace=nats \
--create-namespace=true \
--set nats.jetstream.enabled=true \
--set nats.jetstream.memStorage.enabled=true \
--set nats.jetstream.memStorage.size=1Gi \
--set nats.jetstream.fileStorage.enabled=true \
--set nats.jetstream.fileStorage.size=1Gi \
--set nats.jetstream.fileStorage.storageDirectory=/data/
```
```