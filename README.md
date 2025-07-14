# BranchGen

Ferramenta de linha de comando para gerar nomes de branches Git padronizados, facilitando a adoção de convenções de nomenclatura em projetos.

## Requisitos

| **PACOTE** | **VERSÃO** | **INSTRUÇÕES**               |
|------------|------------|------------------------------|
| Go         | >= 1.21    | <https://go.dev/doc/install> |

## Instalação

### Método 1: Compilação manual

1. Clone este repositório:

```bash
git clone https://github.com/dirsouza/branchgen.git
cd branchgen
```

2. Compile e instale:

```bash
make install
```

Isso irá compilar o programa e instalá-lo em `/usr/local/bin`, permitindo que você execute o comando `branchgen` de qualquer diretório.

### Método 2: Usando Go Install

Se você tem Go instalado, pode instalar diretamente com:

```bash
go install github.com/dirsouza/branchgen@latest
```

## Uso

O BranchGen segue a convenção `tipo/codigo-titulo` para gerar nomes de branches padronizados:

```bash
branchgen -t <tipo> -c <código> -d <título>
```

### Exemplos

```bash
# Gera: feature/XPTO-1234-implementar-nova-funcionalidade
branchgen -t feature -c XPTO-1234 -d "Implementar nova funcionalidade"

# Gera: bugfix/JIRA-5678-corrigir-erro-de-login
branchgen -t bugfix -c JIRA-5678 -d "Corrigir erro de login"
```

### Parâmetros

- `-t, --tipo`: Tipo da branch (feature, bugfix, hotfix, release, support)
- `-c, --codigo`: Código da US (User Story) ou tarefa
- `-d, --titulo`: Título da branch

## Desinstalação

```bash
make uninstall
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

## Licença

Este projeto está licenciado sob a Licença MIT.
