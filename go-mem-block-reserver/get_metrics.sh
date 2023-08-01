#!/bin/bash

value=$(curl --silent --fail "http://127.0.0.1:8080/metrics")
for metric in "memstats.Sys" "memstats.TotalAlloc"; do
    metric_value=$(echo "$value" | jq -r ".$metric")
    printf "%s | %s: %s\n" "$(date --rfc-3339=ns)" "$metric" "$metric_value"
done

