package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint16) int {
    if true {
        fmt.Println(p0)
        var v7 int = 854
        var v8 int64 = 1
        _ = v7
        _ = v8
    } else {
        var v9 int64 = 9223372036854775807
        v9 = (-(v9))
        _ = v9
    }
    v10 := []int{32767, 127}
    fmt.Println(p0)
    _ = p0
    _ = v10
    return -783
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 int = 1
    var v4 int64 = -788
    v5 := make(map[int]int)
    var v6 int16 = 0
    v6 = 32767
    switch 9223372036854775807 {
    case -871:
        v4 -= 485
    case 463:
        v4 -= v4
    }
    var v11 bool = false
    var v12 uint8 = 255
    v3 += -9223372036854775808
    v13 := GenS1{F0: 100}
    var v14 int32 = -847
    v15 := []int{(-(v3)), 234, (v3 | genFn0(241))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
