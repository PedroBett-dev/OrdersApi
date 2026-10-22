package etc

import (
	"crypto/sha256"
	"encoding/base32"
	"regexp"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"github.com/paemuri/brdoc"
)

var nonDigits = regexp.MustCompile(`\D`)

func OnlyDigits(s string) string {
	return nonDigits.ReplaceAllString(s, "")
}

func NewId(s string, len int) string {
	sum := sha256.Sum256([]byte(s))
	return base32.StdEncoding.EncodeToString(sum[:])[:len]
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func IsPhoneValid(s string) bool {
	num, err := phonenumbers.Parse(s, "BR")
	if err != nil || !phonenumbers.IsValidNumber(num) || phonenumbers.GetRegionCodeForNumber(num) != "BR" {
		return false
	}

	return phonenumbers.GetNationalSignificantNumber(num) == s
}

var ufs = map[string]brdoc.FederativeUnit{
	"AC": brdoc.AC, "AL": brdoc.AL, "AP": brdoc.AP, "AM": brdoc.AM,
	"BA": brdoc.BA, "CE": brdoc.CE, "DF": brdoc.DF, "ES": brdoc.ES,
	"GO": brdoc.GO, "MA": brdoc.MA, "MT": brdoc.MT, "MS": brdoc.MS,
	"MG": brdoc.MG, "PA": brdoc.PA, "PB": brdoc.PB, "PR": brdoc.PR,
	"PE": brdoc.PE, "PI": brdoc.PI, "RJ": brdoc.RJ, "RN": brdoc.RN,
	"RS": brdoc.RS, "RO": brdoc.RO, "RR": brdoc.RR, "SC": brdoc.SC,
	"SP": brdoc.SP, "SE": brdoc.SE, "TO": brdoc.TO,
}

// ParseUF converte a sigla (ex: "sp") para o tipo usado pelo brdoc.
func ParseUF(s string) (brdoc.FederativeUnit, bool) {
	uf, ok := ufs[strings.ToUpper(s)]
	return uf, ok
}
