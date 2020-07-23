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

//AddDistributor AddDistributor
func (a *Six910API) AddDistributor(d *sdbi.Distributor, headers *Headers) *ResponseID {
	var rtn ResponseID
	d.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/distributor/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(d)
	if err == nil {
		reqdisa := a.buildRequest(post, url, headers, aJSON)
		sucdisa, stat := a.proxy.Do(reqdisa, &rtn)
		a.log.Debug("suc: ", sucdisa)
		a.log.Debug("stat: ", stat)
		if !sucdisa {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateDistributor UpdateDistributor
func (a *Six910API) UpdateDistributor(d *sdbi.Distributor, headers *Headers) *Response {
	var rtn Response
	d.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/distributor/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(d)
	if err == nil {
		reqdisu := a.buildRequest(put, url, headers, aJSON)
		disusuc, stat := a.proxy.Do(reqdisu, &rtn)
		a.log.Debug("suc: ", disusuc)
		a.log.Debug("stat: ", stat)
		if !disusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetDistributor GetDistributor
func (a *Six910API) GetDistributor(id int64, headers *Headers) *sdbi.Distributor {
	var rtn sdbi.Distributor
	var gdissid = a.getStoreID(headers)
	idStrgdis := strconv.FormatInt(id, 10)
	sidStrgdis := strconv.FormatInt(gdissid, 10)
	var url = a.restURL + "/rs/distributor/get/id/" + idStrgdis + "/" + sidStrgdis
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gdissuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gdissuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetDistributorList GetDistributorList
func (a *Six910API) GetDistributorList(headers *Headers) *[]sdbi.Distributor {
	var rtn []sdbi.Distributor
	var sid = a.getStoreID(headers)
	sidStrGdis := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/distributor/get/list/" + sidStrGdis
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gdissuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gdissuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteDistributor DeleteDistributor
func (a *Six910API) DeleteDistributor(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrddis := strconv.FormatInt(id, 10)
	sidStrddis := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/distributor/delete/" + idStrddis + "/" + sidStrddis
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	ddissuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ddissuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
