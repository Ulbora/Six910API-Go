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

//AddCart AddCart
func (a *Six910API) AddCart(c *sdbi.Cart, headers *Headers) *ResponseID {
	var rtn ResponseID
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/cart/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqca := a.buildRequest(post, url, headers, aJSON)
		succa, stat := a.proxy.Do(reqca, &rtn)
		a.log.Debug("suc: ", succa)
		a.log.Debug("stat: ", stat)
		if !succa {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateCart UpdateCart
func (a *Six910API) UpdateCart(c *sdbi.Cart, headers *Headers) *Response {
	var rtn Response
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/cart/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqcu := a.buildRequest(put, url, headers, aJSON)
		cusuc, stat := a.proxy.Do(reqcu, &rtn)
		a.log.Debug("suc: ", cusuc)
		a.log.Debug("stat: ", stat)
		if !cusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetCart GetCart
func (a *Six910API) GetCart(cid int64, headers *Headers) *sdbi.Cart {
	var rtn sdbi.Cart
	var ctsid = a.getStoreID(headers)
	cidStr := strconv.FormatInt(cid, 10)
	sidStr := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/cart/get/" + cidStr + "/" + sidStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteCart DeleteCart
func (a *Six910API) DeleteCart(id int64, cid int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdc := strconv.FormatInt(id, 10)
	cidStrdc := strconv.FormatInt(cid, 10)
	sidStrdc := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/cart/delete/" + idStrdc + "/" + cidStrdc + "/" + sidStrdc
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dcsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dcsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
