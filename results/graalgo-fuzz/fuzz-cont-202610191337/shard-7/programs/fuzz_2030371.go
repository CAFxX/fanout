package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 int32 = 148
    switch -290 {
    case 128:
        for fi5 := 0; fi5 < 16; fi5++ {
            v4 -= (v4 | 42)
            switch "a" {
            case "":
                _ = len(v2)
            case "\n\t":
                fmt.Println(fi5, fi5, v4)
            default:
                fmt.Println(v4)
            }
            v6 := GenS1{F0: -983, F1: 0, F2: ""}
            _ = fi5
            _ = v6
        }
    case 1:
        switch "\n\t" {
        case "a":
            v7 := GenS1{F0: 772, F1: 4294967295, F2: "a"}
            _ = v7
        case "hello":
            var v8 uint = 348
            _ = v8
        default:
            var v9 int = -568
            _ = v9
        }
    }
    v4 += v4
    switch 892 {
    case -9223372036854775808:
        var v10 uint32 = 2
        _ = v10
    case 586:
        v4 -= v4
    default:
        v11 := []int{}
        _ = v11
    }
    if true {
        var v12 uint64 = 1
        _ = v12
    }
    v4 = (v4 << ((v1) & 63))
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    v4 = 2147483647
    v15 := GenS2{F0: 1000}
    var v16 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
