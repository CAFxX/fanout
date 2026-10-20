package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 uint32 = 2
    v4 = v4
    v5 := GenS1{F0: 605, F1: 464, F2: 41}
    var v6 uint8 = 42
    v6 = 0
    v7 := GenS1{F0: 1, F1: 0, F2: -128}
    var v8 uint32 = v4
    _ = len(v3)
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(fi9, v4, fi9)
        v8 = (v8 - (v4 >> ((0) & 63)))
        if (42 != -9223372036854775808) {
            var v10 int32 = -1
            fmt.Println(fi9, v8)
            switch "\n\t" {
            case "hello":
                var v11 bool = (0 < 18446744073709551615)
                _ = v11
            default:
                var v12 int8 = -1
                _ = v12
            }
            _ = v10
        }
        _ = fi9
    }
    if ((false && true) || (0 != 42)) {
        v8 = v8
    }
    v13 := []int{-317, 553, 529, -404}
    fmt.Println(v8, v8)
    var v14 int16 = int16(v6)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: (^(fi15)), F1: 0, F2: 42}
        _ = fi15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
