package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }

func main() {
    var v1 int16 = 32767
    var v2 uint64 = 972
    var v3 float64 = 1.5
    v4 := []string{(("" + "hello") + "foo bar"), "\n\t"}
    v5 := []string{("\n\t" + (("a" + ("\n\t" + "hello")) + ("foo bar" + "a"))), "", "foo bar"}
    v6 := []int{-2147483648, 9223372036854775807}
    fmt.Println(v1, v1)
    v2 = 86
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 7; fi8++ {
        v7 = v7
        v9 := GenS1{F0: fi8, F1: -1}
        _ = fi8
        _ = v9
    }
    v1 += v1
    v7 -= (v7 | (v7 * (v7 & v7)))
    var v10 int8 = 64
    v11 := GenS1{F0: 32767, F1: 45}
    fmt.Println(v3, v7, v1)
    var v12 uint64 = (^(v2))
    for fi13 := 0; fi13 < 11; fi13++ {
        v10 = int8(fi13)
        _ = fi13
    }
    v14 := GenS2{F0: 65535}
    if len(v4) > 0 {
        v4[0] = ""
    }
    v15 := GenS3{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
