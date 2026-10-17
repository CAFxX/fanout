package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 153
    var v3 uint32 = 1
    v4 := []int{}
    fmt.Println(v2, v3)
    var v5 string = ("" + "foo bar")
    if true {
        var v6 uint16 = 0
        _ = v6
    }
    v3 = v3
    var v7 uint32 = (v3 & (v3 + 658))
    _ = len(v1)
    fmt.Println(v2, v5)
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v2, v3)
        v5 = "\n\t"
        fmt.Println(v7, v7, fi8)
        _ = fi8
    }
    var v9 int16 = 525
    for fi10 := 0; fi10 < 5; fi10++ {
        v5 = v5
        var v11 float32 = -2.25
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: v2, F1: 42}
    switch ((((v5 + ((("a" + (("a" + (v5 + v5)) + "\n\t")) + (v5 + "foo bar")) + "\n\t")) + v5) + (v5 + v5)) + v5) {
    case "a":
        if (2 > 0) {
            v2 += v2
            v5 = ("hello" + "\n\t")
        } else {
            v4 = append(v4, -1)
        }
    case "foo bar":
        var v13 uint32 = v7
        _ = v13
    case "":
        switch "\n\t" {
        case "foo bar":
            v4 = append(v4, 0)
        default:
            var v14 float64 = 3.14159
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 14; fi15++ {
            v16 := GenS2{F0: -455, F1: 2}
            v17 := GenS3{F0: v9}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    if true {
        var v18 uint8 = 128
        v7 = (v3 % ((v7) | 1))
        _ = v18
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
