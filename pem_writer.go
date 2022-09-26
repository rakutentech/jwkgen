package main

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
	"reflect"

	"github.com/pkg/errors"
	"github.com/rakutentech/jwk-go/okp"
)

func pemBlockFor(obj interface{}) (*pem.Block, error) {
	var err error
	var der []byte
	switch o := obj.(type) {
	case *rsa.PrivateKey:
		switch *rsaKeyFormat {
		case "PKCS1":
			return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(o)}, nil
		case "PKCS8":
			der, err := x509.MarshalPKCS8PrivateKey(o)
			if err != nil {
				return nil, errors.Wrap(err, "Unable to marshal RSA private key")
			}
			return &pem.Block{Type: "PRIVATE KEY", Bytes: der}, nil
		default:
			return nil, errors.Errorf("Unknown key format: %v", reflect.TypeOf(obj))
		}
	case *ecdsa.PrivateKey:
		der, err = x509.MarshalECPrivateKey(o)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to marshal ECDSA private key")
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: der}, nil
	case *rsa.PublicKey, *ecdsa.PublicKey:
		der, err = x509.MarshalPKIXPublicKey(obj)
		if err != nil {
			return nil, errors.Wrap(err, "Unable to marshal RSA private key")
		}
		return &pem.Block{Type: "PUBLIC KEY", Bytes: der}, nil
	case okp.CurveOctetKeyPair:
		// ASN.1 representation for PEM is currently an IETF dratf:
		// https://tools.ietf.org/id/draft-ietf-curdle-pkix-04.txt
		// It's bothersome to implement so I skip it for now
		return nil, nil
	default:
		return nil, errors.Errorf("Unknown key type: %v", reflect.TypeOf(obj))
	}
}

func writePem(writer io.Writer, obj interface{}) error {
	block, err := pemBlockFor(obj)
	if err != nil {
		return err
	} else if block == nil {
		return nil // Skipped PEM for this type
	}
	return pem.Encode(writer, block)
}

func writePemFor(objInfo ObjectInfo, obj interface{}) error {
	w, err := writerFor(objInfo)
	if err != nil {
		return err
	}
	return writePem(w, obj)
}
