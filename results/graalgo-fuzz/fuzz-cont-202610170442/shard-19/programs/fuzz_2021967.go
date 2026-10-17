package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 float64 = 0.0
    var v2 uint8 = 42
    var v3 string = "\n\t"
    var v4 string = v3
    v5 := make(map[uint32]int)
    if (true && (!("" != v3))) {
        v6 := []int{2147483647, 739}
        _ = v6
    }
    fmt.Println(v2, v4, v1)
    var v7 int8 = 42
    var v8 int16 = 42
    delete(v5, 1)
    v9 := v5[2147483647]
    v10 := GenS1{F0: (((v2 | (v2 >> ((65535) & 63))) > (v2 << ((4294967295) & 63))) && false)}
    var v11 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
