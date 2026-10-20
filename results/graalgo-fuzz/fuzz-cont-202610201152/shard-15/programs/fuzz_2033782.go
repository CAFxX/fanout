package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 bool }

func main() {
    var v1 string = "a"
    var v2 uint8 = 171
    v3 := []string{}
    var v4 uint8 = 128
    var v5 string = "foo bar"
    var v6 int = -2147483648
    v2 += (v4 >> ((2) & 63))
    v5 = v1
    fmt.Println(v6, v6)
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v7 uint64 = 0
    v5 = ((v5 + v1) + (v5 + "hello"))
    if ((v1 > v1) || true) {
        v8 := GenS1{F0: 42, F1: 2, F2: true}
        _ = v8
    }
    var v9 bool = (true || false)
    if (!v9) {
        var v10 int = (-(v6))
        _ = v10
    } else {
        if false {
            v5 = ""
            v6 = (v6 ^ -433)
        } else {
            v11 := GenS2{F0: 162, F1: 1}
            _ = v11
        }
        var v12 uint32 = 2147483647
        v1 = ("a" + v1)
        _ = v12
    }
    v6 = 32767
    if v9 {
        v13 := GenS3{F0: 2147483647, F1: int32(v4)}
        var v14 int = -2147483648
        _ = v13
        _ = v14
    } else {
        v15 := GenS4{F0: -1}
        _ = v15
    }
    fmt.Println(v4, v7)
    v3 = append(v3, "hello")
    v16 := GenS5{F0: 42, F1: v9}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
