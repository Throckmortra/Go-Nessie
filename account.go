package gomodels

//Account model from Nessie
type Account struct {
	ID         string `json:"id"`
	Nickname   string `json:"nickname"`
	Rewards    int    `json:"rewards"`
	Balance    int    `json:"balance"`
	CustomerID string `json:"customer_id"`
}

//Accounts array
type Accounts []Account

func (a Account) GetID() string {
	return a.ID
}

func (a Account) GetNickname() string {
	return a.Nickname
}
