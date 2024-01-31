
# Securing the app with https
- install cert-manager
- Configure a CA Issuer
- Configure Ingress with the certificate obtained from the CA Issuer
- Verify the connection with openssl


## Steps
- Install cert-manager CRDs and cert-manager with the command:
`kubectl create namespace cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update
helm install cert-manager jetstack/cert-manager --namespace cert-manager --version v1.8.0 --set installCRDs=true`

- run Helm chart
`helm install myapp .`

- Verify the connection with openssl
`openssl s_client -connect myeventsapp.com:443 -servername myeventsapp.com`


### Trust our self-signed Certificate Authority (CA) on Mac to get rid of browser warning

- Extract the CA certificate from the secret in our Kubernetes cluster
`kubectl get secret ca-certificate-secret -n default -o jsonpath="{.data['tls\.crt']}" | base64 --decode > myCA.pem`

Open Keychain Access on your mac: Go to Applications > Utilities > Keychain Access.

Select the System Keychain: In the Keychain Access sidebar, select "System" under the Keychains section.

Import the Certificate:

Drag and drop the myCA.pem file into the Keychain Access window, or use File > Import Items and select the "System" keychain as the destination.
Authenticate as Administrator: You'll be prompted to authenticate with an administrator username and password. This is required because you're modifying the System keychain.

Trust the Certificate:

Find your certificate in the list in the System keychain.
Double-click on it to open the certificate details.
Expand the "Trust" section.
Change "When using this certificate" to "Always Trust".
Close the certificate window, and you'll be prompted for your administrator password again to save the changes.
Restart Your Browser or System: After making these changes, you might need to restart your browser or other applications to recognize the new trust settings. In some cases, a system restart may be required for all changes to take effect.


# Network Policies
- allow-api-to-db-networkpolicy.yaml
How to test:

- allow-frontend-to-api-networkpolicy.yaml
How to test:

- deny-external-db-access-networkpolicy.yaml
How to test:
