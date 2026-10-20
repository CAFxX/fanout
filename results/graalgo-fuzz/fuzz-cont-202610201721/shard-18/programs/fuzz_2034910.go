package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int8, p2 bool) int32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(p1)
        var v7 int8 = 56
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    v9 := GenS1{F0: -28, F1: 0, F2: (p1 % ((100) | 1))}
    p0 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}
func genFn1() int32 {
    var v10 float64 = ((0.0 - (-2.25 / 0.1)) + (0.0 + 0.1))
    var v11 int = 128
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int = -1
    var v2 float64 = 1.5
    v3 := []int{(v1 & (v1 % (((^(v1))) | 1))), v1}
    var v4 string = ""
    v5 := make(map[string]bool)
    _ = len(v3)
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1)
        var v12 bool = false
        _ = v12
    } else {
        v4 = "hello"
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if (857 == 1000) {
            v14 := GenS2{F0: 1}
            var v15 int = fi13
            _ = v14
            _ = v15
        } else {
            v5[v4] = false
            v16 := GenS1{F0: (v1 * -2147483648), F1: -9223372036854775808, F2: 127}
            if (((0.1 / -2.25) <= ((1.5 + ((0.0 + (float32(v2) / (0.0 / (3.14159 / ((0.0 + 0.1) * float32(v2)))))) + -2.25)) / 3.14159)) || false) {
                v17 := GenS1{F0: fi13, F1: 519, F2: -31}
                _ = v17
            }
            _ = v16
        }
        var v18 bool = false
        _ = fi13
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 int = (v1 % ((42) | 1))
        v21 := GenS1{F0: fi19, F1: 9223372036854775807, F2: -30}
        _ = fi19
        _ = v20
        _ = v21
    }
    v1 = v1
    var v22 int64 = 693
    for fi23 := 0; fi23 < 11; fi23++ {
        v1 = (fi23 * fi23)
        _ = len(v3)
        v1 += (v1 << ((1) & 63))
        _ = fi23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if len(v3) > 0 {
            v3[0] = (v1 * (v1 & 127))
        }
        _ = fi24
    }
    var v25 float64 = (1.5 * 1.5)
    var v26 int = v1
    delete(v5, "a")
    switch v1 {
    case 521:
        var v27 uint8 = 101
        _ = v27
    }
    v2 = v25
    var v28 int16 = 0
    var v29 bool = (1 != genFn1())
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
