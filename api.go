package six910api

import (
	"sync"

	lg "github.com/Ulbora/Level_Logger"
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

//CustomerPasswordResponse CustomerPasswordResponse
type CustomerPasswordResponse struct {
	Success  bool   `json:"success"`
	Username string `json:"username"`
	Password string `json:"password"`
	Code     int64  `json:"code"`
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
	GetHierarchicalCategoryList(headers *Headers) *[]sdbi.Category
	GetCategoryList(headers *Headers) *[]sdbi.Category
	GetSubCategoryList(catID int64, headers *Headers) *[]sdbi.Category
	DeleteCategory(id int64, headers *Headers) *Response

	//customer
	AddCustomer(c *sdbi.Customer, headers *Headers) *ResponseID
	UpdateCustomer(c *sdbi.Customer, headers *Headers) *Response
	GetCustomer(email string, headers *Headers) *sdbi.Customer
	GetCustomerID(id int64, headers *Headers) *sdbi.Customer
	GetCustomerList(start int64, end int64, headers *Headers) *[]sdbi.Customer
	DeleteCustomer(id int64, headers *Headers) *Response

	//dataStoreWriteLock
	AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *Headers) *Response
	UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *Headers) *Response
	GetDataStoreWriteLock(dataStore string, headers *Headers) *sdbi.DataStoreWriteLock

	//dataStore
	AddLocalDatastore(d *sdbi.LocalDataStore, headers *Headers) *Response
	UpdateLocalDatastore(d *sdbi.LocalDataStore, headers *Headers) *Response
	GetLocalDatastore(dataStoreName string, headers *Headers) *sdbi.LocalDataStore

	//distrubutor
	AddDistributor(d *sdbi.Distributor, headers *Headers) *ResponseID
	UpdateDistributor(d *sdbi.Distributor, headers *Headers) *Response
	GetDistributor(id int64, headers *Headers) *sdbi.Distributor
	GetDistributorList(headers *Headers) *[]sdbi.Distributor
	DeleteDistributor(id int64, headers *Headers) *Response

	//excluded sub region
	AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, headers *Headers) *ResponseID
	GetExcludedSubRegionList(regionID int64, headers *Headers) *[]sdbi.ExcludedSubRegion
	DeleteExcludedSubRegion(id int64, regionID int64, headers *Headers) *Response

	//included sub region
	AddIncludedSubRegion(e *sdbi.IncludedSubRegion, headers *Headers) *ResponseID
	GetIncludedSubRegionList(regionID int64, headers *Headers) *[]sdbi.IncludedSubRegion
	DeleteIncludedSubRegion(id int64, regionID int64, headers *Headers) *Response

	//instances
	AddInstance(i *sdbi.Instances, headers *Headers) *Response
	UpdateInstance(i *sdbi.Instances, headers *Headers) *Response
	GetInstance(name string, dataStoreName string, headers *Headers) *sdbi.Instances
	GetInstanceList(dataStoreName string, headers *Headers) *[]sdbi.Instances

	// insurance
	AddInsurance(i *sdbi.Insurance, headers *Headers) *ResponseID
	UpdateInsurance(i *sdbi.Insurance, headers *Headers) *Response
	GetInsurance(id int64, headers *Headers) *sdbi.Insurance
	GetInsuranceList(headers *Headers) *[]sdbi.Insurance
	DeleteInsurance(id int64, headers *Headers) *Response

	AddTaxRate(t *sdbi.TaxRate, headers *Headers) *ResponseID
	UpdateTaxRate(t *sdbi.TaxRate, headers *Headers) *Response
	GetTaxRate(country string, state string, headers *Headers) *[]sdbi.TaxRate
	GetTaxRateList(headers *Headers) *[]sdbi.TaxRate
	DeleteTaxRate(id int64, headers *Headers) *Response

	//order
	AddOrder(o *sdbi.Order, headers *Headers) *ResponseID
	UpdateOrder(o *sdbi.Order, headers *Headers) *Response
	GetOrder(id int64, headers *Headers) *sdbi.Order
	GetOrderList(cid int64, headers *Headers) *[]sdbi.Order
	GetStoreOrderList(headers *Headers) *[]sdbi.Order
	GetStoreOrderListByStatus(status string, headers *Headers) *[]sdbi.Order
	GetOrderCountData(headers *Headers) *[]sdbi.OrderCountData
	GetOrderSalesData(headers *Headers) *[]sdbi.OrderSalesData
	DeleteOrder(id int64, headers *Headers) *Response

	//order comments
	AddOrderComments(c *sdbi.OrderComment, headers *Headers) *ResponseID
	GetOrderCommentList(orderID int64, headers *Headers) *[]sdbi.OrderComment

	//order items
	AddOrderItem(i *sdbi.OrderItem, headers *Headers) *ResponseID
	UpdateOrderItem(i *sdbi.OrderItem, headers *Headers) *Response
	GetOrderItem(id int64, headers *Headers) *sdbi.OrderItem
	GetOrderItemList(orderID int64, headers *Headers) *[]sdbi.OrderItem
	DeleteOrderItem(id int64, headers *Headers) *Response

	//order transaction
	AddOrderTransaction(t *sdbi.OrderTransaction, headers *Headers) *ResponseID
	GetOrderTransactionList(orderID int64, headers *Headers) *[]sdbi.OrderTransaction

	//payment gateway
	AddPaymentGateway(pgw *sdbi.PaymentGateway, headers *Headers) *ResponseID
	UpdatePaymentGateway(pgw *sdbi.PaymentGateway, headers *Headers) *Response
	GetPaymentGateway(id int64, headers *Headers) *sdbi.PaymentGateway
	GetPaymentGateways(headers *Headers) *[]sdbi.PaymentGateway
	DeletePaymentGateway(id int64, headers *Headers) *Response

	//plugins
	AddPlugin(p *sdbi.Plugins, headers *Headers) *ResponseID
	UpdatePlugin(p *sdbi.Plugins, headers *Headers) *Response
	GetPlugin(id int64, headers *Headers) *sdbi.Plugins
	GetPluginList(start int64, end int64, headers *Headers) *[]sdbi.Plugins
	DeletePlugin(id int64, headers *Headers) *Response

	//products
	AddProduct(p *sdbi.Product, headers *Headers) *ResponseID
	UpdateProduct(p *sdbi.Product, headers *Headers) *Response
	UpdateProductQuantity(p *sdbi.Product, headers *Headers) *Response
	GetProductByID(id int64, headers *Headers) *sdbi.Product
	GetProductBySku(sku string, did int64, headers *Headers) *sdbi.Product
	GetProductsByPromoted(start int64, end int64, headers *Headers) *[]sdbi.Product
	GetProductsByName(name string, start int64, end int64, headers *Headers) *[]sdbi.Product
	GetProductsByCaterory(catID int64, start int64, end int64, headers *Headers) *[]sdbi.Product
	GetProductList(start int64, end int64, headers *Headers) *[]sdbi.Product
	GetProductIDList(headers *Headers) *[]int64
	GetProductIDListByCategories(idReq *ProdIDReq, headers *Headers) *[]int64
	DeleteProduct(id int64, headers *Headers) *Response

	GetProductManufacturerListByProductName(name string, headers *Headers) *[]string
	GetProductByNameAndManufacturerName(manf string, name string,
		start int64, end int64, headers *Headers) *[]sdbi.Product
	GetProductManufacturerListByCatID(catID int64, headers *Headers) *[]string
	GetProductByCatAndManufacturer(catID int64, manf string,
		start int64, end int64, headers *Headers) *[]sdbi.Product

	//product category
	AddProductCategory(pc *sdbi.ProductCategory, headers *Headers) *Response
	GetProductCategoryList(productID int64, headers *Headers) []int64
	DeleteProductCategory(pc *sdbi.ProductCategory, headers *Headers) *Response

	//region
	AddRegion(r *sdbi.Region, headers *Headers) *ResponseID
	UpdateRegion(r *sdbi.Region, headers *Headers) *Response
	GetRegion(id int64, headers *Headers) *sdbi.Region
	GetRegionList(headers *Headers) *[]sdbi.Region
	DeleteRegion(id int64, headers *Headers) *Response

	//shipment
	AddShipment(s *sdbi.Shipment, headers *Headers) *ResponseID
	UpdateShipment(s *sdbi.Shipment, headers *Headers) *Response
	GetShipment(id int64, headers *Headers) *sdbi.Shipment
	GetShipmentList(orderID int64, headers *Headers) *[]sdbi.Shipment
	DeleteShipment(id int64, headers *Headers) *Response

	//shipment box
	AddShipmentBox(sb *sdbi.ShipmentBox, headers *Headers) *ResponseID
	UpdateShipmentBox(sb *sdbi.ShipmentBox, headers *Headers) *Response
	GetShipmentBox(id int64, headers *Headers) *sdbi.ShipmentBox
	GetShipmentBoxList(shipmentID int64, headers *Headers) *[]sdbi.ShipmentBox
	DeleteShipmentBox(id int64, headers *Headers) *Response

	//shipment item
	AddShipmentItem(si *sdbi.ShipmentItem, headers *Headers) *ResponseID
	UpdateShipmentItem(si *sdbi.ShipmentItem, headers *Headers) *Response
	GetShipmentItem(id int64, headers *Headers) *sdbi.ShipmentItem
	GetShipmentItemList(shipmentID int64, headers *Headers) *[]sdbi.ShipmentItem
	GetShipmentItemListByBox(boxNumber int64, shipmentID int64, headers *Headers) *[]sdbi.ShipmentItem
	DeleteShipmentItem(id int64, headers *Headers) *Response

	//shipment carrier
	AddShippingCarrier(c *sdbi.ShippingCarrier, headers *Headers) *ResponseID
	UpdateShippingCarrier(c *sdbi.ShippingCarrier, headers *Headers) *Response
	GetShippingCarrier(id int64, headers *Headers) *sdbi.ShippingCarrier
	GetShippingCarrierList(headers *Headers) *[]sdbi.ShippingCarrier
	DeleteShippingCarrier(id int64, headers *Headers) *Response

	//shipment method
	AddShippingMethod(s *sdbi.ShippingMethod, headers *Headers) *ResponseID
	UpdateShippingMethod(s *sdbi.ShippingMethod, headers *Headers) *Response
	GetShippingMethod(id int64, headers *Headers) *sdbi.ShippingMethod
	GetShippingMethodList(headers *Headers) *[]sdbi.ShippingMethod
	DeleteShippingMethod(id int64, headers *Headers) *Response

	//store
	AddStore(s *sdbi.Store, headers *Headers) *ResponseID
	UpdateStore(s *sdbi.Store, headers *Headers) *Response
	GetStore(sname string, localDomain string, headers *Headers) *sdbi.Store
	DeleteStore(sname string, localDomain string, headers *Headers) *Response

	//store plugin
	AddStorePlugin(sp *sdbi.StorePlugins, headers *Headers) *ResponseID
	UpdateStorePlugin(sp *sdbi.StorePlugins, headers *Headers) *Response
	GetStorePlugin(id int64, headers *Headers) *sdbi.StorePlugins
	GetStorePluginList(headers *Headers) *[]sdbi.StorePlugins
	DeleteStorePlugin(id int64, headers *Headers) *Response

	//sub region
	AddSubRegion(s *sdbi.SubRegion, headers *Headers) *ResponseID
	UpdateSubRegion(s *sdbi.SubRegion, headers *Headers) *Response
	GetSubRegion(id int64, headers *Headers) *sdbi.SubRegion
	GetSubRegionList(regionID int64, headers *Headers) *[]sdbi.SubRegion
	DeleteSubRegion(id int64, headers *Headers) *Response

	//user
	AddCustomerUser(u *User, headers *Headers) *Response
	AddAdminUser(u *User, headers *Headers) *Response
	UpdateUser(u *User, headers *Headers) *Response
	AdminUpdateUser(u *User, headers *Headers) *Response
	GetUser(u *User, headers *Headers) *UserResponse
	GetAdminUsers(headers *Headers) *[]UserResponse
	GetCustomerUsers(headers *Headers) *[]UserResponse
	GetUsersByCustomer(cid int64, headers *Headers) *[]UserResponse
	ResetCustomerUserPassword(u *User, headers *Headers) *CustomerPasswordResponse

	//zip code zone
	AddZoneZip(z *sdbi.ZoneZip, headers *Headers) *ResponseID
	GetZoneZipListByExclusion(exID int64, headers *Headers) *[]sdbi.ZoneZip
	GetZoneZipListByInclusion(incID int64, headers *Headers) *[]sdbi.ZoneZip
	DeleteZoneZip(id int64, incID int64, exID int64, headers *Headers) *Response

	//visitor
	AddVisit(v *sdbi.Visitor, headers *Headers) *Response
	GetVisitorData(headers *Headers) *[]sdbi.VisitorData

	SetLogger(l *lg.Logger)
}
