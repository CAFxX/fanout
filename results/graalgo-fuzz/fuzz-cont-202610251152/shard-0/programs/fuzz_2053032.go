package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
func genFn0(p0 uint32) string {
    p0 += p0
    switch ((("foo bar" + (("hello" + "a") + ("\n\t" + "\n\t"))) + "foo bar") + "\n\t") {
    case "foo bar":
        var v5 uint16 = 618
        _ = v5
    case "":
        var v6 int64 = -173
        _ = v6
    case "a":
        var v7 uint16 = 357
        _ = v7
    }
    var v8 uint64 = 4294967295
    switch ("" + ("a" + "\n\t")) {
    case "\n\t":
        v8 -= v8
    case "foo bar":
        fmt.Println(v8)
    default:
        v9 := []int{-885, 32767}
        _ = v9
    }
    v10 := []int{-2147483648, 32767, 0}
    _ = p0
    _ = v10
    _ = v8
    return ((("hello" + "a") + ("\n\t" + ("" + "\n\t"))) + "foo bar")
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    var v3 string = "\n\t"
    v4 := make(map[uint32]string)
    _ = len(v2)
    v11 := []int{-1, -641, 32767}
    if (-2147483648 > 9223372036854775807) {
        v12 := GenS1{F0: "hello", F1: 69}
        var v13 bool = true
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 834, F1: -1, F2: 32767}
    var v15 uint64 = 2
    v15 += (^(v15))
    fmt.Println(v15)
    v3 = ""
    if (2 != 255) {
        var v16 int8 = -128
        if (!((v1 + v1) < v1)) {
            if (!(true || true)) {
                v16 = (v16 + v16)
            } else {
                var v17 int64 = 1
                _ = v17
            }
            v15 = (v15 & (v15 % ((202) | 1)))
        }
        _ = len(v4)
        _ = v16
    }
    v18 := v2[0]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
