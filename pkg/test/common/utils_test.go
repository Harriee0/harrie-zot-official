package common_test

import (
	"context"
	"os"
	"path"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"zotregistry.io/zot/pkg/api"
	"zotregistry.io/zot/pkg/api/config"
	tcommon "zotregistry.io/zot/pkg/test/common"
)

func TestWaitTillTrivyDBDownloadStarted(t *testing.T) {
	Convey("finishes successfully", t, func() {
		tempDir := t.TempDir()
		go func() {
			tcommon.WaitTillTrivyDBDownloadStarted(tempDir)
		}()

		time.Sleep(tcommon.SleepTime)

		_, err := os.Create(path.Join(tempDir, "trivy.db"))
		So(err, ShouldBeNil)
	})
}

func TestControllerManager(t *testing.T) {
	Convey("Test StartServer Init() panic", t, func() {
		port := tcommon.GetFreePort()

		conf := config.New()
		conf.HTTP.Port = port

		ctlr := api.NewController(conf)
		ctlrManager := tcommon.NewControllerManager(ctlr)

		// No storage configured
		So(func() { ctlrManager.StartServer() }, ShouldPanic)
	})

	Convey("Test RunServer panic", t, func() {
		tempDir := t.TempDir()

		// Invalid port
		conf := config.New()
		conf.HTTP.Port = "999999"
		conf.Storage.RootDirectory = tempDir

		ctlr := api.NewController(conf)
		ctlrManager := tcommon.NewControllerManager(ctlr)

		ctx := context.Background()

		err := ctlr.Init(ctx)
		So(err, ShouldBeNil)

		So(func() { ctlrManager.RunServer(ctx) }, ShouldPanic)
	})
}
