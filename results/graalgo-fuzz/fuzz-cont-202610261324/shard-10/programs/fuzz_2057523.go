package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint32) int8 {
    v5 := GenS1{F0: true, F1: 2}
    p0 = p0
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v5
    return -128
}

func main() {
    v1 := []string{"\n\t", "foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    for fi6 := 0; fi6 < 16; fi6++ {
        if len(v1) > 0 {
            v1[0] = v4
        }
        var v7 int16 = 784
        if true {
            v4 = (v4 + "a")
        }
        _ = fi6
        _ = v7
    }
    _ = len(v3)
    switch v4 {
    case "a":
        var v8 int = 42
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 6; fi9++ {
            switch (fi9 | (fi9 / ((fi9) | 1))) {
            case 989:
                v4 = ((v4 + v4) + v4)
            default:
                fmt.Println(fi9, v4)
            }
            if (-1 == -32768) {
                var v10 bool = false
                _ = v10
            }
            v11 := GenS2{F0: -53}
            _ = fi9
            _ = v11
        }
    }
    var v12 int64 = 142
    fmt.Println(v12, v12)
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = false
        _ = fi13
        _ = v14
    }
    fmt.Println(v12)
    var v15 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
