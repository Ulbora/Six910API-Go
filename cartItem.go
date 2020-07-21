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

//CartItemReq CartItemReq
type CartItemReq struct {
	CustomerID int64          `json:"customerId"`
	StoreID    int64          `json:"storeId"`
	CartItem   *sdbi.CartItem `json:"cartItem"`
}

//AddCartItem AddCartItem
func (a *Six910API) AddCartItem(ci *sdbi.CartItem, cid int64, headers *Headers) *ResponseID {
	var rtn ResponseID
	var crtia CartItemReq
	crtia.CartItem = ci
	crtia.CustomerID = cid
	crtia.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/cartItem/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(crtia)
	if err == nil {
		reqcia := a.buildRequest(post, url, headers, aJSON)
		succia, stat := a.proxy.Do(reqcia, &rtn)
		a.log.Debug("suc: ", succia)
		a.log.Debug("stat: ", stat)
		if !succia {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateCartItem UpdateCartItem
func (a *Six910API) UpdateCartItem(ci *sdbi.CartItem, cid int64, headers *Headers) *Response {
	var rtn Response
	var crtiu CartItemReq
	crtiu.CartItem = ci
	crtiu.CustomerID = cid
	crtiu.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/cartItem/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(crtiu)
	if err == nil {
		reqciu := a.buildRequest(put, url, headers, aJSON)
		ciusuc, stat := a.proxy.Do(reqciu, &rtn)
		a.log.Debug("suc: ", ciusuc)
		a.log.Debug("stat: ", stat)
		if !ciusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetCartItem GetCartItem
func (a *Six910API) GetCartItem(cid int64, prodID int64, headers *Headers) *sdbi.CartItem {
	var rtn sdbi.CartItem
	var citsid = a.getStoreID(headers)
	cidStrcit := strconv.FormatInt(cid, 10)
	prodIDStr := strconv.FormatInt(prodID, 10)
	sidStr := strconv.FormatInt(citsid, 10)
	var url = a.restURL + "/rs/cartItem/get/" + cidStrcit + "/" + prodIDStr + "/" + sidStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cigsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cigsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

//GetCartItemList GetCartItemList
func (a *Six910API) GetCartItemList(cartID int64, cid int64, headers *Headers) *[]sdbi.CartItem {
	var rtn []sdbi.CartItem
	var sid = a.getStoreID(headers)
	cartIDStrGcil := strconv.FormatInt(cartID, 10)
	cidStrGcil := strconv.FormatInt(cid, 10)
	sidStrGcil := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/cartItem/get/list/" + cartIDStrGcil + "/" + cidStrGcil + "/" + sidStrGcil
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteCartItem DeleteCartItem
func (a *Six910API) DeleteCartItem(id int64, prodID int64, cartID int64, headers *Headers) *Response {
	var rtn Response
	//var sid = a.getStoreID(headers)
	idStrdci := strconv.FormatInt(id, 10)
	prodIDStrdci := strconv.FormatInt(prodID, 10)
	cartIDStrdci := strconv.FormatInt(cartID, 10)
	//sidStrdc := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/cartItem/delete/" + idStrdci + "/" + prodIDStrdci + "/" + cartIDStrdci
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dcisuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dcisuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
