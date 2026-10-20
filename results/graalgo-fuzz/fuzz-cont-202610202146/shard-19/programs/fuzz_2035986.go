package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 659
    var v3 uint64 = (v2 / ((v2) | 1))
    var v4 string = ""
    var v5 uint16 = v1
    v6 := []int{1}
    var v7 uint8 = 42
    var v8 int64 = -992
    if true {
        fmt.Println(v3, v1, v5)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v8 -= (-(v8))
    }
    v9 := GenS1{F0: 1}
    if ((!false) || false) {
        v10 := GenS1{F0: -32768}
        _ = v10
    } else {
        var v11 uint64 = v3
        v12 := GenS1{F0: -1}
        v13 := GenS2{F0: 2147483648}
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(v8, v2, v5)
    var v14 uint = 4294967295
    var v15 float64 = 0.0
    v3 = uint64(v1)
    v16 := GenS2{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
