# RBAC

- using RBAC to control access to the Kubernetes API
- groups are bound to roles
- 3 roles
  - frontend-developer
    - may handle everything related to the frontend
  - backend-developer
    - may handle everything related to the backend and the database
  - application admin
    - may handle everything in the namespace
- roles are bound to resources, since label filtering is not possible for roles
- authentication is handled by google 
- groups are google IAM role

## useful links
[tutorial](https://gcloud.devoteam.com/blog/the-ultimate-security-guide-to-rbac-on-google-kubernetes-engine/)  
[authentication with google](https://cloud.google.com/kubernetes-engine/docs/how-to/api-server-authentication#authenticating_using_oauth)  
[minikube google auth](https://minikube.sigs.k8s.io/docs/handbook/addons/gcp-auth/)