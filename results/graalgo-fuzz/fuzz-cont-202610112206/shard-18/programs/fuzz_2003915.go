package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 int32, p2 uint64) int8 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{(p0 * 9223372036854775807)}
        p1 = (p1 << ((65535) & 63))
        var v6 float64 = (-2.25 / float64(3.14159))
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 int32, p2 uint16) uint {
    var v8 int32 = (p1 ^ (p1 >> (((p0 * (p0 | p0))) & 63)))
    var v9 uint8 = 127
    fmt.Println(v8)
    fmt.Println(p2, p1)
    var v10 uint32 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (p0 + p0)
}

func main() {
    var v1 uint64 = 85
    var v2 int32 = -2147483648
    var v3 int32 = (^(v2))
    v2 = (^(v2))
    v11 := GenS2{F0: v1, F1: v1, F2: 100}
    var v12 uint8 = uint8(v1)
    v13 := []int{2147483647, 127}
    v2 += v3
    v14 := GenS2{F0: (v1 / ((v1) | 1)), F1: 606, F2: 0}
    v15 := GenS3{F0: -1}
    if false {
        if len(v13) > 0 {
            v13[0] = 820
        }
        if (!false) {
            var v16 int32 = (v2 % ((v3) | 1))
            _ = v16
        } else {
            v2 -= 1
        }
    }
    v17 := GenS2{F0: 2, F1: (v1 << ((151) & 63)), F2: 42}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
}
