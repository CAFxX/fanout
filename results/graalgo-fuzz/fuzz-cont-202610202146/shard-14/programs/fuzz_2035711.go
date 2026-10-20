package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int, p1 string, p2 int8) bool {
    var v5 uint64 = 328
    p0 = 1
    var v6 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint16, p1 int8, p2 bool) uint8 {
    var v7 uint8 = 0
    p2 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (v7 & (v7 & v7))
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 uint = 18446744073709551615
    v4 := []string{("hello" + "\n\t")}
    _ = len(v2)
    var v8 int = 127
    var v9 int = 128
    switch v8 {
    case 32767:
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(v9, fi10)
                switch "\n\t" {
                case "\n\t":
                    v11 := v1[""]
                    _ = v11
                case "hello":
                    var v12 float64 = ((-2.25 * float64(1e10)) * ((0.1 - -2.25) - (((0.0 - -2.25) * 1.5) / 1.0)))
                    _ = v12
                default:
                    v13 := GenS1{F0: 488, F1: true}
                    _ = v13
                }
                switch (fi10 | 42) {
                case -236:
                    if (128 == 2) {
                        var v14 float64 = 1.5
                        _ = v14
                    } else {
                        v3 = (^(v3))
                        v15 := GenS2{F0: -1, F1: -32768, F2: 0}
                        for fi16 := 0; fi16 < 8; fi16++ {
                            v17 := GenS2{F0: 816, F1: -179, F2: -892}
                            v3 = 495
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                }
                _ = fi10
            }
        }
    default:
        fmt.Println(v8, v3)
    }
    var v18 int16 = -398
    var v19 uint8 = 215
    v3 -= 18446744073709551615
    _ = len(v1)
    switch "a" {
    case "\n\t":
        var v20 int8 = 42
        _ = v20
    case "hello":
        _ = len(v2)
    case "foo bar":
        var v21 int64 = 42
        _ = v21
    default:
        for fi22 := 0; fi22 < 15; fi22++ {
            v18 = (v18 + v18)
            _ = fi22
        }
    }
    var v23 uint32 = 0
    v19 += (v19 | 42)
    v9 -= 263
    var v24 int = int(v23)
    v24 += v8
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(fi25, v24)
        v26 := GenS2{F0: 124, F1: (v18 - (v18 + v18)), F2: 1}
        _ = fi25
        _ = v26
    }
    v23 = (v23 & (v23 + (v23 | (^(v23)))))
    v23 += (v23 & (v23 & (v23 - v23)))
    v23 = (v23 >> ((v3) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
