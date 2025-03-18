package ablmodels

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDevice(t *testing.T) {
	// Test with a standard device kind
	kind := "testDevice"
	device := NewDevice(kind)

	// Check all fields are correctly initialized
	assert.Equal(t, kind, device.Kind, "Kind should be set correctly")
	assert.Nil(t, device.PresetURI, "PresetURI should be nil")
	assert.Empty(t, device.Name, "Name should be empty")
	assert.Nil(t, device.Parameters, "Parameters should be nil")
	assert.Empty(t, device.Chains, "Chains should be empty")
	assert.Empty(t, device.ReturnChains, "ReturnChains should be empty")
}

func TestDeviceAddChain(t *testing.T) {
	device := NewDevice("testDevice")

	// Define a mock chain
	type mockChain struct {
		Name string
	}
	chain := mockChain{Name: "TestChain"}

	// Add the chain
	result := device.AddChain(chain)

	// Check the chain was added
	assert.Len(t, device.Chains, 1, "Chain should be added")

	// Check that the returned device is the same instance
	assert.Same(t, device, result, "AddChain should return the same device instance")

	// Check that the chain was correctly added
	addedChain, ok := device.Chains[0].(mockChain)
	assert.True(t, ok, "Added chain should be of expected type")
	assert.Equal(t, chain.Name, addedChain.Name, "Chain name should match")
}

func TestDeviceAddReturnChain(t *testing.T) {
	device := NewDevice("testDevice")

	// Define a mock return chain
	type mockChain struct {
		Name string
	}
	returnChain := mockChain{Name: "TestReturnChain"}

	// Add the return chain
	result := device.AddReturnChain(returnChain)

	// Check the return chain was added
	assert.Len(t, device.ReturnChains, 1, "Return chain should be added")

	// Check that the returned device is the same instance
	assert.Same(t, device, result, "AddReturnChain should return the same device instance")

	// Check that the return chain was correctly added
	addedChain, ok := device.ReturnChains[0].(mockChain)
	assert.True(t, ok, "Added return chain should be of expected type")
	assert.Equal(t, returnChain.Name, addedChain.Name, "Return chain name should match")
}

func TestDeviceWithParameters(t *testing.T) {
	device := NewDevice("testDevice")

	// Define mock parameters
	params := map[string]interface{}{
		"param1": 123,
		"param2": "value",
	}

	// Set the parameters
	result := device.WithParameters(params)

	// Check the parameters were set
	assert.True(t, reflect.DeepEqual(device.Parameters, params), "Parameters should be set correctly")

	// Check that the returned device is the same instance
	assert.Same(t, device, result, "WithParameters should return the same device instance")
}
