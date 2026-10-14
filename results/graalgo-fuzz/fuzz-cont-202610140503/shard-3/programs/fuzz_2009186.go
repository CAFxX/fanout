package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint64 }
func genFn0(p0 int64) int32 {
    v7 := GenS1{F0: 313}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{474, 1, 801, 350}
    v2 := []string{"hello", "hello", (("a" + "hello") + ""), "a"}
    var v3 uint = 2
    var v4 uint8 = 128
    var v5 bool = ((int64(v4) < 1) && false)
    v6 := []string{"hello", "foo bar", "foo bar", "\n\t"}
    if v5 {
        v5 = (!(3.14159 == 0.0))
    } else {
        fmt.Println(v4, v5)
    }
    var v8 uint64 = 576
    switch ("" + "a") {
    case "\n\t":
        var v9 bool = ((v3 + v3) <= v3)
        _ = v9
    case "a":
        var v10 uint64 = (v8 + v8)
        _ = v10
    }
    v11 := GenS2{F0: 0}
    fmt.Println(v4, v8, v3)
    v4 = uint8(v3)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v3 -= v3
    v3 -= 18446744073709551615
    v4 = v4
    v12 := GenS3{F0: 0, F1: 9223372036854775807, F2: 2}
    v13 := GenS4{F0: 32767, F1: 65535, F2: (^(v8))}
    if (-1 == genFn0(-757)) {
        var v14 int32 = genFn0(-480)
        _ = len(v2)
        _ = v14
    }
    v8 = v8
    var v15 int = -9223372036854775808
    v6 = append(v6, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
