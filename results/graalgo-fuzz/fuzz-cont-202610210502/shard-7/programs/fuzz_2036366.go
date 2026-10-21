package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int64 = -880
    v2 := []int{-1, 128, 532}
    v3 := []string{}
    v4 := make(map[uint32]bool)
    v1 = 9223372036854775807
    v1 -= (v1 << ((255) & 63))
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v1 = v1
            v7 := GenS1{F0: v1}
            fmt.Println(fi5, v1)
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 42}
        fmt.Println(v1, v1)
        _ = fi5
        _ = v8
    }
    v1 = -179
    switch "hello" {
    case "hello":
        var v9 int16 = 10
        _ = v9
    default:
        for fi10 := 0; fi10 < 19; fi10++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 -= 42
            _ = fi10
        }
    }
    var v11 uint64 = 6
    for fi12 := 0; fi12 < 4; fi12++ {
        v11 += v11
        v11 += (v11 ^ (v11 % ((v11) | 1)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        fmt.Println(v1, v1)
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
