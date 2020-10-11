package nadesiko4

// Env is enviroments for Nadesiko4
type Env struct {
	IsDebug    bool
	IsEvalMode bool
	MainFile   string
	MainSource string
}

// NewEnv returns Env object
func NewEnv() *Env {
	return &Env{
		IsDebug:    true, // Default
		IsEvalMode: false,
	}
}

// Show Enviroment
func (e *Env) Show() {
	if e.IsDebug {
		println("-d")
	}
	if e.IsEvalMode {
		println("-e")
	}
	println("MainFile: " + e.MainFile)
	println("MainSource: " + e.MainSource)
}
