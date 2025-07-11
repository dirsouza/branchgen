package branch

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (gen *Generator) Run(branchType, branchCode, branchTitle string) (string, error) {
	gen.validate(branchType, branchCode, branchTitle)

	slugType := slug.MakeLang(branchType, "pt")
	slugTitle := slug.MakeLang(branchTitle, "pt")

	return fmt.Sprintf("%s/%s-%s", slugType, strings.ToUpper(branchCode), slugTitle), nil
}

func (gen *Generator) validate(branchType, branchCode, branchTitle string) error {
	if strings.TrimSpace(branchType) == "" {
		return errors.New("o tipo da branch não pode ser vazio")
	}

	if strings.TrimSpace(branchCode) == "" {
		return errors.New("o código da US (User Story) ou tarefa não pode ser vazio")
	}

	if strings.TrimSpace(branchTitle) == "" {
		return errors.New("o título da branch não pode ser vazio")
	}

	return nil
}
