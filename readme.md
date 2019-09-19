# SMB Server Boshrelease

## Deploy

```
profile=dev
deployment_name=smbs-$profile

bosh -d ${deployment_name} deploy \
  manifests/smb-server.yml 
  -o manifests/ops/set-external-access.yml  \
  -o manifest/ops/configure-shares-example.yml  \
  --var=deployment_name=${deployment_name} \
  --vars-file=vars/${profile}-vars-file.yml \
  --vars-store=vars/${profile}-vars-store.yml \
  --no-redact --fix

```

