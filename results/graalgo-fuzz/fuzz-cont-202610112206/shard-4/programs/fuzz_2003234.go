package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint64, p1 bool, p2 int16) bool {
    var v4 int16 = (p2 ^ (p2 << ((65535) & 63)))
    var v5 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p0 > p0)
}
func genFn1(p0 uint, p1 int8, p2 uint64) int8 {
    var v6 int16 = -1
    v7 := []int{-1, -9223372036854775808}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (^(p1))
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{65535}
    var v3 uint8 = 0
    _ = len(v2)
    for fi8 := 0; fi8 < 4; fi8++ {
        v3 -= (v3 / ((215) | 1))
        _ = fi8
    }
    var v9 uint16 = 42
    v1[-2147483648] = "a"
    switch -481 {
    case -9223372036854775808:
        v9 = (v9 ^ (v9 - 65535))
    case 128:
        var v10 uint32 = 2
        _ = v10
    case -901:
        v11 := GenS1{F0: v3, F1: 4294967295}
        _ = v11
    }
    if (-2.25 == (-2.25 + (float64(0.1) * -2.25))) {
        if (v9 < 65535) {
            v9 = v9
            fmt.Println(v9, v3)
            _ = len(v1)
        }
    }
    fmt.Println(v9, v3)
    for fi12 := 0; fi12 < 16; fi12++ {
        v9 += (v9 * uint16(v3))
        var v13 int8 = -128
        _ = fi12
        _ = v13
    }
    switch ("hello" + "foo bar") {
    case "hello":
        _ = len(v1)
    case "a":
        var v14 uint8 = (v3 % ((2) | 1))
        _ = v14
    }
    v9 += (v9 << ((149) & 63))
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
