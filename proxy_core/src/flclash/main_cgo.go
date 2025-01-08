//go:build cgo && ohos

package main

import "C"
import (
	bridge "core/dart-bridge"

	napi "github.com/likuai2010/ohos-napi"
	"github.com/likuai2010/ohos-napi/entry"
	"github.com/likuai2010/ohos-napi/js"
)

func startTun(env js.Env, this js.Value, args []js.Value) any {
	tunFd, _ := napi.GetValueInt32(env.Env, args[0].Value)
	StartTUN(int(tunFd))
	return nil
}
func stopTun(env js.Env, this js.Value, args []js.Value) any {
	StopTun()
	return nil
}

func validateConfig(s *C.char, port C.longlong) {
	i := int64(port)
	bytes := []byte(C.GoString(s))
	go func() {
		bridge.SendToPort(i, handleValidateConfig(bytes))
	}()
}

func updateConfig(env js.Env, this js.Value, args []js.Value) any {
	paramsString, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	promise := env.NewPromise()
	bytes := []byte(paramsString)
	go func() {
		promise.Resolve(handleUpdateConfig(bytes))
	}()
	return promise
}

func getProxies(env js.Env, this js.Value, args []js.Value) any {
	return handleGetProxies()
}

func changeProxy(env js.Env, this js.Value, args []js.Value) any {
	paramsString, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	promise := env.NewPromise()
	handleChangeProxy(paramsString, func(value string) {
		promise.Resolve(value)
	})
	return promise
}

func getTraffic(env js.Env, this js.Value, args []js.Value) any {
	onlyProxy := true
	handleGetTraffic(onlyProxy)
	return handleGetTraffic(onlyProxy)
}
func getTotalTraffic(env js.Env, this js.Value, args []js.Value) any {
	onlyProxy := true
	return handleGetTotalTraffic(onlyProxy)
}
func resetTraffic(env js.Env, this js.Value, args []js.Value) any {
	handleResetTraffic()
	return nil
}
func forceGc(env js.Env, this js.Value, args []js.Value) any {
	handleForceGc()
	return nil
}
func asyncTestDelay(env js.Env, this js.Value, args []js.Value) any {
	paramsString, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	promise := env.NewPromise()
	handleAsyncTestDelay(paramsString, func(value string) {
		promise.Resolve(value)
	})
	return promise
}
func getExternalProviders(env js.Env, this js.Value, args []js.Value) any {
	return handleGetExternalProviders()
}
func getExternalProvider(env js.Env, this js.Value, args []js.Value) any {
	externalProviderName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	return handleGetExternalProvider(externalProviderName)
}
func updateGeoData(env js.Env, this js.Value, args []js.Value) any {
	geoType, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	geoName, _ := napi.GetValueStringUtf8(env.Env, args[1].Value)
	promise := env.NewPromise()
	handleUpdateGeoData(geoType, geoName, func(value string) {
		promise.Resolve(value)
	})
	return promise
}
func updateExternalProvider(env js.Env, this js.Value, args []js.Value) any {
	providerName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	promise := env.NewPromise()
	handleUpdateExternalProvider(providerName, func(value string) {
		promise.Resolve(value)
	})
	return promise
}

func sideLoadExternalProvider(env js.Env, this js.Value, args []js.Value) any {
	providerName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	dataChar, _ := napi.GetValueStringUtf8(env.Env, args[1].Value)
	data := []byte(dataChar)
	promise := env.NewPromise()
	handleSideLoadExternalProvider(providerName, data, func(value string) {
		promise.Resolve(value)
	})
	return promise
}
func getConnections(env js.Env, this js.Value, args []js.Value) any {
	return handleGetConnections()
}

func closeConnections(env js.Env, this js.Value, args []js.Value) any {
	return handleCloseConnections()
}

func closeConnection(env js.Env, this js.Value, args []js.Value) any {
	connectionId, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	return handleCloseConnection(connectionId)
}
func startLog(env js.Env, this js.Value, args []js.Value) any {
	tsfn := env.CreateThreadsafeFunction(args[0], "startLog")
	handleStartLog(func(value string) {
		tsfn.Call(env.ValueOf("startLog"), env.ValueOf(value))
	})
	return nil
}
func stopLog(env js.Env, this js.Value, args []js.Value) any {
	handleStopLog()
	return nil
}

func init() {
	entry.Export("startTun", js.AsCallback(startTun))
	entry.Export("stopTun", js.AsCallback(stopTun))
	entry.Export("forceGc", js.AsCallback(forceGc))
	entry.Export("getTraffic", js.AsCallback(getTraffic))
	entry.Export("getTotalTraffic", js.AsCallback(getTotalTraffic))
	entry.Export("resetTraffic", js.AsCallback(resetTraffic))
	entry.Export("asyncTestDelay", js.AsCallback(asyncTestDelay))
	entry.Export("getConnections", js.AsCallback(getConnections))
	entry.Export("updateExternalProvider", js.AsCallback(updateExternalProvider))
}
func main() {
}
