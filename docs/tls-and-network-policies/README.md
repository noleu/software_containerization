
# Securing the app with https
- install cert-manager
- Configure a CA Issuer
- Configure Ingress with the certificate obtained from the CA Issuer
- Verify the connection with openssl


## Steps
- Install latest ingress controller with the command:
`kubectl create namespace ingress-nginx
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
`
- Install cert-manager CRDs and cert-manager with the command:
`kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.8.0 --set installCRDs=true`

- run Helm chart
`helm install myapp .`

- Wait on myapp-ingress-service address to populate and check it with the command: `kubectl get ingress`

- Verify the connection with openssl
`openssl s_client -connect myeventsapp.com:443 -servername myeventsapp.com`


### Trust our self-signed Certificate Authority (CA) on Mac to get rid of browser warning

- Extract the CA certificate from the secret in our Kubernetes cluster
`kubectl get secret ca-certificate-secret -n default -o jsonpath="{.data['tls\.crt']}" | base64 --decode > myCA.pem`

- This command adds the certificate to the System keychain and marks it as trusted for SSL:
`sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain myCA.pem`



# Network Policies
- allow-api-to-db-networkpolicy.yaml
How to test:

- allow-frontend-to-api-networkpolicy.yaml
How to test:

- deny-external-db-access-networkpolicy.yaml
How to test:
