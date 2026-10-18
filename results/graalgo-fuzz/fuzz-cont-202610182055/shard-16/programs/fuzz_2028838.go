package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }

func main() {
    var v1 bool = (!((-2.25 + float32(-2.25)) == ((0.1 * 1e10) / 3.14159)))
    v2 := make(map[string]bool)
    var v3 int64 = 361
    v4 := []string{"hello", (("\n\t" + "hello") + "foo bar"), ("a" + ("" + "foo bar"))}
    var v5 int32 = -585
    switch "" {
    case "":
        v3 = (v3 - v3)
    case "hello":
        if true {
            v6 := GenS1{F0: 9223372036854775807, F1: 116, F2: 255}
            _ = v6
        }
    default:
        v5 = v5
    }
    v7 := GenS1{F0: -1, F1: 247, F2: 0}
    v5 = (v5 / (((v5 + -178)) | 1))
    if (!(214 == 227)) {
        var v8 int32 = v5
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = fi9
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if len(v4) > 0 {
            v4[0] = ("a" + (("\n\t" + "\n\t") + "foo bar"))
        }
        _ = fi11
    }
    var v12 string = "a"
    var v13 uint = 42
    if v1 {
        if false {
            v5 += v5
            var v14 int16 = -121
            v13 -= v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            v3 = 375
            var v16 float32 = (1e10 - float32(((float64(0.1) / 1.0) * float64((0.0 - -2.25)))))
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: 0, F1: 255, F2: 32767}
        _ = v17
    } else {
        fmt.Println(v1, v13)
        v12 = "\n\t"
    }
    var v18 uint16 = 192
    var v19 uint16 = 255
    v3 = -9223372036854775808
    for fi20 := 0; fi20 < 15; fi20++ {
        _ = len(v4)
        _ = fi20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        _ = len(v4)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
