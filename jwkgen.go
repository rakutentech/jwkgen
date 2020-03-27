package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var (
	buildDate string
	version   string

	bareOutput  = false
	allowUnsafe = kingpin.
			Flag("allow-unsafe", "Allow unsafe parameters").Bool()
	useColor = kingpin.
			Flag("color", "Use color in JSON output (true by default)").
			Default("true").Bool()
	curve = kingpin.
		Flag("curve", "Named elliptic curve to use to generate a key. Valid values are P256, P384, P521, X25519, Ed25519").
		Short('e').Default("Ed25519").String()
	bits = kingpin.
		Flag("bits", "Number of bits to use for RSA or octet keys").
		Short('b').Default("2048").Int()
	onlyPEM  = kingpin.Flag("pem", "Print only PEM format").Bool()
	onlyJWK  = kingpin.Flag("jwk", "Print only JWK format").Bool()
	keyType  = kingpin.Arg("key type", "Key type: oct, rsa, ec").Default("ec").Enum("oct", "rsa", "ec")
	filename = kingpin.Arg("filename", "Output filename (without extension)").String()
)

func main() {
	if version == "" {
		kingpin.Version("JWK Generator (jwkgen) development snapshot")
	} else {
		kingpin.Version(fmt.Sprintf("JWK Generator (jwkgen) version %s\nBuild date: %s", version, buildDate))
	}
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if *keyType == "oct" {
		printSymmetricKey()
	} else {
		printKeyPair()
	}

	if !bareOutput {
		fmt.Println()
	}
}

func printSymmetricKey() {
	if *onlyPEM {
		return // Write nothing
	}
	bareOutput = true
	key := generateOctKey()
	err := writeJSONFor(ObjectInfo{".json", "Key (JWK)"}, &key)
	checkKeyError(err)
}

func printKeyPair() {
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
}

func checkKeyError(err error) {
	if err != nil {
		log.Fatalf("Could not write key: %v", err)
	}
}
