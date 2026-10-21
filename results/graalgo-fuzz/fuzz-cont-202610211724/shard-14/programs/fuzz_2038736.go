package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16, p2 uint16) string {
    for fi6 := 0; fi6 < 20; fi6++ {
        p2 += (p0 * p2)
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    fmt.Println(p2, p2, p0)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 9223372036854775807}
        p0 = (p2 & (p2 >> ((2) & 63)))
        if false {
            var v10 string = ""
            v11 := GenS2{F0: v10, F1: 2147483647}
            var v12 float32 = (float32((3.14159 - -2.25)) - (0.0 - -2.25))
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p0, p1)
    if false {
        p1 += -1
        p2 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint16 = 65535
    var v3 int64 = 511
    var v4 float32 = 1.5
    var v5 uint8 = 42
    var v13 int = 9223372036854775807
    v14 := GenS2{F0: "a", F1: v1}
    v15 := GenS1{F0: int64(v5)}
    v3 = v3
    var v16 int64 = v3
    v5 = v5
    v17 := GenS3{F0: (v16 >> ((2) & 63))}
    v18 := []int{}
    var v19 float32 = (0.1 * (float32(((0.1 / (float64((v4 / (-2.25 - (((0.1 * v4) - (float32(3.14159) - (v4 - (v4 + v4)))) + (1.5 * 0.1))))) - ((0.1 * 1e10) * (-2.25 / 1e10)))) * (1.5 / 1.0))) / v4))
    v20 := GenS2{F0: "foo bar", F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
