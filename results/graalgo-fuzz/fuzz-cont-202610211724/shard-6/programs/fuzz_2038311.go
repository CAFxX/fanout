package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
func genFn0(p0 uint16, p1 uint64) string {
    fmt.Println(p1)
    var v4 uint64 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = v4
    return "a"
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]string)
    var v3 int8 = 100
    var v5 int64 = 308
    fmt.Println(v5, v5)
    v5 = (^(v5))
    v6 := GenS1{F0: (!true), F1: 127, F2: 42}
    if ((1000 > 768) && true) {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int64 = (v5 % ((v5) | 1))
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 123
        _ = v9
    }
    v3 = (v3 / ((1) | 1))
    v3 += -82
    var v10 bool = (!(!false))
    fmt.Println(v5, v3, v3)
    var v11 uint = 1
    v12 := []int{-315}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
