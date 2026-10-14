package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int64) uint64 {
    var v4 bool = true
    switch ("hello" + "hello") {
    case "hello":
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint16 = 451
            fmt.Println(v4, v4, p0)
            _ = fi5
            _ = v6
        }
    case "foo bar":
        v4 = true
    default:
        if true {
            var v7 int8 = -128
            _ = v7
        }
    }
    _ = p0
    _ = v4
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := []uint32{2147483648, 2147483648}
    if false {
        var v8 uint8 = 189
        v8 += 255
        _ = v8
    } else {
        _ = len(v3)
        var v9 uint16 = 0
        var v10 int64 = 796
        _ = v10
        _ = v9
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int8 = -1
        var v13 uint64 = 42
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = -2.25
    var v15 int8 = -128
    var v16 uint8 = 118
    var v17 int8 = v15
    var v18 float32 = v14
    v19 := GenS1{F0: "foo bar", F1: (-1 == 128)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
