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

//AddCustomer AddCustomer
func (a *Six910API) AddCustomer(c *sdbi.Customer, headers *Headers) *ResponseID {
	var rtn ResponseID
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/customer/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqcus := a.buildRequest(post, url, headers, aJSON)
		succus, stat := a.proxy.Do(reqcus, &rtn)
		a.log.Debug("suc: ", succus)
		a.log.Debug("stat: ", stat)
		if !succus {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateCustomer UpdateCustomer
func (a *Six910API) UpdateCustomer(c *sdbi.Customer, headers *Headers) *Response {
	var rtn Response
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/customer/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqcusu := a.buildRequest(put, url, headers, aJSON)
		cususuc, stat := a.proxy.Do(reqcusu, &rtn)
		a.log.Debug("suc: ", cususuc)
		a.log.Debug("stat: ", stat)
		if !cususuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetCustomer GetCustomer
func (a *Six910API) GetCustomer(email string, headers *Headers) *sdbi.Customer {
	var rtn sdbi.Customer
	var ctsid = a.getStoreID(headers)
	sidStrGcus := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/customer/get/email/" + email + "/" + sidStrGcus
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetCustomerID GetCustomerID
func (a *Six910API) GetCustomerID(id int64, headers *Headers) *sdbi.Customer {
	var rtn sdbi.Customer
	var ctsid = a.getStoreID(headers)
	idStrGcus := strconv.FormatInt(id, 10)
	sidStrGcus := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/customer/get/id/" + idStrGcus + "/" + sidStrGcus
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cusgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cusgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetCustomerList GetCustomerList
func (a *Six910API) GetCustomerList(start int64, end int64, headers *Headers) *[]sdbi.Customer {
	var rtn []sdbi.Customer
	var sid = a.getStoreID(headers)
	sidStrGcusl := strconv.FormatInt(sid, 10)
	stStrGcusl := strconv.FormatInt(start, 10)
	endStrGcusl := strconv.FormatInt(end, 10)

	var url = a.restURL + "/rs/customer/get/list/" + sidStrGcusl + "/" + stStrGcusl + "/" + endStrGcusl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteCustomer DeleteCustomer
func (a *Six910API) DeleteCustomer(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdcus := strconv.FormatInt(id, 10)
	sidStrdcus := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/customer/delete/" + idStrdcus + "/" + sidStrdcus
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dcussuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dcussuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
