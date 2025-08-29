# SWE Sniper 🎯

> Precision job tracking and change detection for early-bird SWE internship hunters.

SWE Sniper is a **full-stack, production-ready web application** that continuously monitors curated job boards and GitHub repositories for **Summer 2026 SWE internship postings**.

When new opportunities appear or existing ones change, SWE Sniper instantly detects, logs, and surfaces the updates through a clean, intuitive interface. By combining a **high-performance Go backend** with a **modern React frontend**, it helps aspiring engineers **apply first, stay organized, and gain an edge in competitive recruiting cycles.**

---

## 🚀 Live Demo

🔗 **[Try SWE Sniper here](https://swe-sniper.vercel.app/)**  
_(deployed via Vercel + Render)_

![SWE Sniper Demo](./Demo1.gif) <!-- Replace with your actual GIF or video link -->

---

## 📌 Key Features

- 🔍 **Smart Polling Engine** – Monitors GitHub READMEs and job boards at configurable intervals using efficient diff-checking logic.
- ⚡ **Change Detection & Summarization** – Captures added/removed content with precise diffs and human-readable summaries.
- 🔔 **Real-time Notifications** – Alerts sent via email (SendGrid) or optional webhooks for immediate visibility.
- 🧠 **Auto-Adaptable Parsing** – Handles both full GitHub repo views and raw `.md` file URLs intelligently.
- 🌐 Deployed & Always-On – Backend on Render, frontend on Vercel, monitored via UptimeRobot with health-check endpoints.

---

## 🛠️ Tech Stack

| Layer            | Technology                              |
| ---------------- | --------------------------------------- |
| **Backend**      | Go (Golang), Gin, GoQuery               |
| **Frontend**     | React, CSS-in-JS (Vanilla + shadcn UI)  |
| **Email System** | SendGrid API, Go SMTP abstraction       |
| **Database**     | SQLite (for MVP), pluggable to Postgres |
| **Scraper**      | Custom web scraper made in Go           |
| **Diff Engine**  | Custom text diff algorithm              |
| **Deployment**   | Deployment                              |

Render (backend), Vercel (frontend), UptimeRobot monitoring
|

---

## 🧩 Architecture

```plaintext
[Job Board URL] ─▶ [Poller] ─▶ [Parser] ─▶ [Diff Engine] ─▶ [Notifier]
                                      │
                                      ▼
                                [Supabase / Postgres DB]
                                      │
                                      ▼
                                 [React frontend UI]

•	Poller checks each tracked URL on a schedule.
•	Parser scrapes and extracts structured content.
•	Diff Engine compares current vs. previous snapshots.
•	Notifier alerts users of meaningful changes.
•	Database keeps a log of all URLs and their change history as well
•	Frontend provides a clean UI to visualize changes as user specific data
```

## 💡 Future Enhancements

    •	🔐 OAuth-based login (GitHub, Google)
    •	📦 Webhook plug-ins (Slack, Discord, SMS)
