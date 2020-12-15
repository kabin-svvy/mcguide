package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	romanGuide = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	// Material price
	materialGuide = map[string]float64{}

	commandGuild = map[string]string{}

	// Merchant number
	merchantGuide = map[string]int{}

	// ShoutGuide
	shoutGuide = map[string]string{
		"NOIDEA":  "I have no idea what you are talking about",
		"HOWMUCH": "%s is %d",
		"HOWMANY": "%s %s is %s Credits",
	}
)

func mcguide(input string) (msg string, err error) {

	s := strings.Split(input, " ")

	// fmt.Print(s)

	if len(s) >= 3 && s[0] != "how" && s[1] == "is" && romanGuide[s[2]] > 0 {
		// fmt.Println("updateMerchantGuide")
		err = updateMerchantGuide(s)
	} else if len(s) >= 4 && s[0] != "how" && s[len(s)-1] == "Credits" && s[len(s)-3] == "is" {
		// fmt.Println("updateMaterialGuide")
		err = updateMaterialGuide(s)
	} else if len(s) >= 5 && s[len(s)-1] == "?" && s[0] == "how" && s[1] == "much" && s[2] == "is" {
		// fmt.Println("callHowMuch")
		msg, err = callHowMuch(s)
	} else if len(s) >= 6 && s[len(s)-1] == "?" && s[0] == "how" && s[1] == "many" && s[2] == "Credits" && s[3] == "is" && materialGuide[s[len(s)-2]] > 0 {
		// fmt.Println("callHowMany")
		msg, err = callHowMany(s)
	} else {
		msg = shoutGuide["NOIDEA"]
		err = errors.New(msg)
	}

	// fmt.Println(merchantGuide)
	// fmt.Println(materialGuide)
	// fmt.Println(msg)

	if err != nil {
		return err.Error(), err
	}

	return msg, nil
}

func updateMerchantGuide(s []string) error {

	key := s[0]
	value := romanGuide[s[2]]

	// Update
	for k1, v1 := range merchantGuide {
		if key == k1 && value == v1 {
			break
		} else if key == k1 && value != v1 {
			for k2, v2 := range merchantGuide {
				if key != k2 && value == v2 {
					return errors.New("Duplicate symbol in merchant guide")
				}
			}
		} else if value == v1 && key != k1 {
			return errors.New("Duplicate value in merchant guide")
		}
	}

	// Insert
	merchantGuide[key] = value

	// fmt.Println(merchantGuide)

	return nil
}

func updateMaterialGuide(s []string) error {

	credits := s[len(s)-2]

	f, err := strconv.ParseFloat(credits, 64)

	if err != nil {
		return err
	}

	merchantPrice := 0

	if len(s) > 4 {
		_, merchantPrice = convertMerchantNumber(s, 0, len(s)-5)
		if merchantPrice == -1 {
			return errors.New(shoutGuide["NOIDEA"])
		}
	}

	p := float64(merchantPrice)

	if err != nil {
		return err
	}

	if p != 0 {
		f /= p
	}

	materialName := s[len(s)-4]

	// Insert
	materialGuide[materialName] = f

	// fmt.Println(materialGuide)

	return nil
}

func callHowMuch(s []string) (string, error) {
	symbol, p := convertMerchantNumber(s, 3, len(s)-2)
	if p > 0 {
		return fmt.Sprintf(shoutGuide["HOWMUCH"], symbol, p), nil
	}
	return shoutGuide["NOIDEA"], errors.New(shoutGuide["NOIDEA"])
}

func callHowMany(s []string) (string, error) {

	materialName := s[len(s)-2]
	materialPrice := materialGuide[materialName]

	if materialPrice == 0 {
		return shoutGuide["NOIDEA"], errors.New(shoutGuide["NOIDEA"])
	}

	symbol := ""
	merchantPrice := 0

	if len(s) > 6 {
		symbol, merchantPrice = convertMerchantNumber(s, 4, len(s)-3)
		if merchantPrice == -1 {
			return shoutGuide["NOIDEA"], errors.New(shoutGuide["NOIDEA"])
		}
	}

	if merchantPrice > 0 {
		materialPrice *= float64(merchantPrice)
	}

	if materialPrice > 0 {
		return fmt.Sprintf(shoutGuide["HOWMANY"], symbol, materialName, strconv.FormatFloat(materialPrice, 'f', -1, 64)), nil
	}
	return shoutGuide["NODIEA"], errors.New(shoutGuide["NODIEA"])
}

func convertMerchantNumber(s []string, startindex int, lastindex int) (string, int) {
	symbol := ""
	r := 0
	n := 0
	p := 0
	for i := startindex; i <= lastindex; i++ {
		n = merchantGuide[s[i]]
		if n == 0 {
			return "", -1
		}

		if i > 0 {
			if p < n {
				r -= p
				p = n
			} else {
				r += p
				p = n
			}
		} else {
			p = merchantGuide[s[i]]
		}
		if i == lastindex {
			r += n
			symbol = symbol + s[i]
		} else {
			symbol = symbol + s[i] + " "
		}
		if 0 == lastindex {
			r += n
		}
	}
	return symbol, r
}
