## [Установка через Helm](https://istio.io/latest/docs/setup/install/helm/)
Не смог связать ingress и istiod. Лучше использовать установку через istioctl в конфигурации default.
```bash
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
helm install istio-base istio/base -n istio-system --set defaultRevision=default --create-namespace
helm install istiod istio/istiod -n istio-system --wait

kubectl create namespace istio-ingress
helm install istio-ingress istio/gateway -n istio-ingress --wait


```
## [Установка через istioctl](https://istio.io/latest/docs/setup/install/istioctl/)

### 1. Download the Istio release
```bash
cd
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.24.2
export PATH=$PWD/bin:$PATH
```
### 2. [Platform Setup](https://istio.io/latest/docs/setup/platform-setup/)

### 3. [Application Requirements](https://istio.io/latest/docs/ops/deployment/application-requirements/)

### 4. [Install Istio](https://istio.io/latest/docs/tasks/traffic-management/ingress/gateway-api/)
Конфигурация minimal - без istio-ingress, если используется Istio Gateway API.
```bash
istioctl install --set profile=minimal -y
```

Конфигурация default - c istio-ingress, если используется Istio API.
```bash
istioctl install
```

## Добавление istio-injection sidecar в namespace:
```bash
kubectl label namespace stateful-todo istio-injection=enabled
```

Есть вариант не использовать сайдкар, а разворачивать шлюзы на каждой ноде.

## Ресурсы Istio
```bash
kubectl api-resources | grep istio
```

## [Миграция из ingress в Gateway API](https://gateway-api.sigs.k8s.io/guides/migrating-from-ingress/)

### [ingress2gateway](https://github.com/kubernetes-sigs/ingress2gateway)
Automatic Conversion of Ingresses