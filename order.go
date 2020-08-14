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

//AddOrder AddOrder
func (a *Six910API) AddOrder(o *sdbi.Order, headers *Headers) *ResponseID {
	var rtn ResponseID
	o.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/order/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(o)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		sucoda, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", sucoda)
		a.log.Debug("stat: ", stat)
		if !sucoda {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateOrder UpdateOrder
func (a *Six910API) UpdateOrder(o *sdbi.Order, headers *Headers) *Response {
	var rtn Response
	o.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/order/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(o)
	if err == nil {
		reqodu := a.buildRequest(put, url, headers, aJSON)
		odusuc, stat := a.proxy.Do(reqodu, &rtn)
		a.log.Debug("suc: ", odusuc)
		a.log.Debug("stat: ", stat)
		if !odusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetOrder GetOrder
func (a *Six910API) GetOrder(id int64, headers *Headers) *sdbi.Order {
	var rtn sdbi.Order
	var ctsid = a.getStoreID(headers)
	idStrgod := strconv.FormatInt(id, 10)
	sidStrgod := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/order/get/id/" + idStrgod + "/" + sidStrgod
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	odgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", odgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetOrderList GetOrderList
func (a *Six910API) GetOrderList(cid int64, headers *Headers) *[]sdbi.Order {
	var rtn []sdbi.Order
	var sid = a.getStoreID(headers)
	cidStrgodl := strconv.FormatInt(cid, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/order/get/list/" + cidStrgodl + "/" + sidStrGodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	odlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", odlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetStoreOrderList GetStoreOrderList
func (a *Six910API) GetStoreOrderList(headers *Headers) *[]sdbi.Order {
	var rtn []sdbi.Order
	var ssid = a.getStoreID(headers)
	sidStrGodl := strconv.FormatInt(ssid, 10)

	var url = a.restURL + "/rs/order/get/store/list/" + sidStrGodl
	a.log.Debug("url: ", url)

	sreq := a.buildRequest(get, url, headers, nil)
	sodlsuc, sstat := a.proxy.Do(sreq, &rtn)
	a.log.Debug("suc: ", sodlsuc)
	a.log.Debug("stat: ", sstat)

	return &rtn
}

//GetStoreOrderListByStatus GetStoreOrderListByStatus
func (a *Six910API) GetStoreOrderListByStatus(status string, headers *Headers) *[]sdbi.Order {
	var rtn []sdbi.Order
	var sid = a.getStoreID(headers)
	sssidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/order/get/store/list/status/" + status + "/" + sssidStrGodl
	a.log.Debug("url: ", url)

	ssreq := a.buildRequest(get, url, headers, nil)
	ssodlsuc, ssstat := a.proxy.Do(ssreq, &rtn)
	a.log.Debug("suc: ", ssodlsuc)
	a.log.Debug("stat: ", ssstat)

	return &rtn
}

//DeleteOrder DeleteOrder
func (a *Six910API) DeleteOrder(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdod := strconv.FormatInt(id, 10)
	sidStrdod := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/order/delete/" + idStrdod + "/" + sidStrdod
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dodsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dodsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
