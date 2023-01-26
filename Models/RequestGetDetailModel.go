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
	ResponseCode    string  `json:"response_code"`
	ResponseMessage string  `json:"response_message"`
	ListProduct     Product `json:"produk_detail"`
}
