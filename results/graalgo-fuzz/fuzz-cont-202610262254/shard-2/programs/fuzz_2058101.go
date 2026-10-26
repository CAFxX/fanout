package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]int)
    var v3 int64 = -394
    var v4 string = "a"
    switch 2147483647 {
    case 4294967295:
        var v5 uint = 2
        _ = v5
    case 32768:
        var v6 int32 = -756
        _ = v6
    case 9223372036854775807:
        fmt.Println(v3)
    }
    v4 = "\n\t"
    v7 := GenS1{F0: 425}
    fmt.Println(v3)
    for fi8 := 0; fi8 < 3; fi8++ {
        v4 = ("foo bar" + v4)
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v1) > 0 {
                v1[0] = v4
            }
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(v3)
    fmt.Println(v3, v4)
    var v10 uint8 = 2
    var v11 int = 454
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS2{F0: 1, F1: (fi13 % ((128) | 1)), F2: v10}
            v15 := GenS2{F0: uint64(v10), F1: (^(fi13)), F2: 255}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
