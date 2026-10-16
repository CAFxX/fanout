package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0(p0 string, p1 uint8, p2 uint) uint64 {
    p1 = (p1 & (^(p1)))
    switch 0 {
    case 42:
        v7 := []int{}
        _ = v7
    }
    v8 := []int{127}
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}
func genFn1(p0 uint32, p1 uint64, p2 uint64) int32 {
    p2 = (p1 - 4294967295)
    v9 := []int{9223372036854775807, 9223372036854775807, 42}
    switch 1 {
    case -745:
        _ = len(v9)
    }
    var v10 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -497
}

func main() {
    var v1 int8 = -128
    var v2 int64 = 938
    v3 := []uint32{}
    v4 := []int{-2147483648, 2147483647, -761, 128}
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]bool)
    if ((true && true) || true) {
        v11 := GenS1{F0: 1, F1: (v1 ^ (v1 * v1))}
        _ = v11
    }
    v2 -= (v2 | (-(v2)))
    for fi12 := 0; fi12 < 10; fi12++ {
        v5[1] = false
        var v13 uint16 = 65535
        _ = fi12
        _ = v13
    }
    var v14 int64 = v2
    fmt.Println(v2, v1, v1)
    var v15 uint64 = 1
    fmt.Println(v1)
    v14 += 0
    fmt.Println(v14, v14)
    var v16 int32 = genFn1(2, (^(v15)), v15)
    if true {
        var v17 float64 = -2.25
        var v18 uint16 = 65535
        v1 -= (v1 + 127)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v6[42] = (true && false)
        _ = fi19
    }
    switch "\n\t" {
    case "":
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint = 642
            fmt.Println(v1)
            v16 += -1
            _ = fi20
            _ = v21
        }
    case "\n\t":
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint32 = 217
            v24 := GenS1{F0: uint64(v23), F1: v1}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v1, v15)
        _ = len(v6)
        var v26 int64 = (v2 << ((255) & 63))
        _ = fi25
        _ = v26
    }
    v14 = 1
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 bool = false
            v16 = (-(v16))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi28
            _ = v29
        }
        if true {
            var v30 uint16 = 0
            _ = v30
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
