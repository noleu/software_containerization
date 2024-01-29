$primary_account="noleu66@googlemail.com"
alias a-kubectl='kubectl --token="$(gcloud auth print-access-token --account=$primary_account)"'

gcloud auth list
gcloud config set account $primary_account
gcloud iam service-accounts create cluster-user-2 --display-name=cluster-user-2
$account2=$(gcloud iam service-accounts list --format='value(email)' --filter='displayName:cluster-user-2')
gcloud iam service-accounts keys create --iam-account $account2 cluster-user-2.json
gcloud auth activate-service-account $account2 --key-file=cluster-user-2.json
Set-Alias -Name "two-kubectl" -Value 'kubectl --token="$(gcloud auth print-access-token --account=$account2)"'
gcloud config set account $primary_account