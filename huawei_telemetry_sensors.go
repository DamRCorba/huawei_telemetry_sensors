package huawei_sensorPath
import(
      "fmt"
      "reflect"
      "time"
      "strings"
      "unicode"
      "strconv"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-telemetry"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-bfd"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-bgp"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-devm"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-driver"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ifm"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-isis"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-mpls"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ospfv2"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ospfv3"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-qos"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-sem"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-telemEmdi"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-trafficmng"
     // "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huaweiV8R12-debug"
      
      "github.com/golang/protobuf/proto"
      "github.com/influxdata/telegraf/metric"
  )
/* Returns the protoMessage of the sensor path.
  Huawei have only a few sensors paths for metrics.
  The sensors could be known with the command. "display telemetry sensor-path "
  @params: path (string) - The head of the sensor path. Example: "huawei-ifm"
  @returns: sensor-path proto message
*/
func GetMessageType(path string) (proto.Message) {
  fmt.Println("HOLAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
  fmt.Println(path)
  sensorType := strings.Split(path,":")
  fmt.Println(sensorType)
  switch sensorType[0] {
  case "huawei-bfd":
      return &huawei_bfd.Bfd{}

  case "huawei-bgp":
    switch sensorType[1] {
    case "ESTABLISHED":
      return &huawei_bgp.ESTABLISHED{}
    case "BACKWARD":
      return &huawei_bgp.BACKWARD{}
    }
    return &huawei_bgp.ESTABLISHED{}

  case "huawei-devm":
    return &huawei_devm.Devm{}

  case "huawei-driver":
    switch sensorType[1] {
    case "hwEntityInvalid":
        return &huawei_driver.HwEntityInvalid{}
    case "hwEntityResume":
        return &huawei_driver.HwEntityResume{}
    case "hwOpticalInvalid":
      return &huawei_driver.HwOpticalInvalid{}
    case "hwOpticalInvalidResume":
      return &huawei_driver.HwOpticalInvalidResume{}
    }
    return &huawei_driver.HwEntityInvalid{}

  case "huawei-ifm":
    return &huawei_ifm.Ifm{}

  case "huawei-isis":
  case "huawei-isiscomm":
    return &huawei_isiscomm.IsisAdjacencyChange{}

  case "huawei-mpls":
    return &huawei_mpls.Mpls{}

  case "huawei-ospfv2":
    switch sensorType[1] {
    case "ospfNbrStateChange":
      return &huawei_ospfv2.OspfNbrStateChange{}
    case "ospfVirtNbrStateChange":
      return &huawei_ospfv2.OspfVirtNbrStateChange{}
    }
    return &huawei_ospfv2.OspfNbrStateChange{}

  case "huawei-ospfv3":
    return &huawei_ospfv3.Ospfv3NbrStateChange{}

  case "huawei-qos":
    return &huawei_qos.Qos{}

  case "huawei-sem":
    switch sensorType[1] {
    case "hwCPUUtilizationResume":
      return &huawei_sem.HwStorageUtilizationResume{}
    case "hwCPUUtilizationRisingAlarm":
      return &huawei_sem.HwCPUUtilizationRisingAlarm{}
    case "hwStorageUtilizationResume":
      return &huawei_sem.HwStorageUtilizationResume{}
    case "hwStorageUtilizationRisingAlarm":
      return &huawei_sem.HwStorageUtilizationRisingAlarm{}
      }
    return &huawei_sem.HwStorageUtilizationResume{}

  case "huawei-telmEmdi":
  case "huawei-emdi":
    return &huawei_telemEmdi.TelemEmdi{}

  case "huawei-trafficmng":
    return &huawei_trafficmng.Trafficmng{}

  default:
    fmt.Println("Error Sensor Desconocido")
    fmt.Println("HOLAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
    fmt.Println(sensorType)
    return &huawei_devm.Devm{}
  }
    return &huawei_devm.Devm{}
}



/*
  Get the types of the Telemetry EndPoint
  @Params: a string with the telemetry complete path
  @Returns: a Map with keys and types of the endpoint
*/
func GetTypeValue (path string) map[string]reflect.Type {
  resolve := make(map[string]reflect.Type)
  splited := strings.Split(path,":")
  switch splited[0] {
  case "huawei-bfd":

      return resolve

  case "huawei-bgp":
    switch splited[1] {
    case "ESTABLISHED":
      fooType := reflect.TypeOf(huawei_bgp.ESTABLISHED{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
    break;
    case "BACKWARD":
      fooType := reflect.TypeOf(huawei_bgp.BACKWARD{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
    break;

    }
    return resolve

  case "huawei-devm":
    switch splited[1] {
    case "devm/cpuInfos/cpuInfo":
          fooType := reflect.TypeOf(huawei_devm.Devm_CpuInfos_CpuInfo{})
          for i := 0; i < fooType.NumField(); i ++ {
            keys := fooType.Field(i)
            resolve[LcFirst(keys.Name)] = keys.Type
            }
        break;
    case "devm/fans/fan":
      fooType := reflect.TypeOf(huawei_devm.Devm_Fans_Fan{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "devm/memoryInfos/memoryInfo":
      fooType := reflect.TypeOf(huawei_devm.Devm_MemoryInfos_MemoryInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "devm/ports/port":
      fooType := reflect.TypeOf(huawei_devm.Devm_Ports_Port{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "devm/ports/port/opticalInfo":
      fooType := reflect.TypeOf(huawei_devm.Devm_Ports_Port{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      fooType = reflect.TypeOf(huawei_devm.Devm_Ports_Port_OpticalInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "RxPower" || keys.Name == "TxPower" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf(1.0)
          } else {
              resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;
    case "devm/powerSupplys/powerSupply/powerEnvironments/powerEnvironment":
      fooType := reflect.TypeOf(huawei_devm.Devm_PowerSupplys_PowerSupply_PowerEnvironments_PowerEnvironment{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "PowerValue" || keys.Name == "VoltageValue" {
            resolve[LcFirst(keys.Name)] = reflect.TypeOf(1.0)
        } else {
            if keys.Name == "PemIndex" {
              resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
            } else {
                resolve[LcFirst(keys.Name)] = keys.Type
            }
          }
        }
        break;
    case "devm/temperatureInfos/temperatureInfo":
      fooType := reflect.TypeOf(huawei_devm.Devm_TemperatureInfos_TemperatureInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "I2c" || keys.Name == "Channel" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
        } else {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;
    }

    return resolve

  case "huawei-driver":
    switch splited[1] {
    case "hwEntityInvalid":
      fooType := reflect.TypeOf(huawei_driver.HwEntityInvalid{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "I2c" || keys.Name == "Channel" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
        } else {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;
    case "hwEntityResume":
      fooType := reflect.TypeOf(huawei_driver.HwEntityResume{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "I2c" || keys.Name == "Channel" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
        } else {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;
    case "hwOpticalInvalid":
      fooType := reflect.TypeOf(huawei_driver.HwOpticalInvalid{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "I2c" || keys.Name == "Channel" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
        } else {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;
    case "hwOpticalInvalidResume":
      fooType := reflect.TypeOf(huawei_driver.HwOpticalInvalidResume{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        if keys.Name == "I2c" || keys.Name == "Channel" {
          resolve[LcFirst(keys.Name)] = reflect.TypeOf("")
        } else {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
        }
        break;

    }
    return resolve

  case "huawei-ifm":
    fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface{})
    for i := 0; i < fooType.NumField(); i ++ {
      keys := fooType.Field(i)
      if keys.Name == "IfName" {
          resolve[LcFirst(keys.Name)] = keys.Type
          }
      }
    switch splited[1] {
    case "ifm/interfaces/interface": // No trae data mas que IfIndex, IfName e IfAdminStatus_UP si la interface esta Down no devuevle el campo.
      fooType = reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "ifm/interfaces/interface/ifClearedStat":
      fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface_IfClearedStat{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "ifm/interfaces/interface/ifDynamicInfo":
      fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface_IfDynamicInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "ifm/interfaces/interface/ifStatistics":
      fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface_IfStatistics{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "ifm/interfaces/interface/ifStatistics/ethPortErrSts":
      fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface_IfStatistics_EthPortErrSts{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    }
    return resolve
  case "huawei-isiscomm":
  case "huawei-isis":
    fooType := reflect.TypeOf(huawei_isiscomm.IsisAdjacencyChange{})
    for i := 0; i < fooType.NumField(); i ++ {
      keys := fooType.Field(i)
      resolve[LcFirst(keys.Name)] = keys.Type
      }
    return resolve

  case "huawei-mpls":
    fooType := reflect.TypeOf(huawei_mpls.Mpls{})
    for i := 0; i < fooType.NumField(); i ++ {
      keys := fooType.Field(i)
      resolve[LcFirst(keys.Name)] = keys.Type
      }
    return resolve

  case "huawei-ospfv2":
    switch splited[1] {
    case "ospfNbrStateChange":
      fooType := reflect.TypeOf(huawei_ospfv2.OspfNbrStateChange{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      break;
    case "ospfVirtNbrStateChange":
      fooType := reflect.TypeOf(huawei_ospfv2.OspfVirtNbrStateChange{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      break;
    }
    return resolve

  case "huawei-ospfv3":
    fooType := reflect.TypeOf(huawei_ospfv3.Ospfv3NbrStateChange{})
    for i := 0; i < fooType.NumField(); i ++ {
      keys := fooType.Field(i)
      resolve[LcFirst(keys.Name)] = keys.Type
      }
    return resolve

  case "huawei-qos":
    switch splited[1] {
    case "qos/qosBuffers/qosBuffer":
      fooType := reflect.TypeOf(huawei_qos.Qos_QosBuffers_QosBuffer{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "qos/qosIfQoss/qosIfQos/qosPolicyApplys/qosPolicyApply/qosPolicyStats/qosPolicyStat/qosRuleStats/qosRuleStat":
      fooType := reflect.TypeOf(huawei_qos.Qos_QosIfQoss_QosIfQos_QosPolicyApplys_QosPolicyApply_QosPolicyStats_QosPolicyStat_QosRuleStats_QosRuleStat{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "qos/qosPortQueueStatInfos/qosPortQueueStatInfo":
      fooType := reflect.TypeOf(huawei_qos.Qos_QosPortQueueStatInfos_QosPortQueueStatInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    }
    return resolve

  case "huawei-sem":
    switch splited[1] {
    case "hwCPUUtilizationResume":
      fooType := reflect.TypeOf(huawei_sem.HwStorageUtilizationResume{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "hwCPUUtilizationRisingAlarm":
      fooType := reflect.TypeOf(huawei_sem.HwCPUUtilizationRisingAlarm{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "hwStorageUtilizationResume":
      fooType := reflect.TypeOf(huawei_sem.HwStorageUtilizationResume{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "hwStorageUtilizationRisingAlarm":
      fooType := reflect.TypeOf(huawei_sem.HwStorageUtilizationRisingAlarm{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
      }
    return resolve

  case "huawei-telmEmdi":
  case "huawei-emdi":
    switch splited[1] {
    case "emdi/emdiTelemReps/emdiTelemRep":
      fooType := reflect.TypeOf(huawei_telemEmdi.TelemEmdi_EmdiTelemReps_EmdiTelemRep{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      break;
    case "emdi/emdiTelemRtps/emdiTelemRtp":
      fooType := reflect.TypeOf(huawei_telemEmdi.TelemEmdi_EmdiTelemRtps_EmdiTelemRtp{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      break;
    }
    return resolve

  case "huawei-trafficmng":
    switch splited[1] {
    case "trafficmng/tmSlotSFUs/tmSlotSFU/sfuStatisticss/sfuStatistics":
      fooType := reflect.TypeOf(huawei_trafficmng.Trafficmng{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      fooType = reflect.TypeOf(huawei_trafficmng.Trafficmng_TmSlotSFUs{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      fooType = reflect.TypeOf(huawei_trafficmng.Trafficmng_TmSlotSFUs_TmSlotSFU{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      fooType = reflect.TypeOf(huawei_trafficmng.Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      fooType = reflect.TypeOf(huawei_trafficmng.Trafficmng_TmSlotSFUs_TmSlotSFU_SfuStatisticss_SfuStatistics{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
      break;

    }
    return resolve
    // V8R12 Sensors
  case "huawei-debug":
    switch splited[1] {
    case "debug/cpu-infos/cpu-info":
          fooType := reflect.TypeOf(huaweiV8R12_debug.Debug_CpuInfos_CpuInfo{})
          for i := 0; i < fooType.NumField(); i ++ {
            keys := fooType.Field(i)
            resolve[LcFirst(keys.Name)] = keys.Type
            }
        break;
        }
    return resolve
  default:
    fmt.Println("Error Sensor Desconocido")
    return resolve
  }
 return resolve
}

/*
  Change the firts character of a string to Lowercase
*/
func LcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToLower(v)) + str[i+1:]
    }
    return ""
}

/*
  Change the firts character of a string to Uppercase
*/
func UcFirst(str string) string {
    for i, v := range str {
        return string(unicode.ToUpper(v)) + str[i+1:]
    }
    return ""
}

/*
  Append to the tags the telemetry values for position.
  @params:
  k - Key to evaluate
  v - Content of the Key
  tags - Global tags of the metric
  path - Telemetry path
  @returns
  original tag append the key if its a name Key.

*/
func AppendTags(k string, v string, tags map[string]string, path string) map[string]string {
  resolve := tags
  endPointTypes:=GetTypeValue(path)
  if endPointTypes[k] != nil {
    if reflect.TypeOf(decodeVal(endPointTypes[k], v)) == reflect.TypeOf("") {
      if k != "ifAdminStatus" {
          resolve[k] = v
      }
    }
  } else {
    if k == "ifName" || k == "position" || k == "pemIndex" || k == "i2c"{
      resolve[k] = v
    }

  }
  return resolve
}

/*
  Convert the telemetry Data to its type.
  @Params:
  tipo - telemetry path data type
  val - string value
  Returns the converted value
*/
func decodeVal(tipo interface{}, val string) interface{} {
  if tipo == nil {
    return val
  } else {
  value := reflect.New(tipo.(reflect.Type)).Elem().Interface()
  switch value.(type) {
  case uint32: resolve, _ := strconv.ParseUint(val,10,32); return resolve;
  case uint64: resolve,_ :=  strconv.ParseUint(val,10,64); return resolve;
  case int32: resolve,_ :=  strconv.ParseInt(val,10,32);   return resolve;
  case int64: resolve,_ :=  strconv.ParseInt(val,10,64);   return resolve;
  case float64: resolve, err :=  strconv.ParseFloat(val,64);
                if err != nil {
                  name:= strings.Replace(val,"\"","",-1)
                  resolve, _=  strconv.ParseFloat(name,64);
                }
                return resolve;
  case bool: resolve,_ :=  strconv.ParseBool(val); return resolve;
  }
  }
  resolve := val;
  return resolve;
}

/*
  Creates and add metrics from json mapped data in telegraf metrics SeriesGrouper
  @params:
  - grouper (*metric.SeriesGrouper) - pointer of metric series to append data.
  - tags (map[string]string) json data mapped
  - timestamp (time.Time) -
  - path (string) - sensor path
  - subfield (string) - subkey data.
    vals (string) - subkey content

*/
func CreateMetrics(grouper *metric.SeriesGrouper, tags map[string]string, timestamp time.Time, path string, subfield string, vals string)  {
  if subfield == "ifAdminStatus" {
    name:= strings.Replace(subfield,"\"","",-1)
    if vals == "IfAdminStatus_UP" {
      grouper.Add(path, tags, timestamp, string(name), 1)
    } else {
      grouper.Add(path, tags, timestamp, string(name), 0)
    }
  }
  if subfield == "ifOperStatus" {
    name:= strings.Replace(subfield,"\"","",-1)
    if vals == "IfOperStatus_UP" {
      grouper.Add(path, tags, timestamp, string(name), 1)
    } else {
      grouper.Add(path, tags, timestamp, string(name), 0)
    }
  }
  if vals != "" && subfield != "ifName" && subfield != "position" && subfield != "pemIndex" && subfield != "address" && subfield != "i2c" && subfield != "channel" &&
  subfield != "queueType" && subfield != "ifAdminStatus" && subfield != "ifOperStatus" {
    name:= strings.Replace(subfield,"\"","",-1)
    endPointTypes:=GetTypeValue(path)
    grouper.Add(path, tags, timestamp, string(name), decodeVal(endPointTypes[name], vals))
  }
}

/*
  Search de keys and vals of the data row in telemetry message.
  @params:
  - Message (*TelemetryRowGPB) - data buffer GPB of sensor data
  - sensorType (string) - sensor-path group.
  @returns:
  - keys (string) - Keys of the fields
  - vals (string) - Vals of the fields
*/
func SearchKey(Message *telemetry.TelemetryRowGPB, path string)  ([]string, []string){
  sensorType := strings.Split(path,":")[0]
  sensorMsg := GetMessageType(sensorType)
  err := proto.Unmarshal(Message.Content, sensorMsg)
  if (err != nil) {
    panic(err)
  }
  primero := reflect.ValueOf(sensorMsg).Interface()

  str := fmt.Sprintf("%v", primero)
  // format string to JsonString with some modifications.
  jsonString := strings.Replace(str,"<>", "0",-1)
  jsonString = strings.Replace(jsonString,"<", "{\"",-1)
  jsonString= strings.Replace(jsonString,">", "\"}",-1)
  jsonString= strings.Replace(jsonString," ", ",\"",-1)
  jsonString= strings.Replace(jsonString,":", "\":",-1)
  jsonString= strings.Replace(jsonString,",\"\"","",-1)
  jsonString= strings.Replace(jsonString,"},\"", "}",-1)
  jsonString= strings.Replace(jsonString,","," ",-1)
  jsonString= strings.Replace(jsonString,"{"," ",-1)
  jsonString= strings.Replace(jsonString,"}","",-1)
  jsonString="\""+jsonString
  //fmt.Println("-------------------------jsonString-----------------------")
  //fmt.Println(jsonString)
  //fmt.Println("-------------------------jsonString-----------------------")
  if path == "huawei-ifm:ifm/interfaces/interface/ifDynamicInfo" { // caso particular....
    jsonString= strings.Replace(jsonString,"IfOperStatus_UPifName\"","IfOperStatus_UP \"ifName\"",-1)
  }
  fmt.Println("-------------------------jsonString-----------------------")
  fmt.Println(jsonString)
  lastQuote := rune(0)
      f := func(c rune) bool {
          switch {
          case c == lastQuote:
              lastQuote = rune(0)
              return false
          case lastQuote != rune(0):
              return false
          case unicode.In(c, unicode.Quotation_Mark):
              lastQuote = c
              return false
          default:
              return unicode.IsSpace(c)

          }
      }

    // splitting string by space but considering quoted section
    items := strings.FieldsFunc(jsonString, f)

    // create and fill the map
    m := make(map[string]string)
    for _, item := range items {
        x := strings.Split(item, ":")
        m[x[0]] = x[1]
    }
    // get keys and vals of fields
    var keys []string
    var vals []string
    for k, v := range m {
        name:= strings.Replace(k,"\"","",-1) // remove quotes
        keys = append(keys, name)
        vals = append(vals, v)

    }
    // Adaptation to resolve Huawei bad struct Data.
    if path == "huawei-ifm:ifm/interfaces/interface" {
      if Find(keys, "ifAdminStatus") == -1 {
        keys = append(keys, "ifAdminStatus")
        vals = append(vals, "IfAdminStatus_DOWN")
      }
    }
    // Adaptation to resolve Huawei bad struct Data.
    if path == "huawei-ifm:ifm/interfaces/interface/ifDynamicInfo" {
      if Find(keys, "ifOperStatus") == -1 {
        keys = append(keys, "ifOperStatus")
        vals = append(vals, "IfOperStatus_DOWN")
      }
    }

  return keys, vals
}

/* Search for a string in a string array.
  @Params: a String Array
           x String to Search
  @Returns: Returns the index location de x in a or -1 if not Found
*/
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return -1
}