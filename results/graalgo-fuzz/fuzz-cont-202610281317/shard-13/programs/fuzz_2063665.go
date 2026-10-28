package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0() uint {
    v5 := GenS1{F0: 0, F1: -661}
    v6 := GenS2{F0: 2147483647, F1: 127, F2: 684}
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 uint8) bool {
    if false {
        var v7 int64 = 1
        fmt.Println(v7, v7)
        _ = v7
    }
    v8 := []int{127, 128}
    var v9 float64 = 0.1
    _ = p0
    _ = v8
    _ = v9
    return true
}

func main() {
    var v1 int8 = 1
    var v2 uint16 = 748
    var v3 uint = 2
    v4 := []string{"", "\n\t", ""}
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v3)
        _ = fi10
    }
    v1 += (v1 * (v1 * (v1 - (v1 << ((18446744073709551615) & 63)))))
    _ = len(v4)
    if (-1 > 877) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint64 = 42
            var v13 int64 = 723
            _ = fi11
            _ = v12
            _ = v13
        }
        switch -2147483648 {
        case 128:
            fmt.Println(v2)
        case 32768:
            var v14 uint = (^(v3))
            _ = v14
        case -2147483648:
            fmt.Println(v1)
        }
        fmt.Println(v1)
    }
    v2 = v2
    v4 = append(v4, "a")
    var v15 int32 = -1
    for fi16 := 0; fi16 < 5; fi16++ {
        fmt.Println(v15, v2)
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 uint16 = v2
            v4 = append(v4, "foo bar")
            fmt.Println(v3, v18, v18)
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS2{F0: 0, F1: -48, F2: (v2 - (v2 + 2))}
    fmt.Println(v15, v1)
    if false {
        var v20 float32 = (1e10 + float32((1.5 + 3.14159)))
        var v21 uint8 = 71
        fmt.Println(v3)
        _ = v20
        _ = v21
    } else {
        var v22 bool = (true && false)
        v22 = (1 == -624)
        for fi23 := 0; fi23 < 1; fi23++ {
            switch "a" {
            case "foo bar":
                var v24 float64 = 1e10
                _ = v24
            case "a":
                v1 -= v1
            default:
                var v25 int8 = (v1 | -1)
                _ = v25
            }
            _ = fi23
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
