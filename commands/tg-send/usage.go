package tgsend

const usageTemplate = `Example:
🚀tg-send \
  --{{.FlagTgToken}} YOUR_BOT_TOKEN \
  --{{.FlagTgChat}} YOUR_CHAT_ID \
  --{{.FlagFromPath}} PATH_TO_CONTENT_SOURCE

💬Overview:
  This command send request to publish new or change published Telegram content 
  to chat or channel using a specified bot.

By default it processed files inside "--{{.FlagFromPath}}" directory
and pick "main.md" file (it can pick multiple files using "--{{.FlagBlocks}}").
After success uploading a post command update "meta.json" file:
add metadata from Telegram:
  {
    "telegram": {
      "message_id": "{ value }"
      "date":       "{ value }"
    }
  }
If in metadata "message_id" exist send request for modify content.

Example Structure:
📁article1/
    ├── 📄main.md
    ├── 📄common.md (another content block)
    └── 📄meta.json
💡tg-send ... --{{.FlagBlocks}} main.md,common.md`
