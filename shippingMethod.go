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

//AddShippingMethod AddShippingMethod
func (a *Six910API) AddShippingMethod(s *sdbi.ShippingMethod, headers *Headers) *ResponseID {
	var rtn ResponseID
	s.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shippingMethod/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(s)
	if err == nil {
		reqasm := a.buildRequest(post, url, headers, aJSON)
		sucasm, stat := a.proxy.Do(reqasm, &rtn)
		a.log.Debug("suc: ", sucasm)
		a.log.Debug("stat: ", stat)
		if !sucasm {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateShippingMethod UpdateShippingMethod
func (a *Six910API) UpdateShippingMethod(s *sdbi.ShippingMethod, headers *Headers) *Response {
	var rtn Response
	s.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shippingMethod/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(s)
	if err == nil {
		reqsmu := a.buildRequest(put, url, headers, aJSON)
		smusuc, stat := a.proxy.Do(reqsmu, &rtn)
		a.log.Debug("suc: ", smusuc)
		a.log.Debug("stat: ", stat)
		if !smusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetShippingMethod GetShippingMethod
func (a *Six910API) GetShippingMethod(id int64, headers *Headers) *sdbi.ShippingMethod {
	var rtn sdbi.ShippingMethod
	var ctsid = a.getStoreID(headers)
	idStrGsm := strconv.FormatInt(id, 10)
	sidStrGsm := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/shippingMethod/get/id/" + idStrGsm + "/" + sidStrGsm
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	smgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", smgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShippingMethodList GetShippingMethodList
func (a *Six910API) GetShippingMethodList(headers *Headers) *[]sdbi.ShippingMethod {
	var rtn []sdbi.ShippingMethod
	var sid = a.getStoreID(headers)
	sidStrGsml := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shippingMethod/get/list/" + sidStrGsml
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	smlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", smlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteShippingMethod DeleteShippingMethod
func (a *Six910API) DeleteShippingMethod(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsm := strconv.FormatInt(id, 10)
	sidStrdsm := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shippingMethod/delete/" + idStrdsm + "/" + sidStrdsm
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dsmsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dsmsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
