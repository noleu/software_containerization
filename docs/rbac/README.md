# RBAC

- using RBAC to control access to the Kubernetes API
- groups are bound to roles
- 3 roles
  - app developer
    - may handle everything related to the application, like reading logs or executing commands in pods
  - infrastructure admin
    - may handle everything related to the infrastructure of the application, like configuring configs and network policies
  - application admin
    - may handle everything in the namespace
- roles are bound to resources, since label filtering is not possible for roles
- for each role a service account is created
- the token can be retrieved like this

`kubectl --token=$(kubectl create token infrastructure-admin) get pods`

## useful links
[tutorial](https://gcloud.devoteam.com/blog/the-ultimate-security-guide-to-rbac-on-google-kubernetes-engine/)  
[authentication with google](https://cloud.google.com/kubernetes-engine/docs/how-to/api-server-authentication#authenticating_using_oauth)  
[minikube google auth](https://minikube.sigs.k8s.io/docs/handbook/addons/gcp-auth/)