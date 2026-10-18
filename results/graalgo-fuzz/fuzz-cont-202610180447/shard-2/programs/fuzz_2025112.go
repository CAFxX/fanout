package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 157
    v3 := make(map[string]int)
    _ = len(v3)
    if false {
        v2 = (^(v2))
        for fi4 := 0; fi4 < 19; fi4++ {
            var v5 uint = 2
            _ = fi4
            _ = v5
        }
    } else {
        _ = len(v3)
    }
    v1 += v1
    v1 = (v1 * (v1 * v1))
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 uint32 = 278
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(fi6, v2)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{678, 0, 19, -672}
    var v10 int8 = -77
    var v11 int32 = 2147483647
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int8 = 57
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint32 = 1
            _ = fi14
            _ = v15
        }
        var v16 bool = (!true)
        _ = fi12
        _ = v13
        _ = v16
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS1{F0: "hello", F1: -153}
        _ = fi17
        _ = v18
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
