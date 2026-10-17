package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 int16; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 string = (("\n\t" + ((("foo bar" + "foo bar") + "hello") + "")) + "hello")
    v4 := []uint32{2}
    v5 := make(map[uint32]int)
    v6 := []uint32{2147483648}
    _ = len(v5)
    var v7 uint64 = 1
    v8 := GenS1{F0: "", F1: 69, F2: 42}
    var v9 int8 = -1
    v10 := GenS2{F0: -32768, F1: -26, F2: 42}
    v7 += v7
    switch 821 {
    case -1:
        v11 := GenS2{F0: 1000, F1: (-(v9)), F2: 0}
        _ = v11
    case 0:
        var v12 int = -424
        _ = v12
    case -295:
        switch 4294967295 {
        case 128:
            v13 := GenS3{F0: 42, F1: v7}
            _ = v13
        }
    }
    v9 = (v9 ^ v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
