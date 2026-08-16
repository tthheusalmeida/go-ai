#!/usr/bin/env bash

set -Eeuo pipefail

# ============================================================
# Configuração
# ============================================================

GOLANG_INSTALLER_URL="https://go.dev/dl/go1.26.6.linux-amd64.tar.gz"

GO_INSTALL_DIR="/usr/local/go"
GO_PROFILE_FILE="/etc/profile.d/go.sh"

# ============================================================
# Core
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

    trap 'rm -f "$temp_file"' RETURN

    log "Baixando Go..."

    curl \
        --location \
        --fail \
        --show-error \
        --silent \
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
        -xzf "$temp_file"

    if [[ ! -x "$GO_INSTALL_DIR/bin/go" ]]; then
        error "A instalação do Go falhou."
        exit 1
    fi

    log "Go instalado com sucesso."
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

    export PATH="$PATH:$GO_INSTALL_DIR/bin"

    log "PATH do Go configurado."
}

# ============================================================
# Validação
# ============================================================

validate_installation() {
    log "Validando instalação..."

    if ! command -v go >/dev/null 2>&1; then
        error "Go não foi encontrado no PATH."
        exit 1
    fi

    log "Go encontrado."

    go version

    log "Localização do Go:"

    command -v go
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
    validate_installation

    log "Configuração concluída com sucesso!"
}

main "$@"