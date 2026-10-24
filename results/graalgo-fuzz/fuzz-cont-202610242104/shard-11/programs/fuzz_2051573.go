package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 uint8 = 128
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v6 := []int{}
    v1 = (!(!(("hello" + (("a" + "a") + "a")) > "")))
    if false {
        var v7 float32 = 0.0
        delete(v2, "foo bar")
        _ = v7
    } else {
        v1 = true
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 1, F1: -1, F2: 42}
            v3 = (v3 / ((v3) | 1))
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: 65535, F1: 2147483647}
    v11 := GenS3{F0: 1, F1: 4294967295}
    fmt.Println(v1, v1)
    v12 := GenS4{F0: 42}
    v13 := GenS5{F0: -128}
    v3 = (v3 % (((v3 | (v3 / ((v3) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
