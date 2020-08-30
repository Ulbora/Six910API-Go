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

//AddTaxRate AddTaxRate
func (a *Six910API) AddTaxRate(t *sdbi.TaxRate, headers *Headers) *ResponseID {
	var rtn ResponseID
	t.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/taxRate/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(t)
	if err == nil {
		reqtra := a.buildRequest(post, url, headers, aJSON)
		suctra, stat := a.proxy.Do(reqtra, &rtn)
		a.log.Debug("suc: ", suctra)
		a.log.Debug("stat: ", stat)
		if !suctra {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateTaxRate UpdateTaxRate
func (a *Six910API) UpdateTaxRate(t *sdbi.TaxRate, headers *Headers) *Response {
	var rtn Response
	t.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/taxRate/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(t)
	if err == nil {
		reqtru := a.buildRequest(put, url, headers, aJSON)
		trusuc, stat := a.proxy.Do(reqtru, &rtn)
		a.log.Debug("suc: ", trusuc)
		a.log.Debug("stat: ", stat)
		if !trusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetTaxRate GetTaxRate
func (a *Six910API) GetTaxRate(country string, state string, headers *Headers) *[]sdbi.TaxRate {
	var rtn []sdbi.TaxRate
	var sid = a.getStoreID(headers)
	strStrginsul := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/taxRate/get/country/" + country + "/" + state + "/" + strStrginsul
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	trsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", trsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetTaxRateList GetTaxRateList
func (a *Six910API) GetTaxRateList(headers *Headers) *[]sdbi.TaxRate {
	var rtn []sdbi.TaxRate
	var sid = a.getStoreID(headers)
	sidStrgtrl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/taxRate/get/list/" + sidStrgtrl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	trlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", trlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteTaxRate DeleteTaxRate
func (a *Six910API) DeleteTaxRate(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdtrd := strconv.FormatInt(id, 10)
	sidStrdtrd := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/taxRate/delete/" + idStrdtrd + "/" + sidStrdtrd
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dtrdsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dtrdsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
