package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var validFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func main() {
	c, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %+v", err)
	}

	passports := strings.Split(string(c), "\n\n")

	validPassports := 0

	for _, passport := range passports {
		passport = strings.ReplaceAll(passport, "\n", " ")

		for i, f := range validFields {
			if !strings.Contains(passport, f) {
				break
			} else if i == 6 {
				if validatePassport(passport) {
					validPassports++
				}
			}
		}
	}

	fmt.Println(validPassports)
}

func validatePassport(passport string) bool {
	fields := strings.Fields(passport)

	isValid := true

	for _, field := range fields {
		kv := strings.Split(field, ":")
		if validator, ok := fieldValidators[kv[0]]; ok {
			isValid = validator(kv[1])
			if !isValid {
				return false
			}
		}
	}

	return isValid
}

var fieldValidators = map[string]func(v string) bool{
	"byr": func(v string) bool {
		y, _ := strconv.Atoi(v)
		return !(y < 1920 || y > 2002)
	},
	"iyr": func(v string) bool {
		y, _ := strconv.Atoi(v)
		return !(y < 2010 || y > 2020)
	},
	"eyr": func(v string) bool {
		y, _ := strconv.Atoi(v)
		return !(y < 2020 || y > 2030)
	},
	"hgt": func(v string) bool {
		if strings.Contains(v, "in") {
			hs := strings.TrimSuffix(v, "in")
			h, _ := strconv.Atoi(hs)
			return !(h < 59 || h > 76)
		}
		if strings.Contains(v, "cm") {
			hs := strings.TrimSuffix(v, "cm")
			h, _ := strconv.Atoi(hs)
			return !(h < 150 || h > 193)
		}
		return false
	},
	"hcl": func(v string) bool {
		matched, err := regexp.MatchString("#([0-9a-f]){6}", v)
		return matched && err == nil
	},
	"ecl": func(v string) bool {
		eyeColours := map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {}}
		_, ok := eyeColours[v]
		return ok
	},
	"pid": func(v string) bool {
		matched, err := regexp.MatchString(`(\d{9})`, v)
		return matched && err == nil && len(v) == 9
	},
}
