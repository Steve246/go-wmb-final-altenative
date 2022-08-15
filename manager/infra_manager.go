package manager

import (
	"fmt"
	"livecode-wmb-2/config"
	"livecode-wmb-2/model"
	"livecode-wmb-2/service"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
	lopeiClientConn() service.LopeiPaymentClient
}

type infra struct {
	db          *gorm.DB
	lopeiClient service.LopeiPaymentClient
	cfg         config.Config
}

// lopeiClientConn implements Infra
func (i *infra) lopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	infras := infra{
		cfg: config,
		db:  resource,
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	infras.initGrpcClient()
	return &infras
}

func (i *infra) initGrpcClient() {
	dial, err := grpc.Dial(i.cfg.GrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect...", err)
	}
	client := service.NewLopeiPaymentClient(dial)
	i.lopeiClient = client
	fmt.Println("GRPC client connected...")
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	env := os.Getenv("ENV")
	if env == "migrate" {
		err := db.Debug().AutoMigrate(
			&model.Customer{},
			&model.Discount{},
			&model.TransType{},
			&model.Table{},
			&model.Menu{},
			&model.MenuPrice{},
			&model.Bill{},
			&model.BillDetail{},
		)
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
