package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Amount struct {
	Dollars int
	Cents   int
}

func ParseMoney(strInput string, money *Amount) (err error) {

	strInput = strings.TrimPrefix(strInput, "$")
	strInput = strings.Replace(strInput, ",", "", -1)

	parts := strings.Split(strInput, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid money format")
	}

	dollars, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("error parsing dollars: %v", err)
	}

	cents, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("error parsing cents: %v", err)
	}

	money.Cents = cents
	money.Dollars = dollars

	return nil
}

func AddMoney(m1, m2 Amount) Amount {
	totalCents := m1.Cents + m2.Cents
	dollars := m1.Dollars + m2.Dollars + totalCents/100
	cents := totalCents % 100
	return Amount{Dollars: dollars, Cents: cents}
}
