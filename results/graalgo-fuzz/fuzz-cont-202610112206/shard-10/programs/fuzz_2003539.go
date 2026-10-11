package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint) uint16 {
    v5 := []int{505, 127, 32767, -9223372036854775808}
    var v6 int32 = 336
    p0 += (p0 + p0)
    switch "foo bar" {
    case "a":
        if false {
            _ = len(v5)
            var v7 uint = p0
            var v8 uint32 = 2147483647
            _ = v7
            _ = v8
        }
    case "foo bar":
        v5 = append(v5, 32767)
    default:
        v9 := GenS1{F0: -2147483648, F1: 0}
        _ = v9
    }
    v6 = (-(v6))
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 uint = 18446744073709551615
    var v3 float32 = 0.0
    v4 := []uint32{0, 2, 1, 65535}
    v2 += v2
    v2 = 65535
    var v10 uint = 493
    fmt.Println(v3, v10)
    if false {
        var v11 uint8 = 42
        _ = v11
    } else {
        v1 += (^(v1))
        for fi12 := 0; fi12 < 18; fi12++ {
            v4 = append(v4, 4294967295)
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS1{F0: (-(fi12)), F1: 185}
                v10 += v2
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    }
    var v15 int8 = 42
    v1 += v1
    var v16 uint = uint(v1)
    _ = len(v4)
    v17 := GenS2{F0: 298, F1: (v15 * v15)}
    v10 += (v16 - v16)
    var v18 bool = false
    if (v18 && (0.0 == (v3 - v3))) {
        var v19 int16 = 42
        var v20 int16 = v19
        var v21 string = "a"
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 1000, F1: 548}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
