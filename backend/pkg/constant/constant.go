package constant

type URLOperation string

const (
	URLOperationRedirection URLOperation = "redirection"
	URLOperationCanonical   URLOperation = "canonical"
	URLOperationAll         URLOperation = "all"
)
