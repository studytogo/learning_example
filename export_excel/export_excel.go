package export_excel

import (
	"github.com/tealeg/xlsx"
)

func GeneralExcel() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("容器执行耗时信息")
	if err != nil {
		panic(err.Error())
	}
	style := xlsx.NewStyle()
	style.Font.Color = "FFA500"
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "所属空间"
	cell = row.AddCell()
	cell.Value = "地图学习11"

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "计划类型"
	cell = row.AddCell()
	cell.Value = "差分生产-地图学习11"

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "计划名称"
	cell = row.AddCell()
	cell.Value = "23CW08.6_ML-周版本-两江大道-研发中心-01"

	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "节点"
	cell.SetStyle(style)
	cell = row.AddCell()
	cell.Value = "容器名称"
	cell.SetStyle(style)
	cell = row.AddCell()
	cell.Value = "运行耗时"
	cell.SetStyle(style)
	cell = row.AddCell()
	cell.Value = "启动时间"
	cell.SetStyle(style)
	cell = row.AddCell()
	cell.Value = "完成时间"
	cell.SetStyle(style)
	cell = row.AddCell()
	cell.Value = "容器详情（item）"
	cell.SetStyle(style)

	dataList := make([]map[string]string, 0)
	m1 := map[string]string{
		"name":          "节点名1",
		"containerName": "容器名1",
		"costime":       "hh:mm",
		"startime":      "yyyy/mm/dd hh:mm",
		"finishtime":    "yyyy/mm/dd hh:mm",
		"detail":        "-",
	}

	m2 := map[string]string{
		"name":          "节点名1",
		"containerName": "容器名1",
		"costime":       "hh:mm",
		"startime":      "yyyy/mm/dd hh:mm",
		"finishtime":    "yyyy/mm/dd hh:mm",
		"detail":        "-",
	}
	dataList = append(dataList, m1)
	dataList = append(dataList, m2)

	for _, v := range dataList {
		row = sheet.AddRow()
		for _, v1 := range v {
			cell = row.AddCell()
			cell.Value = v1
		}
	}

	err = file.Save("demo.xlsx")
	if err != nil {
		panic(err.Error())
	}
}
