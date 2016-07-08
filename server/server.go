package server

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/martini/errorJson"
	"github.com/tsaikd/godson/server/api"
	"github.com/tsaikd/godson/server/config"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("godson").
	SetUsage("continuous integration server written in golang").
	AddFlag(
		&cli.StringFlag{
			Name:        "host",
			EnvVars:     []string{"HOST"},
			Usage:       "API listen host",
			Destination: &flagHost,
		},
		&cli.IntFlag{
			Name:        "port",
			EnvVars:     []string{"PORT"},
			Value:       3000,
			Usage:       "API listen port",
			Destination: &flagPort,
		},
		&cli.BoolFlag{
			Name:        "tls",
			EnvVars:     []string{"TLS"},
			Usage:       "API with TLS, 'cert.pem', 'key.pem' should be prepared",
			Destination: &flagTLS,
		},
		&cli.StringFlag{
			Name:        "cors",
			EnvVars:     []string{"CORS"},
			Value:       "http://localhost:9000",
			Usage:       "Cross-Origin Resource Sharing",
			Destination: &flagCORS,
		},
		&cli.StringFlag{
			Name:        "config",
			EnvVars:     []string{"GODSON_CONFIG"},
			Value:       "godson.json",
			Usage:       "godson config json file path",
			Destination: &flagConfig,
		},
	).
	SetAction(action)

var flagHost string
var flagPort int
var flagTLS bool
var flagCORS string
var flagConfig string

func action(c *cli.Context) (err error) {
	m := martini.Classic()
	inj := m.Injector

	inj.Map(inj)
	inj.Map(m)
	inj.Map(c)
	inj.Map(logger)

	config, err := config.NewConfigFromFile(flagConfig)
	if err != nil {
		return
	}
	inj.Map(&config.Bitbuckets)

	// martini
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{flagCORS},
	}))
	errorJson.BindMartini(m.Martini)

	// inject api to martini
	if err = api.Inject(inj); err != nil {
		return
	}

	listenAddr := fmt.Sprintf("%s:%d", flagHost, flagPort)
	if flagTLS {
		// HTTPS
		// To generate a development cert and key, run the following from your *nix terminal:
		// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
		logger.Println("HTTPS listening on", listenAddr)
		if err = http.ListenAndServeTLS(listenAddr, "cert.pem", "key.pem", m); err != nil {
			return errutil.New("listen failed", err)
		}
	} else {
		m.RunOnAddr(listenAddr)
	}

	return
}
