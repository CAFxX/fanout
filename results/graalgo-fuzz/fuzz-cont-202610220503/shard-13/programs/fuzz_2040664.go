package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 string }

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 uint32 = 2147483647
    v1 -= v1
    switch "a" {
    case "":
        _ = len(v2)
    case "hello":
        var v5 int32 = int32(v4)
        _ = v5
    }
    v4 = v4
    var v6 uint16 = 792
    v7 := GenS1{F0: 2, F1: 100, F2: (v4 << ((255) & 63))}
    var v8 uint64 = 692
    v9 := GenS1{F0: 1, F1: 100, F2: v4}
    if false {
        v10 := GenS1{F0: 226, F1: -128, F2: v4}
        if false {
            var v11 string = "a"
            _ = v11
        } else {
            v4 += v4
            v12 := GenS2{F0: 2}
            _ = v12
        }
        v13 := GenS3{F0: -281, F1: 355, F2: "a"}
        _ = v10
        _ = v13
    } else {
        v2[2147483648] = ((("" + (((((("" + "") + "hello") + "a") + "hello") + (("hello" + "foo bar") + "\n\t")) + "a")) + "a") + "")
        switch "a" {
        case "foo bar":
            delete(v2, v4)
        }
        v1 = (v1 >> ((4294967295) & 63))
    }
    v14 := GenS1{F0: 2, F1: 1, F2: v4}
    v15 := GenS1{F0: 45, F1: 1, F2: v4}
    fmt.Println(v4)
    fmt.Println(v1)
    v16 := GenS1{F0: 42, F1: 42, F2: 2}
    var v17 string = "foo bar"
    var v18 uint8 = uint8(v6)
    var v19 float64 = ((((-2.25 / 3.14159) * ((-2.25 * (0.0 / 1.5)) * 0.0)) * ((0.0 / 1.0) + 0.0)) / 1.0)
    v20 := GenS1{F0: v18, F1: -128, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
