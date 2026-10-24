package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{815, -665, 2147483647}
    var v2 string = ((("a" + "foo bar") + "") + "\n\t")
    v3 := []uint32{4294967295}
    var v4 int64 = 609
    var v5 int16 = 0
    v2 = v2
    v5 -= 597
    v5 += v5
    var v6 uint32 = 65535
    var v7 int32 = -2147483648
    v8 := GenS1{F0: -219, F1: 42}
    var v9 float32 = 1.5
    var v10 uint64 = 831
    switch 127 {
    case 1:
        var v11 float32 = (v9 * 0.1)
        _ = v11
    default:
        v12 := GenS2{F0: 38}
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
