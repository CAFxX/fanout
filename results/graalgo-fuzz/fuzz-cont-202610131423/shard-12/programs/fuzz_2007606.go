package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
func genFn0() uint16 {
    v4 := []int{-914, 45, -112, 42}
    var v5 float64 = -2.25
    v5 = float64((float32(v5) - 0.0))
    _ = v4
    _ = v5
    return 344
}

func main() {
    var v1 string = "a"
    var v2 int16 = 42
    var v3 int32 = -2147483648
    fmt.Println(v3, v2)
    v6 := GenS1{F0: 0, F1: -45}
    v7 := GenS2{F0: (v3 ^ v3), F1: 253}
    fmt.Println(v1)
    v8 := []int{}
    if true {
        if false {
            _ = len(v8)
            fmt.Println(v2)
        }
    } else {
        switch -2147483648 {
        case 32767:
            v2 = v2
        case -2147483648:
            var v9 int8 = 1
            _ = v9
        default:
            v2 += (v2 % (((v2 << ((1) & 63))) | 1))
        }
        v10 := GenS2{F0: (v3 & (-(v3))), F1: 127}
        v11 := GenS3{F0: 1}
        _ = v10
        _ = v11
    }
    switch "a" {
    case "a":
        v3 = (v3 * (-(v3)))
    case "":
        var v12 int32 = 2147483647
        _ = v12
    case "\n\t":
        switch 314 {
        case 128:
            v2 = v2
        case 413:
            v13 := GenS1{F0: 255, F1: 127}
            _ = v13
        }
    default:
        var v14 int8 = 1
        _ = v14
    }
    if false {
        v15 := GenS4{F0: "hello"}
        switch -236 {
        case -1:
            var v16 int = -9223372036854775808
            _ = v16
        case 2147483647:
            var v17 int64 = 1
            _ = v17
        case -845:
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 int = (^(fi18))
                _ = fi18
                _ = v19
            }
        }
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
