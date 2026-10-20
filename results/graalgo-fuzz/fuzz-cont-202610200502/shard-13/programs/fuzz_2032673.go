package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }

func main() {
    var v1 int32 = 0
    var v2 int16 = 1
    var v3 uint = 1
    fmt.Println(v2)
    v4 := []int{-1}
    if (!true) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v2)
    } else {
        v4 = append(v4, int(v3))
    }
    v2 = (-(v2))
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1, F2: 148}
        var v7 uint16 = 255
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 += (v3 % ((v3) | 1))
    fmt.Println(v3, v3)
    v8 := GenS2{F0: "hello", F1: 200, F2: 42}
    v4 = append(v4, 32768)
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint = 18446744073709551615
        _ = fi9
        _ = v10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
