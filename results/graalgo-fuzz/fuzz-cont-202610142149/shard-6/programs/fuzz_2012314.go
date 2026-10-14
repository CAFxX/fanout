package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 int32) int32 {
    switch "" {
    case "a":
        v4 := []int{int(p1)}
        _ = v4
    case "\n\t":
        var v5 uint32 = 1
        _ = v5
    }
    p1 = (^(p1))
    var v6 int = 299
    _ = p0
    _ = p1
    _ = v6
    return (p1 * 0)
}
func genFn1(p0 int16, p1 string) uint {
    v7 := []int{-1, 853}
    v8 := GenS1{F0: 422}
    for fi9 := 0; fi9 < 15; fi9++ {
        if len(v7) > 0 {
            v7[0] = (fi9 % (((fi9 * (fi9 + fi9))) | 1))
        }
        if (0 == (fi9 % ((fi9) | 1))) {
            if (-257 >= genFn0(2, -2147483648)) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v10 uint = 65535
                _ = v10
            } else {
                p1 = "\n\t"
                v11 := GenS1{F0: 255}
                var v12 uint64 = 42
                _ = v11
                _ = v12
            }
            v13 := GenS2{F0: 1}
            _ = v13
        }
        switch fi9 {
        case 1:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 float64 = 1e10
                p0 -= p0
                v16 := GenS1{F0: 1}
                if len(v7) > 0 {
                    v7[0] = 9223372036854775807
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        case 9223372036854775807:
            p0 -= -1
        default:
            if true {
                fmt.Println(fi9)
                fmt.Println(p1)
            } else {
                switch -347 {
                case -1:
                    if true {
                        for fi17 := 0; fi17 < 13; fi17++ {
                            if (0 > -128) {
                                p1 = p1
                            } else {
                                switch -886 {
                                case 401:
                                    var v18 string = p1
                                    _ = v18
                                default:
                                    p0 -= 1000
                                    var v19 float64 = 1e10
                                    if len(v7) > 0 {
                                        v7[0] = (fi17 + 127)
                                    }
                                    v19 -= 1e10
                                    switch (((((p1 + p1) + (p1 + p1)) + (p1 + (p1 + ""))) + (p1 + "foo bar")) + p1) {
                                    case "":
                                        p0 = (^(p0))
                                    case "a":
                                        fmt.Println(p0, v19)
                                    }
                                    _ = v19
                                }
                                _ = len(v7)
                                if (true && true) {
                                    var v20 float32 = 0.0
                                    v21 := GenS1{F0: 2}
                                    _ = v20
                                    _ = v21
                                } else {
                                    if (1e10 != 0.1) {
                                        p0 = (-(p0))
                                        v22 := GenS1{F0: 18446744073709551615}
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        _ = v22
                                    } else {
                                        var v23 uint = 0
                                        var v24 uint32 = 2
                                        _ = v23
                                        _ = v24
                                    }
                                    v25 := GenS2{F0: -1}
                                    v7 = append(v7, (fi9 | (fi9 + fi9)))
                                    _ = v25
                                }
                            }
                            v26 := GenS3{F0: -27, F1: 18446744073709551615}
                            _ = fi17
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 8; fi27++ {
                            switch p1 {
                            case "foo bar":
                                var v28 float32 = -2.25
                                _ = v28
                            }
                            p1 = p1
                            _ = fi27
                        }
                    }
                }
                p0 = p0
                _ = len(v7)
            }
        }
        _ = fi9
    }
    v29 := GenS1{F0: 65535}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v29
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -32768
    var v3 uint = 42
    if (11 != 0) {
        var v30 bool = true
        _ = v30
    } else {
        for fi31 := 0; fi31 < 4; fi31++ {
            var v32 string = "a"
            var v33 int64 = 1
            var v34 float32 = float32(-2.25)
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
        v35 := GenS3{F0: 81, F1: v3}
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 int = -9223372036854775808
            v38 := []int{32767}
            var v39 uint64 = 267
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
        _ = v35
    }
    var v40 uint = v3
    v3 = (v40 + 42)
    if (!(-1 <= -128)) {
        if (true || (false && true)) {
            v40 = v40
            fmt.Println(v2)
        }
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 float64 = 1e10
            v43 := GenS4{F0: (v3 | 18446744073709551615), F1: 0}
            _ = fi41
            _ = v42
            _ = v43
        }
        for fi44 := 0; fi44 < 16; fi44++ {
            v3 -= v3
            fmt.Println(v3, v2)
            _ = fi44
        }
    }
    var v45 uint8 = 2
    v46 := GenS5{F0: false, F1: 192, F2: 9223372036854775807}
    v40 -= v3
    var v47 uint32 = 42
    v2 += (v2 * (^(v2)))
    v45 = 255
    v40 = 255
    switch 610 {
    case 42:
        v48 := GenS2{F0: 0}
        _ = v48
    case -9223372036854775808:
        var v49 uint = 775
        _ = v49
    default:
        var v50 uint32 = 1
        _ = v50
    }
    var v51 int16 = (v2 % ((v2) | 1))
    v45 = (v45 * v45)
    var v52 int64 = -700
    for fi53 := 0; fi53 < 16; fi53++ {
        v54 := v1[fi53]
        v55 := []int{}
        var v56 bool = true
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    for fi57 := 0; fi57 < 12; fi57++ {
        v3 = (^(v3))
        fmt.Println(v52, v45, fi57)
        _ = fi57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v45
    _ = v46
    _ = v47
    _ = v51
    _ = v52
}
