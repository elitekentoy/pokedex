package repl

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}
