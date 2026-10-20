package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }

func main() {
    v1 := []uint32{793, 493, 72}
    v2 := []uint32{734, 0, 403}
    var v3 int32 = 0
    if len(v2) > 0 {
        v2[0] = 2
    }
    fmt.Println(v3)
    switch "a" {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v3 += v3
    _ = len(v1)
    v3 = v3
    var v4 int64 = -1
    _ = len(v2)
    fmt.Println(v3, v3)
    for fi5 := 0; fi5 < 14; fi5++ {
        fmt.Println(v3, v4, v4)
        var v6 int32 = (v3 | (v3 << ((435) & 63)))
        _ = len(v1)
        _ = fi5
        _ = v6
    }
    var v7 float32 = -2.25
    if (!true) {
        fmt.Println(v3, v4, v3)
    } else {
        var v8 int8 = 0
        v9 := GenS1{F0: -805, F1: v3, F2: 663}
        if (1 == 2) {
            fmt.Println(v3, v7)
            var v10 int16 = int16(v4)
            _ = v10
        } else {
            if (20 < 2147483648) {
                var v11 int16 = -226
                _ = v11
            } else {
                var v12 uint64 = 2
                _ = v12
            }
        }
        _ = v8
        _ = v9
    }
    fmt.Println(v4, v4)
    var v13 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
