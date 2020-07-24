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

//OrderItemReq OrderItemReq
type OrderItemReq struct {
	StoreID   int64           `json:"storeId"`
	OrderItem *sdbi.OrderItem `json:"orderItem"`
}

//AddOrderItem AddOrderItem
func (a *Six910API) AddOrderItem(i *sdbi.OrderItem, headers *Headers) *ResponseID {
	var rtn ResponseID
	var oir OrderItemReq
	oir.OrderItem = i
	oir.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/orderItem/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(oir)
	if err == nil {
		reqoia := a.buildRequest(post, url, headers, aJSON)
		sucoia, stat := a.proxy.Do(reqoia, &rtn)
		a.log.Debug("suc: ", sucoia)
		a.log.Debug("stat: ", stat)
		if !sucoia {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateOrderItem UpdateOrderItem
func (a *Six910API) UpdateOrderItem(i *sdbi.OrderItem, headers *Headers) *Response {
	var rtn Response
	var oir OrderItemReq
	oir.OrderItem = i
	oir.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/orderItem/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(oir)
	if err == nil {
		reqoiu := a.buildRequest(put, url, headers, aJSON)
		oiusuc, stat := a.proxy.Do(reqoiu, &rtn)
		a.log.Debug("suc: ", oiusuc)
		a.log.Debug("stat: ", stat)
		if !oiusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetOrderItem GetOrderItem
func (a *Six910API) GetOrderItem(id int64, headers *Headers) *sdbi.OrderItem {
	var rtn sdbi.OrderItem
	var oisid = a.getStoreID(headers)
	idStrgoi := strconv.FormatInt(id, 10)
	sidStrgoi := strconv.FormatInt(oisid, 10)
	var url = a.restURL + "/rs/orderItem/get/id/" + idStrgoi + "/" + sidStrgoi
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	oigsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", oigsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetOrderItemList GetOrderItemList
func (a *Six910API) GetOrderItemList(orderID int64, headers *Headers) *[]sdbi.OrderItem {
	var rtn []sdbi.OrderItem
	var sid = a.getStoreID(headers)
	orderIDStrgoil := strconv.FormatInt(orderID, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/orderItem/get/list/" + orderIDStrgoil + "/" + sidStrGodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	oilsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", oilsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteOrderItem DeleteOrderItem
func (a *Six910API) DeleteOrderItem(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdoi := strconv.FormatInt(id, 10)
	sidStrdoi := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/orderItem/delete/" + idStrdoi + "/" + sidStrdoi
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	doisuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", doisuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
