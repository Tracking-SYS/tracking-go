# ipfs commandline tool

This is the [ipfs](http://ipfs.io) commandline tool. It contains a full ipfs node.

## Install

To install it, move the binary somewhere in your `$PATH`:

```sh
sudo mv ipfs /usr/local/bin/ipfs
```

Or run `sudo ./install.sh` which does this for you.

## Usage

First, you must initialize your local ipfs node:

```sh
ipfs init
```

This will give you directions to get started with ipfs.
You can always get help with:

```sh
ipfs --help
```

## Pinata
 
 1. Add Pinata credentials
 ```
 ipfs pin remote service add pinata https://api.pinata.cloud/psa {YOUR_JWT}
 ```

 2. pin a CID to Pinata
 ```
 ipfs pin remote add --service=pinata --name={file name} {CID}
 ```

 3. list successful pins
 ```
 ipfs pin remote ls --service=pinata
 ```

 4. list pending pins
 ```
 ipfs pin remote ls --service=pinata --status=queued,pinning,failed

 ```

 [Documentation](https://pinata.cloud/documentation#PinningServicesAPI)
