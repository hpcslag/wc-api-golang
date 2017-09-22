package woocommerce

import (
	"testing"
	"./lib"
	"fmt"
	//"bytes"
)


func TestClient(t *testing.T){
	
	b := lib.Utils{}
	s,_ := b.Oauth_generator("CK HERE","CS HERE", "METHOD HERE","FULL URL HERE")
	fmt.Println(s)

	//simulation:
	/*c:= NewClient("http://example.com","ck_XXXX","cs_XXXX")
	body := c.Get("products",nil)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	newStr := buf.String()
	fmt.Printf(newStr)*/

}