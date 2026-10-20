package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int32, p1 int16) int8 {
    p1 += p1
    v6 := []int{141, 32767, 1}
    var v7 bool = false
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 int16 = 0
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    var v4 int = 32768
    v5 := make(map[string]string)
    var v8 int16 = v1
    switch v4 {
    case -9:
        for fi9 := 0; fi9 < 1; fi9++ {
            v3[-1] = true
            _ = fi9
        }
    case -939:
        if true {
            v1 -= v8
        } else {
            fmt.Println(v1)
        }
    default:
        v10 := GenS1{F0: -128, F1: (-(v4))}
        _ = v10
    }
    v11 := GenS2{F0: (^(v4))}
    if true {
        v12 := GenS3{F0: 9223372036854775807}
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    v13 += (v13 + (^(v13)))
    var v14 string = "\n\t"
    v13 += (v13 * 9223372036854775807)
    switch ("" + (v14 + (v14 + (v14 + v14)))) {
    case "\n\t":
        v15 := GenS4{F0: 1, F1: 127, F2: 88}
        _ = v15
    case "hello":
        v1 = v8
    default:
        var v16 uint32 = 847
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
