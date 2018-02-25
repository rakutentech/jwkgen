package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var (
	bareOutput  = false
	allowUnsafe = kingpin.
			Flag("allow-unsafe", "Allow unsafe parameters").Bool()
	color = kingpin.
		Flag("color", "Use color in JSON output (true by default)").
		Default("true").Bool()
	curve = kingpin.
		Flag("curve", "Named elliptic curve to use to generate a key. Valid values are P256, P384, P521, X25519, Ed25519").
		Short('e').Default("P256").String()
	rsaBits = kingpin.
		Flag("bits", "Number of bits to use for RSA keys").
		Short('b').Default("2048").Int()
	onlyPEM  = kingpin.Flag("pem", "Print only PEM format").Bool()
	onlyJWK  = kingpin.Flag("jwk", "Print only JWK format").Bool()
	keyType  = kingpin.Arg("key type", "Key type: rsa, ec").Default("ec").String()
	filename = kingpin.Arg("filename", "Output filename (without extension)").String()
)

func main() {
	kingpin.Version("1.3.0")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	priv, pub, privJwk, pubJwk := generateKeyPair()

	bareOutput = *filename == "" && (*onlyJWK || *onlyPEM)
	writePublic := !bareOutput
	writeJWK := !*onlyPEM
	writePEM := !*onlyJWK

	var err error
	if writeJWK {
		err = writeJSONFor(ObjectInfo{".json", "Private Key (JWK)"}, &privJwk)
		checkKeyError(err)
	}
	if writePublic {
		err = writeJSONFor(ObjectInfo{".pub.json", "Public Key (JWK)"}, &pubJwk)
		checkKeyError(err)
	}
	if writePEM {
		err = writePemFor(ObjectInfo{".pem", "Private Key (PEM)"}, priv)
		checkKeyError(err)
	}
	if writePublic {
		err = writePemFor(ObjectInfo{".pub.pem", "Public Key (PEM)"}, pub)
		checkKeyError(err)
	}

	if !bareOutput {
		fmt.Println()
	}
}

func checkKeyError(err error) {
	if err != nil {
		log.Fatalf("Could not write key: %v", err)
	}
}
