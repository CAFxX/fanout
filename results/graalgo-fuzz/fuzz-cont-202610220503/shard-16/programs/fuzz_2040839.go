package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }

func main() {
    var v1 string = ""
    var v2 float32 = ((-2.25 + (1.5 * 3.14159)) * 1.5)
    v3 := make(map[string]int)
    var v4 float64 = 1.5
    v5 := []uint32{4294967295, 430, 0}
    var v6 int8 = 0
    v5 = append(v5, 65535)
    switch 4294967295 {
    case -582:
        fmt.Println(v6)
    case 88:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int16 = -32768
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint8 = 2
                var v11 int8 = v6
                v12 := GenS1{F0: 4294967295, F1: 42, F2: fi7}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
    case -376:
        var v13 int32 = 42
        _ = v13
    default:
        var v14 uint16 = 65535
        _ = v14
    }
    var v15 int = -677
    _ = len(v3)
    for fi16 := 0; fi16 < 15; fi16++ {
        v3["a"] = -639
        _ = fi16
    }
    var v17 int32 = 922
    fmt.Println(v17, v17, v17)
    var v18 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
