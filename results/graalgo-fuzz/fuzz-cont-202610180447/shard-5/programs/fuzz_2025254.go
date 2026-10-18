package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 bool = (true && true)
        v7 = ((false || true) || ((42 != 42) || ((0 <= -604) && v7)))
        _ = fi6
        _ = v7
    }
    var v8 string = "hello"
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 float64 = 3.14159
        v8 = ("hello" + v8)
        v10 = v10
        _ = fi9
        _ = v10
    }
    _ = v8
    return 42
}
func genFn1(p0 bool, p1 int32) int {
    if (true || p0) {
        if false {
            if p0 {
                v11 := []int{4294967295, 852, 2147483647}
                var v12 bool = (p0 && (0 < -1))
                _ = v11
                _ = v12
            }
        }
        var v13 uint32 = 2147483648
        if (uint(v13) != 65535) {
            v13 = v13
            p1 = -897
        }
        _ = v13
    } else {
        v14 := []int{128}
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int32 = (p1 ^ 37)
            if ((fi15 >> ((2) & 63)) != fi15) {
                fmt.Println(p0, fi15, p0)
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if p0 {
                        var v18 int = (fi15 | fi17)
                        _ = v18
                    } else {
                        switch "a" {
                        case "a":
                            if p0 {
                                _ = len(v14)
                                p1 = (^(v16))
                                var v19 uint = 2
                                _ = v19
                            }
                        case "foo bar":
                            if (-32768 > 32767) {
                                var v20 int32 = (p1 + p1)
                                _ = v20
                            }
                        default:
                            v16 = -1
                        }
                        if (4294967295 != 2) {
                            var v21 uint8 = 0
                            switch "hello" {
                            case "":
                                if (0 < int8(v16)) {
                                    fmt.Println(p0)
                                    var v22 uint32 = uint32(v21)
                                    if false {
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            var v23 uint16 = 222
                                            _ = v23
                                        }
                                        v24 := GenS1{F0: "foo bar"}
                                        v21 = v21
                                        _ = v24
                                    }
                                    _ = v22
                                } else {
                                    fmt.Println(p0, fi17, p1)
                                    var v25 uint16 = 839
                                    _ = v25
                                }
                            case "a":
                                var v26 int = 32768
                                _ = v26
                            case "foo bar":
                                v27 := GenS1{F0: (("\n\t" + "\n\t") + ("a" + ("\n\t" + "hello")))}
                                _ = v27
                            default:
                                if len(v14) > 0 {
                                    v14[0] = 184
                                }
                            }
                            _ = v21
                        }
                        switch 1 {
                        case -1:
                            var v28 bool = (2147483648 >= 0)
                            _ = v28
                        default:
                            if (373 >= p1) {
                                var v29 int16 = 1000
                                _ = v29
                            } else {
                                fmt.Println(p0, p1, v16)
                                p1 = (v16 * 94)
                                v30 := GenS1{F0: "a"}
                                if (!p0) {
                                    var v31 int16 = 925
                                    _ = v31
                                }
                                _ = v30
                            }
                        }
                    }
                    _ = fi17
                }
                var v32 uint8 = 1
                _ = v32
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    fmt.Println(p0)
    v33 := []int{}
    switch 42 {
    case -147:
        v34 := GenS1{F0: "a"}
        _ = v34
    case -367:
        for fi35 := 0; fi35 < 12; fi35++ {
            p0 = (p0 || (0.1 >= 0.1))
            var v36 int16 = -1
            if false {
                v37 := GenS1{F0: "\n\t"}
                _ = v37
            } else {
                v38 := GenS1{F0: ((("a" + ("hello" + "\n\t")) + "") + "")}
                switch "hello" {
                case "\n\t":
                    var v39 uint = 4294967295
                    _ = v39
                case "foo bar":
                    v36 = (v36 + int16(p1))
                case "a":
                    if true {
                        v40 := GenS2{F0: p0, F1: 42, F2: 176}
                        for fi41 := 0; fi41 < 13; fi41++ {
                            if len(v33) > 0 {
                                v33[0] = fi41
                            }
                            v36 -= (-(v36))
                            _ = fi41
                        }
                        p1 = p1
                        _ = v40
                    } else {
                        _ = len(v33)
                        var v42 int16 = 32767
                        _ = v42
                    }
                default:
                    if (!true) {
                        v43 := GenS2{F0: (100 <= -47), F1: 4294967295, F2: (fi35 | 568)}
                        _ = v43
                    } else {
                        if p0 {
                            if p0 {
                                p1 += p1
                                var v44 int16 = v36
                                _ = v44
                            } else {
                                var v45 float32 = 0.0
                                _ = v45
                            }
                        }
                    }
                }
                _ = v38
            }
            _ = fi35
            _ = v36
        }
    case -1:
        for fi46 := 0; fi46 < 8; fi46++ {
            v47 := GenS1{F0: ""}
            p1 = (p1 / ((p1) | 1))
            var v48 float64 = (0.0 - 3.14159)
            _ = fi46
            _ = v47
            _ = v48
        }
    default:
        v49 := GenS3{F0: 495, F1: 42, F2: 0}
        _ = v49
    }
    var v50 bool = p0
    _ = p0
    _ = p1
    _ = v33
    _ = v50
    return -2147483648
}

func main() {
    v1 := []string{("a" + "a"), "hello", "a", "foo bar"}
    v2 := make(map[string]int)
    var v3 uint64 = 42
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    v51 := GenS3{F0: 396, F1: 0, F2: 1}
    if ("\n\t" >= "a") {
        fmt.Println(v3)
    }
    switch ("foo bar" + ("hello" + (((("hello" + "") + "hello") + "hello") + "foo bar"))) {
    case "":
        v52 := GenS3{F0: (v3 << ((2) & 63)), F1: 42, F2: 287}
        _ = v52
    case "a":
        fmt.Println(v3)
    case "\n\t":
        var v53 int = 9223372036854775807
        _ = v53
    }
    var v54 uint64 = 4294967295
    var v55 int32 = 9
    v54 += (v3 - (v54 * v54))
    if (537 != 2) {
        var v56 int16 = 1
        _ = v56
    } else {
        var v57 int8 = 42
        _ = v57
    }
    fmt.Println(v54)
    switch ("foo bar" + "") {
    case "foo bar":
        fmt.Println(v55)
    }
    var v58 float64 = 0.1
    var v59 uint64 = (^(v54))
    v3 = v59
    v58 = 3.14159
    _ = len(v1)
    var v60 uint64 = 491
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v58
    _ = v59
    _ = v60
}
