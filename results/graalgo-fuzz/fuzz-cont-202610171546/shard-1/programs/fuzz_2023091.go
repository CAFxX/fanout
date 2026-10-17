package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int; F1 int64 }
func genFn0() int64 {
    var v6 int16 = 42
    if (108 == -1) {
        v6 -= v6
        v7 := GenS1{F0: 1, F1: "foo bar", F2: 253}
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    var v9 int32 = 42
    _ = v6
    _ = v9
    return int64(v9)
}
func genFn1(p0 uint, p1 int16, p2 bool) string {
    var v10 float32 = 0.1
    p2 = ((p1 >> ((p0) & 63)) > 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "a"
}

func main() {
    v1 := []int{747}
    var v2 float32 = 1e10
    var v3 bool = (-777 == -676)
    var v4 uint = 18446744073709551615
    var v5 uint = 65535
    fmt.Println(v2, v4)
    fmt.Println(v2, v2)
    if (65535 == 840) {
        var v11 float64 = float64(3.14159)
        _ = v11
    } else {
        switch 128 {
        case 2147483647:
            v4 = 734
        case 1:
            if (4294967295 < 1) {
                var v12 uint32 = 1
                v4 += 42
                v4 = 0
                _ = v12
            } else {
                v5 += (v4 + v4)
            }
        case 42:
            _ = len(v1)
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 uint32 = 2147483647
                _ = fi13
                _ = v14
            }
        }
    }
    if (1 >= 598) {
        var v15 uint64 = 322
        if v3 {
            v16 := GenS1{F0: 303, F1: "foo bar", F2: 42}
            var v17 int32 = -1
            if len(v1) > 0 {
                v1[0] = 128
            }
            _ = v16
            _ = v17
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v15
    }
    v18 := GenS1{F0: 2147483647, F1: "", F2: 128}
    v19 := GenS2{F0: 960, F1: 42}
    if true {
        v20 := GenS3{F0: 2147483647, F1: -224}
        var v21 uint32 = 601
        var v22 float32 = 0.1
        _ = v20
        _ = v21
        _ = v22
    } else {
        var v23 uint64 = 1
        fmt.Println(v4, v5, v23)
        _ = v23
    }
    if ("" > "foo bar") {
        var v24 float64 = 1.5
        _ = v24
    } else {
        v4 = 0
        v25 := GenS1{F0: 42, F1: (("\n\t" + "a") + "foo bar"), F2: uint8(v5)}
        _ = v25
    }
    v3 = false
    switch 9223372036854775807 {
    case -451:
        v5 = 1
    case 9223372036854775807:
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 float32 = (1e10 * 0.0)
            fmt.Println(v2, v3)
            v27 = 0.1
            _ = fi26
            _ = v27
        }
    case 127:
        var v28 int8 = 1
        _ = v28
    default:
        if (-115 > -1) {
            switch -146 {
            case 187:
                v5 -= (^(v5))
            case -377:
                v2 = v2
            }
        }
    }
    fmt.Println(v4, v2, v3)
    for fi29 := 0; fi29 < 10; fi29++ {
        if v3 {
            fmt.Println(fi29, v4, v3)
        }
        v2 = v2
        _ = fi29
    }
    fmt.Println(v2, v4, v2)
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(fi30, v4, v4)
        v1 = append(v1, 439)
        _ = fi30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
