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

//AddRegion AddRegion
func (a *Six910API) AddRegion(r *sdbi.Region, headers *Headers) *ResponseID {
	var rtn ResponseID
	r.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/region/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(r)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		sucarg, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", sucarg)
		a.log.Debug("stat: ", stat)
		if !sucarg {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateRegion UpdateRegion
func (a *Six910API) UpdateRegion(r *sdbi.Region, headers *Headers) *Response {
	var rtn Response
	r.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/region/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(r)
	if err == nil {
		reqrgu := a.buildRequest(put, url, headers, aJSON)
		rgusuc, stat := a.proxy.Do(reqrgu, &rtn)
		a.log.Debug("suc: ", rgusuc)
		a.log.Debug("stat: ", stat)
		if !rgusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetRegion GetRegion
func (a *Six910API) GetRegion(id int64, headers *Headers) *sdbi.Region {
	var rtn sdbi.Region
	var ctsid = a.getStoreID(headers)
	idStrgrg := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/region/get/id/" + idStrgrg + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	rggsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", rggsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetRegionList GetRegionList
func (a *Six910API) GetRegionList(headers *Headers) *[]sdbi.Region {
	var rtn []sdbi.Region
	var sid = a.getStoreID(headers)
	sidStrgrgl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/region/get/list/" + sidStrgrgl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	grglsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", grglsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteRegion DeleteRegion
func (a *Six910API) DeleteRegion(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdgr := strconv.FormatInt(id, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/region/delete/" + idStrdgr + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dgrsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dgrsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
