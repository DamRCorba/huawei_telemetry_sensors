package huawei_sensorPath
import(
      "fmt"
      "reflect"
      "time"
      "strings"
      "unicode"
      "telemetry"          // path to huawei-telemetry.proto
      "huawei_bfd"         // path to huawei-bfd.proto
      "huawei_bgp"         // path to huawei-bgp.proto
      "huawei_devm"        // path to huawei-devm.proto
      "huawei_driver"      // path to huawei-driver.proto
      "huawei_ifm"         // path to huawei-ifm.proto
      "huawei_isis"        // path to huawei-isis.proto
      "huawei_mpls"        // path to huawei-mpls.proto
      "huawei_ospfv2"      // path to huawei-ospfv2.proto
      "huawei_ospfv3"      // path to huawei-ospfv3.proto
      "huawei_qos"         // path to huawei-qos.proto
      "huawei_sem"         // path to huawei-sem.proto
      "huawei_telemEmdi"   // path to huawei-TelemEmdi.proto
      "huawei_trafficmng"  // path to huawei-trafficmng.proto
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
  Creates and add metrics from json mapped data in telegraf metrics SeriesGrouper
  @params:
  - grouper (*metric.SeriesGrouper) - pointer of metric series to append data.
  - tags (map[string]string) json data mapped
  - timestamp (time.Time) -
  - path (string) - sensor path
  - subfield (string) - subkey data.

*/
func CreateMetrics(grouper *metric.SeriesGrouper, tags map[string]string, timestamp time.Time, path string, subfield string)  string{
  grouper.Add(path, tags, timestamp, "name", tags[subfield])
  return "Metica - " + subfield+":"+tags[subfield]
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
        keys = append(keys, k)
        vals = append(vals, v)

    }

  return keys, vals
}
