package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int32 = 0
    var v2 int32 = 0
    var v3 int = 127
    v1 = v1
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(fi4, v1)
        var v5 int16 = 32767
        _ = fi4
        _ = v5
    }
    var v6 bool = (-2.25 >= (3.14159 / ((0.0 / 1.0) - ((-2.25 - 1.5) * (1.5 / ((-2.25 * (0.1 + ((1e10 + (-2.25 - 3.14159)) - (3.14159 / (-2.25 - (1e10 / 1.0)))))) / 1.0))))))
    fmt.Println(v1, v2)
    v7 := GenS1{F0: 9223372036854775807, F1: (-(v1)), F2: -32768}
    v3 = v3
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{-362, v3, (v3 % (((v3 - (-(fi8)))) | 1))}
        fmt.Println(v2, v1)
        var v10 int32 = (-(v1))
        _ = fi8
        _ = v10
        _ = v9
    }
    if (v6 || true) {
        switch v3 {
        case 396:
            for fi11 := 0; fi11 < 12; fi11++ {
                v2 -= v1
                _ = fi11
            }
        case 1:
            fmt.Println(v3, v6)
        case -13:
            v2 -= v1
        default:
            var v12 float64 = ((-2.25 - (-2.25 * -2.25)) * ((1.5 * -2.25) * float64(1.5)))
            _ = v12
        }
        var v13 int = (v3 % ((127) | 1))
        _ = v13
    } else {
        switch v3 {
        case 32767:
            fmt.Println(v6, v3)
        case -1:
            if true {
                var v14 uint = 18446744073709551615
                _ = v14
            }
        case 2147483647:
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS2{F0: (v2 >> ((42) & 63)), F1: 174}
                _ = fi15
                _ = v16
            }
        default:
            v1 = -2147483648
        }
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v2, fi17)
            _ = fi17
        }
    }
    if (870 < 255) {
        fmt.Println(v6, v3)
    }
    v18 := []int{v3, (v3 & -885), v3}
    if v6 {
        if v6 {
            var v19 bool = true
            _ = v19
        } else {
            var v20 int64 = 543
            var v21 int = -9223372036854775808
            v1 -= (v1 | v2)
            _ = v20
            _ = v21
        }
    } else {
        var v22 bool = (v1 == (v2 % ((v2) | 1)))
        _ = v22
    }
    switch "\n\t" {
    case "":
        v6 = false
    default:
        v18 = append(v18, (v3 >> ((770) & 63)))
    }
    switch (v3 % ((v3) | 1)) {
    case 42:
        fmt.Println(v3, v6, v6)
    case -2147483648:
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 int16 = 32767
            v18 = append(v18, fi23)
            v25 := GenS1{F0: 9223372036854775807, F1: (v2 ^ 1), F2: v24}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    v6 = (!v6)
    var v26 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v6
    _ = v7
}
