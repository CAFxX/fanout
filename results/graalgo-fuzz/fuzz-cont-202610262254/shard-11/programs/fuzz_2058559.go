package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint32; F1 uint; F2 int32 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint }

func main() {
    v1 := []int{939, 9223372036854775807}
    var v2 uint = 255
    var v3 float32 = (1.5 - 3.14159)
    var v4 bool = (42 == 42)
    var v5 uint = v2
    if ((-671 == 9223372036854775807) || v4) {
        v6 := GenS1{F0: 1, F1: -42}
        _ = v6
    }
    _ = len(v1)
    if (v4 || true) {
        var v7 int8 = 50
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: (fi8 ^ 32767), F1: 18}
        switch fi8 {
        case -490:
            fmt.Println(v3, v5, v3)
        case 352:
            v10 := GenS2{F0: v5, F1: -418}
            _ = v10
        case 127:
            v11 := GenS3{F0: 2147483647, F1: (v2 - v2), F2: 2147483647}
            _ = v11
        }
        var v12 float64 = 0.1
        _ = fi8
        _ = v12
        _ = v9
    }
    v4 = ((!v4) || false)
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 float64 = (1.5 + -2.25)
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            if (128 <= 42) {
                v17 := GenS3{F0: 4294967295, F1: v5, F2: 1}
                v18 := GenS1{F0: -9223372036854775808, F1: 127}
                _ = v17
                _ = v18
            }
            if true {
                fmt.Println(v2, fi16, v4)
            }
            _ = fi16
        }
        _ = fi13
    }
    if (!v4) {
        v19 := GenS4{F0: 42, F1: v5, F2: 0}
        _ = v19
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS5{F0: v5}
            var v22 uint32 = 65535
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS2{F0: v5, F1: fi23}
            fmt.Println(v4, v5, v5)
            _ = fi23
            _ = v24
        }
        fmt.Println(v2, v4, v4)
    }
    if ((!v4) && (v2 > v5)) {
        v2 -= (^(v5))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
