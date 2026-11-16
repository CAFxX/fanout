package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    var v1 uint16 = 194
    v2 := []uint32{}
    var v3 bool = (1 == 100)
    var v4 float64 = 1e10
    var v5 string = "foo bar"
    v6 := make(map[string]int)
    var v7 float32 = 3.14159
    _ = len(v6)
    for fi8 := 0; fi8 < 9; fi8++ {
        switch (fi8 >> ((2) & 63)) {
        case -1:
            fmt.Println(v3, v4, v7)
        case 2147483647:
            delete(v6, v5)
        }
        var v9 float64 = 0.0
        if false {
            var v10 bool = (9223372036854775807 > -673)
            v11 := GenS1{F0: 951, F1: v10}
            v7 += v7
            _ = v10
            _ = v11
        } else {
            if (v3 && v3) {
                v3 = (false && (4294967295 > 0))
                v12 := GenS1{F0: 853, F1: (!v3)}
                v1 = 0
                _ = v12
            }
            if true {
                v2 = append(v2, 0)
                var v13 int32 = 1
                _ = v13
            } else {
                fmt.Println(v4, v5, fi8)
            }
        }
        _ = fi8
        _ = v9
    }
    v1 -= v1
    _ = len(v6)
    if (v5 <= "\n\t") {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if v3 {
        switch "\n\t" {
        case "foo bar":
            v5 = "hello"
        case "hello":
            var v14 uint8 = 255
            _ = v14
        }
    }
    v15 := GenS2{F0: ("" + (v5 + "\n\t"))}
    var v16 uint64 = 42
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v5)
        _ = fi17
    }
    var v18 uint64 = v16
    v5 = "foo bar"
    switch 9223372036854775807 {
    case -1:
        v19 := GenS3{F0: 489, F1: 144}
        _ = v19
    default:
        for fi20 := 0; fi20 < 13; fi20++ {
            v5 = "\n\t"
            switch fi20 {
            case 254:
                var v21 bool = v3
                _ = v21
            case -1:
                v7 -= float32(v4)
            case -992:
                v22 := GenS2{F0: v5}
                _ = v22
            }
            _ = fi20
        }
    }
    v3 = (v3 || (false && v3))
    var v23 float32 = v7
    var v24 int64 = 153
    if ((v24 & 9223372036854775807) == int64(v18)) {
        v16 += v18
        var v25 float32 = float32(v4)
        _ = v25
    }
    var v26 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
