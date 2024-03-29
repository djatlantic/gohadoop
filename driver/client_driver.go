package main

import (
  "log"
  "github.com/gohadoop"
  "github.com/gohadoop/hadoop_common/ipc/client"
  "github.com/gohadoop/hadoop_common"
  "github.com/gohadoop/hadoop_yarn"
  "github.com/nu7hatch/gouuid"
)

func main() {
  var err error

  clientId, _ := uuid.NewV4()
  ugi, _ := gohadoop.CreateSimpleUGIProto()
  c := &ipc.Client{ClientId: clientId, Ugi: ugi, ServerAddress: "0.0.0.0:28081"}
  var clientProtocolVersion uint64 = 1
  var methodName string
  var protocolName string

  // ApplicationClientProtocol.getApplications 
  methodName = "getApplications"
  protocolName = "org.apache.hadoop.yarn.api.ApplicationClientProtocolPB"
  getAppsRpcProto := hadoop_common.RequestHeaderProto {MethodName: &methodName, DeclaringClassProtocolName: &protocolName, ClientProtocolVersion: &clientProtocolVersion}
  applicationStates := []hadoop_yarn.YarnApplicationStateProto{hadoop_yarn.YarnApplicationStateProto_ACCEPTED, hadoop_yarn.YarnApplicationStateProto_RUNNING, hadoop_yarn.YarnApplicationStateProto_SUBMITTED}
  getAppsReqProto := hadoop_yarn.GetApplicationsRequestProto {ApplicationStates: applicationStates}
  getAppsResProto := hadoop_yarn.GetApplicationsResponseProto{}
  log.Println("Calling rpc method: ", methodName) 
  err = c.Call(&getAppsRpcProto, &getAppsReqProto, &getAppsResProto)
  if err != nil {
    log.Fatal("Client.call failed", err)
  }
  log.Println("Returned response: ", getAppsResProto)

  // ApplicationClientProtocol.getNewApplication
  methodName = "getNewApplication"
  protocolName = "org.apache.hadoop.yarn.api.ApplicationClientProtocolPB"
  getNewAppRpcProto := hadoop_common.RequestHeaderProto {MethodName: &methodName, DeclaringClassProtocolName: &protocolName, ClientProtocolVersion: &clientProtocolVersion}
  getNewAppReqProto := hadoop_yarn.GetNewApplicationRequestProto {}
  getNewAppResProto := hadoop_yarn.GetNewApplicationResponseProto {}
  log.Println("Calling rpc method: ", methodName) 
  err = c.Call(&getNewAppRpcProto, &getNewAppReqProto, &getNewAppResProto)
  if err != nil {
    log.Fatal("Client.call failed", err)
  }
  log.Println("Returned response: ", getNewAppResProto)
}
