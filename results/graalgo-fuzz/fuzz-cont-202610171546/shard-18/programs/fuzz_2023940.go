package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 int64, p1 uint32) int8 {
    p0 = 9223372036854775807
    for fi6 := 0; fi6 < 3; fi6++ {
        switch (^(fi6)) {
        case 1:
            v7 := GenS1{F0: 0}
            _ = v7
        case -593:
            var v8 string = "\n\t"
            _ = v8
        default:
            v9 := GenS2{F0: 1, F1: 42}
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 101
    v3 := []uint32{906, 2147483647}
    var v4 uint = 0
    v5 := []int{457}
    for fi10 := 0; fi10 < 17; fi10++ {
        switch fi10 {
        case 128:
            var v11 bool = false
            _ = v11
        case 43:
            fmt.Println(fi10, fi10, fi10)
        }
        _ = fi10
    }
    v2 -= (v2 % ((127) | 1))
    if (!true) {
        var v12 uint32 = 65535
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        switch "foo bar" {
        case "hello":
            v4 = (^(v4))
        default:
            var v14 int16 = 1000
            _ = v14
        }
        var v15 uint64 = 42
        _ = fi13
        _ = v15
    }
    var v16 uint64 = 2
    v2 = uint8(v4)
    var v17 uint16 = 65535
    v16 = v16
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
