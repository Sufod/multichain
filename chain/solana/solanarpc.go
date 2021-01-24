package solana

// AccountContext ...

type AccountContext struct {
	Slot int `json:"slot"`
}

// AccountValue ...
type AccountValue struct {
	Data       string `json:"data"`
	Executable bool   `json:"executable"`
	Lamports   int    `json:"lamports"`
	Owner      string `json:"owner"`
	RentEpoch  int    `json:"rentEpoch"`
}

// ResponseGetAccountInfo ...
type ResponseGetAccountInfo struct {
	Context AccountContext `json:"context"`
	Value   AccountValue   `json:"value"`
}
