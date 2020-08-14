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

func TestSix910API_AddOrder(t *testing.T) {
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

	var crt sdbi.Order
	crt.BillingAddress = "123"
	crt.CustomerID = 18
	crt.OrderNumber = "1234"
	crt.OrderType = "delivery"
	crt.Subtotal = 100.25
	crt.Total = 100.25

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddOrder(&crt, &head)

	fmt.Println("AddOrder: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_AddOrderFail(t *testing.T) {
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

	var crt sdbi.Order
	crt.BillingAddress = "123"
	crt.CustomerID = 18
	crt.OrderNumber = "1234"
	crt.OrderType = "delivery"
	crt.Subtotal = 100.25
	crt.Total = 100.25

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.AddOrder(&crt, &head)

	fmt.Println("AddOrder: ", res)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_UpdateOrder(t *testing.T) {
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

	var crt sdbi.Order
	crt.ID = 16
	crt.BillingAddress = "123"
	crt.CustomerID = 18
	crt.OrderNumber = "1234"
	crt.OrderType = "delivery"
	crt.Subtotal = 100.25
	crt.Taxes = 5.25
	crt.Total = 105.50

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateOrder(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_UpdateOrderFail(t *testing.T) {
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

	var crt sdbi.Order
	crt.ID = 16
	crt.BillingAddress = "123"
	crt.CustomerID = 18
	crt.OrderNumber = "1234"
	crt.OrderType = "delivery"
	crt.Subtotal = 100.25
	crt.Taxes = 5.25
	crt.Total = 105.50

	var head Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//head.Set("localDomain", "defaultLocalStore.mydomain.com")

	res := api.UpdateOrder(&crt, &head)

	if !res.Success {
		t.Fail()
	}
}

func TestSix910API_GetOrder(t *testing.T) {
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

	res := api.GetOrder(16, &head)
	fmt.Println("GetOrder in get: ", *res)

	if res.ID == 0 {
		t.Fail()
	}
}

func TestSix910API_GetOrderList(t *testing.T) {
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

	res := api.GetOrderList(18, &head)
	fmt.Println("GetOrderList in get: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_GetStoreOrderList(t *testing.T) {
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

	res := api.GetStoreOrderList(&head)
	fmt.Println("GetStoreOrderList in get: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_GetStoreOrderListByStatus(t *testing.T) {
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

	res := api.GetStoreOrderListByStatus("processing", &head)
	fmt.Println("GetOrderList in get: ", *res)

	if (*res)[0].ID == 0 {
		t.Fail()
	}
}

func TestSix910API_DeleteOrder(t *testing.T) {
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

	res := api.DeleteOrder(16, &head)
	fmt.Println("DeleteOrder: ", *res)

	if !res.Success {
		t.Fail()
	}
}
