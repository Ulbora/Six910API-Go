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

//AddInstance AddInstance
func (a *Six910API) AddInstance(i *sdbi.Instances, headers *Headers) *Response {
	var rtn Response
	i.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/instance/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(i)
	if err == nil {
		reqaint := a.buildRequest(post, url, headers, aJSON)
		sucaint, stat := a.proxy.Do(reqaint, &rtn)
		a.log.Debug("suc: ", sucaint)
		a.log.Debug("stat: ", stat)
		if !sucaint {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateInstance UpdateInstance
func (a *Six910API) UpdateInstance(i *sdbi.Instances, headers *Headers) *Response {
	var rtn Response
	i.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/instance/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(i)
	if err == nil {
		reqintu := a.buildRequest(put, url, headers, aJSON)
		intusuc, stat := a.proxy.Do(reqintu, &rtn)
		a.log.Debug("suc: ", intusuc)
		a.log.Debug("stat: ", stat)
		if !intusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetInstance GetInstance
func (a *Six910API) GetInstance(name string, dataStoreName string, headers *Headers) *sdbi.Instances {
	var rtn sdbi.Instances
	var ctsid = a.getStoreID(headers)
	sidStrgint := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/instance/get/name/" + name + "/" + dataStoreName + "/" + sidStrgint
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	intgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", intgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetInstanceList GetInstanceList
func (a *Six910API) GetInstanceList(dataStoreName string, headers *Headers) *[]sdbi.Instances {
	var rtn []sdbi.Instances
	var sid = a.getStoreID(headers)
	sidStrgintl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/instance/get/list/" + dataStoreName + "/" + sidStrgintl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	intlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", intlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
