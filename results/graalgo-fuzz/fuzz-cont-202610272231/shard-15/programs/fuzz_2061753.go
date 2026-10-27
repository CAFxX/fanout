package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int32; F1 string }
func genFn0(p0 int64) int16 {
    if (273 == 0) {
        for fi5 := 0; fi5 < 8; fi5++ {
            p0 = (-(p0))
            switch (fi5 << ((841) & 63)) {
            case 127:
                v6 := GenS1{F0: 0}
                _ = v6
            case -9223372036854775808:
                var v7 int8 = 1
                _ = v7
            default:
                p0 = p0
            }
            _ = fi5
        }
        switch "a" {
        case "foo bar":
            p0 += (p0 + 42)
        }
    }
    if true {
        v8 := []int{-2147483648, 128}
        _ = v8
    } else {
        if (!false) {
            v9 := GenS1{F0: 163}
            p0 += (p0 - (^(p0)))
            v10 := GenS1{F0: 156}
            _ = v10
            _ = v9
        } else {
            p0 += (^(p0))
        }
    }
    fmt.Println(p0)
    p0 = (^(p0))
    _ = p0
    return 1
}
func genFn1(p0 int8, p1 int64, p2 string) uint64 {
    var v11 float32 = float32(((1.5 + float64(-2.25)) - float64(((1e10 * 1e10) + ((-2.25 * float32(0.0)) + 3.14159)))))
    v12 := GenS2{F0: 0, F1: 42}
    var v13 int16 = -292
    switch ((((p2 + (p2 + p2)) + "\n\t") + p2) + p2) {
    case "":
        v14 := GenS3{F0: 989, F1: p2}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 205
    var v4 int8 = 127
    v15 := GenS2{F0: 773, F1: 42}
    fmt.Println(v3, v3)
    switch 32768 {
    case -857:
        for fi16 := 0; fi16 < 8; fi16++ {
            v4 = 127
            for fi17 := 0; fi17 < 1; fi17++ {
                fmt.Println(fi16)
                _ = len(v1)
                v3 += v3
                _ = fi17
            }
            _ = fi16
        }
    case 690:
        if false {
            v2[65535] = 2147483647
        }
    default:
        var v18 uint64 = 18446744073709551615
        _ = v18
    }
    if (-1 == 302) {
        switch 9223372036854775807 {
        case -2147483648:
            var v19 int = 127
            _ = v19
        case 9223372036854775807:
            var v20 string = "\n\t"
            _ = v20
        default:
            if (1e10 != 0.0) {
                fmt.Println(v3, v4)
                var v21 string = ""
                var v22 int8 = (v4 % (((v4 | (-(v4)))) | 1))
                _ = v21
                _ = v22
            }
        }
    } else {
        var v23 uint32 = 65535
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int32 = -824
            v3 = (v3 >> ((0) & 63))
            var v26 uint32 = (v23 * 1)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    var v27 uint64 = 1
    v4 = 1
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 uint8 = 61
        var v30 float32 = (1e10 / float32(0.1))
        switch (fi28 >> ((255) & 63)) {
        case 1:
            var v31 float64 = 1.5
            _ = v31
        case 32767:
            var v32 int64 = 1
            _ = v32
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    v3 += v3
    var v33 float64 = 1e10
    fmt.Println(v4, v4)
    fmt.Println(v4, v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v4
}
