package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 int8, p1 int32, p2 bool) int {
    p2 = p2
    v6 := GenS1{F0: 65535, F1: "\n\t", F2: "a"}
    v7 := []int{}
    p0 -= p0
    var v8 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int8 = -1
    var v2 string = "a"
    v3 := []uint32{195}
    v4 := []int{-2147483648, 32767, 32767, 4294967295}
    var v5 int32 = 525
    v9 := GenS2{F0: 18446744073709551615, F1: v1}
    var v10 int8 = -128
    var v11 uint = 698
    v3 = append(v3, 331)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, 1)
    var v12 uint64 = 665
    v13 := GenS2{F0: 42, F1: (-(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
