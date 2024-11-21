// Code generated by mockery v2.43.2. DO NOT EDIT.

package core

import (
	context "context"
	io "io"

	core "github.com/MAXXXIMUS-tropical-milkshake/drop-audio-streaming/internal/core"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MockBeatStorage is an autogenerated mock type for the BeatStorage type
type MockBeatStorage struct {
	mock.Mock
}

type MockBeatStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBeatStorage) EXPECT() *MockBeatStorage_Expecter {
	return &MockBeatStorage_Expecter{mock: &_m.Mock}
}

// AddBeat provides a mock function with given fields: ctx, beat, beatGenre
func (_m *MockBeatStorage) AddBeat(ctx context.Context, beat core.Beat, beatGenre []core.BeatGenre) (int, error) {
	ret := _m.Called(ctx, beat, beatGenre)

	if len(ret) == 0 {
		panic("no return value specified for AddBeat")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, core.Beat, []core.BeatGenre) (int, error)); ok {
		return rf(ctx, beat, beatGenre)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.Beat, []core.BeatGenre) int); ok {
		r0 = rf(ctx, beat, beatGenre)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.Beat, []core.BeatGenre) error); ok {
		r1 = rf(ctx, beat, beatGenre)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBeatStorage_AddBeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBeat'
type MockBeatStorage_AddBeat_Call struct {
	*mock.Call
}

// AddBeat is a helper method to define mock.On call
//   - ctx context.Context
//   - beat core.Beat
//   - beatGenre []core.BeatGenre
func (_e *MockBeatStorage_Expecter) AddBeat(ctx interface{}, beat interface{}, beatGenre interface{}) *MockBeatStorage_AddBeat_Call {
	return &MockBeatStorage_AddBeat_Call{Call: _e.mock.On("AddBeat", ctx, beat, beatGenre)}
}

func (_c *MockBeatStorage_AddBeat_Call) Run(run func(ctx context.Context, beat core.Beat, beatGenre []core.BeatGenre)) *MockBeatStorage_AddBeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.Beat), args[2].([]core.BeatGenre))
	})
	return _c
}

func (_c *MockBeatStorage_AddBeat_Call) Return(beatID int, err error) *MockBeatStorage_AddBeat_Call {
	_c.Call.Return(beatID, err)
	return _c
}

func (_c *MockBeatStorage_AddBeat_Call) RunAndReturn(run func(context.Context, core.Beat, []core.BeatGenre) (int, error)) *MockBeatStorage_AddBeat_Call {
	_c.Call.Return(run)
	return _c
}

// AddUserSeenBeat provides a mock function with given fields: ctx, userID, beatID
func (_m *MockBeatStorage) AddUserSeenBeat(ctx context.Context, userID int, beatID int) error {
	ret := _m.Called(ctx, userID, beatID)

	if len(ret) == 0 {
		panic("no return value specified for AddUserSeenBeat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) error); ok {
		r0 = rf(ctx, userID, beatID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBeatStorage_AddUserSeenBeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserSeenBeat'
type MockBeatStorage_AddUserSeenBeat_Call struct {
	*mock.Call
}

// AddUserSeenBeat is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
//   - beatID int
func (_e *MockBeatStorage_Expecter) AddUserSeenBeat(ctx interface{}, userID interface{}, beatID interface{}) *MockBeatStorage_AddUserSeenBeat_Call {
	return &MockBeatStorage_AddUserSeenBeat_Call{Call: _e.mock.On("AddUserSeenBeat", ctx, userID, beatID)}
}

func (_c *MockBeatStorage_AddUserSeenBeat_Call) Run(run func(ctx context.Context, userID int, beatID int)) *MockBeatStorage_AddUserSeenBeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *MockBeatStorage_AddUserSeenBeat_Call) Return(_a0 error) *MockBeatStorage_AddUserSeenBeat_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBeatStorage_AddUserSeenBeat_Call) RunAndReturn(run func(context.Context, int, int) error) *MockBeatStorage_AddUserSeenBeat_Call {
	_c.Call.Return(run)
	return _c
}

// ClearUserSeenBeats provides a mock function with given fields: ctx, userID
func (_m *MockBeatStorage) ClearUserSeenBeats(ctx context.Context, userID int) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for ClearUserSeenBeats")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBeatStorage_ClearUserSeenBeats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearUserSeenBeats'
type MockBeatStorage_ClearUserSeenBeats_Call struct {
	*mock.Call
}

// ClearUserSeenBeats is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
func (_e *MockBeatStorage_Expecter) ClearUserSeenBeats(ctx interface{}, userID interface{}) *MockBeatStorage_ClearUserSeenBeats_Call {
	return &MockBeatStorage_ClearUserSeenBeats_Call{Call: _e.mock.On("ClearUserSeenBeats", ctx, userID)}
}

func (_c *MockBeatStorage_ClearUserSeenBeats_Call) Run(run func(ctx context.Context, userID int)) *MockBeatStorage_ClearUserSeenBeats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockBeatStorage_ClearUserSeenBeats_Call) Return(_a0 error) *MockBeatStorage_ClearUserSeenBeats_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBeatStorage_ClearUserSeenBeats_Call) RunAndReturn(run func(context.Context, int) error) *MockBeatStorage_ClearUserSeenBeats_Call {
	_c.Call.Return(run)
	return _c
}

