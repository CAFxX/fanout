package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint16) uint16 {
    v5 := []int{0}
    p0 = p0
    _ = p0
    _ = v5
    return p0
}

func main() {
    var v1 uint64 = 42
    var v2 int8 = -1
    var v3 int16 = 0
    var v4 int32 = 42
    v4 = v4
    if false {
        if true {
            v3 = (-(v3))
        } else {
            v6 := []int{264, -855, 42, -9223372036854775808}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v6
        }
    } else {
        if (3.14159 == -2.25) {
            v7 := GenS1{F0: v2}
            v8 := GenS2{F0: 32767, F1: (v1 & 0)}
            if (false && (true && (1.5 == -2.25))) {
                v9 := []int{-766, 9223372036854775807, 128}
                fmt.Println(v3, v1, v4)
                _ = v9
            }
            _ = v7
            _ = v8
        }
        v2 -= v2
        fmt.Println(v2, v3, v1)
    }
    v10 := []int{42, 32767, -141}
    if len(v10) > 0 {
        v10[0] = 0
    }
    if false {
        var v11 uint64 = (v1 % ((18446744073709551615) | 1))
        _ = v11
    }
    var v12 uint = 6
    v2 += v2
    v2 += (v2 - (v2 % (((v2 - v2)) | 1)))
    var v13 bool = false
    for fi14 := 0; fi14 < 5; fi14++ {
        _ = len(v10)
        var v15 uint64 = 471
        _ = fi14
        _ = v15
    }
    var v16 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
