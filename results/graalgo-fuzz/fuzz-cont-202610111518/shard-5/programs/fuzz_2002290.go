package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 bool }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 int32, p1 int32, p2 int16) bool {
    switch "a" {
    case "foo bar":
        for fi5 := 0; fi5 < 1; fi5++ {
            v6 := GenS1{F0: 358, F1: 588}
            p0 += (p0 ^ p0)
            _ = fi5
            _ = v6
        }
    case "hello":
        switch "" {
        case "\n\t":
            if true {
                p1 -= (-(p1))
            } else {
                p1 = p1
            }
        case "a":
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := GenS2{F0: 0, F1: (-102 != 0), F2: false}
                _ = fi7
                _ = v8
            }
        case "hello":
            var v9 bool = (!(17 < 544))
            _ = v9
        }
    }
    var v10 string = "foo bar"
    var v11 string = ("a" + v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return true
}

func main() {
    var v1 int8 = 100
    var v2 uint64 = 927
    v3 := make(map[int]bool)
    var v4 uint32 = 0
    var v12 string = "a"
    fmt.Println(v1, v2)
    v13 := GenS3{F0: 128, F1: 65535}
    v14 := GenS2{F0: (v4 / ((v4) | 1)), F1: false, F2: false}
    v12 = ("foo bar" + (("a" + "a") + "hello"))
    fmt.Println(v12)
    if false {
        if true {
            v3[0] = false
            fmt.Println(v4, v2)
            switch "hello" {
            case "foo bar":
                v12 = "hello"
            case "hello":
                v1 = 127
            default:
                var v15 string = v12
                _ = v15
            }
        }
    }
    switch 32767 {
    case 32767:
        var v16 int16 = 42
        _ = v16
    case -569:
        fmt.Println(v4)
    case 39:
        var v17 float64 = (0.1 / 1.0)
        _ = v17
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
