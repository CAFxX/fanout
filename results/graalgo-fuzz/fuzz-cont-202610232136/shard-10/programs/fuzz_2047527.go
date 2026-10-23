package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }

func main() {
    v1 := []int{-983}
    var v2 string = "hello"
    var v3 uint16 = 255
    var v4 bool = ((!(!(-820 <= 1000))) || (42 <= 525))
    v5 := make(map[string]int)
    var v6 uint16 = 1
    if (!v4) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 float32 = 0.1
            v5[(v2 + v2)] = fi7
            var v9 uint64 = 18446744073709551615
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
    }
    fmt.Println(v4, v4)
    if v4 {
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 uint16 = (^(v6))
                var v13 uint16 = (v3 - (v12 << ((242) & 63)))
                v14 := GenS1{F0: v4, F1: v3, F2: 20}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 uint64 = 0
            _ = fi10
            _ = v15
        }
    } else {
        v3 = (v6 ^ (v6 << ((0) & 63)))
        fmt.Println(v2)
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int16 = 0
            _ = fi16
            _ = v17
        }
    }
    v3 -= (v6 / ((v6) | 1))
    for fi18 := 0; fi18 < 4; fi18++ {
        v6 -= 764
        _ = fi18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint16 = (v3 >> ((65535) & 63))
        _ = fi19
        _ = v20
    }
    switch 4294967295 {
    case 1:
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    default:
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(v4, v3, v3)
            _ = fi21
        }
    }
    v6 = (v3 << ((2) & 63))
    v6 = 2
    var v22 bool = false
    var v23 float64 = float64(0.0)
    var v24 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
