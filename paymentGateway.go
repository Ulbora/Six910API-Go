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

//PaymentGatewayReq PaymentGatewayReq
type PaymentGatewayReq struct {
	StoreID        int64                `json:"storeId"`
	PaymentGateway *sdbi.PaymentGateway `json:"paymentGateway"`
}

//AddPaymentGateway AddPaymentGateway
func (a *Six910API) AddPaymentGateway(pgw *sdbi.PaymentGateway, headers *Headers) *ResponseID {
	var rtn ResponseID
	var pgr PaymentGatewayReq
	pgr.PaymentGateway = pgw
	pgr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/paymentGateway/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(pgr)
	if err == nil {
		reqapg := a.buildRequest(post, url, headers, aJSON)
		sucapg, stat := a.proxy.Do(reqapg, &rtn)
		a.log.Debug("suc: ", sucapg)
		a.log.Debug("stat: ", stat)
		if !sucapg {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdatePaymentGateway UpdatePaymentGateway
func (a *Six910API) UpdatePaymentGateway(pgw *sdbi.PaymentGateway, headers *Headers) *Response {
	var rtn Response
	var pgr PaymentGatewayReq
	pgr.PaymentGateway = pgw
	pgr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/paymentGateway/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(pgr)
	if err == nil {
		reqodu := a.buildRequest(put, url, headers, aJSON)
		upgsuc, stat := a.proxy.Do(reqodu, &rtn)
		a.log.Debug("suc: ", upgsuc)
		a.log.Debug("stat: ", stat)
		if !upgsuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetPaymentGateway GetPaymentGateway
func (a *Six910API) GetPaymentGateway(id int64, headers *Headers) *sdbi.PaymentGateway {
	var rtn sdbi.PaymentGateway
	var ctsid = a.getStoreID(headers)
	idStrgpg := strconv.FormatInt(id, 10)
	sidStrgpg := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/paymentGateway/get/id/" + idStrgpg + "/" + sidStrgpg
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pggsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pggsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetPaymentGatewayByName GetPaymentGatewayByName
func (a *Six910API) GetPaymentGatewayByName(name string, headers *Headers) *sdbi.PaymentGateway {
	var rtn sdbi.PaymentGateway
	var ctsn = a.getStoreID(headers)
	sidStrgpgn := strconv.FormatInt(ctsn, 10)
	var url = a.restURL + "/rs/paymentGateway/get/name/" + name + "/" + sidStrgpgn
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pggnsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc gbn: ", pggnsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetPaymentGateways GetPaymentGateways
func (a *Six910API) GetPaymentGateways(headers *Headers) *[]sdbi.PaymentGateway {
	var rtn []sdbi.PaymentGateway
	var sid = a.getStoreID(headers)
	sidStrGpgl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/paymentGateway/get/list/" + sidStrGpgl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pglsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pglsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeletePaymentGateway DeletePaymentGateway
func (a *Six910API) DeletePaymentGateway(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdpg := strconv.FormatInt(id, 10)
	sidStrdpg := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/paymentGateway/delete/" + idStrdpg + "/" + sidStrdpg
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dpgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dpgsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
