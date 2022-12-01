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
	commandsRepo       *command.CommandsRepository
}

func Create(cfg config.Configuration) *ToolbxExecutor {
	executor := &ToolbxExecutor{
		config: cfg,
	}
	return executor
}

// Execute the given command. It might be dot command (e.g. '.configure')
// or user defined command that will invoke tool
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

	// validation & initialization
	if e.config.CommandsRepository == "" {
		return MissingRepoError
	}

	dir.Ensure(e.getCommandsDir())
	dir.Ensure(e.getToolsDir())

	err := sync(e.config.CommandsRepository, e.config.CommandsBranch, e.config.GitlabToken, e.getSyncFile(), e.getCommandsDir())
	if err != nil {
		return err
	}

	e.installedToolsRepo = tool.CreateInstallationsRepository(e.getToolsDir())
	e.commandsRepo = command.CreateCommandsRepository(e.getCommandsDir())

	// execution of command
	cmd, err := e.commandsRepo.GetCommand(args[1:])
	if err != nil {
		return err
	}

	subCmds, err := e.commandsRepo.GetSubcommands(cmd)
	if err != nil {
		return err
	}

	// is it a group or final executable command?
	if len(subCmds) > 0 {
		// it's group because having sub commands,
		// let's print help and list of sub-commands
		if len(cmd.Args) > 0 {
			fmt.Printf("Unknown sub-command '%s'\n", cmd.Args[0])
		} else {
			e.printHelpWithList(cmd, subCmds)
		}
	} else {
		// it's final command because there is no sub commands
		// let's invoke tool for this command
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

	for _, c := range cli.DotCommands {
		if c.Name == dotCommand {
			c.Func(os.Args[2:])
			return nil
		}
	}

	log.Fatalf("Unsupported dot command '%s' \n", dotCommand)
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
	tool := e.installedToolsRepo.GetToolForCommand(cmd)
	if tool == nil {
		return nil, false
	}

	if tool.InstalledVersion != cmd.Metadata.Version {
		return nil, false
	}

	return tool, true
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

// this function is executed when you invoke command with other subcommands
// and print the command's  description or help with list of available
// subcommands
func (e *ToolbxExecutor) printHelpWithList(cmd *command.CommandInstance, subCmds []*command.CommandInstance) {
	name := cmd.Name
	if name == "" {
		name = e.config.BrandLabel
	}

	fmt.Printf("\n%s\n", cmd.Metadata.Description)

	d := color.New(color.FgHiWhite, color.Bold)
	d.Printf("\nAvailable sub-commands for %s\n\n", name)

	for _, subcommand := range subCmds {
		fmt.Printf(" %s - %s\n", subcommand.Name, subcommand.Metadata.Description)
	}

	fmt.Printf("\n")
}
