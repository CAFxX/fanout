package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := []int{-999}
    var v2 uint8 = 123
    var v3 float32 = 1e10
    var v4 uint = 18446744073709551615
    v5 := make(map[int]string)
    v6 := make(map[int]uint32)
    fmt.Println(v4)
    v2 = (v2 + v2)
    v2 -= (v2 >> ((v4) & 63))
    fmt.Println(v4, v3, v3)
    v2 += 36
    _ = len(v1)
    if true {
        if (42 <= 0) {
            fmt.Println(v4)
            switch "\n\t" {
            case "foo bar":
                var v7 uint64 = 18446744073709551615
                _ = v7
            case "\n\t":
                var v8 uint32 = 966
                _ = v8
            case "":
                fmt.Println(v4, v2)
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 string = "foo bar"
                v4 -= (^(v4))
                _ = fi9
                _ = v10
            }
        }
        if true {
            if (2147483648 >= 2) {
                var v11 uint = 42
                v12 := GenS1{F0: v2, F1: 127, F2: 42}
                var v13 int32 = 0
                _ = v11
                _ = v12
                _ = v13
            }
            fmt.Println(v4, v2)
        }
        switch "hello" {
        case "foo bar":
            _ = len(v6)
        }
    } else {
        v1 = append(v1, 127)
        v3 -= (float32(-2.25) - v3)
    }
    if (!false) {
        if true {
            switch "hello" {
            case "foo bar":
                var v14 string = "hello"
                _ = v14
            case "":
                _ = len(v1)
            default:
                fmt.Println(v4, v2, v2)
            }
            fmt.Println(v4, v2)
            for fi15 := 0; fi15 < 6; fi15++ {
                var v16 string = "hello"
                _ = fi15
                _ = v16
            }
        } else {
            var v17 int16 = 0
            var v18 uint64 = 457
            _ = v17
            _ = v18
        }
    }
    var v19 uint64 = 42
    _ = len(v1)
    fmt.Println(v3, v4, v2)
    for fi20 := 0; fi20 < 6; fi20++ {
        v2 += (v2 * v2)
        _ = fi20
    }
    v4 -= v4
    v2 -= v2
    var v21 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
