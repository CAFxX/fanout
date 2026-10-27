package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 uint64 }
type GenS3 struct { F0 bool; F1 string; F2 int8 }
type GenS4 struct { F0 uint8; F1 int16; F2 int }

func main() {
    var v1 float32 = 1e10
    v2 := []uint32{65535}
    var v3 bool = true
    switch "hello" {
    case "a":
        v3 = v3
    case "\n\t":
        v1 += v1
    }
    v2 = append(v2, 1)
    switch 42 {
    case -1:
        for fi4 := 0; fi4 < 13; fi4++ {
            if v3 {
                v5 := GenS1{F0: fi4, F1: 9223372036854775807}
                for fi6 := 0; fi6 < 15; fi6++ {
                    v7 := GenS1{F0: (fi4 >> ((631) & 63)), F1: -335}
                    _ = fi6
                    _ = v7
                }
                _ = v5
            } else {
                fmt.Println(v3, v3, v1)
                if v3 {
                    v1 += v1
                    switch "\n\t" {
                    case "a":
                        if (42 <= -48) {
                            v8 := GenS2{F0: -719, F1: (!v3), F2: 886}
                            _ = v8
                        }
                    case "hello":
                        for fi9 := 0; fi9 < 4; fi9++ {
                            _ = len(v2)
                            var v10 string = ("" + (("a" + ("hello" + "hello")) + ((("a" + "") + "hello") + "hello")))
                            _ = fi9
                            _ = v10
                        }
                    case "\n\t":
                        v3 = true
                    default:
                        v3 = ((3.14159 * 0.0) < 0.0)
                    }
                    for fi11 := 0; fi11 < 10; fi11++ {
                        fmt.Println(v1)
                        var v12 float64 = 1e10
                        _ = fi11
                        _ = v12
                    }
                } else {
                    _ = len(v2)
                    fmt.Println(v1, v1, v1)
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v1 += v1
                        var v14 int16 = 365
                        var v15 uint = 1
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
            }
            fmt.Println(fi4, fi4, v3)
            var v16 float32 = 0.0
            _ = fi4
            _ = v16
        }
    case 128:
        v3 = (("" + "hello") <= ("\n\t" + ("" + (("foo bar" + (("hello" + "\n\t") + "foo bar")) + ("a" + ("\n\t" + "\n\t"))))))
    }
    v17 := GenS3{F0: (v3 || true), F1: "foo bar", F2: 0}
    _ = len(v2)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int64 = 240
        if true {
            fmt.Println(fi18)
        } else {
            if v3 {
                switch "" {
                case "hello":
                    fmt.Println(fi18, v19, v3)
                case "foo bar":
                    for fi20 := 0; fi20 < 18; fi20++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi20
                    }
                default:
                    var v21 uint = 177
                    _ = v21
                }
                v22 := GenS2{F0: 2147483647, F1: (!false), F2: 4294967295}
                _ = v22
            } else {
                if (38 < 112) {
                    v23 := GenS4{F0: 87, F1: 32767, F2: (fi18 & fi18)}
                    v19 = (^(v19))
                    _ = v23
                }
            }
        }
        _ = fi18
        _ = v19
    }
    var v24 int8 = 100
    if (-1 >= -32768) {
        v2 = append(v2, 65535)
        var v25 float64 = float64(((((0.1 + float32(3.14159)) + (float32((0.1 / 1e10)) * (((-2.25 + (v1 - v1)) * (v1 * -2.25)) * float32(float64(((v1 * float32(float64(v1))) + ((v1 * v1) - v1))))))) + v1) / 1.0))
        v26 := GenS4{F0: 255, F1: -1, F2: -649}
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 11; fi27++ {
        fmt.Println(v24, v1)
        if false {
            var v28 int = fi27
            v3 = (0.0 <= (3.14159 / float64(0.1)))
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
}
