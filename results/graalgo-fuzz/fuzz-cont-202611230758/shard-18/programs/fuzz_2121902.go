package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() uint8 {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 148}
    v8 := []int{1, -2147483648, 128}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint8 = 148
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int64 = -713
    v2 := []int{int(v1), 4294967295, 42}
    v3 := make(map[int]uint32)
    var v4 int8 = -1
    v5 := make(map[int]string)
    v11 := GenS1{F0: genFn0()}
    fmt.Println(v1, v4)
    v1 += (^(v1))
    v1 += (v1 >> ((821) & 63))
    var v12 int8 = (v4 << ((0) & 63))
    fmt.Println(v4, v4, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 = (v1 & (v1 ^ v1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
