package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint = 1
    v2 := []string{"hello", (("hello" + (("" + "hello") + ("a" + ""))) + "\n\t"), ("" + ("\n\t" + ""))}
    var v3 int8 = 42
    switch 127 {
    case 302:
        v3 = (^(v3))
    case 128:
        var v4 int32 = -751
        _ = v4
    }
    var v5 uint8 = 93
    var v6 float64 = (0.0 * 3.14159)
    var v7 bool = (false && (false && (v1 >= v1)))
    v3 = v3
    v5 += (v5 % ((199) | 1))
    var v8 int32 = 0
    fmt.Println(v7, v6, v3)
    _ = len(v2)
    v5 = v5
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int32 = (v8 / ((v8) | 1))
        switch "foo bar" {
        case "\n\t":
            v11 := GenS1{F0: 433}
            _ = v11
        case "a":
            if (!true) {
                _ = len(v2)
            } else {
                switch int(v5) {
                case -139:
                    v12 := GenS2{F0: (("\n\t" + ("\n\t" + ("hello" + ("\n\t" + "hello")))) + ("a" + ""))}
                    _ = v12
                default:
                    v5 -= uint8(v1)
                }
                fmt.Println(fi9)
            }
        case "":
            var v13 int16 = 42
            _ = v13
        default:
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
        }
        v7 = ((^(v1)) > v1)
        _ = fi9
        _ = v10
    }
    switch 128 {
    case -1:
        var v14 float32 = (3.14159 / -2.25)
        _ = v14
    case 32768:
        if (518 > 2) {
            var v15 bool = (957 <= 0)
            _ = v15
        } else {
            switch "a" {
            case "hello":
                if (v7 || (float64(0.1) == float64(1e10))) {
                    switch 32767 {
                    case 4294967295:
                        v1 = v1
                    case -2147483648:
                        v1 += v1
                    case 749:
                        v16 := GenS3{F0: v8, F1: "a", F2: 32767}
                        _ = v16
                    }
                }
            case "a":
                for fi17 := 0; fi17 < 13; fi17++ {
                    if (((v5 | (v5 & (v5 & 149))) != 2) && false) {
                        v8 = -2147483648
                        v3 -= (v3 - (-(v3)))
                    } else {
                        v18 := GenS2{F0: "a"}
                        _ = v18
                    }
                    if false {
                        if len(v2) > 0 {
                            v2[0] = ((("hello" + "foo bar") + "foo bar") + "foo bar")
                        }
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v20 := GenS4{F0: 2, F1: 0}
                            v6 = (v6 + float64(((0.1 + 0.0) / float32(v6))))
                            _ = fi19
                            _ = v20
                        }
                    }
                    _ = fi17
                }
            }
            if (v7 && (((true || v7) && ("\n\t" >= "a")) && (2147483647 != 42))) {
                fmt.Println(v6, v8, v7)
                v6 = 3.14159
            } else {
                for fi21 := 0; fi21 < 7; fi21++ {
                    v22 := GenS5{F0: 273}
                    v1 += (v1 & (v1 * (^(v1))))
                    _ = fi21
                    _ = v22
                }
            }
        }
    case -712:
        v23 := GenS4{F0: v5, F1: 42}
        _ = v23
    }
    var v24 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
