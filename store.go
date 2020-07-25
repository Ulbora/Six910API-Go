package six910api

import (
	"encoding/json"

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

//AddStore AddStore
func (a *Six910API) AddStore(s *sdbi.Store, headers *Headers) *ResponseID {
	var rtn ResponseID
	var url = a.restURL + "/rs/store/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(s)
	if err == nil {
		reqastr := a.buildRequest(post, url, headers, aJSON)
		sucastr, stat := a.proxy.Do(reqastr, &rtn)
		a.log.Debug("suc: ", sucastr)
		a.log.Debug("stat: ", stat)
		if !sucastr {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateStore UpdateStore
func (a *Six910API) UpdateStore(s *sdbi.Store, headers *Headers) *Response {
	var rtn Response
	var url = a.restURL + "/rs/store/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(s)
	if err == nil {
		reqstru := a.buildRequest(put, url, headers, aJSON)
		strusuc, stat := a.proxy.Do(reqstru, &rtn)
		a.log.Debug("suc: ", strusuc)
		a.log.Debug("stat: ", stat)
		if !strusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetStore GetStore
func (a *Six910API) GetStore(sname string, localDomain string, headers *Headers) *sdbi.Store {
	var rtn sdbi.Store
	var url = a.restURL + "/rs/store/get/" + sname + "/" + localDomain
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	strgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", strgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteStore DeleteStore
func (a *Six910API) DeleteStore(sname string, localDomain string, headers *Headers) *Response {
	var rtn Response

	var url = a.restURL + "/rs/store/delete/" + sname + "/" + localDomain
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dstrsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dstrsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
