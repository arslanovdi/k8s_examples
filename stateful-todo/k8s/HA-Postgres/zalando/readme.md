# [Zalando postgres operator](https://github.com/zalando/postgres-operator)

Оператор Postgres, который создает и управляет кластерами PostgreSQL, работающими в Kubernetes.


### [quickstart](https://github.com/zalando/postgres-operator/blob/master/docs/quickstart.md)

## Установка при помощи helm

```bash
# add repo for postgres-operator
helm repo add postgres-operator-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator

# install the postgres-operator
helm install postgres-operator postgres-operator-charts/postgres-operator

# add repo for postgres-operator-ui
helm repo add postgres-operator-ui-charts https://opensource.zalando.com/postgres-operator/charts/postgres-operator-ui

# install the postgres-operator-ui
helm install postgres-operator-ui postgres-operator-ui-charts/postgres-operator-ui
```

## Доступ к postgres-operator-ui
Проброс порта:
```bash
kubectl port-forward svc/postgres-operator-ui 8081:80
```

Доступ по адресу: `localhost:8081`

## Развертывание HA-Postgres
При развертывании автоматически создается секрет с username/password. По овнеру базы данных указанному в манифесте.
Базовый манифест можно накидать через postgres-operator-ui.