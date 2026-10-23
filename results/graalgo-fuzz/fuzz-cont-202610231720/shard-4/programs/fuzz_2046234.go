package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 string; F1 int8; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    var v3 float32 = float32(((-2.25 + 0.0) + 0.0))
    var v4 float32 = ((v3 * v3) - (v3 + 1e10))
    var v5 uint16 = 858
    var v6 bool = true
    var v7 uint32 = 1
    switch "foo bar" {
    case "hello":
        v8 := []int{2147483647}
        _ = v8
    }
    var v9 uint64 = (v2 % (((v2 * v2)) | 1))
    v10 := []int{}
    v11 := GenS1{F0: (("foo bar" + ("" + "a")) + "hello")}
    fmt.Println(v6, v3)
    v3 = v4
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: ("a" + "hello")}
        var v14 uint8 = 42
        v15 := GenS2{F0: fi12, F1: 105, F2: (v14 * 182)}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int16 = -1
    v17 := GenS3{F0: "hello", F1: -20, F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
