package toolbx

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"net/url"
	"os"
	"os/exec"
	"toolbx/api"
	"toolbx/install"
	"toolbx/install/archive"
)

type Toolbx struct {
	name             string
	commandsDir      string
	gitlabToken      string
	syncRepo         string
	syncRepoBranch   string
	syncFile         string
	installationsDir string
	installers       map[string]install.Installer
}

func Main(options ...ToolbxOption) error {
	tlbx, err := Create(options...)
	if err == MissingRepoError {
		fmt.Println("Hello, and welcome!")
		fmt.Println("")
		fmt.Println("Do one simple configuration step before you start using Toolbx:")
		fmt.Println("")
		fmt.Println("    echo \"repository: https://github.com/sn3d/toolbx-demo.git\" > ~/.toolbx.yaml")
		fmt.Println("")
		return nil
	} else if err != nil {
		log.Fatalln("error starting toolbx:", err)
	}

	err = tlbx.Execute(os.Args)
	if err != nil {
		log.Fatalln("error executing command:", err)
	}

	return nil
}

// Create and initialize new instance with
// given configuration options
func Create(options ...ToolbxOption) (*Toolbx, error) {

	// set default values
	toolbx := &Toolbx{}
	defaultValues(toolbx)

	// apply options
	for _, o := range options {
		o(toolbx)
	}

	// validation & post-initialization
	if toolbx.syncRepo == "" {
		return nil, MissingRepoError
	}

	toolbx.installers = map[string]install.Installer{
		"https+zip": archive.Installer(toolbx.gitlabToken),
	}
	ensureDir(toolbx.commandsDir)
	ensureDir(toolbx.installationsDir)

	err := sync(toolbx.syncRepo, toolbx.syncRepoBranch, toolbx.gitlabToken, toolbx.syncFile, toolbx.commandsDir)
	if err != nil {
		return nil, err
	}

	return toolbx, nil
}

// Execute the given command
func (t *Toolbx) Execute(args []string) error {
	if len(args) < 1 {
		return nil
	}

	commands := CreateCommandsRepository(t.commandsDir)
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
			name = t.name
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
		var installation *api.Installation
		installation, installed := t.isInstalledAndUpdated(cmd)
		if !installed {
			installation, err = t.install(cmd)
			if err != nil {
				return err
			}
		}

		binaryPath := installation.Binary(t.installationsDir)
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

func (t *Toolbx) install(cmd *api.Command) (*api.Installation, error) {
	pkg := cmd.GetPackage()
	fmt.Printf("Installing %s (version:%s platform:%s)...\n", cmd.GetInstallationID(), cmd.Metadata.Version, pkg.Platform)

	installation := &api.Installation{
		ID:               cmd.GetInstallationID(),
		InstalledVersion: cmd.Metadata.Version,
		InstalledCmd:     pkg.Cmd,
	}

	dir := ensureDir(installation.Dir(t.installationsDir))

	uri, err := url.Parse(pkg.Uri)
	if err != nil {
		return nil, err
	}

	installer := t.installers[uri.Scheme]
	if installer == nil {
		return nil, fmt.Errorf("unsupported package scheme %s", uri.Scheme)
	}

	err = installer.Install(*uri, dir)
	if err != nil {
		return nil, err
	}

	err = CreateInstallationsRepository(t.installationsDir).SaveInstallation(installation)
	if err != nil {
		return nil, err
	}

	return installation, nil
}

func (t *Toolbx) isInstalledAndUpdated(cmd *api.Command) (*api.Installation, bool) {
	installation := CreateInstallationsRepository(t.installationsDir).GetInstallationForCommand(cmd)
	if installation == nil {
		return nil, false
	}

	if installation.InstalledVersion != cmd.Metadata.Version {
		return nil, false
	}

	return installation, true
}

func ensureDir(dir string) string {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	return dir
}
