package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 42
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    var v5 uint64 = uint64(v1)
    var v6 bool = false
    var v7 float32 = v3
    var v8 bool = (65535 < 218)
    var v9 int16 = 487
    var v10 uint32 = 65535
    v10 = (v10 >> (((^(v1))) & 63))
    v11 := []int{528, 127, 9223372036854775807}
    v12 := GenS1{F0: 1, F1: 1, F2: 1}
    _ = len(v11)
    if (!(127 == 1)) {
        v13 := GenS1{F0: 0, F1: 42, F2: 1}
        v14 := GenS1{F0: -9223372036854775808, F1: 42, F2: 996}
        _ = v13
        _ = v14
    }
    fmt.Println(v2, v2)
    if ((v2 / ((v2) | 1)) > (v2 - (v2 - v2))) {
        var v15 string = "a"
        var v16 string = "foo bar"
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
    _ = v7
    _ = v8
    _ = v9
}
