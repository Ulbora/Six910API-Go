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

//AddPlugin AddPlugin
func (a *Six910API) AddPlugin(p *sdbi.Plugins, headers *Headers) *ResponseID {
	var rtn ResponseID
	var url = a.restURL + "/rs/plugin/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqapi := a.buildRequest(post, url, headers, aJSON)
		sucpia, stat := a.proxy.Do(reqapi, &rtn)
		a.log.Debug("suc: ", sucpia)
		a.log.Debug("stat: ", stat)
		if !sucpia {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdatePlugin UpdatePlugin
func (a *Six910API) UpdatePlugin(p *sdbi.Plugins, headers *Headers) *Response {
	var rtn Response
	var url = a.restURL + "/rs/plugin/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqpiu := a.buildRequest(put, url, headers, aJSON)
		piusuc, stat := a.proxy.Do(reqpiu, &rtn)
		a.log.Debug("suc: ", piusuc)
		a.log.Debug("stat: ", stat)
		if !piusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetPlugin GetPlugin
func (a *Six910API) GetPlugin(id int64, headers *Headers) *sdbi.Plugins {
	var rtn sdbi.Plugins
	idStrgpi := strconv.FormatInt(id, 10)
	var url = a.restURL + "/rs/plugin/get/id/" + idStrgpi
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pigsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pigsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetPluginList GetPluginList
func (a *Six910API) GetPluginList(start int64, end int64, headers *Headers) *[]sdbi.Plugins {
	var rtn []sdbi.Plugins
	startStr := strconv.FormatInt(start, 10)
	endStr := strconv.FormatInt(end, 10)

	var url = a.restURL + "/rs/plugin/get/list/" + startStr + "/" + endStr
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pilsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pilsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeletePlugin DeletePlugin
func (a *Six910API) DeletePlugin(id int64, headers *Headers) *Response {
	var rtn Response
	idStrdpi := strconv.FormatInt(id, 10)

	var url = a.restURL + "/rs/plugin/delete/" + idStrdpi
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dodsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dodsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
