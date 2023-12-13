// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	models "products/models"

	mock "github.com/stretchr/testify/mock"
)

// ProductServiceType is an autogenerated mock type for the ProductServiceType type
type ProductServiceType struct {
	mock.Mock
}

// DeleteProduct provides a mock function with given fields: id
func (_m *ProductServiceType) DeleteProduct(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllProduct provides a mock function with given fields:
func (_m *ProductServiceType) GetAllProduct() map[int]models.ProductData {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllProduct")
	}

	var r0 map[int]models.ProductData
	if rf, ok := ret.Get(0).(func() map[int]models.ProductData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]models.ProductData)
		}
	}

	return r0
}

// GetAvgPriceAndTotalQuantityByCategory provides a mock function with given fields: categry
func (_m *ProductServiceType) GetAvgPriceAndTotalQuantityByCategory(categry string) (float64, int, error) {
	ret := _m.Called(categry)

	if len(ret) == 0 {
		panic("no return value specified for GetAvgPriceAndTotalQuantityByCategory")
	}

	var r0 float64
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (float64, int, error)); ok {
		return rf(categry)
	}
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(categry)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(categry)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(categry)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// InsertNewProduct provides a mock function with given fields: data
func (_m *ProductServiceType) InsertNewProduct(data models.Product) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for InsertNewProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Product) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchByCategoryAndPriceRange provides a mock function with given fields: _a0
func (_m *ProductServiceType) SearchByCategoryAndPriceRange(_a0 models.SearchByCategoryAndPriceRangeModel) map[int]models.ProductData {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SearchByCategoryAndPriceRange")
	}

	var r0 map[int]models.ProductData
	if rf, ok := ret.Get(0).(func(models.SearchByCategoryAndPriceRangeModel) map[int]models.ProductData); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]models.ProductData)
		}
	}

	return r0
}

// UpdateProduct provides a mock function with given fields: id, data
func (_m *ProductServiceType) UpdateProduct(id int, data models.UpdateProductData) error {
	ret := _m.Called(id, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, models.UpdateProductData) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductServiceType creates a new instance of ProductServiceType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductServiceType(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductServiceType {
	mock := &ProductServiceType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}