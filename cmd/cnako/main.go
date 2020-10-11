package main

import (
	"github.com/kujirahand/nadesiko4"
	"os"
)

func showUsage() {
	println("# nadesiko ver." + nadesiko4.Version)
	println("[USAGE]")
	println("  cnako -e \"source\"")
	println("  cnako file.nako3")
	println("[Options]")
	println("  -d\tDebug mode")
	println("  -e (src)\tEval mode")
}

func main() {
	if len(os.Args) <= 1 {
		showUsage()
		return
	}
	env := checkOptions()
	env.Show()
}

func checkOptions() *nadesiko4.Env {
	env := nadesiko4.NewEnv()
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		// options
		if v[0] == '-' {
			if v == "-d" {
				env.IsDebug = true
				continue
			}
			if v == "-e" {
				env.IsEvalMode = true
				continue
			}
			println("* unknown options: " + v)
			continue
		}
		// get main args
		if env.IsEvalMode {
			env.MainSource = v
			continue
		}
		if env.MainFile == "" {
			env.MainFile = v
		}
	}
	return env
}
