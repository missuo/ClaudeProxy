[![GitHub Workflow][1]](https://github.com/missuo/ClaudeProxy/actions)
[![Go Version][2]](https://github.com/missuo/ClaudeProxy/blob/main/go.mod)
[![Go Report][3]](https://goreportcard.com/badge/github.com/missuo/ClaudeProxy)
[![Maintainability][4]](https://codeclimate.com/github/missuo/ClaudeProxy/maintainability)
[![GitHub License][5]](https://github.com/missuo/ClaudeProxy/blob/main/LICENSE)
[![Docker Pulls][6]](https://hub.docker.com/r/missuo/claude-proxy)
[![Releases][7]](https://github.com/missuo/ClaudeProxy/releases)

[1]: https://img.shields.io/github/actions/workflow/status/missuo/ClaudeProxy/release.yml?logo=github
[2]: https://img.shields.io/github/go-mod/go-version/missuo/ClaudeProxy?logo=go
[3]: https://goreportcard.com/badge/github.com/missuo/ClaudeProxy
[4]: https://api.codeclimate.com/v1/badges/b5b30239174fc6603aca/maintainability
[5]: https://img.shields.io/github/license/missuo/ClaudeProxy
[6]: https://img.shields.io/docker/pulls/missuo/ClaudeProxy?logo=docker
[7]: https://img.shields.io/github/v/release/missuo/ClaudeProxy?logo=smartthings

# ClaudeProxy
Due to the strict restrictions of the Anthropic Claude API, it can only be used in specific countries, and even when enabling the API, you have to choose whether it will be accessed outside the selected countries. If so, you also need to choose which country it will be accessed in. If this rule is violated, the account may be directly blocked. So, Claude Proxy allows you to fix a single IP for easy and secure access to the Claude API.

## Recommendations
- Use the IP of the country you selected when you activated the API
- Use residential IP, not commercial ones.
- Do not frequently change IP

## Start Claude Proxy
### Docker

```bash
docker run -d --restart always -p 8080:8080 ghcr.io/missuo/claudeproxy:latest
```

```bash
docker run -d --restart always -p 8080:8080 missuo/claude-proxy:latest
```

### Docker Compose

```bash
mkdir claude-proxy && cd claude-proxy
wget -O compose.yaml https://raw.githubusercontent.com/missuo/ClaudeProxy/main/compose.yaml
docker compose up -d
```

### Manual

Download the latest release from the [release page](https://github.com/missuo/ClaudeProxy/releases).

```bash
chmod +x claude-proxy
./claude-proxy -p 8080
```
