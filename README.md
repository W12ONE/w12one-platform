# W12onePlatform

This project is my hands-on DevOps lab — a place to build, break, and refine ideas while exploring the full lifecycle of delivering, debugging, and scaling modern software systems. This isn’t about building an app — it’s about building, breaking, and debugging the infrastructure and workflows behind apps.

Right now, the setup includes an Astro frontend and a Go API backend. That gives me a great base to experiment with containerization, CI/CD pipelines, observability, and infrastructure-as-code. The current stack is intentionally minimal — it's here to support workflows, not define them.

Later on, once I’ve built out the complete DevOps pipeline, this project will evolve into a personal dashboard or website that showcases both the tooling and infrastructure behind it.

🛠️ Current Stack

- Frontend: Astro
- Backend: Go API (minimal)
- Monorepo: Nx v21
- Containerization: Docker (WIP)
- CI/CD: GitHub Actions
- Infra as Code: Terraform (WIP)


---

## ⚠ Bleeding Edge Disclaimer

This project is built on **Nx v21**, which — at the time of writing — means:
- I’m ahead of many existing community plugins (most are still built for Nx v19 or older),
- I had to patch and work around breaking changes (e.g. for `@nxtension/astro`), though, I'm thinkering with a more permanent plugin for my purposes.... and the rabbit hole runs far deeper than my quick patch ever intended to go,
- I encountered what appears to be a bug in nx/plugin@21, where a plugin was incorrectly scaffolded into the project root instead of libs/ — though I haven’t yet fully confirmed whether it’s reproducible or covered by an existing issue,
- And if it comes down to it, I’ll fork or rewrite missing integrations to make things work.

If something looks broken, it’s probably because I’m holding together bleeding-edge tech with a mix of duct tape, caffein, and sheer willpower.

---

> This project is an evolving proof of concept — optimized not for polish, but for growth.  
> If you're from Shopify (or anyone hiring in DevOps, cloud, or platform roles): this is me learning in public and building the kind of tooling I’d want to deploy in production — just without the corporate firewall — and with more broken stuff to learn from. 🚀

> ### 🧠 What this teaches me so far

- How to work around tooling friction in evolving ecosystems  
- How CI/CD, IaC, and container orchestration fit together in practice  
- Why platform engineers think in systems, not features

🛠️ Still to come

- Structured Terraform config for Azure deployment
- Multi-service Docker setup with build + orchestration
- DevEx enhancements (e.g., Nx plugin or workspace tooling)


