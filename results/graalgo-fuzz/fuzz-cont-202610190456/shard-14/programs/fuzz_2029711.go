package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 float32 = 1e10
    v2 := []int{-405, 9223372036854775807, -628, 32768}
    var v3 string = ""
    v4 := []uint32{}
    var v5 uint64 = 0
    var v6 string = v3
    v6 = (("a" + v3) + ("\n\t" + "\n\t"))
    v7 := GenS1{F0: 0, F1: 751}
    v5 = (v5 * 675)
    var v8 int32 = -2147483648
    var v9 uint32 = 911
    v2 = append(v2, -943)
    v1 -= v1
    var v10 int64 = -4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
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
