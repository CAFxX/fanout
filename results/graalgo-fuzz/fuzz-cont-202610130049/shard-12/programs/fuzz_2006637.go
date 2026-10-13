package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 uint8 = 61
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 bool = (((42 >= 4294967295) && false) && (!(!true)))
    var v6 float32 = v2
    v7 := []int{99}
    v4[("a" + "hello")] = "hello"
    v2 = -2.25
    fmt.Println(v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := v4[""]
        fmt.Println(v6, v5)
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (42 == -356)}
    fmt.Println(v2, v6, v2)
    var v11 uint = 767
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS2{F0: 379, F1: 0, F2: 2}
        v6 = v6
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
