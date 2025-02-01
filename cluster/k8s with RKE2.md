# Развертывание Kubernetes кластера на базе образов RKE2 Rancher

# Подготовка узлов кластера для развертывания kubernetes
В результате сконфигурируется минимальный отказоустойчивый RKE кластер из 3 узлов, в котором запущены etcd, Kubernetes API, CNI, Ingress контроллер, прочие control plane сервисы и компоненты.

[Требования к железу](https://docs.rke2.io/install/requirements): 2 vCPU и 4 GB оперативной памяти на каждый узел.

В кластере должно быть нечетное количество узлов!

Минимум необходимо 3 узла:
1. control plane
2. worker 1
3. worker 2

### 1. Развернуть [ubuntu-server](https://ubuntu.com/download/server) на каждом узле
Я использовал последнюю на данный момент версию Ubuntu 24.04.1 LTS.

Выдать статические IP.

### 2.  Создать общую A запись
Каждый из узлов кластера должен быть доступен по доменному имени.

Необходимо создать A запись в DNS сервере и добавить IP всех узлов.

Либо как я сделал в домашней среде - в маршрутизаторе (у меня mikrotik) создать DNS Static записи для каждого узла, с одинаковым dns именем.

### 3. Установка RKE2 на control plane
RKE2 – обновленная и улучшенная версия дистрибутива rancher, который не зависит от Docker в отличие от предшественника.

Сервисы control plane запускаются как статические поды под управлением kubelet, а в качестве container runtime – containerd. 

Для начала установки RKE2 нужно подключиться по SSH на первый хост. Всю конфигурацию необходимо выполнять под root пользователем, выполнив `sudo -i`

1. Скачать и запустить инсталлятор:
```bash
sudo -i
curl -sfL https://get.rke2.io | sh -
```
2. Добавить параметр tls-san в конфиг:
```bash
mkdir /etc/rancher/rke2 -p &&
nano  /etc/rancher/rke2/config.yaml
```
Добавить в него dns запись кластера:
```
tls-san:
  - k8s-master.dev.com
```
3. Запустить инсталлятор:
```bash
systemctl enable --now rke2-server.service
```
4. Проверить логи и статус k8s подов
```bash
journalctl -u rke2-server -f

/var/lib/rancher/rke2/bin/kubectl get pods -A --kubeconfig /etc/rancher/rke2/rke2.yaml
```
5. Скопировать токен из файла /var/lib/rancher/rke2/server/node-token
Потребуется для конфигурации остальных узлов кластера.
```bash
cat /var/lib/rancher/rke2/server/node-token
```
`K1094a14b92638ebc42c79a6330143c18e21ba10687fc023dac69bcb204430baf1b::server:ac6f9c671eda0fbf74b0b1313365aeee`
### 4. Установка RKE2 на остальные узлы
Отличается только добавлением токена в конфигурацию.

1. Скачать и запустить инсталлятор:
```bash
sudo -i
curl -sfL https://get.rke2.io | sh -
```
2. Добавить параметр tls-san и токен в конфиг:
```bash
mkdir /etc/rancher/rke2 -p &&
nano  /etc/rancher/rke2/config.yaml
```
Добавить в него dns запись кластера:
```
tls-san:
  - k8s-master.dev.com
token: K1094a14b92638ebc42c79a6330143c18e21ba10687fc023dac69bcb204430baf1b::server:ac6f9c671eda0fbf74b0b1313365aeee
server: https://k8s-master.dev.com:9345
```
3. Запустить инсталлятор:
```bash
systemctl enable --now rke2-server.service
```
4. Проверить логи и статус k8s подов
```bash
journalctl -u rke2-server -f

/var/lib/rancher/rke2/bin/kubectl get pods -A --kubeconfig /etc/rancher/rke2/rke2.yaml
```
5. Проверить работоспособность кластера:
```bash
/var/lib/rancher/rke2/bin/kubectl get nodes  --kubeconfig /etc/rancher/rke2/rke2.yaml

/var/lib/rancher/rke2/bin/kubectl get pods -A --kubeconfig /etc/rancher/rke2/rke2.yaml


