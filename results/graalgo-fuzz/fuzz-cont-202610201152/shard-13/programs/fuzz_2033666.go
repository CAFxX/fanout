package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 int8 = 1
    var v2 string = "foo bar"
    var v3 int64 = -765
    var v4 uint16 = 0
    var v5 bool = ((true || true) && (1 != 1))
    var v6 uint16 = (^(v4))
    v7 := []int{9223372036854775807, 42}
    var v8 int16 = 1000
    _ = len(v7)
    v1 -= v1
    v2 = "foo bar"
    v9 := GenS1{F0: -487, F1: 4294967295}
    for fi10 := 0; fi10 < 4; fi10++ {
        _ = len(v7)
        v1 -= v1
        _ = fi10
    }
    switch -9223372036854775808 {
    case -2147483648:
        var v11 float32 = (0.0 / 3.14159)
        _ = v11
    case -707:
        v12 := GenS2{F0: -2147483648}
        _ = v12
    case 128:
        fmt.Println(v5)
    default:
        v4 = (v4 ^ 0)
    }
    v5 = ((int(v1) <= -996) || v5)
    v13 := GenS2{F0: 32767}
    fmt.Println(v6, v3)
    v14 := GenS1{F0: 0, F1: 42}
    var v15 uint16 = (v6 >> ((4294967295) & 63))
    var v16 bool = (2147483647 == 32768)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v17 int16 = (^(v8))
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v8, v3)
        var v19 float64 = ((0.1 * 1e10) / 1e10)
        if len(v7) > 0 {
            v7[0] = fi18
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
