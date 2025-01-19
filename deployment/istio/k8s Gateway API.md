# [Kubernetes Gateway API](https://gateway-api.sigs.k8s.io/)

## [Istio Gateway API](https://istio.io/latest/docs/tasks/traffic-management/ingress/gateway-api/)
По ссылке хороший мануал по использованию istio gateway api.

Пока Istio Gateway API не дает полного функционала istio api.
### 1. Download the Istio release
```bash
cd
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.24.2
export PATH=$PWD/bin:$PATH
```
### 2. [Platform Setup](https://istio.io/latest/docs/setup/platform-setup/)

### 3. [Application Requirements](https://istio.io/latest/docs/ops/deployment/application-requirements/)

### 4. Install Istio

#### Установка CRD
```bash
kubectl get crd gateways.gateway.networking.k8s.io &> /dev/null || \
  { kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.2.0/standard-install.yaml; }
```
#### Установка istio, профиль minimal
```bash
istioctl install --set profile=minimal -y
```

## Объект Gateway

```bash
kubectl create namespace istio-ingress
```

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: gateway
  namespace: istio-ingress
spec:
  gatewayClassName: istio
  listeners:
    - name: http
      hostname: "stateful-todo.k3s.dev.com"
      port: 80
      protocol: HTTP
      allowedRoutes:
        namespaces:
          from: All #Same  # Same - в том же namespace; all - во всех namespace; selector - по селектору
```

`ingress gateway` в кластере может быть несколько, например какие-то за публичным loadbalancerом, а какие-то за приватным.
Соответственно и объектов `Gateway` может быть несколько. В таком случае прописывать имена хостов практически необходимо, чтобы не было путаницы.

## Объект HTTPRoute

Пример:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: stateful-todo-route
  namespace: stateful-todo
spec:
  parentRefs:
    - name: gateway
      namespace: istio-ingress
  hostnames:
    - "stateful-todo.k3s.dev.com"
  rules:
    - matches:
        - path:
            type: PathPrefix
            value: /tasks
      backendRefs:
        - name: stateful-todo-stable
          port: 5000
          weight: 90
        - name: stateful-todo-canary
          port: 5000
          weight: 10
```

### Установить переменные среды Ingress Host
```bash
kubectl wait -n istio-ingress --for=condition=programmed gateways.gateway.networking.k8s.io gateway
export INGRESS_HOST=$(kubectl get gateways.gateway.networking.k8s.io gateway -n istio-ingress -ojsonpath='{.status.addresses[0].value}')
```
