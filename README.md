# nutrition-overengineering
Application to calc healthcare and nutrition metrics, but using the most over engineered as possible

> "Caraca Matheus, tu é idiota??? 
> "Sim"

# API  

## Usage 

```bash
curl --location --request POST '0.0.0.0:8080/calculator' \
--header 'Content-Type: application/json' \
--data-raw '{ 
   "age": 26,
   "weight": 90.0,
   "height": 1.77,
   "gender": "M", 
   "activity_intensity": "very_active"
} ' --silent | jq .
```

```json
{
    "status": 200,
    "imc": {
        "result": 28.72737719046251,
        "class": "overweight"
    },
    "basal": {
        "bmr": {
            "value": 2011.7,
            "unit": "kcal"
        },
        "necessity": {
            "value": 3470.1825000000003,
            "unit": "kcal"
        }
    },
    "health_info": {
        "age": 26,
        "weight": 90,
        "height": 1.77,
        "gender": "M",
        "activity_intensity": "very_active"
    },
    "recomendations": {
        "protein": {
            "value": 180,
            "unit": "g"
        },
        "water": {
            "value": 3150,
            "unit": "ml"
        },
        "calories": {
            "maintain_weight": {
                "value": 3470.1825000000003,
                "unit": "kcal"
            },
            "loss_weight": {
                "value": 3123.1642500000003,
                "unit": "kcal"
            },
            "gain_weight": {
                "value": 5205.27375,
                "unit": "kcal"
            }
        }
    }
}
```

# TODO
* Create communication across services 
* Refactor communication lib 
* Customize log lib for trace id
* Implement Prometheus HTTP Endpoint
* Implement Jaeger Tracing
* Implement CI / CD 
* Create Kubernetes Deployment

## Notes 

```
brew install grpc protobuf protoc-gen-go-grpc
protoc --proto_path=proto/bmr --go_out=proto/bmr --go_out=plugins=grpc:bmr --go_opt=paths=source_relative proto/bmr/bmr.proto
protoc --go_out=plugins=grpc:bmr bmr.proto
```