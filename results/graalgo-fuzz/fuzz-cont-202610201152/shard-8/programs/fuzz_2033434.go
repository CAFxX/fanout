package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int8, p1 int32, p2 int32) uint16 {
    var v7 uint32 = 0
    var v8 float32 = (0.1 - 0.0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 4294967295
    v3 := make(map[int]int)
    var v4 int = 786
    var v5 uint32 = (v2 | v2)
    v6 := make(map[string]string)
    var v9 uint16 = 21
    v10 := GenS1{F0: 919}
    v3[v4] = 128
    v11 := GenS2{F0: 0, F1: 401, F2: -486}
    v2 -= 1
    fmt.Println(v4, v2, v5)
    var v12 int16 = 436
    var v13 int = v4
    var v14 uint16 = (^(v9))
    var v15 int16 = (^(v12))
    v16 := GenS3{F0: 127}
    var v17 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
    _ = v6
    _ = v9
}
