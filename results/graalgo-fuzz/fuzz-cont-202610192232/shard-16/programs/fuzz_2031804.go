package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{0, 413, 65535}
    v3 := make(map[string]bool)
    var v4 uint64 = 503
    v1 = ("a" + v1)
    fmt.Println(v4, v4)
    var v5 int8 = 0
    v6 := GenS1{F0: v4, F1: ((v1 + (v1 + v1)) + "a"), F2: 939}
    var v7 float32 = 1.5
    switch v1 {
    case "\n\t":
        var v8 float32 = v7
        _ = v8
    case "a":
        if false {
            v4 = v4
        } else {
            v9 := GenS1{F0: v4, F1: v1, F2: 1}
            v5 -= v5
            for fi10 := 0; fi10 < 18; fi10++ {
                switch -948 {
                case 2147483647:
                    v11 := GenS2{F0: fi10, F1: 153}
                    _ = v11
                case 119:
                    v12 := GenS1{F0: v4, F1: (v1 + v1), F2: 42}
                    _ = v12
                case 0:
                    if len(v2) > 0 {
                        v2[0] = uint32(v4)
                    }
                }
                v13 := GenS3{F0: ((((3.14159 / v7) + 0.0) / (1e10 * 1.5)) <= (1.5 + v7))}
                v4 = (^(v4))
                _ = fi10
                _ = v13
            }
            _ = v9
        }
    default:
        if (false || false) {
            v4 = v4
        } else {
            var v14 int64 = 42
            switch v1 {
            case "":
                if (!true) {
                    var v15 int32 = 903
                    v4 -= v4
                    _ = v15
                } else {
                    switch 4294967295 {
                    case 32768:
                        if (false || false) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v16 int16 = -509
                            for fi17 := 0; fi17 < 15; fi17++ {
                                v18 := v3[(((v1 + (("hello" + v1) + (((("a" + v1) + v1) + v1) + v1))) + ("hello" + "hello")) + v1)]
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        }
                    case 9223372036854775807:
                        v19 := GenS4{F0: 255, F1: 32}
                        _ = v19
                    default:
                        v7 += -2.25
                    }
                    for fi20 := 0; fi20 < 14; fi20++ {
                        if ("" > v1) {
                            v21 := GenS4{F0: 73, F1: 1}
                            v7 = float32(float64(v7))
                            _ = v21
                        }
                        var v22 float64 = ((0.0 + 0.1) / ((0.1 + 1e10) / (0.0 - -2.25)))
                        v14 = -540
                        _ = fi20
                        _ = v22
                    }
                }
            case "a":
                var v23 float64 = 3.14159
                _ = v23
            case "\n\t":
                v7 = float32(1e10)
            default:
                v2 = append(v2, 42)
            }
            _ = v14
        }
    }
    v1 = v1
    var v24 string = ""
    var v25 int32 = 2147483647
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            if false {
                var v28 uint = 674
                if true {
                    delete(v3, "a")
                } else {
                    var v29 uint16 = 0
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    v29 = (v29 >> ((18446744073709551615) & 63))
                    _ = v29
                }
                _ = v28
            } else {
                var v30 float64 = 0.0
                _ = v30
            }
            var v31 int8 = (^(v5))
            _ = fi27
            _ = v31
        }
        v7 -= v7
        _ = fi26
    }
    if (-2.25 != -2.25) {
        var v32 int32 = (v25 * v25)
        if true {
            if (!true) {
                var v33 string = (((("hello" + v24) + v24) + (v24 + "foo bar")) + ((v1 + v24) + v24))
                v2 = append(v2, 1)
                for fi34 := 0; fi34 < 14; fi34++ {
                    if true {
                        v35 := GenS2{F0: fi34, F1: 42}
                        _ = v35
                    }
                    _ = fi34
                }
                _ = v33
            }
            delete(v3, v1)
        }
        var v36 float32 = 0.0
        _ = v32
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
