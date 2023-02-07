package models

type Pdst001I struct {
	BarcodeId   string `json:"barcodeId"`
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
}

type Pdst001O []Pdst001IChild

type Pdst001IChild struct {
	ProductId   string `json:"productId" `
	BarcodeId   string `json:"barcodeId"`
	ProductName string `json:"productName" `
	Quantity    int32  `json:"quantity" `
}
