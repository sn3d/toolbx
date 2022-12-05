# Toolbx
Put all your organization CLI tools under one roof like `gcloud` or `aws-cli` for
easy distribution and exploration. 

## Motivation

I'm working as a platform engineer in an organization and one of our jobs is 
creating various tools for developer teams. One of our important goals is to 
provide a self-service platform to dev teams. For that reason we're building 
various small or larger CLI tools, the dev teams can use. We observed 2 problems 
with our internal tooling:

- exploration
- distribution

In other words, quite often people weren't aware that some CLI tools exist. 
Also, dev teams must be using the latest version of tools and not some outdated 
versions. We set up an internal Homebrew repository we're using for distribution. 
But it didn't solve the exploration problem very well.

We like tools like `gcloud` or `aws` CLI. It's a massive CLI, where commands are 
organized as groups and subcommands. Each group and subcommand have some description. 
From an exploration perspective, this is working better than a bunch of isolated CLI 
tools. This approach is also creating more consistent tooling and developer have 
everything in one place.

But developing a monolithic massive CLI tool isn't something you want. Platform 
subsystems might be maintained by various teams which means the CLI will turn into 
a bottleneck, to which each team contributes. I wanted to avoid it by splitting 
subcommands into separated binaries with their lifecycle, own release process, and 
clear ownership.

To address these problems, I've created Toolbx.

Toolbx is installed on the developer's machine, it's syncing with Git [command repository](https://github.com/sn3d/toolbx-demo)
where are defined all command groups and subcommands hierarchically with metadata 
like description text etc. Toolbx isn't just a command definition. Toolbx also does 
the installation of subcommands as separated binaries. Thanks to this separation, 
each subcommand might have its repository, versioning, and lifecycle. If a new 
version of the subcommand is released, you have to only change it in the main command 
repository. When the developer executes the subcommand, and its older version, the 
Toolbx will update it.

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

After you installed `toolbx` binary, the second step is to configure it. All you need 
is to execute `.cofigure` dot-command with a repository where are all command groups 
and subcommands defined. You can use [toolbx-demo](https://github.com/sn3d/toolbx-demo) 
repository.

```
$ ./toolbx .configure -repository https://github.com/sn3d/toolbx-demo.git
```

Now you can start using `toolbx`. If you run the subcommand `toolbx storage`, you
should be synchronized with the command repository, and you should get
the list of all `storage` subcommands defined in [demo repository](https://github.com/sn3d/toolbx-demo/tree/main/storage).

```
$ toolbx storage

managing organization storage systems

Available sub-commands for storage

 kafka - manage Kafka
 postgres - manage PostgreSQL in organization
```

If you run a concrete subcommand, the `toolbx` will install binary for a subcommand,
for the current platform, execute the binary and pass the rest of the arguments.

```
$ toolbx storage kafka list                                                                                                                                                          127 ↵
Installing storage-kafka (version:0.0.4 platform:darwin-amd64)...
organization.com Kafka clusters
cluster-1	[server1:9092 server2:9092]
cluster-2	[server3:9092 server4:9092 server5:9092]
```

In the output above you might see the `toolbx` is installing `storage-kafka` binary
for MacOS, it executes the binary, and pass the rest of the arguments. In this case, 
it's `list`. The rest of the output is the subcommand's output.
