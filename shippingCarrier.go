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

//AddShippingCarrier AddShippingCarrier
func (a *Six910API) AddShippingCarrier(c *sdbi.ShippingCarrier, headers *Headers) *ResponseID {
	var rtn ResponseID
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shippingCarrier/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqsca := a.buildRequest(post, url, headers, aJSON)
		sucsca, stat := a.proxy.Do(reqsca, &rtn)
		a.log.Debug("suc: ", sucsca)
		a.log.Debug("stat: ", stat)
		if !sucsca {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateShippingCarrier UpdateShippingCarrier
func (a *Six910API) UpdateShippingCarrier(c *sdbi.ShippingCarrier, headers *Headers) *Response {
	var rtn Response
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shippingCarrier/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqscu := a.buildRequest(put, url, headers, aJSON)
		scusuc, stat := a.proxy.Do(reqscu, &rtn)
		a.log.Debug("suc: ", scusuc)
		a.log.Debug("stat: ", stat)
		if !scusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetShippingCarrier GetShippingCarrier
func (a *Six910API) GetShippingCarrier(id int64, headers *Headers) *sdbi.ShippingCarrier {
	var rtn sdbi.ShippingCarrier
	var ctsid = a.getStoreID(headers)
	idStrGsc := strconv.FormatInt(id, 10)
	sidStrGsc := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/shippingCarrier/get/id/" + idStrGsc + "/" + sidStrGsc
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	scgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", scgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShippingCarrierList GetShippingCarrierList
func (a *Six910API) GetShippingCarrierList(headers *Headers) *[]sdbi.ShippingCarrier {
	var rtn []sdbi.ShippingCarrier
	var sid = a.getStoreID(headers)
	sidStrGscl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shippingCarrier/get/list/" + sidStrGscl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	sclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", sclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteShippingCarrier DeleteShippingCarrier
func (a *Six910API) DeleteShippingCarrier(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsc := strconv.FormatInt(id, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shippingCarrier/delete/" + idStrdsc + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dscsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dscsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
