package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 int8 = 127
    v4 := make(map[uint32]uint32)
    v5 := []int{-2147483648, 32767}
    if true {
        v6 := GenS1{F0: 0}
        _ = v6
    }
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            v7 = (^(v7))
            _ = len(v2)
            _ = fi9
        }
        _ = fi8
    }
    v10 := GenS2{F0: 32767}
    var v11 uint8 = v7
    _ = len(v2)
    _ = len(v4)
    var v12 int8 = v3
    v3 += 127
    switch 42 {
    case 128:
        switch (("\n\t" + "hello") + "") {
        case "foo bar":
            switch ("" + "hello") {
            case "hello":
                v13 := GenS3{F0: false}
                _ = v13
            case "\n\t":
                v3 = 127
            default:
                v5 = append(v5, 9223372036854775807)
            }
        case "hello":
            var v14 uint = 2
            _ = v14
        }
    default:
        v15 := GenS1{F0: 451}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
