package ablmodels

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewDevice verifies that a new Device is created with the correct initial values
// and that all fields are properly initialized.
func TestNewDevice(t *testing.T) {
	kind := "testDevice"
	device := NewDevice(kind)

	assert.Equal(t, kind, device.Kind, "Kind should be set correctly")
	assert.Nil(t, device.PresetURI, "PresetURI should be nil")
	assert.Empty(t, device.Name, "Name should be empty")
	assert.Nil(t, device.Parameters, "Parameters should be nil")
	assert.Empty(t, device.Chains, "Chains should be empty")
	assert.Empty(t, device.ReturnChains, "ReturnChains should be empty")
}

// TestDeviceAddChain verifies that adding a chain to a device works correctly
// and that the chain is properly stored and retrievable.
func TestDeviceAddChain(t *testing.T) {
	device := NewDevice("testDevice")

	type mockChain struct {
		Name string
	}
	chain := mockChain{Name: "TestChain"}

	result := device.AddChain(chain)

	assert.Len(t, device.Chains, 1, "Chain should be added")
	assert.Same(t, device, result, "AddChain should return the same device instance")
	addedChain, ok := device.Chains[0].(mockChain)
	assert.True(t, ok, "Added chain should be of expected type")
	assert.Equal(t, chain.Name, addedChain.Name, "Chain name should match")
}

// TestDeviceAddReturnChain verifies that adding a return chain to a device works correctly
// and that the return chain is properly stored in the ReturnChains collection.
func TestDeviceAddReturnChain(t *testing.T) {
	device := NewDevice("testDevice")

	type mockChain struct {
		Name string
	}
	returnChain := mockChain{Name: "TestReturnChain"}

	result := device.AddReturnChain(returnChain)

	assert.Len(t, device.ReturnChains, 1, "Return chain should be added")
	assert.Same(t, device, result, "AddReturnChain should return the same device instance")
	addedChain, ok := device.ReturnChains[0].(mockChain)
	assert.True(t, ok, "Added return chain should be of expected type")
	assert.Equal(t, returnChain.Name, addedChain.Name, "Return chain name should match")
}

// TestDeviceWithParameters verifies that setting parameters on a device works correctly
// and that the device instance is returned for method chaining.
func TestDeviceWithParameters(t *testing.T) {
	device := NewDevice("testDevice")

	params := map[string]interface{}{
		"param1": 123,
		"param2": "value",
	}

	result := device.WithParameters(params)

	assert.True(t, reflect.DeepEqual(device.Parameters, params), "Parameters should be set correctly")
	assert.Same(t, device, result, "WithParameters should return the same device instance")
}
