package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0() int16 {
    var v4 uint32 = 2147483647
    v5 := []int{-9223372036854775808, -415}
    fmt.Println(v4)
    v6 := GenS1{F0: "\n\t", F1: "a"}
    _ = v4
    _ = v5
    _ = v6
    return -171
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 1
    var v3 int16 = -803
    v1[654] = "foo bar"
    fmt.Println(v3, v2)
    _ = len(v1)
    var v7 uint16 = 885
    var v8 uint8 = 132
    var v9 uint16 = v7
    v10 := GenS2{F0: 0, F1: false, F2: 127}
    v11 := []int{}
    var v12 int8 = -128
    v13 := GenS3{F0: 2147483647, F1: 0, F2: (v3 % (((v3 & (v3 * v3))) | 1))}
    var v14 uint16 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
