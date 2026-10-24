package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 1
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    fmt.Println(v2)
    v5 := []int{}
    v2 -= v2
    fmt.Println(v2)
    v2 = v2
    var v6 string = "foo bar"
    delete(v1, "\n\t")
    v7 := v3[1]
    v8 := GenS1{F0: 255, F1: 539}
    v9 := GenS2{F0: -128}
    v10 := GenS2{F0: 17}
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS1{F0: 65535, F1: -32768}
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int = fi11
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
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
