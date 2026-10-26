package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 string }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 int; F2 uint16 }

func main() {
    v1 := []int{}
    var v2 uint8 = 255
    v3 := []string{"hello", ("a" + "foo bar")}
    var v4 float32 = -2.25
    fmt.Println(v4, v4)
    v5 := GenS1{F0: 5}
    v6 := GenS2{F0: 941}
    v4 = (0.1 / float32(1e10))
    v7 := GenS3{F0: 9223372036854775807, F1: 1}
    switch "hello" {
    case "a":
        v8 := GenS4{F0: 42, F1: 65535, F2: ""}
        _ = v8
    case "\n\t":
        _ = len(v1)
    case "hello":
        v9 := GenS5{F0: 2, F1: 4294967295}
        _ = v9
    default:
        var v10 int = -1
        _ = v10
    }
    var v11 uint16 = 1
    v2 += (v2 >> ((12) & 63))
    fmt.Println(v11, v11)
    _ = len(v1)
    v12 := GenS6{F0: -447, F1: -69, F2: 917}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
