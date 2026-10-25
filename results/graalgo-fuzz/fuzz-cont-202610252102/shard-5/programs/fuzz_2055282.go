package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }
func genFn0(p0 int, p1 int16, p2 int16) uint32 {
    switch "a" {
    case "a":
        var v5 int = 128
        _ = v5
    case "":
        v6 := GenS1{F0: 2, F1: "a", F2: 127}
        _ = v6
    case "foo bar":
        var v7 uint = 0
        _ = v7
    default:
        var v8 bool = (0 <= 2)
        _ = v8
    }
    v9 := GenS1{F0: 818, F1: "", F2: 142}
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(fi10, fi10, p0)
        v11 := []int{}
        _ = len(v11)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    var v1 uint = 430
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    var v4 int8 = -128
    v4 = (v4 | v4)
    v1 = (v1 << ((v1) & 63))
    v4 = (v4 << ((v1) & 63))
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            var v13 uint64 = 501
            _ = v13
        } else {
            v14 := GenS1{F0: 4294967295, F1: "a", F2: 127}
            _ = v14
        }
        _ = fi12
    }
    var v15 string = "a"
    var v16 int32 = 0
    var v17 uint16 = 1
    v18 := GenS2{F0: -2147483648, F1: 128, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
