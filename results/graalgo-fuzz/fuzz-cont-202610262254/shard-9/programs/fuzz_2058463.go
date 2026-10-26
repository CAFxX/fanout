package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"foo bar", "", ("" + "\n\t")}
    var v3 int8 = -72
    v3 = (v3 << ((18446744073709551615) & 63))
    v4 := v1[9223372036854775807]
    v3 = v3
    for fi5 := 0; fi5 < 6; fi5++ {
        switch fi5 {
        case 4294967295:
            v4 = ("a" + v4)
        default:
            if false {
                v3 = v3
                v1[(fi5 << ((808) & 63))] = (v4 + v4)
            }
        }
        _ = fi5
    }
    switch 2147483647 {
    case 969:
        v6 := GenS1{F0: 279, F1: -221}
        _ = v6
    }
    if false {
        v4 = v4
    }
    var v7 uint = 979
    var v8 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
