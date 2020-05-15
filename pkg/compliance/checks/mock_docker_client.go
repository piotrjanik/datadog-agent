// Code generated by mockery v1.0.0. DO NOT EDIT.

package checks

import (
	context "context"

	types "github.com/docker/docker/api/types"
	mock "github.com/stretchr/testify/mock"
)

// MockDockerClient is an autogenerated mock type for the DockerClient type
type MockDockerClient struct {
	mock.Mock
}

// ImageInspectWithRaw provides a mock function with given fields: ctx, image
func (_m *MockDockerClient) ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error) {
	ret := _m.Called(ctx, image)

	var r0 types.ImageInspect
	if rf, ok := ret.Get(0).(func(context.Context, string) types.ImageInspect); ok {
		r0 = rf(ctx, image)
	} else {
		r0 = ret.Get(0).(types.ImageInspect)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, image)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, image)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImageList provides a mock function with given fields: ctx, options
func (_m *MockDockerClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	ret := _m.Called(ctx, options)

	var r0 []types.ImageSummary
	if rf, ok := ret.Get(0).(func(context.Context, types.ImageListOptions) []types.ImageSummary); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ImageSummary)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ImageListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}