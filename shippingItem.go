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

//ShipmentItemReq ShipmentItemReq
type ShipmentItemReq struct {
	StoreID      int64              `json:"storeId"`
	ShipmentItem *sdbi.ShipmentItem `json:"shipmentItem"`
}

//AddShipmentItem AddShipmentItem
func (a *Six910API) AddShipmentItem(si *sdbi.ShipmentItem, headers *Headers) *ResponseID {
	var rtn ResponseID
	var sir ShipmentItemReq
	sir.ShipmentItem = si
	sir.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipmentItem/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sir)
	if err == nil {
		reqasi := a.buildRequest(post, url, headers, aJSON)
		sucasi, stat := a.proxy.Do(reqasi, &rtn)
		a.log.Debug("suc: ", sucasi)
		a.log.Debug("stat: ", stat)
		if !sucasi {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateShipmentItem UpdateShipmentItem
func (a *Six910API) UpdateShipmentItem(si *sdbi.ShipmentItem, headers *Headers) *Response {
	var rtn Response
	var sir ShipmentItemReq
	sir.ShipmentItem = si
	sir.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/shipmentItem/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sir)
	if err == nil {
		reqsiu := a.buildRequest(put, url, headers, aJSON)
		siusuc, stat := a.proxy.Do(reqsiu, &rtn)
		a.log.Debug("suc: ", siusuc)
		a.log.Debug("stat: ", stat)
		if !siusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetShipmentItem GetShipmentItem
func (a *Six910API) GetShipmentItem(id int64, headers *Headers) *sdbi.ShipmentItem {
	var rtn sdbi.ShipmentItem
	var ctsid = a.getStoreID(headers)
	idStrGsi := strconv.FormatInt(id, 10)
	sidStrGsi := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/shipmentItem/get/id/" + idStrGsi + "/" + sidStrGsi
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShipmentItemList GetShipmentItemList
func (a *Six910API) GetShipmentItemList(shipmentID int64, headers *Headers) *[]sdbi.ShipmentItem {
	var rtn []sdbi.ShipmentItem
	var sid = a.getStoreID(headers)
	spidStrGsi := strconv.FormatInt(shipmentID, 10)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipmentItem/get/list/" + spidStrGsi + "/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	silsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", silsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (a *Six910API) GetShipmentItemListByBox(boxNumber int64, shipmentID int64, headers *Headers) *[]sdbi.ShipmentItem {
	var rtn []sdbi.ShipmentItem
	var sid = a.getStoreID(headers)
	bnumStrGsi := strconv.FormatInt(boxNumber, 10)
	spidStrGsib := strconv.FormatInt(shipmentID, 10)
	sidStrGctlb := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipmentItem/get/list/box/" + bnumStrGsi + "/" + spidStrGsib + "/" + sidStrGctlb
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	silbsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", silbsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteShipmentItem DeleteShipmentItem
func (a *Six910API) DeleteShipmentItem(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsi := strconv.FormatInt(id, 10)
	sidStrdsi := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/shipmentItem/delete/" + idStrdsi + "/" + sidStrdsi
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dsisuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dsisuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
