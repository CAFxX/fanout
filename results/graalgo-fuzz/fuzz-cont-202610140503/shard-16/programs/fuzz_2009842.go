package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 19; fi4++ {
        v2 -= (v2 + v2)
        var v5 uint16 = 166
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 255
    var v7 uint8 = 172
    fmt.Println(v6, v6)
    v6 += (^(v6))
    var v8 uint = 0
    v9 := GenS1{F0: 42, F1: 200}
    var v10 int64 = -1
    v2 = v2
    var v11 int16 = -762
    var v12 int = 127
    delete(v1, 1)
    v7 = (v7 * (v7 ^ 42))
    if (v2 == v2) {
        _ = len(v1)
        v2 -= ((float32(1.5) + v2) + -2.25)
        v7 -= 0
    }
    v10 -= v10
    delete(v1, 2147483648)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