// GetBeatByFilter provides a mock function with given fields: ctx, filter, seen
func (_m *MockBeatStorage) GetBeatByFilter(ctx context.Context, filter core.FeedFilter, seen []string) (*core.Beat, *string, error) {
	ret := _m.Called(ctx, filter, seen)

	if len(ret) == 0 {
		panic("no return value specified for GetBeatByFilter")
	}

	var r0 *core.Beat
	var r1 *string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, core.FeedFilter, []string) (*core.Beat, *string, error)); ok {
		return rf(ctx, filter, seen)
	}
	if rf, ok := ret.Get(0).(func(context.Context, core.FeedFilter, []string) *core.Beat); ok {
		r0 = rf(ctx, filter, seen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, core.FeedFilter, []string) *string); ok {
		r1 = rf(ctx, filter, seen)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, core.FeedFilter, []string) error); ok {
		r2 = rf(ctx, filter, seen)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBeatStorage_GetBeatByFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBeatByFilter'
type MockBeatStorage_GetBeatByFilter_Call struct {
	*mock.Call
}

// GetBeatByFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filter core.FeedFilter
//   - seen []string
func (_e *MockBeatStorage_Expecter) GetBeatByFilter(ctx interface{}, filter interface{}, seen interface{}) *MockBeatStorage_GetBeatByFilter_Call {
	return &MockBeatStorage_GetBeatByFilter_Call{Call: _e.mock.On("GetBeatByFilter", ctx, filter, seen)}
}

func (_c *MockBeatStorage_GetBeatByFilter_Call) Run(run func(ctx context.Context, filter core.FeedFilter, seen []string)) *MockBeatStorage_GetBeatByFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(core.FeedFilter), args[2].([]string))
	})
	return _c
}

func (_c *MockBeatStorage_GetBeatByFilter_Call) Return(beat *core.Beat, genre *string, err error) *MockBeatStorage_GetBeatByFilter_Call {
	_c.Call.Return(beat, genre, err)
	return _c
}

func (_c *MockBeatStorage_GetBeatByFilter_Call) RunAndReturn(run func(context.Context, core.FeedFilter, []string) (*core.Beat, *string, error)) *MockBeatStorage_GetBeatByFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetBeatByID provides a mock function with given fields: ctx, id, param
func (_m *MockBeatStorage) GetBeatByID(ctx context.Context, id int, param core.IsDownloaded) (*core.Beat, error) {
	ret := _m.Called(ctx, id, param)

	if len(ret) == 0 {
		panic("no return value specified for GetBeatByID")
	}

	var r0 *core.Beat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, core.IsDownloaded) (*core.Beat, error)); ok {
		return rf(ctx, id, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, core.IsDownloaded) *core.Beat); ok {
		r0 = rf(ctx, id, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, core.IsDownloaded) error); ok {
		r1 = rf(ctx, id, param)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBeatStorage_GetBeatByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBeatByID'
type MockBeatStorage_GetBeatByID_Call struct {
	*mock.Call
}

// GetBeatByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int
//   - param core.IsDownloaded
func (_e *MockBeatStorage_Expecter) GetBeatByID(ctx interface{}, id interface{}, param interface{}) *MockBeatStorage_GetBeatByID_Call {
	return &MockBeatStorage_GetBeatByID_Call{Call: _e.mock.On("GetBeatByID", ctx, id, param)}
}

func (_c *MockBeatStorage_GetBeatByID_Call) Run(run func(ctx context.Context, id int, param core.IsDownloaded)) *MockBeatStorage_GetBeatByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(core.IsDownloaded))
	})
	return _c
}

