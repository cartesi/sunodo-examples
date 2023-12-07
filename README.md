# Sunodo dApp Examples

This repository includes examples of decentralized applications built with Sunodo.

Sunodo streamlines development across different language ecosystems, eliminating the need to navigate intricate configuration settings.

With Sunodo, you can bootstrap a complete Cartesi machine template in your programming language of choice through the CLI without getting lost in configurations.

## Why Sunodo?

- **Building Applications:** Sunodo simplifies the process of compiling applications from source code into a Cartesi Machine.
- **Local Development Environment:** Run, test, and debug applications in a local development environment.
- **Deployment:** A simplified deployment process with options to deploy applications on both public and private networks.
- **Monitoring:** Keep a close eye on the status and performance of applications that are already up and running.

## Requirements and Installation

The CLI heavily uses Docker under the hood, so you must have it installed and up-to-date.

The recommended way to have all plugins ready for building is to install [Docker Desktop](https://www.docker.com/products/docker-desktop/).

### macOS

If you have [Homebrew](https://brew.sh/) installed, you can install Sunodo by running this command:

```bash
brew install sunodo/tap/sunodo
```

Alternatively, you can install Sunodo with Node.js by running:

```bash
npm install -g @sunodo/cli
```

### Linux

You can either use [Homebrew on Linux](https://docs.brew.sh/Homebrew-on-Linux), or install Sunodo with:

```
npm install -g @sunodo/cli
```

### Windows

Install [WSL2](https://learn.microsoft.com/en-us/windows/wsl/install) and the Ubuntu dsitro from Microsoft Store and install Sunodo with:

```
npm install -g @sunodo/cli
```


## Creating an application

Use the `sunodo create` command to quickly start a Cartesi dApp from scratch. It sets up everything you need with template code.

Here are the available templates:

- `cpp`: A template for C++ development.
- `cpp-low-level`: C++ template using the low level API, instead of the HTTP server
- `go`: Go lang template
- `javascript`: A node.js 20 template tailored for JavaScript developers
- `lua`: Lua 5.4 template
- `python`: python 3 template
- `ruby`: ruby template
- `rust`: rust template
- `typescript`: TypeScript template

To create a new application from a basic Python template, run:

```
> sunodo create dapp-name --template python
✔ Application created at dapp-name
```

## Building the application

To build an application, run:

```
sunodo build
```

When you run the `sunodo build` command:

- Your program's code gets compiled into the RISC-V architecture
- A Cartesi Machine is assembled.
- The end result of this process is a Cartesi Machine snapshot, ready to receive inputs.

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
[INFO  rollup_http_server::dapp_process] starting dapp
INFO:__main__:HTTP rollup_server url is http://127.0.0.1:5004
INFO:__main__:Sending finish

Manual yield rx-accepted (0x100000000 data)
Cycles: 2730825066
2730825066: a8a8ac815729549ca97aad13cdf10a5f51205f64e7f3fc77fc4e05050c49caee
Storing machine: please wait
```

## Running the application

This executes a Cartesi node for the application previously built with `sunodo build`.

```
sunodo run
```

The `sunodo run` command should print this output:

```
43272d70-prompt-1     | Anvil running at http://localhost:8545
43272d70-prompt-1     | GraphQL running at http://localhost:8080/graphql
43272d70-prompt-1     | Inspect running at http://localhost:8080/inspect/
43272d70-prompt-1     | Explorer running at http://localhost:8080/explorer/
43272d70-prompt-1     | Press Ctrl+C to stop the node
```


### No backend mode

Sunodo run also supports running a node without your application backend packaged as a Cartesi machine.

```
sunodo run --no-backend
```

In this case, your application can be executed on the host without being compiled to RISC-V. Here are some drawbacks with `--no-backend`:

- **Compilation Requirement**: Your application must eventually be compiled to RISC-V during deployment.

- **Sandbox Restrictions**: In `--no-backend` mode, the application won't run within the Cartesi machine's sandbox, enabling operations that are otherwise restricted.

- **API Compatibility**: This mode is compatible only with applications using the Cartesi Rollups HTTP API, not those using the low-level API like the [C++ low-level dApp](./echo-low-level/)

- **Performance Impact**: Expect lower performance inside a Cartesi machine compared to running on the host.

When launching a node with the `--no-backend` you must then start your application on the host and fetch inputs from the endpoint running at http://127.0.0.1:8080/host-runner.

### Verbosity

By default, the Cartesi node runs in non-verbose mode, providing logs only from your backend application. For more information, use the `--verbose` command option.

```
sunodo run --verbose
```

### Blockchain Configuration

Sunodo runs a local private chain powered by Anvil at port `8545`.

All contracts of the Cartesi Rollups framework are deployed and you can inspect their addresses by running:

```
sunodo address-book
```

The private chain has a default block time of 5 seconds and it runs on auto-mine mode.

You can manually configure block time by running:

```
sunodo run --block-time <seconds>
```

### Epoch Configuration

By default the node closes an epoch once a day.

You can manually configure epoch-duration by running:

```
sunodo run --epoch-duration <seconds>
```

### Rollups Node Configuration

You can configure Rollups Node services using environment variables.

Create a `.sunodo.env` file in your project's root.

If you wish to modify the default deadline for advancing the state in the rollups-advance-runner service, open the `.sunodo.env` file and add the desired configuration, for example:

```shell
SM_DEADLINE_ADVANCE_STATE=360000
```

## Sending inputs to running applications

Your applications can receive inputs by sending transactions with the input payload to the `InputBox` smart contracts of the Cartesi Rollups framework. Sunodo provides a convenient CLI command to facilitate this process.

To send inputs, use the command:

```
sunodo send
```

This command guides you through the process of sending inputs interactively.

```
? Select send sub-command (Use arrow keys)
❯ Send DApp address input to the application.
  Send ERC-20 deposit to the application.
  Send ERC-721 deposit to the application.
  Send ether deposit to the application.
  Send generic input to the application.
```

### Input Types

There are five types of inputs you can send using sub-commands:

#### 1. dApp Address

Useful for applications that need to know their own address. The input payload is the address of the application and the sender is the [`DAppAddressRelay`](https://github.com/cartesi/docs/blob/392265797d8cfebca7bb1ad7b63cbf9a0569fc12/cartesi-rollups/api/json-rpc/relays/DAppAddressRelay.md#L4) smart contract.

```
sunodo send dapp-address
```

#### 2. ERC-20 Deposit

Deposits ERC-20 tokens to the application. Refer to the [`ERC20Portal`](https://github.com/cartesi/docs/blob/392265797d8cfebca7bb1ad7b63cbf9a0569fc12/cartesi-rollups/api/json-rpc/portals/ERC20Portal.md) documentation for payload format details.

```
sunodo send erc20
```

#### 3. ERC-721 Deposit (NFT)

Deposits ERC-721 tokens (NFT) to the application. Refer to the [`ERC721Portal`](https://github.com/cartesi/docs/blob/392265797d8cfebca7bb1ad7b63cbf9a0569fc12/cartesi-rollups/api/json-rpc/portals/ERC721Portal.md) documentation for payload format details.

```
sunodo send erc721
```

#### 4. Ether Deposit

Deposits Ether (native token) to the application. Refer to the [`EtherPortal`](https://github.com/cartesi/docs/blob/392265797d8cfebca7bb1ad7b63cbf9a0569fc12/cartesi-rollups/api/json-rpc/portals/EtherPortal.md) documentation for payload format details.

```
sunodo send ether
```

#### 5. Generic Input:

Send inputs with any payload format. This is suitable for specialized input requirements.

```
sunodo send generic
```

The encoding of the payload can be specified with the `--input-encoding` option. Supported encodings include:

- **hex**: Parses user input as a hex-string and converts it to bytes.
- **string**: Converts user input from a UTF-8 string to bytes.
- **abi**: Allows the user to specify an ABI-encoded input format in a human-readable format.

Example:

```
sunodo send generic --input-encoding hex
```

## Example dApps

### 1. [Echo Python dApp](./echo-python)

A basic "hello world" dApp written in Python that simply copies each input received as a corresponding output notice.

### 2. [Echo C++ dApp](./echo-cpp)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in C++.

### 3. [Echo Rust ](./echo-rust)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in Rust.

### 4. [Echo Lua dApp](./echo-lua)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in Lua.

### 5. [Echo JS dApp](./echo-js)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in JavaScript.

### 6. [Echo TS dApp](./echo-ts)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in TypeScript.

### 7. [Echo Low-Level dApp](./echo-low-level)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in C++ using the low-level Cartesi Rollups API.

### 8. [Echo Ruby dApp](./echo-ruby)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in Ruby.

### 9. [Echo Go dApp](./echo-go)

Implements the same behavior as the [Echo Python dApp](#2-echo-python-dapp) above, but with a back-end written in Go.
