package model

import (
	"database/sql"
	_ "database/sql"

	l "github.com/zanshin/atempo/internal/logger"
)

type HrefClick struct {
	IPAddress  string `json:"ipAddress"`
	URL        string `json:"url"`
	Href       string `json:"href"`
	HrefTop    int    `json:"hrefTop"`
	HrefRight  int    `json:"hrefRight"`
	HrefBottom int    `json:"hrefBottom"`
	HrefLeft   int    `json:"hrefLeft"`
}

type HrefClicks struct {
	HrefClicks []HrefClick `json:"hrefClick"`
}

func GetHrefClicks(dbConn sql.DB) HrefClicks {
	sqlStatement := "SELECT ipAddress, url, href, hrefTop, hrefRight, hrefBottom, hrefLeft"

	rows, err := dbConn.Query(sqlStatement)
	if err != nil {
		l.Error.Printf("Unable to select HrefClicks. Reason: %s", err)
		return HrefClicks{}
	}

	defer rows.Close()

	result := HrefClicks{}
	for rows.Next() {
		hc := HrefClick{}

		err = rows.Scan(&hc.IPAddress, &hc.URL, &hc.Href, &hc.HrefTop, &hc.HrefRight, &hc.HrefBottom, &hc.HrefLeft)
		if err != nil {
			l.Error.Printf("Unable to scan row into PaveView struct. Reason: %s", err)
			return HrefClicks{}
		}

		result.HrefClicks = append(result.HrefClicks, hc)
	}

	return result

}
