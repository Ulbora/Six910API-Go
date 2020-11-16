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

//AddVisit AddVisit
func (a *Six910API) AddVisit(v *sdbi.Visitor, headers *Headers) *Response {
	var rtn Response
	v.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/visit/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(v)
	if err == nil {
		reqva := a.buildRequest(post, url, headers, aJSON)
		sucvaa, stat := a.proxy.Do(reqva, &rtn)
		a.log.Debug("suc: ", sucvaa)
		a.log.Debug("stat: ", stat)
		if !sucvaa {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetVisitorData GetVisitorData
func (a *Six910API) GetVisitorData(headers *Headers) *[]sdbi.VisitorData {
	var rtn []sdbi.VisitorData
	var sid = a.getStoreID(headers)
	sidStrGvdl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/visitor/data/list/" + sidStrGvdl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	vdlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", vdlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
