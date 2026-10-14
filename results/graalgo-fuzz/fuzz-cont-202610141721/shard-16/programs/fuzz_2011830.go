package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 uint16 = 2
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    v5 := []int{-1}
    v5 = append(v5, 32767)
    for fi6 := 0; fi6 < 7; fi6++ {
        v2 -= (^(v2))
        _ = fi6
    }
    v7 := GenS1{F0: ("\n\t" + "foo bar"), F1: 580}
    v8 := GenS1{F0: "hello", F1: 42}
    var v9 int32 = -1
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint64 = 2
        v4 = ("foo bar" + ((((v4 + "") + "\n\t") + v4) + v4))
        _ = fi10
        _ = v11
    }
    v2 += v2
    v12 := GenS2{F0: -1, F1: -32768, F2: 108}
    fmt.Println(v2, v1)
    _ = len(v3)
    v2 += (v2 + (v2 ^ (-(v2))))
    var v13 int8 = -13
    fmt.Println(v2, v9, v9)
    if (2147483647 <= 32767) {
        v2 += 0
        v14 := GenS3{F0: v13, F1: (!((^(v9)) < v9))}
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
