package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -1
    var v3 int8 = 100
    var v4 uint16 = 0
    var v5 string = "hello"
    if false {
        var v6 int64 = -571
        if ((v3 >> ((1) & 63)) == (v3 - (v3 % (((v3 ^ -94)) | 1)))) {
            fmt.Println(v2)
            if (0 <= 4294967295) {
                var v7 uint32 = 4294967295
                v3 = v3
                _ = v7
            } else {
                fmt.Println(v2, v5, v4)
            }
        }
        v3 = v3
        _ = v6
    }
    var v8 float32 = 0.1
    v9 := GenS1{F0: 74, F1: v3}
    if (1 == 2) {
        v10 := GenS2{F0: 255}
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    v12 := []int{9223372036854775807, -295, 9223372036854775807, 4294967295}
    v5 = ("hello" + v5)
    fmt.Println(v5, v8)
    var v13 float32 = v8
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
