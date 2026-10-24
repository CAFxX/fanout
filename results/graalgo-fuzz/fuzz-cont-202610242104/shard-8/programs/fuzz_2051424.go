package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int16) uint16 {
    var v7 uint8 = 95
    var v8 bool = (!true)
    fmt.Println(p1, v7, v7)
    v7 = v7
    var v9 uint64 = 817
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint8, p1 uint64) uint8 {
    var v10 float32 = (float32(0.0) + 1e10)
    p1 += 2
    _ = p0
    _ = p1
    _ = v10
    return 127
}

func main() {
    var v1 int8 = -128
    v2 := []string{("" + "foo bar"), "\n\t"}
    var v3 int = 9223372036854775807
    var v4 int64 = -1
    var v5 uint64 = 2
    var v6 float32 = ((0.0 / ((0.1 + (0.0 * 1.5)) + 1.5)) + 0.1)
    v11 := GenS1{F0: false, F1: (v4 * v4)}
    switch (v3 | v3) {
    case 378:
        fmt.Println(v5)
    }
    fmt.Println(v5, v4, v4)
    var v12 uint = 964
    v1 = v1
    var v13 int64 = (v4 - 0)
    v14 := GenS2{F0: 509}
    v6 = (v6 / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
