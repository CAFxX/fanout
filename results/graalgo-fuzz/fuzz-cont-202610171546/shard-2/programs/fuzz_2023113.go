package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 122
    var v3 int = -264
    var v4 int16 = 42
    var v5 uint = 4294967295
    var v6 uint32 = 2147483648
    v7 := []int{v3, -1, -1, v3}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(v6)
    if true {
        switch (^(v3)) {
        case 9223372036854775807:
            v4 = (^(v4))
        case -1:
            v7 = append(v7, v3)
        default:
            _ = len(v7)
        }
        var v8 int8 = 0
        _ = v8
    }
    v2 = v2
    v5 -= v5
    v9 := GenS1{F0: 255, F1: 2}
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 = (-(v3))
        for fi11 := 0; fi11 < 14; fi11++ {
            v3 += 32768
            delete(v1, v3)
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
