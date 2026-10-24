package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 bool; F2 bool }
func genFn0(p0 int8, p1 int64) uint16 {
    v4 := GenS1{F0: 2147483647, F1: p0, F2: -611}
    v5 := GenS2{F0: 0, F1: 4294967295}
    p1 = p1
    for fi6 := 0; fi6 < 12; fi6++ {
        p1 += (p1 + p1)
        p0 = p0
        p1 += p1
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int64, p1 int32) uint32 {
    if false {
        v7 := []int{}
        if (((-2.25 - float64(1e10)) * (1.5 + 1.5)) > -2.25) {
            fmt.Println(p1)
        } else {
            p1 -= p1
            var v8 int = -2147483648
            _ = v8
        }
        _ = v7
    }
    p1 += (-(p1))
    _ = p0
    _ = p1
    return 593
}

func main() {
    var v1 int32 = -1
    var v2 int16 = 1
    var v3 int8 = 100
    if true {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v3, v1)
            _ = fi9
        }
        v10 := []int{32767, 0}
        if true {
            if (!false) {
                var v11 float32 = (-2.25 * 1.5)
                _ = v11
            }
            if (true && false) {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v3)
                }
            } else {
                _ = len(v10)
                v12 := GenS1{F0: 32767, F1: 127, F2: 1}
                v2 += v2
                _ = v12
            }
            var v13 uint64 = 2
            _ = v13
        }
        _ = v10
    }
    if (true && false) {
        v1 += v1
    }
    switch 381 {
    case -2147483648:
        v3 += v3
    case 681:
        if (true && ("foo bar" == "hello")) {
            if (("\n\t" + "hello") >= "foo bar") {
                var v14 uint32 = 2147483647
                _ = v14
            } else {
                var v15 uint = 4294967295
                v16 := []int{}
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(fi17, v2)
                v1 = (v1 + (v1 * 42))
                _ = fi17
            }
        } else {
            fmt.Println(v2, v2)
            switch 580 {
            case 2147483647:
                var v18 uint16 = 255
                _ = v18
            case 9223372036854775807:
                if true {
                    v2 = v2
                }
            default:
                var v19 int64 = 42
                _ = v19
            }
        }
    default:
        v2 = 374
    }
    v20 := []int{371, -1, 128}
    var v21 int = 127
    fmt.Println(v1, v1)
    if (!(0.0 != 3.14159)) {
        var v22 float32 = 1.5
        v22 -= v22
        _ = v22
    } else {
        var v23 bool = false
        if true {
            if (v21 < (v21 / (((-(v21))) | 1))) {
                fmt.Println(v23)
                if ((v2 % ((v2) | 1)) == (-(v2))) {
                    if false {
                        if (!(1e10 < 0.1)) {
                            fmt.Println(v1)
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            var v24 uint16 = 2
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 15; fi25++ {
                                var v26 bool = ((v3 == (v3 | v3)) && true)
                                fmt.Println(v2, fi25, v23)
                                _ = fi25
                                _ = v26
                            }
                            var v27 float32 = 1e10
                            _ = v27
                        }
                        if len(v20) > 0 {
                            v20[0] = 128
                        }
                    } else {
                        v28 := GenS3{F0: 4294967295, F1: 496}
                        v29 := GenS4{F0: 255, F1: ("a" + ((("\n\t" + "hello") + "") + ("\n\t" + "\n\t")))}
                        for fi30 := 0; fi30 < 1; fi30++ {
                            fmt.Println(fi30, fi30)
                            _ = fi30
                        }
                        _ = v28
                        _ = v29
                    }
                    v23 = (((v3 & (^(v3))) < v3) || false)
                }
            } else {
                var v31 uint16 = 2
                _ = v31
            }
            v32 := GenS5{F0: 1}
            _ = v32
        } else {
            v33 := GenS4{F0: 43, F1: (("" + (((("hello" + "\n\t") + "foo bar") + (("" + ("" + ("" + "hello"))) + "a")) + "hello")) + "")}
            _ = v33
        }
        var v34 uint32 = 799
        _ = v23
        _ = v34
    }
    _ = len(v20)
    v1 = 1
    switch v21 {
    case 2147483647:
        var v35 bool = (65535 <= 518)
        _ = v35
    case 61:
        fmt.Println(v21, v21)
    case -2147483648:
        for fi36 := 0; fi36 < 19; fi36++ {
            v37 := GenS1{F0: v21, F1: -112, F2: (v1 % ((499) | 1))}
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            var v38 int = v21
            _ = fi36
            _ = v37
            _ = v38
        }
    default:
        v1 += v1
    }
    _ = len(v20)
    if false {
        var v39 uint16 = 65535
        var v40 int32 = v1
        _ = v39
        _ = v40
    } else {
        v41 := GenS4{F0: 255, F1: ("a" + ("foo bar" + ""))}
        v42 := GenS6{F0: 42, F1: false, F2: ((1.5 + -2.25) < (0.0 * ((0.1 - 1.5) - 0.1)))}
        v43 := GenS4{F0: 42, F1: ""}
        _ = v41
        _ = v42
        _ = v43
    }
    v20 = append(v20, v21)
    v2 -= (v2 | v2)
    var v44 string = ("a" + ("" + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v44
}
