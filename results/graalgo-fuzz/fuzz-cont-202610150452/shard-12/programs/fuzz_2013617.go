package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 int64; F2 bool }
func genFn0(p0 int) int32 {
    v5 := GenS1{F0: 1, F1: p0}
    var v6 uint16 = 2
    var v7 float32 = (1.5 * 0.1)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{("a" + ("" + ("a" + "\n\t"))), "hello", "hello"}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v8 uint = 289
    fmt.Println(v8)
    v9 := GenS1{F0: 308, F1: 42}
    for fi10 := 0; fi10 < 15; fi10++ {
        if len(v1) > 0 {
            v1[0] = ("a" + "hello")
        }
        fmt.Println(v8, v8)
        _ = fi10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v8, fi11)
        if (fi11 == (fi11 * fi11)) {
            var v12 uint64 = 42
            _ = v12
        } else {
            switch "\n\t" {
            case "\n\t":
                var v13 uint = (v8 - (v8 - v8))
                _ = v13
            case "":
                v14 := GenS2{F0: 42}
                _ = v14
            case "foo bar":
                v15 := GenS3{F0: ("foo bar" + ("hello" + ("foo bar" + "\n\t"))), F1: 1}
                _ = v15
            default:
                for fi16 := 0; fi16 < 6; fi16++ {
                    v8 = v8
                    v8 -= (v8 & (v8 | (v8 * v8)))
                    v1 = append(v1, "\n\t")
                    var v17 uint16 = 0
                    if (9223372036854775807 == -1) {
                        var v18 int8 = 1
                        v18 -= 100
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
        if (true || ((!false) || (0 < 255))) {
            switch ("foo bar" + ("\n\t" + "\n\t")) {
            case "a":
                fmt.Println(fi11, fi11)
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    var v19 int8 = 1
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if (false && (!true)) {
                            delete(v2, 188)
                            if true {
                                if (false || (!(!true))) {
                                    v21 := GenS4{F0: 0, F1: 4294967295, F2: uint64(v8)}
                                    _ = v21
                                } else {
                                    v8 -= v8
                                }
                                var v22 float32 = (-2.25 - 0.0)
                                _ = v22
                            } else {
                                var v23 uint = 255
                                _ = v23
                            }
                            v8 = (^(v8))
                        } else {
                            v8 -= v8
                        }
                        v24 := GenS4{F0: 42, F1: 223, F2: 0}
                        v8 += v8
                        _ = fi20
                        _ = v24
                    }
                }
            default:
                var v25 float64 = float64(3.14159)
                _ = v25
            }
            v26 := GenS5{F0: 772}
            _ = v26
        }
        _ = fi11
    }
    v8 = (^(v8))
    fmt.Println(v8)
    if (!true) {
        v27 := GenS1{F0: -1, F1: 9223372036854775807}
        var v28 uint32 = 42
        _ = v27
        _ = v28
    }
    var v29 uint8 = 2
    v29 -= (v29 / ((v29) | 1))
    if (-94 < -1) {
        if false {
            v30 := v4[2147483647]
            _ = v30
        } else {
            fmt.Println(v8)
        }
        v29 -= v29
        v31 := GenS2{F0: 1}
        _ = v31
    }
    var v32 uint64 = 349
    v33 := GenS6{F0: 9223372036854775807, F1: -1, F2: (-32768 > -272)}
    for fi34 := 0; fi34 < 16; fi34++ {
        var v35 int16 = -794
        var v36 string = ("" + ("hello" + ((("a" + "foo bar") + ((("hello" + "\n\t") + "") + (("\n\t" + ("" + "\n\t")) + ("hello" + "a")))) + "hello")))
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v8
    _ = v9
}
