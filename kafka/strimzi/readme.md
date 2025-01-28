# install

```bash
kubectl create namespace kafka
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
```

# uninstall

```bash
kubectl -n kafka delete -f 'https://strimzi.io/install/latest?namespace=kafka'
kubectl delete namespace kafka
```