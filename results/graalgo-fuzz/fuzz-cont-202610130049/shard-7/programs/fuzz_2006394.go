package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint8) uint8 {
    p1 = (p1 % ((p1) | 1))
    p1 -= p1
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 int64 = -1
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 uint16, p1 int32, p2 int16) int32 {
    if (!false) {
        if (!(true && true)) {
            v7 := GenS1{F0: 884}
            _ = v7
        }
        if true {
            var v8 int = -327
            switch ("foo bar" + "hello") {
            case "hello":
                p0 = (p0 + (p0 * p0))
            case "foo bar":
                v9 := GenS1{F0: 65535}
                _ = v9
            case "\n\t":
                fmt.Println(p2, p1)
            default:
                v10 := GenS1{F0: 2}
                _ = v10
            }
            var v11 uint16 = 286
            _ = v11
            _ = v8
        }
    }
    v12 := GenS2{F0: true, F1: p0}
    switch ("foo bar" + (((("foo bar" + "hello") + ("foo bar" + "hello")) + "hello") + "")) {
    case "a":
        var v13 float64 = 1e10
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -125
}

func main() {
    var v1 int8 = 107
    var v2 int64 = 1
    v3 := []uint32{65535, 2147483648, 1}
    var v4 uint = 958
    v14 := GenS1{F0: 0}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint64 = 42
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: ((!((-97 >= -8) && (v1 == 100))) || false), F1: 2}
    if (!true) {
        var v18 uint16 = 255
        _ = v18
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 string = ""
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 uint16 = 65535
            switch 128 {
            case 42:
                switch fi19 {
                case 32767:
                    if false {
                        var v23 uint32 = 2147483648
                        _ = len(v3)
                        fmt.Println(v4, fi21)
                        _ = v23
                    } else {
                        v24 := GenS3{F0: v4}
                        v25 := GenS2{F0: false, F1: uint16(v4)}
                        _ = v24
                        _ = v25
                    }
                case -9223372036854775808:
                    v2 += v2
                case 9223372036854775807:
                    var v26 int8 = v1
                    _ = v26
                default:
                    v2 = (v2 | (v2 % ((v2) | 1)))
                }
            default:
                _ = len(v3)
            }
            fmt.Println(fi19, v22, v20)
            _ = fi21
            _ = v22
        }
        fmt.Println(v4, fi19)
        _ = fi19
        _ = v20
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        if (32767 <= 42) {
            switch fi27 {
            case 1:
                if false {
                    fmt.Println(fi27)
                } else {
                    if false {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v29 := GenS4{F0: (v4 ^ (v4 ^ v4)), F1: 2, F2: 123}
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        fmt.Println(fi27)
                    }
                    var v30 uint8 = 0
                    v30 -= (v30 - v30)
                    _ = v30
                }
            }
            v2 += 115
            v31 := GenS5{F0: (v2 % (((-(v2))) | 1))}
            _ = v31
        }
        _ = fi27
    }
    var v32 int = 1
    var v33 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
