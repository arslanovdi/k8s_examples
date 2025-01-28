
## kube-state-metrics

сбор метрик из k8s и передача в prometheus

prometheus конфиг
```yaml
---
scrape_configs:
  - job_name: 'kube-state-metrics'
    static_configs:
      - targets: ['kube-state-metrics.default.svc.cluster.local:8080']
---
```