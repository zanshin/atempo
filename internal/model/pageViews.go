package model

import (
	"database/sql"
	_ "database/sql"

	l "github.com/zanshin/atempo/internal/logger"
)

type PageView struct {
	Domain       string `json:"domain"`
	IPAddress    string `json:"ipAddress"`
	URL          string `json:"url"`
	UserAgent    string `json:"userAgent"`
	ScreenHeight int    `json:"screenHeight"`
	ScreenWidth  int    `json:"screenWidth"`
}

type PageViews struct {
	PageViews []PageView `json:"pageView"`
}

func GetPageViews(dbConn sql.DB) PageViews {
	sqlStatement := "SELECT domain, ipAddress, url, userAgent, screenHeight, screenWidth"

	rows, err := dbConn.Query(sqlStatement)
	if err != nil {
		l.Error.Printf("Unable to select PageViews. Reason: %s", err)
		return PageViews{}
	}

	defer rows.Close()

	result := PageViews{}
	for rows.Next() {
		pv := PageView{}

		err = rows.Scan(&pv.Domain, &pv.IPAddress, &pv.URL, &pv.UserAgent, &pv.ScreenHeight, &pv.ScreenWidth)
		if err != nil {
			l.Error.Printf("Unable to scan row into PaveView struct. Reason: %s", err)
			return PageViews{}
		}

		result.PageViews = append(result.PageViews, pv)
	}

	return result

}
