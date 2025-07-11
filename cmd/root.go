package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dirsouza/branchgen/internal/branch"

	"github.com/spf13/cobra"
)

var branchType, branchCode, branchTitle string

var rootCmd = &cobra.Command{
	Use:   "branchgen",
	Short: "Gera nomes de branches Git padronizados",
	Run: func(cmd *cobra.Command, args []string) {
		generator := branch.NewGenerator()
		branchName, err := generator.Run(branchType, branchCode, branchTitle)
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		fmt.Println("\n✅ Branch gerada com sucesso!")
		fmt.Printf("📋 Nome: \033[1;32m%s\033[0m\n", branchName)
		fmt.Printf("🔄 Comando: \033[1;33mgit checkout -b %s\033[0m\n\n", branchName)

		copyToClipboard(branchName)
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&branchType, "tipo", "t", "", "Tipo da branch: feature, bugfix, hotfix, release, support")
	rootCmd.Flags().StringVarP(&branchCode, "codigo", "c", "", "Código da US (User Story) ou tarefa")
	rootCmd.Flags().StringVarP(&branchTitle, "titulo", "d", "", "Título da branch")

	rootCmd.MarkFlagRequired("tipo")
	rootCmd.MarkFlagRequired("codigo")
	rootCmd.MarkFlagRequired("titulo")

	cobra.CheckErr(rootCmd.Execute())
}

func copyToClipboard(text string) {
	fmt.Print("\nDeseja copiar o comando para a área de transferência? [s/N]: ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	var validResponses = map[string]bool{
		"s":   true,
		"sim": true,
	}

	if validResponses[response] {
		gitCommand := fmt.Sprintf("git checkout -b %s", text)
		copy := branch.NewCopyToClipboard()
		if err := copy.Run(gitCommand); err != nil {
			fmt.Printf("\n❌ Erro: %s\n", err)
			fmt.Println("\n💡 Dica: Instale o xclip ou xsel para usar a funcionalidade de copiar para a área de transferência.")
			fmt.Println("   Ubuntu/Debian: sudo apt-get install xclip")
			fmt.Println("   Fedora/CentOS: sudo dnf install xclip")
			fmt.Println("   Arch Linux: sudo pacman -S xclip")
			return
		}

		fmt.Println("\n✅ Comando copiado para a área de transferência!")
	}
}
