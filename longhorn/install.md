# [Longhorn](https://longhorn.io/)
Распределенное хранилище поверх k8s кластера.

## [Установка через helm](https://longhorn.io/docs/1.7.2/deploy/install/install-with-helm/)
```bash
helm repo add longhorn https://charts.longhorn.io
helm repo update
helm install longhorn longhorn/longhorn --namespace longhorn-system --create-namespace --version 1.7.2
```


## [Доступ к longhorn-ui](https://longhorn.io/docs/1.7.2/deploy/accessing-the-ui/longhorn-ingress/)

### Create a basic auth file auth.
```bash
USER=<USERNAME_HERE>; PASSWORD=<PASSWORD_HERE>; echo "${USER}:$(openssl passwd -stdin -apr1 <<< ${PASSWORD})" >> auth
```

### Создаем секрет

Из файла
```bash
kubectl -n longhorn-system create secret generic basic-auth --from-file=auth
```

Либо через деплой auth.yaml, который подготовили на базе файла auth

### Create an Ingress manifest longhorn-ingress.yml

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: longhorn-ingress
  namespace: longhorn-system
  annotations:
    # type of authentication
    nginx.ingress.kubernetes.io/auth-type: basic
    # prevent the controller from redirecting (308) to HTTPS
    nginx.ingress.kubernetes.io/ssl-redirect: 'false'
    # name of the secret that contains the user/password definitions
    nginx.ingress.kubernetes.io/auth-secret: basic-auth
    # message to display with an appropriate context why the authentication is required
    nginx.ingress.kubernetes.io/auth-realm: 'Authentication Required '
    # custom max body size for file uploading like backing image uploading
    nginx.ingress.kubernetes.io/proxy-body-size: 10000m
spec:
  ingressClassName: nginx
  rules:
  - http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: longhorn-frontend
            port:
              number: 80
```