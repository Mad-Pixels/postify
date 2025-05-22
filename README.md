<picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
    <source media="(prefers-color-scheme: light)" srcset="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
    <img
        alt="MadPixels"
        src="https://github.com/Mad-Pixels/.github/raw/main/profile/banner.png">
</picture>

# Postify

Help you to streamline content management through a unified structure, enabling one source for multiple publishing endpoints.

## Motivation💡

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
# send / edit telegram post:
# *it use meta.json for find telegram message.
# *if message_id exist send request for change post, else send new post.
$ ./postify tg-send \
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
CONTENT_ALL=$(cd "${ROOT_DIR}" && ls -d */)
CONTENT_NEW=$(git diff --name-only HEAD^ HEAD --diff-filter=A | xargs -n 1 dirname | sort -u)
CONTENT_CHANGED=$(git diff --name-only HEAD^ HEAD --diff-filter=M | xargs -n 1 dirname | sort -u)

for dir in $CONTENT_NEW; do
  ./postify tg-send ...
done
...
for dir in $CONTENT_CHANGED; do
  ./postify tg-send ...
done
...
for dir in $CONTENT_ALL; do
  ./postify html-content ...
done
```

## Deploy
```bash
git tag -a v0.0.1 -m "New Release" 
git push origin v0.0.1
```


# Contributing
We're open to any new ideas and contributions. We also have some rules and taboos here, so please read this page and our [Code of Conduct](/CODE_OF_CONDUCT.md) carefully.

## I want to report an issue
If you've found an issue and want to report it, please check our [Issues](https://github.com/Mad-Pixels/postify/issues) page.