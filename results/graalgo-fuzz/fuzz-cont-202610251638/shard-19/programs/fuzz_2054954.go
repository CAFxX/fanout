package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = 0
    var v2 int64 = v1
    v3 := make(map[uint32]int)
    var v4 uint16 = 612
    var v5 int = 32768
    v6 := []int{}
    if true {
        var v7 float32 = 1e10
        _ = v7
    }
    if ("" < ("\n\t" + "a")) {
        var v8 uint8 = 1
        switch 825 {
        case 0:
            v9 := GenS1{F0: v5, F1: true}
            _ = v9
        case -2147483648:
            var v10 int64 = v2
            _ = v10
        case 2147483647:
            var v11 int8 = -128
            _ = v11
        }
        _ = v8
    }
    var v12 float32 = 0.0
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 float32 = 1.5
        v15 := GenS2{F0: 2}
        var v16 int64 = v1
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint8 = 2
    v5 = v5
    switch "" {
    case "foo bar":
        v12 = ((v12 - 0.1) + ((3.14159 / (0.0 * (float32(1.5) - (v12 - 0.1)))) * v12))
    }
    v5 += v5
    v6 = append(v6, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
