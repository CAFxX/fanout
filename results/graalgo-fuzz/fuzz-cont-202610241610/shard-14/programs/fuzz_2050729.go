package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[uint32]string)
    v3 := []int{9223372036854775807, 42, 413}
    var v4 uint16 = 464
    v5 := GenS1{F0: 42, F1: -200}
    switch "foo bar" {
    case "\n\t":
        switch "foo bar" {
        case "":
            v4 += v4
        case "hello":
            v6 := GenS2{F0: "a"}
            _ = v6
        case "foo bar":
            fmt.Println(v4)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "foo bar":
        var v7 float32 = v1
        _ = v7
    }
    if false {
        v4 = 42
        if false {
            v1 += ((v1 * v1) * v1)
        }
    } else {
        var v8 int64 = 474
        _ = v8
    }
    var v9 uint64 = 0
    switch "foo bar" {
    case "hello":
        var v10 int32 = 1
        _ = v10
    }
    v9 += 151
    v4 = (v4 >> ((1) & 63))
    v11 := GenS3{F0: uint8(v9), F1: 386, F2: 1000}
    v1 = v1
    v12 := GenS4{F0: 522, F1: 127, F2: 0}
    v13 := GenS5{F0: 42}
    var v14 uint8 = 42
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
