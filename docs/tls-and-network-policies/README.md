
# Securing the app with https
- install cert-manager
- Configure a CA Issuer
- Configure Ingress with the certificate obtained from the CA Issuer
- Verify the connection


## Steps

- Prereq: Install latest ingress controller, cert-manager CRDs and cert-manager

- run Helm chart
`helm install myapp .`

- retrieve the CA certificate secret that was created
`kubectl get secrets`

- To verify the validity of the certificate view the website or with the following openssl command
`openssl s_client -connect myeventsapp.com:443 -servername myeventsapp.com`

### To Trust our self-signed Certificate Authority (CA) on Mac to get rid of browser warning extract the CA certificate from the secret to a PEM file
`kubectl get secret myapp-helmcharts-ca-certificate-secret -n default -o jsonpath="{.data['tls\.crt']}" | base64 --decode > myCA.pem`

- add the CA certificate to the trusted root certificates in your keychain
`sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain myCA.pem`


# Network Policies

Check for allowed networkpolicies:
`kubectl get networkpolicies --all-namespaces`

1- denyExternalDbAccess

`helm upgrade myapp . --set networkPolicies.denyExternalDbAccess.enabled=true`

- Try to connect to database from inside the api pod. It'd fail: 
`kubectl get pods`  
`kubectl exec -it myapp-helmcharts-api-7884d6c575-qlqvq  -- /bin/sh `

`PGPASSWORD=cG9zdGdyZXM= psql -h myapp-helmcharts-database-cluster-ip-service -U cG9zdGdyZXM= -d Events`

Network policies are additive apply the second network policy. api pod connection to the database is successful

2- allowApiToDb

`helm upgrade myapp . --set networkPolicies.allowApiToDb.enabled=true`

`kubectl get pods`  
`kubectl exec -it myapp-helmcharts-api-7884d6c575-qlqvq  -- /bin/sh `

`PGPASSWORD=cG9zdGdyZXM= psql -h myapp-helmcharts-database-cluster-ip-service -U cG9zdGdyZXM= -d Events`

Connected.

3- allowFrontendToApi

- Go inside a frontend pod and see that it connects
`kubectl exec -it myapp-helmcharts-frontend-65875f8989-4xfwh -- /bin/sh`
Execute the command
`curl -v http://api-cluster-ip-service/api/events`


- Create a test pod and exec into it 
`kubectl run curlpod --image=alpine --restart=Never --rm -it -- /bin/sh`

`apk add curl`

`curl -v http://api-cluster-ip-service/api/events`

Fails.

