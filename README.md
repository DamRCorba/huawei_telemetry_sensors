# Huawei Telemetry Sensors

Este proyecto interpreta los mensajes de telemetria de los routers Huawei NE40E en V8R10

## Sensores de Metricas

```
<Huawei>dis telemetry sensor-path | i Sample
Info: It will take a long time if the content you search is too much or the string you input is too long, you can press CTRL_C to break.
Sample(S) : Serial sample.
Sample(P) : Parallel sample.
Parallel sampled data may not be sent through UDP in out-of-band mode.
------------------------------------------------------------------------------------------------------------------
Type        MinPeriod(ms)  MaxEachPeriod  Path      
------------------------------------------------------------------------------------------------------------------
Sample(S)   10000          --             huawei-devm:devm/cpuInfos/cpuInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/fans/fan - Ok
Sample(S)   10000          --             huawei-devm:devm/memoryInfos/memoryInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/ports/port - Solo devuelve posicion y entIndex
Sample(S)   300000         --             huawei-devm:devm/ports/port/opticalInfo - Probado Ok
Sample(S)   300000         --             huawei-devm:devm/powerSupplys/powerSupply/powerEnvironments/powerEnvironment - Ok.
Sample(S)   300000         --             huawei-devm:devm/temperatureInfos/temperatureInfo - Ok
Sample(P)   300000         --             huawei-emdi:emdi/emdiTelemReps/emdiTelemRep -
Sample(P)   10000          --             huawei-emdi:emdi/emdiTelemRtps/emdiTelemRtp -
Sample(S)   10000          --             huawei-ifm:ifm/interfaces/interface - Solo devuelve AdminStatus si la interface esta UP. Sino solo index
Sample(P)   100            20             huawei-ifm:ifm/interfaces/interface/ifClearedStat - No Interesa
Sample(S)   10000          --             huawei-ifm:ifm/interfaces/interface/ifDynamicInfo - Devuelve el estado operativo de la interfaz
Sample(P)   100            20             huawei-ifm:ifm/interfaces/interface/ifStatistics - Probado ok
Sample(P)   1000           20             huawei-ifm:ifm/interfaces/interface/ifStatistics/ethPortErrSts - OK
Sample(S)   10000          --             huawei-mpls:mpls/mplsTe/rsvpTeTunnels/rsvpTeTunnel/tunnelPaths/tunnelPath - No hay RSVP
Sample(P)   1000           200            huawei-qos:qos/qosBuffers/qosBuffer - No da Informacion de metricas.
Sample(P)   100            20             huawei-qos:qos/qosIfQoss/qosIfQos/qosPolicyApplys/qosPolicyApply/qosPolicyStats/qosPolicyStat/qosRuleStats/qosRuleStat Da informacion de buffer por politica por puerto.
Sample(P)   1000           10             huawei-qos:qos/qosPortQueueStatInfos/qosPortQueueStatInfo - Probado Ok
Sample(S)   300000         --             huawei-trafficmng:trafficmng/tmSlotSFUs/tmSlotSFU/sfuStatisticss/sfuStatistics Probado Ok.
------------------------------------------------------------------------------------------------------------------

```

### Sensor ISSUE

The NE40-X3/8/16A in VRP V8R10 sends the information of the sensor "huawei-qos:qos/qosPortQueueStatInfos/qosPortQueueStatInfo" misordered. The plugin not correct this input for future compatibility. Thats an ISSUE from Huawei so this is the information needed to use this plugin. Check if your device sends it correctly.


| Sensor Queue Indication | Real Queue     |
| :---------------------- | :------------- |
| Null                    | BE             |
| AF2                     | AF1            |
| AF3                     | AF2            |
| AF4                     | AF3            |
| BE                      | AF4            |
| CS6                     | EF             |
| CS7                     | CS6            |
| EF                      | CS7            |
