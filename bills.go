package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/augustyip/bills/services"
	log "github.com/sirupsen/logrus"
)

// Certification struct
type Certification struct {
	Service  string
	Username string
	Password string
}

// Bill struct
type Bill struct {
	AccountNo       string
	Balance         string
	LastPaymentDate string
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {

	file, _ := os.Open("cert.json")
	decoder := json.NewDecoder(file)
	certs := make([]Certification, 0)
	err := decoder.Decode(&certs)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, cert := range certs {

		switch s := cert.Service; s {
		case "towngas":
			log.Info("Starting to run towngas service...")
			towngas := services.Towngas{cert.Username, cert.Password}
			r := services.GetNewsNoticeAsync(towngas)
			// fmt.Printf(r)
			log.Info(r)

		case "clp":
			clp := services.Clp{cert.Username, cert.Password}
			var clpBill services.Bill
			clpBill.GetServiceDashboard(clp)
			log.Info(clpBill)

		case "wsd":
			wsd := services.Wsd{cert.Username, cert.Password}
			r := services.ElectronicBill(wsd)
			// fmt.Printf(r)
			log.Info(r)

		}
	}
}
