package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32) uint8 {
    var v5 uint8 = 111
    v6 := []int{-2147483648}
    v5 -= v5
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 bool = (false && false)
    var v2 int64 = 0
    var v3 bool = v1
    v4 := make(map[uint32]bool)
    v1 = (!false)
    fmt.Println(v3, v3)
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 float32 = 0.0
            var v11 uint8 = genFn0(-2147483648)
            _ = fi9
            _ = v10
            _ = v11
        }
        v3 = v1
        v7 += 914
        _ = fi8
    }
    fmt.Println(v2, v7, v1)
    v12 := GenS1{F0: 18}
    v13 := GenS1{F0: 2147483647}
    if v1 {
        v14 := []int{-1, -2147483648}
        _ = v14
    } else {
        if (-640 >= 350) {
            if (32767 <= 0) {
                v1 = (471 > 1)
            }
            if v1 {
                v2 = 9223372036854775807
            } else {
                var v15 int32 = 0
                v16 := v4[443]
                v17 := GenS2{F0: (v2 | (v2 | -9223372036854775808)), F1: 862}
                _ = v15
                _ = v16
                _ = v17
            }
            v4[1] = ((v2 + v2) == -9223372036854775808)
        }
    }
    v4[4294967295] = (3.14159 == float32(0.0))
    if (!(417 != 32767)) {
        v18 := GenS1{F0: 0}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 int8 = 14
            _ = fi19
            _ = v20
        }
        v7 = (v7 % ((v7) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
