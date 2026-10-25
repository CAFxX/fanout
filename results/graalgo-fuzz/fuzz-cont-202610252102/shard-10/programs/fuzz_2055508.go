package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 int; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int64 }

func main() {
    v1 := []string{}
    var v2 float64 = 3.14159
    var v3 uint64 = 1
    var v4 float32 = 1.5
    var v5 float64 = float64(0.0)
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: (^(fi7)), F1: 9223372036854775807, F2: true}
        fmt.Println(fi7, v5)
        var v9 int8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    switch -239 {
    case 42:
        for fi10 := 0; fi10 < 16; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
    case 1:
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                v3 += (v3 / ((v3) | 1))
            }
        }
    case -524:
        var v11 string = "a"
        _ = v11
    default:
        v12 := GenS2{F0: 42, F1: false, F2: 4294967295}
        _ = v12
    }
    v5 = 0.1
    if (0 == 46) {
        _ = len(v6)
        for fi13 := 0; fi13 < 18; fi13++ {
            v4 = 1.5
            var v14 uint = 2
            _ = fi13
            _ = v14
        }
        var v15 float64 = (((3.14159 - (float64(v4) * 1e10)) + -2.25) + (((v2 + v5) + v5) + (v2 + 1e10)))
        _ = v15
    } else {
        v2 -= ((0.0 * (3.14159 + v5)) * v5)
    }
    v16 := GenS1{F0: 127, F1: 0, F2: ((2147483647 > 1) && true)}
    if false {
        if true {
            var v17 bool = (!true)
            if v17 {
                if (!true) {
                    fmt.Println(v5)
                    var v18 uint64 = v3
                    _ = v18
                }
                _ = len(v6)
            } else {
                switch "foo bar" {
                case "":
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                case "\n\t":
                    var v19 float64 = 0.1
                    _ = v19
                case "foo bar":
                    if ("hello" <= "\n\t") {
                        fmt.Println(v17, v2, v3)
                    }
                default:
                    _ = len(v1)
                }
                var v20 int32 = 876
                _ = v20
            }
            v2 -= ((v5 + v5) * v5)
            _ = v17
        }
        v5 -= ((v5 + v2) * (v2 * -2.25))
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int16 = -1
            _ = fi21
            _ = v22
        }
    } else {
        v23 := v6[(("" + "a") + (("a" + "hello") + "a"))]
        if len(v1) > 0 {
            v1[0] = ("foo bar" + "foo bar")
        }
        _ = v23
    }
    if (9223372036854775807 < 964) {
        var v24 uint8 = 255
        v25 := GenS3{F0: v24, F1: v3, F2: 168}
        _ = v24
        _ = v25
    }
    if (2 == 2) {
        fmt.Println(v4)
    }
    var v26 float32 = (0.1 / float32(v5))
    var v27 uint64 = v3
    var v28 string = ("foo bar" + (("hello" + "") + "\n\t"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
