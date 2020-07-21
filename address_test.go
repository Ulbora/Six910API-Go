package six910api

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910API_AddAddress(t *testing.T) {
	var sapi Six910API
	sapi.SetAPIKey("123")
	//sapi.localDomain = ""
	//sapi.storeName = ""
	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	var add sdbi.Address
	add.Address = "test"
	var head Headers
	head.Set("storeName", "defaultLocalStore")
	head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddAddress(&add, &head)

	if !res.Success {
		t.Fail()
	}

}
