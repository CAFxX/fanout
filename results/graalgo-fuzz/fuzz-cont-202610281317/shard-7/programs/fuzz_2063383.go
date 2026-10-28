package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
func genFn0(p0 bool) int16 {
    switch "a" {
    case "":
        var v7 string = ("a" + "a")
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(p0, p0)
            _ = fi8
        }
    default:
        var v9 int64 = -785
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: 42, F2: p0}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return 32767
}
func genFn1(p0 uint64) uint8 {
    if true {
        fmt.Println(p0)
        v11 := []int{}
        p0 = (p0 << ((0) & 63))
        _ = v11
    }
    var v12 string = (("a" + "a") + "foo bar")
    p0 = p0
    _ = p0
    _ = v12
    return 42
}

func main() {
    v1 := []int{32767, 42}
    var v2 uint16 = 2
    v3 := make(map[string]bool)
    v4 := []uint32{42, 4294967295, 148, 540}
    v5 := make(map[uint32]int)
    var v6 uint32 = 2147483647
    fmt.Println(v2, v6)
    fmt.Println(v6)
    _ = len(v4)
    _ = len(v4)
    fmt.Println(v2, v2)
    var v13 uint8 = 135
    var v14 int = 2147483647
    var v15 int = (v14 ^ (v14 * (v14 | v14)))
    v14 = v15
    var v16 uint16 = (v2 % (((v2 ^ 52)) | 1))
    for fi17 := 0; fi17 < 17; fi17++ {
        v13 = (v13 % ((v13) | 1))
        var v18 int8 = -1
        _ = fi17
        _ = v18
    }
    v14 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
