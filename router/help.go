package router

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"
	"unicode/utf8"
)

var helpCommand = Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "Shows a list of commands or help for one command",
	ArgsUsage: "[command]",
	Action: func(c *Context) error {
		args := c.Args()
		if args.Present() {
			if err := ShowCommandHelp(c, args.First()); err != nil {
				return err
			}
			return nil
		}

		ShowAppHelp(c)
		return nil
	},
}

var helpSubcommand = Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "Shows a list of commands or help for one command",
	ArgsUsage: "[command]",
	Action: func(c *Context) error {
		args := c.Args()
		if args.Present() {
			if err := ShowCommandHelp(c, args.First()); err != nil {
				return err
			}
			return nil
		}

		if err := ShowSubcommandHelp(c); err != nil {
			return err
		}
		return nil
	},
}

// Prints help for the App or Command
type helpPrinter func(w io.Writer, templ string, data interface{})

// Prints help for the App or Command with custom template function.
type helpPrinterCustom func(w io.Writer, templ string, data interface{}, customFunc map[string]interface{})

// HelpPrinter is a function that writes the help output. If not set a default
// is used. The function signature is:
// func(w io.Writer, templ string, data interface{})
var HelpPrinter helpPrinter = printHelp

// HelpPrinterCustom is same as HelpPrinter but
// takes a custom function for template function map.
var HelpPrinterCustom helpPrinterCustom = printHelpCustom

// ShowAppHelpAndExit - Prints the list of subcommands for the app and exits with exit code.
func ShowAppHelpAndExit(c *Context, exitCode int) {
	ShowAppHelp(c)
	os.Exit(exitCode)
}

// ShowAppHelp is an action that displays the help.
func ShowAppHelp(c *Context) (err error) {
	// if c.App.CustomAppHelpTemplate == "" {
	// 	HelpPrinter(c.App.Writer, AppHelpTemplate, c.App)
	// 	return
	// }
	// customAppData := func() map[string]interface{} {
	// 	if c.App.ExtraInfo == nil {
	// 		return nil
	// 	}
	// 	return map[string]interface{}{
	// 		"ExtraInfo": c.App.ExtraInfo,
	// 	}
	// }
	// HelpPrinterCustom(c.App.Writer, c.App.CustomAppHelpTemplate, c.App, customAppData())
	return nil
}

// DefaultAppComplete prints the list of subcommands as the default app completion method
// func DefaultAppComplete(c *Context) {
// 	DefaultCompleteWithFlags(nil)(c)
// }

func printCommandSuggestions(commands []Command, writer io.Writer) {
	for _, command := range commands {
		if command.Hidden {
			continue
		}
		if os.Getenv("_CLI_ZSH_AUTOCOMPLETE_HACK") == "1" {
			for _, name := range command.Names() {
				fmt.Fprintf(writer, "%s:%s\n", name, command.Usage)
			}
		} else {
			for _, name := range command.Names() {
				fmt.Fprintf(writer, "%s\n", name)
			}
		}
	}
}

func cliArgContains(flagName string) bool {
	for _, name := range strings.Split(flagName, ",") {
		name = strings.TrimSpace(name)
		count := utf8.RuneCountInString(name)
		if count > 2 {
			count = 2
		}
		flag := fmt.Sprintf("%s%s", strings.Repeat("-", count), name)
		for _, a := range os.Args {
			if a == flag {
				return true
			}
		}
	}
	return false
}

func printFlagSuggestions(lastArg string, flags []Flag, writer io.Writer) {
	cur := strings.TrimPrefix(lastArg, "-")
	cur = strings.TrimPrefix(cur, "-")
	for _, flag := range flags {
		if bflag, ok := flag.(BoolFlag); ok && bflag.Hidden {
			continue
		}
		for _, name := range strings.Split(flag.GetName(), ",") {
			name = strings.TrimSpace(name)
			// this will get total count utf8 letters in flag name
			count := utf8.RuneCountInString(name)
			if count > 2 {
				count = 2 // resuse this count to generate single - or -- in flag completion
			}
			// if flag name has more than one utf8 letter and last argument in cli has -- prefix then
			// skip flag completion for short flags example -v or -x
			if strings.HasPrefix(lastArg, "--") && count == 1 {
				continue
			}
			// match if last argument matches this flag and it is not repeated
			if strings.HasPrefix(name, cur) && cur != name && !cliArgContains(flag.GetName()) {
				flagCompletion := fmt.Sprintf("%s%s", strings.Repeat("-", count), name)
				fmt.Fprintln(writer, flagCompletion)
			}
		}
	}
}

// func DefaultCompleteWithFlags(cmd *Command) func(c *Context) {
// 	return func(c *Context) {
// 		if len(os.Args) > 2 {
// 			lastArg := os.Args[len(os.Args)-2]
// 			if strings.HasPrefix(lastArg, "-") {
// 				printFlagSuggestions(lastArg, c.App.Flags, c.App.Writer)
// 				if cmd != nil {
// 					printFlagSuggestions(lastArg, cmd.Flags, c.App.Writer)
// 				}
// 				return
// 			}
// 		}
// 		if cmd != nil {
// 			printCommandSuggestions(cmd.Subcommands, c.App.Writer)
// 		} else {
// 			printCommandSuggestions(c.App.Commands, c.App.Writer)
// 		}
// 	}
// }

// ShowCommandHelpAndExit - exits with code after showing help
func ShowCommandHelpAndExit(c *Context, command string, code int) {
	ShowCommandHelp(c, command)
	os.Exit(code)
}

// ShowCommandHelp prints help for the given command
func ShowCommandHelp(ctx *Context, command string) error {
	// show the subcommand help for a command with subcommands
	buf := new(bytes.Buffer)
	if command == "" {
		HelpPrinter(buf, SubcommandHelpTemplate, ctx.App)
		return nil
	}

	for _, c := range ctx.App.Commands {
		if c.HasName(command) {
			if c.CustomHelpTemplate != "" {
				HelpPrinterCustom(buf, c.CustomHelpTemplate, c, nil)
			} else {
				HelpPrinter(buf, CommandHelpTemplate, c)
			}
			return nil
		}
	}
	return nil
}

// ShowSubcommandHelp prints help for the given subcommand
func ShowSubcommandHelp(c *Context) error {
	return ShowCommandHelp(c, c.Command.Name)
}

func printHelpCustom(out io.Writer, templ string, data interface{}, customFunc map[string]interface{}) {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}
	for key, value := range customFunc {
		funcMap[key] = value
	}

	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	err := t.Execute(w, data)
	if err != nil {
		// If the writer is closed, t.Execute will fail, and there's nothing
		// we can do to recover.
		if os.Getenv("CLI_TEMPLATE_ERROR_DEBUG") != "" {
			fmt.Fprintf(ErrWriter, "CLI TEMPLATE ERROR: %#v\n", err)
		}
		return
	}
	w.Flush()
}

func printHelp(out io.Writer, templ string, data interface{}) {
	printHelpCustom(out, templ, data, nil)
}

func checkVersion(c *Context) bool {
	found := false
	if VersionFlag.GetName() != "" {
		eachName(VersionFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkHelp(c *Context) bool {
	found := false
	if HelpFlag.GetName() != "" {
		eachName(HelpFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkCommandHelp(c *Context, name string) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowCommandHelp(c, name)
		return true
	}

	return false
}

func checkSubcommandHelp(c *Context) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowSubcommandHelp(c)
		return true
	}

	return false
}
