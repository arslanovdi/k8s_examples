# ArgoCD

## install

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

Для доступа к ArgoCD UI нужно внести изменение в деплой `argocd-server`.

Добавить аргумент `--insecure` 

```yaml
      containers:
      - args:
        - /usr/local/bin/argocd-server
        - --insecure
```

И настроить в моем случае istio virtualservice + istio gateway для маршутизации трафика с 80 порта в argocd-server:80

Имя пользователя по умолчанию: `admin`.

Пароль находится в секрете `argocd-initial-admin-secret`.

## Установка ArgoCD CLI

```bash
curl -sSL -o argocd-linux-amd64 https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
sudo install -m 555 argocd-linux-amd64 /usr/local/bin/argocd
rm argocd-linux-amd64
```

### Подключение к ArgoCD через CLI
```bash
export password=$(kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo)
argocd login http://argocd.k3s.dev.com:80 --username admin --password $password --insecure
```
