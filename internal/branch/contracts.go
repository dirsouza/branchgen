package branch

type IGenerator interface {
	Run(branchType, branchCode, branchTitle string) (string, error)
}

type ICopyToClipboard interface {
	Run(text string) error
}
