FROM mcr.microsoft.com/vscode/devcontainers/go:1.15 as dev

ENTRYPOINT ["tail", "-f", "/dev/null"]