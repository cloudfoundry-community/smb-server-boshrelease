# SMB Server Boshrelease

## Bosh Deploy

```
profile=dev
deployment_name=smbs-$profile

bosh -d smbs-dev deploy manifests/smb-server.yml
  -o manifests/ops/set-external-access.yml  \
  -o manifests/ops/configure-shares-example.yml  \
  --var=deployment_name=smbs-dev
```

## Cloud Foundry

1. Deploy SMB Volume Service to Cloud Foundry https://docs.cloudfoundry.org/running/deploy-vol-services.html#smb-example
2. Enable it: `cf enable-service-access smb`
3. Create service: `cf create-service smb Existing smbs-dev-vol-a -c '{"share":"//IP-ADDRESS/vol-a"}'`
4. Push app [voltest](cf/voltest-app): `cf push --no-start`
5. Bind service with credentials: `cf bind-service voltest smbs-dev-vol-a -c '{"username":"user-a","password":"pass-a"}'`
6. Start the app: `cf voltest start`
7. Navigate to the app and test it.
