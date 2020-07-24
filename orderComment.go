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

//OrderCommentReq OrderCommentReq
type OrderCommentReq struct {
	StoreID      int64              `json:"storeId"`
	OrderComment *sdbi.OrderComment `json:"orderComment"`
}

//AddOrderComments AddOrderComments
func (a *Six910API) AddOrderComments(c *sdbi.OrderComment, headers *Headers) *ResponseID {
	var rtn ResponseID
	var odcReq OrderCommentReq
	odcReq.OrderComment = c
	odcReq.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/orderComment/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(odcReq)
	if err == nil {
		reqoca := a.buildRequest(post, url, headers, aJSON)
		sucoca, stat := a.proxy.Do(reqoca, &rtn)
		a.log.Debug("suc: ", sucoca)
		a.log.Debug("stat: ", stat)
		if !sucoca {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetOrderCommentList GetOrderCommentList
func (a *Six910API) GetOrderCommentList(orderID int64, headers *Headers) *[]sdbi.OrderComment {
	var rtn []sdbi.OrderComment
	var sid = a.getStoreID(headers)
	oidStrgodl := strconv.FormatInt(orderID, 10)
	sidStrGodl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/orderComment/get/list/" + oidStrgodl + "/" + sidStrGodl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	oclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", oclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}
