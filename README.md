# GO

- https://github.com/uber-go/fx
- https://github.com/uber-go/zap
- https://github.com/knadh/koanf
- https://github.com/spf13/cobra

```go
package main

import (
	"github.com/starudream/go-lib/v2/app"
)

var rootCmd = app.RootCommand(func(c *app.Command) {
	c.Use = "example-app"
})
```

```go
package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"

	"github.com/starudream/go-lib/v2/app"
	"github.com/starudream/go-lib/v2/config"
	"github.com/starudream/go-lib/v2/log"
)

var serveCmd = app.NewCommand(func(c *app.Command) {
	c.Use = "serve"
	c.Short = "serve app"
	c.Run = app.Run(
		app.Provide(serveRun),
		app.Invoke(serveRoute),
	)

	c.PersistentFlags().String("addr", "0.0.0.0:9999", "server listen addr")
})

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveRun(lc app.Lifecycle, k *config.Koanf, logger *log.SugaredLogger) *http.ServeMux {
	mux := http.NewServeMux()
	addr := k.MustString("addr")
	server := &http.Server{Handler: mux}

	start, end :=
		func() error {
			ln, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			logger.Named("serve").Infof("listen on %s", addr)
			go func() { _ = server.Serve(ln) }()
			return nil
		},
		func(ctx context.Context) error {
			logger.Named("serve").Infof("shutdown")
			return server.Shutdown(ctx)
		}

	lc.Append(app.StartStopHook(start, end))

	return mux
}

func serveRoute(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bs, err := httputil.DumpRequest(r, true)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
		} else {
			_, _ = w.Write(bs)
		}
	})
}
```

```go
package main

import (
	"fmt"

	"github.com/starudream/go-lib/v2/app"
)

var actionCmd = app.NewCommand(func(c *app.Command) {
	c.Use = "action"
	c.Short = "action app"
	c.Run = app.Action(actionRun)
})

func init() {
	rootCmd.AddCommand(actionCmd)
}

func actionRun(_ *app.Command, args []string) error {
	fmt.Println("action", args)
	return nil
}
```
