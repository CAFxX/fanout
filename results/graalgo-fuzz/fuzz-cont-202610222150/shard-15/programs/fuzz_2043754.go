package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint8, p2 int16) int {
    var v7 uint = 0
    var v8 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32768
}
func genFn1(p0 uint, p1 uint32, p2 int) int {
    v9 := []int{p2, p2}
    switch (p2 ^ (p2 >> ((42) & 63))) {
    case 127:
        fmt.Println(p2)
    case 9223372036854775807:
        for fi10 := 0; fi10 < 11; fi10++ {
            v11 := GenS1{F0: 96}
            p1 -= p1
            for fi12 := 0; fi12 < 1; fi12++ {
                switch fi12 {
                case 4294967295:
                    fmt.Println(p0, p2, p0)
                case 719:
                    v9 = append(v9, (p2 ^ (-(p2))))
                }
                v13 := GenS1{F0: int8(fi10)}
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
    }
    switch "foo bar" {
    case "hello":
        p2 = genFn0(uint32(p0), 84, -1)
    case "foo bar":
        p2 = p2
    }
    var v14 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v9
    return p2
}

func main() {
    var v1 uint64 = 4294967295
    v2 := []uint32{4294967295, 2147483648}
    var v3 uint8 = 17
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    var v6 int64 = 9223372036854775807
    for fi15 := 0; fi15 < 20; fi15++ {
        if (-599 != 0) {
            fmt.Println(v6, v1, fi15)
        }
        _ = fi15
    }
    fmt.Println(v1)
    v1 -= (v1 | (v1 - v1))
    _ = len(v2)
    var v16 int32 = -1
    v17 := v5[4294967295]
    fmt.Println(v6, v17)
    var v18 int64 = 42
    var v19 uint16 = 435
    _ = len(v2)
    switch 494 {
    case -322:
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint = 4294967295
            for fi22 := 0; fi22 < 4; fi22++ {
                _ = len(v5)
                _ = fi22
            }
            _ = fi20
            _ = v21
        }
    case 127:
        var v23 uint8 = v3
        _ = v23
    default:
        var v24 uint16 = 95
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
