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

func TestSix910API_AddZoneZip(t *testing.T) {
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

	var crt sdbi.ZoneZip
	crt.ExcludedSubRegionID = 2
	crt.ZipCode = "1234588"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddZoneZip(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_AddZoneZipFail(t *testing.T) {
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

	var crt sdbi.ZoneZip
	crt.ExcludedSubRegionID = 2
	crt.ZipCode = "1234588"

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddZoneZip(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_GetZoneZipListByExclusion(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"id":1}]`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	res := api.GetZoneZipListByExclusion(2, &head)
	fmt.Println("GetZoneZipListByExclusion: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_GetZoneZipListByInclusion(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"id":1}]`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	res := api.GetZoneZipListByInclusion(2, &head)
	fmt.Println("GetZoneZipListByInclusion: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_DeleteZoneZip(t *testing.T) {
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

	res := api.DeleteZoneZip(5, 2, 0, &head)
	fmt.Println("DeleteZoneZip: ", *res)

	if !res.Success {
		t.Fail()
	}
}
