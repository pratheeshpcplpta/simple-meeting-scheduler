package main

import (
	"fmt"

	"github.com/pratheeshpcplpta/simple-meeting-scheduler/helper"
)

func main() {
	TestEncrypDecryp()
}
func TestEncrypDecryp() {
	enc := helper.AES_Encrypt("test", "login")
	fmt.Println(enc)

	dec := helper.AES_Decrypt(enc, "login")
	fmt.Println(dec)

	// enc := helper.Base64Encode([]byte("test123456789"))
	// fmt.Println(string(enc))

	// dec, _ := helper.Base64Decode([]byte(enc))
	// fmt.Println(string(dec))

}
