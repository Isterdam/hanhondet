package main

import (
	"database/sql"
	"fmt"
)

type Page interface {
	SearchInfo(info *Info) ([]*Info, error)
}

var page Page

type dbPage struct {
	db *sql.DB
}

func (page *dbPage) SearchInfo(info *Info) ([]*Info, error) {
	_, err := page.db.Query("UPDATE infos SET sokningar = sokningar + 1 WHERE ord = $1", info.Ord)
	if err != nil {
		fmt.Errorf("Could not increment searches!")
	}

	rows, err := page.db.Query("SELECT ord, genus, sokningar from infos WHERE ord = $1", info.Ord)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	infos := []*Info{}
	for rows.Next() {
		info := &Info{}
		if err := rows.Scan(&info.Ord, &info.Genus, &info.Sokningar); err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	if len(infos) == 0 {
		return infos, fmt.Errorf("This word could not be found!")
	}
	return infos, nil
}

func InitPage(p Page) {
	page = p
}
