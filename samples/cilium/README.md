while true; do curl --location --request POST 'health.k8s.cluster.local:30080/calculator' \
--header 'Host: health.k8s.cluster.local' \
--header 'Content-Type: application/json' \
--data-raw '{ 
   "age": 26,
   "weight": 90.0,
   "height": 1.77,
   "gender": "M", 
   "activity_intensity": "very_active"
} ' --silent | jq . ; echo ; done; 