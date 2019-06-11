# jwkgen - JSON Web Key Generator [![Go Report Card](https://goreportcard.com/badge/github.com/rakutentech/jwkgen)](https://goreportcard.com/report/github.com/rakutentech/jwkgen)

## Overview

jwkgen is a small command-line tool that generates asymmetric JSON Web Keys for the
following algorithms and curves:

* RSA
* Curve25519 (ECDH, RFC 8037 compliant)
* Ed25519 (EdDSA, RFC 8037 compliant)
* P-256 (ECDSA and ECDH)
* P-384 (ECDSA and ECDH)
* P-521 (ECDSA and ECDH)

## Installation


On Mac, you can just use Homebrew:
```sh
> brew tap rakutentech/tap
> brew install jwkgen
```
On other platforms, you can just download the [latest release
archive](https://github.com/rakutentech/jwkgen/releases/latest) for your
platform and extract the binary to any location.

If you have Go installed, you can also install the latest version from master
branch:
```sh
> go get -u github.com/rakutentech/jwkgen
```

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
