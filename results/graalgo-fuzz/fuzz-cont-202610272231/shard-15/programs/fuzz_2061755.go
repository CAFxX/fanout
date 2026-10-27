package main

import (
    "fmt"
)

func genFn0(p0 uint8) int8 {
    v6 := []int{2147483647, 32767, -259}
    var v7 int64 = 1
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            p0 -= 42
            _ = fi9
        }
        _ = len(v6)
        var v10 int8 = 100
        _ = fi8
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    return 106
}

func main() {
    var v1 uint16 = 431
    v2 := make(map[string]bool)
    var v3 float32 = (float32(1.5) + -2.25)
    v4 := make(map[int]int)
    var v5 uint32 = 911
    v5 = v5
    switch -9223372036854775808 {
    case -1:
        fmt.Println(v5, v5)
    default:
        if true {
            v5 -= v5
        }
    }
    v11 := []int{}
    var v12 int32 = 610
    var v13 uint = 18446744073709551615
    v1 -= (^(v1))
    if true {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = len(v11)
        v12 -= (v12 / ((v12) | 1))
    }
    fmt.Println(v1, v12, v3)
    var v14 float32 = -2.25
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
