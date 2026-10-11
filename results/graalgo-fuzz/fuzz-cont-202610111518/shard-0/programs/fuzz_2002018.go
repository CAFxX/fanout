package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }

func main() {
    var v1 uint8 = 133
    var v2 bool = false
    v3 := []string{}
    var v4 float32 = (-2.25 / 1.0)
    v5 := make(map[uint32]string)
    v4 = 0.1
    fmt.Println(v2)
    var v6 float32 = v4
    _ = len(v3)
    v5[45] = "a"
    v7 := GenS1{F0: 1, F1: 2, F2: 32767}
    var v8 string = "hello"
    v9 := GenS2{F0: 42}
    if (!v2) {
        if (4294967295 > 42) {
            v10 := GenS3{F0: -376}
            _ = v10
        }
        var v11 float32 = -2.25
        delete(v5, 42)
        _ = v11
    }
    v4 -= 0.0
    var v12 uint32 = 0
    v6 -= ((v4 / 3.14159) + v4)
    var v13 int = 127
    v14 := GenS1{F0: 127, F1: 491, F2: 96}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
