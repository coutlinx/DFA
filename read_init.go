package daf

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Word struct {
	Value string
}

// 数据库链接信息
var (
	dbUser = "dbUser"
	dbPass = "123456"
	host   = "127.0.0.1"
	port   = "3306"
	dbBase = "dbBase"
)

func (sw *Word) TableName() string {
	return "sensitive_word"
}

// InitTrie 初始化敏感词树
func InitTrie(excelDir string) error {
	var err error
	if err = readExcel(excelDir); err != nil {
		return err
	}
	if err = readDB(); err != nil {
		return err
	}
	return nil
}

// 读取excel文件
func readExcel(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, excelFileName := range files {
		xlFile, err := xlsx.OpenFile(dir + "\\" + excelFileName.Name())
		if err != nil {
			return err
		}
		columnIndex := 0 // 0 代表第一列，因为列索引是从 0 开始的
		//只读第一个sheet
		sheet := xlFile.Sheets[0]
		// 跳过第一行
		for _, row := range sheet.Rows[1:] {
			if columnIndex < len(row.Cells) {
				cell := row.Cells[0]
				text := cell.String()
				SensitiveTire.insert(text)
			}
		}
	}

	return nil
}

// 读取数据数据库信息
func readDB() error {
	var totalWords int64
	// 替换以下变量为实际的数据库连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, host, port, dbBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db.Debug()
	fmt.Println("Connected to database successfully.")
	db.Model(&Word{}).Count(&totalWords)
	if totalWords > 1000 {
		// 分页读取数据
		var pageSize int64 = 1000
		var totalPages = totalWords / pageSize
		if totalWords%pageSize != 0 {
			totalPages += 1
		}
		for page := int64(0); page < totalPages; page++ {
			var words []Word
			offset := page * pageSize
			if err = db.Limit(int(pageSize)).Offset(int(offset)).Find(&words).Error; err != nil {
				return err
			}
			// 查出来的数据添加进树中
			for _, w := range words {
				SensitiveTire.insert(w.Value)
			}
		}
	} else {
		// 如果不超过1000，直接读取所有数据
		var words []Word
		if err = db.Find(&words).Error; err != nil {
			return err
		}
		// 查出来的数据添加进树中
		for _, w := range words {
			SensitiveTire.insert(w.Value)
		}
	}
	return nil
}
