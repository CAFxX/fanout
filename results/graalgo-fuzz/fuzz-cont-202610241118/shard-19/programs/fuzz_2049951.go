package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32) uint {
    fmt.Println(p0)
    p0 = p0
    var v6 float32 = 0.1
    _ = p0
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]bool)
    var v3 int64 = 9223372036854775807
    var v4 int8 = -1
    var v5 uint = 65535
    v4 = (v4 - v4)
    var v7 uint8 = 222
    v8 := GenS1{F0: -13}
    v7 -= 2
    switch 816 {
    case 9223372036854775807:
        v2[-303] = true
    case 0:
        var v9 int8 = (v4 % (((v4 | -1)) | 1))
        _ = v9
    }
    v4 -= v4
    fmt.Println(v1, v5)
    if (42 < v7) {
        var v10 float32 = 0.0
        v11 := []int{}
        var v12 int64 = (v3 ^ v3)
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{}
    var v14 float32 = ((((1e10 - -2.25) * 1.5) * 0.1) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
