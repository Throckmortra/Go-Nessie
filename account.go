package gonessie

//Account model from Nessie
type Account struct {
	id 				string `json:"id"`
	accountType 	string `json:"type"`
	nickname 		string `json:"nickname"`
	rewards 		int    `json:"rewards"`
	balance 		int    `json:"balance"`
	accountNumber	string `json:"account_number"`
	customerId 		string `json:"customer_id"`
}

//Accounts array
type Accounts []Account

func (a Account) GetId() string {
	return a.id
}

func (a Account) GetType() string {
	return a.accountType
}

func (a Account) GetNickname() string {
	return a.nickname
}

func (a Account) GetRewards() string {
	return a.rewards
}

func (a Account) GetBalance() int {
	return a.balance
}

func (a Account) GetAccountNumber() string {
	return a.accountNumber
}

func (a Account) GetCustomerId() string {
	return a.customerId
}