package Models

type Product struct {
	Id     int    `json:"id"`
	Produk string `json:"produk"`
	Stock  int    `json:"stok"`
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
