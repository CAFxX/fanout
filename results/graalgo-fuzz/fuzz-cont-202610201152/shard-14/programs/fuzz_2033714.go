package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int16 = 1
    var v2 string = "a"
    var v3 float32 = (0.1 - float32(0.1))
    v4 := GenS1{F0: 0}
    var v5 uint8 = 255
    v6 := []int{2147483647}
    var v7 int64 = 568
    v7 -= v7
    v1 += v1
    v8 := GenS1{F0: 42}
    fmt.Println(v3)
    var v9 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
