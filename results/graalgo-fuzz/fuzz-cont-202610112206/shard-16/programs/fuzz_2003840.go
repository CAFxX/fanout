package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) int {
    v5 := GenS1{F0: (p0 << ((65535) & 63)), F1: 42}
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]int)
    var v3 uint16 = 255
    v4 := make(map[int]string)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int64 = (v1 % ((0) | 1))
        v7 = 0
        _ = fi6
        _ = v7
    }
    v3 -= (v3 - 1)
    if true {
        v3 = 42
    }
    var v8 uint = 759
    switch 0 {
    case 476:
        var v9 int32 = 62
        _ = v9
    case 729:
        v3 = v3
    }
    v1 -= (v1 + v1)
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v1)
        delete(v2, "\n\t")
        if false {
            var v11 int = (-(fi10))
            var v12 int8 = -128
            _ = v11
            _ = v12
        } else {
            v8 = (v8 | (v8 ^ (^(v8))))
        }
        _ = fi10
    }
    v8 = v8
    v13 := GenS1{F0: 455, F1: 2}
    var v14 uint8 = 1
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS2{F0: "\n\t", F1: (^(v8)), F2: (v1 % ((v1) | 1))}
        _ = fi15
        _ = v16
    }
    var v17 uint64 = uint64(v3)
    v18 := []int{}
    fmt.Println(v8)
    switch ("foo bar" + "\n\t") {
    case "a":
        var v19 uint64 = (v17 & uint64(v3))
        _ = v19
    }
    v8 += 2
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
