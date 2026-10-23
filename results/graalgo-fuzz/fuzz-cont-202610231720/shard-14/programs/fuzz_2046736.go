package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16; F1 uint64; F2 bool }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0() int16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 float32 = 0.1
        v8 := GenS1{F0: 1}
        if (741 < 42) {
            v7 -= (v7 + float32(1e10))
            for fi9 := 0; fi9 < 4; fi9++ {
                if (255 <= 42) {
                    if true {
                        v7 += v7
                        v5 = append(v5, fi9)
                        v5 = append(v5, fi9)
                    }
                }
                fmt.Println(fi9, fi6, fi9)
                _ = fi9
            }
            v10 := GenS1{F0: -1}
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if (!false) {
        if len(v5) > 0 {
            v5[0] = -958
        }
        var v11 uint32 = 298
        _ = v11
    }
    var v12 float32 = (3.14159 * float32(float64(3.14159)))
    var v13 float32 = (v12 + (((v12 / float32(1.5)) - (v12 * (v12 / v12))) + 0.1))
    _ = v12
    _ = v13
    _ = v5
    return -32768
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 float64 = 1e10
    v4 := make(map[int]uint32)
    v14 := []int{42, -9223372036854775808, 1, -9223372036854775808}
    v3 = 3.14159
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int32 = -1
        if false {
            for fi17 := 0; fi17 < 5; fi17++ {
                switch "a" {
                case "":
                    if true {
                        var v18 uint64 = 0
                        _ = v18
                    } else {
                        v19 := GenS2{F0: 985, F1: (!(!(true || false))), F2: -275}
                        delete(v2, "\n\t")
                        _ = v19
                    }
                case "\n\t":
                    v20 := GenS2{F0: 2, F1: false, F2: -1}
                    _ = v20
                }
                switch ("a" + "\n\t") {
                case "hello":
                    v21 := GenS3{F0: 42, F1: 42}
                    _ = v21
                default:
                    if (false || (32767 != 32767)) {
                        if true {
                            v4[fi17] = 518
                            fmt.Println(fi17, v3, fi17)
                        } else {
                            switch ("a" + "hello") {
                            case "a":
                                fmt.Println(fi17, v3)
                            case "\n\t":
                                var v22 uint = 430
                                _ = v22
                            }
                            var v23 bool = false
                            if (!v23) {
                                v24 := v2[("foo bar" + "foo bar")]
                                if (!((false || v23) && false)) {
                                    if true {
                                        var v25 int8 = 1
                                        var v26 uint8 = 42
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 string = "hello"
                                        var v28 string = v27
                                        switch fi17 {
                                        case 565:
                                            fmt.Println(fi15, v27, fi15)
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                    v23 = (false && v23)
                                    if (42 == 65535) {
                                        var v29 int64 = -647
                                        if false {
                                            v30 := v2[""]
                                            _ = v30
                                        } else {
                                            v24 = (v24 | (v24 ^ v24))
                                        }
                                        _ = v29
                                    }
                                    var v31 uint8 = 72
                                    _ = v31
                                } else {
                                    var v32 int8 = 127
                                    _ = v32
                                }
                                _ = v24
                            }
                            _ = v23
                        }
                        switch fi15 {
                        case 128:
                            v3 = v3
                        case 2147483647:
                            v16 += (^(v16))
                        case 32768:
                            v33 := GenS2{F0: 2147483648, F1: false, F2: 0}
                            _ = v33
                        }
                        if false {
                            var v34 int32 = (v16 - -2147483648)
                            _ = v34
                        }
                    } else {
                        switch (fi15 | (fi15 % (((fi17 & (fi15 * fi17))) | 1))) {
                        case 649:
                            var v35 int16 = 32767
                            _ = v35
                        }
                        v36 := v2["hello"]
                        var v37 uint = 292
                        _ = v36
                        _ = v37
                    }
                }
                if false {
                    v3 = (v3 - v3)
                    delete(v4, fi15)
                    fmt.Println(fi17, fi15, fi17)
                }
                _ = fi17
            }
            v3 -= 1.5
            if (4294967295 != 478) {
                var v38 float64 = float64(float32((v3 * float64(-2.25))))
                v3 -= v3
                _ = v38
            }
        } else {
            switch "a" {
            case "":
                v3 = v3
            default:
                switch "\n\t" {
                case "":
                    var v39 float32 = (-2.25 - ((3.14159 * 1e10) / (-2.25 - 3.14159)))
                    _ = v39
                default:
                    v16 = v16
                }
            }
            v40 := GenS4{F0: 32767, F1: 42, F2: (!((v16 - v16) != 1))}
            _ = v40
        }
        v41 := GenS5{F0: false, F1: 65535}
        _ = fi15
        _ = v16
        _ = v41
    }
    switch "\n\t" {
    case "hello":
        var v42 uint = 0
        _ = v42
    case "":
        for fi43 := 0; fi43 < 17; fi43++ {
            for fi44 := 0; fi44 < 13; fi44++ {
                var v45 int = fi43
                _ = len(v4)
                v3 -= (v3 * float64(0.1))
                _ = fi44
                _ = v45
            }
            var v46 uint32 = 1
            if (-440 <= genFn0()) {
                v3 = v3
                switch ((("hello" + ("foo bar" + "hello")) + ("" + (("a" + "\n\t") + ""))) + ("\n\t" + ("\n\t" + "\n\t"))) {
                case "hello":
                    var v47 int8 = 79
                    _ = v47
                case "":
                    var v48 int = fi43
                    _ = v48
                case "foo bar":
                    for fi49 := 0; fi49 < 8; fi49++ {
                        if false {
                            v46 += (^(v46))
                            v50 := v2[("hello" + "\n\t")]
                            v51 := GenS4{F0: genFn0(), F1: 389, F2: false}
                            _ = v50
                            _ = v51
                        }
                        var v52 uint16 = 89
                        v4[fi43] = v46
                        _ = fi49
                        _ = v52
                    }
                }
                if len(v14) > 0 {
                    v14[0] = fi43
                }
            }
            _ = fi43
            _ = v46
        }
    case "foo bar":
        v14 = append(v14, 42)
    default:
        v3 -= v3
    }
    var v53 bool = true
    v53 = false
    for fi54 := 0; fi54 < 1; fi54++ {
        switch fi54 {
        case -1:
            v55 := GenS3{F0: 213, F1: 700}
            _ = v55
        case 1:
            fmt.Println(fi54)
        case -23:
            v56 := GenS6{F0: 453, F1: 2, F2: 2}
            _ = v56
        }
        var v57 bool = v53
        var v58 int64 = -899
        _ = fi54
        _ = v57
        _ = v58
    }
    fmt.Println(v3, v3)
    v53 = (255 != 255)
    switch -2147483648 {
    case 127:
        var v59 int16 = -1
        _ = v59
    case 7:
        var v60 int = 588
        _ = v60
    }
    for fi61 := 0; fi61 < 19; fi61++ {
        if v53 {
            v3 = 0.1
            v62 := GenS4{F0: -29, F1: 761, F2: v53}
            _ = v62
        } else {
            v63 := GenS2{F0: 2, F1: (((v3 + 3.14159) * (v3 * (1.5 * 0.1))) <= v3), F2: 42}
            if (635 < 4294967295) {
                v4[(fi61 >> ((34) & 63))] = 2
            }
            _ = v63
        }
        switch (fi61 >> ((18446744073709551615) & 63)) {
        case 0:
            v53 = v53
        case -345:
            _ = len(v14)
        case 127:
            fmt.Println(v3, v3)
        }
        _ = fi61
    }
    v53 = v53
    var v64 int = -135
    v3 = v3
    if ((v53 || true) && v53) {
        v65 := GenS3{F0: 18446744073709551615, F1: -1}
        v3 -= float64((3.14159 + -2.25))
        _ = v65
    }
    var v66 float32 = 0.0
    switch ("" + ("\n\t" + ("" + ("foo bar" + (((("\n\t" + (("" + "foo bar") + ("" + "hello"))) + "") + "hello") + ""))))) {
    case "":
        v67 := GenS2{F0: 485, F1: (1 >= int8(v64)), F2: genFn0()}
        _ = v67
    case "foo bar":
        for fi68 := 0; fi68 < 5; fi68++ {
            if (297 > 588) {
                v3 = v3
                v53 = ((2 <= 65535) || v53)
                v64 -= 32768
            } else {
                v69 := GenS5{F0: (!v53), F1: 2}
                for fi70 := 0; fi70 < 6; fi70++ {
                    var v71 int = fi68
                    _ = fi70
                    _ = v71
                }
                _ = v69
            }
            _ = len(v14)
            _ = fi68
        }
    }
    var v72 int = (v64 + (v64 - (^(v64))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v64
    _ = v66
    _ = v72
}
