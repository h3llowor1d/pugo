package main

import (
	"github.com/codegangsta/cli"
	"github.com/go-xiaohei/pugo/app/builder"
	"github.com/go-xiaohei/pugo/app/command"
	"gopkg.in/inconshreveable/log15.v2"
	"gopkg.in/inconshreveable/log15.v2/ext"
	"path"
)

//go:generate go-bindata -o=app/asset/asset.go -pkg=asset source/... template/... doc/...

const (
	VERSION  = "0.8.6.1229"
	VER_DATE = "2015-12-29"

	SRC_DIR   = "source"   // source contents dir
	TPL_DIR   = "template" // template dir
	MEDIA_DIR = "media"    // upload dir
)

var (
	app = cli.NewApp()
	opt = new(builder.BuildOption)
)

func init() {
	app.Name = "PuGo"
	app.Usage = "a static website generator & deployer in Go"
	app.Author = "fuxiaohei"
	app.Email = "fuxiaohei@vip.qq.com"
	app.Version = VERSION + "(" + VER_DATE + ")"
	opt.SrcDir = SRC_DIR
	opt.TplDir = TPL_DIR
	opt.MediaDir = path.Join(SRC_DIR, MEDIA_DIR)
	opt.Version = VERSION
	opt.VerDate = VER_DATE

	log15.Root().SetHandler(log15.LvlFilterHandler(log15.LvlInfo, ext.FatalHandler(log15.StderrHandler)))
}

func main() {
	// app.Action = action
	app.Commands = []cli.Command{
		command.New(SRC_DIR, TPL_DIR),
		command.Build(opt),
		command.Server(opt),
		command.Doc(opt),
	}
	app.RunAndExitOnError()
}
