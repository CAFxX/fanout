package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 int8; F2 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint; F1 int32; F2 uint64 }
type GenS11 struct { F0 uint; F1 uint; F2 int }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8; F1 int32; F2 int8 }
type GenS14 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS15 struct { F0 uint64 }
func genFn0(p0 string) uint64 {
    if false {
        fmt.Println(p0)
        fmt.Println(p0)
    }
    v5 := GenS1{F0: 299, F1: "", F2: 1}
    var v6 uint = 1
    fmt.Println(v6, p0)
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32) int32 {
    if true {
        var v7 int64 = 9223372036854775807
        p0 = p0
        p0 = (p0 - 2147483647)
        _ = v7
    } else {
        v8 := GenS2{F0: -1, F1: "a"}
        switch "a" {
        case "\n\t":
            fmt.Println(p0)
        }
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS3{F0: 0}
        _ = fi9
        _ = v10
    }
    _ = p0
    return 1
}

func main() {
    var v1 float64 = 0.1
    var v2 bool = true
    v3 := []string{"hello", "hello", ""}
    var v4 bool = v2
    var v11 uint64 = 18446744073709551615
    switch "" {
    case "\n\t":
        var v12 string = "a"
        _ = v12
    case "":
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS4{F0: 4294967295}
            var v15 int8 = 42
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(v15, fi16)
                _ = len(v3)
                switch (-(fi16)) {
                case -490:
                    v17 := GenS5{F0: (^(v11))}
                    _ = v17
                case 32768:
                    switch "hello" {
                    case "a":
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    case "\n\t":
                        switch (("hello" + ("" + "a")) + "\n\t") {
                        case "":
                            _ = len(v3)
                        case "foo bar":
                            if ((1 < 0) && (((v11 % (((v11 + v11)) | 1)) <= genFn0("a")) && ((v15 - v15) != 3))) {
                                switch fi16 {
                                case -9223372036854775808:
                                    v15 = -107
                                }
                                v2 = v2
                            } else {
                                v15 = (v15 + (v15 << ((65535) & 63)))
                                _ = len(v3)
                                if false {
                                    _ = len(v3)
                                }
                                v2 = ((!v4) && false)
                                if false {
                                    fmt.Println(v15)
                                    v18 := GenS6{F0: genFn0(("" + "\n\t"))}
                                    _ = v18
                                } else {
                                    if (42 < -87) {
                                        if len(v3) > 0 {
                                            v3[0] = ("" + "hello")
                                        }
                                        switch 0 {
                                        case 9223372036854775807:
                                            v3 = append(v3, "foo bar")
                                        }
                                        v19 := GenS7{F0: 0, F1: (-(v15)), F2: (-(v15))}
                                        var v20 uint = 613
                                        v21 := GenS1{F0: fi16, F1: "hello", F2: v15}
                                        v22 := GenS8{F0: 127}
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        if ((!v4) && (v1 >= v1)) {
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                v1 += v1
                                            case "foo bar":
                                                var v23 uint32 = 0
                                                _ = v23
                                            case "\n\t":
                                                if false {
                                                    v24 := GenS3{F0: 200}
                                                    var v25 uint = 18446744073709551615
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    switch 4294967295 {
                                                    case -409:
                                                        v2 = (1 >= 65535)
                                                    case -9223372036854775808:
                                                        v11 -= (^(v11))
                                                    case -414:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                    fmt.Println(v2, v1)
                                                }
                                                switch (fi13 << ((255) & 63)) {
                                                case 1:
                                                    fmt.Println(v11, v11)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v26 int32 = 1
                                                    _ = v26
                                                default:
                                                    if false {
                                                        v27 := GenS9{F0: 2}
                                                        _ = v27
                                                    }
                                                }
                                                v1 -= ((0.0 - (v1 - float64((((1.5 + ((0.0 + -2.25) + ((3.14159 / (float32(-2.25) - float32(v1))) + float32(float64(float32((1e10 + (v1 * (1.5 - 1.5))))))))) / 1.0) + (1.5 * ((3.14159 * -2.25) * -2.25)))))) + ((float64(1e10) * (v1 - v1)) / (v1 / float64(3.14159))))
                                            default:
                                                v1 -= (v1 + (((0.1 + v1) - v1) * 1.5))
                                            }
                                            var v28 string = ("" + (("\n\t" + ("hello" + "\n\t")) + ""))
                                            switch fi13 {
                                            case 128:
                                                v29 := GenS5{F0: 2}
                                                _ = v29
                                            case -1:
                                                if v2 {
                                                    _ = len(v3)
                                                    var v30 uint32 = 1
                                                    if v2 {
                                                        var v31 uint64 = (v11 * (v11 >> ((18446744073709551615) & 63)))
                                                        var v32 uint8 = 11
                                                        v33 := GenS1{F0: (fi16 / (((^(fi16))) | 1)), F1: (v28 + v28), F2: 100}
                                                        _ = v31
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    v11 = genFn0((((v28 + ((v28 + "\n\t") + v28)) + v28) + (v28 + (v28 + (((v28 + "\n\t") + v28) + ("" + "hello"))))))
                                                    _ = v30
                                                } else {
                                                    v34 := GenS2{F0: -1, F1: "a"}
                                                    var v35 int8 = -1
                                                    _ = v34
                                                    _ = v35
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            _ = v28
                                        } else {
                                            v2 = (v11 != genFn0("foo bar"))
                                            v15 = (v15 & v15)
                                            v2 = v4
                                        }
                                        switch 4294967295 {
                                        case -2147483648:
                                            v4 = (!(945 != -25))
                                        case 127:
                                            var v36 uint16 = 65535
                                            _ = v36
                                        case 42:
                                            v4 = (128 <= 128)
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v4 = v4
                                    v37 := GenS6{F0: 2}
                                    _ = v37
                                }
                                var v38 uint8 = 255
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                                v1 -= 3.14159
                                if false {
                                    switch int(v11) {
                                    case 4294967295:
                                        v39 := GenS1{F0: 127, F1: "a", F2: (v15 | (v15 << ((171) & 63)))}
                                        _ = v39
                                    case 261:
                                        var v40 float64 = v1
                                        _ = v40
                                    case 1:
                                        var v41 int64 = 605
                                        _ = v41
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "":
                                        v3 = append(v3, "a")
                                    case "foo bar":
                                        v15 -= 103
                                    default:
                                        v1 += ((((0.0 * float64(0.0)) - (((0.0 / (v1 / 0.1)) - float64(0.0)) * (3.14159 - v1))) + v1) - ((v1 / v1) + (1.5 + -2.25)))
                                    }
                                } else {
                                    v4 = (v4 && (!false))
                                    v4 = v4
                                    v38 += 198
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v42 bool = v4
                                v1 -= v1
                                _ = v38
                                _ = v42
                            }
                        case "hello":
                            fmt.Println(v4, v15)
                            switch "hello" {
                            case "\n\t":
                                var v43 bool = (42 != 1)
                                _ = v43
                            }
                            switch fi13 {
                            case -482:
                                var v44 int8 = v15
                                _ = v44
                            case 274:
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            }
                        }
                    case "":
                        var v45 float64 = v1
                        _ = v45
                    }
                }
                switch "" {
                case "foo bar":
                    switch fi16 {
                    case 712:
                        var v46 float64 = float64((1.5 * (1e10 - (((1.5 * float32(float64(-2.25))) * 0.1) / 1.0))))
                        _ = v46
                    default:
                        fmt.Println(v11)
                    }
                case "":
                    v15 = 42
                case "a":
                    _ = len(v3)
                }
                switch 32767 {
                case -1:
                    v1 -= v1
                case 127:
                    var v47 int32 = 2147483647
                    v48 := GenS10{F0: 18446744073709551615, F1: (v47 | -652), F2: (v11 / (((^(v11))) | 1))}
                    _ = v47
                    _ = v48
                case 9223372036854775807:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi16
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi49 := 0; fi49 < 3; fi49++ {
            for fi50 := 0; fi50 < 11; fi50++ {
                fmt.Println(v11, fi50, v1)
                var v51 int32 = 980
                if v2 {
                    v2 = true
                    fmt.Println(fi50)
                    if (("hello" + ("hello" + (("" + ((("\n\t" + "foo bar") + ("hello" + "")) + ("\n\t" + "a"))) + "a"))) != (("" + "hello") + "a")) {
                        fmt.Println(v2)
                        var v52 int8 = -81
                        var v53 int32 = (-(v51))
                        _ = v52
                        _ = v53
                    } else {
                        v3 = append(v3, "a")
                    }
                } else {
                    v54 := GenS11{F0: 1, F1: 748, F2: 1}
                    var v55 float32 = 1e10
                    _ = v54
                    _ = v55
                }
                _ = fi50
                _ = v51
            }
            var v56 uint8 = 127
            var v57 uint8 = 255
            _ = fi49
            _ = v56
            _ = v57
        }
    }
    v4 = (2147483647 == 2147483648)
    for fi58 := 0; fi58 < 6; fi58++ {
        if (!(!v4)) {
            var v59 string = ("\n\t" + "a")
            switch "" {
            case "foo bar":
                v59 = v59
            case "\n\t":
                v11 = v11
            case "a":
                var v60 float32 = (0.0 - (((1e10 * 0.1) / 1.0) / (0.0 - -2.25)))
                _ = v60
            }
            v2 = true
            _ = v59
        } else {
            var v61 int8 = 100
            var v62 uint16 = 2
            _ = v61
            _ = v62
        }
        v63 := GenS7{F0: -62, F1: -128, F2: 119}
        var v64 bool = v2
        _ = fi58
        _ = v63
        _ = v64
    }
    if (false || v4) {
        var v65 uint8 = 128
        v66 := GenS12{F0: (v11 + 123)}
        v1 -= ((v1 / 1.5) + float64(1.5))
        _ = v65
        _ = v66
    }
    fmt.Println(v1, v2)
    for fi67 := 0; fi67 < 19; fi67++ {
        var v68 uint64 = genFn0(("\n\t" + ("hello" + "")))
        switch "\n\t" {
        case "hello":
            for fi69 := 0; fi69 < 11; fi69++ {
                v70 := GenS11{F0: 2, F1: 255, F2: (fi67 - fi67)}
                if true {
                    v4 = ((2147483647 > fi67) && true)
                } else {
                    v11 += genFn0("foo bar")
                    if (!(!true)) {
                        fmt.Println(v1, v4)
                        if true {
                            var v71 uint = 255
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v71
                        } else {
                            v68 += (v11 % ((v68) | 1))
                            v68 = genFn0(("\n\t" + ""))
                        }
                    } else {
                        _ = len(v3)
                    }
                }
                var v72 uint = 65535
                _ = fi69
                _ = v70
                _ = v72
            }
        case "a":
            fmt.Println(v1)
        default:
            var v73 uint64 = genFn0(("a" + ("\n\t" + (("a" + "\n\t") + ""))))
            _ = v73
        }
        v74 := GenS5{F0: v68}
        _ = fi67
        _ = v68
        _ = v74
    }
    v3 = append(v3, (("\n\t" + "") + ("" + "\n\t")))
    switch -358 {
    case 128:
        var v75 int = 42
        _ = v75
    }
    switch ("a" + ("hello" + ("\n\t" + (((("foo bar" + "") + ("foo bar" + ("" + "a"))) + "a") + "hello")))) {
    case "foo bar":
        fmt.Println(v11)
    case "":
        v76 := GenS13{F0: -128, F1: 1, F2: 0}
        _ = v76
    case "a":
        for fi77 := 0; fi77 < 5; fi77++ {
            var v78 int64 = 366
            if ((genFn1(2) < 1) || v4) {
                if false {
                    var v79 uint16 = 2
                    for fi80 := 0; fi80 < 10; fi80++ {
                        var v81 int16 = -1
                        var v82 string = "foo bar"
                        fmt.Println(v1, v2, v79)
                        var v83 float64 = (0.1 - (v1 + v1))
                        if false {
                            if v4 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v3) > 0 {
                                    v3[0] = (v82 + ((v82 + v82) + "foo bar"))
                                }
                                if (true || false) {
                                    fmt.Println(fi77, v2, v1)
                                } else {
                                    if v4 {
                                        v83 -= 0.0
                                        v84 := GenS14{F0: 0, F1: v79, F2: 1}
                                        fmt.Println(fi77, v2, v78)
                                        _ = v84
                                    } else {
                                        v2 = (((!(v4 && (!v4))) || ((4294967295 == 65535) && v4)) && false)
                                        var v85 int8 = 42
                                        _ = v85
                                    }
                                    if ((v79 % (((^(v79))) | 1)) >= 641) {
                                        _ = len(v3)
                                        switch (((("\n\t" + v82) + (v82 + "\n\t")) + v82) + ((v82 + "hello") + v82)) {
                                        case "a":
                                            v3 = append(v3, ("foo bar" + (("" + v82) + (v82 + v82))))
                                        case "\n\t":
                                            var v86 uint32 = 2
                                            _ = v86
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v87 uint64 = v11
                                        _ = v87
                                    }
                                }
                                fmt.Println(v83)
                                v4 = v2
                            } else {
                                v79 -= (v79 - v79)
                                fmt.Println(fi80)
                                v82 = v82
                            }
                        } else {
                            v2 = true
                        }
                        _ = fi80
                        _ = v81
                        _ = v82
                        _ = v83
                    }
                    _ = v79
                } else {
                    v11 += (v11 | (v11 + genFn0((((("a" + ("foo bar" + "\n\t")) + "\n\t") + "\n\t") + "\n\t"))))
                }
                if false {
                    for fi88 := 0; fi88 < 12; fi88++ {
                        v1 += v1
                        fmt.Println(v2)
                        _ = fi88
                    }
                    var v89 int16 = 870
                    _ = v89
                }
            }
            var v90 int = (fi77 - fi77)
            _ = fi77
            _ = v78
            _ = v90
        }
    }
    v11 -= (v11 >> ((0) & 63))
    _ = len(v3)
    for fi91 := 0; fi91 < 12; fi91++ {
        fmt.Println(fi91)
        fmt.Println(fi91, v1)
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi91
    }
    v92 := GenS15{F0: v11}
    if (-9223372036854775808 == 42) {
        if (!v2) {
            var v93 uint16 = 2
            v4 = false
            var v94 uint = 915
            _ = v93
            _ = v94
        } else {
            var v95 int = -396
            if v2 {
                if (v95 > 2147483647) {
                    switch "\n\t" {
                    case "a":
                        v96 := GenS4{F0: 687}
                        _ = v96
                    case "\n\t":
                        v95 += int(v11)
                    default:
                        if true {
                            var v97 uint16 = 514
                            var v98 string = ("\n\t" + "foo bar")
                            if (255 == 56) {
                                fmt.Println(v1, v4)
                                v3 = append(v3, (((v98 + (v98 + (v98 + v98))) + (v98 + "hello")) + v98))
                                var v99 int8 = -1
                                _ = v99
                            }
                            _ = v97
                            _ = v98
                        }
                    }
                }
            }
            _ = v95
        }
        switch -9223372036854775808 {
        case 128:
            v11 += v11
        case 42:
            v3 = append(v3, "foo bar")
        default:
            var v100 float64 = (0.0 / v1)
            _ = v100
        }
        var v101 int = 127
        _ = v101
    }
    if true {
        _ = len(v3)
        v3 = append(v3, ((("foo bar" + "foo bar") + "\n\t") + ("a" + ("" + "foo bar"))))
    } else {
        fmt.Println(v4, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v92
}
