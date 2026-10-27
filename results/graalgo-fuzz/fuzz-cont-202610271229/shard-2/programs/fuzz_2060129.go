package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint; F2 uint64 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]bool)
    var v3 uint = 619
    v4 := make(map[string]string)
    var v5 int64 = 0
    v6 := GenS1{F0: -1, F1: v5}
    v5 -= v5
    v7 := GenS1{F0: -962, F1: 9223372036854775807}
    var v8 string = "\n\t"
    if false {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 uint16 = 2
            if false {
                var v11 bool = true
                var v12 float64 = 0.1
                _ = v11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            fmt.Println(v8, v1)
            v14 := GenS1{F0: 2147483647, F1: v5}
            var v15 uint8 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
    } else {
        fmt.Println(v1, v8, v5)
        v1 = (v1 | v1)
    }
    v16 := GenS2{F0: 42}
    v17 := []int{2147483647, 32767, -879}
    for fi18 := 0; fi18 < 4; fi18++ {
        fmt.Println(v8)
        switch "hello" {
        case "foo bar":
            var v19 uint64 = 271
            _ = v19
        default:
            if false {
                _ = len(v4)
            }
        }
        _ = fi18
    }
    v20 := GenS3{F0: 1, F1: 2, F2: 514}
    v4[v8] = "foo bar"
    var v21 float32 = 1e10
    v3 -= (v3 << ((4294967295) & 63))
    var v22 bool = false
    var v23 int64 = v5
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 int8 = 112
        _ = fi24
        _ = v25
    }
    fmt.Println(v21, v23, v3)
    v26 := GenS4{F0: v22, F1: v3, F2: 25}
    v27 := GenS5{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
