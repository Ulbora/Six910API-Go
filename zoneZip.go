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

//ZoneZipReq ZoneZipReq
type ZoneZipReq struct {
	StoreID int64         `json:"storeId"`
	ZoneZip *sdbi.ZoneZip `json:"zoneZip"`
}

//AddZoneZip AddZoneZip
func (a *Six910API) AddZoneZip(z *sdbi.ZoneZip, headers *Headers) *ResponseID {
	var rtn ResponseID
	var zzr ZoneZipReq
	zzr.ZoneZip = z
	zzr.StoreID = a.getStoreID(headers)
	var url = a.restURL + "/rs/zoneZip/add"
	a.log.Debug("url: ", url)
	aJSON, err := json.Marshal(zzr)
	if err == nil {
		reqazz := a.buildRequest(post, url, headers, aJSON)
		sucazz, stat := a.proxy.Do(reqazz, &rtn)
		a.log.Debug("suc: ", sucazz)
		a.log.Debug("stat: ", stat)
		if !sucazz {
			rtn.Code = int64(stat)
		}
	}
	a.log.Debug("rtn: ", rtn)
	return &rtn
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (a *Six910API) GetZoneZipListByExclusion(exID int64, headers *Headers) *[]sdbi.ZoneZip {
	var rtn []sdbi.ZoneZip
	var sid = a.getStoreID(headers)
	excidStrGctl := strconv.FormatInt(exID, 10)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/zoneZip/exc/get/list/" + excidStrGctl + "/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	gexclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", gexclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (a *Six910API) GetZoneZipListByInclusion(incID int64, headers *Headers) *[]sdbi.ZoneZip {
	var rtn []sdbi.ZoneZip
	var sid = a.getStoreID(headers)
	inccidStrGctl := strconv.FormatInt(incID, 10)
	sidStrGctl := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/zoneZip/inc/get/list/" + inccidStrGctl + "/" + sidStrGctl
	a.log.Debug("url: ", url)

	req := a.buildRequest(get, url, headers, nil)
	ginclsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", ginclsuc)
	a.log.Debug("stat: ", stat)

	return &rtn
}

//DeleteZoneZip DeleteZoneZip
func (a *Six910API) DeleteZoneZip(id int64, incID int64, exID int64, headers *Headers) *Response {
	var rtn Response
	var sid = a.getStoreID(headers)
	zzidStrdct := strconv.FormatInt(id, 10)
	zzincIDStrdct := strconv.FormatInt(incID, 10)
	zzexIDStrdct := strconv.FormatInt(exID, 10)
	sidStrdct := strconv.FormatInt(sid, 10)

	var url = a.restURL + "/rs/zoneZip/delete/" + zzidStrdct + "/" + zzincIDStrdct + "/" + zzexIDStrdct + "/" + sidStrdct
	a.log.Debug("url: ", url)

	req := a.buildRequest(delete, url, headers, nil)
	dzzsuc, stat := a.proxy.Do(req, &rtn)
	a.log.Debug("suc: ", dzzsuc)
	a.log.Debug("stat: ", stat)
	return &rtn
}
