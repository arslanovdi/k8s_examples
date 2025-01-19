# API Istio

В будущем API по умолчанию для управления трафиком будет Kubernetes Gateway API.
Этот вариант я пока не опробовал, пользуюсь Gateway API.


## [ingress2gateway](https://github.com/kubernetes-sigs/ingress2gateway)
Automatic Conversion of Ingresses

## Объект Gateway
Используется для конфигурации `ingress gateway`.

С помощью него в `istio ingress` можно открывать порты, добавлять TLS сертификаты, фильтровать запросы по именам хостов и т.д.

Пример:

```yaml
apiVersion: networking.istio.io/v1
kind: Gateway
metadata:
  name: istio-gateway
spec:
  selector:
    istio: ingressgateway   # use istio default ingress gateway
  servers:
    - port:
        number: 80
        name: http
        protocol: HTTP
      hosts:
        - "*"
```

`ingress gateway` в кластере может быть несколько, например какие-то за публичным loadbalancerом, а какие-то за приватным.
Соответственно и объектов `Gateway` может быть несколько. В таком случае прописывать имена хостов практически необходимо, чтобы не было путаницы.

## Объект VirtualService
Аналог обекта `ingress k8s`, но намного гибче.

- Конфигурирует L7 уровень.
- Настраивает роутинг на конкретные сервисы.
- Можено добавлять:
  - таймауты
  - ретраи
  - балансировать трафик по процентам

Здесь настраиваем куда дальше должен пойти трафик, пришедший на конкретный http endpoint.

Пример:

```yaml
kind: VirtualService
metadata:
  name: demo
spec:
  hosts:
    - "*"           # имена хостов
  gateways:
    - istio-gateway # имя gateway, с трафиком которого будем работать
  http:             # роутинг
    - match:        # все запросы
      - uri:
          prefix: /tasks    # с префиксом /tasks
      - uri:
          prefix: /users    # и с префиксом /users
      route:
      - destination:        # будут направлены
          host: demo        # в сервис demo, можно указать полное имя с namespace
          port:
            number: 8000    # на порт 8000

```

## Объект DestinationRule
- Настраивает поведение трафика после выполнения роутинга
- Группирует версии приложения в сабсеты
- Добавляет mTLS

Можно группировать поды по какому-то `label`, например все поды с версией v3 будут в группе с именем testversion.
Соответственно с помощью `VirtualService` можно будет навесить на эту группу процент трафика, что-то типа A/B тестов.

Также можно настраивать mTLS - взаимное шифрование трафика между подами.

Пример:

```yaml
apiVersion: networking.istio.io/v1
kind: DestinationRule
metadata:
  name: demo
spec:
  host: "*.foo.com"
  trafficPolicy:
    tls:
      mode: SIMPLE
  subsets:
  - name: testversion
    labels:
      version: v3
    trafficPolicy:
      loadBalancer:
        simple: ROUND_ROBIN
```

## Схема прохождения пакетов
- пакет -> 
- Внешний loadBalancer -> 
- порт kubernetes worker нод ->
- istio ingress gateway service ->
- istio ingress gateway pod (этот pod настраивается через объекты `Gateway` и `VirtualService`) ->
  - `Gateway` описывает порты, протоколы с SSL сертификатом и т.д.
  - `VirtualService` описывает роутинг пакета к нужному kubernetes сервису
- Service приложения -> 
- к запросу применяются правила, описанные в `DestinationRule`.
