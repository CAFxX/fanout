package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint }
func genFn0(p0 uint16) string {
    v5 := []int{2147483647}
    var v6 int16 = 1
    _ = p0
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 bool, p1 int32, p2 int16) uint {
    v7 := []int{-9223372036854775808, int(p2), 73}
    fmt.Println(p2, p0)
    for fi8 := 0; fi8 < 20; fi8++ {
        if p0 {
            var v9 uint = 147
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint16 = 175
                p2 -= (p2 - int16(p1))
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            _ = len(v7)
            for fi12 := 0; fi12 < 1; fi12++ {
                v13 := GenS1{F0: p1, F1: -1}
                v14 := GenS2{F0: 207, F1: "a", F2: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
            p2 -= p2
        }
        fmt.Println(p0)
        p2 = (p2 << ((65535) & 63))
        _ = fi8
    }
    p0 = (p0 && (411 != 65535))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 577
}

func main() {
    v1 := []int{-2147483648}
    var v2 uint16 = 46
    var v3 bool = (false && (-1 != -1))
    var v4 int16 = -32768
    var v15 float32 = (1.5 / -2.25)
    v1 = append(v1, 368)
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS2{F0: 204, F1: "a", F2: 2}
        var v18 int = (^(fi16))
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int16 = v4
    fmt.Println(v19, v2)
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: int32(fi20), F1: 0}
        _ = fi20
        _ = v21
    }
    _ = len(v1)
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
