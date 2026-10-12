package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 int64 = -1
    var v2 bool = false
    v3 := make(map[string]uint32)
    switch "hello" {
    case "a":
        v4 := GenS1{F0: -1}
        _ = v4
    }
    switch 32768 {
    case -41:
        var v5 uint32 = 175
        _ = v5
    default:
        fmt.Println(v2, v1)
    }
    v6 := []int{370}
    v2 = (!v2)
    var v7 uint64 = 2
    v7 -= v7
    switch (("foo bar" + "foo bar") + "a") {
    case "hello":
        v8 := GenS2{F0: (v7 & 848), F1: 127}
        _ = v8
    }
    fmt.Println(v7, v1)
    v2 = (!(1 != 42))
    if false {
        if v2 {
            var v9 string = ("foo bar" + "foo bar")
            fmt.Println(v7)
            _ = v9
        }
    }
    v10 := GenS1{F0: 1}
    switch 32768 {
    case 32768:
        var v11 float32 = ((1e10 + float32((3.14159 * 0.1))) + 1.5)
        _ = v11
    case 581:
        var v12 int8 = -11
        _ = v12
    case -9223372036854775808:
        var v13 bool = v2
        _ = v13
    default:
        switch 835 {
        case -9223372036854775808:
            var v14 int8 = 100
            _ = v14
        }
    }
    switch "\n\t" {
    case "a":
        v7 -= v7
    default:
        delete(v3, "\n\t")
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
