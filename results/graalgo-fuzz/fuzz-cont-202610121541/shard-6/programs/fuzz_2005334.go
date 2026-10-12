package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }

func main() {
    var v1 int32 = 42
    var v2 int16 = 42
    var v3 float32 = float32(0.0)
    var v4 float32 = (((v3 - 0.1) * 1e10) - v3)
    v3 = float32(float64(-2.25))
    switch "a" {
    case "hello":
        v5 := []int{127, 127}
        _ = v5
    case "a":
        v6 := GenS1{F0: (-(v2)), F1: v1}
        _ = v6
    case "\n\t":
        v4 += 1e10
    default:
        v3 = 1.5
    }
    v7 := []int{4294967295, -801, 9223372036854775807}
    v7 = append(v7, 0)
    for fi8 := 0; fi8 < 10; fi8++ {
        if (("\n\t" >= ("" + ("\n\t" + "foo bar"))) && (!(-1 != -9223372036854775808))) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v3 += 3.14159
        } else {
            fmt.Println(v1)
            if (false && false) {
                var v9 int32 = (v1 + v1)
                _ = v9
            } else {
                var v10 uint64 = 0
                v4 += (-2.25 * v4)
                var v11 int64 = -904
                _ = v10
                _ = v11
            }
            v2 -= v2
        }
        _ = fi8
    }
    v3 -= float32(0.1)
    var v12 uint32 = 145
    v1 -= v1
    if (!("a" <= ("a" + "\n\t"))) {
        var v13 float64 = -2.25
        v3 = float32((float64(0.0) - v13))
        _ = v13
    }
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    var v14 int32 = v1
    v2 += -146
    var v15 string = ("" + "a")
    var v16 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
