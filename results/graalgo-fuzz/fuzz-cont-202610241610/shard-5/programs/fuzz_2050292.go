package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 string, p1 int8) string {
    if (-403 < -686) {
        for fi4 := 0; fi4 < 12; fi4++ {
            v5 := GenS1{F0: (fi4 << ((1) & 63))}
            for fi6 := 0; fi6 < 2; fi6++ {
                p1 = 42
                v7 := GenS2{F0: (p1 / ((p1) | 1))}
                _ = fi6
                _ = v7
            }
            _ = fi4
            _ = v5
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 string = ""
            var v10 float32 = (0.0 - 1.5)
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(p0)
    } else {
        p0 = ((p0 + ((p0 + ("\n\t" + "a")) + ("hello" + ((p0 + p0) + "")))) + p0)
        fmt.Println(p1)
    }
    v11 := GenS3{F0: 9223372036854775807, F1: -643}
    _ = p0
    _ = p1
    _ = v11
    return "hello"
}

func main() {
    var v1 bool = false
    var v2 uint64 = 843
    var v3 float32 = -2.25
    v2 = (^(v2))
    if ((((3.14159 / v3) * float32((0.1 / (0.0 - ((0.0 + 3.14159) * 1.5))))) + v3) > ((v3 + -2.25) * 0.0)) {
        v2 += v2
    } else {
        if ((956 > (v2 << ((0) & 63))) && false) {
            fmt.Println(v2, v3)
        }
    }
    v3 -= (v3 * v3)
    v2 += (v2 << ((4294967295) & 63))
    fmt.Println(v2, v3, v1)
    var v12 float64 = 0.1
    v1 = true
    fmt.Println(v1, v1, v1)
    var v13 int8 = -128
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS2{F0: (v13 & (-(v13)))}
        var v16 int8 = v13
        _ = fi14
        _ = v15
        _ = v16
    }
    switch 0 {
    case 4294967295:
        v17 := []int{128}
        _ = v17
    case 9223372036854775807:
        if v1 {
            var v18 uint32 = 686
            var v19 int = 127
            _ = v18
            _ = v19
        }
    }
    switch "" {
    case "hello":
        v1 = v1
    default:
        v2 = (v2 >> ((18446744073709551615) & 63))
    }
    v13 = v13
    var v20 uint8 = 206
    fmt.Println(v2, v12, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
}
