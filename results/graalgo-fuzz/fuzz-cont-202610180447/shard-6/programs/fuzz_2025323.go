package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 int16, p1 int16) string {
    fmt.Println(p1, p0)
    var v5 int64 = 42
    var v6 bool = true
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (("" + ("hello" + (("a" + "hello") + (("a" + "") + "hello")))) + "\n\t")
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    var v4 float64 = 0.0
    var v7 uint8 = 133
    for fi8 := 0; fi8 < 4; fi8++ {
        v2 = (v2 / ((v2) | 1))
        _ = fi8
    }
    fmt.Println(v4, v2)
    var v9 string = ""
    var v10 int8 = (v2 ^ (^(v2)))
    v11 := GenS1{F0: 0, F1: false}
    delete(v1, 2)
    v12 := GenS2{F0: -96, F1: (!false)}
    var v13 uint = 18446744073709551615
    v13 = v13
    var v14 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
