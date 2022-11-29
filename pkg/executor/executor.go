package executor

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sn3d/toolbx/pkg/cli"
	"github.com/sn3d/toolbx/pkg/command"
	"github.com/sn3d/toolbx/pkg/config"
	"github.com/sn3d/toolbx/pkg/dir"
	"github.com/sn3d/toolbx/pkg/installer"
	"github.com/sn3d/toolbx/pkg/tool"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

type ToolbxExecutor struct {
	config             config.Configuration
	installedToolsRepo *tool.InstalledToolsRepository
}

// Create and initialize new instance with
// given configuration options
func Initialize(cfg config.Configuration) (*ToolbxExecutor, error) {

	// set default values
	executor := &ToolbxExecutor{
		config: cfg,
	}

	// validation & post-initialization
	if executor.config.CommandsRepository == "" {
		return nil, MissingRepoError
	}

	dir.Ensure(executor.getCommandsDir())
	dir.Ensure(executor.getToolsDir())

	err := sync(cfg.CommandsRepository, cfg.CommandsBranch, cfg.GitlabToken, executor.getSyncFile(), executor.getCommandsDir())
	if err != nil {
		return nil, err
	}

	executor.installedToolsRepo = tool.CreateInstallationsRepository(executor.getToolsDir())

	return executor, nil
}

// Execute the given command
func (e *ToolbxExecutor) Execute(args []string) error {
	if len(os.Args) >= 2 && strings.HasPrefix(os.Args[1], ".") {
		return e.runDotCommand(args)
	} else {
		return e.runCommand(args)
	}
}

func (e *ToolbxExecutor) runCommand(args []string) error {
	if len(args) < 1 {
		return nil
	}

	commands := command.CreateCommandsRepository(e.getCommandsDir())
	cmd, err := commands.GetCommand(args[1:])
	if err != nil {
		return err
	}

	subCmds, err := commands.GetSubcommands(cmd)
	if err != nil {
		return err
	}

	// is it a group or final executable command?
	if len(subCmds) > 0 {
		// it's group because having sub commands,
		// let's print list of sub-commands
		name := cmd.Name
		if name == "" {
			name = e.config.BrandLabel
		}

		if len(cmd.Args) > 0 {
			fmt.Printf("Unknown sub-command '%s'\n", cmd.Args[0])
		}

		fmt.Printf("\n%s\n", cmd.Metadata.Description)

		d := color.New(color.FgHiWhite, color.Bold)
		d.Printf("\nAvailable sub-commands for %s\n\n", name)

		for _, subcommand := range subCmds {
			fmt.Printf(" %s - %s\n", subcommand.Name, subcommand.Metadata.Description)
		}

		fmt.Printf("\n")
	} else {
		// it's command because there is no sub commands
		// let's execute it
		var t *tool.ToolInstance
		t, isInstalled := e.isInstalledAndUpdated(cmd)
		if !isInstalled {
			t, err = e.install(cmd)
			if err != nil {
				return err
			}
		}

		binaryPath := t.Binary(e.getToolsDir())
		cmd := exec.Command(binaryPath, cmd.Args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()

		if err != nil {
			return err
		}
	}

	return nil
}

func (e *ToolbxExecutor) runDotCommand(args []string) error {
	dotCommand := args[1][1:]
	switch dotCommand {
	case "configure":
		cli.ConfigureCmd(os.Args[2:])
	default:
		log.Fatalf("Unsupported dot command '%s' \n", dotCommand)
	}

	return nil
}

func (e *ToolbxExecutor) install(cmd *command.CommandInstance) (*tool.ToolInstance, error) {
	pkg := cmd.GetPackage()
	d := color.New(color.FgHiBlack)
	d.Printf("Installing %s (version:%s platform:%s)...\n", cmd.GetToolID(), cmd.Metadata.Version, pkg.Platform)

	tool := &tool.ToolInstance{
		ID:               cmd.GetToolID(),
		InstalledVersion: cmd.Metadata.Version,
		InstalledCmd:     pkg.Cmd,
	}

	toolDir := dir.Ensure(tool.Dir(e.getToolsDir()))

	uri, err := url.Parse(pkg.Uri)
	if err != nil {
		return nil, err
	}

	opts := installer.InstallationOptions{BearerToken: e.config.GitlabToken}

	err = installer.Install(*uri, toolDir, opts)
	if err != nil {
		return nil, err
	}

	err = e.installedToolsRepo.SaveTool(tool)
	if err != nil {
		return nil, err
	}

	return tool, nil
}

func (e *ToolbxExecutor) isInstalledAndUpdated(cmd *command.CommandInstance) (*tool.ToolInstance, bool) {
	installation := e.installedToolsRepo.GetToolForCommand(cmd)
	if installation == nil {
		return nil, false
	}

	if installation.InstalledVersion != cmd.Metadata.Version {
		return nil, false
	}

	return installation, true
}

// function returns directory where are commands defined. Usually it's
// $TOOLBX_DATA/commands
func (e *ToolbxExecutor) getCommandsDir() string {
	return path.Join(e.config.DataDir, "commands")
}

// function returns directory where are all tools installed. Usually it's
// $TOOLBX_DATA/installed_tools
func (e *ToolbxExecutor) getToolsDir() string {
	return path.Join(e.config.DataDir, "installed_tools")
}

// function returns path to sync fileUsually it's
// $TOOLBX_DATA/sync
func (e *ToolbxExecutor) getSyncFile() string {
	return path.Join(e.config.DataDir, "sync")
}
