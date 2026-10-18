package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    var v4 uint = 976
    v5 := []int{128, 951, -616, -1}
    var v6 int = 522
    var v7 uint = (v4 * v4)
    fmt.Println(v6, v7)
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := GenS1{F0: ("hello" + "\n\t"), F1: 0}
        v10 := GenS2{F0: (v1 ^ (^(v1))), F1: 4294967295, F2: int16(v7)}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = "a"
    v7 -= 18446744073709551615
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
