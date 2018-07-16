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

func Get(
	feilds []string,
	db *gorm.DB,
	query interface{},
	getOptions *GetOptions,
	list interface{},
) error {

	chain := commonQuery(feilds, db, query, getOptions, list)
	check := chain.Find(list, query)
	if check == nil {
		return errors.New("db error")
	}
	return nil
}

func UpdateIfExsit(
	feilds []string,
	db *gorm.DB,
	query interface{},
	new interface{},
	list interface{},
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

	if len(byted) > 2 {
		Put(db, query, new)
	} else {
		Post(db, new)
	}
	return nil
}

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

func Count(
	db *gorm.DB,
	query interface{},
	getOptions *GetOptions,
	list interface{},
) (int, error) {

	chain := commonQuery([]string{}, db, query, getOptions, list)
	count := 0
	check := chain.Where(query).Limit(1).Find(list).Limit(-1).Count(&count)
	if check == nil {
		return 0, errors.New("db error")
	}
	return count, nil
}

func commonQuery(
	feilds []string,
	db *gorm.DB,
	query interface{},
	getOptions *GetOptions,
	list interface{},
) *gorm.DB {

	fds := "*"
	if len(feilds) > 0 {
		fds = strings.Join(feilds, ",")
	}

	chain := db.Select(fds)
	if getOptions.Filter != nil {
		for _, item := range getOptions.Filter {
			filtername := item[0].(string)
			filtervalue := item[1]
			filtertype := item[2].(string)

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
