package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint32; F1 int32 }
func genFn0(p0 int8, p1 int32, p2 bool) uint32 {
    p1 -= p1
    var v4 bool = p2
    fmt.Println(p2, v4)
    for fi5 := 0; fi5 < 18; fi5++ {
        p2 = (v4 && v4)
        _ = fi5
    }
    fmt.Println(v4, p1, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}

func main() {
    var v1 bool = false
    var v2 int16 = 905
    var v3 int64 = -9223372036854775808
    var v6 int16 = (v2 - v2)
    var v7 string = "foo bar"
    switch v7 {
    case "hello":
        v6 = (v2 - v2)
    case "":
        switch v7 {
        case "a":
            v8 := []int{127, -1, -56, 2147483647}
            _ = v8
        case "":
            v9 := GenS1{F0: (42 < 1)}
            _ = v9
        default:
            var v10 float32 = (float32(1.5) / 1e10)
            _ = v10
        }
    }
    v11 := GenS1{F0: (false && v1)}
    switch 42 {
    case 127:
        v12 := GenS2{F0: 2147483648, F1: 696}
        _ = v12
    case -1:
        v6 = -32768
    case 678:
        var v13 int16 = v6
        _ = v13
    }
    fmt.Println(v1, v6)
    if (42 >= -128) {
        v6 -= v2
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := GenS3{F0: 2, F1: -12}
            v3 -= (v3 / (((v3 * v3)) | 1))
            _ = fi14
            _ = v15
        }
        fmt.Println(v1, v7, v7)
        var v16 int8 = 127
        _ = v16
    }
    v17 := []int{}
    v17 = append(v17, 1)
    if len(v17) > 0 {
        v17[0] = -9223372036854775808
    }
    v6 += v2
    fmt.Println(v1)
    switch "a" {
    case "hello":
        fmt.Println(v7)
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS3{F0: 1, F1: (fi18 & fi18)}
            _ = fi18
            _ = v19
        }
    default:
        var v20 string = v7
        _ = v20
    }
    switch v7 {
    case "":
        fmt.Println(v3, v6)
    case "foo bar":
        var v21 float32 = (((0.1 - 0.0) / 1.5) / (0.1 * 1e10))
        _ = v21
    default:
        var v22 bool = v1
        _ = v22
    }
    v23 := GenS4{F0: (v3 % ((-482) | 1)), F1: -1}
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS5{F0: 2147483647, F1: 58}
        var v26 bool = v1
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
