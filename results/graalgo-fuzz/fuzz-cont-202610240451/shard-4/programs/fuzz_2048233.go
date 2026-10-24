package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16 }
func genFn0() uint32 {
    v6 := []int{32767}
    switch "\n\t" {
    case "hello":
        v7 := GenS1{F0: 1, F1: 2}
        _ = v7
    case "\n\t":
        var v8 string = ""
        _ = v8
    case "foo bar":
        v6 = append(v6, 127)
    }
    _ = v6
    return 42
}
func genFn1(p0 int32) int32 {
    v9 := GenS2{F0: 4294967295}
    var v10 int64 = -9223372036854775808
    _ = p0
    _ = v10
    _ = v9
    return (p0 >> ((0) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 891
    var v3 float32 = (-2.25 * (3.14159 * float32(-2.25)))
    var v4 float64 = 0.1
    var v5 int64 = -216
    v11 := GenS3{F0: false, F1: 42}
    var v12 float32 = v3
    fmt.Println(v5, v5, v2)
    v13 := GenS4{F0: 32767}
    v14 := []int{-36}
    var v15 int8 = 127
    v2 -= (v2 << ((1) & 63))
    fmt.Println(v4, v15, v15)
    var v16 uint8 = uint8(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
