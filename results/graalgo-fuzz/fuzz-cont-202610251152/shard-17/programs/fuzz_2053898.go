package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0() uint8 {
    v4 := []int{-9223372036854775808, -12, 128}
    var v5 uint32 = 2147483648
    var v6 uint64 = uint64(v5)
    _ = v4
    _ = v5
    _ = v6
    return 218
}

func main() {
    var v1 int = 392
    var v2 int16 = -249
    v3 := []string{"a"}
    var v7 int = (v1 >> ((0) & 63))
    switch v7 {
    case -9223372036854775808:
        v8 := GenS1{F0: "foo bar", F1: 2147483647, F2: -128}
        _ = v8
    }
    var v9 uint8 = 83
    switch "" {
    case "foo bar":
        var v10 uint8 = genFn0()
        _ = v10
    case "":
        v11 := GenS1{F0: "\n\t", F1: 1, F2: -86}
        _ = v11
    }
    v12 := GenS2{F0: -112, F1: false}
    var v13 int32 = -1
    v14 := GenS2{F0: 0, F1: (1 < 42)}
    var v15 int32 = (v13 + v13)
    var v16 float64 = ((0.1 - 1e10) / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
