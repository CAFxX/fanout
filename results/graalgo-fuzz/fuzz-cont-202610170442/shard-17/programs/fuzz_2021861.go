package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 uint }
func genFn0(p0 int8, p1 uint64) int32 {
    p1 += p1
    fmt.Println(p1, p1)
    v5 := GenS1{F0: (p1 * (p1 + (p1 ^ 167)))}
    v6 := GenS1{F0: (p1 >> ((1) & 63))}
    p1 -= 18446744073709551615
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int8, p1 uint8) string {
    p0 = (p0 % ((p0) | 1))
    fmt.Println(p0, p0)
    p0 = p0
    _ = p0
    _ = p1
    return "\n\t"
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"a", (("a" + ("\n\t" + "hello")) + ""), ("hello" + ("" + "foo bar"))}
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    v1 -= v1
    var v7 uint = 4294967295
    v8 := GenS1{F0: (v1 ^ (v1 / (((v1 | 4294967295)) | 1)))}
    var v9 uint16 = 66
    v7 -= (v7 | uint(v1))
    v4 -= v4
    v1 = (v1 - v1)
    v10 := GenS2{F0: -24, F1: -9223372036854775808, F2: (v7 >> (((v7 ^ (^(v7)))) & 63))}
    switch -72 {
    case 4294967295:
        if ((("hello" + "") <= ("" + "")) && (932 != 2147483647)) {
            var v11 float32 = (3.14159 / 1e10)
            _ = v11
        } else {
            v1 += 33
            v4 += -1
            var v12 uint64 = 2
            _ = v12
        }
    case -1:
        var v13 uint = v7
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
