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

### CF Operator
1. Deploy SMB Volume Service to Cloud Foundry https://docs.cloudfoundry.org/running/deploy-vol-services.html#smb-example
2. Enable it: `cf enable-service-access smb`

### Developer 
- Create service: 
  ```
  cf create-service smb Existing smbs-dev-share-a -c '{"share":"//IP-ADDRESS/share-a"}'
  ```
- Push the app e.g. [voltest](cf/voltest-app): `cf push --no-start`
- Bind service with credentials to the app: 
  ```
  cf bind-service voltest smbs-dev-share-a -c '{"username":"user-a","password":"pass-a", "mount":"/var/path"}'
  ```
- Start the app `cf voltest start` and test it.

## Windows

- create connection: `net use \\<smb-server-ip>\<share-name> /USER:<user-name>`
- list connections: `net use`
- delete connection: `net use \\<smb-server-ip>\<share-name> /delete # delete all connections: net use * /delete`

## TODO

1. HA Cluster: https://wiki.samba.org/index.php/Configuring_clustered_Samba
2. NFS Server Boshrelease based on: https://github.com/nfs-ganesha/nfs-ganesha/wiki/NFS-Ganesha-and-High-Availability
