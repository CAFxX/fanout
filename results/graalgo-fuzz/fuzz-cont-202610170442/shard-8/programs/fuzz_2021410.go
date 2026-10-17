package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint }
func genFn0(p0 int64, p1 uint) int8 {
    var v6 int16 = 42
    var v7 int = 490
    var v8 uint32 = 2
    v9 := GenS1{F0: -419, F1: 2, F2: "hello"}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 100
}
func genFn1(p0 int, p1 uint64, p2 uint) uint {
    v10 := []int{}
    var v11 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    var v1 float32 = (float32(3.14159) - 0.0)
    var v2 uint16 = 418
    var v3 float32 = (v1 / v1)
    v4 := make(map[uint32]bool)
    v5 := []int{}
    if false {
        var v12 uint32 = 65535
        if ((true || true) || true) {
            fmt.Println(v12, v1)
            var v13 float32 = v3
            if ((3.14159 <= -2.25) && (!false)) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        v16 := v4[(v12 + 211)]
                        if (!(411 < 0)) {
                            if len(v5) > 0 {
                                v5[0] = (fi15 & fi15)
                            }
                            switch (fi14 % ((fi15) | 1)) {
                            case -1:
                                if false {
                                    if true {
                                        switch (fi15 / ((int(v2)) | 1)) {
                                        case -1:
                                            var v17 int16 = -963
                                            _ = v17
                                        case 127:
                                            switch (fi15 % ((32768) | 1)) {
                                            case -9223372036854775808:
                                                var v18 uint8 = 127
                                                _ = v18
                                            }
                                        case 1:
                                            var v19 int8 = -128
                                            _ = v19
                                        default:
                                            v2 += 42
                                        }
                                        v4[(^(v12))] = (938 < 97)
                                        v20 := GenS1{F0: 1000, F1: uint8(v2), F2: ("a" + ("a" + "\n\t"))}
                                        _ = v20
                                    } else {
                                        v2 -= 65535
                                    }
                                    if ((fi15 & fi14) >= (^(fi15))) {
                                        var v21 uint32 = v12
                                        _ = v21
                                    }
                                    var v22 int = (fi15 >> ((383) & 63))
                                    _ = v22
                                } else {
                                    v16 = true
                                    var v23 bool = (v16 || v16)
                                    _ = v23
                                }
                            }
                            var v24 uint64 = 4294967295
                            _ = v24
                        }
                        _ = fi15
                        _ = v16
                    }
                    for fi25 := 0; fi25 < 14; fi25++ {
                        if false {
                            var v26 uint = 4294967295
                            if true {
                                v1 -= float32(-2.25)
                                v2 += v2
                                switch (-(fi25)) {
                                case -9223372036854775808:
                                    if (-9223372036854775808 != 373) {
                                        v27 := GenS1{F0: -409, F1: 42, F2: ("\n\t" + "")}
                                        v28 := GenS2{F0: ("hello" + (("" + "\n\t") + (("\n\t" + (("a" + "hello") + "hello")) + "\n\t"))), F1: false, F2: 835}
                                        if true {
                                            fmt.Println(v3)
                                        } else {
                                            var v29 uint8 = 40
                                            v3 = v1
                                            _ = v29
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                case 128:
                                    v2 = v2
                                    v30 := GenS3{F0: "", F1: v12}
                                    _ = v30
                                case 4294967295:
                                    if false {
                                        if (v12 > v12) {
                                            v13 = 3.14159
                                        } else {
                                            var v31 int = fi14
                                            v32 := GenS1{F0: -832, F1: 128, F2: "hello"}
                                            _ = v31
                                            _ = v32
                                        }
                                    } else {
                                        fmt.Println(fi25, fi14)
                                        if (true || (42 != genFn0(743, 507))) {
                                            if false {
                                                v3 -= v3
                                                var v33 uint16 = 166
                                                var v34 int16 = 32767
                                                _ = v33
                                                _ = v34
                                            }
                                            v1 -= 0.0
                                        } else {
                                            v12 += (^(v12))
                                            v13 += 1.5
                                        }
                                        v35 := GenS4{F0: "hello", F1: v12}
                                        var v36 uint8 = 1
                                        _ = v35
                                        _ = v36
                                    }
                                    v4[(v12 ^ (v12 * v12))] = (true && true)
                                    if ((((-2.25 / 1.0) * 1.5) + ((0.1 + 0.1) + 3.14159)) == 1.5) {
                                        v26 += 1
                                        var v37 int32 = -1
                                        _ = v37
                                    }
                                default:
                                    var v38 int64 = -9223372036854775808
                                    _ = v38
                                }
                            } else {
                                fmt.Println(v2)
                                var v39 int16 = -1
                                v2 = (v2 ^ 65535)
                                _ = v39
                            }
                            var v40 int = fi14
                            _ = v26
                            _ = v40
                        }
                        _ = fi25
                    }
                    _ = fi14
                }
                var v41 string = "a"
                _ = v41
            } else {
                fmt.Println(v13, v12)
                _ = len(v4)
            }
            _ = v13
        }
        if true {
            var v42 uint16 = (v2 >> ((338) & 63))
            if false {
                for fi43 := 0; fi43 < 17; fi43++ {
                    if true {
                        if (734 <= genFn1(fi43, 4294967295, 691)) {
                            var v44 uint = 42
                            var v45 float64 = (((((1.5 - (float64(v1) - (1.5 * (-2.25 + float64(v1))))) - (1e10 - -2.25)) / 1.0) * 3.14159) + ((-2.25 + 0.0) * ((-2.25 * 1.5) - 1e10)))
                            var v46 uint64 = 755
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        v47 := GenS1{F0: 1, F1: 255, F2: ("hello" + "hello")}
                        v12 = (v12 & 797)
                        _ = v47
                    } else {
                        var v48 string = ("a" + ("a" + (((("\n\t" + "hello") + "hello") + "\n\t") + "a")))
                        if false {
                            v1 -= 1e10
                        } else {
                            var v49 float64 = (0.0 * 1e10)
                            _ = v49
                        }
                        switch fi43 {
                        case -2147483648:
                            for fi50 := 0; fi50 < 17; fi50++ {
                                v51 := GenS4{F0: "a", F1: 2147483648}
                                _ = fi50
                                _ = v51
                            }
                        case 4294967295:
                            v42 = v2
                        case 32768:
                            v2 = (v42 % (((v42 / (((^(v2))) | 1))) | 1))
                        default:
                            if len(v5) > 0 {
                                v5[0] = int(v12)
                            }
                        }
                        _ = v48
                    }
                    v2 -= (v42 << ((18446744073709551615) & 63))
                    _ = fi43
                }
                v3 += ((v1 * ((v3 / (float32((float64((v3 * v1)) + 3.14159)) + v1)) / v3)) / v3)
            } else {
                v52 := GenS3{F0: ((("foo bar" + (((("hello" + "foo bar") + "") + (("" + "a") + "")) + "hello")) + "a") + "hello"), F1: v12}
                _ = v52
            }
            fmt.Println(v1)
            _ = v42
        } else {
            var v53 uint8 = 2
            _ = v53
        }
        _ = v12
    }
    v54 := GenS5{F0: 18446744073709551615}
    for fi55 := 0; fi55 < 6; fi55++ {
        v56 := GenS5{F0: 4294967295}
        _ = fi55
        _ = v56
    }
    var v57 uint = genFn1(-2147483648, 676, genFn1(127, 117, genFn1(42, 91, 532)))
    v57 += v57
    var v58 float64 = 1e10
    if false {
        switch "foo bar" {
        case "":
            fmt.Println(v2, v1)
        case "hello":
            var v59 bool = (!(false && false))
            _ = v59
        case "foo bar":
            v58 = ((0.1 - v58) / 3.14159)
        default:
            var v60 int64 = 163
            _ = v60
        }
        v1 = 3.14159
    } else {
        v61 := GenS1{F0: -104, F1: uint8(v2), F2: ("a" + "foo bar")}
        _ = v61
    }
    var v62 uint = v57
    var v63 string = ""
    v5 = append(v5, 32768)
    fmt.Println(v62, v1)
    v64 := GenS4{F0: "foo bar", F1: 2}
    var v65 int8 = -128
    fmt.Println(v1, v2)
    var v66 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v57
    _ = v58
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
