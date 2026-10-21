package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
func genFn0(p0 int8) uint {
    v4 := []int{}
    p0 = p0
    v5 := GenS1{F0: 128, F1: 1}
    p0 -= 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = v4
    _ = v5
    return 65535
}

func main() {
    var v1 uint16 = 42
    v2 := []string{"\n\t", ""}
    v3 := []uint32{}
    var v6 string = ""
    var v7 bool = false
    v8 := GenS1{F0: -743, F1: 4294967295}
    v7 = v7
    var v9 uint = genFn0(-128)
    var v10 float64 = (1e10 / -2.25)
    v11 := GenS2{F0: 18446744073709551615, F1: 716, F2: v7}
    fmt.Println(v1, v10, v6)
    var v12 string = "\n\t"
    _ = len(v2)
    v12 = (v6 + (v6 + ("a" + v12)))
    v13 := GenS2{F0: 4294967295, F1: 2147483647, F2: (!(!(42 > 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
