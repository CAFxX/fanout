package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 uint64, p2 int16) uint8 {
    v7 := []int{32768, -1, int(p0), 2147483647}
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 204
}

func main() {
    var v1 string = ((("hello" + "foo bar") + "") + "\n\t")
    var v2 float64 = ((1.5 + float64((((0.1 - float32(1e10)) * 0.0) / 3.14159))) / (1e10 + 0.0))
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    var v5 string = "a"
    var v6 uint = 636
    fmt.Println(v5, v6)
    var v9 int = -1
    var v10 int = (v9 << ((18446744073709551615) & 63))
    var v11 float32 = 1e10
    var v12 int = 215
    v13 := []int{(v10 << (((v6 | (v6 / (((v6 % (((v6 * v6)) | 1))) | 1)))) & 63)), (v10 | 785), 42, v12}
    fmt.Println(v5, v9, v10)
    v14 := GenS1{F0: 0, F1: 46}
    v15 := GenS1{F0: 493, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
