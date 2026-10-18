package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() uint64 {
    v7 := []int{}
    var v8 uint = 42
    v8 += v8
    var v9 int16 = 32767
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{}
    v3 := []uint32{1, 4294967295}
    var v4 uint32 = 387
    var v5 bool = true
    v6 := []uint32{v4, v4}
    v3 = append(v3, v4)
    var v10 float64 = ((-2.25 + 1.5) + 1e10)
    var v11 uint16 = 32767
    v12 := GenS1{F0: 18446744073709551615}
    v4 += (v4 ^ v4)
    var v13 uint = uint(v1)
    v14 := GenS1{F0: (^(v13))}
    fmt.Println(v4)
    v10 = v10
    v1 += (v1 + (^(v1)))
    var v15 string = (("hello" + "hello") + "")
    var v16 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
