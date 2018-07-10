package store

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type GetOptions struct {
	Sort      [][]string               `json:"sort"`
	Filter    [][3]interface{}         `json:"filter"`
	Page      int                      `json:"page"`
	PageCount int                      `json:"page_count"`
	IN        map[string][]interface{} `json:"in"`
}

// Get 应付单表的一切查询
func Get(
	feilds []string,
	db *gorm.DB, // 数据库连接
	query interface{}, // 查询条件
	getOptions *GetOptions, // 过滤条件二进制
	list interface{}, // 列表
) error {

	chain := commonQuery(feilds, db, query, getOptions, list)
	check := chain.Find(list, query)
	if check == nil {
		return errors.New("db error")
	}
	return nil
}

// UpdateIfExsit
func UpdateIfExsit(
	feilds []string,
	db *gorm.DB, // 数据库连接
	query interface{}, // 查询条件
	new interface{}, // 新的
	list interface{}, // 列表
) error {
	chain := commonQuery(feilds, db, query, &GetOptions{}, list)
	check := chain.Find(list, query)
	if check == nil {
		return errors.New("db error")
	}
	byted, err := json.Marshal(list)
	if err != nil {
		return errors.New("db error")
	}

	if len(byted) > 2 { // exsit
		Put(db, query, new)
	} else { // new
		Post(db, new)
	}
	return nil
}

// Post 通用添加数据
func Post(
	db *gorm.DB,
	entity interface{},
) error {
	check := db.Create(entity)
	if check == nil {
		return errors.New("db error")
	}
	return nil
}

// Put 通用数据库put方法
func Put(
	db *gorm.DB,
	query interface{},
	new interface{},
) error {
	check := db.Model(query).Updates(new)
	if check == nil {
		return errors.New("db error")
	}
	return nil
}

// Delete 通用的数据库删除方法
func Delete(
	db *gorm.DB,
	queryOnlyID interface{},
) error {
	check := db.Delete(queryOnlyID)
	if check == nil {
		return errors.New("db error")
	}
	return nil
}

// Count 通用计数
func Count(
	db *gorm.DB, // 数据库连接
	query interface{}, // 查询条件
	getOptions *GetOptions, // 过滤条件二进制
	list interface{}, // 列表
) (int, error) {

	chain := commonQuery([]string{}, db, query, getOptions, list)
	count := 0
	// 这里有个问题是只有调用Find，gorm才识别真正的Model，否则识别不到表，暂且用limit1，保证性能
	check := chain.Where(query).Limit(1).Find(list).Limit(-1).Count(&count)
	if check == nil {
		return 0, errors.New("db error")
	}
	return count, nil
}

func commonQuery(
	feilds []string,
	db *gorm.DB, // 数据库连接
	query interface{}, // 查询条件
	getOptions *GetOptions, // 过滤条件二进制
	list interface{}, // 列表
) *gorm.DB {

	fds := "*"
	if len(feilds) > 0 {
		fds = strings.Join(feilds, ",")
	}

	println("fdseeee", fds)

	chain := db.Select(fds)
	if getOptions.Filter != nil {
		for _, item := range getOptions.Filter {
			filtername := item[0].(string)
			filtervalue := item[1]
			filtertype := item[2].(string)

			// 如果未指定前缀和后缀，表示contain
			if filtertype == "like" {
				if !strings.Contains(filtervalue.(string), "%") {
					filtervalue = "%" + filtervalue.(string) + "%"
				}
			}
			if filtername != "" && filtervalue != "" && filtertype != "" {
				if filtertype == "like" {
					chain = chain.Where(filtername+" LIKE ? ", filtervalue)
				} else if filtertype == "eq" {
					chain = chain.Where(filtername+" = ?", filtervalue)
				} else if filtertype == "lt" {
					chain = chain.Where(filtername+" < ?", filtervalue)
				} else if filtertype == "gt" {
					chain = chain.Where(filtername+" > ?", filtervalue)
				} else if filtertype == "lte" {
					chain = chain.Where(filtername+" <= ?", filtervalue)
				} else if filtertype == "gte" {
					chain = chain.Where(filtername+" >= ?", filtervalue)
				} else if filtertype == "not" {
					chain = chain.Where(filtername+" <> ?", filtervalue)
				}
			}
		}
	}

	if getOptions.Sort == nil || len(getOptions.Sort) == 0 {
		chain = chain.Order("id desc")
	} else {
		for _, item := range getOptions.Sort {
			sortcol := item[0]
			sorttype := item[1]

			if sorttype == "ascend" || sorttype == "asc" {
				sorttype = "asc"
			} else if sorttype == "descend" || sorttype == "desc" {
				sorttype = "desc"
			} else {
				sorttype = ""
			}

			if sorttype != "" {
				chain = chain.Order(sortcol + " " + sorttype)
			}

		}
	}

	// page start from 1
	if getOptions.Page != 0 && getOptions.PageCount != 0 {
		chain = chain.Limit(getOptions.PageCount).Offset(getOptions.PageCount * (getOptions.Page - 1))
	}

	if getOptions.IN != nil {
		for col, values := range getOptions.IN {
			chain = chain.Where(col+" IN (?)", values)
		}
	}

	return chain
}
