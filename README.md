# SMB Server Boshrelease

## BOSH

1. Create and upload dev release: `bosh create-release --force && bosh upload-release`
2. Deploy:
```
bosh -d smbs-dev deploy manifests/smb-server.yml
  -o manifests/ops/set-external-access.yml  \
  -o manifests/ops/configure-shares-example.yml  \
  --var=deployment_name=smbs-dev
```

## Cloud Foundry

1. Deploy SMB Volume Service to Cloud Foundry https://docs.cloudfoundry.org/running/deploy-vol-services.html#smb-example
2. Enable it: `cf enable-service-access smb`
3. Create service: `cf create-service smb Existing smbs-dev-vol-a -c '{"share":"//IP-ADDRESS/vol-a"}'`
4. Push the test app [voltest](cf/voltest-app): `cf push --no-start`
5. Bind service with credentials: `cf bind-service voltest smbs-dev-vol-a -c '{"username":"user-a","password":"pass-a"}'`
6. Start the test app: `cf voltest start`
7. Navigate to the test app and test it.

## TODO

1. HA Cluster: https://wiki.samba.org/index.php/Configuring_clustered_Samba
2. NFS Server Boshrelease based on: https://github.com/nfs-ganesha/nfs-ganesha/wiki/NFS-Ganesha-and-High-Availability
