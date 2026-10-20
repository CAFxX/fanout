package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }

func main() {
    v1 := []int{2147483647, -1, 683}
    v2 := make(map[uint32]uint32)
    var v3 bool = (!true)
    var v4 int32 = -1
    var v5 int = -100
    v6 := make(map[string]uint32)
    var v7 uint64 = 923
    for fi8 := 0; fi8 < 18; fi8++ {
        switch "foo bar" {
        case "a":
            var v9 string = "\n\t"
            _ = v9
        }
        v3 = true
        _ = fi8
    }
    v5 += v5
    fmt.Println(v7)
    switch "foo bar" {
    case "a":
        v7 -= (^(v7))
    case "hello":
        var v10 uint8 = 2
        _ = v10
    }
    var v11 int64 = 42
    if (v3 && v3) {
        var v12 uint = 1
        if (uint64(v12) < v7) {
            _ = len(v2)
            v13 := GenS1{F0: 4294967295, F1: v7}
            var v14 float64 = float64((0.1 + (((1e10 + 1e10) / 1e10) - 3.14159)))
            _ = v13
            _ = v14
        }
        _ = v12
    } else {
        v5 -= v5
        for fi15 := 0; fi15 < 9; fi15++ {
            if (!((false && v3) && (!v3))) {
                v16 := GenS2{F0: -783, F1: -128}
                _ = v16
            }
            switch "a" {
            case "":
                v17 := GenS3{F0: 1, F1: v11, F2: (v4 % (((v4 & v4)) | 1))}
                _ = v17
            }
            var v18 uint32 = 2147483648
            _ = fi15
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(v11, v3)
        v4 = (^(v4))
        if v3 {
            _ = len(v1)
        }
        _ = fi19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        v11 -= (v11 | v11)
        fmt.Println(v5, v11, v5)
        _ = fi20
    }
    v7 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
