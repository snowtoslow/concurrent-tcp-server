package server

type executor struct {
	commands []Command
}

func (executor *executor) executeCommand(input string) {
	for _, executor := range executor.commands {
		executor.execute(input)
	}
}
