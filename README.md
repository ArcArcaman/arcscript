# ArcScript
Arcscript serves as a general Command Line Interface (CLI) tool for software developer.

# Installation
1. Install Golang from Golang's official website.
2. Clone this repo in your machine.
3. Run the following command in the directory:
    ```bash
    go install
    ``````
4. All done!

# Commands
## Convert
Converters that encode or decode strings to or from their specific formats.

### base64
Encode / Decode base64 strings.


## Random
Generates random data.

### string
Generates random strings. Constraints can be applied using the `charset` flag. Default is `none`.

### eddsa
Generates random EDDSA key pair. The results are generated as two PEM files. You can change the private key output file name and path with the `out` flag. Public key file name will always be `private_key_name.pub`. Default is `./eddsa_key`.