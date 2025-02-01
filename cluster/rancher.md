# Rancher
Rancher - ПО для разворачивания и управления kubernetes кластерами.

Для отказоустойчивости rancher должен быть развернут в kubernetes кластере.

Для теста можно скачать докер образ rancher/rancher. В этом образе одноузловой k3s кластер с установленным rancher.

## Запуск в докере.

```bash
docker run -d --restart=unless-stopped --name rancher -p 80:80 -p 443:443 --privileged rancher/rancher:v2.10-head
```
или

```bash
sudo docker run -d --restart=unless-stopped --name rancher -p 80:80 -p 443:443 --privileged rancher/rancher --acme-domain <ваш домен>
```

## [Установка Rancher на сервер](https://ranchermanager.docs.rancher.com/getting-started/installation-and-upgrade/install-upgrade-on-a-kubernetes-cluster)
Требования:
1. Kubectl
2. Helm
3. cert-manager

### Установка kubectl (не требуется)
```bash
curl -LO https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
chmod +x kubectl
mkdir -p ~/.local/bin
mv ./kubectl ~/.local/bin/kubectl
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
sudo chmod 755 /etc/rancher/k3s/k3s.yaml
```

### Символическая ссылка на существующий kubectl
```bash
ln -s /var/lib/rancher/rke2/bin/kubectl /usr/local/bin/kubectl
cp /etc/rancher/rke2/rke2.yaml ~/.kube/config
```

Проверка версии:
```bash
kubectl version
```

### [Установка helm](https://cert-manager.io/docs/installation/helm/#installing-with-helm)
```bash
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

Добавить репозиторий Helm:
```bash
helm repo add rancher-latest https://releases.rancher.com/server-charts/latest
helm repo add rancher-stable https://releases.rancher.com/server-charts/stable
helm repo add rancher-alpha https://releases.rancher.com/server-charts/alpha
kubectl create namespace cattle-system
```

## [Установка cert-manager](https://ranchermanager.docs.rancher.com/getting-started/installation-and-upgrade/resources/upgrade-cert-manager)
```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.16.2/cert-manager.crds.yaml
```

```bash
helm repo add jetstack https://charts.jetstack.io
helm repo update
```

```bash
helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.16.2 \
  --set crds.enabled=true
```

Проверим:
```bash
kubectl get pods --namespace cert-manager
```

## Установка Rancher
hostname - A запись DNS на узел(ы) k3s.
```bash
helm install rancher rancher-stable/rancher \
  --namespace cattle-system \
  --set hostname=k8s-master.dev.com \
  --set bootstrapPassword=admin
```

Дождаться завершения установки:
```bash
kubectl -n cattle-system rollout status deploy/rancher
```

Подключаемся к Rancher по доменному имени: `https://k8s-master.dev.com`

Вуаля!


## Развертывание нового Custom кластера

В веб интерфейсе rancher -> Clusters -> Add custom cluster.

Потребуются узлы с установленной Ubuntu server/ 