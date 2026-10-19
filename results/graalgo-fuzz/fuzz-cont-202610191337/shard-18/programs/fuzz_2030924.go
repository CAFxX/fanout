package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 uint64, p1 string, p2 int) uint32 {
    var v5 int = (p2 << ((255) & 63))
    var v6 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8, p2 uint16) int {
    fmt.Println(p1)
    fmt.Println(p0, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 45
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -159
    v3 := []string{("a" + "foo bar"), ("" + "\n\t"), "hello"}
    v4 := make(map[uint32]string)
    switch ("hello" + "foo bar") {
    case "\n\t":
        v2 += (^(v2))
    case "a":
        var v7 uint64 = 0
        _ = v7
    case "":
        fmt.Println(v2)
    }
    v2 -= (v2 >> ((1) & 63))
    v2 += (-(v2))
    if false {
        if false {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int32 = -768
                var v10 uint8 = 2
                v9 -= (v9 << ((2) & 63))
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        var v11 int64 = -419
        _ = v11
    }
    var v12 string = ("a" + "hello")
    v13 := GenS1{F0: 65535}
    v14 := GenS1{F0: 18446744073709551615}
    v4[17] = v12
    fmt.Println(v12, v2)
    if (((v12 + (v12 + v12)) + v12) < v12) {
        var v15 int32 = -2147483648
        var v16 float32 = -2.25
        var v17 int64 = 1
        _ = v15
        _ = v16
        _ = v17
    } else {
        fmt.Println(v12)
    }
    if true {
        v18 := GenS2{F0: "foo bar", F1: 42}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
