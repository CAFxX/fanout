package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 int64; F2 uint }

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 42
    v3 := []uint32{}
    var v4 uint8 = (v2 % ((199) | 1))
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 1
    v7 := GenS1{F0: 100, F1: 620, F2: ("" + "hello")}
    v1 = v1
    switch "foo bar" {
    case "":
        v8 := GenS2{F0: 42, F1: -223, F2: 42}
        _ = v8
    default:
        switch -470 {
        case 4294967295:
            v9 := GenS2{F0: 1, F1: 42, F2: 65535}
            _ = v9
        }
    }
    v10 := GenS2{F0: 1, F1: 0, F2: 536}
    v11 := GenS1{F0: -128, F1: -9223372036854775808, F2: ("hello" + "\n\t")}
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS1{F0: 1, F1: -583, F2: "a"}
            _ = fi13
            _ = v14
        }
        var v15 float64 = v1
        _ = fi12
        _ = v15
    }
    switch "a" {
    case "hello":
        v5[uint32(v6)] = 847
    }
    var v16 uint = 302
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
