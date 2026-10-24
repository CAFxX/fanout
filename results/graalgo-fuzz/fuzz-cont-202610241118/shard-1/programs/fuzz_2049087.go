package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int }
type GenS2 struct { F0 int; F1 string; F2 int }
func genFn0(p0 uint64, p1 uint8, p2 int) uint {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: (p1 | p1), F1: 552, F2: fi6}
        v8 := []int{4294967295, (p2 | 9223372036854775807), p2, 32768}
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 += 22
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1() int64 {
    var v9 int8 = 100
    switch ("hello" + (("\n\t" + "hello") + "")) {
    case "hello":
        v9 += -128
    case "a":
        fmt.Println(v9)
    }
    v9 -= (v9 >> ((4294967295) & 63))
    v9 += (v9 - (v9 ^ v9))
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := []uint32{572}
    v2 := []uint32{2, 244, 2147483647}
    var v3 int8 = 1
    var v4 uint16 = 1
    var v5 uint32 = 42
    v3 += v3
    if (-1 != 0) {
        v2 = append(v2, v5)
    } else {
        var v10 int32 = -1
        v11 := GenS2{F0: 587, F1: "\n\t", F2: -498}
        _ = v10
        _ = v11
    }
    var v12 int = 911
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 += v5
    _ = len(v1)
    if (1 <= (v12 >> ((42) & 63))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 float64 = 3.14159
        _ = v13
    }
    v2 = append(v2, (v5 * v5))
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
