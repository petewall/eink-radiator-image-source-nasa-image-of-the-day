// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"image"
	"io"
	"sync"

	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/internal"
)

type FakeImageEncoder struct {
	Stub        func(io.Writer, image.Image) error
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 io.Writer
		arg2 image.Image
	}
	returns struct {
		result1 error
	}
	returnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageEncoder) Spy(arg1 io.Writer, arg2 image.Image) error {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 io.Writer
		arg2 image.Image
	}{arg1, arg2})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("ImageEncoder", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return returns.result1
}

func (fake *FakeImageEncoder) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeImageEncoder) Calls(stub func(io.Writer, image.Image) error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeImageEncoder) ArgsForCall(i int) (io.Writer, image.Image) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *FakeImageEncoder) Returns(result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEncoder) ReturnsOnCall(i int, result1 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageEncoder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageEncoder) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ internal.ImageEncoder = new(FakeImageEncoder).Spy
