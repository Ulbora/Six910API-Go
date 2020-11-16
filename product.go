package six910api

import (
	"encoding/json"
	"strconv"

	sdbi "github.com/Ulbora/six910-database-interface"
)

/*
 Six910 is a shopping cart and E-commerce system.

 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

//ProdIDReq ProdIDReq
type ProdIDReq struct {
	StoreID      int64    `json:"storeId"`
	CategoryList *[]int64 `json:"categoryList"`
}

//AddProduct AddProduct
func (a *Six910API) AddProduct(p *sdbi.Product, headers *Headers) *ResponseID {
	var rtn ResponseID
	p.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/product/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		sucpda, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", sucpda)
		a.log.Debug("stat: ", stat)
		if !sucpda {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateProduct UpdateProduct
func (a *Six910API) UpdateProduct(p *sdbi.Product, headers *Headers) *Response {
	var rtn Response
	p.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/product/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqpdu := a.buildRequest(put, url, headers, aJSON)
		pdusuc, stat := a.proxy.Do(reqpdu, &rtn)
		a.log.Debug("suc: ", pdusuc)
		a.log.Debug("stat: ", stat)
		if !pdusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateProductQuantity UpdateProductQuantity
func (a *Six910API) UpdateProductQuantity(p *sdbi.Product, headers *Headers) *Response {
	var rtn Response
	p.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/product/update/quantity"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqpduq := a.buildRequest(put, url, headers, aJSON)
		pduqsuc, stat := a.proxy.Do(reqpduq, &rtn)
		a.log.Debug("suc: ", pduqsuc)
		a.log.Debug("stat: ", stat)
		if !pduqsuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetProductByID GetProductByID
func (a *Six910API) GetProductByID(id int64, headers *Headers) *sdbi.Product {
	var rtn sdbi.Product
	var ctsid = a.getStoreID(headers)
	idStrgpd := strconv.FormatInt(id, 10)
	sidStrgod := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/product/get/id/" + idStrgpd + "/" + sidStrgod
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	odgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", odgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductBySku GetProductBySku
func (a *Six910API) GetProductBySku(sku string, did int64, headers *Headers) *sdbi.Product {
	var rtn sdbi.Product
	var ctsid = a.getStoreID(headers)
	didStrgpd := strconv.FormatInt(did, 10)
	sidStrgod := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/product/get/sku/" + sku + "/" + didStrgpd + "/" + sidStrgod
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	odgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", odgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductsByPromoted GetProductsByPromoted
func (a *Six910API) GetProductsByPromoted(start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var prsid = a.getStoreID(headers)
	prstStrgodl := strconv.FormatInt(start, 10)
	prendStrgodl := strconv.FormatInt(end, 10)
	prsidStrGodl := strconv.FormatInt(prsid, 10)

	var url = a.restURL + "/rs/product/get/promoted/" + prsidStrGodl + "/" + prstStrgodl + "/" + prendStrgodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	prpdnlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", prpdnlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductsByName GetProductsByName
func (a *Six910API) GetProductsByName(name string, start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var sid = a.getStoreID(headers)
	stStrgodl := strconv.FormatInt(start, 10)
	endStrgodl := strconv.FormatInt(end, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/name/" + name + "/" + sidStrGodl + "/" + stStrgodl + "/" + endStrgodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdnlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdnlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductsByCaterory GetProductsByCaterory
func (a *Six910API) GetProductsByCaterory(catID int64, start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var sid = a.getStoreID(headers)
	catStrgodl := strconv.FormatInt(catID, 10)
	stStrgodl := strconv.FormatInt(start, 10)
	endStrgodl := strconv.FormatInt(end, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/category/" + catStrgodl + "/" + sidStrGodl + "/" + stStrgodl + "/" + endStrgodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductList GetProductList
func (a *Six910API) GetProductList(start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var sid = a.getStoreID(headers)
	stStrgpdl := strconv.FormatInt(start, 10)
	endStrgpdl := strconv.FormatInt(end, 10)
	sidStrGpdl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/list/" + sidStrGpdl + "/" + stStrgpdl + "/" + endStrgpdl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductIDList GetProductIDList
func (a *Six910API) GetProductIDList(headers *Headers) *[]int64 {
	var rtn []int64
	var sid = a.getStoreID(headers)
	//stStrgpdl := strconv.FormatInt(start, 10)
	//endStrgpdl := strconv.FormatInt(end, 10)
	sidStrGidl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/ids/" + sidStrGidl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductIDListByCategories GetProductIDListByCategories
func (a *Six910API) GetProductIDListByCategories(idReq *ProdIDReq, headers *Headers) *[]int64 {
	var rtn []int64
	idReq.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/product/get/ids/cat"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(idReq)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		sucpdid, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", sucpdid)
		a.log.Debug("stat: ", stat)
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//DeleteProduct DeleteProduct
func (a *Six910API) DeleteProduct(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdpd := strconv.FormatInt(id, 10)
	sidStrdpd := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/delete/" + idStrdpd + "/" + sidStrdpd
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dpdsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dpdsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
