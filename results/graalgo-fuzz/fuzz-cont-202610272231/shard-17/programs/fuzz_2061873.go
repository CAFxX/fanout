package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32, p1 uint16) bool {
    p1 += (p1 + (p1 >> ((65535) & 63)))
    for fi6 := 0; fi6 < 15; fi6++ {
        if (!(!(0.0 <= (0.1 / float32(1.5))))) {
            var v7 uint64 = 1
            _ = v7
        }
        v8 := GenS1{F0: (fi6 & (fi6 + fi6))}
        _ = fi6
        _ = v8
    }
    _ = p0
    _ = p1
    return ((255 == 1) || (((-(p0)) < p0) || true))
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]string)
    v3 := []uint32{816, 65535, 4294967295}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]bool)
    v1 = v1
    if (610 > 2) {
        v9 := GenS2{F0: 1}
        v10 := GenS1{F0: -428}
        v11 := GenS3{F0: ((773 == 0) && true), F1: 568}
        _ = v10
        _ = v11
        _ = v9
    }
    if (4294967295 == 1) {
        fmt.Println(v1)
        if (true || false) {
            if true {
                var v12 float32 = float32(v1)
                _ = v12
            }
            fmt.Println(v1)
            var v13 string = ""
            _ = v13
        } else {
            fmt.Println(v1)
            fmt.Println(v1)
        }
    }
    v14 := GenS1{F0: -1}
    for fi15 := 0; fi15 < 16; fi15++ {
        _ = len(v5)
        _ = fi15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int8 = -128
        var v18 float64 = 0.1
        _ = fi16
        _ = v17
        _ = v18
    }
    v1 -= v1
    if ((42 > 42) || true) {
        fmt.Println(v1)
    }
    v1 += v1
    switch "a" {
    case "a":
        if false {
            var v19 uint = 1
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            switch "foo bar" {
            case "foo bar":
                if false {
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v1 -= -2.25
                        _ = fi21
                    }
                }
            case "\n\t":
                var v22 int16 = -720
                _ = v22
            case "":
                v1 = v1
            default:
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
            var v23 int32 = 42
            _ = fi20
            _ = v23
        }
    }
    v1 = (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
