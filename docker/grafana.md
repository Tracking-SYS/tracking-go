# GRAFANA BOARD SETUP

## Query setup 

1. Network IO (Bytes transfer) *(unit: megabytes)*
```
sum by (direction) (increase(system_network_io[1m])) / 1000000
```
2. System Memory Usage *(unit: megabytes)*
```
avg_over_time(system_memory_usage[1m]) / 1000000
```
3. API Latency - TP95 *(unit: milliseconds)*
```
histogram_quantile(0.95, rate(http_server_requests_duration_seconds_bucket[1m])) * 1000
```
4. API Throughput *(unit: rpm)*
```
increase(http_server_requests_number[1m])
```
5. New Product Creation
```
SELECT
  creation_time AS "time",
  id
FROM products
WHERE
  $__timeFilter(creation_time)
ORDER BY creation_time
```
6. Number Goroutine
```
runtime_go_goroutines
```


## Alert setting
1. Get webhook url from **Slack**
2. Create notification channel on **Grafana**
3. Set Alert rule to specific pannel on **Grafana** (_Alert tab_) 
