# Connecting your AI agent to Taskwondo (MCP)

Taskwondo ships an MCP (Model Context Protocol) server that lets an AI agent — Devin, Claude, Cursor, or any MCP-compatible client — read and manage work items **as you**. The agent only sees what your account can see: your project memberships and role apply.

## 1. Create your API key

In Taskwondo, go to **Preferences → API Keys → Create New Key**:

- **Name**: anything, e.g. `my-agent`
- **Permissions**:
  - `Read only` — agent can only report on work (safest for assistants that just summarize/query)
  - `Read & Write` — agent can create and update work items, comments, time entries
  - `Full access` — same as read & write today; kept for future scopes
- **Expiration**: pick a duration — the key stops working after it

Copy the `twk_…` key — **it is shown only once**. You can revoke it anytime from the same page.

## 2. Get the MCP server binary

The server is a small stdio binary that runs on the machine where your MCP client runs. Download the build for your platform from the releases page:

```
https://github.com/Wameuh/taskwondo/releases
```

(`taskwondo-mcp-linux-amd64`, `taskwondo-mcp-darwin-arm64`, `taskwondo-mcp-windows-amd64.exe`, …) — put it somewhere on your `PATH` or note its full path.

To build it yourself instead (requires Go ≥ 1.25):

```bash
git clone https://github.com/Wameuh/taskwondo
cd taskwondo/mcp && go build -o taskwondo-mcp .
```

## 3. Configure your MCP client

Point your client at the binary with two environment variables:

- `TASKWONDO_URL` → `https://taskwondo.wameuh.cloud`
- `TASKWONDO_API_KEY` → your `twk_…` key

**Devin CLI** (`~/.config/devin/mcp_config.json`):

```json
{
  "mcpServers": {
    "taskwondo": {
      "command": "/home/YOU/.local/bin/taskwondo-mcp",
      "env": {
        "TASKWONDO_URL": "https://taskwondo.wameuh.cloud",
        "TASKWONDO_API_KEY": "twk_YOUR_KEY_HERE"
      }
    }
  }
}
```

**Claude Desktop** (`claude_desktop_config.json`) — same `mcpServers` block.

**Other MCP clients** — configure a stdio server with the same command + env vars.

Restart your client/agent session; the `taskwondo` tools should appear (53 tools: work items, comments, projects, milestones, time tracking, inbox, relations, attachments, members…).

## Notes

- **Your key, your rights**: the agent can only do what your account can — project memberships, roles, and limits all apply. If you need an agent to *not* touch a project, simply don't be a member of it (or use a dedicated account).
- **No admin access needed**: any user can create keys; nothing is installed server-side.
- **Revoking**: delete the key in **Preferences → API Keys** to instantly cut the agent's access.
- Keep the key out of committed files — treat it like a password.
