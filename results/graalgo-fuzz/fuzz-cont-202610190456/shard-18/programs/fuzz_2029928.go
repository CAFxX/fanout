package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint8, p1 int32, p2 int32) uint32 {
    if true {
        var v4 int8 = 42
        _ = v4
    } else {
        if (!(!(976 > 4294967295))) {
            p1 = p1
            v5 := []int{}
            _ = v5
        }
    }
    v6 := GenS1{F0: 1, F1: 42}
    var v7 uint64 = 816
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 int8 = 72
    var v2 uint = 4294967295
    v3 := make(map[int]bool)
    var v8 int32 = 1
    v1 += (v1 >> ((v2) & 63))
    fmt.Println(v1)
    v9 := []int{}
    var v10 int = -181
    var v11 float64 = 0.1
    switch ("\n\t" + "hello") {
    case "":
        _ = len(v9)
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 float64 = (0.1 / (v11 / 3.14159))
            _ = fi13
            _ = v14
        }
        v8 += v8
        _ = fi12
    }
    v10 -= (v10 | (v10 >> ((65535) & 63)))
    for fi15 := 0; fi15 < 8; fi15++ {
        v8 += (v8 >> ((4294967295) & 63))
        fmt.Println(v2)
        _ = len(v3)
        _ = fi15
    }
    _ = len(v9)
    var v16 int32 = (v8 * 597)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
