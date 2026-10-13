package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32) uint32 {
    var v4 int16 = -937
    var v5 string = "a"
    var v6 bool = (!true)
    p0 += p0
    fmt.Println(v5, p0)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint32) int64 {
    v7 := []int{128, 4294967295, -2147483648}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint = 0
    var v2 int32 = int32(v1)
    var v3 bool = false
    v8 := []int{798, 4294967295, 127}
    if len(v8) > 0 {
        v8[0] = 42
    }
    var v9 bool = (v3 || v3)
    v10 := GenS1{F0: v1}
    if (0.0 == (0.0 + 3.14159)) {
        v11 := GenS1{F0: (v1 | (v1 - v1))}
        fmt.Println(v3, v1, v9)
        _ = v11
    }
    v8 = append(v8, -1)
    v1 = (v1 >> (((v1 >> ((v1) & 63))) & 63))
    v12 := GenS2{F0: 0}
    var v13 bool = true
    v14 := GenS1{F0: (v1 / ((v1) | 1))}
    v13 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
