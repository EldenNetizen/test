package main

import (
	"fmt"
	"time"

	"github.com/EldenNetizen/test/internal/handler"
	"github.com/EldenNetizen/test/internal/model"
	"github.com/EldenNetizen/test/pkg/utils"
)

func threadTest(ch chan int) {
	for i := range 5 {
		fmt.Println("Thread Test: ", i)
		ch <- i
		time.Sleep(5 * time.Second)
	}
	close(ch)
}

func main() {
	handler := handler.NewAgreementHandler()
	idGetter := utils.NewSnowflakeIDGenerator(1)
	handler.CreateAgreement(model.NewAgreement(model.WithAgreementAmountSum("10000"), model.WithId(uint64(idGetter.Generate()))))
}
