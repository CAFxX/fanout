package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 int = 180
    v4 := []int{(v3 << ((18446744073709551615) & 63)), -856, v3, v3}
    var v5 int64 = -851
    v5 += 42
    switch "a" {
    case "foo bar":
        fmt.Println(v5, v5)
    default:
        switch v3 {
        case -614:
            v6 := GenS1{F0: 2}
            _ = v6
        }
    }
    v7 := GenS2{F0: "hello", F1: (v5 << ((4294967295) & 63))}
    v8 := GenS3{F0: 16, F1: 2}
    var v9 int = v3
    _ = len(v1)
    var v10 float64 = -2.25
    switch (("foo bar" + "\n\t") + (("foo bar" + ((("hello" + "foo bar") + ("a" + "foo bar")) + "a")) + ("foo bar" + "hello"))) {
    case "":
        v5 -= 1
    }
    v11 := GenS4{F0: 1, F1: 42, F2: (false && (32767 < (v3 & (v3 | v9))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
