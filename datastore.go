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

//AddLocalDatastore AddLocalDatastore
func (a *Six910API) AddLocalDatastore(d *sdbi.LocalDataStore, headers *Headers) *Response {
	var rtn Response
	d.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/datastore/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(d)
	if err == nil {
		reqdsa := a.buildRequest(post, url, headers, aJSON)
		sucdsa, stat := a.proxy.Do(reqdsa, &rtn)
		a.log.Debug("suc: ", sucdsa)
		a.log.Debug("stat: ", stat)
		if !sucdsa {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateLocalDatastore UpdateLocalDatastore
func (a *Six910API) UpdateLocalDatastore(d *sdbi.LocalDataStore, headers *Headers) *Response {
	var rtn Response
	d.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/datastore/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(d)
	if err == nil {
		reqdsu := a.buildRequest(put, url, headers, aJSON)
		dsusuc, stat := a.proxy.Do(reqdsu, &rtn)
		a.log.Debug("suc: ", dsusuc)
		a.log.Debug("stat: ", stat)
		if !dsusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetLocalDatastore GetLocalDatastore
func (a *Six910API) GetLocalDatastore(dataStoreName string, headers *Headers) *sdbi.LocalDataStore {
	var rtn sdbi.LocalDataStore
	var gdssid = a.getStoreID(headers)
	sidStrGct := strconv.FormatInt(gdssid, 10)
	var url = a.restURL + "/rs/datastore/get/" + dataStoreName + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gdssuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gdssuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
