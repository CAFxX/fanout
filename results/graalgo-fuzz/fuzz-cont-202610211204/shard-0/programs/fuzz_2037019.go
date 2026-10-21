package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0() uint32 {
    var v6 uint8 = 1
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 = (v6 & (^(v6)))
        var v8 uint32 = 1
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 865
    _ = v6
    _ = v9
    return 428
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]uint32)
    var v3 string = "foo bar"
    var v4 int32 = -1
    var v5 uint8 = 213
    v10 := []int{-1, 127, 4294967295, 127}
    var v11 int32 = -2147483648
    v12 := GenS1{F0: 42, F1: 4294967295}
    fmt.Println(v11, v5, v4)
    v13 := GenS1{F0: -168, F1: 1}
    var v14 int64 = 1
    var v15 uint64 = 863
    v15 = v15
    v15 = 42
    v16 := GenS1{F0: 1, F1: (v15 >> ((65535) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
