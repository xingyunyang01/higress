package tairvector

type KnnSearchResp struct {
	Key        string `json:"key"`
	Similarity string `json:"similarity"`
}

type GetAllResp struct {
	Vector   string `json:"vector"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
