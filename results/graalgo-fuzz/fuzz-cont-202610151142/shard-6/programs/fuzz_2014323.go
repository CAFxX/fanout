package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 int64, p1 int, p2 bool) uint {
    var v5 uint32 = 584
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int16 = -32768
        var v8 int64 = p0
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}
func genFn1(p0 int32, p1 int8) uint64 {
    var v9 uint8 = 46
    var v10 uint32 = 638
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = ((1.5 - 0.0) - (1e10 - 1e10))
    v3 := make(map[uint32]string)
    var v4 float64 = 1e10
    var v11 int64 = 1
    var v12 float32 = (-2.25 + ((0.1 * 0.0) - 1.5))
    var v13 int16 = 354
    var v14 uint32 = 1
    fmt.Println(v2)
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v3[(v14 * (^(v14)))] = "foo bar"
            var v17 int64 = (v11 & (v11 * 9223372036854775807))
            v2 = (v4 - -2.25)
            _ = fi16
            _ = v17
        }
        var v18 uint32 = (v14 & (v14 | (v14 | (v14 ^ v14))))
        v13 = 491
        _ = fi15
        _ = v18
    }
    var v19 string = "hello"
    v20 := GenS1{F0: -812, F1: 315, F2: 4294967295}
    var v21 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
