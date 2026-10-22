package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 312
    v3 := []uint32{}
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := GenS1{F0: fi5, F1: false}
        var v7 uint8 = 2
        v8 := GenS1{F0: (fi5 << ((273) & 63)), F1: true}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(v2, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    v2 -= 0
    v2 += (v2 >> ((42) & 63))
    v9 := GenS2{F0: true, F1: -55}
    if len(v3) > 0 {
        v3[0] = 2
    }
    v10 := GenS1{F0: 101, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
