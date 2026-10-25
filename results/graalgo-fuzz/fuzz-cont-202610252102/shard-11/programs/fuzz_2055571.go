package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 607
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 uint64 = 2
    v5 -= (v5 % (((^(v5))) | 1))
    fmt.Println(v2, v5)
    v6 := GenS1{F0: "", F1: 964}
    v3[32767] = ("" + "foo bar")
    var v7 uint16 = 58
    fmt.Println(v5, v5)
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 uint16 = 32767
        _ = fi9
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
