import './style.css';

const app = document.querySelector('#app');

app.innerHTML = `
  <main class="page">
    <section class="hero">
      <p class="eyebrow">Deployment demo</p>
      <h1>deployment-test</h1>
      <p class="lede">
        这是一个用于验证一键部署服务的最小 Vite 站点。
      </p>

      <div class="meta">
        <span>Vite</span>
        <span>Port 4173</span>
        <span>.qiniu/deploy.yaml</span>
      </div>
    </section>

    <section class="panel">
      <div>
        <h2>部署约定</h2>
        <p>安装后执行构建，再用预览服务对外提供静态站点。</p>
      </div>

      <dl class="specs">
        <div>
          <dt>installCommand</dt>
          <dd>npm ci &amp;&amp; npm run build</dd>
        </div>
        <div>
          <dt>startCommand</dt>
          <dd>npm run start</dd>
        </div>
        <div>
          <dt>healthcheck</dt>
          <dd>/</dd>
        </div>
      </dl>
    </section>
  </main>
`;
