package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 uint16) int {
    var v6 uint = 255
    var v7 int64 = 42
    var v8 int = 4294967295
    v7 += v7
    var v9 uint = v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return int(p1)
}
func genFn1() int8 {
    var v10 uint64 = 0
    fmt.Println(v10)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint8 = 127
        fmt.Println(v10, v10)
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: 255}
    _ = v10
    _ = v13
    return 42
}

func main() {
    v1 := []int{-2147483648}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    v5 := make(map[uint32]string)
    var v14 bool = true
    v14 = (!(2 == 807))
    var v15 int8 = genFn1()
    var v16 int16 = -381
    v16 = (^(v16))
    v16 = v16
    v14 = v14
    var v17 float32 = 3.14159
    v18 := GenS1{F0: 42, F1: 1000}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
