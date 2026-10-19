package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0(p0 int16, p1 uint32) int {
    var v5 int8 = 120
    if false {
        if ((((3.14159 + 0.0) + 0.1) - 0.0) > 3.14159) {
            var v6 uint8 = 42
            var v7 string = "a"
            _ = v6
            _ = v7
        }
        v8 := GenS1{F0: 9223372036854775807, F1: (960 >= 1), F2: (p1 & (p1 * 42))}
        _ = v8
    } else {
        fmt.Println(v5, p0, p0)
        switch int(p1) {
        case 32768:
            fmt.Println(v5, v5, p1)
        case 128:
            p1 = (^(p1))
        }
        if (0.1 < 0.1) {
            if true {
                var v9 int8 = (v5 - v5)
                v10 := []int{32767, -1000, 4294967295, -1}
                v5 = -1
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{-60, -1, 213, 22}
            p0 -= p0
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 int64, p1 int8) uint32 {
    p0 -= p0
    v12 := GenS1{F0: (^(p0)), F1: (!false), F2: 625}
    p1 = (p1 >> ((237) & 63))
    var v13 int64 = p0
    if true {
        fmt.Println(p1)
        p0 += -833
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 1
    v3 := []string{("foo bar" + "hello")}
    var v4 float64 = 0.0
    switch -9223372036854775808 {
    case -1:
        _ = len(v3)
    case 128:
        if true {
            v14 := GenS1{F0: (^(v1)), F1: (!(("foo bar" + "\n\t") >= "hello")), F2: 0}
            if (v4 > v4) {
                var v15 int64 = (v1 / ((v1) | 1))
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 int8 = 100
                    v15 -= v1
                    _ = fi16
                    _ = v17
                }
                fmt.Println(v1, v4, v1)
                _ = v15
            } else {
                var v18 float64 = v4
                switch ("hello" + ("foo bar" + "hello")) {
                case "":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 int64 = (v1 * (v1 >> ((255) & 63)))
                        var v21 uint = 255
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case "hello":
                    var v22 uint8 = uint8(v2)
                    _ = v22
                case "a":
                    _ = len(v3)
                default:
                    fmt.Println(v18)
                }
                _ = v18
            }
            _ = v14
        }
    }
    fmt.Println(v4, v2)
    fmt.Println(v1)
    v4 -= v4
    v2 = 32767
    v23 := GenS2{F0: (v1 << ((2) & 63)), F1: -364}
    v24 := GenS3{F0: -32768, F1: 180, F2: 2147483647}
    v1 = (^(v1))
    var v25 float32 = 1e10
    v25 = float32((v4 - ((((float64((((3.14159 + 1.5) / 1.0) + 3.14159)) + float64(1e10)) - v4) + float64(0.1)) / (1e10 * v4))))
    v4 = (v4 + -2.25)
    var v26 float64 = ((((v4 / v4) / float64(1e10)) - v4) / (v4 / (float64(v25) * v4)))
    switch ((("foo bar" + ("a" + "\n\t")) + "hello") + "a") {
    case "":
        v26 = ((v26 / v4) / (-2.25 / v4))
    case "foo bar":
        v25 = 1.5
    case "a":
        fmt.Println(v25, v2)
    default:
        _ = len(v3)
    }
    var v27 int16 = 1
    v27 -= 0
    var v28 uint8 = 42
    if true {
        var v29 float32 = ((3.14159 - 1.5) / v25)
        fmt.Println(v25, v26, v27)
        if (v27 < -1) {
            v1 -= v1
        } else {
            v3 = append(v3, (("a" + ("\n\t" + ("\n\t" + "\n\t"))) + "foo bar"))
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
