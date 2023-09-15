package mod

import (
	entity "github.com/mingmingshiliyu/fastflow/pkg/entity"
	mock "github.com/stretchr/testify/mock"
)

// MockCloser is an autogenerated mock type for the Closer type
type MockCloser struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockCloser) Close() {
	_m.Called()
}

// MockParser is an autogenerated mock type for the Parser type
type MockParser struct {
	mock.Mock
}

// EntryTaskIns provides a mock function with given fields: taskIns
func (_m *MockParser) EntryTaskIns(taskIns *entity.TaskInstance) {
	_m.Called(taskIns)
}

// InitialDagIns provides a mock function with given fields: dagIns
func (_m *MockParser) InitialDagIns(dagIns *entity.DagInstance) {
	_m.Called(dagIns)
}

// MockExecutor is an autogenerated mock type for the Executor type
type MockExecutor struct {
	mock.Mock
}

// CancelTaskIns provides a mock function with given fields: taskInsId
func (_m *MockExecutor) CancelTaskIns(taskInsId []string) error {
	ret := _m.Called(taskInsId)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(taskInsId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Push provides a mock function with given fields: dagIns, taskIns
func (_m *MockExecutor) Push(data *entity.DagInstance, taskIns *entity.TaskInstance) {
	_m.Called(data, taskIns)
}

// MockKeeper is an autogenerated mock type for the Keeper type
type MockKeeper struct {
	mock.Mock
}

// AliveNodes provides a mock function with given fields:
func (_m *MockKeeper) AliveNodes() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockKeeper) IsAlive(workerKey string) (bool, error) {
	ret := _m.Called(workerKey)
	return ret.Bool(0), ret.Error(1)
}

// Close provides a mock function with given fields:
func (_m *MockKeeper) Close() {
	_m.Called()
}

// IsLeader provides a mock function with given fields:
func (_m *MockKeeper) IsLeader() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// WorkerKey provides a mock function with given fields:
func (_m *MockKeeper) WorkerKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WorkerNumber provides a mock function with given fields:
func (_m *MockKeeper) WorkerNumber() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// NewMutex provides a mock function with given fields: key
func (_m *MockKeeper) NewMutex(key string) DistributedMutex {
	ret := _m.Called(key)

	var r0 DistributedMutex
	if rf, ok := ret.Get(0).(func(string) DistributedMutex); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DistributedMutex)
		}
	}

	return r0
}

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

