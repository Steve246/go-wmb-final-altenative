package repository

import (
	"context"
	"fmt"
	"livecode-wmb-2/service"
	"livecode-wmb-2/utils"
	"log"
	"strconv"
	"strings"
)

type LopeiRepository interface {
	CheckBalance(lopeId int32) (float32, error)
	DoPayment(lopeId int32, amount float32) error
}

type lopeiRepository struct {
	client service.LopeiPaymentClient
}

// CheckBalance implements CustomerRepository
func (c *lopeiRepository) CheckBalance(lopeId int32) (float32, error) {
	balance, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{
		LopeId: lopeId,
	})
	if err != nil {
		log.Fatalln("error when calling check balance...", err)
	}
	ini := strings.Split(balance.GetResult(), ",")
	_, number := utils.ParseData(ini[1])
	nah, _ := strconv.Atoi(number)
	// fmt.Println(nah)
	return float32(nah), err
}

// DoPayment implements CustomerRepository
func (c *lopeiRepository) DoPayment(lopeId int32, amount float32) error {
	payment, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		LopeId: lopeId,
		Amount: amount,
	})
	if err != nil {
		log.Fatalln("error when calling do payment...", err)
	}
	fmt.Println(payment)
	return nil
}

func NewLopeiRepository(client service.LopeiPaymentClient) LopeiRepository {
	repo := new(lopeiRepository)
	repo.client = client
	return repo
}
