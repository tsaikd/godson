package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/flagutil"
	"github.com/tsaikd/KDGoLib/martini/errorJson"
	"github.com/tsaikd/KDGoLib/version"

	"github.com/tsaikd/godson/server/api"
	"github.com/tsaikd/godson/server/config"
)

var (
	// basic server info
	FlagHost = flagutil.AddStringFlag(cli.StringFlag{
		Name:   "host",
		EnvVar: "HOST",
		Value:  "",
		Usage:  "API listen host",
	})
	FlagPort = flagutil.AddIntFlag(cli.IntFlag{
		Name:   "port",
		EnvVar: "PORT",
		Value:  3000,
		Usage:  "API listen port",
	})
	FlagTLS = flagutil.AddBoolFlag(cli.BoolFlag{
		Name:   "tls",
		EnvVar: "TLS",
		Usage:  "API with TLS, 'cert.pem', 'key.pem' should be prepared",
	})
	FlagCORS = flagutil.AddStringFlag(cli.StringFlag{
		Name:   "cors",
		EnvVar: "CORS",
		Value:  "http://localhost:9000",
		Usage:  "Cross-Origin Resource Sharing",
	})

	FlagConfig = flagutil.AddStringFlag(cli.StringFlag{
		Name:   "config",
		EnvVar: "GODSON_CONFIG",
		Value:  "godson.json",
		Usage:  "godson config json file path",
	})
)

func init() {
	version.VERSION = "0.0.1"
}

func Main() {
	app := cli.NewApp()
	app.Name = "godson"
	app.Usage = "continuous integration server written in golang"
	app.Version = version.String()
	app.Action = actionWrapper(MainAction)
	app.Flags = flagutil.AllFlags()

	app.Run(os.Args)
}

func MainAction(c *cli.Context) (err error) {
	m := martini.Classic()
	inj := m.Injector

	inj.Map(inj)
	inj.Map(m)
	inj.Map(c)
	inj.Map(logger)

	config, err := config.NewConfigFromFile(c.GlobalString(FlagConfig.Name))
	if err != nil {
		return
	}
	inj.Map(&config.Bitbuckets)

	// martini
	m.Map(errorJson.ReturnErrorProvider())
	m.Use(render.Renderer())
	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{c.GlobalString(FlagCORS.Name)},
	}))

	// inject api to martini
	if err = api.Inject(inj); err != nil {
		return
	}

	listenAddr := fmt.Sprintf("%s:%d", c.GlobalString(FlagHost.Name), c.GlobalInt(FlagPort.Name))
	if c.GlobalBool(FlagTLS.Name) {
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
