package models

type ParkModel struct {
	Model
	ParkName      string `json:"parkName"`
	Level         string `json:"level"`           // 园区级别
	ParkType      string `json:"parkType"`        // 园区类型
	LandProperty  string `json:"natureOfLand"`    // 土地性质
	MainIndustry  string `json:"leadingIndustry"` // 主导产业
	Introduce     string `json:"introduce"`       // 园区介绍
	Advantage     string `json:"advantage"`       // 园区优势
	EstablishTime string `json:"establishTime"`   // 成立时间
}
