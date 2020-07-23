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

//ExcludedSubRegionReq ExcludedSubRegionReq
type ExcludedSubRegionReq struct {
	StoreID           int64                   `json:"storeId"`
	ExcludedSubRegion *sdbi.ExcludedSubRegion `json:"excludedSubRegion"`
}

//AddExcludedSubRegion AddExcludedSubRegion
func (a *Six910API) AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, headers *Headers) *ResponseID {
	var rtn ResponseID
	var exreq ExcludedSubRegionReq
	exreq.ExcludedSubRegion = e
	exreq.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/excludedSubRegion/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(exreq)
	if err == nil {
		reqex := a.buildRequest(post, url, headers, aJSON)
		succt, stat := a.proxy.Do(reqex, &rtn)
		a.log.Debug("suc: ", succt)
		a.log.Debug("stat: ", stat)
		if !succt {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (a *Six910API) GetExcludedSubRegionList(regionID int64, headers *Headers) *[]sdbi.ExcludedSubRegion {
	var rtn []sdbi.ExcludedSubRegion
	var sid = a.getStoreID(headers)
	ridStrgex := strconv.FormatInt(regionID, 10)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/excludedSubRegion/get/list/" + ridStrgex + "/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (a *Six910API) DeleteExcludedSubRegion(id int64, regionID int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrder := strconv.FormatInt(id, 10)
	ridStrder := strconv.FormatInt(regionID, 10)
	sidStrder := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/excludedSubRegion/delete/" + idStrder + "/" + ridStrder + "/" + sidStrder
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dersuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dersuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
