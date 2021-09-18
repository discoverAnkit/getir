package repository

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type InMemoryTestClient struct {
	inMemoryClient *inMemoryClient
	t              *testing.T
}

func NewTestClient(t *testing.T) *InMemoryTestClient {

	cacheClient := cache.New(1*time.Hour,24*time.Hour)

	return &InMemoryTestClient{
		inMemoryClient : &inMemoryClient{
			cacheClient:cacheClient,
		},
		t:  t,
	}
}

//Note : Setting value for a key is a pre requisite before getting value for a key
//Thus, this test also covers SetKeyValue function
func TestInMemoryClient_GetValue_Success(t *testing.T) {
	// setup test client.
	c := NewTestClient(t)

	// given
	key := "ankit"
	value := "goyal"

	// when
	err := c.inMemoryClient.SetKeyValue(context.Background(),key,value)

	// then
	assert.NoError(t, err)

	// when
	valueFromDB := c.inMemoryClient.GetValue(context.Background(),key)

	// then
	assert.Equal(t,value,valueFromDB)
}

func TestInMemoryClient_GetValue_Failure(t *testing.T) {
	// setup test client.
	c := NewTestClient(t)

	// when
	valueFromDB := c.inMemoryClient.GetValue(context.Background(),"someKey")

	// then
	assert.Equal(t,valueFromDB, "")
}