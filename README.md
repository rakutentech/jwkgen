# jwkgen - JSON Web Key Generator [![Go Report Card](https://goreportcard.com/badge/github.com/rakutentech/jwkgen)](https://goreportcard.com/report/github.com/rakutentech/jwkgen)

jwkgen is a small command-line tool that generates asymmetric JSON Web Keys for the
following algorithms and curves:

* RSA
* Curve25519 (ECDH, RFC 8037 compliant)
* Ed25519 (EdDSA, RFC 8037 compliant)
* P-256 (ECDSA and ECDH)
* P-384 (ECDSA and ECDH)
* P-521 (ECDSA and ECDH)

## Usage

**jwkgen [options] <key type> [filename]**

**-h, --help**

Show context-sensitive help (also try --help-long and --help-man).

**--allow-unsafe**

Allow unsafe parameters

**--color**

Use color in JSON output (true by default)

**-e, --curve="P256"**

Named elliptic curve to use to generate a key. Valid values are P256, P384, P521, X25519, Ed25519

**-b, --bits=2048**

Number of bits to use when generating RSA keys

**--pem**

Output only PEM format (useful for pipelining results and shell scripting)

**--jwk**

Output only JWK format (useful for pipelining results and shell scripting)

**--version**

Show jwkgen version.
