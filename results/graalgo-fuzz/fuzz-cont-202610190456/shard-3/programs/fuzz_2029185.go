package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    var v1 float32 = -2.25
    var v2 string = "a"
    var v3 int64 = 0
    switch -968 {
    case -806:
        v4 := []int{1}
        _ = v4
    }
    if true {
        v1 = v1
        switch 42 {
        case 132:
            var v5 int64 = (v3 + v3)
            _ = v5
        }
    }
    v6 := GenS1{F0: ("" + (v2 + v2)), F1: 219}
    v7 := GenS1{F0: (v2 + v2), F1: 2}
    if ("hello" < ("a" + "hello")) {
        fmt.Println(v1, v3, v3)
        var v8 float32 = 1e10
        switch ((v2 + v2) + v2) {
        case "":
            v9 := []int{9223372036854775807}
            _ = v9
        default:
            v10 := []int{-2147483648, 1}
            _ = v10
        }
        _ = v8
    } else {
        v3 -= v3
        switch "\n\t" {
        case "hello":
            v1 = v1
        }
    }
    v11 := []int{134, 2147483647, 42}
    var v12 int8 = -68
    v1 += 0.0
    v12 = -82
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
