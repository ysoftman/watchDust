package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

// versionInfo : 빌드/배포 식별 정보
func versionInfo() string {
	var b strings.Builder
	fmt.Fprintf(&b, "go.version: %s\n", runtime.Version())

	// GAE에서 배포 시 자동 주입되는 환경변수
	for _, key := range []string{"GAE_SERVICE", "GAE_VERSION", "GAE_DEPLOYMENT_ID", "GAE_INSTANCE"} {
		if v := os.Getenv(key); v != "" {
			fmt.Fprintf(&b, "%s: %s\n", strings.ToLower(key), v)
		}
	}

	// VCS(git) 메타데이터 — go build 시 자동 주입 (Go 1.18+)
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Fprintf(&b, "module: %s\n", info.Main.Path)
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision", "vcs.time", "vcs.modified", "GOOS", "GOARCH":
				fmt.Fprintf(&b, "%s: %s\n", s.Key, s.Value)
			}
		}
	}
	return b.String()
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	logInfo(r, "/version 요청 처리")
	SetCommonResponseHeader(w)
	if _, err := fmt.Fprint(w, versionInfo()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
