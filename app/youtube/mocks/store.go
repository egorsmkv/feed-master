// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"sync"
	"time"

	ytfeed "github.com/umputun/feed-master/app/youtube/feed"
)

// StoreServiceMock is a mock implementation of youtube.StoreService.
//
// 	func TestSomethingThatUsesStoreService(t *testing.T) {
//
// 		// make and configure a mocked youtube.StoreService
// 		mockedStoreService := &StoreServiceMock{
// 			CheckProcessedFunc: func(entry ytfeed.Entry) (bool, time.Time, error) {
// 				panic("mock out the CheckProcessed method")
// 			},
// 			CountProcessedFunc: func() int {
// 				panic("mock out the CountProcessed method")
// 			},
// 			ExistFunc: func(entry ytfeed.Entry) (bool, error) {
// 				panic("mock out the Exist method")
// 			},
// 			LastFunc: func() (ytfeed.Entry, error) {
// 				panic("mock out the Last method")
// 			},
// 			LoadFunc: func(channelID string, max int) ([]ytfeed.Entry, error) {
// 				panic("mock out the Load method")
// 			},
// 			RemoveFunc: func(entry ytfeed.Entry) error {
// 				panic("mock out the Remove method")
// 			},
// 			RemoveOldFunc: func(channelID string, keep int) ([]string, error) {
// 				panic("mock out the RemoveOld method")
// 			},
// 			ResetProcessedFunc: func(entry ytfeed.Entry) error {
// 				panic("mock out the ResetProcessed method")
// 			},
// 			SaveFunc: func(entry ytfeed.Entry) (bool, error) {
// 				panic("mock out the Save method")
// 			},
// 			SetProcessedFunc: func(entry ytfeed.Entry) error {
// 				panic("mock out the SetProcessed method")
// 			},
// 		}
//
// 		// use mockedStoreService in code that requires youtube.StoreService
// 		// and then make assertions.
//
// 	}
type StoreServiceMock struct {
	// CheckProcessedFunc mocks the CheckProcessed method.
	CheckProcessedFunc func(entry ytfeed.Entry) (bool, time.Time, error)

	// CountProcessedFunc mocks the CountProcessed method.
	CountProcessedFunc func() int

	// ExistFunc mocks the Exist method.
	ExistFunc func(entry ytfeed.Entry) (bool, error)

	// LastFunc mocks the Last method.
	LastFunc func() (ytfeed.Entry, error)

	// LoadFunc mocks the Load method.
	LoadFunc func(channelID string, max int) ([]ytfeed.Entry, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(entry ytfeed.Entry) error

	// RemoveOldFunc mocks the RemoveOld method.
	RemoveOldFunc func(channelID string, keep int) ([]string, error)

	// ResetProcessedFunc mocks the ResetProcessed method.
	ResetProcessedFunc func(entry ytfeed.Entry) error

	// SaveFunc mocks the Save method.
	SaveFunc func(entry ytfeed.Entry) (bool, error)

	// SetProcessedFunc mocks the SetProcessed method.
	SetProcessedFunc func(entry ytfeed.Entry) error

	// calls tracks calls to the methods.
	calls struct {
		// CheckProcessed holds details about calls to the CheckProcessed method.
		CheckProcessed []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
		// CountProcessed holds details about calls to the CountProcessed method.
		CountProcessed []struct {
		}
		// Exist holds details about calls to the Exist method.
		Exist []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
		// Last holds details about calls to the Last method.
		Last []struct {
		}
		// Load holds details about calls to the Load method.
		Load []struct {
			// ChannelID is the channelID argument value.
			ChannelID string
			// Max is the max argument value.
			Max int
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
		// RemoveOld holds details about calls to the RemoveOld method.
		RemoveOld []struct {
			// ChannelID is the channelID argument value.
			ChannelID string
			// Keep is the keep argument value.
			Keep int
		}
		// ResetProcessed holds details about calls to the ResetProcessed method.
		ResetProcessed []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
		// SetProcessed holds details about calls to the SetProcessed method.
		SetProcessed []struct {
			// Entry is the entry argument value.
			Entry ytfeed.Entry
		}
	}
	lockCheckProcessed sync.RWMutex
	lockCountProcessed sync.RWMutex
	lockExist          sync.RWMutex
	lockLast           sync.RWMutex
	lockLoad           sync.RWMutex
	lockRemove         sync.RWMutex
	lockRemoveOld      sync.RWMutex
	lockResetProcessed sync.RWMutex
	lockSave           sync.RWMutex
	lockSetProcessed   sync.RWMutex
}

// CheckProcessed calls CheckProcessedFunc.
func (mock *StoreServiceMock) CheckProcessed(entry ytfeed.Entry) (bool, time.Time, error) {
	if mock.CheckProcessedFunc == nil {
		panic("StoreServiceMock.CheckProcessedFunc: method is nil but StoreService.CheckProcessed was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockCheckProcessed.Lock()
	mock.calls.CheckProcessed = append(mock.calls.CheckProcessed, callInfo)
	mock.lockCheckProcessed.Unlock()
	return mock.CheckProcessedFunc(entry)
}

// CheckProcessedCalls gets all the calls that were made to CheckProcessed.
// Check the length with:
//     len(mockedStoreService.CheckProcessedCalls())
func (mock *StoreServiceMock) CheckProcessedCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockCheckProcessed.RLock()
	calls = mock.calls.CheckProcessed
	mock.lockCheckProcessed.RUnlock()
	return calls
}

// CountProcessed calls CountProcessedFunc.
func (mock *StoreServiceMock) CountProcessed() int {
	if mock.CountProcessedFunc == nil {
		panic("StoreServiceMock.CountProcessedFunc: method is nil but StoreService.CountProcessed was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCountProcessed.Lock()
	mock.calls.CountProcessed = append(mock.calls.CountProcessed, callInfo)
	mock.lockCountProcessed.Unlock()
	return mock.CountProcessedFunc()
}

// CountProcessedCalls gets all the calls that were made to CountProcessed.
// Check the length with:
//     len(mockedStoreService.CountProcessedCalls())
func (mock *StoreServiceMock) CountProcessedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCountProcessed.RLock()
	calls = mock.calls.CountProcessed
	mock.lockCountProcessed.RUnlock()
	return calls
}

// Exist calls ExistFunc.
func (mock *StoreServiceMock) Exist(entry ytfeed.Entry) (bool, error) {
	if mock.ExistFunc == nil {
		panic("StoreServiceMock.ExistFunc: method is nil but StoreService.Exist was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockExist.Lock()
	mock.calls.Exist = append(mock.calls.Exist, callInfo)
	mock.lockExist.Unlock()
	return mock.ExistFunc(entry)
}

// ExistCalls gets all the calls that were made to Exist.
// Check the length with:
//     len(mockedStoreService.ExistCalls())
func (mock *StoreServiceMock) ExistCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockExist.RLock()
	calls = mock.calls.Exist
	mock.lockExist.RUnlock()
	return calls
}

// Last calls LastFunc.
func (mock *StoreServiceMock) Last() (ytfeed.Entry, error) {
	if mock.LastFunc == nil {
		panic("StoreServiceMock.LastFunc: method is nil but StoreService.Last was just called")
	}
	callInfo := struct {
	}{}
	mock.lockLast.Lock()
	mock.calls.Last = append(mock.calls.Last, callInfo)
	mock.lockLast.Unlock()
	return mock.LastFunc()
}

// LastCalls gets all the calls that were made to Last.
// Check the length with:
//     len(mockedStoreService.LastCalls())
func (mock *StoreServiceMock) LastCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockLast.RLock()
	calls = mock.calls.Last
	mock.lockLast.RUnlock()
	return calls
}

// Load calls LoadFunc.
func (mock *StoreServiceMock) Load(channelID string, max int) ([]ytfeed.Entry, error) {
	if mock.LoadFunc == nil {
		panic("StoreServiceMock.LoadFunc: method is nil but StoreService.Load was just called")
	}
	callInfo := struct {
		ChannelID string
		Max       int
	}{
		ChannelID: channelID,
		Max:       max,
	}
	mock.lockLoad.Lock()
	mock.calls.Load = append(mock.calls.Load, callInfo)
	mock.lockLoad.Unlock()
	return mock.LoadFunc(channelID, max)
}

// LoadCalls gets all the calls that were made to Load.
// Check the length with:
//     len(mockedStoreService.LoadCalls())
func (mock *StoreServiceMock) LoadCalls() []struct {
	ChannelID string
	Max       int
} {
	var calls []struct {
		ChannelID string
		Max       int
	}
	mock.lockLoad.RLock()
	calls = mock.calls.Load
	mock.lockLoad.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *StoreServiceMock) Remove(entry ytfeed.Entry) error {
	if mock.RemoveFunc == nil {
		panic("StoreServiceMock.RemoveFunc: method is nil but StoreService.Remove was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(entry)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedStoreService.RemoveCalls())
func (mock *StoreServiceMock) RemoveCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// RemoveOld calls RemoveOldFunc.
func (mock *StoreServiceMock) RemoveOld(channelID string, keep int) ([]string, error) {
	if mock.RemoveOldFunc == nil {
		panic("StoreServiceMock.RemoveOldFunc: method is nil but StoreService.RemoveOld was just called")
	}
	callInfo := struct {
		ChannelID string
		Keep      int
	}{
		ChannelID: channelID,
		Keep:      keep,
	}
	mock.lockRemoveOld.Lock()
	mock.calls.RemoveOld = append(mock.calls.RemoveOld, callInfo)
	mock.lockRemoveOld.Unlock()
	return mock.RemoveOldFunc(channelID, keep)
}

// RemoveOldCalls gets all the calls that were made to RemoveOld.
// Check the length with:
//     len(mockedStoreService.RemoveOldCalls())
func (mock *StoreServiceMock) RemoveOldCalls() []struct {
	ChannelID string
	Keep      int
} {
	var calls []struct {
		ChannelID string
		Keep      int
	}
	mock.lockRemoveOld.RLock()
	calls = mock.calls.RemoveOld
	mock.lockRemoveOld.RUnlock()
	return calls
}

// ResetProcessed calls ResetProcessedFunc.
func (mock *StoreServiceMock) ResetProcessed(entry ytfeed.Entry) error {
	if mock.ResetProcessedFunc == nil {
		panic("StoreServiceMock.ResetProcessedFunc: method is nil but StoreService.ResetProcessed was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockResetProcessed.Lock()
	mock.calls.ResetProcessed = append(mock.calls.ResetProcessed, callInfo)
	mock.lockResetProcessed.Unlock()
	return mock.ResetProcessedFunc(entry)
}

// ResetProcessedCalls gets all the calls that were made to ResetProcessed.
// Check the length with:
//     len(mockedStoreService.ResetProcessedCalls())
func (mock *StoreServiceMock) ResetProcessedCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockResetProcessed.RLock()
	calls = mock.calls.ResetProcessed
	mock.lockResetProcessed.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *StoreServiceMock) Save(entry ytfeed.Entry) (bool, error) {
	if mock.SaveFunc == nil {
		panic("StoreServiceMock.SaveFunc: method is nil but StoreService.Save was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(entry)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedStoreService.SaveCalls())
func (mock *StoreServiceMock) SaveCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}

// SetProcessed calls SetProcessedFunc.
func (mock *StoreServiceMock) SetProcessed(entry ytfeed.Entry) error {
	if mock.SetProcessedFunc == nil {
		panic("StoreServiceMock.SetProcessedFunc: method is nil but StoreService.SetProcessed was just called")
	}
	callInfo := struct {
		Entry ytfeed.Entry
	}{
		Entry: entry,
	}
	mock.lockSetProcessed.Lock()
	mock.calls.SetProcessed = append(mock.calls.SetProcessed, callInfo)
	mock.lockSetProcessed.Unlock()
	return mock.SetProcessedFunc(entry)
}

// SetProcessedCalls gets all the calls that were made to SetProcessed.
// Check the length with:
//     len(mockedStoreService.SetProcessedCalls())
func (mock *StoreServiceMock) SetProcessedCalls() []struct {
	Entry ytfeed.Entry
} {
	var calls []struct {
		Entry ytfeed.Entry
	}
	mock.lockSetProcessed.RLock()
	calls = mock.calls.SetProcessed
	mock.lockSetProcessed.RUnlock()
	return calls
}
