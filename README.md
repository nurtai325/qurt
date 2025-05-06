# 🚀 Qurt: An Interactive Programming Language in Kazakh

![Qurt Logo](https://qurt.tech/images/logo.png)

**Qurt** is an innovative, interactive programming language designed to teach the fundamentals of coding in **Kazakh**. Tailored for educators, students, and developers, Qurt bridges the gap between technology and native language, fostering an inclusive learning environment.

> 🏫**Real world usage** — currently used in schools near Almaty to help students learn programming through a culturally relevant and intuitive platform.

> ⚛️ **Built with React** — providing a modern, fast, and dynamic user experience. Editor code in the ./editor directory is written in React.

> ⚡**Blazingly fast** — the project takes advantage of advanced caching and load time optimization techniques. Also, the editor code in React is highly optimized to provide blazingly fast user experience.

---

## 🌐 Explore Qurt

- **Official Website**: [qurt.tech](https://qurt.tech)
- **GitHub Repository**: [github.com/nurtai325/qurt](https://github.com/nurtai325/qurt)

---

## 🛠️ Run Qurt Locally

It is really easy to run **Qurt** locally with docker. You'll need to have the following packages installed on your machine:

### Required Packages:

- Node.js
- Docker

Before running you would have to install **certbot** and generate certificates for qurt.tech for nginx.
Once you have the required packages installed, follow these steps to run the project locally:

```bash
git clone https://github.com/nurtai325/qurt.git
cd qurt/editor
npm i
npm run build
cd ..
docker compose up
```

Open localhost in the browser.

Optionally you can run the custom CI/CD server. 
```bash
GIT_SECRET=webhook_secret go run main.go
```
