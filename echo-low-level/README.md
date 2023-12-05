# Echo Low-Level dApp

This example implements the same behavior as the [Echo dApp written in Python](../echo-python/), but here the back-end is written in C++ and uses the low-level Cartesi Rollup API instead of the HTTP API.

As the other example, the dApp simply copies (or "echoes") each input received as a corresponding output notice.

## Building the application

To build the application, run:

```
sunodo build
```

## Running the application

This executes a Cartesi node for the application previously built with `sunodo build`.

```
sunodo run
```

## Interacting with the application

You can use the `sunodo send` command to send input payloads to your applications.

With your node running, open a new terminal tab. You can send a generic input to your application as follows:

```shell
 sunodo send generic
```

For local testing, select `Foundry` which gives you mock and test faucets to submit transactions:

```
> sunodo send generic
? Chain (Use arrow keys)
❯ Foundry
? Chain Foundry
? RPC URL http://127.0.0.1:8545
? Wallet Mnemonic
? Mnemonic test test test test test test test test test test test junk
? Account 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 9999.969170031383970357 ETH
? DApp address 0x70ac08179605AF2D9e75782b8DEcDD3c22aA4D0C
? Input String encoding
? Input (as string) Hello world, this is the echo-low-level dApp!
✔ Input sent: 0xd30150ee888a2bbf6b491812ee9ca28cb5754381eba3415ce4087322768c191f
```

Check [this documentation](../README.md/#sending-inputs-to-running-applications) for a comprehensive list of input types that a dApp can receive.
