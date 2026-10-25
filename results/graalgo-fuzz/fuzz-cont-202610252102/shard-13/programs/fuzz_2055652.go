package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
func genFn0(p0 uint16) uint64 {
    v5 := []int{}
    p0 -= p0
    _ = p0
    _ = v5
    return 919
}
func genFn1() uint64 {
    v6 := GenS1{F0: -1}
    v7 := []int{}
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float32 = (3.14159 * (((-2.25 + (0.1 * 1.5)) * -2.25) - (1.5 / 1.0)))
    var v2 uint16 = 1
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    switch -1 {
    case 4294967295:
        var v8 float64 = 3.14159
        _ = v8
    case 128:
        _ = len(v3)
    case 127:
        v9 := GenS2{F0: "a", F1: 415}
        _ = v9
    default:
        v1 = v1
    }
    fmt.Println(v2, v2)
    delete(v3, 1)
    if (157 != 255) {
        v1 = float32(1e10)
    }
    if (42 < 127) {
        v1 = 3.14159
    }
    if false {
        switch ("\n\t" + (("hello" + "") + "")) {
        case "\n\t":
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := GenS3{F0: "\n\t"}
                _ = fi10
                _ = v11
            }
        }
    }
    if (127 < 42) {
        v12 := GenS3{F0: "\n\t"}
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v2 -= (v2 * 255)
        v14 := GenS4{F0: (fi13 - fi13)}
        fmt.Println(v1)
        _ = fi13
        _ = v14
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
