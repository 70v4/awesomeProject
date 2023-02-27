package request

type BatchMintNFTService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	TokenURI        string `form:"token_uri" json:"token_uri" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
	ContractAddress int    `form:"contract_address" json:"contract_address" binding:"required"`
}

type BatchWithdrawalNFTService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	TokenID         int64  `form:"token_id" json:"token_id" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}

type BatchWithdrawalTokenService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	ToAddress       string `form:"to_address" json:"to_address" binding:"required"`
	Amount          string `form:"amount" json:"amount" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
}

type DiversifyFundsService struct {
	ContractAddress string `json:"contract_address" form:"contract_address"`
	MinBalance      string `json:"min_balance" form:"min_balance"`
}

type RecoveryFundsService struct {
	ContractAddress string `json:"contract_address" form:"contract_address"`
}
