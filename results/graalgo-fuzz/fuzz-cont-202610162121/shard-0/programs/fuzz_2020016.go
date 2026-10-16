package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    var v3 int64 = -1
    var v4 int64 = (v3 % ((v3) | 1))
    var v5 string = "a"
    var v6 uint = 179
    var v7 float32 = (0.0 - 1.5)
    fmt.Println(v4, v4)
    fmt.Println(v7)
    v8 := GenS1{F0: 255, F1: 18446744073709551615, F2: 1}
    for fi9 := 0; fi9 < 17; fi9++ {
        v4 = v4
        v6 = v6
        v10 := []int{226, (fi9 - (fi9 & (fi9 << ((v6) & 63)))), fi9}
        _ = fi9
        _ = v10
    }
    var v11 int = -546
    var v12 uint32 = 0
    var v13 uint64 = 18446744073709551615
    v6 = 42
    var v14 int16 = 1
    fmt.Println(v11, v5)
    var v15 int32 = -658
    v5 = ""
    fmt.Println(v11, v7)
    v6 -= v6
    if (true || true) {
        var v16 int32 = (v15 / ((0) | 1))
        v13 += v13
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
