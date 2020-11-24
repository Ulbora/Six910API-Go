package six910api

import (
	"encoding/json"
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

//AddAdminUser AddAdminUser
func (a *Six910API) AddAdminUser(au *User, headers *Headers) *Response {
	var rtn Response
	au.StoreID = a.getStoreID(headers)
	var aurl = a.restURL + "/rs/user/add"
	a.log.Debug("url: ", aurl)
	aJSON, err := json.Marshal(au)
	if err == nil {
		reqacu := a.buildRequest(post, aurl, headers, aJSON)
		sucaacu, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", sucaacu)
		a.log.Debug("stat: ", stat)
		if !sucaacu {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//AdminUpdateUser AdminUpdateUser
func (a *Six910API) AdminUpdateUser(au *User, headers *Headers) *Response {
	var rtn Response
	au.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/user/admin/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(au)
	if err == nil {
		reqacuu := a.buildRequest(put, url, headers, aJSON)
		acuusuc, stat := a.proxy.Do(reqacuu, &rtn)
		a.log.Debug("suc: ", acuusuc)
		a.log.Debug("stat: ", stat)
		if !acuusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}
