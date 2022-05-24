## Fabric Ca Demo
This repo demonstrates how to use fabric-ca to create certificates.
Happy joining.

## Identity managerment
![identity](./doc/d01.png)

## Certificate enrollment
![enroll](doc/d02.png)

## Affiliation managerment
![affi](doc/d03.png)

## Idmix credentials
![idmix](doc/d04.png)

## Certificate store structure
Certificate stores at ./cahome/organizations

## Config
See ./config.yarml for more infomation. At this time, pg is required.

## About key store
The user private key is stored in the client, not in the Fabric-CA. The client never shares its private key (as it is expected to). The client generates a key pair, submits a CSR to the CA and receives the signed certificate. The client saves the received certificate (the CA saves a copy, too)