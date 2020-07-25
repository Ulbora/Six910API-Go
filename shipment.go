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

//ShipmentReq ShipmentReq
type ShipmentReq struct {
	StoreID  int64          `json:"storeId"`
	Shipment *sdbi.Shipment `json:"shipment"`
}

//AddShipment AddShipment
func (a *Six910API) AddShipment(s *sdbi.Shipment, headers *Headers) *ResponseID {
	var rtn ResponseID
	var sr ShipmentReq
	sr.Shipment = s
	sr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipment/add"
	a.log.Debug("url: ", url)
	a.log.Debug("sr: ", *sr.Shipment)
	aJSON, err := json.Marshal(sr)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		sucasp, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", sucasp)
		a.log.Debug("stat: ", stat)
		if !sucasp {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateShipment UpdateShipment
func (a *Six910API) UpdateShipment(s *sdbi.Shipment, headers *Headers) *Response {
	var rtn Response
	var sr ShipmentReq
	sr.Shipment = s
	sr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipment/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sr)
	if err == nil {
		reqspu := a.buildRequest(put, url, headers, aJSON)
		spusuc, stat := a.proxy.Do(reqspu, &rtn)
		a.log.Debug("suc: ", spusuc)
		a.log.Debug("stat: ", stat)
		if !spusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetShipment GetShipment
func (a *Six910API) GetShipment(id int64, headers *Headers) *sdbi.Shipment {
	var rtn sdbi.Shipment
	var gspsid = a.getStoreID(headers)
	idStrGsp := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(gspsid, 10)
	var url = a.restURL + "/rs/shipment/get/id/" + idStrGsp + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	spgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", spgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShipmentList GetShipmentList
func (a *Six910API) GetShipmentList(orderID int64, headers *Headers) *[]sdbi.Shipment {
	var rtn []sdbi.Shipment
	var sid = a.getStoreID(headers)
	oidStrGsp := strconv.FormatInt(orderID, 10)
	sidStrGspl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipment/get/list/" + oidStrGsp + "/" + sidStrGspl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gsplsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gsplsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteShipment DeleteShipment
func (a *Six910API) DeleteShipment(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsp := strconv.FormatInt(id, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipment/delete/" + idStrdsp + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dspsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dspsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
