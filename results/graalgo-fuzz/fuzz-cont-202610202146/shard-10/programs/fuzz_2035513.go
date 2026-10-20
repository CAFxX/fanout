package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 uint, p1 bool) uint64 {
    p1 = p1
    switch "hello" {
    case "foo bar":
        p0 += (^(p0))
    case "hello":
        v5 := GenS1{F0: p1, F1: 0}
        _ = v5
    default:
        var v6 int64 = 381
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint8 = 1
        if true {
            v8 = (v8 / ((v8) | 1))
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    return 186
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 2
    v3 := make(map[uint32]bool)
    v4 := []int{-962}
    var v9 int32 = -692
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 int64 = 9223372036854775807
        _ = fi10
        _ = v11
    }
    switch "\n\t" {
    case "a":
        if (!(895 == 379)) {
            var v12 int32 = 672
            var v13 int32 = (v12 >> ((2) & 63))
            var v14 int32 = v13
            _ = v12
            _ = v13
            _ = v14
        } else {
            v9 = (^(v9))
        }
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        fmt.Println(v9, v2)
    }
    switch 9223372036854775807 {
    case -2147483648:
        switch (((("\n\t" + ("a" + "hello")) + "foo bar") + "") + "") {
        case "":
            v9 += (v9 << ((0) & 63))
        }
    case -754:
        switch "" {
        case "\n\t":
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint8 = 128
                _ = fi15
                _ = v16
            }
        }
    default:
        var v17 float32 = -2.25
        _ = v17
    }
    fmt.Println(v2, v2)
    v2 -= v2
    var v18 float64 = 1e10
    fmt.Println(v2, v18)
    v1["foo bar"] = true
    v19 := GenS1{F0: (!true), F1: -61}
    v20 := GenS2{F0: 125, F1: 19, F2: 127}
    var v21 uint32 = 1
    if (v18 <= v18) {
        fmt.Println(v2, v9, v9)
        v22 := GenS3{F0: v21, F1: -1}
        if true {
            var v23 string = "hello"
            v18 -= v18
            var v24 uint64 = 0
            _ = v23
            _ = v24
        } else {
            var v25 float32 = 3.14159
            _ = v25
        }
        _ = v22
    }
    v26 := GenS4{F0: ("\n\t" + ((("\n\t" + "hello") + ("" + "hello")) + "foo bar")), F1: 42, F2: -1}
    _ = len(v1)
    fmt.Println(v18, v9, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
