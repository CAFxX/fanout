package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
func genFn0(p0 string, p1 uint64) uint16 {
    var v4 uint8 = uint8(p1)
    var v5 uint32 = 4294967295
    v6 := GenS1{F0: v5, F1: (p0 + (p0 + p0)), F2: 2}
    p1 -= (p1 & p1)
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{(fi7 ^ 2147483647), fi7}
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int16 = -19
    var v2 float64 = (-2.25 * float64(1e10))
    var v3 int = 88
    v9 := []int{v3, (^(v3)), v3}
    var v10 int16 = (v1 % (((v1 ^ 1000)) | 1))
    var v11 uint16 = 2
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 1, F1: "a", F2: 427}
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int64 = -1
            v10 = (v10 ^ 32767)
            _ = len(v9)
            _ = fi14
            _ = v15
        }
        fmt.Println(v1, v11)
        _ = fi12
        _ = v13
    }
    switch 1 {
    case -896:
        fmt.Println(v10)
    default:
        v10 = v10
    }
    fmt.Println(v3, v2, v1)
    v16 := GenS2{F0: 255, F1: ((v3 + 128) > 0)}
    switch v3 {
    case -2147483648:
        v2 = (((((v2 + v2) - (0.0 - v2)) / 3.14159) + (v2 * 0.0)) + (0.0 * ((0.0 * float64(0.0)) * float64((0.0 * (1.5 + 0.1))))))
    case 0:
        var v17 bool = true
        _ = v17
    }
    _ = len(v9)
    if true {
        v18 := GenS2{F0: 4294967295, F1: false}
        for fi19 := 0; fi19 < 4; fi19++ {
            v9 = append(v9, (^(v3)))
            v20 := GenS1{F0: 2, F1: "a", F2: (v11 << ((uint(v11)) & 63))}
            _ = fi19
            _ = v20
        }
        _ = len(v9)
        _ = v18
    } else {
        var v21 int16 = -32768
        _ = v21
    }
    var v22 int64 = 42
    var v23 int16 = v10
    var v24 uint = 18446744073709551615
    v22 = v22
    fmt.Println(v10)
    var v25 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v9
}
