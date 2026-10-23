package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int; F2 int }
type GenS3 struct { F0 int; F1 uint32; F2 int64 }

func main() {
    var v1 float64 = (1e10 * (((3.14159 - 3.14159) + -2.25) / (((((3.14159 / 1.0) / 1e10) - ((1.5 * 0.1) - 1.5)) + float64((float32(1.5) * 0.1))) * ((0.1 + (0.0 / 3.14159)) * (3.14159 + ((0.1 * 0.1) + (-2.25 - (((1.5 * 1e10) - ((1e10 - ((1e10 * (0.0 + 3.14159)) * (1e10 + 3.14159))) - 1e10)) / 1.0))))))))
    var v2 float64 = (v1 + ((-2.25 * v1) - v1))
    v3 := []uint32{410, 146}
    v4 := []int{9223372036854775807, -2147483648}
    v5 := []uint32{}
    fmt.Println(v1, v1)
    v6 := GenS1{F0: (v1 <= -2.25)}
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            var v8 float64 = v2
            v2 -= (v2 + float64(3.14159))
            _ = v8
        }
        v1 -= v1
        var v9 uint64 = 1
        _ = fi7
        _ = v9
    }
    if true {
        v5 = append(v5, 586)
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint64 = 259
            if true {
                v12 := GenS2{F0: 100, F1: (fi10 << ((585) & 63)), F2: fi10}
                v13 := GenS3{F0: (fi10 >> ((42) & 63)), F1: 4294967295, F2: 1}
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    _ = len(v5)
    v14 := GenS1{F0: true}
    v1 = v1
    if false {
        fmt.Println(v1)
        var v15 float64 = v1
        _ = v15
    } else {
        v1 += 0.0
    }
    var v16 string = (("a" + "foo bar") + ("a" + ("a" + "")))
    var v17 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
