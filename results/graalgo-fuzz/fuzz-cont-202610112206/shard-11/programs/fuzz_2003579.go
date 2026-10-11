package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    var v3 string = ((v2 + v2) + v2)
    v4 := make(map[string]bool)
    var v5 uint8 = 128
    v2 = v2
    _ = len(v4)
    v6 := GenS1{F0: (^(v5))}
    if (3.14159 <= float32(1.5)) {
        fmt.Println(v2)
    }
    var v7 uint8 = 128
    v8 := []int{32767, 66}
    fmt.Println(v5, v3, v3)
    if (true || (2147483647 > -2147483648)) {
        _ = len(v8)
    }
    v3 = v2
    if len(v8) > 0 {
        v8[0] = int(v5)
    }
    if (18446744073709551615 > 2) {
        v9 := GenS2{F0: 4294967295, F1: -173, F2: 42}
        _ = v9
    }
    v3 = ""
    v10 := v4["foo bar"]
    fmt.Println(v3, v10, v5)
    v11 := GenS1{F0: v7}
    fmt.Println(v3, v7, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
