# Рекомендации Vertical Pod Autoscaler (VPA)

Можно стресс тест сервиса устроить и посмотреть сколько требуется ресурсов, потом уже эти цифры заложить в деплой.

Нагрузить сервис:
```bash
kubectl run -i --tty load-generator --rm --namespace=stateful-todo --image=busybox:1.28 --restart=Never -- /bin/sh -c "while sleep 0.01; do wget -q -O- stateful-todo-clusterip-service:5000/tasks; done"
```


Рекомендации по лимитам.
```bash
kubectl -n stateful-todo describe vpa stateful-todo-vpa
```

Текущие параметры смотреть в describe пода.