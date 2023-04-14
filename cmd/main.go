// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"encoding/base64"
	"log"
	"net/http"
	"time"
	"errors"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/lib/pq"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	dbConn *gorm.DB
)

func init() {
   // err is pre-declared to avoid shadowing client.
   var err error

    dbConn, err = gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN: "user=postgres port=5432 password=1234 host=plagiarism-checker-377309:us-central1:pl-checker-psql dbname=plagiarism_checker sslmode=disable",
	}))
	if err != nil {
		log.Fatal(err)
	}

   log.Println("GetPlagiarismReportsList fn inited")
}



func GetPlagiarismReportsList(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Write([]byte(""))
		return
	}

	tokenParam := r.Header.Get("Authorization")
	log.Println(tokenParam)

	// if tokenParam == "" {
	// 	log.Println("ERROR: authorization token is empty")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(Response{ Errors: []ResponseError{ { Type: "error", Key: "authorization_failed" } } })
	// 	return
	// }

	sDec, _ := base64.StdEncoding.DecodeString(strings.Split(tokenParam, ".")[1])
    log.Println(string(sDec))
    log.Println()

	if err := verifyToken(tokenParam, "fd583962-3439-4226-9729-2b00c41b343c"); err != nil {
		log.Println("ERROR: ", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{ Error: &ResponseError{ Type: "error", Key: "authorization_failed" } })
		return
	}

	reportsList, err := getReportsList()
	if err != nil {
		log.Println("ERROR ->", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{ Error: &ResponseError{ Type: "error", Key: "internal_server_error" } })
		return
	}
	
	if err := json.NewEncoder(w).Encode(Response{ Data: reportsList }); err != nil {
		log.Println("ERROR ->", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{ Error: &ResponseError{ Type: "error", Key: "internal_server_error" } })
		return
	}
}

func verifyToken(tokenParam, key string) error {
	if tokenParam == "" {
		return errors.New("authorization token is required")
	}

	token, err := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("verify token error")
		}
		return []byte(key), nil
	})
	if !token.Valid {
		err = errors.New("auth token is invalid")
	}
	return err
}

type ReportsList []ReportsListItem
type ReportsListItem struct {
	ID string `json:"id"`
	CreatedAt time.Time `json:"-"`
	CreatedAtPresenter int `json:"created_at"`
  	Title string `json:"title"`
  	PlagiarisedPercent   int `json:"plagiarised_percent"`
	Seen time.Time `json:"-"`
	SeenPresenter int `json:"seen"`
  	State string `json:"state"`
}

type ResponseError struct {
	Type string `json:"type"`
	Key string `json:"key"`
}

type Response struct {
	Error *ResponseError `json:"error,omitempty"`
	Data any `json:"data"`
}

func getReportsList() (ReportsList, error) {
	var (
		r ReportsListItem
		rr = make(ReportsList, 0)
	)

	rows, err := dbConn.Table("tasks as t").
		Select("t.id, t.created_at, t.title, coalesce(r.plagiarised_percent, 0) as plagiarised_percent, coalesce(r.seen, to_timestamp(0)) as seen, t.state").
		Joins("left join reports as r on r.task_id = t.id").
		Order("t.created_at desc").
		Limit(5).
		Rows()
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&r.ID, &r.CreatedAt, &r.Title, &r.PlagiarisedPercent, &r.Seen, &r.State); err != nil {
			log.Println(err)
		}
		log.Println("report line --->>>", r)
		r.SeenPresenter = int(r.Seen.Unix())
		r.CreatedAtPresenter = int(r.CreatedAt.Unix())
		r.Title = strings.ReplaceAll(r.Title, `\n`, " ")
		rr = append(rr, r)
	}
	
	//select r.created_at, t.title, r.plagiarised_percent, r.seen, t.state from reports as r join tasks as t  on r.task_id = t.id order by r.created_at desc limit 5;

	return rr, nil
} 