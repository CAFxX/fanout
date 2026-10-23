package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32; F1 string; F2 uint }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0(p0 int, p1 int32) string {
    p0 = p0
    if true {
        fmt.Println(p1, p1)
    }
    p1 = (p1 & (-(p1)))
    p1 -= (p1 & (p1 >> ((658) & 63)))
    v5 := GenS1{F0: (986 >= 158), F1: "a"}
    _ = p0
    _ = p1
    _ = v5
    return "hello"
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 2
    var v3 int16 = 738
    v4 := make(map[string]string)
    if (0 < 754) {
        v6 := GenS2{F0: false, F1: 0, F2: -225}
        if (!(!false)) {
            var v7 bool = (!true)
            fmt.Println(v2, v2, v3)
            switch "\n\t" {
            case "foo bar":
                if v7 {
                    v2 = v2
                    if v7 {
                        var v8 int8 = 53
                        _ = v8
                    }
                } else {
                    fmt.Println(v7)
                    v9 := GenS3{F0: -236, F1: "\n\t", F2: 2}
                    _ = v9
                }
            }
            _ = v7
        } else {
            if false {
                v10 := GenS4{F0: 127, F1: int16(v2)}
                fmt.Println(v3, v2)
                _ = v10
            } else {
                v11 := GenS4{F0: 127, F1: -32768}
                _ = v11
            }
            v2 += (v2 * (v2 / ((v2) | 1)))
        }
        _ = v6
    }
    v12 := []int{}
    fmt.Println(v3, v2)
    v3 = 32767
    if (("foo bar" >= "foo bar") || (false || ((("\n\t" + "\n\t") + "foo bar") <= (((((("a" + ("a" + ("" + "foo bar"))) + "\n\t") + ("a" + "\n\t")) + ("\n\t" + "a")) + "a") + "hello")))) {
        v1[-1] = "a"
    } else {
        if (-2147483648 > -814) {
            v3 = v3
        }
    }
    v2 -= (v2 ^ v2)
    v3 = 42
    var v13 float64 = 3.14159
    var v14 bool = (false || false)
    var v15 string = "\n\t"
    switch 595 {
    case 9223372036854775807:
        for fi16 := 0; fi16 < 20; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
    case -431:
        if ((-1 <= 4294967295) && false) {
            switch (v15 + v15) {
            case "hello":
                v17 := GenS1{F0: v14, F1: v15}
                _ = v17
            case "":
                switch v15 {
                case "a":
                    if v14 {
                        v2 = 1
                        v18 := GenS2{F0: false, F1: 2, F2: 9223372036854775807}
                        v19 := GenS5{F0: -37, F1: 1}
                        _ = v18
                        _ = v19
                    } else {
                        var v20 float64 = ((v13 - float64(1.5)) - 1.5)
                        _ = v20
                    }
                default:
                    if v14 {
                        v2 -= 18446744073709551615
                    }
                }
            case "a":
                var v21 uint32 = 42
                _ = v21
            }
            fmt.Println(v14, v2)
            var v22 int32 = 175
            _ = v22
        } else {
            fmt.Println(v14)
        }
    default:
        v15 = v15
    }
    v15 = "a"
    var v23 uint64 = v2
    if false {
        v12 = append(v12, -9223372036854775808)
        if len(v12) > 0 {
            v12[0] = 211
        }
        if false {
            switch -2147483648 {
            case 1:
                var v24 int64 = 9223372036854775807
                _ = v24
            }
        }
    } else {
        var v25 int8 = -128
        fmt.Println(v15)
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
