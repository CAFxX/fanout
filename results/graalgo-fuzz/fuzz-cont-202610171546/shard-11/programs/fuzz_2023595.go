package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 string; F1 int16; F2 int }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 uint32, p2 uint16) bool {
    switch "foo bar" {
    case "hello":
        v6 := []int{32767}
        _ = v6
    case "":
        var v7 bool = false
        _ = v7
    case "foo bar":
        switch "hello" {
        case "a":
            switch int(p0) {
            case 42:
                fmt.Println(p1, p2)
            case 127:
                var v8 int32 = 2147483647
                _ = v8
            case -1:
                v9 := GenS1{F0: 985, F1: (p0 + (p0 - 42))}
                _ = v9
            }
        default:
            p2 -= p2
        }
    default:
        p0 -= (p0 << (((p0 / ((p0) | 1))) & 63))
    }
    switch -167 {
    case 42:
        if true {
            fmt.Println(p2, p1, p0)
        }
    case -52:
        if true {
            var v10 uint64 = 324
            _ = v10
        }
    case -28:
        var v11 int32 = 42
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return (false || true)
}

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 1
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    var v12 string = "\n\t"
    _ = len(v3)
    var v13 float64 = 1e10
    var v14 bool = false
    switch "" {
    case "hello":
        var v15 float64 = v13
        _ = v15
    case "a":
        v16 := GenS2{F0: v1, F1: 32767, F2: 42}
        _ = v16
    case "foo bar":
        fmt.Println(v13, v13, v14)
    default:
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 int64 = 0
            var v19 int64 = (v18 + (v18 << ((18446744073709551615) & 63)))
            var v20 int64 = (-(v19))
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v14)
        _ = fi21
    }
    switch v12 {
    case "\n\t":
        var v22 float64 = v13
        _ = v22
    default:
        v23 := GenS3{F0: 32768}
        _ = v23
    }
    v24 := GenS3{F0: 2147483647}
    v25 := GenS2{F0: (((("a" + v1) + "") + "\n\t") + ((v12 + "\n\t") + "\n\t")), F1: int16(v2), F2: 4}
    v12 = ("" + (("hello" + ((v12 + v12) + v12)) + "hello"))
    var v26 int64 = -885
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
