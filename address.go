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

//AddressReq AddressReq
type AddressReq struct {
	StoreID int64         `json:"storeId"`
	Address *sdbi.Address `json:"address"`
}

//AddAddress AddAddress
func (a *Six910API) AddAddress(ad *sdbi.Address, headers *Headers) *ResponseID {
	var rtn ResponseID
	var add AddressReq
	add.Address = ad
	add.StoreID = a.getStoreID(headers)

	var url = a.restURL + "/rs/address/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(add)
	if err == nil {
		req := a.buildRequest(post, url, headers, aJSON)
		suc, stat := a.proxy.Do(req, &rtn)
		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		if !suc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateAddress UpdateAddress
func (a *Six910API) UpdateAddress(ad *sdbi.Address, headers *Headers) *Response {
	var rtn Response
	var addu AddressReq
	addu.Address = ad
	addu.StoreID = a.getStoreID(headers)

	var url = a.restURL + "/rs/address/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(addu)
	if err == nil {
		req := a.buildRequest(put, url, headers, aJSON)
		ausuc, stat := a.proxy.Do(req, &rtn)
		a.log.Debug("suc: ", ausuc)
		a.log.Debug("stat: ", stat)
		if !ausuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetAddress GetAddress
func (a *Six910API) GetAddress(id int64, cid int64, headers *Headers) *sdbi.Address {
	var rtn sdbi.Address
	var sid = a.getStoreID(headers)
	idStr := strconv.FormatInt(id, 10)
	cidStr := strconv.FormatInt(cid, 10)
	sidStr := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/address/get/id/" + idStr + "/" + cidStr + "/" + sidStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetAddressList GetAddressList
func (a *Six910API) GetAddressList(cid int64, headers *Headers) *[]sdbi.Address {
	var rtn []sdbi.Address
	var sid = a.getStoreID(headers)
	cidStr := strconv.FormatInt(cid, 10)
	sidStr := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/address/get/list/" + cidStr + "/" + sidStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteAddress DeleteAddress
func (a *Six910API) DeleteAddress(id int64, cid int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStr := strconv.FormatInt(id, 10)
	cidStr := strconv.FormatInt(cid, 10)
	sidStr := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/address/delete/" + idStr + "/" + cidStr + "/" + sidStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
