package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint }
func genFn0() string {
    var v4 int16 = -32768
    v4 -= (v4 | (v4 ^ -1))
    var v5 int = -2147483648
    var v6 int64 = int64(v5)
    _ = v4
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 int64, p1 uint, p2 string) uint8 {
    v7 := GenS1{F0: -170, F1: 218}
    fmt.Println(p1, p2, p2)
    v8 := GenS1{F0: 9223372036854775807, F1: p1}
    p1 += (p1 - 597)
    if true {
        v9 := GenS1{F0: -2147483648, F1: (^(p1))}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 102
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]string)
    var v3 int16 = -32768
    v3 -= v3
    if false {
        fmt.Println(v3)
    } else {
        v10 := GenS1{F0: 32767, F1: 65535}
        _ = v10
    }
    v11 := GenS2{F0: 828}
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := []int{}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := v2[(v1 * 65535)]
    v16 := GenS2{F0: 2147483647}
    v1 = 2147483648
    var v17 bool = ((false || false) && (("hello" >= "hello") && true))
    if (2 != 42) {
        var v18 float64 = 0.1
        _ = v18
    }
    v19 := GenS3{F0: 91, F1: 0}
    v15 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
