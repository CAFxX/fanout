package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
func genFn0(p0 uint32) int16 {
    fmt.Println(p0)
    v4 := GenS1{F0: -1, F1: 974, F2: -606}
    p0 = 843
    p0 -= p0
    _ = p0
    _ = v4
    return -816
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v5 float64 = 3.14159
    v5 = 1e10
    v5 = 3.14159
    var v6 int8 = 42
    fmt.Println(v6)
    var v7 uint64 = 625
    v6 = -128
    v7 = v7
    var v8 string = "foo bar"
    fmt.Println(v8, v8, v6)
    if false {
        v6 += 42
    } else {
        v8 = "\n\t"
        switch 556 {
        case -614:
            v8 = "\n\t"
        case 32768:
            var v9 uint64 = 4294967295
            _ = v9
        case -1:
            v10 := []int{}
            _ = v10
        }
        v8 = v8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
