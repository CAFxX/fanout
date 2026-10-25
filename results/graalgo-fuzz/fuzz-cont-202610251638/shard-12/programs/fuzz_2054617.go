package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 bool; F1 int8; F2 int16 }
func genFn0() int {
    if true {
        var v6 uint64 = 821
        v6 += v6
        _ = v6
    } else {
        v7 := []int{2147483647}
        _ = v7
    }
    var v8 int16 = 42
    var v9 string = "\n\t"
    v9 = "a"
    v8 = -1
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 bool, p1 int32, p2 uint8) uint32 {
    var v10 int16 = 1000
    p1 = (p1 << ((961) & 63))
    v11 := []int{-210}
    v11 = append(v11, 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 7
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 541
    var v3 bool = ("\n\t" != "\n\t")
    v4 := make(map[string]bool)
    var v5 uint = 894
    var v12 float32 = 3.14159
    var v13 uint = (^(v5))
    v14 := GenS1{F0: 127, F1: 33}
    v15 := GenS2{F0: ((3.14159 * v12) > v12), F1: -128, F2: -109}
    v16 := []int{}
    v5 = (v13 - (v13 % ((v5) | 1)))
    _ = len(v1)
    v17 := GenS1{F0: -535, F1: (^(v5))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
