package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }

func main() {
    var v1 uint64 = 908
    var v2 int32 = -1
    var v3 int32 = v2
    var v4 int = 970
    v5 := make(map[int]int)
    var v6 bool = true
    v7 := GenS1{F0: 65535, F1: v4, F2: 32768}
    v4 = 631
    v8 := []int{1, 128, (v4 - -1)}
    var v9 int8 = 100
    v2 = (v3 | v2)
    switch "a" {
    case "":
        v4 = int(v9)
    case "foo bar":
        v5[4294967295] = -1
    case "\n\t":
        var v10 uint8 = 255
        _ = v10
    default:
        fmt.Println(v1, v1, v4)
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 string = "hello"
        for fi13 := 0; fi13 < 14; fi13++ {
            fmt.Println(fi11, fi13)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint64 = v1
        var v16 int = (fi14 * fi14)
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = len(v5)
    for fi17 := 0; fi17 < 8; fi17++ {
        if false {
            var v18 int64 = 653
            v18 = 42
            _ = v18
        }
        _ = fi17
    }
    v1 = (v1 ^ 4294967295)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
