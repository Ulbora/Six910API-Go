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

//SubRegionReq SubRegionReq
type SubRegionReq struct {
	StoreID   int64           `json:"storeId"`
	SubRegion *sdbi.SubRegion `json:"subRegion"`
}

//AddSubRegion AddSubRegion
func (a *Six910API) AddSubRegion(s *sdbi.SubRegion, headers *Headers) *ResponseID {
	var rtn ResponseID
	var srr SubRegionReq
	srr.SubRegion = s
	srr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/subRegion/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(srr)
	if err == nil {
		reqasr := a.buildRequest(post, url, headers, aJSON)
		sucasr, stat := a.proxy.Do(reqasr, &rtn)
		a.log.Debug("suc: ", sucasr)
		a.log.Debug("stat: ", stat)
		if !sucasr {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateSubRegion UpdateSubRegion
func (a *Six910API) UpdateSubRegion(s *sdbi.SubRegion, headers *Headers) *Response {
	var rtn Response
	var srr SubRegionReq
	srr.SubRegion = s
	srr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/subRegion/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(srr)
	if err == nil {
		reqsru := a.buildRequest(put, url, headers, aJSON)
		srusuc, stat := a.proxy.Do(reqsru, &rtn)
		a.log.Debug("suc: ", srusuc)
		a.log.Debug("stat: ", stat)
		if !srusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetSubRegion GetSubRegion
func (a *Six910API) GetSubRegion(id int64, headers *Headers) *sdbi.SubRegion {
	var rtn sdbi.SubRegion
	var ctsid = a.getStoreID(headers)
	idStrGsr := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/subRegion/get/id/" + idStrGsr + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	srgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", srgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetSubRegionList GetSubRegionList
func (a *Six910API) GetSubRegionList(regionID int64, headers *Headers) *[]sdbi.SubRegion {
	var rtn []sdbi.SubRegion
	var sid = a.getStoreID(headers)
	rgidStrGsrl := strconv.FormatInt(regionID, 10)
	sidStrGsrl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/subRegion/get/list/" + rgidStrGsrl + "/" + sidStrGsrl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	srlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", srlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteSubRegion DeleteSubRegion
func (a *Six910API) DeleteSubRegion(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdsr := strconv.FormatInt(id, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/subRegion/delete/" + idStrdsr + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dsrsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dsrsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
