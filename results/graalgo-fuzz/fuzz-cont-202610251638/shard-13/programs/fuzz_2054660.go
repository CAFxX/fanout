package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    var v4 int = 127
    fmt.Println(v1, v1)
    var v5 int32 = -148
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint32 = 28
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 669, F1: 2147483648}
    _ = len(v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "":
        _ = len(v3)
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v2[42] = v4
        _ = fi9
    }
    var v10 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
