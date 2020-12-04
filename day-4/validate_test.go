package main

import "testing"

func TestValidatePassport(t *testing.T) {
	testCases := []struct {
		passport string
		valid    bool
	}{
		{
			passport: "pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
			valid:    true,
		},
		{
			passport: "eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			valid:    true,
		},
		{
			passport: "hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
			valid:    true,
		}, {
			passport: "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			valid:    true,
		},
		{
			passport: "eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			valid:    false,
		},
		{
			passport: "iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946",
			valid:    false,
		},
		{
			passport: "hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			valid:    false,
		},
		{
			passport: "hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023",
			valid:    false,
		},
		{
			passport: "pid:3556412378 byr:2007",
			valid:    false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.passport, func(t *testing.T) {
			actual := validatePassport(tC.passport)
			if actual != tC.valid {
				t.Fatalf("passport failed validation: %s %t", tC.passport, tC.valid)
			}
		})
	}
}

func TestValidatePID(t *testing.T) {
	testCases := []struct {
		pid   string
		valid bool
	}{
		{
			pid:   "123456789",
			valid: true,
		},
		{
			pid:   "12345",
			valid: false,
		},
		{
			pid:   "1234567890",
			valid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.pid, func(t *testing.T) {
			v, _ := fieldValidators["pid"]
			r := v(tC.pid)
			if tC.valid != r {
				t.Fatalf("pid failed validation: %s, %t", tC.pid, r)
			}
		})
	}
}
