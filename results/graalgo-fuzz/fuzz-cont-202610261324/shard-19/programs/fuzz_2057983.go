package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 string; F1 int64 }
func genFn0(p0 uint32) string {
    switch 264 {
    case 42:
        var v7 uint64 = 14
        _ = v7
    case 32768:
        var v8 uint = 506
        _ = v8
    case 588:
        v9 := []int{}
        _ = v9
    default:
        var v10 int16 = -263
        _ = v10
    }
    fmt.Println(p0)
    var v11 string = ""
    _ = p0
    _ = v11
    return ("hello" + "hello")
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = 2
    v3 := []int{128, 171}
    v4 := make(map[string]bool)
    var v5 uint16 = 1
    var v6 int8 = 42
    v12 := GenS1{F0: -9223372036854775808, F1: 42, F2: 2147483647}
    if false {
        v13 := GenS2{F0: v6, F1: 841}
        _ = v13
    } else {
        var v14 bool = true
        fmt.Println(v5)
        _ = v14
    }
    var v15 bool = true
    for fi16 := 0; fi16 < 20; fi16++ {
        v3 = append(v3, 4294967295)
        v3 = append(v3, (fi16 + (fi16 | -1)))
        if ((42 < -177) && (v15 || ((v6 * (v6 >> ((v2) & 63))) > int8(fi16)))) {
            var v17 bool = v15
            if (!(!(0.0 <= 0.0))) {
                if (132 <= 59) {
                    v6 += (v6 % (((v6 * v6)) | 1))
                    v18 := GenS3{F0: 4294967295, F1: v6}
                    _ = v18
                }
                switch "" {
                case "":
                    v19 := GenS3{F0: 1, F1: v6}
                    _ = v19
                case "foo bar":
                    if v15 {
                        v2 += v2
                    }
                case "hello":
                    if (false || v15) {
                        v2 = (v2 - (v2 + v2))
                        fmt.Println(v6, v1)
                        var v20 int8 = v6
                        _ = v20
                    }
                default:
                    fmt.Println(v6)
                }
            } else {
                fmt.Println(v6, v5, v15)
            }
            switch fi16 {
            case -1:
                v2 -= (v2 | 226)
            }
            _ = v17
        }
        _ = fi16
    }
    if (4294967295 != 95) {
        if v15 {
            v4[("" + "hello")] = (!(9223372036854775807 > 9223372036854775807))
        } else {
            v21 := GenS4{F0: (!true), F1: (("" + ("\n\t" + "\n\t")) + ("" + ("foo bar" + (((("foo bar" + "\n\t") + ("\n\t" + "hello")) + ("hello" + ("foo bar" + "\n\t"))) + ("" + ("foo bar" + "\n\t"))))))}
            var v22 uint32 = v1
            var v23 string = "hello"
            _ = v21
            _ = v22
            _ = v23
        }
        v24 := GenS5{F0: "", F1: -1}
        var v25 int64 = int64(v5)
        _ = v24
        _ = v25
    } else {
        var v26 uint32 = (v1 | (v1 ^ v1))
        _ = v26
    }
    var v27 int = 518
    if true {
        fmt.Println(v15, v6, v5)
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 float64 = (3.14159 / 0.1)
        var v30 float32 = (0.0 + 3.14159)
        v4["a"] = (0 == fi28)
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
