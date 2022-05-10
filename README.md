# Toolbx
Put all your organization CLI tools under one roof like `gcloud` or `aws-cli` for
easy distribution and exploration. 

## Motivation

I'm working as platform engineer in organization and one of our jobs is creating
various tools for developer teams. One of our important goals is to provide 
self-service platform to dev teams. For that reason we're building various small 
or larger CLI tools, the dev teams can use. We observed 2 problems with our 
internal tooling:

- exploration
- distribution

In other words, quite often people wasn't aware that some CLI tools exist. Also, 
it's important that dev teams are using the latest version of tools and not some 
outdated versions. We setup internal Homebrew repository we're using for 
distribution. But it didn't solve exploration problem very well.

We like tools like `gcloud` or `aws` CLI. It's massive CLI, where commands are
organized as groups and subcommands. Each group and subcommand have some 
description. From exploration perspective this is working better than a bunch
of isolated CLI tools. This approach is also creating more consistent tooling 
and developer have everything on one place.  

But developing monolithic massive CLI tool isn't something you want. Platform 
subsystems might be maintained by various teams which means the CLI will turn 
into bottleneck where each team contribute. I wanted to avoid it by splitting 
subcommands into separated binaries with own lifecycle, own release process 
and clear ownership. 

To address these problems, I've created Toolbx. 

Toolbx is installed on developer's machine, it's syncing with Git [command repository](https://github.com/sn3d/toolbx-demo)
where are defined all command groups and subcommands hierarchically with metadata like 
description text etc. Toolbx isn't just command definition. Toolbx also do installation 
of subcommands as separated binaries. Thanks to this separation, each subcommand might 
have own repository, versioning, lifecycle. If new version of subcommand is released,
you have to only change it in main command repository. When developer execute
the subcommand, and it's older version, the toolbx will update it.

## Installation

Toolbx is simple binary and installation is easy. Toolbx currently provides 
binaries for the following:

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
[command repository](https://github.com/sn3d/toolbx-demo). It's repository 
where are all command groups and subcommands defined. You can use `toolbx-demo` repository:

```
$ echo "repository: https://github.com/sn3d/toolbx-demo.git" > ~/.toolbx.yaml
```

Now you can start using `toolbx`. If you run subcommand `toolbx storage`, you
should be synchronized with command repository, and you should get 
the list of all `storage` subcommands defined in [demo repository](https://github.com/sn3d/toolbx-demo/tree/main/storage).
```
$ toolbx storage

managing organization storage systems

Available subcommands for storage

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
