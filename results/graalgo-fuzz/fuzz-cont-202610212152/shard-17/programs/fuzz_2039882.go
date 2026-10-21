package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 int64) int {
    if (65535 > 65535) {
        var v5 uint = 18446744073709551615
        v6 := GenS1{F0: ((24 >= 385) && (!(("a" + "hello") != "\n\t")))}
        _ = v5
        _ = v6
    }
    var v7 uint = 65535
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 0
    var v2 int = 0
    v3 := make(map[int]uint32)
    v4 := []string{"", "hello"}
    v8 := GenS2{F0: 65535, F1: "foo bar"}
    if ((true || true) || true) {
        if true {
            switch ("foo bar" + ("foo bar" + "hello")) {
            case "":
                v2 = (v2 * (-(v2)))
            case "foo bar":
                v2 = (-(v2))
            case "a":
                v1 -= 42
            }
            var v9 int32 = -2147483648
            v10 := GenS1{F0: (32767 != 2)}
            _ = v10
            _ = v9
        } else {
            var v11 int32 = 42
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := GenS3{F0: "", F1: (v11 >> ((0) & 63))}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    } else {
        v1 -= 0
    }
    var v14 int = v2
    for fi15 := 0; fi15 < 12; fi15++ {
        switch ((("\n\t" + "hello") + ("foo bar" + "foo bar")) + "hello") {
        case "\n\t":
            var v16 int8 = 100
            _ = v16
        }
        v2 += (fi15 | (v2 & v14))
        _ = fi15
    }
    switch "\n\t" {
    case "foo bar":
        var v17 string = "hello"
        _ = v17
    }
    var v18 float64 = 0.0
    v3[(v2 & v2)] = 55
    switch "hello" {
    case "foo bar":
        v18 -= float64(0.0)
    default:
        var v19 int32 = 1
        _ = v19
    }
    if true {
        var v20 int32 = 1
        _ = v20
    } else {
        delete(v3, (v14 ^ v2))
        var v21 uint16 = 2
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
