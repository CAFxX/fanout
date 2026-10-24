package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
func genFn0() int {
    var v4 int = 127
    v5 := []int{(-(v4)), (v4 | v4)}
    _ = len(v5)
    v6 := GenS1{F0: 666, F1: (false && (2147483648 <= 2)), F2: 4294967295}
    _ = v4
    _ = v5
    _ = v6
    return (^(v4))
}

func main() {
    var v1 float32 = ((((float32(0.0) + (0.0 * 0.1)) / (1e10 + (0.1 / ((3.14159 + ((float32(1e10) + (((0.1 / 1.0) + float32(1.5)) * 0.1)) - (0.0 / 1.5))) * ((0.1 * -2.25) * 1e10))))) * 0.0) * 3.14159)
    var v2 int32 = 103
    var v3 int = 128
    var v7 uint64 = 556
    v8 := []int{(v3 >> ((uint(v7)) & 63)), genFn0()}
    fmt.Println(v7, v1, v1)
    var v9 int8 = 1
    _ = len(v8)
    v10 := GenS1{F0: 676, F1: (!(false || (0 >= 203))), F2: uint32(v7)}
    var v11 int = (v3 - v3)
    v3 = v11
    if (!false) {
        fmt.Println(v3, v9, v7)
    }
    var v12 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
