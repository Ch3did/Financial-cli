#!/bin/bash

set -e


APP_NAME="financial-cli"
INSTALL_DIR="/opt/$APP_NAME"
ALIAS="cx"
BIN_LINK="/usr/local/bin/$ALIAS"
MAIN_FILE="cmd/main.go"
GO_VERSION="1.22.3"


# 2. Copia o projeto para /opt/financial-cli
echo ">> Instalando aplicação em $INSTALL_DIR"
sudo rm -rf "$INSTALL_DIR"
sudo mkdir -p "$INSTALL_DIR"
sudo cp -r . "$INSTALL_DIR"
echo "Diretório atual: $(pwd)"

# 3. Cria o script de execução
echo ">> Criando script executável em $BIN_LINK"

sudo tee "$BIN_LINK" > /dev/null <<EOF
#!/bin/bash
cd "$INSTALL_DIR"
exec go run $MAIN_FILE "\$@"
EOF

sudo chmod +x "$BIN_LINK"

# 4. Executa init (inicialização do banco)
echo ">> Executando '$ALIAS init' para configurar banco de dados..."
$BIN_LINK init

echo ">> Instalação concluída!"
echo "Você pode rodar o CLI com o comando: $ALIAS"