// BatchCreatTaskIns provides a mock function with given fields: taskIns
func (_m *MockStore) BatchCreatTaskIns(taskIns []*entity.TaskInstance) error {
	ret := _m.Called(taskIns)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*entity.TaskInstance) error); ok {
		r0 = rf(taskIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchUpdateDagIns provides a mock function with given fields: dagIns
func (_m *MockStore) BatchUpdateDagIns(dagIns []*entity.DagInstance) error {
	ret := _m.Called(dagIns)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*entity.DagInstance) error); ok {
		r0 = rf(dagIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchUpdateTaskIns provides a mock function with given fields: taskIns
func (_m *MockStore) BatchUpdateTaskIns(taskIns []*entity.TaskInstance) error {
	ret := _m.Called(taskIns)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*entity.TaskInstance) error); ok {
		r0 = rf(taskIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *MockStore) Close() {
	_m.Called()
}

// CreateDag provides a mock function with given fields: dag
func (_m *MockStore) CreateDag(dag *entity.Dag) error {
	ret := _m.Called(dag)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Dag) error); ok {
		r0 = rf(dag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateDagIns provides a mock function with given fields: dagIns
func (_m *MockStore) CreateDagIns(dagIns *entity.DagInstance) error {
	ret := _m.Called(dagIns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.DagInstance) error); ok {
		r0 = rf(dagIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDag provides a mock function with given fields: dagId
func (_m *MockStore) GetDag(dagId string) (*entity.Dag, error) {
	ret := _m.Called(dagId)

	var r0 *entity.Dag
	if rf, ok := ret.Get(0).(func(string) *entity.Dag); ok {
		r0 = rf(dagId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dagId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDagInstance provides a mock function with given fields: dagInsId
func (_m *MockStore) GetDagInstance(dagInsId string) (*entity.DagInstance, error) {
	ret := _m.Called(dagInsId)

	var r0 *entity.DagInstance
	if rf, ok := ret.Get(0).(func(string) *entity.DagInstance); ok {
		r0 = rf(dagInsId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.DagInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dagInsId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskIns provides a mock function with given fields: taskIns
func (_m *MockStore) GetTaskIns(taskIns string) (*entity.TaskInstance, error) {
	ret := _m.Called(taskIns)

	var r0 *entity.TaskInstance
	if rf, ok := ret.Get(0).(func(string) *entity.TaskInstance); ok {
		r0 = rf(taskIns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TaskInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(taskIns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDagInstance provides a mock function with given fields: input
func (_m *MockStore) ListDagInstance(input *ListDagInstanceInput) ([]*entity.DagInstance, error) {
	ret := _m.Called(input)

	var r0 []*entity.DagInstance
	if rf, ok := ret.Get(0).(func(*ListDagInstanceInput) []*entity.DagInstance); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.DagInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ListDagInstanceInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTaskInstance provides a mock function with given fields: input
func (_m *MockStore) ListTaskInstance(input *ListTaskInstanceInput) ([]*entity.TaskInstance, error) {
	ret := _m.Called(input)

	var r0 []*entity.TaskInstance
	if rf, ok := ret.Get(0).(func(*ListTaskInstanceInput) []*entity.TaskInstance); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.TaskInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ListTaskInstanceInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PatchTaskIns provides a mock function with given fields: taskIns
func (_m *MockStore) PatchTaskIns(taskIns *entity.TaskInstance) error {
	ret := _m.Called(taskIns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.TaskInstance) error); ok {
		r0 = rf(taskIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PatchDagIns provides a mock function with given fields: dagIns, mustsPatchFields
func (_m *MockStore) PatchDagIns(dagIns *entity.DagInstance, mustsPatchFields ...string) error {
	_va := make([]interface{}, len(mustsPatchFields))
	for _i := range mustsPatchFields {
		_va[_i] = mustsPatchFields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, dagIns)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.DagInstance, ...string) error); ok {
		r0 = rf(dagIns, mustsPatchFields...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDagIns provides a mock function with given fields: dagIns
func (_m *MockStore) UpdateDag(dag *entity.Dag) error {
	ret := _m.Called(dag)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Dag) error); ok {
		r0 = rf(dag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDagIns provides a mock function with given fields: dagIns
func (_m *MockStore) UpdateDagIns(dagIns *entity.DagInstance) error {
	ret := _m.Called(dagIns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.DagInstance) error); ok {
		r0 = rf(dagIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTaskIns provides a mock function with given fields: taskIns
func (_m *MockStore) UpdateTaskIns(taskIns *entity.TaskInstance) error {
	ret := _m.Called(taskIns)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.TaskInstance) error); ok {
		r0 = rf(taskIns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockStore) Marshal(obj interface{}) ([]byte, error) {
	ret := _m.Called(obj)
	return ret.Get(0).([]byte), ret.Error(1)
}
func (_m *MockStore) Unmarshal(bytes []byte, ptr interface{}) error {
	ret := _m.Called(bytes, ptr)
	return ret.Error(0)
}

type MockTaskInfoGetter struct {
	ID     string
	Depend []string
	Status entity.TaskInstanceStatus
}

// GetDepend provides a mock function with given fields:
func (_m *MockTaskInfoGetter) GetDepend() []string {
	return _m.Depend
}

// GetGraphID provides a mock function with given fields:
func (_m *MockTaskInfoGetter) GetGraphID() string {
	return _m.ID
}

// GetID provides a mock function with given fields:
func (_m *MockTaskInfoGetter) GetID() string {
	return _m.ID
}

// GetStatus provides a mock function with given fields:
func (_m *MockTaskInfoGetter) GetStatus() entity.TaskInstanceStatus {
	return _m.Status
}
