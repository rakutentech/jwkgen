package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"github.com/pkg/errors"
	"github.com/rakutentech/jwk-go/jwk"
	"github.com/rakutentech/jwk-go/okp"
	"log"
	"os"
)

func generateKeyPair() (interface{}, interface{}, *jwk.KeySpec, *jwk.KeySpec) {
	priv, err := generateKey()
	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
		os.Exit(1)
	}
	pub := publicKey(priv)
	if pub == nil {
		log.Fatal("failed to deduce public key from private key")
	}
	privJwk := jwk.NewSpec(priv)
	if err := privJwk.Normalize(jwk.NormalizationSettings{
		RequireKeyID: true,
	}); err != nil {
		panic(err)
	}
	pubJwk := jwk.NewSpec(pub)
	if err := pubJwk.Normalize(jwk.NormalizationSettings{
		RequireKeyID: true,
	}); err != nil {
		panic(err)
	}

	return priv, pub, privJwk, pubJwk
}

func generateKey() (interface{}, error) {
	switch *keyType {
	case "ec":
		return generateECKey()
	case "rsa":
		return generateRSAKey()
	default:
		return nil, errors.Errorf("Unknown key type: %s", *keyType)
	}
}

func generateECKey() (interface{}, error) {
	switch *curve {
	case "P256", "P-256":
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384", "P-384":
		return ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521", "P-521":
		return ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	case "Ed25519":
		return okp.GenerateEd25519(rand.Reader)
	case "X25519":
		return okp.GenerateCurve25519(rand.Reader)
	default:
		return nil, errors.Errorf("Unknown Elliptic Curve: %s", *curve)
	}
}

func generateRSAKey() (interface{}, error) {
	if *rsaBits < 512 || *rsaBits > 8192 {
		return nil, errors.Errorf("Invalid RSA key size: %d", *rsaBits)
	} else if *rsaBits < 2048 && !*allowUnsafe {
		return nil, errors.Errorf("RSA key size (%d) is too small. NIST recommends at least 2048 bits.", *rsaBits)
	}
	return rsa.GenerateKey(rand.Reader, *rsaBits)
}

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case okp.CurveOctetKeyPair:
		pubKey, err := okp.CurveExtractPublic(k)
		if err != nil {
			panic(err)
		}
		return pubKey
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}
