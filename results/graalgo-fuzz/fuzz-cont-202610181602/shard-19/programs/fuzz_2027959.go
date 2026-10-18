package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = 96
    v2 := make(map[int]int)
    var v3 int64 = -279
    v4 := v2[79]
    if false {
        var v5 int64 = (-(v3))
        _ = v5
    }
    v4 -= v4
    switch "hello" {
    case "\n\t":
        var v6 bool = false
        _ = v6
    }
    v2[-580] = v4
    var v7 float32 = (0.1 / 1.0)
    if true {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(v3, v1, v3)
            if ((("\n\t" + ("a" + "\n\t")) <= (("foo bar" + ("hello" + ("foo bar" + "hello"))) + "hello")) && false) {
                fmt.Println(v7, fi8)
                var v9 int8 = v1
                _ = v9
            }
            var v10 string = "hello"
            _ = fi8
            _ = v10
        }
        if false {
            fmt.Println(v1, v7, v1)
            for fi11 := 0; fi11 < 8; fi11++ {
                v3 -= (v3 >> ((2) & 63))
                v12 := v2[576]
                _ = fi11
                _ = v12
            }
        }
        v13 := GenS1{F0: -387, F1: 2}
        _ = v13
    }
    v7 += v7
    switch ("foo bar" + ((("hello" + "") + "\n\t") + (("a" + "hello") + (((("" + ("hello" + "hello")) + ("" + "")) + (("foo bar" + "hello") + "hello")) + (("a" + "a") + ("" + "")))))) {
    case "\n\t":
        var v14 float64 = ((1.5 / 0.1) + (0.1 * 1.5))
        _ = v14
    default:
        v15 := []int{}
        _ = v15
    }
    var v16 int32 = -1
    v17 := []int{(v4 | v4), v4, 2147483647}
    if false {
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(v4)
            v19 := GenS1{F0: (fi18 << ((1) & 63)), F1: 1}
            switch v4 {
            case -1:
                var v20 uint8 = 116
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        var v21 float64 = (1e10 / 1e10)
        _ = v21
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        v1 -= (v1 + 1)
        var v23 uint = 65535
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: (v1 % ((v1) | 1))}
    var v25 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v7
}
