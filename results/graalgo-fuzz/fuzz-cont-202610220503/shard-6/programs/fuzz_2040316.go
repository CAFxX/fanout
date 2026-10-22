package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    var v1 int = 134
    v2 := make(map[int]int)
    var v3 string = "hello"
    var v4 string = ("a" + (v3 + v3))
    for fi5 := 0; fi5 < 5; fi5++ {
        if true {
            v1 -= v1
            var v6 int = -1
            fmt.Println(v1, v6, v1)
            _ = v6
        }
        v7 := []int{(v1 - fi5)}
        _ = fi5
        _ = v7
    }
    v4 = "hello"
    v8 := []int{(v1 + v1), v1, -1, 128}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 int64 = 487
    v1 -= 127
    delete(v2, (v1 & (-(v1))))
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: -1, F1: 1}
        v12 := GenS1{F0: -319, F1: 42}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint = 18446744073709551615
    fmt.Println(v4)
    v1 += (^(v1))
    v13 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
