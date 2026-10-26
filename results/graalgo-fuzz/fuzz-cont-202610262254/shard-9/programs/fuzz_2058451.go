package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 bool }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 318, F1: "foo bar"}
            v9 := GenS1{F0: 1000, F1: (("hello" + "foo bar") + (("\n\t" + "hello") + ((("foo bar" + ("" + ("" + ("foo bar" + ("hello" + ("foo bar" + "")))))) + ("a" + ("hello" + ((("a" + "foo bar") + "hello") + ("foo bar" + ""))))) + "foo bar")))}
            v10 := GenS2{F0: 1000, F1: (fi7 >> ((255) & 63))}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 uint = 255
        _ = fi6
        _ = v11
    }
    var v12 float64 = (0.0 - float64((1.5 - float32((3.14159 / 1.0)))))
    switch "hello" {
    case "foo bar":
        v13 := GenS2{F0: 1000, F1: 2147483647}
        _ = v13
    default:
        var v14 int8 = -1
        _ = v14
    }
    if false {
        switch ("a" + "a") {
        case "a":
            var v15 float32 = float32(v12)
            _ = v15
        default:
            var v16 float64 = v12
            _ = v16
        }
        fmt.Println(v12)
    }
    v12 += (v12 + v12)
    _ = v12
    return 255
}
func genFn1(p0 uint64, p1 uint64, p2 int) uint32 {
    if (127 != 157) {
        p2 = p2
        v17 := GenS2{F0: 3, F1: p2}
        _ = v17
    }
    p1 -= 0
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = float32(3.14159)
    v3 := make(map[int]uint32)
    v4 := []string{((v1 + (v1 + (v1 + "hello"))) + v1), "\n\t", ("hello" + v1), v1}
    v5 := []string{(v1 + v1)}
    v18 := GenS1{F0: -377, F1: v1}
    v3[9223372036854775807] = 1
    var v19 int64 = -301
    v5 = append(v5, v1)
    fmt.Println(v1, v1)
    switch "hello" {
    case "hello":
        if true {
            var v20 uint8 = 42
            v1 = v1
            _ = v20
        }
    case "foo bar":
        v21 := GenS1{F0: 0, F1: v1}
        _ = v21
    case "\n\t":
        v19 = (v19 ^ -179)
    default:
        var v22 float32 = 0.1
        _ = v22
    }
    v2 = (v2 / -2.25)
    v23 := GenS3{F0: true, F1: 514, F2: 65535}
    if (-6 == 100) {
        var v24 int32 = -409
        for fi25 := 0; fi25 < 11; fi25++ {
            v2 -= (-2.25 - (0.1 - 0.0))
            v26 := GenS4{F0: 68, F1: 2, F2: true}
            _ = fi25
            _ = v26
        }
        v27 := GenS5{F0: -1, F1: 1}
        _ = v24
        _ = v27
    }
    if false {
        if false {
            fmt.Println(v1, v2)
        }
        if false {
            if false {
                switch "a" {
                case "":
                    var v28 int8 = -1
                    _ = v28
                case "foo bar":
                    _ = len(v4)
                case "a":
                    fmt.Println(v2, v1)
                default:
                    var v29 bool = false
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                v1 = (v1 + v1)
                v3[fi30] = 2147483648
                v19 -= (v19 + (v19 / ((v19) | 1)))
                _ = fi30
            }
            fmt.Println(v19, v19, v2)
        } else {
            fmt.Println(v19)
            v2 = 1e10
        }
        fmt.Println(v19, v1, v19)
    } else {
        v1 = v1
    }
    v31 := GenS5{F0: 0, F1: 53}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
