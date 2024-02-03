# Helm instructions 

## Helm create
We used helm create to create the helm chart. 
```bash
helm create helmcharts
```

## Helm install
```bash
helm install myapp helmcharts
```

## Helm uninstall
```bash
helm uninstall myapp
```

## Helm upgrade
```bash
helm upgrade myapp helmcharts
```

## Helm rollout
Update values.yaml frontend image tag to rollout and run:
```bash
helm upgrade myapp helmcharts
```

## Helm canary
Update values.yaml frontend image tag to canary, and change the replica count to for canary-frontend to 1. The other replicas are set to 2. After run:
```bash
helm upgrade myapp helmcharts
```