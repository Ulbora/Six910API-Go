package six910api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

func TestSix910API_getStoreID(t *testing.T) {
	var sapi Six910API
	//sapi.SetAPIKey("123")
	//sapi.storeID = 59

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"id":2}`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := sapi.getStoreID(&head)

	if res != 2 {
		t.Fail()
	}
}

func TestSix910API_SetStoreID(t *testing.T) {
	var sapi Six910API

	sapi.GetNew()
	sapi.SetLogLever(lg.AllLevel)

	sapi.SetStoreID(5)
	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	res := sapi.getStoreID(&head)

	if res != 5 {
		t.Fail()
	}
}
