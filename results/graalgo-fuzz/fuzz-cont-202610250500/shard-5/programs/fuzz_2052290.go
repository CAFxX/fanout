package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 0
    var v2 uint64 = 546
    var v3 string = "\n\t"
    v2 += v2
    var v4 float32 = (1e10 / 1e10)
    v5 := []int{int(v2), 4294967295}
    fmt.Println(v2, v2)
    v1 = v1
    var v6 int = -1
    if false {
        v1 = (v1 * v1)
        var v7 int16 = 1
        _ = v7
    }
    v8 := GenS1{F0: 795}
    v5 = append(v5, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
