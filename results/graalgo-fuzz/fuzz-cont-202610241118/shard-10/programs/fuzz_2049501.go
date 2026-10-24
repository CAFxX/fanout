package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 bool; F1 uint16; F2 string }
func genFn0() uint {
    var v5 float64 = (3.14159 + 1e10)
    var v6 string = "hello"
    var v7 int8 = -59
    if false {
        v7 += 42
    } else {
        v6 = ("foo bar" + "foo bar")
    }
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 int32) int {
    var v8 string = "\n\t"
    var v9 int32 = p1
    if ((!false) || true) {
        var v10 int = -9223372036854775808
        _ = v10
    } else {
        var v11 int16 = 32767
        v12 := GenS1{F0: 589, F1: genFn0(), F2: true}
        _ = v11
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    _ = v9
    return -133
}

func main() {
    var v1 int8 = 127
    var v2 uint64 = 710
    var v3 int = -435
    var v4 bool = true
    if v4 {
        v14 := []int{v3, (-(v3)), v3}
        if ((v2 >> ((1) & 63)) == v2) {
            var v15 string = ("" + "")
            for fi16 := 0; fi16 < 11; fi16++ {
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS1{F0: 60, F1: 830, F2: (423 >= 211)}
                    if (v3 == int(v1)) {
                        var v19 int = (v3 | fi17)
                        v1 = 100
                        _ = v19
                    }
                    v20 := GenS2{F0: 127, F1: v4, F2: (v1 << ((0) & 63))}
                    _ = fi17
                    _ = v18
                    _ = v20
                }
                _ = fi16
            }
            _ = v15
        } else {
            if (!v4) {
                var v21 float64 = (0.1 / 0.1)
                var v22 bool = false
                _ = v21
                _ = v22
            }
            var v23 int64 = -578
            _ = len(v14)
            _ = v23
        }
        _ = v14
    } else {
        var v24 int16 = -827
        fmt.Println(v1, v24, v2)
        for fi25 := 0; fi25 < 6; fi25++ {
            v3 += int(v1)
            fmt.Println(v2, v4, v3)
            fmt.Println(v3, fi25)
            _ = fi25
        }
        _ = v24
    }
    if ("hello" >= "\n\t") {
        for fi26 := 0; fi26 < 20; fi26++ {
            v4 = ("" > "hello")
            fmt.Println(v4, v3, v3)
            if v4 {
                fmt.Println(v2, v2, v4)
            }
            _ = fi26
        }
        var v27 int32 = 0
        _ = v27
    } else {
        v2 -= v2
        v1 -= (-(v1))
        if ((float64(float32(-2.25)) != 3.14159) || (false && (564 < -448))) {
            v2 = 18446744073709551615
        }
    }
    v3 = v3
    v28 := []int{}
    fmt.Println(v4, v4, v3)
    if len(v28) >= 2 {
        v28 = v28[1:len(v28)]
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        if (2 >= 729) {
            if v4 {
                v2 = 4294967295
                v30 := GenS3{F0: 0, F1: 42, F2: -1}
                var v31 string = "hello"
                _ = v30
                _ = v31
            }
            v4 = (v4 && (3.14159 < 1.5))
            if (v4 || true) {
                switch fi29 {
                case 269:
                    if (((("\n\t" + (("" + "") + "\n\t")) + "\n\t") + "\n\t") == (("\n\t" + "a") + "\n\t")) {
                        switch "hello" {
                        case "foo bar":
                            if v4 {
                                var v32 float64 = ((float64(-2.25) / 1e10) - (1.5 * (float64(1.5) - (1e10 * float64(((0.1 * (1.5 / 3.14159)) + (1e10 - 0.0)))))))
                                _ = v32
                            }
                        case "hello":
                            var v33 float32 = ((float32(3.14159) - (-2.25 - -2.25)) - 1.5)
                            _ = v33
                        case "a":
                            for fi34 := 0; fi34 < 16; fi34++ {
                                var v35 float32 = (float32((0.0 / (((3.14159 * -2.25) / 3.14159) - 1e10))) / 1e10)
                                fmt.Println(v3, v2)
                                _ = fi34
                                _ = v35
                            }
                        default:
                            if v4 {
                                v2 = (v2 % ((v2) | 1))
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            if v4 {
                                switch (-(v3)) {
                                case 4294967295:
                                    v4 = (!false)
                                case -2147483648:
                                    if len(v28) > 0 {
                                        v28[0] = v3
                                    }
                                }
                                fmt.Println(v1, v4)
                                switch (^(v3)) {
                                case 127:
                                    fmt.Println(v1)
                                }
                            } else {
                                v2 = 1
                                fmt.Println(v2, fi29)
                            }
                        case "foo bar":
                            var v36 int8 = -128
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 17; fi37++ {
                            v1 += (-(v1))
                            v28 = append(v28, fi29)
                            _ = len(v28)
                            v28 = append(v28, (^(v3)))
                            _ = fi37
                        }
                    } else {
                        for fi38 := 0; fi38 < 1; fi38++ {
                            if len(v28) > 0 {
                                v28[0] = (fi38 * fi29)
                            }
                            if true {
                                v39 := GenS3{F0: 9223372036854775807, F1: 127, F2: 0}
                                v40 := GenS4{F0: -346, F1: -461}
                                if len(v28) > 0 {
                                    v28[0] = v3
                                }
                                _ = v39
                                _ = v40
                            } else {
                                if (1 <= 42) {
                                    var v41 uint8 = 56
                                    var v42 int16 = 32767
                                    if (789 < 42) {
                                        var v43 float64 = 3.14159
                                        v43 += v43
                                        switch int(v41) {
                                        case 0:
                                            var v44 int8 = v1
                                            _ = v44
                                        case -389:
                                            var v45 int32 = 0
                                            _ = v45
                                        case -518:
                                            v3 += fi29
                                        }
                                        _ = v43
                                    } else {
                                        v1 = (v1 - 100)
                                        var v46 int32 = -25
                                        var v47 int = genFn1(v2, (-(v46)))
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := GenS3{F0: -1, F1: (v41 * v41), F2: 1}
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    _ = v41
                                    _ = v42
                                    _ = v48
                                } else {
                                    v4 = (!(1 > 117))
                                    var v49 uint8 = 42
                                    _ = v49
                                }
                            }
                            _ = fi38
                        }
                    }
                case 1:
                    var v50 uint8 = 255
                    _ = v50
                case 32768:
                    v3 = (fi29 % ((fi29) | 1))
                default:
                    v2 = 846
                }
            }
        }
        if (1e10 >= float32(1.5)) {
            _ = len(v28)
        }
        _ = fi29
    }
    v51 := GenS3{F0: 1, F1: 2, F2: 294}
    v3 += -9223372036854775808
    v2 = v2
    if false {
        if v4 {
            var v52 uint32 = 65535
            for fi53 := 0; fi53 < 5; fi53++ {
                v54 := GenS5{F0: (3.14159 < (1.5 * 3.14159)), F1: 0, F2: "foo bar"}
                v1 = (v1 / ((100) | 1))
                _ = fi53
                _ = v54
            }
            fmt.Println(v2, v1, v2)
            _ = v52
        } else {
            fmt.Println(v2)
            if v4 {
                if (1 > 1) {
                    var v55 int64 = 1
                    if (v4 || (v4 || (v4 && v4))) {
                        var v56 uint = 2
                        if v4 {
                            fmt.Println(v56, v4)
                        }
                        for fi57 := 0; fi57 < 5; fi57++ {
                            for fi58 := 0; fi58 < 14; fi58++ {
                                v1 = 104
                                var v59 uint32 = 2147483647
                                var v60 uint16 = 1
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            _ = fi57
                        }
                        _ = v56
                    } else {
                        if len(v28) > 0 {
                            v28[0] = v3
                        }
                        if v4 {
                            v1 = -110
                            var v61 int64 = v55
                            v3 -= v3
                            _ = v61
                        }
                        var v62 uint64 = v2
                        _ = v62
                    }
                    _ = v55
                }
                v1 += (v1 - (^(v1)))
                for fi63 := 0; fi63 < 13; fi63++ {
                    if len(v28) > 0 {
                        v28[0] = genFn1((v2 << ((genFn0()) & 63)), 2147483647)
                    }
                    v4 = false
                    _ = fi63
                }
            } else {
                var v64 string = "hello"
                _ = v64
            }
            _ = len(v28)
        }
    } else {
        switch "a" {
        case "":
            var v65 uint = 1
            _ = v65
        case "\n\t":
            var v66 int = v3
            _ = v66
        }
        var v67 uint64 = 42
        _ = v67
    }
    var v68 float32 = 1.5
    v28 = append(v28, (v3 * (v3 + 128)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v51
    _ = v68
}
