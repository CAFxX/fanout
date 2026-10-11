package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 int) int32 {
    p0 -= (p0 ^ (-(p0)))
    v4 := GenS1{F0: 0, F1: 65535}
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int8 = 42
        fmt.Println(p0, v6, p1)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 2147483647
}

func main() {
    var v1 bool = (2 > 498)
    var v2 int32 = 42
    v3 := make(map[string]uint32)
    fmt.Println(v1, v1)
    v1 = (-9223372036854775808 > -1)
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 1, F1: 821}
            _ = fi8
            _ = v9
        }
        _ = len(v3)
        _ = fi7
    }
    fmt.Println(v2, v1)
    var v10 int8 = -128
    var v11 float32 = (3.14159 / 1.5)
    switch -9223372036854775808 {
    case 127:
        v2 += 0
    }
    v12 := GenS1{F0: 410, F1: 4294967295}
    var v13 string = ""
    var v14 int16 = 32767
    fmt.Println(v14, v11, v11)
    var v15 int = 9223372036854775807
    var v16 int = (v15 & v15)
    var v17 int32 = v2
    v10 = (v10 >> ((0) & 63))
    v1 = (42 >= 206)
    if v1 {
        v18 := GenS2{F0: (0.1 > (v11 / -2.25))}
        v17 = 116
        _ = v18
    }
    v19 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
