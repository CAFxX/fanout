package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    v3 := []int{227, 32768}
    v4 := make(map[int]string)
    fmt.Println(v1)
    v5 := GenS1{F0: "a", F1: 0}
    var v6 int64 = 1
    v6 = 1
    var v7 int64 = (v1 + (v6 % ((v1) | 1)))
    delete(v4, -1)
    if true {
        var v8 uint8 = 128
        v9 := GenS2{F0: (222 == 18446744073709551615)}
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: "", F1: 1}
    var v11 int16 = -315
    if (!(42 == 23)) {
        if (1 < 65535) {
            var v12 int = 2147483647
            v13 := GenS1{F0: "\n\t", F1: 0}
            var v14 uint8 = 127
            _ = v12
            _ = v13
            _ = v14
        } else {
            fmt.Println(v6)
        }
        var v15 int64 = v6
        _ = v15
    }
    v6 = (v1 / ((v7) | 1))
    if (488 != 408) {
        var v16 uint8 = 42
        var v17 int64 = v6
        v11 += (v11 ^ (v11 - (-(v11))))
        _ = v16
        _ = v17
    }
    if ((!false) || (4294967295 <= 128)) {
        v4[9223372036854775807] = ("hello" + "")
        var v18 string = "\n\t"
        _ = v18
    } else {
        var v19 float32 = 0.1
        var v20 float64 = 0.0
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
