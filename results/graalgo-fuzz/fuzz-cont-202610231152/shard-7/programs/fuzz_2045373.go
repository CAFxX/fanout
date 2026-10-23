package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
type GenS5 struct { F0 uint; F1 int }
func genFn0(p0 string, p1 string) uint16 {
    var v7 int16 = 489
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]string)
    v3 := []string{(((("foo bar" + v1) + v1) + ("foo bar" + v1)) + ("hello" + v1)), "a"}
    var v4 float32 = -2.25
    var v5 uint = 174
    v6 := make(map[int]bool)
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v6)
        var v9 bool = true
        v10 := GenS1{F0: v1, F1: -60}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!false) {
        switch 1 {
        case 32768:
            _ = len(v6)
        case 9223372036854775807:
            v11 := GenS2{F0: 255}
            _ = v11
        default:
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(v1)
                var v13 int64 = -9223372036854775808
                var v14 uint64 = 42
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        var v15 int = 32768
        _ = v15
    }
    var v16 int64 = 404
    switch 9223372036854775807 {
    case 2147483647:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch "hello" {
    case "hello":
        v17 := GenS3{F0: -558, F1: int16(v5)}
        _ = v17
    case "foo bar":
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(v16, v4, v16)
            _ = fi18
        }
    case "":
        v19 := GenS4{F0: (0 < genFn0(v1, v1)), F1: 648, F2: false}
        _ = v19
    }
    v20 := GenS5{F0: v5, F1: 0}
    var v21 int8 = -118
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 int64 = -1
        fmt.Println(v1, v1, v4)
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
