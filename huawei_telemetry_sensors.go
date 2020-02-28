package huawei_sensorPath
import(
      "fmt"
      "reflect"
      "time"
      "strings"
      "unicode"
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-telemetry"  //"telemetry"          // path to huawei-telemetry.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-bfd"        //"huawei_bfd"         // path to huawei-bfd.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-bgp"        //"huawei_bgp"         // path to huawei-bgp.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-devm"       //"huawei_devm"        // path to huawei-devm.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-driver"     //"huawei_driver"      // path to huawei-driver.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ifm"        //"huawei_ifm"         // path to huawei-ifm.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-isis"       //"huawei_isis"        // path to huawei-isis.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-mpls"       //"huawei_mpls"        // path to huawei-mpls.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ospfv2"     //"huawei_ospfv2"      // path to huawei-ospfv2.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-ospfv3"     //"huawei_ospfv3"      // path to huawei-ospfv3.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-qos"        //"huawei_qos"         // path to huawei-qos.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-sem"        //"huawei_sem"         // path to huawei-sem.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-telemEmdi"  //"huawei_telemEmdi"   // path to huawei-TelemEmdi.proto
      "github.com/DamRCorba/huawei_telemetry_sensors/sensors/huawei-trafficmng" //"huawei_trafficmng"  // path to huawei-trafficmng.proto
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
  switch path {
  case "huawei-bfd":
      return &huawei_bfd.Bfd{}

  case "huawei-bgp":
    return &huawei_bgp.ESTABLISHED{}

  case "huawei-devm":
    return &huawei_devm.Devm{}

  case "huawei-driver":
    return &huawei_driver.HwEntityInvalid{}

  case "huawei-ifm":
    return &huawei_ifm.Ifm{}

  case "huawei-isis":
    return &huawei_isiscomm.IsisAdjacencyChange{}

  case "huawei-mpls":
    return &huawei_mpls.Mpls{}

  case "huawei-ospfv2":
    return &huawei_ospfv2.OspfNbrStateChange{}

  case "huawei-ospfv3":
    return &huawei_ospfv3.Ospfv3NbrStateChange{}

  case "huawei-qos":
    return &huawei_qos.Qos{}

  case "huawei-sem":
    return &huawei_sem.HwStorageUtilizationResume{}

  case "huawei-telmEmdi":
  case "huawei-emdi":
    return &huawei_telemEmdi.TelemEmdi{}

  case "huawei-trafficmng":
    return &huawei_trafficmng.Trafficmng{}

  default:
    fmt.Println("Error Sensor Desconocido")
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
      fooType := reflect.TypeOf(huawei_devm.Devm_Ports_Port_OpticalInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "devm/powerSupplys/powerSupply/powerEnvironments/powerEnvironment":
      fooType := reflect.TypeOf(huawei_devm.Devm_PowerSupplys_PowerSupply_PowerEnvironments_PowerEnvironment{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    case "devm/temperatureInfos/temperatureInfo":
      fooType := reflect.TypeOf(huawei_devm.Devm_TemperatureInfos_TemperatureInfo{})
      for i := 0; i < fooType.NumField(); i ++ {
        keys := fooType.Field(i)
        resolve[LcFirst(keys.Name)] = keys.Type
        }
        break;
    }

    return resolve

  case "huawei-driver":
    return resolve

  case "huawei-ifm":
    switch splited[1] {
    case "ifm/interfaces/interface":
      fooType := reflect.TypeOf(huawei_ifm.Ifm_Interfaces_Interface{})
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

  case "huawei-isis":
    return resolve

  case "huawei-mpls":
    return resolve

  case "huawei-ospfv2":
    return resolve

  case "huawei-ospfv3":
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
    return resolve

  case "huawei-telmEmdi":
  case "huawei-emdi":
    return resolve

  case "huawei-trafficmng":
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
  case uint32: resolve, _ := strconv.ParseUint(val,10,32);
                //fmt.Println("uint32 selected", resolve, da)
                return resolve;
  case uint64: resolve,_ :=  strconv.ParseUint(val,10,64); return resolve;
  case int32: resolve,_ :=  strconv.ParseInt(val,10,32); return resolve;
  case int64: resolve,_ :=  strconv.ParseInt(val,10,64); return resolve;
  case float64: resolve,_ :=  strconv.ParseFloat(val,64); return resolve;
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
  name:= strings.Replace(subfield,"\"","",-1)
  endPointTypes:=GetTypeValue(path)
  grouper.Add(path, tags, timestamp, string(name), decodeVal(endPointTypes[name], vals))
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
func SearchKey(Message *telemetry.TelemetryRowGPB, sensorType string)  ([]string, []string){
  sensorMsg := GetMessageType(sensorType)
  err := proto.Unmarshal(Message.Content, sensorMsg)
  if (err != nil) {
    panic(err)
  }
  primero := reflect.ValueOf(sensorMsg).Interface()
  fmt.Println(reflect.TypeOf(primero))
  str := fmt.Sprintf("%v", primero)
  // format string to JsonString with some modifications.
  jsonString := strings.Replace(str,"<>", "NoStats",-1)
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

  return keys, vals
}
