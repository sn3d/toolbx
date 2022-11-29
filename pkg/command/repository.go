package command

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type CommandsRepository struct {
	commandsDir string
}

func CreateCommandsRepository(commandsDir string) *CommandsRepository {
	return &CommandsRepository{
		commandsDir: commandsDir,
	}
}

func (repo *CommandsRepository) GetCommand(args []string) (*CommandInstance, error) {
	rootCmd := &CommandInstance{
		Name: "",
		Dir:  repo.commandsDir,
		Args: args,
	}

	subCmd, err := getSubCommand(rootCmd, args)
	if errors.Is(err, NoChildError) {
		subCmd = rootCmd
	}

	err = loadMetadata(subCmd)
	if err != nil {
		return nil, err
	}

	return subCmd, nil
}

func (repo *CommandsRepository) GetSubcommands(cmd *CommandInstance) ([]*CommandInstance, error) {
	subcommands := make([]*CommandInstance, 0)
	items, _ := ioutil.ReadDir(cmd.Dir)
	for _, item := range items {
		if item.IsDir() && item.Name()[0] != '.' {
			cmd := &CommandInstance{
				Parent: cmd,
				Name:   item.Name(),
				Dir:    filepath.Join(cmd.Dir, item.Name()),
			}

			err := loadMetadata(cmd)
			if err != nil {
				return subcommands, err
			}

			subcommands = append(subcommands, cmd)
		}
	}

	return subcommands, nil
}

func getSubCommand(cmd *CommandInstance, args []string) (*CommandInstance, error) {
	if len(args) == 0 {
		return nil, NoChildError
	}

	subCmd := &CommandInstance{
		Parent: cmd,
		Name:   args[0],
		Dir:    filepath.Join(cmd.Dir, args[0]),
		Args:   args[1:],
	}

	if _, err := os.Stat(subCmd.Dir); os.IsNotExist(err) {
		return nil, NoChildError
	}

	child, err := getSubCommand(subCmd, subCmd.Args)
	if errors.Is(err, NoChildError) {
		return subCmd, nil
	} else {
		return child, nil
	}
}

// load metadata of command from YAML file on given path.
//
// Each command or subcommand is represented by own folder. The folder
// might contain 'command.yaml' file with all command's metadata.
//
// If file is not present, then is returned empty Metadata with no
// error.
func loadMetadata(cmd *CommandInstance) error {

	path := filepath.Join(cmd.Dir, "command.yaml")
	meta := &Metadata{}

	// if 'command.yaml' is missing, we will return empty
	// metadata.
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cmd.Metadata = meta
		return nil
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unkown error in reading metadata from %s", path)
	}

	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return err
	}

	cmd.Metadata = meta
	return nil
}
