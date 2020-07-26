package six910api

import (
	"encoding/json"
	"strconv"
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

//AddCustomerUser AddCustomerUser
func (a *Six910API) AddCustomerUser(u *User, headers *Headers) *Response {
	var rtn Response
	u.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/user/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(u)
	if err == nil {
		reqacu := a.buildRequest(post, url, headers, aJSON)
		sucacu, stat := a.proxy.Do(reqacu, &rtn)
		a.log.Debug("suc: ", sucacu)
		a.log.Debug("stat: ", stat)
		if !sucacu {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateUser UpdateUser
func (a *Six910API) UpdateUser(u *User, headers *Headers) *Response {
	var rtn Response
	u.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/user/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(u)
	if err == nil {
		reqcuu := a.buildRequest(put, url, headers, aJSON)
		cuusuc, stat := a.proxy.Do(reqcuu, &rtn)
		a.log.Debug("suc: ", cuusuc)
		a.log.Debug("stat: ", stat)
		if !cuusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetUser GetUser
func (a *Six910API) GetUser(u *User, headers *Headers) *UserResponse {
	var rtn UserResponse
	var ussid = a.getStoreID(headers)

	sidStrGus := strconv.FormatInt(ussid, 10)
	var url = a.restURL + "/rs/user/" + u.Username + "/" + sidStrGus
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	usgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", usgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetAdminUsers GetAdminUsers
func (a *Six910API) GetAdminUsers(headers *Headers) *[]UserResponse {
	var rtn []UserResponse
	var sid = a.getStoreID(headers)
	sidStrGuaul := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/user/get/admin/list/" + sidStrGuaul
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gaulsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gaulsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetCustomerUsers GetCustomerUsers
func (a *Six910API) GetCustomerUsers(headers *Headers) *[]UserResponse {
	var rtn []UserResponse
	var sid = a.getStoreID(headers)
	sidStrGucul := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/user/get/customer/list/" + sidStrGucul
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gculsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gculsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
