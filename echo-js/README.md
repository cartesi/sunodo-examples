# Echo JavaScript dApp

This is a template for JavaScript Cartesi DApps. It uses node to execute the backend application.
The application entrypoint is the `src/index.js` file. It is bundled with [esbuild](https://esbuild.github.io), but any bundler can be used.



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

For local testing, select `Foundary` which gives you mock and test faucets to submit transactions:

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
? Input (as string) Hello world, this is the echo-js dApp!
✔ Input sent: 0xd30150ee888a2bbf6b491812ee9ca28cb5754381eba3415ce4087322768c191f
```


Check [this documentation](../README.md/#sending-inputs-to-running-applications) for a comprehensive list of input types that a dApp can receive.