func (_c *MockBeatStorage_GetBeatByID_Call) Return(_a0 *core.Beat, _a1 error) *MockBeatStorage_GetBeatByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBeatStorage_GetBeatByID_Call) RunAndReturn(run func(context.Context, int, core.IsDownloaded) (*core.Beat, error)) *MockBeatStorage_GetBeatByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetBeatFromS3 provides a mock function with given fields: ctx, beatPath, start, end
func (_m *MockBeatStorage) GetBeatFromS3(ctx context.Context, beatPath string, start int, end *int) (io.ReadCloser, int, string, error) {
	ret := _m.Called(ctx, beatPath, start, end)

	if len(ret) == 0 {
		panic("no return value specified for GetBeatFromS3")
	}

	var r0 io.ReadCloser
	var r1 int
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, *int) (io.ReadCloser, int, string, error)); ok {
		return rf(ctx, beatPath, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int, *int) io.ReadCloser); ok {
		r0 = rf(ctx, beatPath, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int, *int) int); ok {
		r1 = rf(ctx, beatPath, start, end)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int, *int) string); ok {
		r2 = rf(ctx, beatPath, start, end)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, string, int, *int) error); ok {
		r3 = rf(ctx, beatPath, start, end)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockBeatStorage_GetBeatFromS3_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBeatFromS3'
type MockBeatStorage_GetBeatFromS3_Call struct {
	*mock.Call
}

// GetBeatFromS3 is a helper method to define mock.On call
//   - ctx context.Context
//   - beatPath string
//   - start int
//   - end *int
func (_e *MockBeatStorage_Expecter) GetBeatFromS3(ctx interface{}, beatPath interface{}, start interface{}, end interface{}) *MockBeatStorage_GetBeatFromS3_Call {
	return &MockBeatStorage_GetBeatFromS3_Call{Call: _e.mock.On("GetBeatFromS3", ctx, beatPath, start, end)}
}

func (_c *MockBeatStorage_GetBeatFromS3_Call) Run(run func(ctx context.Context, beatPath string, start int, end *int)) *MockBeatStorage_GetBeatFromS3_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int), args[3].(*int))
	})
	return _c
}

func (_c *MockBeatStorage_GetBeatFromS3_Call) Return(obj io.ReadCloser, size int, contentType string, err error) *MockBeatStorage_GetBeatFromS3_Call {
	_c.Call.Return(obj, size, contentType, err)
	return _c
}

func (_c *MockBeatStorage_GetBeatFromS3_Call) RunAndReturn(run func(context.Context, string, int, *int) (io.ReadCloser, int, string, error)) *MockBeatStorage_GetBeatFromS3_Call {
	_c.Call.Return(run)
	return _c
}

