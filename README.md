# BackupForge

<div align="center">
  <img src="frontend/public/logo.svg" alt="BackupForge Logo" width="150"/>
  <h1>BackupForge</h1>
  <p>A premium native database backup & restore application for Windows (PostgreSQL, MySQL, MariaDB, MongoDB).</p>
</div>

---

## ✨ Features

* **Zero-Configuration Portable Environment**: Comes bundled with local portable Go and Valkey/Redis cache servers.
* **Pure Desktop App Experience**: Launches directly as a borderless native application window (`msedge --app` mode) on execution.
* **Offline-Focused & Private**: No external web links, star-buttons, or documentation trackers. All data stays local.
* **Disabled Browser-Like Artifacts**: Intercepts and disables browser right-click context menus and layout text-selections to feel like a high-performance system utility.
* **Seamless Reloading**: Supports `Ctrl + R` and `F5` keyboard hotkeys to reload the workspace environment smoothly.
* **Vibrant Forge Theme**: Styled with a cohesive warm forge-orange theme preset across all UI components and dialogs.
* **Smart Security**: Password input masking using native styling to prevent modern browsers from showing intrusive "Save Password?" autofill prompts.

---

## 🚀 Getting Started on Windows

Running BackupForge locally on your machine is incredibly simple. 

1. **Prerequisites**: Ensure you have `pnpm` installed on your machine.
2. **Launch**: Double-click or run the local launch script:
   ```powershell
   .\run-local.bat
   ```
3. **What happens behind the scenes**:
   * Detects and links the local Go compiler.
   * Boots the local Redis cache server automatically.
   * Compiles the high-fidelity frontend assets directly into the Go backend.
   * Starts the Go web server and auto-launches the application window.

---

## 🛠️ Configuration & Architecture

* **Embedded Frontend**: All React assets are built and compiled into the Go binary at `backend/ui/build`, meaning the entire server is served from a single port (`4005`).
* **Valkey/Redis Cache Client**: Configured to disable client-side tracking, ensuring absolute stability on local Windows developer machines without version or RESP compatibility conflicts.
* **Database Drivers**: Windows database client bypasses allow logical and physical backups to run without requiring pre-compiled Linux database executables.

---

## 📝 License

This project is licensed under the Apache 2.0 License.
