package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int64 }
func genFn0(p0 uint32, p1 int, p2 uint32) int {
    for fi5 := 0; fi5 < 2; fi5++ {
        p1 += (p1 | (p1 & 9223372036854775807))
        var v6 int64 = -1
        for fi7 := 0; fi7 < 8; fi7++ {
            if (-1 >= 42) {
                v8 := []int{}
                _ = v8
            } else {
                var v9 int = -905
                v10 := GenS1{F0: 0, F1: v9}
                _ = v10
                _ = v9
            }
            var v11 float64 = (3.14159 / (float64(0.0) - 1e10))
            _ = fi7
            _ = v11
        }
        _ = fi5
        _ = v6
    }
    v12 := []int{p1, 1, p1, (p1 << ((372) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}
func genFn1() uint64 {
    var v13 int = 9223372036854775807
    var v14 bool = true
    _ = v13
    _ = v14
    return 130
}

func main() {
    var v1 bool = (!true)
    var v2 float32 = -2.25
    v3 := make(map[string]uint32)
    v4 := []int{-461}
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            fmt.Println(v2, v2)
            _ = fi16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            if v1 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v18 uint16 = 791
                var v19 float64 = 0.0
                _ = v18
                _ = v19
            } else {
                fmt.Println(fi17, fi15)
                v20 := GenS1{F0: 1, F1: fi17}
                v2 -= (1e10 / 3.14159)
                _ = v20
            }
            v2 -= ((v2 - (float32(-2.25) * 3.14159)) / 3.14159)
            if (-1 <= 42) {
                fmt.Println(fi15, v2)
                switch "foo bar" {
                case "a":
                    if (!(2147483647 > 0)) {
                        var v21 uint = 135
                        if (!true) {
                            if v1 {
                                _ = len(v4)
                            }
                            if v1 {
                                _ = len(v3)
                                var v22 int16 = 0
                                _ = v22
                            }
                        }
                        v2 = v2
                        _ = v21
                    } else {
                        fmt.Println(fi15, fi17)
                    }
                case "foo bar":
                    switch fi17 {
                    case 127:
                        var v23 float32 = v2
                        _ = v23
                    case 128:
                        v1 = v1
                    default:
                        switch "\n\t" {
                        case "":
                            delete(v3, "hello")
                        case "hello":
                            var v24 uint = 42
                            _ = v24
                        default:
                            var v25 int32 = 0
                            _ = v25
                        }
                    }
                default:
                    v26 := v3["\n\t"]
                    _ = v26
                }
            }
            switch (fi15 + fi17) {
            case -327:
                if (255 == 1) {
                    v2 = v2
                    if true {
                        v27 := GenS2{F0: v1}
                        _ = v27
                    }
                    if true {
                        if v1 {
                            switch "" {
                            case "hello":
                                v28 := GenS1{F0: -25, F1: 2147483647}
                                _ = v28
                            case "\n\t":
                                v2 = v2
                            default:
                                fmt.Println(v1, fi17, fi15)
                            }
                        } else {
                            var v29 int8 = 42
                            if true {
                                v2 = 1e10
                                v30 := GenS3{F0: 127, F1: 65535, F2: 18446744073709551615}
                                var v31 uint32 = 609
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        }
                    } else {
                        var v32 uint8 = 42
                        fmt.Println(fi17, fi15, fi17)
                        _ = v32
                    }
                }
            case 1:
                if (2 < 114) {
                    v33 := v3[((("hello" + "\n\t") + ("" + ((((("a" + "") + "hello") + "a") + "") + ("" + "")))) + (("a" + "foo bar") + "foo bar"))]
                    _ = v33
                }
            case 42:
                var v34 int = (fi15 | (^(fi15)))
                _ = v34
            }
            _ = fi17
        }
        _ = fi15
    }
    var v35 uint32 = 0
    if false {
        for fi36 := 0; fi36 < 10; fi36++ {
            switch 633 {
            case 924:
                if false {
                    var v37 uint8 = uint8(v35)
                    var v38 uint8 = (^(v37))
                    v1 = (false && (!v1))
                    _ = v37
                    _ = v38
                }
            case 42:
                v35 = (v35 * (v35 & 36))
            case 452:
                v39 := GenS3{F0: (fi36 * (fi36 * (fi36 & fi36))), F1: 18446744073709551615, F2: 18446744073709551615}
                _ = v39
            }
            fmt.Println(v35, v2)
            fmt.Println(v1, v2, fi36)
            _ = fi36
        }
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS3{F0: 386, F1: 4294967295, F2: 42}
        v1 = true
        _ = fi40
        _ = v41
    }
    if ((((v2 * ((v2 - (float32(((float64(v2) - (1e10 / ((0.1 + (1e10 * -2.25)) * (3.14159 / -2.25)))) / 0.1)) * ((v2 + float32(0.0)) / (float32(float64(1.5)) + v2)))) / 1e10)) <= 3.14159) || v1) && v1) {
        var v42 int8 = 127
        v43 := v3[("foo bar" + (("" + "") + "foo bar"))]
        _ = v42
        _ = v43
    }
    switch "hello" {
    case "foo bar":
        v1 = v1
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        if v1 {
            switch genFn0(213, 32767, (v35 % (((^(v35))) | 1))) {
            case 9223372036854775807:
                if v1 {
                    v1 = true
                }
            default:
                v35 = (v35 | (v35 * v35))
            }
        } else {
            var v45 uint32 = v35
            for fi46 := 0; fi46 < 5; fi46++ {
                v35 = (v35 / (((^(v35))) | 1))
                _ = fi46
            }
            _ = v45
        }
        var v47 uint8 = 42
        for fi48 := 0; fi48 < 11; fi48++ {
            v47 = (^(v47))
            _ = fi48
        }
        _ = fi44
        _ = v47
    }
    v2 -= (v2 + (v2 - (v2 + v2)))
    if (!(-785 != 9223372036854775807)) {
        if (v35 >= v35) {
            v49 := GenS1{F0: 0, F1: -9223372036854775808}
            fmt.Println(v2, v35)
            v1 = ("" >= "foo bar")
            _ = v49
        }
        v50 := GenS4{F0: 2147483648, F1: 65, F2: -1}
        _ = v50
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = v2
    if v1 {
        var v51 float32 = float32(3.14159)
        _ = v51
    } else {
        v52 := GenS3{F0: -129, F1: 440, F2: 0}
        _ = v52
    }
    var v53 int8 = 127
    v1 = (1 > 4294967295)
    v35 = (v35 - v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v53
}
