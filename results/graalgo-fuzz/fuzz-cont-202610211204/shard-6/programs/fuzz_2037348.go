package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := []string{"foo bar", "a"}
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    switch 42 {
    case 127:
        v4 := GenS1{F0: 572, F1: 9223372036854775807}
        _ = v4
    case 535:
        v1 = append(v1, "foo bar")
    case -2147483648:
        var v5 uint8 = 88
        _ = v5
    default:
        switch 127 {
        case 760:
            for fi6 := 0; fi6 < 4; fi6++ {
                for fi7 := 0; fi7 < 8; fi7++ {
                    v8 := GenS2{F0: 4294967295}
                    var v9 uint8 = 128
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                var v10 int64 = 873
                _ = fi6
                _ = v10
            }
        case -59:
            var v11 bool = (!true)
            _ = v11
        default:
            var v12 float64 = 1e10
            _ = v12
        }
    }
    var v13 float32 = (3.14159 - 0.1)
    v13 = ((1.5 / v13) * v13)
    var v14 int16 = -1
    var v15 int32 = 2147483647
    delete(v2, 32767)
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := GenS3{F0: 500}
        if true {
            fmt.Println(fi16, v14, v14)
        }
        _ = fi16
        _ = v17
    }
    if (-1 <= 127) {
        switch "foo bar" {
        case "\n\t":
            v18 := GenS1{F0: 4294967295, F1: 367}
            _ = v18
        case "foo bar":
            v19 := GenS4{F0: (v15 ^ (v15 | v15))}
            _ = v19
        case "a":
            v15 = v15
        default:
            for fi20 := 0; fi20 < 12; fi20++ {
                if len(v1) > 0 {
                    v1[0] = ("hello" + "a")
                }
                var v21 uint8 = 13
                _ = fi20
                _ = v21
            }
        }
        switch "hello" {
        case "foo bar":
            var v22 int8 = 42
            _ = v22
        case "\n\t":
            fmt.Println(v13)
        case "":
            v13 -= 1.5
        }
        v23 := GenS5{F0: 156}
        _ = v23
    } else {
        var v24 int8 = -1
        switch 258 {
        case -60:
            switch "hello" {
            case "\n\t":
                for fi25 := 0; fi25 < 16; fi25++ {
                    fmt.Println(v15, v24, v14)
                    _ = fi25
                }
            case "":
                if false {
                    var v26 string = "hello"
                    _ = v26
                }
            default:
                for fi27 := 0; fi27 < 9; fi27++ {
                    v24 += (v24 / (((v24 >> ((558) & 63))) | 1))
                    _ = fi27
                }
            }
        case 321:
            v28 := GenS5{F0: 2147483647}
            _ = v28
        default:
            v29 := GenS1{F0: 804, F1: 0}
            _ = v29
        }
        _ = v24
    }
    if (1 < -110) {
        fmt.Println(v13, v13)
        v15 += -72
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
