package six910api

import (
	"sync"

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

	sdbi "github.com/Ulbora/six910-database-interface"
)

const (
	post   = "POST"
	put    = "PUT"
	get    = "GET"
	delete = "DELETE"
)

//ResponseID ResponseID
type ResponseID struct {
	ID      int64  `json:"id"`
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//Response Response
type Response struct {
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

//User User
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	OldPassword string `json:"oldPassword"`
	Role        string `json:"role"`
	CustomerID  int64  `json:"customerId"`
	StoreID     int64  `json:"storeId"`
	Enabled     bool   `json:"enabled"`
}

//UserResponse UserResponse
type UserResponse struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	CustomerID int64  `json:"customerId"`
	StoreID    int64  `json:"storeId"`
	Enabled    bool   `json:"enabled"`
}

//Headers Headers
type Headers struct {
	headers map[string]string
	mu      sync.Mutex
}

//Set Set
func (h *Headers) Set(key string, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

// go mod init github.com/Ulbora/Six910API-Go

//API API
type API interface {
	//address
	AddAddress(a *sdbi.Address, headers *Headers) *ResponseID
	UpdateAddress(a *sdbi.Address, headers *Headers) *Response
	GetAddress(id int64, cid int64, headers *Headers) *sdbi.Address
	GetAddressList(cid int64, headers *Headers) *[]sdbi.Address
	DeleteAddress(id int64, cid int64, headers *Headers) *Response

	//cart
	AddCart(c *sdbi.Cart, headers *Headers) *ResponseID
	UpdateCart(c *sdbi.Cart, headers *Headers) *Response
	GetCart(cid int64, headers *Headers) *sdbi.Cart
	DeleteCart(id int64, cid int64, headers *Headers) *Response

	//cartItem
	AddCartItem(ci *sdbi.CartItem, cid int64, headers *Headers) *ResponseID
	UpdateCartItem(ci *sdbi.CartItem, cid int64, headers *Headers) *Response
	GetCartItem(cid int64, prodID int64, headers *Headers) *sdbi.CartItem
	GetCartItemList(cartID int64, cid int64, headers *Headers) *[]sdbi.CartItem
	DeleteCartItem(id int64, prodID int64, cartID int64, headers *Headers) *Response

	//category
	AddCategory(c *sdbi.Category, headers *Headers) *ResponseID
	UpdateCategory(c *sdbi.Category, headers *Headers) *Response
	GetCategory(id int64, headers *Headers) *sdbi.Category
	GetCategoryList(headers *Headers) *[]sdbi.Category
	GetSubCategoryList(catID int64, headers *Headers) *[]sdbi.Category
	DeleteCategory(id int64, headers *Headers) *Response

	// //customer
	// AddCustomer(c *sdbi.Customer) *ResponseID
	// UpdateCustomer(c *sdbi.Customer) *Response
	// GetCustomer(email string, storeID int64) *sdbi.Customer
	// GetCustomerID(id int64, storeID int64) *sdbi.Customer
	// GetCustomerList(storeID int64) *[]sdbi.Customer
	// DeleteCustomer(id int64, storeID int64) *Response

	// //dataStoreWriteLock
	// AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response
	// UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock) *Response
	// GetDataStoreWriteLock(dataStore string, storeID int64) *sdbi.DataStoreWriteLock

	// //dataStore
	// AddLocalDatastore(d *sdbi.LocalDataStore) *Response
	// UpdateLocalDatastore(d *sdbi.LocalDataStore) *Response
	// GetLocalDatastore(storeID int64, dataStoreName string) *sdbi.LocalDataStore

	// //distrubutor
	// AddDistributor(d *sdbi.Distributor) *ResponseID
	// UpdateDistributor(d *sdbi.Distributor) *Response
	// GetDistributor(id int64, storeID int64) *sdbi.Distributor
	// GetDistributorList(storeID int64) *[]sdbi.Distributor
	// DeleteDistributor(id int64, storeID int64) *Response

	// //excluded sub region
	// AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *ResponseID
	// UpdateExcludedSubRegion(e *sdbi.ExcludedSubRegion, sid int64) *Response
	// GetExcludedSubRegion(id int64, sid int64) *sdbi.ExcludedSubRegion
	// GetExcludedSubRegionList(regionID int64, sid int64) *[]sdbi.ExcludedSubRegion
	// DeleteExcludedSubRegion(id int64, regionID int64, sid int64) *Response

	// //included sub region
	// AddIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *ResponseID
	// UpdateIncludedSubRegion(e *sdbi.IncludedSubRegion, sid int64) *Response
	// GetIncludedSubRegion(id int64, sid int64) *sdbi.IncludedSubRegion
	// GetIncludedSubRegionList(regionID int64, sid int64) *[]sdbi.IncludedSubRegion
	// DeleteIncludedSubRegion(id int64, regionID int64, sid int64) *Response

	// //instances
	// AddInstance(i *sdbi.Instances) *Response
	// UpdateInstance(i *sdbi.Instances) *Response
	// GetInstance(name string, dataStoreName string, storeID int64) *sdbi.Instances

	// //insurance
	// AddInsurance(i *sdbi.Insurance) *ResponseID
	// UpdateInsurance(i *sdbi.Insurance) *Response
	// GetInsurance(id int64, sid int64) *sdbi.Insurance
	// GetInsuranceList(storeID int64) *[]sdbi.Insurance
	// DeleteInsurance(id int64, sid int64) *Response

	// //order
	// AddOrder(o *sdbi.Order) *ResponseID
	// UpdateOrder(o *sdbi.Order) *Response
	// GetOrder(id int64, sid int64) *sdbi.Order
	// GetOrderList(cid int64, sid int64) *[]sdbi.Order
	// DeleteOrder(id int64, sid int64) *Response

	// //order comments
	// AddOrderComments(c *sdbi.OrderComment, sid int64) *ResponseID
	// GetOrderCommentList(orderID int64, sid int64) *[]sdbi.OrderComment

	// //order items
	// AddOrderItem(i *sdbi.OrderItem, sid int64) *ResponseID
	// UpdateOrderItem(i *sdbi.OrderItem, sid int64) *Response
	// GetOrderItem(id int64, sid int64) *sdbi.OrderItem
	// GetOrderItemList(orderID int64, sid int64) *[]sdbi.OrderItem
	// DeleteOrderItem(id int64, sid int64) *Response

	// //order transaction
	// AddOrderTransaction(t *sdbi.OrderTransaction, sid int64) *ResponseID
	// GetOrderTransactionList(orderID int64, sid int64) *[]sdbi.OrderTransaction

	// //payment gateway
	// AddPaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *ResponseID
	// UpdatePaymentGateway(pgw *sdbi.PaymentGateway, sid int64) *Response
	// GetPaymentGateway(id int64, sid int64) *sdbi.PaymentGateway
	// GetPaymentGateways(storeID int64) *[]sdbi.PaymentGateway
	// DeletePaymentGateway(id int64, sid int64) *Response

	// //plugins
	// AddPlugin(p *sdbi.Plugins) *ResponseID
	// UpdatePlugin(p *sdbi.Plugins) *Response
	// GetPlugin(id int64) *sdbi.Plugins
	// GetPluginList(start int64, end int64) *[]sdbi.Plugins
	// DeletePlugin(id int64) *Response

	// //products
	// AddProduct(p *sdbi.Product) *ResponseID
	// UpdateProduct(p *sdbi.Product) *Response
	// GetProductByID(id int64, sid int64) *sdbi.Product
	// GetProductsByName(name string, sid int64, start int64, end int64) *[]sdbi.Product
	// GetProductsByCaterory(catID int64, sid int64, start int64, end int64) *[]sdbi.Product
	// GetProductList(storeID int64, start int64, end int64) *[]sdbi.Product
	// DeleteProduct(id int64, sid int64) *Response

	// //product category
	// AddProductCategory(pc *sdbi.ProductCategory, sid int64) *Response
	// DeleteProductCategory(pc *sdbi.ProductCategory, sid int64) *Response

	// //region
	// AddRegion(r *sdbi.Region) *ResponseID
	// UpdateRegion(r *sdbi.Region) *Response
	// GetRegion(id int64, sid int64) *sdbi.Region
	// GetRegionList(storeID int64) *[]sdbi.Region
	// DeleteRegion(id int64, sid int64) *Response

	// //shipment
	// AddShipment(s *sdbi.Shipment, sid int64) *ResponseID
	// UpdateShipment(s *sdbi.Shipment, sid int64) *Response
	// GetShipment(id int64, sid int64) *sdbi.Shipment
	// GetShipmentList(orderID int64, sid int64) *[]sdbi.Shipment
	// DeleteShipment(id int64, sid int64) *Response

	// //shipment box
	// AddShipmentBox(sb *sdbi.ShipmentBox, sid int64) *ResponseID
	// UpdateShipmentBox(sb *sdbi.ShipmentBox, sid int64) *Response
	// GetShipmentBox(id int64, sid int64) *sdbi.ShipmentBox
	// GetShipmentBoxList(shipmentID int64, sid int64) *[]sdbi.ShipmentBox
	// DeleteShipmentBox(id int64, sid int64) *Response

	// //shipment item
	// AddShipmentItem(si *sdbi.ShipmentItem, sid int64) *ResponseID
	// UpdateShipmentItem(si *sdbi.ShipmentItem, sid int64) *Response
	// GetShipmentItem(id int64, sid int64) *sdbi.ShipmentItem
	// GetShipmentItemList(shipmentID int64, sid int64) *[]sdbi.ShipmentItem
	// GetShipmentItemListByBox(boxNumber int64, shipmentID int64, sid int64) *[]sdbi.ShipmentItem
	// DeleteShipmentItem(id int64, sid int64) *Response

	// //shipment carrier
	// AddShippingCarrier(c *sdbi.ShippingCarrier) *ResponseID
	// UpdateShippingCarrier(c *sdbi.ShippingCarrier) *Response
	// GetShippingCarrier(id int64, sid int64) *sdbi.ShippingCarrier
	// GetShippingCarrierList(storeID int64) *[]sdbi.ShippingCarrier
	// DeleteShippingCarrier(id int64, sid int64) *Response

	// //shipment method
	// AddShippingMethod(s *sdbi.ShippingMethod) *ResponseID
	// UpdateShippingMethod(s *sdbi.ShippingMethod) *Response
	// GetShippingMethod(id int64, sid int64) *sdbi.ShippingMethod
	// GetShippingMethodList(storeID int64) *[]sdbi.ShippingMethod
	// DeleteShippingMethod(id int64, sid int64) *Response

	// //store
	// AddStore(s *sdbi.Store) *ResponseID
	// UpdateStore(s *sdbi.Store) *Response
	// GetStore(sname string, localDomain string) *sdbi.Store
	// DeleteStore(sname string, localDomain string) *Response

	// //store plugin
	// AddStorePlugin(sp *sdbi.StorePlugins) *ResponseID
	// UpdateStorePlugin(sp *sdbi.StorePlugins) *Response
	// GetStorePlugin(id int64, sid int64) *sdbi.StorePlugins
	// GetStorePluginList(storeID int64) *[]sdbi.StorePlugins
	// DeleteStorePlugin(id int64, sid int64) *Response

	// //sub region
	// AddSubRegion(s *sdbi.SubRegion, sid int64) *ResponseID
	// UpdateSubRegion(s *sdbi.SubRegion, sid int64) *Response
	// GetSubRegion(id int64, sid int64) *sdbi.SubRegion
	// GetSubRegionList(regionID int64, sid int64) *[]sdbi.SubRegion
	// DeleteSubRegion(id int64, sid int64) *Response

	// //user
	// AddCustomerUser(u *User) *Response
	// UpdateUser(u *User) *Response
	// GetUser(u *User) *UserResponse
	// GetAdminUsers(storeID int64) *[]UserResponse
	// GetCustomerUsers(storeID int64) *[]UserResponse

	// //zip code zone
	// AddZoneZip(z *sdbi.ZoneZip, sid int64) *ResponseID
	// GetZoneZipListByExclusion(exID int64, sid int64) *[]sdbi.ZoneZip
	// GetZoneZipListByInclusion(incID int64, sid int64) *[]sdbi.ZoneZip
	// DeleteZoneZip(id int64, incID int64, exID int64, sid int64) *Response
}
