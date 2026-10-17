package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    v3 := []string{}
    var v4 uint64 = 423
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 42}
    var v7 uint16 = 255
    delete(v5, "")
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 -= uint64(v2)
        _ = fi8
    }
    v2 += v2
    fmt.Println(v2, v2)
    switch 42 {
    case 4294967295:
        var v9 string = ("\n\t" + (("a" + "\n\t") + "\n\t"))
        _ = v9
    }
    switch ("\n\t" + "") {
    case "":
        fmt.Println(v7)
    case "\n\t":
        v10 := GenS2{F0: v4}
        _ = v10
    }
    v2 = (v2 - v2)
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: 128}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(fi13, v7)
        v14 := GenS3{F0: v2, F1: -2147483648, F2: (^(v4))}
        v4 = v4
        _ = fi13
        _ = v14
    }
    v2 = 2
    v4 += (v4 % ((v4) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
