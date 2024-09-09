FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# ユーザーのパスワードを設定
ARG USERNAME=vscode
ARG PASSWORD=password

RUN echo "${USERNAME}:${PASSWORD}" | chpasswd