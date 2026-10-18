package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int32 = 870
    v2 := make(map[int]int)
    var v3 int32 = v1
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 387, F1: 0}
        v7 := GenS1{F0: 2147483647, F1: 940}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = v4
    v9 := []int{128, 1, 42, 32768}
    v1 += 42
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((-(fi10)) != (fi10 - (^(fi10)))) {
            v11 := GenS2{F0: (!true)}
            v9 = append(v9, fi10)
            _ = v11
        }
        v1 = (v1 + -870)
        _ = fi10
    }
    var v12 bool = (42 < 2147483648)
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v8, v12, fi13)
        v14 := GenS2{F0: v12}
        v8 = 3.14159
        _ = fi13
        _ = v14
    }
    v8 = v4
    for fi15 := 0; fi15 < 15; fi15++ {
        switch "foo bar" {
        case "":
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint32 = 2147483647
                v2[(fi15 / (((fi15 ^ fi15)) | 1))] = (fi16 % (((-(fi16))) | 1))
                _ = fi16
                _ = v17
            }
        }
        _ = fi15
    }
    v18 := GenS2{F0: false}
    v19 := GenS3{F0: 100}
    switch "hello" {
    case "\n\t":
        var v20 uint32 = 2147483648
        _ = v20
    default:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
