package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    var v6 int16 = 1
    fmt.Println(v6)
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 189
    var v2 int8 = -1
    v3 := make(map[uint32]uint32)
    v4 := []int{2147483647, 4294967295, 9223372036854775807, -984}
    v5 := []uint32{4294967295}
    v7 := GenS1{F0: 42}
    fmt.Println(v2)
    v2 = -15
    var v8 string = "a"
    v9 := GenS1{F0: 255}
    switch 4294967295 {
    case 500:
        v1 = (v1 | 255)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2, v8)
    var v10 uint = 255
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
