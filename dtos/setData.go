package dtos

type SetDataReq struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Country string `json:"Country"`
	State   string `json:"State"`
}

type SteDataRes struct {
	Success string `json:"Success"`
}
