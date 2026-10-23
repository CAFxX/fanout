package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint16 = 32
    v4 := make(map[string]string)
    v5 := []int{}
    v6 := []uint32{836}
    var v7 uint16 = 1
    switch ("a" + ("" + "a")) {
    case "foo bar":
        var v8 uint = 18446744073709551615
        _ = v8
    case "":
        if false {
            fmt.Println(v7)
            switch 1 {
            case 9223372036854775807:
                delete(v4, "")
            case 576:
                var v9 float32 = (1.5 + (0.1 - 3.14159))
                _ = v9
            }
        }
    }
    fmt.Println(v3, v3)
    v7 = 1
    v7 -= (v7 / ((v7) | 1))
    v10 := GenS1{F0: 255, F1: 9223372036854775807, F2: 2147483647}
    v3 = v7
    fmt.Println(v7, v7)
    if true {
        fmt.Println(v3)
        v11 := GenS1{F0: 18446744073709551615, F1: -1, F2: 32767}
        fmt.Println(v7)
        _ = v11
    } else {
        v12 := GenS1{F0: 0, F1: 2147483647, F2: -1}
        if false {
            fmt.Println(v3, v7)
            v7 = v3
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v3 = (v7 | v7)
        _ = v12
    }
    var v13 uint64 = 1
    switch ("hello" + "hello") {
    case "hello":
        if true {
            v14 := GenS1{F0: 12, F1: 127, F2: 9223372036854775807}
            _ = v14
        } else {
            var v15 int64 = -1
            v7 -= (v7 ^ v3)
            var v16 int8 = 74
            _ = v15
            _ = v16
        }
    default:
        v13 = (v13 % ((v13) | 1))
    }
    v17 := GenS2{F0: -1}
    v3 = (v3 + (v3 & 201))
    var v18 float32 = 3.14159
    v3 = (v7 << ((329) & 63))
    v19 := GenS3{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
