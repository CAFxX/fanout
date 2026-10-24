package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int8; F1 int32; F2 string }
type GenS5 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0(p0 int64, p1 string, p2 int32) int {
    fmt.Println(p2, p0)
    var v4 int = 1
    switch p1 {
    case "foo bar":
        var v5 float32 = 0.1
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := []int{v4, (fi6 % (((fi7 ^ v4)) | 1))}
                if (true && ((p0 % ((p0) | 1)) >= (p0 >> ((2) & 63)))) {
                    if (42 > 0) {
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        switch ("a" + p1) {
                        case "a":
                            var v9 int8 = -128
                            _ = v9
                        default:
                            switch fi7 {
                            case -591:
                                fmt.Println(p1, p1)
                            default:
                                v4 = (fi6 | (v4 % ((-268) | 1)))
                            }
                        }
                        switch fi7 {
                        case -1:
                            p1 = p1
                        case 9223372036854775807:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                    }
                } else {
                    var v10 int32 = p2
                    if true {
                        var v11 bool = (255 > 2)
                        var v12 bool = v11
                        fmt.Println(p2, v4)
                        _ = v11
                        _ = v12
                    }
                    _ = v10
                }
                v13 := GenS1{F0: 1, F1: 2, F2: 83}
                switch "\n\t" {
                case "":
                    v14 := GenS1{F0: 90, F1: 255, F2: 42}
                    _ = v14
                default:
                    fmt.Println(v4, fi7)
                }
                var v15 int64 = p0
                p0 = p0
                _ = fi7
                _ = v13
                _ = v15
                _ = v8
            }
            fmt.Println(p0)
            fmt.Println(p1, fi6, p2)
            _ = fi6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 | (-(v4)))
}
func genFn1() uint8 {
    var v16 int32 = 1
    var v17 int16 = -1
    _ = v16
    _ = v17
    return 163
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 1
    var v3 int32 = v2
    v3 = v2
    v2 = -8
    if true {
        var v18 float64 = float64(0.1)
        if true {
            v19 := []int{}
            v20 := GenS2{F0: 878}
            var v21 uint8 = 128
            _ = v19
            _ = v20
            _ = v21
        }
        switch 42 {
        case 128:
            v1[790] = "foo bar"
        case 695:
            v1[127] = (("\n\t" + "hello") + "foo bar")
        case -1:
            var v22 uint32 = 4294967295
            _ = v22
        default:
            if true {
                if false {
                    v2 = v2
                    v3 += v3
                } else {
                    v18 = (0.1 - 1e10)
                }
                var v23 int16 = 700
                v24 := []int{128, 61, -838}
                _ = v23
                _ = v24
            }
        }
        _ = v18
    }
    var v25 uint32 = 2147483647
    v26 := GenS3{F0: -400, F1: 572}
    for fi27 := 0; fi27 < 19; fi27++ {
        fmt.Println(v25)
        var v28 uint = 65535
        v29 := GenS4{F0: -120, F1: -2147483648, F2: (((("\n\t" + ("a" + ("hello" + (("hello" + "foo bar") + ("\n\t" + ""))))) + ("a" + "a")) + ("foo bar" + "foo bar")) + "hello")}
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v3, v3)
    v30 := GenS5{F0: 32767, F1: false, F2: 130}
    fmt.Println(v2)
    switch "foo bar" {
    case "":
        var v31 int32 = v3
        _ = v31
    case "foo bar":
        v3 -= v3
    case "a":
        var v32 int8 = 1
        _ = v32
    }
    var v33 int32 = (^(v2))
    var v34 uint64 = 2
    v3 -= v3
    fmt.Println(v3, v33, v25)
    switch 0 {
    case 42:
        v25 += 4294967295
    case 380:
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 int = (fi35 * 127)
            v37 := GenS1{F0: 42, F1: 672, F2: 54}
            v3 -= (v3 % (((v2 % ((v33) | 1))) | 1))
            _ = fi35
            _ = v36
            _ = v37
        }
    case 4294967295:
        var v38 int64 = 9223372036854775807
        _ = v38
    }
    v2 += (v3 - v33)
    for fi39 := 0; fi39 < 7; fi39++ {
        v1[fi39] = "a"
        var v40 float32 = 3.14159
        _ = fi39
        _ = v40
    }
    delete(v1, 9223372036854775807)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v33
    _ = v34
}
