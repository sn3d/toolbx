# toolbx
Create your own organization cli tool like gcloud or aws-cli. Toolbx helps you with exploration and distribution of your internal CLI tools

## Install Toolbx

Toolbx currently provides binaries for the following:

- MacOS (64bit only)
- Linux 
- Windows

Download appropriate version for your platform from [Toolbx Releases](https://github.com/sn3d/toolbx/releases). 
Once it's downloaded, you will unarchive the `toolbx` binary. You don't need 
to install the binary to some global location. The binary works well also from 
any place.

But for easy to use you should move the `toolbx` binary somewhere in your `PATH` 
(e.g. to `/usr/local/bin`).

## Quick Start

When you installed `toolbx` binary, as second step is configuration of your 
command repository. It's repository where are all commands and subcommands 
defined. You can use `toolbx-demo` repository:

```
$ echo "repository: https://github.com/sn3d/toolbx-demo.git" > ~/.toolbx.yaml
```

Now you can start using `toolbx`. If you run subcommand `toolbx storage`, you
should be synchronized with command repository, and you should get 
the list of all `storage` subcommands defined in [demo repository](https://github.com/sn3d/toolbx-demo/tree/main/storage).
```
$ toolbx storage

managing organization storage systems

Available sub-commands for storage

 kafka - manage Kafka
 postgres - manage PostgreSQL in organization
```

If you run concrete subcommand, the `toolbx` will install binary for subcommand, 
for current platform, execute the binary and pass the rest of the arguments.

```
$ toolbx storage kafka list                                                                                                                                                          127 ↵
Installing storage-kafka (version:0.0.4 platform:darwin-amd64)...
organization.com Kafka clusters
cluster-1	[server1:9092 server2:9092]
cluster-2	[server3:9092 server4:9092 server5:9092]
```

In output above you might see the `toolbx` is installing `storage-kafka` binary
for MacOS, execute the binary and pass the rest of the arguments, in this case
it's `list`. The rest of the output is subcommand's output.