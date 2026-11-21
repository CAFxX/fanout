package main

import (
    "fmt"
)

func genFn0(p0 uint16, p1 uint) bool {
    var v5 int16 = -32768
    var v6 float32 = 1.5
    v7 := []int{0, 127, -2147483648}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1() uint32 {
    var v8 uint16 = 0
    var v9 uint = 451
    v10 := []int{}
    _ = len(v10)
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int64 = 0
        _ = len(v10)
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v8
    _ = v9
    return 112
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{2147483647}
    v3 := make(map[int]uint32)
    v4 := []string{}
    v1 += -1
    var v13 int8 = 0
    fmt.Println(v13)
    v1 += -1
    fmt.Println(v1, v1)
    if (!(true && false)) {
        v1 += (v1 | v1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v1)
    v13 = (v13 % ((0) | 1))
    fmt.Println(v13, v13)
    switch 9223372036854775807 {
    case 42:
        var v14 uint16 = 187
        _ = v14
    default:
        var v15 float32 = 1.5
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
