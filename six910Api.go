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
	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

//Six910API Six910API
type Six910API struct {
	proxy   px.Proxy
	log     *lg.Logger
	RestURL string
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

//SetLogLever SetLogLever
func (a *Six910API) SetLogLever(level int) {
	a.log.LogLevel = level
}
