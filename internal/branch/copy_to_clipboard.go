package branch

import (
	"errors"
	"os/exec"
	"strings"
)

type CopyToClipboard struct{}

func NewCopyToClipboard() *CopyToClipboard {
	return &CopyToClipboard{}
}

func (copy *CopyToClipboard) Run(text string) error {
	if err := copy.xclip(text); err != nil {
		if err := copy.xsel(text); err != nil {
			return errors.New("nenhum dos comandos xclip ou xsel de cópia para a área de transferência disponível")
		}
	}

	return nil
}

func (copy *CopyToClipboard) xclip(text string) error {
	if err := copy.hasPackage("xclip"); err != nil {
		return err
	}

	cmd := exec.Command("xclip", "-selection", "clipboard")
	return copy.execute(text, cmd)
}

func (copy *CopyToClipboard) xsel(text string) error {
	if err := copy.hasPackage("xsel"); err != nil {
		return err
	}

	cmd := exec.Command("xsel", "--clipboard", "--input")
	return copy.execute(text, cmd)
}

func (copy *CopyToClipboard) hasPackage(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return err
	}

	return nil
}

func (copy *CopyToClipboard) execute(text string, cmd *exec.Cmd) error {
	cmd.Stdin = strings.NewReader(text)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
