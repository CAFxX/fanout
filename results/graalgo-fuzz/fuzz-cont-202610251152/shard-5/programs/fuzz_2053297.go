package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 bool; F2 int32 }
type GenS9 struct { F0 string; F1 bool; F2 int32 }
type GenS10 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint32) int {
    p0 = (p0 * p0)
    var v6 int = -946
    _ = p0
    _ = v6
    return (v6 >> ((562) & 63))
}
func genFn1() bool {
    var v7 bool = (true || (18446744073709551615 < 2))
    var v8 uint16 = 42
    if (255 != 4294967295) {
        v8 = 761
        v8 -= v8
        v9 := GenS1{F0: 0, F1: 23}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            switch (fi10 / (((^(fi10))) | 1)) {
            case 9223372036854775807:
                var v11 int8 = 40
                _ = v11
            default:
                v12 := GenS1{F0: -1, F1: 9}
                _ = v12
            }
            fmt.Println(v7)
            v8 = (v8 / (((v8 * v8)) | 1))
            _ = fi10
        }
        var v13 uint = 688
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS1{F0: 195, F1: uint8(v8)}
            for fi16 := 0; fi16 < 7; fi16++ {
                v7 = (!v7)
                if (!(9223372036854775807 <= -9223372036854775808)) {
                    v17 := []int{}
                    if (!(((v7 || v7) && v7) && ((int64(fi14) == -1) && v7))) {
                        if len(v17) > 0 {
                            v17[0] = (fi14 & (fi16 | fi16))
                        }
                    }
                    _ = v17
                } else {
                    fmt.Println(v8)
                    v7 = v7
                }
                var v18 int16 = 42
                _ = fi16
                _ = v18
            }
            if v7 {
                var v19 bool = true
                var v20 int = 0
                _ = v19
                _ = v20
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v8 -= (^(v8))
    _ = v7
    _ = v8
    return (0 != 0)
}

func main() {
    var v1 int32 = -397
    var v2 float64 = 0.0
    var v3 bool = (v2 != v2)
    v4 := make(map[string]bool)
    var v5 float32 = (((float32(v2) / -2.25) + (1.5 + -2.25)) + 0.0)
    if v3 {
        v21 := []int{}
        switch 0 {
        case 74:
            var v22 uint32 = 0
            _ = v22
        case 42:
            var v23 float64 = ((v2 / v2) + v2)
            _ = v23
        default:
            v24 := GenS2{F0: 1000, F1: 128}
            _ = v24
        }
        _ = v21
    }
    v25 := GenS3{F0: 18446744073709551615, F1: 127, F2: 1}
    v4[("" + "a")] = (v3 || v3)
    v26 := GenS3{F0: 4294967295, F1: 1, F2: 1}
    v27 := GenS4{F0: v1, F1: 629}
    v28 := v4[("foo bar" + ("" + "foo bar"))]
    if (v1 < 42) {
        switch (("\n\t" + ("foo bar" + "foo bar")) + "a") {
        case "a":
            switch "hello" {
            case "a":
                var v29 bool = (!(42 > 42))
                _ = v29
            case "foo bar":
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := []int{genFn0(4294967295), (^(fi30)), (fi30 + (fi30 / ((0) | 1))), fi30}
                    _ = fi30
                    _ = v31
                }
            case "":
                v32 := GenS4{F0: 910, F1: 2147483648}
                _ = v32
            }
        case "":
            if (-128 != -92) {
                v5 = v5
                var v33 int64 = 9223372036854775807
                fmt.Println(v2, v28, v2)
                _ = v33
            } else {
                v34 := GenS2{F0: 32767, F1: 255}
                var v35 int8 = 0
                var v36 bool = v3
                _ = v34
                _ = v35
                _ = v36
            }
        case "\n\t":
            for fi37 := 0; fi37 < 6; fi37++ {
                if (!v3) {
                    var v38 uint8 = 255
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 15; fi39++ {
                        var v40 bool = true
                        _ = fi39
                        _ = v40
                    }
                }
                if true {
                    v4["hello"] = v3
                    for fi41 := 0; fi41 < 7; fi41++ {
                        v42 := v4["foo bar"]
                        var v43 int32 = 932
                        if (54 >= 2) {
                            switch genFn0(2) {
                            case -2147483648:
                                v3 = ((v1 == v1) || (9223372036854775807 > -1))
                            case 4294967295:
                                switch (fi37 + 128) {
                                case 9223372036854775807:
                                    v42 = false
                                default:
                                    v44 := GenS5{F0: -255, F1: ("hello" + "foo bar")}
                                    _ = v44
                                }
                            case 9223372036854775807:
                                var v45 uint64 = 4294967295
                                switch fi41 {
                                case 618:
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            if ((("\n\t" + "") + ("\n\t" + (("hello" + "a") + ("foo bar" + ("a" + "\n\t"))))) != "") {
                                                switch ((("a" + "\n\t") + "\n\t") + "") {
                                                case "\n\t":
                                                    var v46 int64 = 42
                                                    _ = v46
                                                case "":
                                                    v2 = ((v2 / v2) / 1.0)
                                                case "hello":
                                                    v47 := GenS6{F0: 255, F1: 4294967295, F2: 32767}
                                                    _ = v47
                                                }
                                            } else {
                                                v48 := GenS1{F0: (v1 ^ v1), F1: 255}
                                                _ = v48
                                            }
                                        case "":
                                            var v49 float64 = (float64(1.5) + (v2 / v2))
                                            _ = v49
                                        }
                                    } else {
                                        switch fi41 {
                                        case -522:
                                            _ = len(v4)
                                        case 2147483647:
                                            switch (fi41 + fi41) {
                                            case 42:
                                                var v50 uint = 1
                                                _ = v50
                                            case -775:
                                                if (v3 && (false || (42 == 127))) {
                                                    v43 += v43
                                                    v51 := GenS1{F0: (v1 % ((789) | 1)), F1: 113}
                                                    var v52 bool = (!(true && false))
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    v5 -= (v5 / 1e10)
                                                    var v53 uint64 = 2
                                                    v54 := GenS2{F0: -295, F1: 1}
                                                    var v55 int16 = 1000
                                                    v3 = false
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                            case -9223372036854775808:
                                                var v56 uint8 = 42
                                                _ = v56
                                            default:
                                                v43 += 2147483647
                                                if ((v43 * v1) != (^(v1))) {
                                                    if (127 > -128) {
                                                        v57 := []int{(^(fi41)), fi41, (fi37 * 127)}
                                                        var v58 uint16 = 32767
                                                        switch "" {
                                                        case "\n\t":
                                                            v28 = v28
                                                        default:
                                                            delete(v4, ("" + "a"))
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                }
                                                if (!v3) {
                                                    v59 := GenS7{F0: 255}
                                                    v60 := v4["hello"]
                                                    var v61 bool = false
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                        }
                                    }
                                    var v62 uint16 = 32767
                                    _ = v62
                                case 2147483647:
                                    if false {
                                        var v63 int32 = -713
                                        v2 += (v2 + 0.0)
                                        _ = v63
                                    }
                                    v64 := GenS3{F0: (^(v45)), F1: 0, F2: 42}
                                    _ = v64
                                }
                                _ = v45
                            default:
                                var v65 uint = 42
                                _ = v65
                            }
                            v66 := GenS2{F0: 32767, F1: 42}
                            switch "foo bar" {
                            case "foo bar":
                                var v67 bool = ((32767 != 42) && v28)
                                _ = v67
                            }
                            _ = v66
                        } else {
                            switch fi37 {
                            case 0:
                                switch "hello" {
                                case "foo bar":
                                    v68 := GenS8{F0: 1, F1: (2147483648 > 65535), F2: (^(v43))}
                                    _ = v68
                                case "hello":
                                    v28 = ("\n\t" <= (("" + ((("hello" + "foo bar") + "") + "foo bar")) + "a"))
                                default:
                                    var v69 float64 = (-2.25 * (1e10 - v2))
                                    _ = v69
                                }
                            case 32767:
                                fmt.Println(v2, v28)
                            case -866:
                                var v70 float32 = v5
                                _ = v70
                            }
                            switch "a" {
                            case "":
                                switch fi41 {
                                case 128:
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        v71 := GenS8{F0: 181, F1: (0 != 428), F2: (v43 << ((441) & 63))}
                                        _ = v71
                                    }
                                }
                            case "hello":
                                v3 = v28
                            case "a":
                                switch fi37 {
                                case 0:
                                    if v28 {
                                        v5 += 0.0
                                        var v72 uint16 = 432
                                        var v73 uint8 = 255
                                        _ = v72
                                        _ = v73
                                    }
                                case 9223372036854775807:
                                    fmt.Println(v2)
                                case 32768:
                                    if ((!v42) || (true && (!(v2 != float64(float32(((v2 * v2) / v2))))))) {
                                        fmt.Println(v2)
                                        var v74 int32 = (v1 * (v1 % ((v1) | 1)))
                                        v75 := GenS3{F0: 18446744073709551615, F1: 255, F2: 65535}
                                        _ = v74
                                        _ = v75
                                    }
                                default:
                                    if v3 {
                                        fmt.Println(fi37, v3)
                                        if (2 >= 4294967295) {
                                            fmt.Println(v28, v28)
                                            if v42 {
                                                v43 -= (-(v1))
                                                v2 += 1.5
                                                fmt.Println(v28)
                                                if true {
                                                    v76 := GenS7{F0: 65535}
                                                    switch "a" {
                                                    case "foo bar":
                                                        if v42 {
                                                            v1 = v43
                                                            v77 := GenS7{F0: 974}
                                                            var v78 int8 = 0
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            v5 = v5
                                                            var v79 uint = 315
                                                            var v80 int32 = 0
                                                            v81 := GenS9{F0: "a", F1: v3, F2: v1}
                                                            v82 := GenS5{F0: -9223372036854775808, F1: "a"}
                                                            v83 := []int{(fi37 & -187), fi37}
                                                            fmt.Println(v2)
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                    case "hello":
                                                        fmt.Println(v1)
                                                        fmt.Println(v43)
                                                        var v84 uint = 42
                                                        _ = len(v4)
                                                        _ = v84
                                                    }
                                                    _ = v76
                                                }
                                            } else {
                                                fmt.Println(v1, v2, fi41)
                                            }
                                            if (42 >= 747) {
                                                var v85 uint8 = 2
                                                fmt.Println(v2, v2)
                                                _ = v85
                                            }
                                        } else {
                                            v86 := v4[("hello" + "foo bar")]
                                            _ = v86
                                        }
                                        var v87 int8 = 42
                                        _ = v87
                                    }
                                }
                            }
                            v88 := []int{(-(fi41))}
                            _ = v88
                        }
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                } else {
                    switch genFn0(11) {
                    case -9223372036854775808:
                        for fi89 := 0; fi89 < 11; fi89++ {
                            var v90 uint32 = 65535
                            v91 := []int{(fi89 % ((fi37) | 1)), (fi37 >> ((266) & 63)), fi37}
                            v90 -= v90
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                    default:
                        v92 := GenS4{F0: v1, F1: 341}
                        _ = v92
                    }
                    fmt.Println(v28, v5, v3)
                }
                _ = fi37
            }
        }
    } else {
        for fi93 := 0; fi93 < 11; fi93++ {
            for fi94 := 0; fi94 < 11; fi94++ {
                v95 := []int{fi94, (fi93 >> ((65535) & 63)), fi94, (^(fi93))}
                _ = fi94
                _ = v95
            }
            _ = fi93
        }
        fmt.Println(v28)
        var v96 uint = 4294967295
        _ = v96
    }
    if (0 >= 476) {
        if v28 {
            v3 = (v3 || (!(119 == 18446744073709551615)))
            v97 := GenS6{F0: 734, F1: 2, F2: 566}
            _ = v97
        }
        v98 := GenS10{F0: 948, F1: 4294967295}
        v28 = v3
        _ = v98
    }
    v28 = true
    fmt.Println(v1, v2)
    v99 := []int{170, -2147483648, genFn0(2147483648)}
    fmt.Println(v5, v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
