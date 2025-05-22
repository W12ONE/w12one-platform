# W12onePlatform

This is my hands-on DevOps lab — a place to build, break, and occasionally curse at things while figuring out what makes modern software delivery tick. It's not about building *an* app — it's about breaking *everything around* the app and seeing how gracefully (or not) it fails.

Right now, the setup includes:
- ✅ A working Astro frontend, containerized and nginx’d like a responsible adult
- 🛠️ A Go backend that exists more in spirit than in code — the Dockerfile stub is there, but I decided learning the language might be smart before pretending it's production-ready
- 🧱 A base Docker image scaffolded to reduce duplication I haven't duplicated yet
- ⚙️ Nx v21 running the monorepo — because if I'm going to suffer, I want dependency graphs

This stack is intentionally minimal, mostly because I’ve refused to complicate it before it breaks at least twice. Eventually, this will become a personal dashboard or something equally self-indulgent — built entirely with the DevOps pipeline it's meant to showcase.

---

## 🛠️ Current Stack

| Thing                  | Status                      |
|------------------------|-----------------------------|
| Frontend               | Astro (running, Dockerized) |
| Backend                | Go (barebones, in progress) |
| Monorepo tooling       | Nx v21                      |
| Containerization       | Docker (actively used)      |
| CI/CD                  | GitHub Actions              |
| Infra as Code          | Terraform (scaffolded, staring at me) |

---

## ⚠ Bleeding Edge Disclaimer

This project runs on **Nx v21**, which at the time of writing:
- broke multiple plugins that weren’t ready yet (hi `@nxtension/astro` 👋),
- forced me to patch things manually while questioning my life choices,
- includes a plugin generator that appears to drop files into wrong directories (possibly by design, possibly a feature),
- and may require me to just fork everything and rewrite it because I chose chaos.

If something looks broken, it’s probably because I’m holding together bleeding-edge tech with a mix of duct tape, caffeine, and sheer willpower.

---

> This isn’t optimized for polish. It’s optimized for getting my hands dirty.  
> If you're from Shopify (or anyone hiring in DevOps, cloud, or platform roles):  
> consider this the unfiltered version of what I’d be doing behind a firewall — only with more mess, more fun, and fewer Jira tickets. 🚀

---

## 🧠 What I'm learning (on purpose or by accident)

- How to navigate tools that promise "plugin architecture" and deliver "good luck"
- How CI, Docker, and IaC fight for attention in a monorepo
- Why platform engineers start talking about “systems” and stop talking about “features”

---

## 🧩 Still to come

- Actually finish wiring up the Go API (after I stop googling what a goroutine is)
- Terraform configs that might one day talk to Azure
- Multi-service Docker Compose setup, probably with too many ports
- Developer tooling that makes me feel like less of a sysadmin
