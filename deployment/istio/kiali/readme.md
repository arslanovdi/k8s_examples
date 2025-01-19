# Kiali
Графическое представление топологии сервисов. Только для Istio API.

## [Установка](https://istio.io/latest/docs/ops/integrations/kiali/)

```bash
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.24/samples/addons/kiali.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.24/samples/addons/prometheus.yaml

```

### Access to the UI
```bash
kubectl port-forward svc/kiali 20001:20001 -n istio-system
```
or
```bash
./istioctl dashboard kiali
```
Then, access Kiali by visiting https://localhost:20001/ in your preferred web browser.