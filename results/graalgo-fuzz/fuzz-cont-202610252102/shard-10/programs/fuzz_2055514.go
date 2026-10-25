package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 bool }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 int; F2 bool }

func main() {
    var v1 bool = (-122 == 100)
    var v2 int8 = 49
    v3 := []int{42}
    if len(v3) > 0 {
        v3[0] = -827
    }
    v4 := GenS1{F0: (2 < 602), F1: 0, F2: ((2147483648 < 0) && v1)}
    if false {
        var v5 uint32 = 399
        _ = v5
    } else {
        var v6 int8 = (v2 / (((-(v2))) | 1))
        _ = v6
    }
    v7 := GenS2{F0: 1, F1: 2}
    v2 = (v2 | (v2 | (v2 & 42)))
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: 996, F1: 65535}
        _ = fi8
        _ = v9
    }
    v10 := GenS3{F0: 701, F1: 128, F2: ("\n\t" <= "foo bar")}
    v11 := GenS3{F0: 0, F1: 127, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
