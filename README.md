# Sunodo dApp Examples

This repository includes examples of decentralized applications  built with Sunodo.

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

Linux users can either use Homebrew on Linux, or install Node.js and then install sunodo with:

```
npm install -g @sunodo/cli
```

### Windows 

Install [WSL2](https://learn.microsoft.com/en-us/windows/wsl/install) and the Ubuntu dsitro from Microsoft Store and install Sunodo with:

```
npm install -g @sunodo/cli
```

## Building the application

To build an application simply run:

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


### No backend mode

Sunodo run also supports running a node without the user application backend packaged as a Cartesi machine. 

```
sunodo run --no-backend
```

In this case, your application can be executed on the host without being compiled to RISC-V. Here are some drawbacks with `--no-backend`:

- Compilation Requirement: Your application must eventually be compiled to RISC-V during deployment.

- Sandbox Restrictions: In `--no-backend` mode, the application won't run within the Cartesi machine's sandbox, enabling operations that are otherwise restricted.

- API Compatibility: This mode is compatible only with applications using the Cartesi Rollups HTTP API, not those using the low-level API.

- Performance Impact: Expect lower performance inside a Cartesi machine compared to running on the host.

When launching a node with the `--no-backend` you must then start your application on the host and fetch inputs from the endpoint running at http://127.0.0.1:8080/host-runner.


### Verbosity

By default, the Cartesi node works in non-verbose mode, providing logs only from your backend application. For more information, use the `--verbose` command option.

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
