package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint8 {
    v4 := []int{1, -1, 9223372036854775807, 2147483647}
    switch 9223372036854775807 {
    case -1:
        var v5 uint64 = 810
        _ = v5
    case 27:
        _ = len(v4)
    }
    if len(v4) > 0 {
        v4[0] = 736
    }
    switch 990 {
    case 4294967295:
        var v6 uint64 = 4294967295
        _ = v6
    case 9223372036854775807:
        var v7 uint16 = 865
        _ = v7
    }
    _ = v4
    return 1
}

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    fmt.Println(v3, v3)
    v8 := GenS1{F0: -1}
    v9 := []int{}
    for fi10 := 0; fi10 < 8; fi10++ {
        v3 = v3
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v1)
        _ = fi11
    }
    v12 := GenS1{F0: 1000}
    v13 := GenS1{F0: 0}
    fmt.Println(v3)
    v3 = -164
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
