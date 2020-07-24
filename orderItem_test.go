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

func TestSix910API_AddOrderItem(t *testing.T) {
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

	var crt sdbi.OrderItem
	crt.OrderID = 15
	crt.ProductID = 7
	crt.ProductName = "test"
	crt.Quantity = 5

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddOrderItem(&crt, &head)

	fmt.Println("AddOrderItem: ", res)

	if !res.Success {
		t.Fail()
	}

}

func TestSix910API_AddOrderItemFail(t *testing.T) {
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

	var crt sdbi.OrderItem
	crt.OrderID = 15
	crt.ProductID = 7
	crt.ProductName = "test"
	crt.Quantity = 5

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddOrderItem(&crt, &head)

	fmt.Println("AddOrderItem: ", res)

	if !res.Success {
		t.Fail()
	}

}

func TestSix910API_UpdateOrderItem(t *testing.T) {
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

	var crt sdbi.OrderItem
	crt.ID = 5
	crt.OrderID = 15
	crt.ProductID = 7
	crt.ProductName = "test"
	crt.Quantity = 10

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateOrderItem(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_UpdateOrderItemFail(t *testing.T) {
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

	var crt sdbi.OrderItem
	crt.ID = 5
	crt.OrderID = 15
	crt.ProductID = 7
	crt.ProductName = "test"
	crt.Quantity = 10

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateOrderItem(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_GetOrderItem(t *testing.T) {
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

	res := api.GetOrderItem(5, &head)
	fmt.Println("GetOrderItem in get: ", *res)

	if res.ID == 0 {
		t.Fail()
	}
}

func TestSix910API_GetOrderItemList(t *testing.T) {
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

	res := api.GetOrderItemList(15, &head)
	fmt.Println("GetOrderItemList in get: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_DeleteOrderItem(t *testing.T) {
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

	res := api.DeleteOrderItem(6, &head)
	fmt.Println("DeleteOrderItem: ", *res)

	if !res.Success {
		t.Fail()
	}
}
