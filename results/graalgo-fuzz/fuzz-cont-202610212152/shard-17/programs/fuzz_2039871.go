package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{65535, 0}
    var v3 int64 = 42
    v3 = (^(v3))
    v2 = append(v2, 433)
    if true {
        var v4 int8 = 1
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int64 = v3
            v4 += (v4 << ((18446744073709551615) & 63))
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 100, F1: 2}
        _ = v4
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi9
    }
    v10 := GenS1{F0: -37, F1: 2}
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 int8 = 113
            var v14 float64 = 0.1
            v3 = v3
            _ = fi12
            _ = v13
            _ = v14
        }
        if false {
            var v15 uint = 842
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 float64 = (float64(-2.25) / (float64(1.5) / 1e10))
                var v18 float64 = (((1.5 + v17) / ((v17 + ((float64(1.5) * 1.5) / v17)) * v17)) / v17)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
        _ = fi11
    }
    v3 = (v3 << ((957) & 63))
    fmt.Println(v1, v3)
    v19 := GenS1{F0: -1, F1: 0}
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if true {
        v1 += 0.0
        var v20 float64 = (0.0 / 0.1)
        _ = v20
    } else {
        fmt.Println(v3)
        v21 := GenS2{F0: (v3 ^ 49), F1: "hello"}
        fmt.Println(v3, v1)
        _ = v21
    }
    var v22 int = 42
    for fi23 := 0; fi23 < 13; fi23++ {
        v1 -= 1e10
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
