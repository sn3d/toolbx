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
the list of all `storage` subcommands.
```
$ toolbx storage
```