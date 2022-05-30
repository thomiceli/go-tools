package main

import (
	"log"
	"os"
	"runtime/pprof"

	"golang.org/x/tools/go/packages"
	"honnef.co/go/tools/go/loader"
	"honnef.co/go/tools/lintcmd/cache"
	"honnef.co/go/tools/unused"
)

// OPT(dh): we don't need full graph merging if we're not flagging exported objects. In that case, we can reuse the old
// list-based merging approach.

// OPT(dh): we can either merge graphs as we process packages, or we can merge them all in one go afterwards (then
// reloading them from cache). The first approach will likely lead to higher peak memory usage, but the latter may take
// more wall time to finish if we had spare CPU resources while processing packages.

func main() {
	pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// XXX set cache key for this tool

	c, err := cache.Default()
	if err != nil {
		log.Fatal(err)
	}
	cfg := &packages.Config{
		Tests: true,
	}
	specs, err := loader.Graph(c, cfg, os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// var sg unused.SerializedGraph

	for _, spec := range specs {
		if len(spec.Errors) != 0 {
			continue
		}
		lpkg, _, err := loader.Load(spec)
		if err != nil {
			continue
		}
		if len(lpkg.Errors) != 0 {
			continue
		}
		g := unused.Graph(lpkg.Fset, lpkg.Syntax, lpkg.Types, lpkg.TypesInfo, nil)
		_ = g
		// sg.Merge(g)
	}
}
