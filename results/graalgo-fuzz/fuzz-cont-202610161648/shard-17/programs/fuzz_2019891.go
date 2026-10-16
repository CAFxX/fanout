package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint16 {
    var v5 int8 = 0
    switch -253 {
    case -478:
        var v6 uint64 = 880
        _ = v6
    case 32768:
        switch 127 {
        case 32767:
            v7 := GenS1{F0: 0}
            _ = v7
        default:
            fmt.Println(v5)
        }
    }
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 0
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]uint32)
    v2 = v2
    v8 := []int{}
    var v9 string = "hello"
    delete(v4, v3)
    var v10 uint16 = 1
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if ((1 <= 871) || (!true)) {
        v11 := v1[809]
        var v12 int = v11
        _ = v11
        _ = v12
    }
    v3 = (v3 / ((v3) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
