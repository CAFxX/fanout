package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
func genFn0() int64 {
    var v6 float64 = float64((0.0 + 0.0))
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{fi7, (fi7 ^ (fi7 * (fi7 - fi7))), 128, fi7}
        v6 -= v6
        _ = fi7
        _ = v8
    }
    _ = v6
    return -1
}
func genFn1() uint16 {
    var v9 uint = 42
    var v10 float32 = ((1.5 + (3.14159 - 1.5)) - 1e10)
    var v11 int8 = 9
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 0
    var v3 bool = (0 > 2147483647)
    v4 := []int{128, -1, -61, 9223372036854775807}
    var v5 uint = 502
    if false {
        var v12 string = "hello"
        _ = v12
    } else {
        var v13 int64 = -1
        v3 = (793 < 510)
        var v14 uint64 = 18446744073709551615
        _ = v13
        _ = v14
    }
    v2 -= v2
    for fi15 := 0; fi15 < 18; fi15++ {
        v4 = append(v4, fi15)
        v2 = (v2 | (v2 + v2))
        _ = fi15
    }
    v4 = append(v4, -1)
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(fi16, v5, v2)
        _ = fi16
    }
    if (!(50 == -128)) {
        v1 = ("" + v1)
        var v17 uint32 = 2147483648
        _ = v17
    }
    _ = len(v4)
    var v18 uint8 = 0
    var v19 uint64 = 18446744073709551615
    if (3.14159 == 1.5) {
        fmt.Println(v18)
    } else {
        v5 = v5
        if len(v4) > 0 {
            v4[0] = -221
        }
    }
    var v20 uint16 = 603
    var v21 int32 = 943
    if v3 {
        v22 := GenS1{F0: (0.0 == 1e10), F1: 2147483648, F2: v5}
        switch "" {
        case "\n\t":
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 string = "foo bar"
                v25 := GenS1{F0: v3, F1: 2147483647, F2: (v5 + (v5 << ((v5) & 63)))}
                _ = fi23
                _ = v24
                _ = v25
            }
        case "":
            v19 = (^(v19))
        default:
            var v26 uint32 = 714
            _ = v26
        }
        _ = v22
    } else {
        fmt.Println(v18, v20)
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
