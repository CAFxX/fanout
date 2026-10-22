package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint }

func main() {
    var v1 string = "a"
    var v2 float32 = 0.0
    v3 := make(map[int]bool)
    var v4 int64 = -725
    v5 := []int{-2147483648, 127, 9223372036854775807}
    switch 1 {
    case -2147483648:
        if false {
            v6 := GenS1{F0: 515, F1: -9223372036854775808}
            _ = v6
        }
    case 128:
        fmt.Println(v2, v1)
    case 9223372036854775807:
        v7 := GenS2{F0: 42}
        _ = v7
    }
    fmt.Println(v4, v2)
    var v8 uint64 = 2
    var v9 string = "hello"
    v2 += (v2 - (v2 - float32(0.0)))
    var v10 int64 = (v4 & v4)
    if false {
        var v11 int16 = -770
        var v12 float64 = (1.5 * 3.14159)
        var v13 int32 = -562
        _ = v11
        _ = v12
        _ = v13
    } else {
        if ((v10 ^ (v10 % ((v4) | 1))) > v10) {
            for fi14 := 0; fi14 < 16; fi14++ {
                v4 = v4
                for fi15 := 0; fi15 < 11; fi15++ {
                    v4 += (v4 + v10)
                    _ = fi15
                }
                var v16 float64 = float64(3.14159)
                _ = fi14
                _ = v16
            }
        }
    }
    v17 := GenS2{F0: 65535}
    v18 := GenS3{F0: 32768, F1: 9223372036854775807}
    var v19 uint16 = uint16(v8)
    v20 := GenS3{F0: 0, F1: -584}
    if false {
        v4 -= (v10 * (-(v10)))
    }
    if true {
        v21 := GenS4{F0: -57, F1: 2147483647, F2: 65535}
        _ = v21
    }
    switch 397 {
    case 127:
        fmt.Println(v10, v8)
    }
    var v22 int = int(v8)
    v23 := GenS5{F0: v4, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
