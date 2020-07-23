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

//AddInsurance AddInsurance
func (a *Six910API) AddInsurance(i *sdbi.Insurance, headers *Headers) *ResponseID {
	var rtn ResponseID
	i.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/insurance/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(i)
	if err == nil {
		reqinsu := a.buildRequest(post, url, headers, aJSON)
		sucinsu, stat := a.proxy.Do(reqinsu, &rtn)
		a.log.Debug("suc: ", sucinsu)
		a.log.Debug("stat: ", stat)
		if !sucinsu {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateInsurance UpdateInsurance
func (a *Six910API) UpdateInsurance(i *sdbi.Insurance, headers *Headers) *Response {
	var rtn Response
	i.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/insurance/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(i)
	if err == nil {
		reqinsuu := a.buildRequest(put, url, headers, aJSON)
		insuusuc, stat := a.proxy.Do(reqinsuu, &rtn)
		a.log.Debug("suc: ", insuusuc)
		a.log.Debug("stat: ", stat)
		if !insuusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetInsurance GetInsurance
func (a *Six910API) GetInsurance(id int64, headers *Headers) *sdbi.Insurance {
	var rtn sdbi.Insurance
	var ctsid = a.getStoreID(headers)
	idStrginsu := strconv.FormatInt(id, 10)
	sidStrginsu := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/insurance/get/id/" + idStrginsu + "/" + sidStrginsu
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ginsusuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ginsusuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetInsuranceList GetInsuranceList
func (a *Six910API) GetInsuranceList(headers *Headers) *[]sdbi.Insurance {
	var rtn []sdbi.Insurance
	var sid = a.getStoreID(headers)
	sidStrginsul := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/insurance/get/list/" + sidStrginsul
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	insulsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", insulsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteInsurance DeleteInsurance
func (a *Six910API) DeleteInsurance(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdinsu := strconv.FormatInt(id, 10)
	sidStrdinsu := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/insurance/delete/" + idStrdinsu + "/" + sidStrdinsu
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dinsusuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dinsusuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
