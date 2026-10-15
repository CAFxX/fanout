package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 bool; F1 bool }
func genFn0(p0 uint32, p1 uint16, p2 int64) string {
    fmt.Println(p2)
    v5 := GenS1{F0: 0}
    switch "a" {
    case "\n\t":
        fmt.Println(p1, p0)
    default:
        v6 := GenS2{F0: -114, F1: p0}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]string)
    var v4 string = "foo bar"
    var v7 int32 = 42
    fmt.Println(v7)
    switch (v4 + ("a" + ("hello" + "hello"))) {
    case "\n\t":
        var v8 float32 = 1.5
        _ = v8
    case "":
        v7 -= -728
    }
    var v9 int8 = -1
    v10 := []int{9223372036854775807, 32767}
    var v11 int32 = v7
    v12 := GenS2{F0: 1, F1: 65535}
    v13 := GenS3{F0: 42, F1: 68}
    for fi14 := 0; fi14 < 9; fi14++ {
        if (false && (401 > v2)) {
            v2 += v2
        }
        v4 = v4
        if (4294967295 <= 65535) {
            var v15 bool = (fi14 <= -470)
            _ = v15
        } else {
            fmt.Println(v4, fi14, v9)
            var v16 float32 = 3.14159
            v1[fi14] = (v4 + "foo bar")
            _ = v16
        }
        _ = fi14
    }
    v11 = v11
    var v17 uint8 = 0
    var v18 int8 = (^(v9))
    var v19 int16 = 0
    switch "a" {
    case "\n\t":
        fmt.Println(v17, v2, v17)
    case "":
        v20 := GenS4{F0: v4, F1: v2}
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 float32 = ((0.1 + (0.0 + 1.5)) * 1.5)
            _ = fi21
            _ = v22
        }
    default:
        v23 := GenS5{F0: false, F1: false}
        _ = v23
    }
    fmt.Println(v4)
    v9 -= v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
