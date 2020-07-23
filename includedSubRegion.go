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

//IncludedSubRegionReq IncludedSubRegionReq
type IncludedSubRegionReq struct {
	StoreID           int64                   `json:"storeId"`
	IncludedSubRegion *sdbi.IncludedSubRegion `json:"includedSubRegion"`
}

//AddIncludedSubRegion AddIncludedSubRegion
func (a *Six910API) AddIncludedSubRegion(e *sdbi.IncludedSubRegion, headers *Headers) *ResponseID {
	var rtn ResponseID
	var inreq IncludedSubRegionReq
	inreq.IncludedSubRegion = e
	inreq.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/includedSubRegion/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(inreq)
	if err == nil {
		reqin := a.buildRequest(post, url, headers, aJSON)
		sucin, stat := a.proxy.Do(reqin, &rtn)
		a.log.Debug("suc: ", sucin)
		a.log.Debug("stat: ", stat)
		if !sucin {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (a *Six910API) GetIncludedSubRegionList(regionID int64, headers *Headers) *[]sdbi.IncludedSubRegion {
	var rtn []sdbi.IncludedSubRegion
	var sid = a.getStoreID(headers)
	ridStrginr := strconv.FormatInt(regionID, 10)
	sidStrinr := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/includedSubRegion/get/list/" + ridStrginr + "/" + sidStrinr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	inrsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", inrsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (a *Six910API) DeleteIncludedSubRegion(id int64, regionID int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdinr := strconv.FormatInt(id, 10)
	ridStrdinr := strconv.FormatInt(regionID, 10)
	sidStrdinr := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/includedSubRegion/delete/" + idStrdinr + "/" + ridStrdinr + "/" + sidStrdinr
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dinrsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dinrsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
