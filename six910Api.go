package six910api

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

import (
	"bytes"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	sdbi "github.com/Ulbora/six910-database-interface"
)

//Six910API Six910API
type Six910API struct {
	proxy       px.Proxy
	log         *lg.Logger
	restURL     string
	storeName   string
	localDomain string
	storeID     int64
	apiKey      string
}

//GetNew GetNew
func (a *Six910API) GetNew() API {
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	a.log = &l

	var p px.GoProxy
	a.proxy = &p

	return a
}

func (a *Six910API) buildRequest(method string, url string, headers *Headers, aJSON []byte) *http.Request {
	headers.Set("storeName", a.storeName)
	headers.Set("localDomain", a.localDomain)
	if a.apiKey != "" {
		headers.Set("apiKey", a.apiKey)
	}
	var req *http.Request
	var err error
	if method == post || method == put {
		headers.Set("Content-Type", "application/json")
		req, err = http.NewRequest(method, url, bytes.NewBuffer(aJSON))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	a.log.Debug("err in build req: ", err)
	if err == nil {
		for k, v := range headers.headers {
			a.log.Debug("header: ", k, v)
			req.Header.Set(k, v)
		}
	}
	return req
}

func (a *Six910API) getStoreID(headers *Headers) int64 {
	if a.storeID == 0 {
		var url = a.restURL + "/rs/store/get/" + a.storeName + "/" + a.localDomain
		var str sdbi.Store
		req := a.buildRequest(get, url, headers, nil)
		suc, stat := a.proxy.Do(req, &str)
		a.log.Debug("suc: ", suc)
		a.log.Debug("stat: ", stat)
		if suc {
			a.storeID = str.ID
		}
	}
	a.log.Debug("storeId: ", a.storeID)
	return a.storeID
}

//SetLogLever SetLogLever
func (a *Six910API) SetLogLever(level int) {
	a.log.LogLevel = level
}

//SetStore SetStore
func (a *Six910API) SetStore(storeName string, localDomain string) {
	a.storeName = storeName
	a.localDomain = localDomain
}

//SetRestURL SetRestURL
func (a *Six910API) SetRestURL(url string) {
	a.restURL = url
}

//SetAPIKey SetAPIKey
func (a *Six910API) SetAPIKey(key string) {
	a.apiKey = key
}

//OverrideProxy OverrideProxy
func (a *Six910API) OverrideProxy(proxy px.Proxy) {
	a.proxy = proxy
}

//SetStoreID SetStoreID
func (a *Six910API) SetStoreID(sid int64) {
	a.storeID = sid
}