// GetBeatGenres provides a mock function with given fields: ctx, beatID
func (_m *MockBeatStorage) GetBeatGenres(ctx context.Context, beatID int) ([]core.BeatGenre, error) {
	ret := _m.Called(ctx, beatID)

	if len(ret) == 0 {
		panic("no return value specified for GetBeatGenres")
	}

	var r0 []core.BeatGenre
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]core.BeatGenre, error)); ok {
		return rf(ctx, beatID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []core.BeatGenre); ok {
		r0 = rf(ctx, beatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.BeatGenre)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, beatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBeatStorage_GetBeatGenres_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBeatGenres'
type MockBeatStorage_GetBeatGenres_Call struct {
	*mock.Call
}

// GetBeatGenres is a helper method to define mock.On call
//   - ctx context.Context
//   - beatID int
func (_e *MockBeatStorage_Expecter) GetBeatGenres(ctx interface{}, beatID interface{}) *MockBeatStorage_GetBeatGenres_Call {
	return &MockBeatStorage_GetBeatGenres_Call{Call: _e.mock.On("GetBeatGenres", ctx, beatID)}
}

func (_c *MockBeatStorage_GetBeatGenres_Call) Run(run func(ctx context.Context, beatID int)) *MockBeatStorage_GetBeatGenres_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockBeatStorage_GetBeatGenres_Call) Return(beatGenres []core.BeatGenre, err error) *MockBeatStorage_GetBeatGenres_Call {
	_c.Call.Return(beatGenres, err)
	return _c
}

func (_c *MockBeatStorage_GetBeatGenres_Call) RunAndReturn(run func(context.Context, int) ([]core.BeatGenre, error)) *MockBeatStorage_GetBeatGenres_Call {
	_c.Call.Return(run)
	return _c
}

// GetBeatsByBeatmakerID provides a mock function with given fields: ctx, beatmakerID, p
func (_m *MockBeatStorage) GetBeatsByBeatmakerID(ctx context.Context, beatmakerID int, p core.GetBeatsParams) ([]core.Beat, int, error) {
	ret := _m.Called(ctx, beatmakerID, p)

	if len(ret) == 0 {
		panic("no return value specified for GetBeatsByBeatmakerID")
	}

	var r0 []core.Beat
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, core.GetBeatsParams) ([]core.Beat, int, error)); ok {
		return rf(ctx, beatmakerID, p)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, core.GetBeatsParams) []core.Beat); ok {
		r0 = rf(ctx, beatmakerID, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, core.GetBeatsParams) int); ok {
		r1 = rf(ctx, beatmakerID, p)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, core.GetBeatsParams) error); ok {
		r2 = rf(ctx, beatmakerID, p)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockBeatStorage_GetBeatsByBeatmakerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBeatsByBeatmakerID'
type MockBeatStorage_GetBeatsByBeatmakerID_Call struct {
	*mock.Call
}

// GetBeatsByBeatmakerID is a helper method to define mock.On call
//   - ctx context.Context
//   - beatmakerID int
//   - p core.GetBeatsParams
func (_e *MockBeatStorage_Expecter) GetBeatsByBeatmakerID(ctx interface{}, beatmakerID interface{}, p interface{}) *MockBeatStorage_GetBeatsByBeatmakerID_Call {
	return &MockBeatStorage_GetBeatsByBeatmakerID_Call{Call: _e.mock.On("GetBeatsByBeatmakerID", ctx, beatmakerID, p)}
}

func (_c *MockBeatStorage_GetBeatsByBeatmakerID_Call) Run(run func(ctx context.Context, beatmakerID int, p core.GetBeatsParams)) *MockBeatStorage_GetBeatsByBeatmakerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int), args[2].(core.GetBeatsParams))
	})
	return _c
}

func (_c *MockBeatStorage_GetBeatsByBeatmakerID_Call) Return(beats []core.Beat, total int, err error) *MockBeatStorage_GetBeatsByBeatmakerID_Call {
	_c.Call.Return(beats, total, err)
	return _c
}

func (_c *MockBeatStorage_GetBeatsByBeatmakerID_Call) RunAndReturn(run func(context.Context, int, core.GetBeatsParams) ([]core.Beat, int, error)) *MockBeatStorage_GetBeatsByBeatmakerID_Call {
	_c.Call.Return(run)
	return _c
}

