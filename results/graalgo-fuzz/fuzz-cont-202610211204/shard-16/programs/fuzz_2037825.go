package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -798
    var v3 uint16 = 195
    var v4 float64 = 0.1
    var v5 bool = (!true)
    var v6 uint8 = 128
    v7 := GenS1{F0: v5}
    v8 := v1[int(v3)]
    fmt.Println(v4, v8, v5)
    v6 = 2
    fmt.Println(v3)
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v6)
        _ = fi9
    }
    switch "" {
    case "a":
        var v10 float64 = v4
        _ = v10
    case "hello":
        if (false && v5) {
            var v11 uint16 = v3
            _ = v11
        }
    default:
        switch (v8 + v8) {
        case "\n\t":
            var v12 float32 = (0.0 / (1.5 / 1.0))
            _ = v12
        case "a":
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int32 = 985
                v15 := GenS1{F0: false}
                if (65535 >= 65535) {
                    fmt.Println(v6)
                } else {
                    v2 -= 1
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        case "hello":
            var v16 uint32 = 614
            _ = v16
        default:
            v17 := GenS1{F0: v5}
            _ = v17
        }
    }
    if (-1 <= 42) {
        var v18 float32 = -2.25
        _ = v18
    } else {
        if (!(2 == 1)) {
            v4 = v4
            var v19 uint8 = (v6 & v6)
            _ = v19
        } else {
            v3 += v3
            var v20 int8 = 100
            _ = v20
        }
    }
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
