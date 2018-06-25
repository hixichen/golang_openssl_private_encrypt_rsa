// Copyright (c) 2018,  xichen0425@gmail.com. All Rights Reserved.
//

package chenxi_package

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// RSA PKCS8 private key format
// "
// -----BEGIN RSA PRIVATE KEY-----
//	MIIE****
// -----END RSA PRIVATE KEY-----
// "
//
// var inputPrivateKey string

func signWithRsaPrivateKey(message, inputPrivateKey string) (string, error) {

	block, _ := pem.Decode([]byte(inputPrivateKey))
	if block == nil {
		return "", fmt.Errorf("no key found")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("unsupported key type %q", block.Type)
	}

	rsaPrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Sign secret with rsa with PKCS 1.5 as the padding algorithm
	// The result is exactly same as "openssl rsautl -sign -inkey "inputPrivateKey_in_file" -in "YOUR_PLAIN_TEXT""
	signer, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey.(*rsa.PrivateKey), crypto.Hash(0), []byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signer), nil
}
