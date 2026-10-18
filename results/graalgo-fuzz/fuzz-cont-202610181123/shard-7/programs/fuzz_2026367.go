package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() uint16 {
    v4 := []int{2147483647, -2147483648}
    var v5 bool = (-13 >= -1)
    for fi6 := 0; fi6 < 12; fi6++ {
        _ = len(v4)
        _ = fi6
    }
    _ = v4
    _ = v5
    return 791
}
func genFn1(p0 uint8, p1 uint16, p2 int) uint32 {
    p2 -= (p2 ^ -2147483648)
    p2 += p2
    v7 := []int{4294967295, p2, 144}
    switch 2147483647 {
    case -9223372036854775808:
        p1 = (p1 ^ genFn0())
    }
    switch "a" {
    case "\n\t":
        v7 = append(v7, p2)
    case "hello":
        var v8 int32 = -1
        _ = v8
    case "":
        for fi9 := 0; fi9 < 20; fi9++ {
            _ = len(v7)
            _ = fi9
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 390
}

func main() {
    var v1 uint8 = 187
    var v2 int = 32767
    var v3 uint8 = 47
    var v10 uint64 = 1
    v11 := GenS1{F0: -38}
    fmt.Println(v1, v10)
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: (!true), F1: genFn1((v1 + v1), 2, fi12), F2: 0}
        var v14 int32 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 += -18
    v15 := []int{(v2 ^ v2)}
    v16 := GenS3{F0: 2, F1: 2147483647}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v17 float64 = 3.14159
    var v18 uint32 = 65535
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
