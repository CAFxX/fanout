package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2, 762}
    v3 := make(map[int]bool)
    v4 := []string{""}
    switch -1 {
    case 76:
        if true {
            var v5 string = "hello"
            v6 := GenS1{F0: 0, F1: -1, F2: v5}
            _ = v5
            _ = v6
        }
    }
    var v7 float32 = ((float32((-2.25 + 0.1)) + 3.14159) * float32((-2.25 + 3.14159)))
    v8 := GenS1{F0: 379, F1: -2147483648, F2: "\n\t"}
    v9 := GenS2{F0: false}
    v10 := GenS1{F0: 97, F1: 0, F2: "\n\t"}
    v11 := GenS2{F0: false}
    v7 -= float32((0.1 + float64(0.0)))
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := v3[fi12]
        v13 = (!(1 <= 42))
        _ = fi12
        _ = v13
    }
    if false {
        var v14 string = "foo bar"
        var v15 string = v14
        _ = v14
        _ = v15
    } else {
        switch 42 {
        case 502:
            v7 -= v7
        case -9223372036854775808:
            v16 := GenS2{F0: (-1 < 127)}
            _ = v16
        case -418:
            v17 := GenS1{F0: 1, F1: -1, F2: "\n\t"}
            _ = v17
        }
        var v18 bool = (-9223372036854775808 < 213)
        v18 = v18
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
