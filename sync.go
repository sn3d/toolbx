package toolbx

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"time"
)

func sync(repo, branch, token, syncFile string, commandsDir string) error {
	var err error

	// first we will check if sync file exist, and it's older than current
	// date. This mechanism ensure the sync will run only one per day
	info, err := os.Stat(syncFile)
	if err == nil {
		currentTime := time.Now()
		lastSyncTime := info.ModTime()
		if lastSyncTime.Add(24 * time.Hour).After(currentTime) {
			// skip sync
			return nil
		}
		currentTime.Add(24 * time.Hour)
	}

	// sync the commands
	fmt.Printf("updating commands...\n")
	gitRepo, err := git.PlainOpen(commandsDir)
	if err != nil {
		// let's clone the brand-new repo
		gitOpts := &git.CloneOptions{
			URL:           repo,
			ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		}

		if token != "" {
			gitOpts.Auth = &http.BasicAuth{
				Username: "",
				Password: token,
			}
		}

		_, err = git.PlainClone(commandsDir, false, gitOpts)
		if err != nil {
			return err
		}
	} else {
		// let's update the existing repo
		wtree, err := gitRepo.Worktree()
		if err != nil {
			return err
		}

		err = wtree.Pull(&git.PullOptions{
			RemoteName: "origin",
			Auth: &http.BasicAuth{
				Username: "",
				Password: token,
			},
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return err
		}
	}

	// create sync file, this file is used for ensuring the sync is
	// executed only once per day. The sync is executed only when
	// date of this file is older than current date, or if file is
	// missing
	f, err := os.Create(syncFile)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}
