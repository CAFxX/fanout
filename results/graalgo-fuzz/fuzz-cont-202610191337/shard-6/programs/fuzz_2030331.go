package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint }

func main() {
    v1 := []int{-702, -9223372036854775808, 936}
    var v2 int32 = -1
    v3 := []string{("foo bar" + "hello"), "hello", "a"}
    v4 := []int{713, 394, -706, -9223372036854775808}
    var v5 int32 = (v2 / (((v2 / ((v2) | 1))) | 1))
    v2 -= (-(v5))
    v3 = append(v3, "\n\t")
    v5 = (v2 - v2)
    v4 = append(v4, 32768)
    v2 -= v5
    v5 += (v2 + -897)
    v6 := GenS1{F0: 42}
    v7 := GenS2{F0: (v5 * (v2 << ((105) & 63))), F1: 1}
    switch 32768 {
    case -2147483648:
        if len(v1) > 0 {
            v1[0] = 127
        }
    case 0:
        v8 := GenS3{F0: 4294967295}
        _ = v8
    }
    var v9 uint8 = 210
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
