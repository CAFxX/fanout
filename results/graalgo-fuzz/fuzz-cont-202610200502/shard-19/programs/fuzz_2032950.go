package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint }
func genFn0() uint16 {
    v6 := []int{2147483647}
    switch "hello" {
    case "":
        var v7 uint = 1
        _ = v7
    case "foo bar":
        var v8 uint8 = 187
        _ = v8
    default:
        var v9 float64 = (-2.25 / -2.25)
        _ = v9
    }
    _ = v6
    return 42
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    var v3 string = v1
    var v4 string = ("foo bar" + v1)
    var v5 float32 = (3.14159 - float32(1e10))
    var v10 int16 = 32767
    v5 = v5
    var v11 int8 = -1
    var v12 int8 = 127
    var v13 int = 9223372036854775807
    for fi14 := 0; fi14 < 8; fi14++ {
        if (false || ((1 == 1) && false)) {
            var v15 int32 = int32(v12)
            _ = v15
        } else {
            v1 = ("foo bar" + "foo bar")
        }
        _ = fi14
    }
    v12 -= v12
    v2[v4] = 2
    var v16 int16 = v10
    v12 += v11
    v17 := []int{(v13 ^ v13), (^(v13)), (^(v13)), (v13 << ((42) & 63))}
    if false {
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := GenS1{F0: v11, F1: (v10 % (((-(v10))) | 1)), F2: 42}
            var v20 bool = false
            fmt.Println(v1, v11)
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int8 = (v12 - 100)
        _ = v21
    } else {
        v12 = v11
        v22 := GenS1{F0: (v11 & (^(v11))), F1: (v10 - (v10 + 572)), F2: 42}
        switch (^(v13)) {
        case -1:
            for fi23 := 0; fi23 < 15; fi23++ {
                fmt.Println(v1, v12)
                _ = fi23
            }
        }
        _ = v22
    }
    v24 := GenS1{F0: v11, F1: 1, F2: 742}
    v25 := GenS1{F0: v12, F1: -734, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
