package Models

type RequestGetDetails struct {
	Id     string `json:"id"`
	Produk string `json:"produk"`
	Stock  string `json:"stok"`
}

type Output struct {
	ResponseCode    string              `json:"response_code"`
	ResponseMessage string              `json:"response_message"`
	ListElektronik  []RequestGetDetails `json:"list_elektronik"`
	ListBuku        []RequestGetDetails `json:"list_buku"`
	ListMinuman     []RequestGetDetails `json:"list_minuman"`
}
