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

//OrderTransactionReq OrderTransactionReq
type OrderTransactionReq struct {
	StoreID          int64                  `json:"storeId"`
	OrderTransaction *sdbi.OrderTransaction `json:"orderTransaction"`
}

//AddOrderTransaction AddOrderTransaction
func (a *Six910API) AddOrderTransaction(t *sdbi.OrderTransaction, headers *Headers) *ResponseID {
	var rtn ResponseID
	var ot OrderTransactionReq
	ot.OrderTransaction = t
	ot.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/orderTransaction/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(ot)
	if err == nil {
		reqota := a.buildRequest(post, url, headers, aJSON)
		sucota, stat := a.proxy.Do(reqota, &rtn)
		a.log.Debug("suc: ", sucota)
		a.log.Debug("stat: ", stat)
		if !sucota {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetOrderTransactionList GetOrderTransactionList
func (a *Six910API) GetOrderTransactionList(orderID int64, headers *Headers) *[]sdbi.OrderTransaction {
	var rtn []sdbi.OrderTransaction
	var sid = a.getStoreID(headers)
	orderIDStrgodl := strconv.FormatInt(orderID, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/orderTransaction/get/list/" + orderIDStrgodl + "/" + sidStrGodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	otlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", otlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
