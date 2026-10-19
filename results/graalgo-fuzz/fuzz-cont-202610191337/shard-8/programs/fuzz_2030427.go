package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16 }
func genFn0() int32 {
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: 42, F1: 18}
        var v10 uint64 = 42
        var v11 string = "a"
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v7
    return 2147483647
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 1
    var v3 int64 = 324
    v4 := make(map[uint32]int)
    var v5 uint64 = 147
    v6 := []int{4294967295}
    v12 := GenS1{F0: 777, F1: 792}
    v5 -= 4294967295
    v2 += 116
    v13 := GenS2{F0: 32767}
    if false {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint16 = 2
            _ = fi14
            _ = v15
        }
        v3 = (v3 ^ 531)
        delete(v4, 65535)
    } else {
        var v16 float32 = ((float32((0.0 - (1e10 / -2.25))) - (3.14159 - 3.14159)) + 0.0)
        _ = v16
    }
    var v17 uint64 = v5
    fmt.Println(v2)
    switch 32768 {
    case -9223372036854775808:
        v18 := GenS1{F0: 0, F1: 42}
        _ = v18
    case -818:
        fmt.Println(v3, v1)
    default:
        v5 -= v17
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
