package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint }
type GenS4 struct { F0 uint; F1 int8; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int; F2 uint32 }
func genFn0() int32 {
    var v7 int = -2147483648
    var v8 uint16 = 1
    var v9 float32 = -2.25
    if true {
        if (128 == 0) {
            v9 = (v9 - (0.1 * v9))
            v9 = 0.0
            v10 := GenS1{F0: (-(v7))}
            _ = v10
        }
        var v11 string = ("\n\t" + "foo bar")
        switch v11 {
        case "hello":
            var v12 string = "\n\t"
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := []int{fi13}
                v7 += (v7 / ((v7) | 1))
                _ = fi13
                _ = v14
            }
        }
        _ = v11
    }
    if (255 != 2) {
        var v15 uint64 = 1
        v16 := GenS2{F0: uint8(v8), F1: -128}
        _ = v15
        _ = v16
    }
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    var v2 uint64 = 1
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    v17 := GenS2{F0: 128, F1: 1}
    v18 := []int{729, 128, -705}
    v19 := GenS3{F0: (true && true), F1: 65535, F2: 2}
    fmt.Println(v6, v4)
    v1 = "foo bar"
    switch (v1 + (v1 + v1)) {
    case "foo bar":
        var v20 int64 = 42
        _ = v20
    case "":
        if (!false) {
            fmt.Println(v1, v6, v4)
            v1 = ("\n\t" + ("\n\t" + ((("" + v1) + v1) + v1)))
            fmt.Println(v1, v1, v1)
        }
    case "a":
        v21 := GenS4{F0: uint(v2), F1: -128, F2: 874}
        _ = v21
    default:
        var v22 float32 = 0.1
        _ = v22
    }
    v1 = "hello"
    v23 := GenS5{F0: true}
    v24 := GenS6{F0: v1, F1: 4294967295, F2: 261}
    fmt.Println(v2, v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
