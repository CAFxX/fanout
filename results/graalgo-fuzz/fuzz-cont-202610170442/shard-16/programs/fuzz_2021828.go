package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int64, p1 uint) uint64 {
    v7 := GenS1{F0: p0}
    if (!(872 <= 534)) {
        v8 := GenS1{F0: -636}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int = 128
    var v2 float64 = 1e10
    v3 := make(map[string]bool)
    var v4 int64 = -628
    var v5 int16 = -665
    var v6 int8 = 93
    var v9 float64 = v2
    v10 := GenS2{F0: (v1 >> ((355) & 63)), F1: 127, F2: v5}
    v11 := v3[("" + ("a" + ("hello" + "\n\t")))]
    fmt.Println(v5, v1)
    var v12 string = "\n\t"
    switch (v1 % (((^(v1))) | 1)) {
    case 318:
        v13 := GenS3{F0: 685}
        _ = v13
    case 1:
        switch (v12 + ((v12 + v12) + (v12 + (v12 + v12)))) {
        case "\n\t":
            v4 = (v4 + (v4 >> ((255) & 63)))
        }
    default:
        fmt.Println(v1)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        switch (-(fi14)) {
        case -1:
            var v15 int = (^(fi14))
            _ = v15
        case 32767:
            v16 := []int{}
            _ = v16
        case -252:
            if v11 {
                v11 = (false || ((((v2 + v9) - (v2 / v9)) + (float64((-2.25 * 0.0)) / v2)) >= v2))
                v17 := GenS4{F0: -43}
                _ = v17
            } else {
                if v11 {
                    var v18 float64 = ((-2.25 * v2) - (0.0 + (0.1 + (v9 / v9))))
                    fmt.Println(v4, v2)
                    _ = v18
                } else {
                    fmt.Println(v6, v5, v1)
                    switch "" {
                    case "":
                        v5 = v5
                    }
                }
                v19 := []int{v1, v1, (fi14 >> ((998) & 63))}
                _ = v19
            }
        default:
            v6 += -128
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            v3[(v12 + v12)] = (2147483647 < -2147483648)
            _ = fi20
        }
        _ = fi14
    }
    switch v1 {
    case 191:
        var v21 bool = false
        _ = v21
    case 145:
        delete(v3, "\n\t")
    default:
        var v22 uint32 = 1
        _ = v22
    }
    if v11 {
        v23 := []int{}
        v5 = (v5 * (v5 ^ v5))
        _ = v23
    } else {
        var v24 string = v12
        _ = v24
    }
    v12 = "hello"
    v1 = v1
    v2 = 0.1
    _ = len(v3)
    v25 := GenS5{F0: 0}
    v26 := GenS2{F0: v1, F1: 203, F2: 32767}
    v27 := []int{(v1 & (^(v1))), (v1 >> ((497) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
