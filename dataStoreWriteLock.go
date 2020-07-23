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

//AddDataStoreWriteLock AddDataStoreWriteLock
func (a *Six910API) AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *Headers) *Response {
	var rtn Response
	w.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/dataStoreWriteLock/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(w)
	if err == nil {
		reqdwa := a.buildRequest(post, url, headers, aJSON)
		sucdwa, stat := a.proxy.Do(reqdwa, &rtn)
		a.log.Debug("suc: ", sucdwa)
		a.log.Debug("stat: ", stat)
		if !sucdwa {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (a *Six910API) UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *Headers) *Response {
	var rtn Response
	w.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/dataStoreWriteLock/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(w)
	if err == nil {
		reqdwu := a.buildRequest(put, url, headers, aJSON)
		dwusuc, stat := a.proxy.Do(reqdwu, &rtn)
		a.log.Debug("suc: ", dwusuc)
		a.log.Debug("stat: ", stat)
		if !dwusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetDataStoreWriteLock GetDataStoreWriteLock
func (a *Six910API) GetDataStoreWriteLock(dataStore string, headers *Headers) *sdbi.DataStoreWriteLock {
	var rtn sdbi.DataStoreWriteLock
	var gdwsid = a.getStoreID(headers)
	sidStrGct := strconv.FormatInt(gdwsid, 10)
	var url = a.restURL + "/rs/dataStoreWriteLock/get/" + dataStore + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gdwsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gdwsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
