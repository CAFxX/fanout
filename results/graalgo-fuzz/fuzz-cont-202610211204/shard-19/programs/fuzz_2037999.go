package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 uint8 = 42
    var v2 uint32 = 1
    var v3 int16 = 498
    var v4 int32 = 0
    var v5 int16 = (v3 >> ((18446744073709551615) & 63))
    if true {
        var v6 float64 = 1e10
        var v7 int32 = -2147483648
        _ = v6
        _ = v7
    } else {
        var v8 uint64 = 598
        if false {
            fmt.Println(v1)
        }
        _ = v8
    }
    fmt.Println(v5, v3)
    v2 += v2
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float64 = -2.25
        var v11 int64 = -1
        if true {
            v12 := GenS1{F0: 255, F1: v4, F2: (^(v1))}
            if true {
                v13 := []int{-2147483648, fi9}
                v11 = (v11 & (v11 / ((v11) | 1)))
                _ = v13
            }
            if false {
                for fi14 := 0; fi14 < 13; fi14++ {
                    v3 = (-(v5))
                    var v15 string = "foo bar"
                    _ = fi14
                    _ = v15
                }
            }
            _ = v12
        } else {
            fmt.Println(v1)
            var v16 int16 = 0
            _ = v16
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v17 uint32 = v2
    v2 -= 645
    v1 = (v1 | v1)
    v2 += v17
    if false {
        v18 := GenS2{F0: 42, F1: uint64(v17)}
        var v19 float32 = -2.25
        _ = v18
        _ = v19
    } else {
        v3 = v3
        fmt.Println(v3)
        var v20 int64 = -1
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
