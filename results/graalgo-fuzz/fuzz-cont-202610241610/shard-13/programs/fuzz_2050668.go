package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int32; F1 uint32; F2 uint32 }

func main() {
    v1 := []string{"\n\t", "\n\t", ""}
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := []int{}
    var v5 bool = true
    var v6 string = "foo bar"
    v6 = v6
    v7 := GenS1{F0: 9223372036854775807, F1: ("a" + v6)}
    var v8 string = v6
    v9 := GenS2{F0: v5, F1: v6}
    v10 := GenS2{F0: (-32768 >= 1), F1: v6}
    switch ((v8 + (v8 + "")) + "\n\t") {
    case "foo bar":
        fmt.Println(v5, v6)
    case "":
        v11 := GenS2{F0: v5, F1: v6}
        _ = v11
    }
    switch 32767 {
    case 1:
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint32 = 825
            var v14 int8 = 100
            _ = fi12
            _ = v13
            _ = v14
        }
    case -703:
        fmt.Println(v8, v8)
    case 32768:
        for fi15 := 0; fi15 < 18; fi15++ {
            v8 = "hello"
            _ = fi15
        }
    default:
        v16 := GenS3{F0: 828, F1: 65535, F2: 715}
        _ = v16
    }
    v17 := GenS3{F0: -278, F1: 4294967295, F2: 2}
    var v18 uint32 = 0
    v19 := GenS1{F0: -9223372036854775808, F1: (((v6 + v8) + "hello") + "a")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
