# Carbonclient

This is a lightweight client for the Graphite data interface called Carbon. I
created this specifically to feed in batches of data using the "pickle"
protocol, so that is the only method currently supported.

More details about Carbon and the pickle protocol can be found [in the Graphite
docs](https://graphite.readthedocs.io/en/latest/feeding-carbon.html).

## Feeding data

Create a client using the IP and port of your Graphite/Carbon server:

```go
client := carbonclient.NewCarbonClient("192.168.10.100", carbonclient.PICKLE_PORT)
```

NB: only the pickle protocol is supported, and the default port is available as
a package-level constant as demonstrated above.

Compose your stats as `[]carbonclient.TimedMetric`. Each metric has a `Path`,
and a `Value`. The `Value` must be an instance of
`carbonclient.TimedMetricValue`:

```go
var stats []carbonclient.TimedMetric

stats = append(stats,
  carbonclient.TimedMetric{
    Path: "stats.gauges.widgetcount",
    Value: carbonclient.TimedMetricValue{
      Timestamp: time.Now(),
      Value:     42}})
```

The intended purpose of the Carbon pickle protocol is to send data collected in
the past, but for the purposes of illustration, `time.Now()` is used above.

Finally, send the stats!

```go
err := client.SendMetrics(stats[:])
if err != nil {
  panic(err)
}
```

## License

```text
        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
                    Version 2, December 2004 

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net> 

 Everyone is permitted to copy and distribute verbatim or modified 
 copies of this license document, and changing it is allowed as long 
 as the name is changed. 

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION 

  0. You just DO WHAT THE FUCK YOU WANT TO.
```
