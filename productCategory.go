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

//ProductCategoryReq ProductCategoryReq
type ProductCategoryReq struct {
	StoreID         int64                 `json:"storeId"`
	ProductCategory *sdbi.ProductCategory `json:"productCategory"`
}

//AddProductCategory AddProductCategory
func (a *Six910API) AddProductCategory(pc *sdbi.ProductCategory, headers *Headers) *Response {
	var rtn Response
	var pcr ProductCategoryReq
	pcr.ProductCategory = pc
	pcr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/productCategory/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(pcr)
	if err == nil {
		reqapc := a.buildRequest(post, url, headers, aJSON)
		sucapc, stat := a.proxy.Do(reqapc, &rtn)
		a.log.Debug("suc: ", sucapc)
		a.log.Debug("stat: ", stat)
		if !sucapc {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//DeleteProductCategory DeleteProductCategory
func (a *Six910API) DeleteProductCategory(pc *sdbi.ProductCategory, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	catidStrdct := strconv.FormatInt(pc.CategoryID, 10)
	pidStrdct := strconv.FormatInt(pc.ProductID, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/productCategory/delete/" + catidStrdct + "/" + pidStrdct + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dpcsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dpcsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
