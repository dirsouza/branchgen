# BranchGen

Ferramenta de linha de comando para gerar nomes de branches Git padronizados.

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

```bash
branchgen -t <tipo> -c <código> -d <título>
```

Exemplo:

```bash
branchgen -t feature -c XPTO-1234 -d "Implementar nova funcionalidade"
```

### Parâmetros

- `-t, --tipo`: Tipo da branch (feature, bugfix, hotfix, release, support)
- `-c, --codigo`: Código da US (User Story) ou tarefa
- `-d, --titulo`: Título da branch

## Desinstalação

```bash
make uninstall
```
