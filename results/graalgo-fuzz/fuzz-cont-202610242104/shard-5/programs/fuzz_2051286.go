package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint, p1 int) uint64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int64 = 9223372036854775807
        if true {
            var v6 uint16 = 1
            v7 := []int{p1, p1, (-(fi4)), -773}
            fmt.Println(p1, v6, p0)
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int = fi4
                switch ("a" + "foo bar") {
                case "\n\t":
                    var v10 uint = 0
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            var v11 uint = (p0 >> (((p0 << (((p0 / (((p0 ^ (p0 << ((227) & 63)))) | 1))) & 63))) & 63))
            var v12 int16 = 1
            _ = v11
            _ = v12
        }
        _ = fi4
        _ = v5
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 989}
        fmt.Println(fi13, p1)
        _ = fi13
        _ = v14
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p1)
    case "hello":
        v15 := GenS2{F0: 42, F1: (p1 << ((p0) & 63))}
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    fmt.Println(v1)
    v16 := []int{-9223372036854775808, 127}
    switch -9223372036854775808 {
    case 0:
        v1 += 90
    case 127:
        v1 = (v1 & v1)
    case -286:
        var v17 bool = (2 < 1)
        _ = v17
    default:
        v1 = (v1 << ((4294967295) & 63))
    }
    var v18 uint32 = 2147483647
    v1 += (v1 << ((2) & 63))
    var v19 int = 127
    if true {
        v18 += v18
        var v20 int16 = 1
        var v21 uint8 = 249
        _ = v20
        _ = v21
    }
    if false {
        switch ((("\n\t" + ("" + "foo bar")) + "\n\t") + "a") {
        case "foo bar":
            v22 := GenS3{F0: 2}
            _ = v22
        }
        v23 := v3[v19]
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
