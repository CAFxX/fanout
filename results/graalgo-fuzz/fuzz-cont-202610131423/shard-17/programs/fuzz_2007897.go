package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "\n\t"
    var v3 uint32 = 2
    v4 := []uint32{(^(v3)), v3}
    var v5 int16 = 239
    v4 = append(v4, (^(v3)))
    v6 := GenS1{F0: 42, F1: 128}
    fmt.Println(v5)
    var v7 uint = 255
    fmt.Println(v7, v2)
    if true {
        if true {
            for fi8 := 0; fi8 < 12; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v2 = v2
                    _ = len(v4)
                    _ = fi9
                }
                _ = fi8
            }
        }
    }
    var v10 uint32 = v3
    fmt.Println(v10)
    var v11 int8 = 13
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v7, v5, v7)
        _ = len(v4)
        _ = fi12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v7 = (v7 % ((v7) | 1))
        v7 = v7
        _ = fi13
    }
    if true {
        var v14 int = 728
        var v15 bool = (255 == 0)
        var v16 bool = false
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v7)
        _ = fi17
    }
    fmt.Println(v5, v10, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
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