// GetPresignedURL provides a mock function with given fields: ctx, beatPath, expiry
func (_m *MockBeatStorage) GetPresignedURL(ctx context.Context, beatPath string, expiry time.Duration) (string, error) {
	ret := _m.Called(ctx, beatPath, expiry)

	if len(ret) == 0 {
		panic("no return value specified for GetPresignedURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) (string, error)); ok {
		return rf(ctx, beatPath, expiry)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) string); ok {
		r0 = rf(ctx, beatPath, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, beatPath, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBeatStorage_GetPresignedURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPresignedURL'
type MockBeatStorage_GetPresignedURL_Call struct {
	*mock.Call
}

// GetPresignedURL is a helper method to define mock.On call
//   - ctx context.Context
//   - beatPath string
//   - expiry time.Duration
func (_e *MockBeatStorage_Expecter) GetPresignedURL(ctx interface{}, beatPath interface{}, expiry interface{}) *MockBeatStorage_GetPresignedURL_Call {
	return &MockBeatStorage_GetPresignedURL_Call{Call: _e.mock.On("GetPresignedURL", ctx, beatPath, expiry)}
}

func (_c *MockBeatStorage_GetPresignedURL_Call) Run(run func(ctx context.Context, beatPath string, expiry time.Duration)) *MockBeatStorage_GetPresignedURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(time.Duration))
	})
	return _c
}

func (_c *MockBeatStorage_GetPresignedURL_Call) Return(url string, err error) *MockBeatStorage_GetPresignedURL_Call {
	_c.Call.Return(url, err)
	return _c
}

func (_c *MockBeatStorage_GetPresignedURL_Call) RunAndReturn(run func(context.Context, string, time.Duration) (string, error)) *MockBeatStorage_GetPresignedURL_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserSeenBeats provides a mock function with given fields: ctx, userID
func (_m *MockBeatStorage) GetUserSeenBeats(ctx context.Context, userID int) ([]string, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserSeenBeats")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]string, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []string); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBeatStorage_GetUserSeenBeats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserSeenBeats'
type MockBeatStorage_GetUserSeenBeats_Call struct {
	*mock.Call
}

// GetUserSeenBeats is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
func (_e *MockBeatStorage_Expecter) GetUserSeenBeats(ctx interface{}, userID interface{}) *MockBeatStorage_GetUserSeenBeats_Call {
	return &MockBeatStorage_GetUserSeenBeats_Call{Call: _e.mock.On("GetUserSeenBeats", ctx, userID)}
}

func (_c *MockBeatStorage_GetUserSeenBeats_Call) Run(run func(ctx context.Context, userID int)) *MockBeatStorage_GetUserSeenBeats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockBeatStorage_GetUserSeenBeats_Call) Return(_a0 []string, _a1 error) *MockBeatStorage_GetUserSeenBeats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBeatStorage_GetUserSeenBeats_Call) RunAndReturn(run func(context.Context, int) ([]string, error)) *MockBeatStorage_GetUserSeenBeats_Call {
	_c.Call.Return(run)
	return _c
}

// PopUserSeenBeat provides a mock function with given fields: ctx, userID
func (_m *MockBeatStorage) PopUserSeenBeat(ctx context.Context, userID int) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for PopUserSeenBeat")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBeatStorage_PopUserSeenBeat_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PopUserSeenBeat'
type MockBeatStorage_PopUserSeenBeat_Call struct {
	*mock.Call
}

// PopUserSeenBeat is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int
func (_e *MockBeatStorage_Expecter) PopUserSeenBeat(ctx interface{}, userID interface{}) *MockBeatStorage_PopUserSeenBeat_Call {
	return &MockBeatStorage_PopUserSeenBeat_Call{Call: _e.mock.On("PopUserSeenBeat", ctx, userID)}
}

func (_c *MockBeatStorage_PopUserSeenBeat_Call) Run(run func(ctx context.Context, userID int)) *MockBeatStorage_PopUserSeenBeat_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *MockBeatStorage_PopUserSeenBeat_Call) Return(_a0 error) *MockBeatStorage_PopUserSeenBeat_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBeatStorage_PopUserSeenBeat_Call) RunAndReturn(run func(context.Context, int) error) *MockBeatStorage_PopUserSeenBeat_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBeatStorage creates a new instance of MockBeatStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBeatStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBeatStorage {
	mock := &MockBeatStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
