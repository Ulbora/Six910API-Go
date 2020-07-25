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

//AddStorePlugin AddStorePlugin
func (a *Six910API) AddStorePlugin(sp *sdbi.StorePlugins, headers *Headers) *ResponseID {
	var rtn ResponseID
	sp.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/storePlugin/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sp)
	if err == nil {
		reqspi := a.buildRequest(post, url, headers, aJSON)
		sucspi, stat := a.proxy.Do(reqspi, &rtn)
		a.log.Debug("suc: ", sucspi)
		a.log.Debug("stat: ", stat)
		if !sucspi {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateStorePlugin UpdateStorePlugin
func (a *Six910API) UpdateStorePlugin(sp *sdbi.StorePlugins, headers *Headers) *Response {
	var rtn Response
	sp.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/storePlugin/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(sp)
	if err == nil {
		reqspu := a.buildRequest(put, url, headers, aJSON)
		spusuc, stat := a.proxy.Do(reqspu, &rtn)
		a.log.Debug("suc: ", spusuc)
		a.log.Debug("stat: ", stat)
		if !spusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetStorePlugin GetStorePlugin
func (a *Six910API) GetStorePlugin(id int64, headers *Headers) *sdbi.StorePlugins {
	var rtn sdbi.StorePlugins
	var ctsid = a.getStoreID(headers)
	idStrGspi := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/storePlugin/get/id/" + idStrGspi + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	spigsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", spigsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetStorePluginList GetStorePluginList
func (a *Six910API) GetStorePluginList(headers *Headers) *[]sdbi.StorePlugins {
	var rtn []sdbi.StorePlugins
	var sid = a.getStoreID(headers)
	sidStrGspil := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/storePlugin/get/list/" + sidStrGspil
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	spiglsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", spiglsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteStorePlugin DeleteStorePlugin
func (a *Six910API) DeleteStorePlugin(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdspi := strconv.FormatInt(id, 10)
	sidStrdspi := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/storePlugin/delete/" + idStrdspi + "/" + sidStrdspi
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dspisuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dspisuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
