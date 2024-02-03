# install, scale, and uninstall commands

The following commands are used to install, scale, and uninstall the application on GKE.

```bash

docker login ghcr.io
cd ./frontend
docker build -t ghcr.io/noleu/event_calendar_frontend:v1 .
docker push ghcr.io/noleu/event_calendar_frontend:v1

gcloud container clusters get-credentials cluster-1 --region=europe-west1-b
kubectl config current-context
helm install my-events-app .\helmcharts\
kubectl get deployments


helm upgrade my-events-app .\helmcharts\ --set api.replicas=3
kubectl get deployments
helm uninstall my-events-app
kubectl get deployments
```
The following commands are executed when setting up the cluster, not part of the presentation.

```bash
kubectl create namespace cert-manager

helm repo add jetstack https://charts.jetstack.io

helm repo update

helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.8.0 --set installCRDs=true

kubectl create namespace ingress-nginx

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml

helm install myapp.

```