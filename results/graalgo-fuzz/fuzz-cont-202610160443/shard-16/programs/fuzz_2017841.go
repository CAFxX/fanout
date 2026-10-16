package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 70
    v3 := []int{325, 127, 0}
    v2 -= v2
    if (127 > 57) {
        fmt.Println(v2)
    } else {
        fmt.Println(v2)
        v2 = (v2 & v2)
    }
    var v4 int8 = -1
    v3 = append(v3, 9223372036854775807)
    v4 += v4
    if (1 <= 42) {
        v5 := GenS1{F0: "\n\t", F1: 65535, F2: -576}
        fmt.Println(v2)
        v4 += v4
        _ = v5
    } else {
        v6 := GenS1{F0: "foo bar", F1: 4294967295, F2: 9223372036854775807}
        _ = v6
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: "hello", F1: 2, F2: fi7}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: "hello", F1: 491, F2: -9223372036854775808}
    _ = len(v3)
    _ = len(v1)
    var v10 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
