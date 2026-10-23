package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 int32 = 2147483647
    v2 := []int{-172, -9223372036854775808, 4294967295, -374}
    var v3 int64 = 1
    v4 := make(map[int]string)
    v1 += (v1 << ((2) & 63))
    for fi5 := 0; fi5 < 12; fi5++ {
        v3 = v3
        v6 := GenS1{F0: v3}
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 868
    fmt.Println(v7)
    var v8 float32 = 1e10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 string = "foo bar"
    switch 32767 {
    case 4294967295:
        fmt.Println(v7)
    default:
        var v10 float32 = float32(1.5)
        _ = v10
    }
    v11 := GenS2{F0: 2147483647, F1: (true && (!false))}
    v9 = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
