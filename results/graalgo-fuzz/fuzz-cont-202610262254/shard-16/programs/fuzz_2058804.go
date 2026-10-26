package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 bool, p2 uint) int16 {
    p2 += 4294967295
    var v6 int16 = 454
    fmt.Println(v6)
    switch 32767 {
    case -9223372036854775808:
        v7 := GenS1{F0: 1}
        _ = v7
    case 2147483647:
        var v8 int64 = 42
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    var v1 int16 = -474
    var v2 uint8 = 127
    v3 := []uint32{2147483647, 233, 2, 2}
    var v4 int8 = 42
    var v5 uint16 = 1
    if true {
        _ = len(v3)
    }
    v2 = 42
    for fi9 := 0; fi9 < 19; fi9++ {
        v1 = 1
        var v10 int32 = 4
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 128
    var v12 uint16 = (v5 & (^(v5)))
    var v13 uint = uint(v12)
    for fi14 := 0; fi14 < 1; fi14++ {
        if len(v3) > 0 {
            v3[0] = 1
        }
        if false {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 int8 = 42
                fmt.Println(fi14, v16)
                if len(v3) > 0 {
                    v3[0] = 764
                }
                _ = fi15
                _ = v16
            }
        }
        fmt.Println(v13)
        _ = fi14
    }
    _ = len(v3)
    if (false && true) {
        var v17 float32 = (0.0 - float32(1.5))
        v18 := GenS1{F0: -983}
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
