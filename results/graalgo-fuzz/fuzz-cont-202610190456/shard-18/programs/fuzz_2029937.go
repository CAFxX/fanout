package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint32; F1 int64; F2 int }
func genFn0(p0 int16, p1 string) uint {
    v4 := []int{32767, -941}
    var v5 float32 = 1e10
    v4 = append(v4, 127)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 uint32 = 893
    var v2 string = "a"
    v3 := []int{1, -36}
    if true {
        v1 = (v1 % ((306) | 1))
        var v6 uint8 = 201
        _ = v6
    }
    v1 = (v1 | 2147483648)
    v2 = ""
    if true {
        var v7 bool = (2147483647 > 0)
        v3 = append(v3, 42)
        _ = v7
    }
    _ = len(v3)
    var v8 int64 = -9223372036854775808
    v2 = v2
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            if false {
                v1 += v1
                var v11 uint64 = 321
                fmt.Println(v1, fi9, fi10)
                v12 := GenS1{F0: 1}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    switch "a" {
    case "a":
        v13 := GenS2{F0: 520, F1: v2}
        _ = v13
    default:
        v14 := GenS3{F0: 4294967295, F1: (v8 - (v8 - (v8 >> ((18446744073709551615) & 63)))), F2: 127}
        _ = v14
    }
    v15 := GenS3{F0: v1, F1: -9223372036854775808, F2: 2147483647}
    fmt.Println(v8)
    var v16 float64 = 1.5
    var v17 int64 = -9223372036854775808
    v18 := GenS3{F0: v1, F1: (v17 / ((v8) | 1)), F2: 9223372036854775807}
    var v19 int = 830
    fmt.Println(v8, v1, v2)
    if len(v3) > 0 {
        v3[0] = (v19 - v19)
    }
    var v20 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
