package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 128
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    var v5 int64 = v4
    v6 := []int{}
    _ = len(v6)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 string = "foo bar"
        fmt.Println(v8)
        fmt.Println(v8)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 2, F1: 0, F2: (^(v4))}
    var v10 float64 = 0.1
    for fi11 := 0; fi11 < 15; fi11++ {
        switch 128 {
        case -65:
            v4 = v4
        default:
            var v12 int64 = (v4 ^ v5)
            _ = v12
        }
        _ = fi11
    }
    v13 := GenS2{F0: 340}
    v5 -= v4
    fmt.Println(v1)
    v1 = (^(v1))
    var v14 float64 = (v10 - 0.1)
    v15 := GenS3{F0: -911}
    if false {
        for fi16 := 0; fi16 < 1; fi16++ {
            v4 -= int64(v1)
            _ = fi16
        }
        var v17 uint16 = 55
        v2 -= (v2 + 128)
        _ = v17
    } else {
        v18 := GenS3{F0: 1}
        _ = v18
    }
    v19 := GenS2{F0: 65535}
    delete(v3, 1)
    for fi20 := 0; fi20 < 14; fi20++ {
        v2 += v2
        for fi21 := 0; fi21 < 15; fi21++ {
            v4 = v5
            v22 := v3[2147483648]
            _ = fi21
            _ = v22
        }
        if true {
            v1 = v1
        } else {
            v23 := GenS2{F0: 470}
            switch (fi20 % ((fi20) | 1)) {
            case 42:
                if len(v6) > 0 {
                    v6[0] = (fi20 - 801)
                }
            default:
                var v24 int8 = 100
                _ = v24
            }
            _ = v23
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
