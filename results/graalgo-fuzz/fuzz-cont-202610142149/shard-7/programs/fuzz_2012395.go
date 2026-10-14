package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 18446744073709551615
    var v3 int64 = 1
    var v4 int32 = 2147483647
    var v5 int = 127
    v6 := []int{484}
    var v7 uint32 = 1
    switch v5 {
    case -471:
        v8 := GenS1{F0: v3}
        _ = v8
    }
    if (float32(-2.25) <= 0.1) {
        v6 = append(v6, 42)
    }
    v4 = v4
    v9 := GenS2{F0: (v3 - (v3 % ((v3) | 1))), F1: 127}
    var v10 bool = (!(true && ("\n\t" >= "a")))
    var v11 float64 = (1e10 / -2.25)
    var v12 uint16 = 327
    fmt.Println(v11, v10, v11)
    v6 = append(v6, int(v4))
    var v13 uint8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
