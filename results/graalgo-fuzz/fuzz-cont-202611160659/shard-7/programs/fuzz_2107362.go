package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }

func main() {
    var v1 uint32 = 65535
    var v2 int64 = 1
    var v3 uint16 = 615
    v4 := make(map[int]bool)
    var v5 bool = false
    v6 := GenS1{F0: ("" + ("foo bar" + (("" + "hello") + "hello"))), F1: -534, F2: -1}
    var v7 int64 = v2
    if (v7 < v2) {
        v8 := GenS1{F0: ("a" + "\n\t"), F1: 1000, F2: 115}
        var v9 int32 = -501
        _ = v8
        _ = v9
    }
    v1 = v1
    v2 -= 1
    v3 -= 31
    fmt.Println(v7, v1)
    var v10 uint64 = 4294967295
    var v11 float64 = 1.5
    var v12 float32 = 1e10
    for fi13 := 0; fi13 < 15; fi13++ {
        if (!true) {
            v14 := GenS1{F0: "\n\t", F1: 0, F2: 42}
            _ = v14
        } else {
            v1 = (v1 >> ((uint(v10)) & 63))
        }
        _ = fi13
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
