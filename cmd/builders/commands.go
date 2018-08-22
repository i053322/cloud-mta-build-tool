package builders

import (
	"cloud-mta-build-tool/cmd/mta/models"
)

type CommandList struct {
	Info    string
	Command []string
}

// CommandProvider - Get build command's to execute
func CommandProvider(modules models.Modules) []CommandList {
	//Get config from ./commands_cfg.yaml as generated artifacts from source
	commands := Parse(CommandsConfig)
	return mesh(modules, commands)
}

// Match the object according to type and provide the respective command
func mesh(modules models.Modules, commands Builders) []CommandList {
	//The object support deep struct for future use, can be simplified to flat object
	cmds := make([]CommandList, 1)
	for i, b := range commands.Builders {
		//Return only matching types
		if modules.Type == b.Name {
			cmds[i].Info = b.Info
			for _, cmd := range b.Type {
				cmds[i].Command = append(cmds[i].Command, cmd.Command)
			}
		}
	}
	return cmds
}