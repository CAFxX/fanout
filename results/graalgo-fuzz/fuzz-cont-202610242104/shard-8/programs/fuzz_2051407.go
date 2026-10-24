package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 string; F1 string; F2 int }
type GenS3 struct { F0 int16; F1 int32 }
func genFn0(p0 int, p1 int64) uint8 {
    v7 := []int{}
    if false {
        if true {
            var v8 int = p0
            _ = v8
        }
        if true {
            if true {
                v9 := GenS1{F0: 0, F1: 1000, F2: 1}
                _ = v9
            } else {
                if (!true) {
                    var v10 uint32 = 1
                    var v11 bool = true
                    p0 = 32768
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: -155, F1: 0, F2: 375}
                _ = v12
            }
            switch ("foo bar" + "") {
            case "\n\t":
                if false {
                    var v13 float32 = ((0.1 / 0.1) + 0.0)
                    _ = v13
                }
            case "":
                var v14 int8 = 15
                _ = v14
            default:
                var v15 uint64 = 387
                _ = v15
            }
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p0 = p0
    var v16 int16 = 0
    _ = p0
    _ = p1
    _ = v16
    _ = v7
    return 2
}
func genFn1(p0 uint8) string {
    fmt.Println(p0)
    p0 = p0
    p0 = (p0 ^ (^(p0)))
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 float32 = 3.14159
        v18 -= (1.5 + v18)
        _ = fi17
        _ = v18
    }
    _ = p0
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (3.14159 / 1e10)
    v3 := make(map[string]bool)
    v4 := []uint32{1, 2147483647, 761, 2147483647}
    v5 := make(map[uint32]bool)
    var v6 bool = false
    var v19 uint32 = 739
    fmt.Println(v19, v19, v2)
    v1[(v19 * 1)] = (("\n\t" + "\n\t") + ("a" + "\n\t"))
    v19 += 0
    v20 := GenS2{F0: ("foo bar" + "hello"), F1: "foo bar", F2: 288}
    switch -2147483648 {
    case 42:
        v19 = v19
    case 9223372036854775807:
        if (((0 <= 2147483647) && true) || false) {
            switch "\n\t" {
            case "a":
                switch ("" + ("hello" + "a")) {
                case "hello":
                    v21 := GenS3{F0: 32767, F1: 42}
                    _ = v21
                case "\n\t":
                    var v22 int8 = -30
                    _ = v22
                case "":
                    var v23 float64 = ((float64(float32(0.0)) * (v2 + (v2 - 0.1))) - v2)
                    _ = v23
                }
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "":
                fmt.Println(v6)
            default:
                var v24 int8 = 1
                _ = v24
            }
            var v25 int64 = 0
            fmt.Println(v6, v25)
            _ = v25
        }
    case -120:
        v26 := v5[1]
        _ = v26
    }
    _ = len(v4)
    if v6 {
        var v27 uint64 = 417
        var v28 string = ("a" + ((("hello" + "foo bar") + ("hello" + "\n\t")) + "a"))
        _ = v27
        _ = v28
    }
    delete(v1, v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
