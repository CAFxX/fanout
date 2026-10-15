package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
func genFn0(p0 bool) int16 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 274
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]bool)
    var v3 int8 = -77
    v4 := make(map[int]bool)
    var v5 uint32 = 2147483647
    v6 := make(map[uint32]string)
    fmt.Println(v5)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v1)
        switch -940 {
        case -177:
            fmt.Println(v1, v3, v5)
        case 4294967295:
            for fi8 := 0; fi8 < 4; fi8++ {
                v5 = v5
                var v9 uint64 = 2
                _ = fi8
                _ = v9
            }
        default:
            v3 -= -79
        }
        _ = fi7
    }
    v10 := GenS1{F0: (!((float64(((1.5 / 1e10) / 1e10)) / 3.14159) > (3.14159 + 3.14159))), F1: "foo bar", F2: 0}
    var v11 int = -728
    v12 := []int{}
    if len(v12) > 0 {
        v12[0] = v11
    }
    v1 -= 152
    v11 -= (^(v11))
    v13 := GenS1{F0: false, F1: ("a" + ("\n\t" + "foo bar")), F2: 0}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int8 = v3
        var v16 int8 = 100
        var v17 uint32 = v5
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v5 -= v5
    var v18 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
