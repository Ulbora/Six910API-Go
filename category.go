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

//AddCategory AddCategory
func (a *Six910API) AddCategory(c *sdbi.Category, headers *Headers) *ResponseID {
	var rtn ResponseID
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/category/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqct := a.buildRequest(post, url, headers, aJSON)
		succt, stat := a.proxy.Do(reqct, &rtn)
		a.log.Debug("suc: ", succt)
		a.log.Debug("stat: ", stat)
		if !succt {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//UpdateCategory UpdateCategory
func (a *Six910API) UpdateCategory(c *sdbi.Category, headers *Headers) *Response {
	var rtn Response
	c.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/category/update"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(c)
	if err == nil {
		reqctu := a.buildRequest(put, url, headers, aJSON)
		ctusuc, stat := a.proxy.Do(reqctu, &rtn)
		a.log.Debug("suc: ", ctusuc)
		a.log.Debug("stat: ", stat)
		if !ctusuc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetCategory GetCategory
func (a *Six910API) GetCategory(id int64, headers *Headers) *sdbi.Category {
	var rtn sdbi.Category
	var ctsid = a.getStoreID(headers)
	idStrGct := strconv.FormatInt(id, 10)
	sidStrGct := strconv.FormatInt(ctsid, 10)
	var url = a.restURL + "/rs/category/get/id/" + idStrGct + "/" + sidStrGct
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	cgsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", cgsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetHierarchicalCategoryList GetHierarchicalCategoryList
func (a *Six910API) GetHierarchicalCategoryList(headers *Headers) *[]sdbi.Category {
	var rtn []sdbi.Category
	var sid = a.getStoreID(headers)
	hsidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/category/get/list/hierarchical/" + hsidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	hausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", hausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetCategoryList GetCategoryList
func (a *Six910API) GetCategoryList(headers *Headers) *[]sdbi.Category {
	var rtn []sdbi.Category
	var sid = a.getStoreID(headers)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/category/get/list/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetSubCategoryList GetSubCategoryList
func (a *Six910API) GetSubCategoryList(catID int64, headers *Headers) *[]sdbi.Category {
	var rtn []sdbi.Category
	//var sid = a.getStoreID(headers)
	ctidStrGctl := strconv.FormatInt(catID, 10)

	var url = a.restURL + "/rs/category/get/sub/list/" + ctidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ausuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ausuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteCategory DeleteCategory
func (a *Six910API) DeleteCategory(id int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	idStrdct := strconv.FormatInt(id, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/category/delete/" + idStrdct + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dctsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dctsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
