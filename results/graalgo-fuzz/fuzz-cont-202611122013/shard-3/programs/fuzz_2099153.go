package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 int }
type GenS4 struct { F0 uint; F1 uint; F2 uint32 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 int8; F1 int32; F2 int }
type GenS10 struct { F0 uint64; F1 uint }
func genFn0() bool {
    var v4 string = (("" + (("hello" + "foo bar") + ("a" + ("a" + "foo bar")))) + "foo bar")
    fmt.Println(v4)
    v4 = v4
    _ = v4
    return (9223372036854775807 == 1)
}
func genFn1(p0 int64, p1 int16, p2 int) uint8 {
    var v5 float32 = (3.14159 / 3.14159)
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 253
}

func main() {
    var v1 uint = 792
    var v2 int8 = 127
    var v3 bool = true
    v6 := []int{128}
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    switch "foo bar" {
    case "a":
        v6 = append(v6, 2147483647)
    }
    _ = len(v6)
    v2 = 100
    switch 127 {
    case -9223372036854775808:
        v2 -= (v2 * (v2 & int8(v1)))
    case -497:
        v7 := GenS1{F0: 222, F1: -806}
        _ = v7
    case 127:
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: genFn1(42, 669, fi8), F1: 1}
            v6 = append(v6, fi8)
            _ = fi8
            _ = v9
        }
    }
    if v3 {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v10 int16 = 1000
        for fi11 := 0; fi11 < 14; fi11++ {
            v10 += v10
            _ = fi11
        }
        _ = v10
    } else {
        var v12 int8 = 12
        v13 := GenS2{F0: 65535, F1: "a", F2: 4294967295}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        _ = len(v6)
        if v3 {
            var v15 bool = (fi14 != (fi14 / (((^(fi14))) | 1)))
            _ = v15
        }
        _ = fi14
    }
    var v16 uint32 = 649
    if (127 <= 255) {
        switch "" {
        case "foo bar":
            var v17 string = "foo bar"
            _ = v17
        default:
            switch -605 {
            case -386:
                if (!((0.0 * 1.5) >= 0.1)) {
                    if v3 {
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if v3 {
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    v20 := GenS2{F0: 42, F1: "foo bar", F2: v16}
                                    if v3 {
                                        v21 := GenS3{F0: (int32(v16) >= -568), F1: fi19, F2: (fi18 + 127)}
                                        var v22 int32 = 0
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                var v23 uint32 = (v16 | v16)
                                var v24 float32 = 3.14159
                                _ = v23
                                _ = v24
                            } else {
                                if v3 {
                                    var v25 int8 = v2
                                    _ = v25
                                } else {
                                    if (true && v3) {
                                        if (18446744073709551615 <= 210) {
                                            for fi26 := 0; fi26 < 15; fi26++ {
                                                if false {
                                                    switch 1 {
                                                    case 42:
                                                        switch "hello" {
                                                        case "":
                                                            v1 -= v1
                                                            var v27 float32 = 0.0
                                                            if (v3 && v3) {
                                                                fmt.Println(v16)
                                                            }
                                                            _ = v27
                                                        default:
                                                            v2 = 1
                                                        }
                                                    case 2147483647:
                                                        if v3 {
                                                            if v3 {
                                                                _ = len(v6)
                                                                v2 = (v2 & v2)
                                                            } else {
                                                                var v28 int8 = 127
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                var v29 int16 = -32768
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        } else {
                                                            if len(v6) > 0 {
                                                                v6[0] = fi26
                                                            }
                                                            v30 := GenS4{F0: v1, F1: 4294967295, F2: (v16 % ((v16) | 1))}
                                                            v1 -= (v1 >> (((^(v1))) & 63))
                                                            _ = v30
                                                        }
                                                    case 128:
                                                        if false {
                                                            if false {
                                                                v6 = append(v6, 194)
                                                                switch (("" + "\n\t") + "hello") {
                                                                case "foo bar":
                                                                    var v31 float64 = 0.0
                                                                    _ = v31
                                                                default:
                                                                    if len(v6) > 0 {
                                                                        v6[0] = fi18
                                                                    }
                                                                }
                                                                v32 := GenS5{F0: 1, F1: 9223372036854775807}
                                                                _ = v32
                                                            }
                                                            if v3 {
                                                                v33 := GenS5{F0: 1000, F1: -9223372036854775808}
                                                                var v34 int64 = 746
                                                                var v35 int32 = 42
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                        }
                                                    }
                                                    if false {
                                                        if false {
                                                            var v36 int64 = 919
                                                            _ = v36
                                                        } else {
                                                            if false {
                                                                v37 := GenS1{F0: 127, F1: -584}
                                                                if true {
                                                                    v38 := GenS2{F0: uint16(v1), F1: ("hello" + "foo bar"), F2: v16}
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    _ = v38
                                                                } else {
                                                                    if ((v2 << (((v1 ^ (v1 >> (((^(v1))) & 63)))) & 63)) <= 42) {
                                                                        fmt.Println(fi18, v16, v16)
                                                                    }
                                                                    switch ("\n\t" + ("hello" + ("" + ("\n\t" + ("hello" + (("foo bar" + "\n\t") + "")))))) {
                                                                    case "foo bar":
                                                                        var v39 string = ((("" + "\n\t") + "a") + "foo bar")
                                                                        _ = v39
                                                                    default:
                                                                        v40 := GenS3{F0: (!v3), F1: fi26, F2: -1}
                                                                        _ = v40
                                                                    }
                                                                    var v41 string = ("\n\t" + (("hello" + "") + "a"))
                                                                    _ = v41
                                                                }
                                                                if (-2.25 == 0.1) {
                                                                    if false {
                                                                        if (0.0 > (3.14159 * ((float32((-2.25 - float64(float32(0.1)))) * 3.14159) * 1.5))) {
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                        }
                                                                        var v42 uint64 = 0
                                                                        v43 := GenS2{F0: 65535, F1: "\n\t", F2: (v16 % ((v16) | 1))}
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v6 = append(v6, (-(fi18)))
                                                                        case "hello":
                                                                            if (!((1 < uint16(v16)) || (203 > 0))) {
                                                                                var v44 int64 = int64(v2)
                                                                                _ = v44
                                                                            }
                                                                            switch "hello" {
                                                                            case "":
                                                                                v42 = (v42 & 347)
                                                                            default:
                                                                                switch fi26 {
                                                                                case 127:
                                                                                    v42 += 1
                                                                                case -882:
                                                                                    var v45 int16 = 1
                                                                                    _ = v45
                                                                                }
                                                                            }
                                                                        }
                                                                        var v46 int64 = -1
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v46
                                                                    }
                                                                    var v47 int32 = -1
                                                                    _ = v47
                                                                } else {
                                                                    if (1 >= 2) {
                                                                        if (fi18 == 128) {
                                                                            var v48 uint8 = 0
                                                                            v3 = v3
                                                                            var v49 uint16 = 225
                                                                            v50 := GenS3{F0: v3, F1: fi26, F2: 128}
                                                                            _ = v48
                                                                            _ = v49
                                                                            _ = v50
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                var v51 int32 = -1
                                                                                _ = v51
                                                                            case "foo bar":
                                                                                v1 = v1
                                                                            default:
                                                                                v2 = v2
                                                                            }
                                                                            v1 = v1
                                                                            v3 = ((v16 / (((v16 | (v16 % (((v16 >> ((18446744073709551615) & 63))) | 1)))) | 1)) == v16)
                                                                        }
                                                                        var v52 int = (fi18 & fi18)
                                                                        if v3 {
                                                                            v53 := GenS6{F0: 0}
                                                                            if true {
                                                                                fmt.Println(v52, fi26, v16)
                                                                            } else {
                                                                                var v54 bool = (v1 == v1)
                                                                                fmt.Println(fi18)
                                                                                if v54 {
                                                                                    var v55 float64 = -2.25
                                                                                    var v56 int64 = 0
                                                                                    _ = v55
                                                                                    _ = v56
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        v57 := GenS1{F0: 42, F1: 0}
                                                                                        _ = v57
                                                                                    case "foo bar":
                                                                                        if len(v6) > 0 {
                                                                                            v6[0] = (v52 >> ((42) & 63))
                                                                                        }
                                                                                    }
                                                                                    var v58 int16 = -1
                                                                                    v6 = append(v6, fi26)
                                                                                    _ = v58
                                                                                }
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                                _ = v54
                                                                            }
                                                                            switch (fi18 % (((fi26 * -9223372036854775808)) | 1)) {
                                                                            case -622:
                                                                                v59 := GenS7{F0: (v52 * fi26), F1: (-(v52))}
                                                                                _ = v59
                                                                            default:
                                                                                v60 := GenS8{F0: -1, F1: 9223372036854775807}
                                                                                _ = v60
                                                                            }
                                                                            _ = v53
                                                                        } else {
                                                                            if len(v6) > 0 {
                                                                                v6[0] = v52
                                                                            }
                                                                            v61 := GenS2{F0: 42, F1: "", F2: (v16 ^ (v16 + 2))}
                                                                            v52 = (fi18 << ((v1) & 63))
                                                                            _ = v61
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                    v3 = (v16 > v16)
                                                                }
                                                                var v62 string = ("a" + "\n\t")
                                                                if (65535 < 1) {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        if v3 {
                                                                            _ = len(v6)
                                                                        } else {
                                                                            fmt.Println(v2, v3)
                                                                        }
                                                                    case "foo bar":
                                                                        switch (fi18 - fi26) {
                                                                        case 0:
                                                                            var v63 uint = 249
                                                                            _ = v63
                                                                        }
                                                                    default:
                                                                        v3 = true
                                                                    }
                                                                    v62 = ((v62 + (v62 + "foo bar")) + ((v62 + ("a" + ((v62 + v62) + v62))) + v62))
                                                                }
                                                                v64 := GenS6{F0: 2}
                                                                var v65 int8 = v2
                                                                if (v3 || v3) {
                                                                    switch (v62 + v62) {
                                                                    case "":
                                                                        if ((!v3) || true) {
                                                                            if true {
                                                                                var v66 uint = (^(v1))
                                                                                switch (fi26 ^ -195) {
                                                                                case 502:
                                                                                    v62 = v62
                                                                                    if v3 {
                                                                                        v67 := GenS3{F0: v3, F1: (fi18 + (fi26 / ((-346) | 1))), F2: fi26}
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = len(v6)
                                                                                        _ = v67
                                                                                    }
                                                                                case 220:
                                                                                    switch "" {
                                                                                    case "a":
                                                                                        var v68 uint32 = (v16 / (((v16 & v16)) | 1))
                                                                                        _ = v68
                                                                                    case "":
                                                                                        var v69 float32 = float32(-2.25)
                                                                                        _ = v69
                                                                                    case "hello":
                                                                                        var v70 int64 = -872
                                                                                        var v71 string = (v62 + v62)
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    }
                                                                                default:
                                                                                    var v72 int = fi18
                                                                                    _ = v72
                                                                                }
                                                                                _ = v66
                                                                            }
                                                                        } else {
                                                                            v73 := GenS6{F0: 863}
                                                                            _ = v73
                                                                        }
                                                                    case "a":
                                                                        v74 := GenS2{F0: 32767, F1: (((v62 + v62) + v62) + (v62 + "")), F2: 65535}
                                                                        _ = v74
                                                                    default:
                                                                        var v75 int8 = 22
                                                                        _ = v75
                                                                    }
                                                                    var v76 bool = false
                                                                    _ = v76
                                                                } else {
                                                                    fmt.Println(v62, v16, fi26)
                                                                    fmt.Println(fi26, v62)
                                                                }
                                                                _ = v37
                                                                _ = v62
                                                                _ = v64
                                                                _ = v65
                                                            } else {
                                                                var v77 uint16 = 255
                                                                _ = v77
                                                            }
                                                        }
                                                        fmt.Println(v2, v3)
                                                        fmt.Println(v2)
                                                    } else {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                        v78 := GenS1{F0: 42, F1: 32767}
                                                        if v3 {
                                                            if v3 {
                                                                v6 = append(v6, fi26)
                                                            }
                                                        }
                                                        _ = v78
                                                    }
                                                    switch (fi26 % ((fi26) | 1)) {
                                                    case 2147483647:
                                                        v79 := GenS4{F0: 4294967295, F1: v1, F2: 2147483648}
                                                        _ = v79
                                                    }
                                                } else {
                                                    if (42 >= 32767) {
                                                        fmt.Println(v1, v3)
                                                        var v80 uint32 = (v16 >> (((v1 * (v1 % ((v1) | 1)))) & 63))
                                                        var v81 uint8 = 112
                                                        var v82 float32 = (-2.25 / -2.25)
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                    } else {
                                                        fmt.Println(v3, v1, v3)
                                                    }
                                                }
                                                if (26 >= 207) {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    fmt.Println(v1)
                                                    var v83 int8 = v2
                                                    v3 = (!false)
                                                    _ = v83
                                                } else {
                                                    v1 -= (^(v1))
                                                    v1 += v1
                                                }
                                                v16 = 42
                                                v2 -= (v2 + (v2 / ((v2) | 1)))
                                                var v84 int = 494
                                                switch ("" + ("" + ("\n\t" + ("a" + "a")))) {
                                                case "a":
                                                    if ("a" > ("" + "hello")) {
                                                        fmt.Println(fi18)
                                                        var v85 uint8 = 0
                                                        if v3 {
                                                            var v86 uint = v1
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = v86
                                                        } else {
                                                            v85 = (v85 | genFn1(9223372036854775807, 32767, (fi26 + 42)))
                                                            fmt.Println(v85, v2, v16)
                                                            if (!v3) {
                                                                var v87 uint64 = uint64(v1)
                                                                _ = v87
                                                            }
                                                        }
                                                        _ = v85
                                                    }
                                                case "":
                                                    if (!(4294967295 == fi26)) {
                                                        _ = len(v6)
                                                    }
                                                }
                                                _ = fi26
                                                _ = v84
                                            }
                                        }
                                        var v88 int8 = v2
                                        _ = len(v6)
                                        _ = v88
                                    }
                                    var v89 int32 = -2147483648
                                    for fi90 := 0; fi90 < 7; fi90++ {
                                        if (-1 < 42) {
                                            fmt.Println(fi18, v89, fi90)
                                            var v91 bool = v3
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v89 -= v89
                                                case "":
                                                    var v92 int64 = 9223372036854775807
                                                    _ = v92
                                                default:
                                                    v16 += 65535
                                                }
                                            case "a":
                                                v93 := GenS9{F0: v2, F1: (v89 & v89), F2: fi18}
                                                _ = v93
                                            case "\n\t":
                                                v16 = (^(v16))
                                            }
                                            v94 := GenS10{F0: 983, F1: (v1 % (((v1 * (v1 - v1))) | 1))}
                                            _ = v91
                                            _ = v94
                                        } else {
                                            if (1 == -170) {
                                                _ = len(v6)
                                                if ((v2 << ((v1) & 63)) == 127) {
                                                    v95 := GenS1{F0: 145, F1: 953}
                                                    _ = v95
                                                } else {
                                                    var v96 string = "\n\t"
                                                    _ = v96
                                                }
                                            } else {
                                                v89 = (v89 ^ v89)
                                            }
                                            var v97 uint8 = 1
                                            _ = v97
                                        }
                                        _ = fi90
                                    }
                                    _ = v89
                                }
                            }
                            var v98 int64 = -665
                            _ = fi18
                            _ = v98
                        }
                        v99 := GenS3{F0: false, F1: 909, F2: 4294967295}
                        _ = v99
                    } else {
                        v6 = append(v6, 32768)
                    }
                    var v100 float32 = (3.14159 / (3.14159 - (((((-2.25 * (0.0 / (float32((3.14159 / 1e10)) / 3.14159))) * (0.0 / float32(3.14159))) * (1e10 / ((1e10 / 1.0) / 1.5))) - float32(0.1)) + (-2.25 * -2.25))))
                    _ = v100
                }
            case 4294967295:
                var v101 bool = (v1 < 18446744073709551615)
                _ = v101
            }
        }
    } else {
        var v102 uint64 = 42
        var v103 int32 = 0
        _ = v102
        _ = v103
    }
    v3 = (!(!((1 >= 1) && false)))
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
