package models

type EncryptData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type EncryptDetails struct {
	InternalKey map[string]string
	OriginalKey map[byte]int
	NumericKey  []int
	TransPosKey []int
	ColumnSize  int
	RowSize     int
}

type Internal struct {
	OriginalKey string
	InternalKey string
}
