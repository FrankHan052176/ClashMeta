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
	//bytes := []byte(paramsString)
	go func() {
		//bridge.SendToPort(i, handleUpdateConfig(bytes))
	}()
	return paramsString
}

func getProxies(env js.Env, this js.Value, args []js.Value) any {
	return handleGetProxies()
}

func changeProxy(env js.Env, this js.Value, args []js.Value) any {
	paramsString, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	handleChangeProxy(paramsString, func(value string) {
		//bridge.SendToPort(i, value)
	})
	return nil
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
	handleAsyncTestDelay(paramsString, func(value string) {
		//bridge.SendToPort(i, value)
	})
	return nil
}
func getExternalProviders(env js.Env, this js.Value, args []js.Value) any {
	return handleGetExternalProviders()
}
func getExternalProvider(env js.Env, this js.Value, args []js.Value) any {
	externalProviderName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	return handleGetExternalProvider(externalProviderName)
}
func updateExternalProvider(env js.Env, this js.Value, args []js.Value) {
	providerName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	handleUpdateExternalProvider(providerName, func(value string) {
		//bridge.SendToPort(i, value)
	})
}
func sideLoadExternalProvider(env js.Env, this js.Value, args []js.Value) {
	providerName, _ := napi.GetValueStringUtf8(env.Env, args[0].Value)
	dataChar, _ := napi.GetValueStringUtf8(env.Env, args[1].Value)
	data := []byte(dataChar)
	handleSideLoadExternalProvider(providerName, data, func(value string) {
		//bridge.SendToPort(i, value)
	})
}

func init() {
	entry.Export("startTun", js.AsCallback(startTun))
	entry.Export("stopTun", js.AsCallback(stopTun))
	entry.Export("forceGc", js.AsCallback(forceGc))
	entry.Export("getTraffic", js.AsCallback(getTraffic))
	entry.Export("getTotalTraffic", js.AsCallback(getTotalTraffic))
	// entry.Export("resetTraffic", js.AsCallback(OnEvent))
	// entry.Export("asyncTestDelay", js.AsCallback(GetPeers))
	// entry.Export("getConnections", js.AsCallback(GetNodeStatus))
	// entry.Export("GetNetworkConfigs", js.AsCallback(GetNetworkConfigs))
	// entry.Export("TestPromise", js.AsCallback(GetPromiseResolve))
}
func main() {
}
