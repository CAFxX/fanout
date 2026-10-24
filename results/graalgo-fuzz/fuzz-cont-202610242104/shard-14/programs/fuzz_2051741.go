package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
func genFn0(p0 uint, p1 int64) int {
    v6 := GenS1{F0: (p1 % (((p1 / ((p1) | 1))) | 1))}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v6
    return -511
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 381
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := []string{("hello" + ("foo bar" + "\n\t"))}
    fmt.Println(v2)
    v2 = v2
    v7 := GenS1{F0: 1}
    v8 := GenS2{F0: 221}
    if (-502 > 127) {
        var v9 int64 = v1
        _ = v9
    }
    var v10 int8 = 0
    var v11 uint64 = 0
    v12 := GenS2{F0: 255}
    delete(v3, 4294967295)
    var v13 uint32 = v2
    v14 := v3[127]
    v15 := GenS3{F0: 2147483647, F1: (^(v11)), F2: 42}
    v16 := GenS1{F0: (^(v1))}
    v11 -= v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
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
    _ = v7
    _ = v8
}
