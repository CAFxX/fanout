package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }
type GenS2 struct { F0 int }

func main() {
    var v1 uint = 923
    var v2 string = ("hello" + "foo bar")
    var v3 uint16 = 65535
    fmt.Println(v2)
    v4 := []int{-1, 0, 0, 2147483647}
    fmt.Println(v3, v2)
    fmt.Println(v3, v1, v3)
    var v5 uint32 = uint32(v1)
    v4 = append(v4, -2147483648)
    v6 := GenS1{F0: v3, F1: 4294967295, F2: v2}
    fmt.Println(v3, v3, v3)
    fmt.Println(v5, v1, v2)
    var v7 int8 = -128
    v7 += v7
    v8 := GenS2{F0: 0}
    var v9 uint = (v1 + v1)
    for fi10 := 0; fi10 < 19; fi10++ {
        if len(v4) > 0 {
            v4[0] = (^(fi10))
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = v2
        _ = fi10
    }
    var v11 int16 = 103
    var v12 bool = (32 >= 128)
    v13 := GenS1{F0: 42, F1: 42, F2: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
