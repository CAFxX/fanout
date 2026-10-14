package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    var v1 string = ("foo bar" + "foo bar")
    v2 := make(map[int]bool)
    v3 := []int{-9223372036854775808}
    v4 := []int{-1, 273, 9223372036854775807, 2147483647}
    var v5 int8 = -128
    if false {
        switch v1 {
        case "hello":
            var v6 int8 = v5
            _ = v6
        case "":
            switch 1 {
            case 2147483647:
                var v7 uint64 = 471
                _ = v7
            }
        case "a":
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float32 = 1e10
                _ = fi8
                _ = v9
            }
        }
        fmt.Println(v5, v1)
        var v10 uint64 = 590
        _ = v10
    } else {
        _ = len(v2)
        var v11 uint8 = 0
        v11 = (v11 >> ((2) & 63))
        _ = v11
    }
    fmt.Println(v1)
    fmt.Println(v1, v5)
    v12 := GenS1{F0: -9223372036854775808}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 bool = false
        v14 = v14
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: v1, F1: 0}
    v16 := GenS3{F0: -2147483648, F1: 2}
    switch v1 {
    case "hello":
        v17 := GenS1{F0: -9223372036854775808}
        _ = v17
    case "a":
        v5 = (v5 + v5)
    case "foo bar":
        v4 = append(v4, -370)
    }
    var v18 string = v1
    var v19 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
