.PHONY: build install uninstall

# Variáveis
BINARY_NAME=branchgen
INSTALL_PATH=/usr/local/bin

build:
    @echo "Compilando $(BINARY_NAME)..."
    @go build -o $(BINARY_NAME) main.go

install: build
    @echo "Instalando $(BINARY_NAME) em $(INSTALL_PATH)..."
    @sudo cp $(BINARY_NAME) $(INSTALL_PATH)/$(BINARY_NAME)
    @echo "✅ $(BINARY_NAME) instalado com sucesso!"
    @echo "Agora você pode executar '$(BINARY_NAME)' de qualquer diretório."

uninstall:
    @echo "Desinstalando $(BINARY_NAME) de $(INSTALL_PATH)..."
    @sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
    @echo "✅ $(BINARY_NAME) desinstalado com sucesso!"