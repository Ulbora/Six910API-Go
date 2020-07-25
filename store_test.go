package six910api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910API_AddStore(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":1}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var crt sdbi.Store
	crt.Company = "tester"
	crt.Email = "tester@tester.com"
	crt.StoreName = "teststore"
	crt.LocalDomain = "teststore.teststore.com"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddStore(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_AddStoreFail(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":1}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var crt sdbi.Store
	crt.Company = "tester"
	crt.Email = "tester@tester.com"
	crt.StoreName = "teststore"
	crt.LocalDomain = "teststore.teststore.com"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddStore(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_UpdateStore(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":1}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var crt sdbi.Store
	crt.ID = 59
	crt.Company = "tester"
	crt.Email = "tester@tester.com"
	crt.StoreName = "teststore"
	crt.LocalDomain = "teststore.teststore.com"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateStore(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_UpdateStoreFail(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":1}`))
	gp.MockResp = &mres
	//gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var crt sdbi.Store
	crt.ID = 59
	crt.Company = "tester"
	crt.Email = "tester@tester.com"
	crt.StoreName = "teststore"
	crt.LocalDomain = "teststore.teststore.com"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateStore(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_GetStore(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"id":1}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	res := api.GetStore("defaultLocalStore", "defaultLocalStore.mydomain.com", &head)
	fmt.Println("GetStore in get: ", *res)

	if res.ID == 0 {
		t.Fail()
	}
}

func TestSix910API_DeleteStore(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	api := sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"success":true, "id":1}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	res := api.DeleteStore("defaultLocalStore", "defaultLocalStore.mydomain.com", &head)
	fmt.Println("DeleteStore: ", *res)

	if !res.Success {
		t.Fail()
	}
}
