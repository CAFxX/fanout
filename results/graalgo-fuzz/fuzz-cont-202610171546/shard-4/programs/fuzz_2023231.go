package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int8 = -128
    var v2 int32 = 466
    v3 := make(map[int]int)
    var v4 uint16 = 0
    var v5 int64 = 1
    v6 := GenS1{F0: v2}
    v4 = (v4 * (v4 / (((v4 | v4)) | 1)))
    v2 += v2
    v7 := []int{793}
    switch (("a" + "") + ("" + "hello")) {
    case "hello":
        if true {
            var v8 int64 = -165
            for fi9 := 0; fi9 < 15; fi9++ {
                switch "\n\t" {
                case "a":
                    switch "foo bar" {
                    case "":
                        fmt.Println(v2, fi9, v4)
                    default:
                        fmt.Println(v1, v4)
                    }
                case "\n\t":
                    v4 -= v4
                }
                _ = fi9
            }
            v4 = v4
            _ = v8
        }
    case "a":
        fmt.Println(v5, v1, v4)
    }
    v10 := GenS2{F0: 702, F1: 2147483648}
    v1 += (v1 / (((v1 % (((v1 << ((158) & 63))) | 1))) | 1))
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
