package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 42
    v2 := make(map[string]string)
    var v3 int32 = -548
    v4 := []string{}
    var v5 int16 = -32768
    v6 := GenS1{F0: v3}
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int = v1
            _ = fi7
            _ = v8
        }
    case "":
        v3 = 947
    }
    var v9 int16 = v5
    v5 = 42
    var v10 int64 = 1
    for fi11 := 0; fi11 < 18; fi11++ {
        fmt.Println(v9)
        var v12 uint = 0
        _ = fi11
        _ = v12
    }
    switch (v1 ^ (v1 - -9223372036854775808)) {
    case 32767:
        fmt.Println(v3)
    case 128:
        var v13 int16 = (v9 << ((838) & 63))
        _ = v13
    case -1:
        var v14 uint8 = 1
        _ = v14
    default:
        v15 := GenS1{F0: v3}
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := v2["\n\t"]
        v3 += (v3 * v3)
        _ = fi16
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
