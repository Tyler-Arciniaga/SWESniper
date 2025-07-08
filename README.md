# SWE Sniper 🎯

> Precision job tracking and change detection for early-bird SWE internship hunters.

SWE Sniper is a fully functional, production-grade backend monitoring tool paired with a lightweight frontend interface. It automatically tracks curated job boards and GitHub repositories for **Summer 2026 SWE internship opportunities**, notifies users about updates, and logs historical changes with visual clarity — empowering aspiring engineers to **apply first and stay ahead**.

---

## 🚀 Demo

![SWE Sniper Demo](./Demo1.gif) <!-- Replace with your actual GIF or video link -->

---

## 📌 Key Features

- 🔍 **Smart Polling Engine** – Monitors GitHub READMEs and job boards at configurable intervals using efficient diff-checking logic.
- ⚡ **Change Detection & Summarization** – Captures added/removed content with precise diffs and human-readable summaries.
- 🔔 **Real-time Notifications** – Alerts sent via email (SendGrid) or optional webhooks for immediate visibility.
- 🧠 **Auto-Adaptable Parsing** – Handles both full GitHub repo views and raw `.md` file URLs intelligently.
- 🧱 **Modular, Scalable Architecture** – Designed for future multi-user support, user dashboards, and cloud deployment.

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
| **Deployment**   | Docker (future), Localhost for testing  |

---

## 🧩 Architecture

```plaintext
[Job Board URL] ─▶ [Poller] ─▶ [Parser] ─▶ [Diff Engine] ─▶ [Notifier]
                                      │
                                      ▼
                                [SQLite DB]

•	Poller checks each tracked URL on a schedule.
•	Parser scrapes and extracts structured content.
•	Diff Engine compares current vs. previous snapshots.
•	Notifier alerts users of meaningful changes.
•	Database keeps a log of all URLs and their change history.
```

## 💡 Future Enhancements

    •	👥 Multi-user authentication & dashboards
    •	☁️ Cloud deployment via Fly.io or Render
    •	📈 Analytics dashboard for tracked URLs
    •	🔐 OAuth-based login (GitHub, Google)
    •	📦 Webhook plug-ins (Slack, Discord, SMS)
