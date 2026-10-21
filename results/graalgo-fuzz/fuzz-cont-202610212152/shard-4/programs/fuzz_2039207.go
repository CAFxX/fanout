package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
func genFn0(p0 uint, p1 uint16) uint {
    fmt.Println(p1)
    var v5 uint32 = 10
    _ = p0
    _ = p1
    _ = v5
    return (p0 & p0)
}

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 32767
    v3 := []uint32{42, 2, 251, 0}
    v4 := make(map[uint32]int)
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 1, F1: 127, F2: fi6}
        _ = fi6
        _ = v7
    }
    v1 = (((v1 - v1) + v1) * (v1 + ((float32(3.14159) - (v1 + v1)) - v1)))
    fmt.Println(v1, v1)
    v2 += v2
    v1 += v1
    fmt.Println(v1, v2)
    v2 += (v2 % (((v2 * v2)) | 1))
    switch ("" + "hello") {
    case "a":
        var v8 uint16 = v2
        _ = v8
    case "hello":
        if (!false) {
            if (0 != 757) {
                for fi9 := 0; fi9 < 20; fi9++ {
                    switch "\n\t" {
                    case "hello":
                        var v10 uint32 = 42
                        _ = v10
                    case "":
                        for fi11 := 0; fi11 < 15; fi11++ {
                            var v12 float32 = (v1 - (float32((0.1 - 1.5)) - 0.0))
                            _ = fi11
                            _ = v12
                        }
                    case "\n\t":
                        fmt.Println(v1, fi9, v1)
                    }
                    _ = fi9
                }
                v1 += v1
            }
            var v13 int8 = -120
            _ = v13
        } else {
            var v14 uint8 = 127
            v1 += v1
            _ = v14
        }
    }
    switch "foo bar" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            var v15 uint32 = 0
            _ = v15
        case "":
            fmt.Println(v2)
        case "a":
            var v16 float32 = (1e10 - 3.14159)
            _ = v16
        }
    case "foo bar":
        switch 128 {
        case 0:
            v2 += v2
        case -9223372036854775808:
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 18; fi18++ {
                    v1 -= v1
                    switch (fi18 * 4294967295) {
                    case -1:
                        v1 = 1e10
                    case 1:
                        fmt.Println(fi18, v2, fi17)
                    case 128:
                        delete(v4, uint32(v2))
                    default:
                        var v19 string = "\n\t"
                        _ = v19
                    }
                    delete(v4, 834)
                    _ = fi18
                }
                fmt.Println(fi17, v2)
                fmt.Println(v1, v1, v2)
                _ = fi17
            }
        }
    default:
        var v20 bool = false
        _ = v20
    }
    v1 = (float32((1.5 / ((0.1 * 3.14159) - 3.14159))) - -2.25)
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 5; fi22++ {
            _ = len(v3)
            var v23 float64 = (-2.25 - (1e10 / (1e10 + 1.5)))
            _ = fi22
            _ = v23
        }
        var v24 float64 = 1e10
        var v25 uint32 = 1
        _ = fi21
        _ = v24
        _ = v25
    }
    v1 = ((0.0 * v1) * (v1 + (1.5 + v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
