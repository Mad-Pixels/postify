package htmlcontent

const usageTemplate = `Example:
🚀html-content \
  --{{.FlagFromPath}} PATH_TO_CONTENT_SOURCE
  --{{.FlagToPath}} PATH_TO_DIRECTORY_WITH_RESULT

💬Overview:
  This command convert Markdown article to HTML static page.

By default it processed files inside "--{{.FlagFromPath}}" directory
and pick "main.md" file (it can pick multiple files using "--{{.FlagBlocks}}") generate HTML content and
create new file --{{.FlagContentName}} in --{{.FlagToPath}}

Example Structure:
📁my_article/
   ├── 📄main.md
   ├── 📄common.md (another content block)
   └── 📄meta.json

💡For generating "router" use "meta.json" file:
  {
    "static": {
      "title": "{ value }"
      "url":   "{ value }"
    }
  }
  It can be created automatically where:
    title = filepath.Base(--{{.FlagFromPath}})
    url = filepath.Base("content/"--{{.FlagFromPath}})
  Or you can set it manually before running.
🚀html-content ... --{{.FlagRouterPath}} /my_static/router.json
  
💡Template wrapping:
  <html lang="ru">
    <head>
      <meta charset="UTF-8">
    </head>
    <body>
      {{"{{"}} index . "main.md" {{"}}"}}
    </body>
  </html>
🚀html-content ... --{{.FlagTmplPath}} /source/my_html_template`
