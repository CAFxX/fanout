package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 bool }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint }

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]string)
    var v3 int16 = -1
    var v4 uint16 = 166
    var v5 string = "foo bar"
    v6 := make(map[int]string)
    fmt.Println(v4, v1)
    v4 = v4
    var v7 int64 = 997
    if (2 <= 2147483648) {
        fmt.Println(v7, v4, v3)
        if false {
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 12; fi9++ {
                    v10 := []int{}
                    v11 := GenS1{F0: true, F1: 42, F2: true}
                    v12 := GenS2{F0: 2147483647, F1: 42}
                    v3 -= (v3 << ((0) & 63))
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = fi8
            }
        }
        var v13 uint64 = 0
        _ = v13
    }
    var v14 int32 = v1
    var v15 uint32 = 833
    switch "\n\t" {
    case "":
        fmt.Println(v15, v5, v14)
    case "a":
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(v7, v4, v1)
            switch (v5 + v5) {
            case "\n\t":
                var v17 float64 = 3.14159
                _ = v17
            }
            _ = fi16
        }
    case "hello":
        fmt.Println(v14)
    default:
        v18 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615}
        _ = v18
    }
    if false {
        fmt.Println(v5, v5, v3)
        switch "a" {
        case "foo bar":
            for fi19 := 0; fi19 < 11; fi19++ {
                v14 -= -908
                var v20 string = (("" + ((v5 + "") + "a")) + v5)
                fmt.Println(v3, v5, v4)
                _ = fi19
                _ = v20
            }
        default:
            v21 := GenS2{F0: 1, F1: 2}
            _ = v21
        }
    }
    delete(v6, 42)
    v22 := []int{32768, -2147483648}
    v6[-2147483648] = "foo bar"
    var v23 string = v5
    v24 := GenS3{F0: uint(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
