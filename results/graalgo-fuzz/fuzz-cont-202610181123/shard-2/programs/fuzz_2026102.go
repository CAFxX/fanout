package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8 }

func main() {
    var v1 bool = (2 >= 0)
    var v2 uint8 = 255
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 32767, F1: 542}
    fmt.Println(v2, v1)
    if v1 {
        v5 := GenS2{F0: 0}
        var v6 int16 = 56
        if false {
            var v7 uint = 1
            v2 = (v2 ^ 1)
            var v8 uint8 = v2
            _ = v7
            _ = v8
        } else {
            switch -536 {
            case 0:
                var v9 uint = 169
                _ = v9
            }
            v10 := GenS2{F0: 255}
            var v11 bool = v1
            _ = v10
            _ = v11
        }
        _ = v5
        _ = v6
    }
    fmt.Println(v2, v2)
    var v12 uint = 442
    v13 := []int{}
    v14 := GenS3{F0: "foo bar", F1: -128}
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint32 = 749
        _ = fi15
        _ = v16
    }
    switch "" {
    case "\n\t":
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = fi17
            var v19 string = "hello"
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int = 1
            v2 = uint8(v12)
            _ = fi20
            _ = v21
        }
    }
    switch ("a" + ("" + "")) {
    case "\n\t":
        var v22 float64 = (float64(((1e10 + 1e10) / 3.14159)) / 3.14159)
        _ = v22
    case "hello":
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v13) > 0 {
                v13[0] = (fi23 | (fi23 - int(v2)))
            }
            var v24 uint64 = 18446744073709551615
            _ = fi23
            _ = v24
        }
    case "":
        var v25 string = "\n\t"
        _ = v25
    }
    v12 = 467
    var v26 uint16 = 38
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
