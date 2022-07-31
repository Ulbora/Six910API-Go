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

//GetProductManufacturerListByProductName GetProductManufacturerListByProductName
func (a *Six910API) GetProductManufacturerListByProductName(name string, headers *Headers) *[]string {
	var rtn []string
	var sid = a.getStoreID(headers)
	sidStrGml := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/manufacturer/get/product/name/" + name + "/" + sidStrGml
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	mlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", mlsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

//GetProductManufacturerListByProductSearch GetProductManufacturerListByProductSearch
func (a *Six910API) GetProductManufacturerListByProductSearch(search string, headers *Headers) *[]string {
	var rtn []string
	var sid = a.getStoreID(headers)
	sidDesStrGml := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/manufacturer/get/product/desc/" + search + "/" + sidDesStrGml
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	dmlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dmlsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

//GetProductByNameAndManufacturerName GetProductByNameAndManufacturerName
func (a *Six910API) GetProductByNameAndManufacturerName(manf string, name string,
	start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var sid = a.getStoreID(headers)
	stStrgpml := strconv.FormatInt(start, 10)
	endStrgpml := strconv.FormatInt(end, 10)
	sidStrGpml := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/manufacturer/name/" + manf + "/" + name + "/" + sidStrGpml + "/" + stStrgpml + "/" + endStrgpml
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdmlsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdmlsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetProductManufacturerListByCatID GetProductManufacturerListByCatID
func (a *Six910API) GetProductManufacturerListByCatID(catID int64, headers *Headers) *[]string {
	var rtn []string
	var sid = a.getStoreID(headers)
	catStrGmcl := strconv.FormatInt(catID, 10)
	sidStrGmcl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/manufacturer/get/category/" + catStrGmcl + "/" + sidStrGmcl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	mclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", mclsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}

//GetProductByCatAndManufacturer GetProductByCatAndManufacturer
func (a *Six910API) GetProductByCatAndManufacturer(catID int64, manf string,
	start int64, end int64, headers *Headers) *[]sdbi.Product {
	var rtn []sdbi.Product
	var sid = a.getStoreID(headers)
	catStrgpmcl := strconv.FormatInt(catID, 10)
	stStrgpmcl := strconv.FormatInt(start, 10)
	endStrgpmcl := strconv.FormatInt(end, 10)
	sidStrGpmcl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/product/get/category/manufacturer/" + catStrgpmcl + "/" + manf + "/" + sidStrGpmcl + "/" + stStrgpmcl + "/" + endStrgpmcl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	pdmclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", pdmclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//ProductSearch ProductSearch
func (a *Six910API) ProductSearch(p *sdbi.ProductSearch, headers *Headers) *[]sdbi.Product {
	p.StoreID = a.getStoreID(headers)
	var prrtn []sdbi.Product
	var url = a.restURL + "/rs/product/search"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(p)
	if err == nil {
		reqpsr := a.buildRequest(post, url, headers, aJSON)
		sucpsr, stat := a.proxy.Do(reqpsr, &prrtn)
		a.log.Debug("suc: ", sucpsr)
		a.log.Debug("stat: ", stat)
	}
	return &prrtn
}
