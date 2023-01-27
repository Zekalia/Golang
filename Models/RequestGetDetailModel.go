package Models

type Product struct {
	Id     string `json:"id"`
	Produk string `json:"produk"`
	Stock  string `json:"stok"`
}

type ReqGetBuku struct {
	Buku string `json:"buku"`
}

type Output struct {
	ResponseCode    string    `json:"response_code"`
	ResponseMessage string    `json:"response_message"`
	ListElektronik  []Product `json:"list_elektronik"`
	ListBuku        []Product `json:"list_buku"`
	ListMainan      []Product `json:"list_mainan"`
}

type ReqGetElectronicById struct {
	Id string `json:"id"`
}

type ReqInsertDataElectronic struct {
	Id         string `json:"ID"`
	Elektronik string `json:"ELEKTRONIK"`
	Mainan     string `json:"MAINAN"`
	Buku       string `json:"BUKU"`
	Stok       string `json:"STOK"`
}

type ReqUpdateDataElectronic struct {
	Id         string `json:"ID"`
	Elektronik string `json:"ELEKTRONIK"`
	Mainan     string `json:"MAINAN"`
	Buku       string `json:"BUKU"`
	Stok       string `json:"STOK"`
}

type OutputElectronic struct {
	ResponseCode    string  `json:"response_code"`
	ResponseMessage string  `json:"response_message"`
	Elektronik      Product `json:"elektronik"`
}
