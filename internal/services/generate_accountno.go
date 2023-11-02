package service

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
)

func (a *accountService) GenerateAccountNo(c context.Context, m *model.Account) (err error) {
	rand.Seed(time.Now().UnixNano())
	digit_1 := rand.Intn(10)
	digit_2 := rand.Intn(10)
	digit_3 := rand.Intn(10)
	digit_4 := rand.Intn(10)
	digit_5 := rand.Intn(10)
	digit_6 := rand.Intn(10)
	digit_7 := rand.Intn(10)
	digit_8 := rand.Intn(10)
	digit_9 := rand.Intn(10)
	digit_10 := (digit_1 + digit_2 + digit_3 + digit_4 + digit_5 + digit_6 + digit_7 + digit_8 + digit_9) % 10
	str_digit_1 := strconv.Itoa(digit_1)
	str_digit_2 := strconv.Itoa(digit_2)
	str_digit_3 := strconv.Itoa(digit_3)
	str_digit_4 := strconv.Itoa(digit_4)
	str_digit_5 := strconv.Itoa(digit_5)
	str_digit_6 := strconv.Itoa(digit_6)
	str_digit_7 := strconv.Itoa(digit_7)
	str_digit_8 := strconv.Itoa(digit_8)
	str_digit_9 := strconv.Itoa(digit_9)
	str_digit_10 := strconv.Itoa(digit_10)
	m.AccountNo = (str_digit_1 + str_digit_2 + str_digit_3 + str_digit_4 + str_digit_5 + str_digit_6 + str_digit_7 + str_digit_8 + str_digit_9 + str_digit_10)

	return
}
