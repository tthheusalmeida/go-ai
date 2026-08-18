# Go-AI

> Um serviço de ferramentas de IA em Go.

## Configurar a máquina

Dê permissão de execução ao script:

```bash
chmod +x ./scripts/setup.sh
```

Execute o setup:

```bash
sudo ./scripts/setup.sh
```

Em caso de sucesso, aparecerá a mensagem:

```bash
[INFO] Validando instalação...
[INFO] ✓ curl
[INFO] ✓ wget
[INFO] ✓ git
[INFO] ✓ unzip
[INFO] ✓ tar
[INFO] ✓ build-essential
.
.
.
[INFO] Todas as dependências estão funcionando.
[INFO] Configuração concluída com sucesso!
```

## Rodar o projeto

```bash
go run ./cmd/api
```

## Instalar o modelo

Para executar o projeto, é necessário instalar um modelo.
Exemplo de instalação do `Quen3:4b`:

```bash
ollama pull qwen3:4b
```

Listar os modelos:

```bash
ollama list
```
