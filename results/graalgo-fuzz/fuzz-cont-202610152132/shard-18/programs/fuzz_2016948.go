package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = (!false)
    var v2 uint = 0
    var v3 float32 = (1e10 * 0.0)
    var v4 uint64 = 702
    var v5 uint8 = 255
    fmt.Println(v4, v5)
    v6 := GenS1{F0: 42}
    var v7 uint16 = 42
    fmt.Println(v4, v3, v1)
    fmt.Println(v4)
    var v8 uint64 = (v4 * uint64(v5))
    v9 := []int{78, 861, 32768}
    v1 = v1
    if (false && (!v1)) {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
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
