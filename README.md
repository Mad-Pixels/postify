![CI/CD Status](https://github.com//Mad-Pixels/go-postify/actions/workflows/publish.yml/badge.svg?branch=main)

# Postify

Help you to streamline content management through a unified structure, enabling one source for multiple publishing endpoints.

## Motivation 💡

The inception of Postify was driven by a need for an efficient tool to manage content publication across different 
platforms seamlessly. As a developer and content creator, I sought a solution that could:

- Facilitate the transition of content from Markdown to various endpoints, such as a Telegram channel and a website.
- Ensure that content posted to one platform could be easily accessed and linked from another, enhancing the reader's experience by providing expanded content views and direct links to the original posts.
- Streamline the content update process, allowing for simultaneous updates across all platforms with minimal effort.

This vision was particularly geared towards managing documentation sites or running thematic Telegram channels, like my own `golang-etc`, where I explore different features of the Go programming language. The ability to link posts across platforms means that content is not only easily accessible but also interactive and engaging.

Leveraging Git as the content repository, Postify integrates smoothly into development workflows. Adding or updating articles triggers automatic publishing to Telegram and the static site's rebuild and deployment processes. This flexible system simplifies content management, making it an ideal tool for developers and content creators who manage dynamic content across multiple platforms.

## Usage 🛠

Postify is designed to work with a specific structure of content directories and metadata files to manage content efficiently across different platforms. Here's how to organize your content and metadata for Postify:
### Content Directory Structure 📁
```sql
./my_cool_article
    ├── main.md # main content file, can be multiple, specified by the `--with-blocks` flag.
    └── meta.json # system file, automatically filled with metadata but it can be changed manually.
```

### Metadata File (meta.json) 📄
The `meta.json` file is crucial for managing content across platforms. It stores metadata for each piece of content, including information for Telegram posts and static website generation.
```json
{
  "telegram": {
    "message_id": "{{ auto-filled when content will be sent to Telegram }}",
    "date": "{{ auto-filled when content will be sent to Telegram }}"
  },
  "static": {
    "title": "{{ can be filled manually, if not get `{{ article_dir_name }}` }}",
    "url": "can be filled manually, if not get `content/{{ article_dir_name }}` "
  }
}
```
`static` block used for generating a correct `router.json` file to ensure that links correctly lead to the content on website.  
For work with telegram you should create your telegram bot and add it to your channel or group.
### Cli-app  💻
```bash
# send post to telegram:
$ ./postify tg-send \
  --from {{ ./article }} \
  --chat-id {{ channel or chat id }} \
  --bot-token {{ telegram bot token }}

# edit telegram post:
# *it use meta.json for find telegram message.
$ ./postify tg-edit \
  --from {{ ./article }} \
  --chat-id {{ channel or chat id }} \
  --bot-token {{ telegram bot token }}

# generate static content:
$ ./postify html-content \
  --from {{ ./article }} \
  --to {{ directory with generated content }} \
  --with-assets {{ optional: path to additional files witch will be copy to content directory }} \
  --with-tmpl {{ optional: path to template file for adding generating content to template }} \
  --with-router {{ optional: path to router file }}
```
  
**By the way it can be used as a part of CI/CD process:**
```bash
ROOT_DIR=$(git rev-parse --show-toplevel)
ALL_CONTENT=$(cd "${ROOT_DIR}" && ls -d */)
NEW_CONTENT=$(git status --porcelain | awk '/^A|^AM/ {if (!($2 ~ /^\.|\/\./)) print $2}' | xargs -n1 dirname | sort -u | grep -vE '^\.?$')
CHANGED_CONTENT=$(git status --porcelain | awk '/^ M/ {if (!($2 ~ /^\.|\/\./)) print $2}' | xargs -n1 dirname | sort -u | grep -vE '^\.?$')

for dir in $NEW_CONTENT; do
  "postify tg-send ...
done
...
for dir in $CHANGED_CONTENT; do
  "postify tg-edit ...
done
...
for dir in $ALL_CONTENT; do
  "postify html-content ...
done
```

## Contributing
We welcome contributions!  
Please feel free to submit pull requests or open issues to discuss new features or improvements.

## License
This project is licensed under the GPL-3.0 License - see the 
[LICENSE](https://github.com/Mad-Pixels/go-postify/blob/main/LICENSE) file for details.

## Deploy
```bash
git tag -a v0.0.1 -m "New Release" 
git push origin v0.0.1
```
