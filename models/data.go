package models

type RootIn struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
}

type RootOut struct {
	A      int `json:"a"`
	B      int `json:"b"`
	C      int `json:"c"`
	Nroots int `json:"n_roots"`
}
