package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int }

func main() {
    var v1 int64 = 1
    var v2 uint = 383
    var v3 int64 = -9223372036854775808
    v4 := []int{-2147483648, 791, -9223372036854775808, 1}
    var v5 int32 = 42
    v6 := GenS1{F0: 68}
    v7 := GenS1{F0: 18446744073709551615}
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    var v8 int32 = v5
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            _ = len(v4)
            v2 -= (^(v2))
            _ = fi10
        }
        if false {
            v11 := GenS1{F0: uint64(v2)}
            for fi12 := 0; fi12 < 17; fi12++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi12
            }
            _ = v11
        }
        v13 := GenS2{F0: 85, F1: 900}
        _ = fi9
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
