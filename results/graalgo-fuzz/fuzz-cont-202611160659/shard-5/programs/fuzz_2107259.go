package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 bool = true
    v2 := []int{4294967295, 974, 9223372036854775807}
    var v3 int16 = -1
    v4 := []uint32{829, 4294967295, 0}
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: 32767, F1: 42, F2: 18446744073709551615}
        _ = v5
    case "a":
        v1 = (!(!(!v1)))
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = 18446744073709551615
        _ = fi6
        _ = v7
    }
    if (435 == 0) {
        var v8 uint16 = 726
        v8 = 32767
        fmt.Println(v1, v1, v8)
        _ = v8
    }
    if (127 <= -191) {
        v1 = true
        v1 = v1
    }
    v9 := GenS2{F0: 42}
    v3 -= v3
    switch "foo bar" {
    case "":
        v1 = (-1 != 1)
    case "a":
        var v10 int8 = 100
        _ = v10
    case "hello":
        var v11 int = 4294967295
        _ = v11
    default:
        switch 128 {
        case -9223372036854775808:
            v1 = v1
        default:
            v12 := GenS3{F0: (^(v3))}
            _ = v12
        }
    }
    v1 = (!false)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
