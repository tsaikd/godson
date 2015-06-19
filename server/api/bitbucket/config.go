package bitbucket

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"os/exec"

	"github.com/tsaikd/KDGoLib/errutil"
)

type Configs []Config

type Config struct {
	// (required) regist repository for trigger, ex: myorganization/myproject
	Repository string `json:"repository,omitempty"`

	// (optional) use provider to eval template result, default: sh, ex: sh
	Provider string `json:"provider,omitempty"`

	// (required) indicate path of template file, to execute, ex: myshell.sh.tmpl
	TemplateFile string `json:"template_file,omitempty"`

	tmpl *template.Template
}

func NewConfigFromFile(filename string) (retconf *Config, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return NewConfigFromData(data)
}

func NewConfigFromData(data []byte) (retconf *Config, err error) {
	config := Config{}
	if err = json.Unmarshal(data, &config); err != nil {
		return
	}

	if err = InitConfig(&config); err != nil {
		return
	}

	retconf = &config
	return
}

func InitConfig(config *Config) (err error) {
	if config.tmpl, err = template.ParseFiles(config.TemplateFile); err != nil {
		return
	}

	if config.Provider == "" {
		config.Provider = "sh"
	}

	return
}

func (t *Config) Execute(w io.Writer, context interface{}) (err error) {
	switch t.Provider {
	case "sh":
		return t.executeShell(w, context)
	default:
		return errutil.New("unknown provider: " + t.Provider)
	}
}

func (t *Config) executeShell(w io.Writer, context interface{}) (err error) {
	buffer := bytes.NewBuffer(nil)
	if err = t.tmpl.Execute(buffer, context); err != nil {
		return
	}

	cmd := exec.Command(t.Provider, "-c", buffer.String())
	cmd.Stdout = w
	cmd.Stderr = w
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
