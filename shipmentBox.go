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

//ShipmentBoxReq ShipmentBoxReq
type ShipmentBoxReq struct {
	StoreID     int64             `json:"storeId"`
	ShipmentBox *sdbi.ShipmentBox `json:"shipmentBox"`
}

//AddShipmentBox AddShipmentBox
func (a *Six910API) AddShipmentBox(sb *sdbi.ShipmentBox, headers *Headers) *ResponseID {
	var rtn ResponseID
	var sbr ShipmentBoxReq
	sbr.ShipmentBox = sb
	sbr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipmentBox/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sbr)
	if err == nil {
		reqasb := a.buildRequest(post, url, headers, aJSON)
		sucasb, stat := a.proxy.Do(reqasb, &rtn)
		a.log.Debug("suc: ", sucasb)
		a.log.Debug("stat: ", stat)
		if !sucasb {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateShipmentBox UpdateShipmentBox
func (a *Six910API) UpdateShipmentBox(sb *sdbi.ShipmentBox, headers *Headers) *Response {
	var rtn Response
	var sbr ShipmentBoxReq
	sbr.ShipmentBox = sb
	sbr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipmentBox/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sbr)
	if err == nil {
		reqsbu := a.buildRequest(put, url, headers, aJSON)
		sbusuc, stat := a.proxy.Do(reqsbu, &rtn)
		a.log.Debug("suc: ", sbusuc)
		a.log.Debug("stat: ", stat)
		if !sbusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetShipmentBox GetShipmentBox
func (a *Six910API) GetShipmentBox(id int64, headers *Headers) *sdbi.ShipmentBox {
	var rtn sdbi.ShipmentBox
	var ctsid = a.getStoreID(headers)
	idStrGsb := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/shipmentBox/get/id/" + idStrGsb + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gsbsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gsbsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShipmentBoxList GetShipmentBoxList
func (a *Six910API) GetShipmentBoxList(shipmentID int64, headers *Headers) *[]sdbi.ShipmentBox {
	var rtn []sdbi.ShipmentBox
	var sid = a.getStoreID(headers)
	spidStrGsb := strconv.FormatInt(shipmentID, 10)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipmentBox/get/list/" + spidStrGsb + "/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	sblsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", sblsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteShipmentBox DeleteShipmentBox
func (a *Six910API) DeleteShipmentBox(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsb := strconv.FormatInt(id, 10)
	sidStrdsb := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipmentBox/delete/" + idStrdsb + "/" + sidStrdsb
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dsbsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dsbsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
