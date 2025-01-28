## fluentbit

Собирает логи с подов и передает их в систему логирования.
Graylog в этом случае.

### install

```bash
helm repo add fluent https://fluent.github.io/helm-charts
helm upgrade --install fluent-bit fluent/fluent-bit -n observability
```
