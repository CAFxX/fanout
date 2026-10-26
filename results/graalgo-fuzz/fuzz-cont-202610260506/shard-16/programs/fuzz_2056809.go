package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 string }

func main() {
    v1 := []string{(("a" + "\n\t") + ""), "hello", ("" + "hello"), ""}
    v2 := make(map[int]int)
    var v3 float64 = (-2.25 / 1e10)
    var v4 uint16 = 32767
    var v5 int16 = 32767
    v3 -= ((-2.25 - 0.1) / v3)
    var v6 uint = 18446744073709551615
    for fi7 := 0; fi7 < 5; fi7++ {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = fi7
    }
    v6 += v6
    v8 := GenS1{F0: v4, F1: 652}
    var v9 int64 = 794
    switch int(v6) {
    case 127:
        var v10 float64 = v3
        _ = v10
    case 107:
        v6 = (v6 % ((v6) | 1))
    default:
        v11 := GenS2{F0: uint(v4), F1: 4294967295}
        _ = v11
    }
    v12 := GenS3{F0: "hello"}
    fmt.Println(v6, v4)
    var v13 int64 = v9
    v14 := GenS3{F0: ("hello" + "foo bar")}
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println(v13, v3)
    switch -975 {
    case 4294967295:
        if (true || false) {
            v6 -= 865
            v5 -= (v5 - (v5 / ((v5) | 1)))
        }
    case -135:
        switch "\n\t" {
        case "":
            var v15 float64 = v3
            _ = v15
        case "foo bar":
            var v16 uint8 = 2
            _ = v16
        case "a":
            _ = len(v2)
        }
    case -136:
        v6 += (v6 % (((^(v6))) | 1))
    default:
        v17 := GenS2{F0: (v6 | 1), F1: 476}
        _ = v17
    }
    v6 = (v6 & (^(v6)))
    var v18 float32 = 0.1
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint8 = 99
        var v21 uint16 = (v4 * v4)
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
