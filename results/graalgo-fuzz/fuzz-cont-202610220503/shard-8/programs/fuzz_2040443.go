package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int16, p2 string) uint32 {
    v4 := GenS1{F0: 2, F1: 2, F2: p2}
    var v5 float32 = 0.1
    v6 := GenS1{F0: 2147483647, F1: uint(p0), F2: "\n\t"}
    p1 -= p1
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 int32 = -185
    v2 := make(map[int]int)
    var v3 uint8 = 250
    var v7 uint8 = v3
    switch "a" {
    case "\n\t":
        var v8 uint8 = (^(v3))
        _ = v8
    case "":
        var v9 float32 = 1e10
        _ = v9
    }
    v10 := GenS2{F0: 1}
    var v11 int8 = -128
    var v12 int8 = 42
    v13 := GenS1{F0: 366, F1: 163, F2: "foo bar"}
    for fi14 := 0; fi14 < 5; fi14++ {
        switch (fi14 / ((170) | 1)) {
        case -1:
            delete(v2, -180)
        }
        var v15 int16 = 32767
        v16 := GenS1{F0: 520, F1: 255, F2: "hello"}
        _ = fi14
        _ = v15
        _ = v16
    }
    v1 -= v1
    v12 += (v12 | v12)
    if true {
        fmt.Println(v12)
    }
    switch ("hello" + "foo bar") {
    case "\n\t":
        var v17 uint8 = v3
        _ = v17
    case "foo bar":
        var v18 int64 = -1
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 += (v1 ^ (v1 ^ 42))
        v20 := []int{9223372036854775807, fi19, (^(fi19))}
        _ = fi19
        _ = v20
    }
    var v21 int8 = (v11 - v11)
    var v22 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
