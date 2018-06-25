**Copyright (c) 2018 Chen Xi. All Rights Reserved.**

## openssl private encrypt with RSA pkcs1.5 in Golang

### Input

- message you want to encrypt
- RSA private key in pkcs8.

This function equals to:


Bash:
```
$ openssl rsautl -sign -inkey "YOUR_RSA_PRIVATE_KEY" -in "YOUR_SIGN_CONTENT" | base64
```

python:
```
M2Crypto.RSA.private_encrypt(message, RSA.pkcs1_padding)
```

Javescript:
```JS
const NodeRSA = require('node-rsa');
const key = new NodeRSA('RSA_PRIVATE_KEY');
key.encryptPrivate(message, 'base64');
```
