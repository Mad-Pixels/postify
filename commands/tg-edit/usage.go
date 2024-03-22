package tgedit

const usageTemplate = `Example:
🚀tg-edit \
  --{{.FlagTgToken}} YOUR_BOT_TOKEN \
  --{{.FlagTgChat}} YOUR_CHAT_ID \
  --{{.FlagFromPath}} PATH_TO_CONTENT_SOURCE

💬Overview:
  This command edit content already posted telegram content.
  It used "meta.json" file
  {
    "telegram": {
      "message_id": "{ value }"
      "date":       "{ value }"
    }
  }
  For getting message_id which should be edit.

  By default it processed files inside "--{{.FlagFromPath}}" directory
  and pick "main.md" file (it can pick multiple files using "--{{.FlagBlocks}}").

Example Structure:
📁article1/
    ├── 📄main.md
    ├── 📄common.md (another content block)
    └── 📄meta.json
💡tg-edit ... --{{.FlagBlocks}} main.md,common.md`
