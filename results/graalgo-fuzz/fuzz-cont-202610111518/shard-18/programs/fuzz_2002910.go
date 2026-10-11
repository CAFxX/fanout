package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
func genFn0(p0 uint32, p1 int16) uint {
    v5 := []int{-2147483648, 94, 128}
    fmt.Println(p1)
    p1 += (p1 / (((p1 + (p1 - 42))) | 1))
    var v6 int8 = 1
    p1 += p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    var v3 uint8 = 103
    v4 := make(map[string]uint32)
    var v7 uint = v1
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            if false {
                v10 := GenS1{F0: 127, F1: 42, F2: "a"}
                _ = v10
            } else {
                var v11 uint32 = (v8 & v8)
                _ = v11
            }
        } else {
            var v12 string = "\n\t"
            switch v12 {
            case "foo bar":
                fmt.Println(v7, fi9, v1)
            case "\n\t":
                fmt.Println(v3, v3)
            case "a":
                v13 := v2["foo bar"]
                _ = v13
            default:
                if (1 > 9223372036854775807) {
                    v14 := v4[v12]
                    _ = v14
                }
            }
            fmt.Println(v12, v12)
            _ = v12
        }
        _ = fi9
    }
    switch "a" {
    case "hello":
        var v15 float32 = (-2.25 + -2.25)
        _ = v15
    }
    var v16 int8 = 1
    v16 += (v16 * -29)
    v4["\n\t"] = (^(v8))
    var v17 bool = false
    fmt.Println(v8)
    var v18 float64 = (0.0 * ((1.5 * 3.14159) * float64(1.5)))
    var v19 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
