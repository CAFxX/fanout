package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 int8 = 0
    var v5 string = "\n\t"
    var v6 float64 = 1.5
    var v7 bool = (!false)
    fmt.Println(v4, v4)
    var v8 uint32 = 224
    v9 := GenS1{F0: int(v4), F1: (v4 % ((v4) | 1))}
    _ = len(v2)
    if (!v7) {
        v10 := GenS2{F0: 4294967295}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v8 = (^(v8))
            _ = fi11
        }
    }
    v12 := GenS1{F0: 32768, F1: 19}
    v8 -= (v8 << ((4294967295) & 63))
    if v7 {
        switch 0 {
        case -390:
            v13 := []int{1}
            _ = v13
        }
    } else {
        var v14 uint32 = (v8 >> ((0) & 63))
        v5 = "a"
        _ = v14
    }
    fmt.Println(v4)
    fmt.Println(v6, v7, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
