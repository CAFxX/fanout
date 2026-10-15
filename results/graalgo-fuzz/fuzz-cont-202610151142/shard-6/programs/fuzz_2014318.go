package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = true
    var v3 uint16 = 255
    v4 := make(map[string]int)
    if (-1 < 0) {
        v2 = (42 < 1)
        v5 := []int{911, 9223372036854775807, -987, 526}
        _ = v5
    } else {
        v2 = true
        v2 = v2
        if v2 {
            var v6 uint32 = uint32(v3)
            _ = v6
        } else {
            var v7 uint64 = 0
            if (32768 < 4294967295) {
                v8 := v1[753]
                if (248 < 42) {
                    if (!v2) {
                        for fi9 := 0; fi9 < 6; fi9++ {
                            v4["foo bar"] = fi9
                            var v10 float64 = (3.14159 / (float64(3.14159) / 1.0))
                            var v11 float64 = 0.1
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                        v12 := GenS1{F0: 32767, F1: 65535, F2: (v8 + v8)}
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            for fi14 := 0; fi14 < 9; fi14++ {
                                fmt.Println(v8, v2, v7)
                                var v15 uint8 = 127
                                var v16 int64 = int64(v3)
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            _ = fi13
                        }
                        v17 := v1[(v8 * v8)]
                        _ = v17
                    }
                    var v18 uint = 28
                    v19 := GenS1{F0: -32768, F1: v3, F2: (^(v8))}
                    _ = v18
                    _ = v19
                } else {
                    var v20 int32 = 0
                    switch 2147483647 {
                    case 511:
                        v7 += v7
                    case -251:
                        if (32767 < 1) {
                            switch (("\n\t" + "") + ("" + "a")) {
                            case "":
                                v20 -= v20
                            case "\n\t":
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    var v22 bool = (v20 <= v20)
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            var v23 string = "\n\t"
                            _ = v23
                        } else {
                            v24 := GenS2{F0: v20, F1: 4294967295, F2: -998}
                            switch 621 {
                            case 32768:
                                var v25 float64 = (0.0 - (((((0.0 + (0.1 - 1.5)) - ((3.14159 / 1.0) - 3.14159)) / 1e10) - 1.5) - 3.14159))
                                _ = v25
                            default:
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    v20 = v20
                                    v7 = 426
                                    _ = fi26
                                }
                            }
                            _ = v24
                        }
                    default:
                        v1[v8] = (v8 | (v8 * v8))
                    }
                    switch (("hello" + "foo bar") + "foo bar") {
                    case "\n\t":
                        var v27 int8 = 1
                        _ = v27
                    case "":
                        for fi28 := 0; fi28 < 1; fi28++ {
                            v29 := GenS1{F0: 1000, F1: (v3 + v3), F2: v8}
                            _ = fi28
                            _ = v29
                        }
                    case "a":
                        v30 := []int{32768, 0, -464}
                        _ = v30
                    }
                    _ = v20
                }
                v8 = v8
                _ = v8
            }
            _ = v7
        }
    }
    var v31 float32 = ((3.14159 - 1e10) - ((1e10 * 1.5) - (-2.25 / 1.0)))
    for fi32 := 0; fi32 < 11; fi32++ {
        if (("hello" + "") > ((((("hello" + (((("" + (("a" + "hello") + "")) + ("hello" + "foo bar")) + ("foo bar" + "\n\t")) + "\n\t")) + ("hello" + "hello")) + (("a" + ("\n\t" + "\n\t")) + ("hello" + ("a" + ("foo bar" + ("hello" + ("" + "\n\t"))))))) + "hello") + ("hello" + ("" + "")))) {
            if ((-1 <= 9) && false) {
                v3 = 65535
                for fi33 := 0; fi33 < 20; fi33++ {
                    if true {
                        var v34 float32 = v31
                        fmt.Println(fi32, v3, v31)
                        fmt.Println(v31, v2, fi32)
                        _ = v34
                    } else {
                        v35 := v1[2147483647]
                        if v2 {
                            v36 := v4[("" + "a")]
                            v37 := GenS1{F0: 364, F1: 1, F2: uint32(v3)}
                            _ = v36
                            _ = v37
                        }
                        v38 := []int{(fi32 | fi32), fi33, fi33, 4294967295}
                        if len(v38) > 0 {
                            v38[0] = 42
                        }
                        var v39 int32 = 0
                        _ = v35
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS2{F0: 2147483647, F1: 275, F2: 42}
                    if true {
                        fmt.Println(fi32, v31)
                        switch 1 {
                        case 128:
                            fmt.Println(fi32, v2, v31)
                        default:
                            v41 := GenS2{F0: -76, F1: 2, F2: -1}
                            _ = v41
                        }
                        if (v3 > (v3 / (((v3 / ((v3) | 1))) | 1))) {
                            fmt.Println(fi33)
                            _ = len(v1)
                            v42 := GenS2{F0: -1, F1: 585, F2: 1}
                            _ = v42
                        } else {
                            var v43 float32 = ((float32(3.14159) * (v31 * ((float32(1.5) - ((v31 / v31) / 1.5)) * v31))) + (((v31 - (v31 / float32(1.5))) * v31) / float32(0.1)))
                            v31 += 0.1
                            switch fi33 {
                            case 32768:
                                switch (fi33 / ((-542) | 1)) {
                                case -830:
                                    v3 = 2
                                case -694:
                                    if (!(v2 || (!true))) {
                                        v44 := []int{}
                                        _ = v44
                                    }
                                case 128:
                                    v43 += ((((1.5 - v31) / (v43 * v31)) - v43) * (v43 + 3.14159))
                                default:
                                    _ = len(v1)
                                }
                            }
                            v45 := GenS2{F0: 0, F1: 65535, F2: 914}
                            _ = v43
                            _ = v45
                        }
                        v46 := GenS1{F0: 1000, F1: 2, F2: 2147483648}
                        var v47 bool = (("hello" > ("\n\t" + "foo bar")) && v2)
                        _ = v46
                        _ = v47
                    } else {
                        v48 := GenS1{F0: 1000, F1: v3, F2: 78}
                        switch fi32 {
                        case -138:
                            if (false || v2) {
                                var v49 int8 = 0
                                v50 := GenS1{F0: -503, F1: v3, F2: uint32(v3)}
                                var v51 int64 = 42
                                _ = v49
                                _ = v50
                                _ = v51
                            }
                        case 128:
                            if false {
                                var v52 uint16 = 2
                                _ = v52
                            } else {
                                var v53 int64 = -6
                                v3 = v3
                                if (!v2) {
                                    switch (^(fi33)) {
                                    case 0:
                                        v31 = (1.5 + float32(float64(v31)))
                                    case 32768:
                                        switch "a" {
                                        case "\n\t":
                                            delete(v1, 65535)
                                        }
                                    }
                                    if (!(!true)) {
                                        if (false && v2) {
                                            fmt.Println(fi33, v53, fi33)
                                            var v54 float32 = (((v31 * v31) - float32(3.14159)) * (float32(3.14159) * v31))
                                            var v55 int16 = -1
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 uint16 = (^(v3))
                                            v2 = true
                                            v57 := GenS2{F0: -1, F1: 18446744073709551615, F2: 585}
                                            _ = v56
                                            _ = v57
                                        }
                                        fmt.Println(v2)
                                    } else {
                                        v58 := GenS3{F0: 65535, F1: v3, F2: "\n\t"}
                                        _ = v58
                                    }
                                }
                                _ = v53
                            }
                        case 42:
                            switch ("" + "foo bar") {
                            case "a":
                                v59 := []int{42}
                                _ = v59
                            default:
                                if (!((100 <= 127) && (v2 && (v3 >= v3)))) {
                                    var v60 int64 = 42
                                    v1[2] = 42
                                    v61 := GenS2{F0: 180, F1: 114, F2: 1000}
                                    _ = v60
                                    _ = v61
                                }
                            }
                        default:
                            switch "\n\t" {
                            case "a":
                                var v62 uint16 = v3
                                _ = v62
                            default:
                                delete(v1, 381)
                            }
                        }
                        v2 = (v3 > v3)
                        _ = v48
                    }
                    _ = fi33
                    _ = v40
                }
            }
            var v63 float64 = (1e10 - float64(3.14159))
            switch "foo bar" {
            case "":
                if v2 {
                    v2 = (!(!v2))
                }
            case "hello":
                v64 := v4[("foo bar" + (("hello" + "\n\t") + ("foo bar" + "\n\t")))]
                _ = v64
            case "a":
                fmt.Println(v31, v31, v3)
            default:
                v31 = v31
            }
            _ = v63
        }
        _ = fi32
    }
    var v65 int32 = -2147483648
    v66 := GenS2{F0: (v65 - 734), F1: 781, F2: 1000}
    v65 = v65
    var v67 int32 = v65
    v67 = -2147483648
    v67 -= (v67 - (v67 & (^(v65))))
    fmt.Println(v3)
    for fi68 := 0; fi68 < 2; fi68++ {
        v69 := GenS4{F0: 0}
        if ((v65 | 0) == (v65 ^ v65)) {
            v70 := GenS3{F0: 65535, F1: 0, F2: "foo bar"}
            var v71 string = "\n\t"
            var v72 int8 = 100
            _ = v70
            _ = v71
            _ = v72
        }
        _ = fi68
        _ = v69
    }
    for fi73 := 0; fi73 < 11; fi73++ {
        var v74 string = ("" + "hello")
        switch v74 {
        case "\n\t":
            v75 := GenS5{F0: 255}
            _ = v75
        case "a":
            v3 -= v3
        case "foo bar":
            var v76 int16 = 42
            _ = v76
        default:
            v67 += (v65 * v65)
        }
        v31 -= v31
        _ = fi73
        _ = v74
    }
    fmt.Println(v67, v67)
    switch 285 {
    case 0:
        v3 += v3
    case 32768:
        var v77 string = "hello"
        _ = v77
    default:
        _ = len(v4)
    }
    fmt.Println(v67, v67, v67)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v65
    _ = v66
    _ = v67
}
