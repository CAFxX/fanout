package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int, p1 int8, p2 uint16) uint {
    if (42 >= 18446744073709551615) {
        fmt.Println(p0, p2)
    }
    var v6 int64 = 0
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint8 = 0
    var v2 bool = true
    v3 := []int{4294967295, 128, int(v1), 32768}
    var v4 int64 = -1
    v5 := make(map[string]int)
    for fi7 := 0; fi7 < 10; fi7++ {
        v1 = (v1 - v1)
        var v8 int32 = 959
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 2147483647
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v2, v4)
        _ = fi10
    }
    if true {
        fmt.Println(v1, v2, v1)
    } else {
        fmt.Println(v2, v2)
        v9 -= v9
    }
    v11 := GenS1{F0: 369, F1: 1}
    var v12 int32 = 2147483647
    v13 := GenS1{F0: 4294967295, F1: 42}
    v14 := GenS1{F0: 42, F1: 18446744073709551615}
    var v15 int8 = -128
    v16 := GenS1{F0: uint64(v9), F1: 4294967295}
    fmt.Println(v12, v1, v9)
    var v17 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
