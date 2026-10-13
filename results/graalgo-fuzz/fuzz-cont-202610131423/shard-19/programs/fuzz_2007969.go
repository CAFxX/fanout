package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }

func main() {
    v1 := []string{""}
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 uint = 18446744073709551615
    v2 = (v2 % ((4294967295) | 1))
    fmt.Println(v2)
    v6 := GenS1{F0: 88, F1: 9223372036854775807}
    var v7 int64 = 0
    v8 := GenS1{F0: 127, F1: 4294967295}
    v2 = (v2 / ((397) | 1))
    v2 = v2
    if true {
        fmt.Println(v7)
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            v5 = (v5 % ((1) | 1))
            switch -9223372036854775808 {
            case 9223372036854775807:
                var v10 float32 = float32(0.1)
                _ = v10
            case -23:
                var v11 int16 = -32768
                _ = v11
            case 32768:
                var v12 uint32 = 4294967295
                _ = v12
            }
            _ = fi9
        }
        switch "" {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "":
            var v13 string = ("foo bar" + "foo bar")
            _ = v13
        }
    }
    v7 = (v7 & -9223372036854775808)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
