//+build tools

package main

//go:generate mockery -name=InMemoryClient -dir=repository/ -output=mocks
//go:generate mockery -name=MongoClient -dir=repository/ -output=mocks
