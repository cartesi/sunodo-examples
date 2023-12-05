# Echo Python dApp

This dApp represents a minimalistic Cartesi Rollups application that simply copies (or "echoes") each input received as a corresponding output notice. This dApp's back-end is written in Python.

## Building the application

To build the application, run:

```
sunodo build
```

A Cartesi Machine snapshot will be created:

```
         .
        / \
      /    \
\---/---\  /----\
 \       X       \
  \----/  \---/---\
       \    / CARTESI
        \ /   MACHINE
         '

[INFO  rollup_http_server] starting http dispatcher service...
[INFO  rollup_http_server::http_service] starting http dispatcher http service!
[INFO  actix_server::builder] starting 1 workers
[INFO  actix_server::server] Actix runtime found; starting in Actix runtime
[INFO  rollup_http_server::dapp_process] starting dapp: python3 dapp.py
INFO:__main__:HTTP rollup_server url is http://127.0.0.1:5004
INFO:__main__:Sending finish

Manual yield rx-accepted (0x100000000 data)
Cycles: 2801607569
2801607569: 949fc4c3fc34333a0add4c1f367d2cc1d8c802c3ab3e5b96a78a116b78070dba
Storing machine: please wait
```

## Running the application

This executes a Cartesi node for the application previously built with `sunodo build`.

```
sunodo run
```

It should print this output:

```
 ✔ Network 949fc4c3_default                       Created                                                                                                                                              0.0s
 ✔ Volume "949fc4c3_blockchain-data"              Created                                                                                                                                              0.0s
 ✔ Volume "949fc4c3_traefik-conf"                 Created
949fc4c3-prompt-1     | Anvil running at http://localhost:8545
949fc4c3-prompt-1     | GraphQL running at http://localhost:8080/graphql
949fc4c3-prompt-1     | Inspect running at http://localhost:8080/inspect/
949fc4c3-prompt-1     | Explorer running at http://localhost:8080/explorer/
949fc4c3-prompt-1     | Press Ctrl+C to stop the node
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
? Input (as string) Hello world, this is the echo-python dApp!
✔ Input sent: 0xd30150ee888a2bbf6b491812ee9ca28cb5754381eba3415ce4087322768c191f
```

Check [this documentation](../README.md/#sending-inputs-to-running-applications) for a comprehensive list of input types that a dApp can receive.
