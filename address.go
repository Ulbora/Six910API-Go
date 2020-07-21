package six910api

import sdbi "github.com/Ulbora/six910-database-interface"

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

//AddAddress AddAddress
func (a *Six910API) AddAddress(ad *sdbi.Address, headers *Headers) *ResponseID {
	var rtn ResponseID
	var url = a.restURL + "/rs/address/add"
	for k, v := range headers.headers {
		a.log.Debug("header: ", k, v)
	}
	a.log.Debug("url: ", url)

	return &rtn
}
