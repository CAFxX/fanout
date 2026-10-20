package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }

func main() {
    v1 := []string{"foo bar", "", ("foo bar" + "")}
    var v2 uint8 = 127
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: "\n\t", F1: 0, F2: v2}
    if len(v1) > 0 {
        v1[0] = ""
    }
    v2 = (^(v2))
    if (-32768 != 1000) {
        var v5 int = 1
        v5 = (v5 % ((v5) | 1))
        _ = v5
    }
    var v6 uint8 = (v2 >> ((42) & 63))
    if ((!(true && (!false))) && true) {
        fmt.Println(v2, v6)
        var v7 uint64 = 2
        var v8 int16 = 709
        _ = v7
        _ = v8
    }
    v6 = (v6 / (((v2 << ((0) & 63))) | 1))
    v9 := v3[723]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
