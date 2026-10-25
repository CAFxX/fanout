package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 uint8, p1 int8, p2 string) int32 {
    v4 := GenS1{F0: 2, F1: 42, F2: p0}
    var v5 uint32 = 42
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -447
}

func main() {
    var v1 int16 = 921
    var v2 int8 = -128
    v3 := make(map[uint32]int)
    fmt.Println(v1)
    switch 240 {
    case -524:
        if (v1 <= (v1 ^ v1)) {
            v6 := GenS2{F0: 4294967295, F1: v2, F2: v1}
            v2 = -128
            _ = v6
        }
    }
    switch "a" {
    case "foo bar":
        v7 := GenS1{F0: 913, F1: 1, F2: 255}
        _ = v7
    case "":
        v8 := GenS3{F0: ""}
        _ = v8
    case "hello":
        var v9 bool = true
        _ = v9
    }
    fmt.Println(v1)
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: 819, F1: (v2 >> ((30) & 63)), F2: -32768}
        _ = fi10
        _ = v11
    }
    v12 := []int{-130, 32767, 1}
    v2 += -62
    v2 -= (v2 * (v2 - v2))
    fmt.Println(v2, v2)
    if true {
        v1 -= (v1 % ((601) | 1))
        if true {
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: 2, F1: v2, F2: v1}
                _ = fi13
                _ = v14
            }
            var v15 uint16 = 32767
            v1 += int16(v15)
            _ = v15
        }
    }
    switch ("a" + ("" + (("\n\t" + "hello") + ""))) {
    case "hello":
        var v16 bool = false
        _ = v16
    case "foo bar":
        v2 += (v2 + v2)
    }
    v1 -= (-(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
