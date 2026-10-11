package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 int32 = 0
    v2 := []uint32{2147483648, 2}
    v3 := make(map[string]bool)
    if len(v2) > 0 {
        v2[0] = 468
    }
    switch "" {
    case "":
        _ = len(v2)
    case "foo bar":
        v1 += 2147483647
    case "hello":
        var v4 uint64 = 1
        _ = v4
    }
    if false {
        var v5 int = 241
        v5 = (v5 % (((^(v5))) | 1))
        _ = v5
    }
    fmt.Println(v1)
    for fi6 := 0; fi6 < 6; fi6++ {
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    v7 := v3[""]
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int16 = 32767
        _ = fi8
        _ = v9
    }
    var v10 string = ("" + "hello")
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 float32 = ((-2.25 * 0.0) + -2.25)
        fmt.Println(v7)
        v1 += (v1 << ((664) & 63))
        _ = fi11
        _ = v12
    }
    v7 = v7
    v10 = (v10 + (v10 + "a"))
    v13 := GenS1{F0: 42, F1: 88}
    if false {
        if ("" >= ((v10 + ("" + v10)) + (v10 + v10))) {
            for fi14 := 0; fi14 < 15; fi14++ {
                v15 := GenS1{F0: v1, F1: 0}
                switch (v10 + "") {
                case "\n\t":
                    var v16 int32 = v1
                    _ = v16
                case "hello":
                    for fi17 := 0; fi17 < 2; fi17++ {
                        fmt.Println(v1, v10)
                        var v18 uint64 = 4294967295
                        _ = fi17
                        _ = v18
                    }
                case "foo bar":
                    v2 = append(v2, 65535)
                default:
                    fmt.Println(v10, v7, v10)
                }
                _ = fi14
                _ = v15
            }
            switch "foo bar" {
            case "hello":
                var v19 int32 = (v1 << ((251) & 63))
                _ = v19
            case "a":
                v20 := v3[v10]
                _ = v20
            default:
                v21 := GenS2{F0: 250}
                _ = v21
            }
        } else {
            if (!v7) {
                v7 = v7
                var v22 bool = v7
                _ = v22
            } else {
                v23 := GenS3{F0: -288, F1: 1}
                _ = v23
            }
        }
    } else {
        for fi24 := 0; fi24 < 16; fi24++ {
            v1 += 2147483647
            _ = fi24
        }
        var v25 uint = 2
        _ = v25
    }
    v26 := GenS1{F0: 2147483647, F1: 100}
    var v27 float64 = 0.0
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v7
}
