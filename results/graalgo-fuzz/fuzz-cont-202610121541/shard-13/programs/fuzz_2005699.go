package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint8, p1 uint16) uint64 {
    var v5 uint32 = 0
    if (!(float32((((3.14159 * (0.0 + 1e10)) / 1.0) + (1.5 * ((1e10 - 1.5) - (0.0 + float64(1e10)))))) <= 0.0)) {
        p0 = 3
        switch 42 {
        case 42:
            p0 = 219
        case -2147483648:
            p1 += (^(p1))
        case 32767:
            var v6 int = 2147483647
            _ = v6
        default:
            p1 = (p1 | p1)
        }
        p1 = p1
    } else {
        p0 -= (p0 << ((620) & 63))
        v7 := GenS1{F0: 65535, F1: p0, F2: "foo bar"}
        _ = v7
    }
    v8 := []int{}
    v9 := GenS1{F0: p1, F1: p0, F2: "foo bar"}
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    _ = v9
    return 723
}
func genFn1(p0 uint8, p1 int8) string {
    if false {
        v10 := GenS2{F0: 1, F1: 210}
        if false {
            var v11 string = ""
            var v12 int64 = -165
            var v13 int16 = 1000
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := []int{0, -248, 4294967295, 0}
            p1 += p1
            _ = v14
        }
        fmt.Println(p0)
        _ = v10
    } else {
        p0 = p0
    }
    var v15 uint16 = 0
    _ = p0
    _ = p1
    _ = v15
    return ("a" + "")
}

func main() {
    var v1 int8 = 100
    var v2 float64 = (0.1 / 1.0)
    var v3 int32 = 0
    var v4 float32 = (3.14159 + float32((0.1 + (v2 * (v2 * (float64(3.14159) - 0.0))))))
    var v16 float64 = ((3.14159 * 1.5) - 1e10)
    v4 += v4
    var v17 int16 = 1
    v18 := []int{128, 128, 32768, -1}
    for fi19 := 0; fi19 < 13; fi19++ {
        if len(v18) > 0 {
            v18[0] = 274
        }
        switch "\n\t" {
        case "\n\t":
            v2 = (v16 + ((v16 - (1.5 - (v16 - v16))) * float64(v4)))
        case "hello":
            fmt.Println(v2, v16, v4)
        default:
            var v20 bool = ("a" < "foo bar")
            _ = v20
        }
        _ = fi19
    }
    var v21 uint16 = 489
    v21 -= (v21 + 255)
    v22 := GenS2{F0: 2147483648, F1: 104}
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            v4 -= v4
            _ = fi24
        }
        _ = fi23
    }
    switch -838 {
    case 42:
        if ((v17 ^ v17) >= 42) {
            v4 = v4
            v1 += v1
            fmt.Println(v2, v4)
        }
    case -9223372036854775808:
        var v25 int = -1
        _ = v25
    case -495:
        var v26 int32 = v3
        _ = v26
    default:
        v3 = int32(v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
