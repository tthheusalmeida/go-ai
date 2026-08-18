#!/usr/bin/env bash

set -Eeuo pipefail

# ============================================================
# Configuração
# ============================================================

GOLANG_INSTALLER_URL="https://go.dev/dl/go1.26.6.linux-amd64.tar.gz"

GO_INSTALL_DIR="/usr/local/go"
GO_PROFILE_FILE="/etc/profile.d/go.sh"

OLLAMA_INSTALLER_URL="https://ollama.com/install.sh"

# ============================================================
# Cores
# ============================================================

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# ============================================================
# Funções auxiliares
# ============================================================

log() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

require_root() {
    if [[ "$EUID" -ne 0 ]]; then
        error "Este script precisa ser executado como root."
        error "Execute: sudo $0"
        exit 1
    fi
}

validate_installed_command() {
    local command="$1"
    local name="$2"

    if ! command -v "$command" >/dev/null 2>&1; then
        error "A instalação do $name falhou."
        exit 1
    fi

    log "$name instalado com sucesso."
}

validate_executable() {
    local path="$1"
    local name="$2"

    if [[ ! -x "$path" ]]; then
        error "A instalação do $name falhou."
        exit 1
    fi

    log "$name instalado com sucesso."
}

spinner() {
    local pid="$1"
    local message="$2"
    local chars='|/-\'
    local i=0

    while kill -0 "$pid" 2>/dev/null; do
        printf "\r${GREEN}[INFO]${NC} %s %s" \
            "$message" \
            "${chars:i++%${#chars}:1}"

        sleep 0.1
    done

    printf "\r${GREEN}[INFO]${NC} %s ✓\n" "$message"
}

# ============================================================
# Sistema
# ============================================================

update_system() {
    log "Atualizando índices do APT..."

    apt update

    log "Atualizando pacotes instalados..."

    apt upgrade -y
}

install_base_packages() {
    log "Instalando pacotes básicos..."

    apt install -y \
        curl \
        wget \
        git \
        ca-certificates \
        unzip \
        tar \
        build-essential
}

# ============================================================
# Go
# ============================================================

install_golang() {
    log "Instalando Go..."

    if [[ -z "$GOLANG_INSTALLER_URL" ]]; then
        error "GOLANG_INSTALLER_URL não foi configurada."
        exit 1
    fi

    local temp_file

    temp_file="$(mktemp --suffix=.tar.gz)"

    log "Baixando Go..."

    curl \
        --location \
        --fail \
        --show-error \
        --progress-bar \
        "$GOLANG_INSTALLER_URL" \
        --output "$temp_file"

    log "Download concluído."

    if [[ -d "$GO_INSTALL_DIR" ]]; then
        log "Removendo instalação anterior do Go..."

        rm -rf "$GO_INSTALL_DIR"
    fi

    log "Extraindo Go em /usr/local..."

    tar \
        -C /usr/local \
        -xzf "$temp_file" &

    local tar_pid=$!

    spinner "$tar_pid" "Extraindo Go"

    wait "$tar_pid"

    rm -f "$temp_file"

    validate_executable "$GO_INSTALL_DIR/bin/go" "Go"
}

# ============================================================
# PATH
# ============================================================

configure_go_path() {
    log "Configurando PATH do Go..."

    cat > "$GO_PROFILE_FILE" <<EOF
export PATH="\$PATH:$GO_INSTALL_DIR/bin"
EOF

    chmod 644 "$GO_PROFILE_FILE"

    source "$GO_PROFILE_FILE"

    log "PATH do Go configurado."
}

# ============================================================
# Ollama
# ============================================================

install_ollama() {
    log "Instalando Ollama..."

    if [[ -z "$OLLAMA_INSTALLER_URL" ]]; then
        error "OLLAMA_INSTALLER_URL não foi configurada."
        exit 1
    fi

    curl \
        --fail \
        --show-error \
        --location \
        "$OLLAMA_INSTALLER_URL" | sh

    validate_installed_command "ollama" "Ollama"
}

# ============================================================
# Validação
# ============================================================

validate_command() {
    local command="$1"
    local name="$2"

    if command -v "$command" >/dev/null 2>&1; then
        log "✓ $name"
        return 0
    fi

    error "✗ $name"
    return 1
}

validate_installation() {
    log "Validando instalação..."

    local failed=0

    validate_command "curl" "curl" || failed=1
    validate_command "wget" "wget" || failed=1
    validate_command "git" "git" || failed=1
    validate_command "unzip" "unzip" || failed=1
    validate_command "tar" "tar" || failed=1
    validate_command "gcc" "build-essential" || failed=1
    validate_command "ollama" "Ollama" || failed=1

    if command -v go >/dev/null 2>&1; then
        local go_version

        go_version="$(go version | awk '{print $3}')"

        log "✓ Go ${go_version#go}"
    else
        error "✗ Go"
        failed=1
    fi

    if command -v ollama >/dev/null 2>&1; then
        local ollama_version

        ollama_version="$(ollama --version)"

        log "✓ ${ollama_version}"
    else
        error "✗ Ollama"
        failed=1
    fi

    if [[ ":$PATH:" == *":$GO_INSTALL_DIR/bin:"* ]]; then
        log "✓ Go PATH"
    else
        error "✗ Go PATH"
        failed=1
    fi

    if [[ "$failed" -ne 0 ]]; then
        error "A validação encontrou problemas."
        exit 1
    fi

    log "Todas as dependências estão funcionando."
}

# ============================================================
# Main
# ============================================================

main() {
    log "Iniciando configuração da máquina..."

    require_root

    update_system
    install_base_packages
    install_golang
    configure_go_path
    install_ollama
    validate_installation

    log "Configuração concluída com sucesso!"
    log "Para atualizar o terminal atual:"
    log "source /etc/profile.d/go.sh"
}

main "$@"
