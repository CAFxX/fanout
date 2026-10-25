package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
func genFn0(p0 int32) int16 {
    var v5 uint16 = 42
    v6 := []int{498}
    _ = p0
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 bool = (!(!(0.0 != (1.5 * (3.14159 * ((float64(float32(3.14159)) / 1.0) + 1.5))))))
    v2 := make(map[string]int)
    var v3 uint32 = 323
    v4 := make(map[int]string)
    var v7 int = 4294967295
    v8 := []int{v7, 71, v7}
    v9 := GenS1{F0: 28, F1: 2}
    v1 = true
    v8 = append(v8, (^(v7)))
    delete(v2, ("hello" + ""))
    v7 += 2147483647
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
