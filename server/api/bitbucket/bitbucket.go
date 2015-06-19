package bitbucket

import (
	"log"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/godson/server/api"
	"github.com/tsaikd/godson/server/api/bitbucket/payload"
)

var (
	confmap map[string]Configs = map[string]Configs{}
)

func init() {
	api.Regist(func(m *martini.ClassicMartini, configs *Configs) (err error) {
		for _, config := range *configs {
			confs, ok := confmap[config.Repository]
			if !ok {
				confs = Configs{}
			}
			confs = append(confs, config)
			confmap[config.Repository] = confs
		}

		m.Any("/1/bitbucket/:username/:projname", binding.Bind(ApiReq{}), apiHandler)
		return
	})
}

type ApiReq struct {
	payload.Push
}

func apiHandler(logger *log.Logger, r render.Render, apireq ApiReq, params martini.Params) (err error) {
	repo := params["username"] + "/" + params["projname"]
	confs, ok := confmap[repo]
	if !ok {
		err = errutil.New("unknown repository: " + repo)
		return
	}

	for _, config := range confs {
		go func(apireq ApiReq) {
			if err := config.Execute(os.Stdout, apireq.Push); err != nil {
				logger.Println(err.Error())
			}
		}(apireq)
	}

	return
}
