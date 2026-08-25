package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const page = `<!doctype html>
<html lang="zh-CN">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="description" content="用于验证一键部署流程的 Go demo 服务" />
    <title>deployment-test</title>
    <style>
      :root {
        color-scheme: light;
        font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
        line-height: 1.5;
        color: #172033;
        background: radial-gradient(circle at top, #f4f7ff 0, #eef3ff 30%, #f8fafc 70%);
      }
      * { box-sizing: border-box; }
      html, body { margin: 0; min-width: 320px; min-height: 100%; }
      body { min-height: 100vh; }
      .page {
        display: grid;
        gap: 24px;
        width: min(960px, calc(100vw - 32px));
        margin: 0 auto;
        padding: 56px 0;
      }
      .hero, .panel {
        border: 1px solid rgba(23, 32, 51, 0.12);
        border-radius: 16px;
        background: rgba(255, 255, 255, 0.82);
        box-shadow: 0 18px 50px rgba(23, 32, 51, 0.08);
        backdrop-filter: blur(18px);
      }
      .hero { padding: 40px; }
      .eyebrow {
        margin: 0 0 12px;
        text-transform: uppercase;
        letter-spacing: 0.12em;
        font-size: 0.78rem;
        color: #5d6a85;
      }
      h1, h2, p, dl { margin: 0; }
      h1 { font-size: clamp(2.5rem, 5vw, 4.2rem); line-height: 1.02; }
      .lede { margin-top: 16px; max-width: 34rem; font-size: 1.08rem; color: #475269; }
      .meta { display: flex; flex-wrap: wrap; gap: 10px; margin-top: 24px; }
      .meta span {
        padding: 8px 12px;
        border-radius: 999px;
        background: #e8eefc;
        color: #27406f;
        font-size: 0.92rem;
      }
      .panel { display: grid; gap: 20px; padding: 28px 32px; }
      .panel h2 { font-size: 1.15rem; }
      .panel p { margin-top: 8px; color: #566179; }
      .specs { display: grid; gap: 14px; }
      .specs div { display: grid; gap: 4px; padding-top: 14px; border-top: 1px solid rgba(23, 32, 51, 0.1); }
      .specs dt { font-size: 0.88rem; color: #6a768d; }
      .specs dd {
        margin: 0;
        font-family: ui-monospace, SFMono-Regular, Consolas, "Liberation Mono", monospace;
        color: #162033;
      }
      @media (max-width: 640px) {
        .page { width: min(calc(100vw - 20px), 960px); padding: 20px 0; }
        .hero, .panel { border-radius: 12px; }
        .hero { padding: 24px; }
        .panel { padding: 20px 24px; }
      }
    </style>
  </head>
  <body>
    <main class="page">
      <section class="hero">
        <p class="eyebrow">Deployment demo</p>
        <h1>deployment-test</h1>
        <p class="lede">这是一个用于验证一键部署服务的最小 Go HTTP 服务。</p>
        <div class="meta">
          <span>Go</span>
          <span>Port 4173</span>
          <span>.qiniu/deploy.yaml</span>
        </div>
      </section>
      <section class="panel">
        <div>
          <h2>部署约定</h2>
          <p>使用 Go 标准库构建并启动 HTTP 服务。</p>
        </div>
        <dl class="specs">
          <div><dt>installCommand</dt><dd>go build -o deployment-test .</dd></div>
          <div><dt>startCommand</dt><dd>./deployment-test</dd></div>
          <div><dt>healthcheck</dt><dd>/</dd></div>
        </dl>
      </section>
    </main>
  </body>
</html>`

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = fmt.Fprint(w, page)
	})
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4173"
	}

	addr := ":" + port
	log.Printf("HTTP server listening on %s", addr)
	if err := http.ListenAndServe(addr, newHandler()); err != nil {
		log.Fatalf("HTTP server stopped: %v", err)
	}
}
