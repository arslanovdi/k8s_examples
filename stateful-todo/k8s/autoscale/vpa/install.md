# [Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/docs/installation.md)

```bash
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler
./hack/vpa-up.sh
```

Почему-то из WSL скрипт не работал.
Установил с ноды k8s.

vpa-admission-controller выдавал ошибку сертификата при установке.
Выаполнить:

```bash
bash ./pkg/admission-controller/gencerts.sh
```

и убить под vpa-admission-controller, чтобы пересоздался.