package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 uint32; F1 string; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    v5 := make(map[string]string)
    var v6 string = ""
    v7 := GenS1{F0: -859}
    v6 = v6
    var v8 float32 = 0.0
    var v9 float64 = 1.5
    if false {
        v10 := v5[((v6 + ("foo bar" + ("\n\t" + v6))) + v6)]
        v11 := GenS2{F0: v4, F1: v6, F2: 65535}
        _ = v10
        _ = v11
    }
    var v12 string = v6
    if (!(!(!v4))) {
        delete(v5, "foo bar")
        var v13 bool = true
        _ = v13
    }
    fmt.Println(v6)
    var v14 int16 = 42
    fmt.Println(v8)
    v15 := []int{-9223372036854775808, 235}
    v16 := GenS3{F0: 2147483647, F1: "foo bar", F2: 1}
    v8 = float32(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
