package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint32) uint64 {
    var v4 uint64 = 632
    v5 := []int{128, -2147483648, -9223372036854775808}
    v4 -= v4
    _ = p0
    _ = v4
    _ = v5
    return (v4 - (v4 << ((18446744073709551615) & 63)))
}
func genFn1(p0 int16, p1 int) bool {
    fmt.Println(p1, p1)
    var v6 uint8 = 1
    v6 = (v6 >> ((0) & 63))
    v6 -= 0
    p1 -= p1
    _ = p0
    _ = p1
    _ = v6
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 42
    v3 := make(map[string]bool)
    v2 += v2
    var v7 int8 = 127
    v8 := GenS1{F0: (v2 + (v2 >> ((4294967295) & 63)))}
    v7 = (v7 % (((v7 - v7)) | 1))
    v9 := []int{126, -439}
    var v10 float64 = 1e10
    var v11 float32 = 1e10
    v12 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
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
