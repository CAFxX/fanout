package main

import (
    "fmt"
)

func genFn0() int32 {
    var v7 int8 = -1
    fmt.Println(v7)
    v7 -= v7
    v8 := []int{}
    _ = v7
    _ = v8
    return -51
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = 42
    v3 := make(map[int]uint32)
    var v4 uint = 4294967295
    v5 := []int{127, int(v2), 127, 128}
    v6 := make(map[int]int)
    var v9 int64 = v2
    v2 += v9
    var v10 int8 = -128
    var v11 float32 = (0.1 + -2.25)
    fmt.Println(v11, v2, v2)
    var v12 string = ""
    for fi13 := 0; fi13 < 12; fi13++ {
        v10 += (v10 + (v10 << ((v4) & 63)))
        _ = fi13
    }
    _ = len(v6)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
