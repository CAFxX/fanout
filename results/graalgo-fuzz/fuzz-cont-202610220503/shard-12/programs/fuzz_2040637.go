package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint8 = 246
    v2 := []string{(("\n\t" + "hello") + ""), "a"}
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    v5 := []string{}
    v3[563] = ((-9223372036854775808 > 42) || false)
    switch "" {
    case "":
        switch 837 {
        case 381:
            _ = len(v4)
        case 424:
            v1 = (v1 ^ (v1 * v1))
        }
    case "hello":
        v1 = 127
    }
    v2 = append(v2, "\n\t")
    v6 := GenS1{F0: false, F1: 509}
    fmt.Println(v1)
    v7 := GenS2{F0: 654, F1: 1}
    var v8 bool = (114 > -1)
    _ = len(v5)
    if (32767 >= 1) {
        fmt.Println(v8, v8)
        v9 := GenS3{F0: "\n\t", F1: 949}
        v1 = 237
        _ = v9
    } else {
        v8 = (!v8)
        var v10 int8 = 100
        _ = v10
    }
    var v11 uint8 = v1
    var v12 bool = v8
    v13 := GenS4{F0: -1, F1: 4}
    v11 = v1
    fmt.Println(v12, v11, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
