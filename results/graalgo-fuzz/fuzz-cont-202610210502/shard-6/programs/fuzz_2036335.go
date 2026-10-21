package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0() uint64 {
    var v5 uint = 691
    var v6 int = 2147483647
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{(v6 & 0), (^(v6)), fi7, -1}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int, p1 int32, p2 uint64) uint {
    fmt.Println(p1)
    var v9 uint16 = 557
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    v2 := []string{"hello", "\n\t", "\n\t", ""}
    var v3 uint32 = 227
    var v4 string = "\n\t"
    var v10 int = 2147483647
    var v11 float32 = 1e10
    var v12 uint8 = 250
    v1 = append(v1, "\n\t")
    var v13 uint64 = 693
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v10 = v10
    for fi14 := 0; fi14 < 20; fi14++ {
        v2 = append(v2, "hello")
        if ("" != "a") {
            var v15 uint = genFn1(v10, 1, v13)
            var v16 string = v4
            _ = v15
            _ = v16
        }
        switch v4 {
        case "\n\t":
            v4 = v4
        case "foo bar":
            v17 := GenS1{F0: (1e10 > float64((float32(-2.25) + v11)))}
            _ = v17
        case "a":
            v11 = (float32(0.1) * 0.0)
        default:
            fmt.Println(v4, v11, v3)
        }
        _ = fi14
    }
    if true {
        v18 := GenS2{F0: 132, F1: (2 == v3), F2: -1}
        fmt.Println(v10)
        var v19 int32 = 2147483647
        _ = v18
        _ = v19
    } else {
        v20 := GenS3{F0: 42, F1: 65535, F2: int64(v3)}
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
