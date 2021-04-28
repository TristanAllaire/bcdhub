package domains

// Repository -
type Repository interface {
	TokenBalances(network, contract, address string, size, offset int64) (TokenBalanceResponse, error)
}