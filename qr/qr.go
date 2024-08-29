package qr

import (
	"regexp"
	"strconv"

	"github.com/skip2/go-qrcode"
	"github.com/yousifnimah/Cryptx/CRC16"

	promptpayqr "github.com/kazekim/promptpay-qr-go"
)

type qrPromptPayField struct {
	id       string
	name     string
	value    string
	length   int
	subField []*qrPromptPayField
}

func CreateQRCode(target string, amount int) string {
	amtStt := strconv.Itoa(amount)
	qr := promptpayqr.New().GeneratePayload(target, &amtStt)

	return qr
}

func QRCodeToImage(qrCode string) []byte {
	png, _ := qrcode.Encode(qrCode, qrcode.Medium, 256)
	return png
}

func calChecksum(data string) string {
	Input := []byte(data)          //string to slice of bytes
	AlgorithmName := "CCITT_FALSE" //CRC-8 algorithm name from supported table
	checksumHex := CRC16.ResultHex(Input, AlgorithmName)
	return checksumHex
}

func CreatePrompPayQRCode(target string, amount float64) string {
	payloadFormatIndicator := qrPromptPayField{id: "00", name: "Payload Format Indicator", length: 2, value: "01"}

	var poiMethod qrPromptPayField
	if amount == 0.0 {
		poiMethod = qrPromptPayField{id: "01", length: 2, name: "Point of Initiation Method", value: "11"}
	} else {
		poiMethod = qrPromptPayField{id: "01", length: 2, name: "Point of Initiation Method", value: "12"}
	}

	applicationIDSubField := qrPromptPayField{id: "00", name: "Application ID", length: 16, value: "A000000677010111"}
	promptPayPhoneNumberSubField := qrPromptPayField{id: "01", name: "Prompt phone number", length: 13, value: formatTraget(target)}
	merchantAccountInformation := qrPromptPayField{id: "29", length: 37, name: "Marchant Account Information",
		subField: []*qrPromptPayField{&applicationIDSubField, &promptPayPhoneNumberSubField}}

	thbISOCurrencyCode := "764"
	currency := qrPromptPayField{id: "53", name: "Transaction Currency", length: 3, value: thbISOCurrencyCode}
	var transactionAmt qrPromptPayField
	if amount != 0.0 {
		amountVal := formatAmount(amount)
		transactionAmt = qrPromptPayField{id: "54", name: "Transcation amount", length: len(amountVal), value: amountVal}
	}
	country := qrPromptPayField{id: "58", name: "Country Code", length: 2, value: "TH"}

	qr := serialize(payloadFormatIndicator) + serialize(poiMethod) + serialize(merchantAccountInformation) + serialize(transactionAmt) + serialize(country) + serialize(currency)
	checksum := qrPromptPayField{id: "63", name: "CRC 16 Chcksum", length: 4}
	qr = qr + serialize(checksum)
	qr = qr + calChecksum(qr)[2:6]

	return qr
}

func serialize(field qrPromptPayField) string {
	if field.id == "" {
		return ""
	}
	if field.subField == nil {
		return field.id + formatFieldLengthStr(field.length) + field.value
	} else {
		var subFieldOutput string
		for _, sub := range field.subField {
			subFieldOutput = subFieldOutput + serialize(*sub)
		}
		return field.id + formatFieldLengthStr(field.length) + subFieldOutput
	}
}

func formatAmount(amount float64) string {
	return strconv.FormatFloat(amount, 'f', 2, 64)
}

func formatFieldLengthStr(l int) string {
	var lengthStr string
	if l < 10 {
		lengthStr = "0" + strconv.Itoa(l)
	} else {
		lengthStr = strconv.Itoa(l)
	}
	return lengthStr
}

func formatTraget(target string) string {
	allNumCheck := regexp.MustCompile(`^[0-9]`)
	if allNumCheck.MatchString(target) {
		if len(target) == 10 { // is phone number
			return "0066" + target[1:]
		}
		if len(target) == 13 { // is citizen number
			return target
		}
	}
	return ""
}
