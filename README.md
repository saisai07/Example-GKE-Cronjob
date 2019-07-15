# GKE Cronjob Sample
## Overview
A sample of a Go project connecting to CloudSQL on the GKE Cronjob type

## Requirement
1. Create Cloud SQL Instance
2. Create GKE Clusters
3. Create Service Account and Download Json Key file

## Usage
### Create own envirement Type
1. Docker Container Up
```
docker-compose up -d
```

2. Login to My Contaier
```
docker exec -it {CONTAINER_ID} sh
```

3. Run Application
```
go run main.go
```

### Deploy To GKE
1. Create Container to Container Registry
  ```
  gcloud builds submit --tag=gcr.io/{PROJECT_ID}/batch:v1 .
  ```

2. Deploy to GKE
```
kubectl apply -f manifests/deployment.yaml
```

3. Check Pods
```
kubectl get pods
```

4. Check Logs
```
kubectl logs {CONTAINER_ID}
```