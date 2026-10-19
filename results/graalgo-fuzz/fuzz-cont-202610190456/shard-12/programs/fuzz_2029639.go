package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 uint; F2 int16 }
func genFn0(p0 int, p1 uint64, p2 int8) uint64 {
    switch "a" {
    case "\n\t":
        var v4 int = (p0 | (p0 + p0))
        _ = v4
    default:
        var v5 int = p0
        _ = v5
    }
    switch (^(p0)) {
    case -464:
        v6 := []int{p0, p0, -152}
        _ = v6
    }
    var v7 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (^(v7))
}
func genFn1(p0 uint) uint {
    switch -597 {
    case -706:
        v8 := GenS1{F0: 719, F1: int32(p0), F2: -128}
        _ = v8
    case 127:
        p0 = 18446744073709551615
    case 699:
        v9 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: -9}
        _ = v9
    }
    v10 := []int{-109, -2147483648, -1, 127}
    if false {
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                if false {
                    var v13 int16 = 1
                    var v14 uint = 255
                    _ = v13
                    _ = v14
                }
                var v15 float32 = 3.14159
                _ = fi12
                _ = v15
            }
            _ = fi11
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    p0 += (^(p0))
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 864
    var v3 bool = (((1 == 42) && false) && (false && true))
    var v16 uint8 = 2
    var v17 int64 = 9223372036854775807
    var v18 int32 = 1
    v2 = v2
    var v19 int16 = -32768
    var v20 int64 = v17
    v21 := GenS2{F0: 255, F1: 56, F2: (v16 ^ v16)}
    var v22 bool = v3
    v19 = 1
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS3{F0: 1, F1: 255, F2: v19}
        var v25 float64 = 0.0
        var v26 float32 = (3.14159 + 0.0)
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: -1, F2: (v16 + v16)}
    for fi28 := 0; fi28 < 4; fi28++ {
        if v3 {
            var v29 float32 = 1.5
            _ = v29
        }
        v18 = (v18 - (v18 & (v18 & int32(v17))))
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
}
