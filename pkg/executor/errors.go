package executor

import "errors"

// MissingRepoError indicates you have probably not configured toolbx. You should
// create configuration file and define Git repository where are commands
// defined:
//
//     echo "repository: https://github.com/sn3d/toolbx-demo.git" > ~/.toolbx.yaml
//
var MissingRepoError = errors.New("no repository with commands defined")
