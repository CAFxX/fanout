package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int16; F2 int16 }
type GenS7 struct { F0 uint; F1 uint; F2 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 uint; F1 uint64 }
func genFn0() uint8 {
    var v4 uint64 = 1
    switch "foo bar" {
    case "a":
        v5 := GenS1{F0: -91}
        _ = v5
    default:
        v6 := GenS1{F0: 2147483647}
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 32767
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{}
            if false {
                var v11 uint = 263
                v12 := GenS1{F0: -398}
                _ = v11
                _ = v12
            }
            var v13 float64 = (((float64(float32(0.0)) / 1.5) - ((0.0 + (float64(0.0) + -2.25)) * 3.14159)) / float64(0.1))
            _ = fi9
            _ = v10
            _ = v13
        }
        _ = fi7
        _ = v8
    }
    _ = v4
    return 52
}

func main() {
    var v1 int8 = 1
    var v2 uint = 42
    var v3 int32 = -1
    var v14 bool = false
    for fi15 := 0; fi15 < 7; fi15++ {
        switch -9223372036854775808 {
        case 32768:
            var v16 int8 = (v1 + -111)
            _ = v16
        default:
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := []int{(fi17 & -1), 480}
                _ = fi17
                _ = v18
            }
        }
        _ = fi15
    }
    v2 = (v2 | (v2 ^ (v2 - (v2 & 2))))
    v19 := GenS1{F0: v3}
    var v20 int8 = v1
    v21 := []int{}
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 float64 = float64((3.14159 * 1.5))
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if true {
                v24 := GenS2{F0: 4294967295}
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 string = (("a" + (("foo bar" + "a") + "\n\t")) + "foo bar")
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                var v27 int = fi22
                v28 := GenS3{F0: -32768, F1: v3}
                _ = v27
                _ = v28
            }
        }
        _ = fi22
        _ = v23
    }
    if false {
        if true {
            v14 = v14
            var v29 int = 9223372036854775807
            fmt.Println(v29)
            _ = v29
        } else {
            if len(v21) > 0 {
                v21[0] = 32768
            }
            switch 0 {
            case 32768:
                for fi30 := 0; fi30 < 12; fi30++ {
                    v1 = v1
                    var v31 bool = false
                    fmt.Println(v14, v3, v1)
                    _ = fi30
                    _ = v31
                }
            case -9223372036854775808:
                v1 += (v1 - v1)
            }
            v14 = v14
        }
        var v32 string = ""
        var v33 int16 = 0
        _ = v32
        _ = v33
    } else {
        var v34 int = 1
        _ = v34
    }
    switch 2147483647 {
    case 505:
        if v14 {
            switch 160 {
            case -2147483648:
                var v35 string = "a"
                _ = v35
            case 9223372036854775807:
                v1 = v1
            case 32768:
                for fi36 := 0; fi36 < 8; fi36++ {
                    fmt.Println(v1, v3)
                    if v14 {
                        v20 -= v1
                        for fi37 := 0; fi37 < 7; fi37++ {
                            var v38 string = "a"
                            if false {
                                if (v14 && false) {
                                    v1 = v20
                                }
                            }
                            switch (-(fi37)) {
                            case 32768:
                                var v39 int8 = v1
                                _ = v39
                            case 32767:
                                fmt.Println(fi36, fi37)
                                if v14 {
                                    v40 := GenS3{F0: -1, F1: (-(v3))}
                                    _ = v40
                                } else {
                                    if (1000 < 42) {
                                        fmt.Println(v1, fi37, v2)
                                        if false {
                                            if ((!v14) && v14) {
                                                if len(v21) > 0 {
                                                    v21[0] = 0
                                                }
                                                var v41 uint32 = 0
                                                v42 := GenS2{F0: (-(fi36))}
                                                if (v14 && (!(int32(fi36) != -1))) {
                                                    var v43 string = (v38 + ((("hello" + v38) + "a") + "hello"))
                                                    _ = v43
                                                }
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v44 int16 = 42
                                                if ((0.1 - -2.25) < -2.25) {
                                                    switch fi36 {
                                                    case 944:
                                                        var v45 bool = (!v14)
                                                        _ = v45
                                                    case 1:
                                                        var v46 int64 = 9223372036854775807
                                                        _ = v46
                                                    default:
                                                        v14 = false
                                                    }
                                                    var v47 string = v38
                                                    var v48 string = "foo bar"
                                                    if (!(v14 && v14)) {
                                                        v14 = false
                                                        fmt.Println(v44)
                                                    } else {
                                                        v49 := GenS4{F0: v2}
                                                        v50 := GenS4{F0: (v2 | 65535)}
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    var v51 uint64 = 42
                                                    var v52 int32 = (v3 / ((42) | 1))
                                                    _ = v47
                                                    _ = v48
                                                    _ = v51
                                                    _ = v52
                                                }
                                                if v14 {
                                                    v53 := GenS5{F0: 9223372036854775807}
                                                    v20 = v20
                                                    v21 = append(v21, (fi36 & (fi36 / (((fi36 | -2147483648)) | 1))))
                                                    if (!v14) {
                                                        _ = len(v21)
                                                        fmt.Println(fi36)
                                                    }
                                                    _ = v53
                                                }
                                                _ = v44
                                            }
                                            v21 = append(v21, fi37)
                                            var v54 bool = true
                                            if (v38 <= ("a" + v38)) {
                                                var v55 string = v38
                                                switch v55 {
                                                case "foo bar":
                                                    var v56 uint = 42
                                                    _ = v56
                                                case "\n\t":
                                                    fmt.Println(v1)
                                                default:
                                                    v20 -= v20
                                                }
                                                _ = v55
                                            } else {
                                                var v57 float32 = (float32(float64((3.14159 / 0.1))) * 0.1)
                                                if (int8(v3) < -117) {
                                                    var v58 uint64 = 628
                                                    var v59 int16 = -1
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    if v54 {
                                                        switch v38 {
                                                        case "":
                                                            v57 += (v57 + v57)
                                                        case "foo bar":
                                                            var v60 uint = (v2 * (v2 - v2))
                                                            _ = v60
                                                        default:
                                                            v61 := GenS6{F0: (v20 + v1), F1: 1, F2: -32768}
                                                            _ = v61
                                                        }
                                                    }
                                                }
                                                if true {
                                                    if true {
                                                        v62 := GenS5{F0: 9223372036854775807}
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                        var v63 uint = v2
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                } else {
                                                    if len(v21) > 0 {
                                                        v21[0] = -102
                                                    }
                                                    v3 += -1
                                                    var v64 uint8 = 42
                                                    var v65 float32 = (v57 + (((((v57 + 1e10) - v57) / 3.14159) - v57) * v57))
                                                    v3 = v3
                                                    _ = v64
                                                    _ = v65
                                                }
                                                _ = v57
                                            }
                                            v2 = (v2 % ((v2) | 1))
                                            _ = v54
                                        } else {
                                            fmt.Println(v14)
                                            var v66 uint8 = 128
                                            if (float64(1.5) < (3.14159 * 0.1)) {
                                                var v67 int64 = -1
                                                _ = v67
                                            } else {
                                                if true {
                                                    v68 := GenS6{F0: v20, F1: int16(v3), F2: -956}
                                                    if v14 {
                                                        var v69 uint16 = 9
                                                        _ = v69
                                                    }
                                                    var v70 uint = 0
                                                    _ = v68
                                                    _ = v70
                                                }
                                            }
                                            _ = v66
                                        }
                                    }
                                }
                            }
                            _ = fi37
                            _ = v38
                        }
                    }
                    _ = fi36
                }
            }
            if ((v2 / (((v2 / ((0) | 1))) | 1)) > v2) {
                v1 -= 83
            }
        }
    case 0:
        var v71 string = "hello"
        _ = v71
    case -371:
        for fi72 := 0; fi72 < 9; fi72++ {
            v73 := GenS3{F0: 32767, F1: 1}
            _ = fi72
            _ = v73
        }
    }
    v74 := GenS7{F0: (^(v2)), F1: v2, F2: 462}
    var v75 uint64 = 18446744073709551615
    switch ("\n\t" + ("hello" + "\n\t")) {
    case "foo bar":
        v75 = (v75 % ((251) | 1))
    case "a":
        v1 += v20
    default:
        switch 361 {
        case 9223372036854775807:
            if v14 {
                switch 32767 {
                case 42:
                    v20 = (v20 - (^(v20)))
                default:
                    var v76 int16 = 1
                    _ = v76
                }
            }
        }
    }
    var v77 uint16 = 1
    if v14 {
        v78 := GenS5{F0: 0}
        _ = v78
    } else {
        for fi79 := 0; fi79 < 1; fi79++ {
            var v80 uint32 = 65535
            switch (^(fi79)) {
            case -6:
                fmt.Println(v3, v14, v75)
            case -1:
                for fi81 := 0; fi81 < 8; fi81++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    v82 := GenS8{F0: ""}
                    _ = fi81
                    _ = v82
                }
            default:
                v77 = v77
            }
            for fi83 := 0; fi83 < 15; fi83++ {
                v20 = (v20 + 13)
                if (1.5 >= float64((1.5 - -2.25))) {
                    var v84 int64 = 42
                    _ = v84
                } else {
                    var v85 int8 = (-(v20))
                    _ = v85
                }
                switch "" {
                case "\n\t":
                    v21 = append(v21, -9223372036854775808)
                    fmt.Println(v2)
                case "":
                    var v86 int8 = (v1 - (v20 << ((65535) & 63)))
                    _ = v86
                default:
                    v2 = 9
                }
                _ = fi83
            }
            _ = fi79
            _ = v80
        }
        if v14 {
            fmt.Println(v75, v20, v75)
        } else {
            v87 := GenS5{F0: 9223372036854775807}
            for fi88 := 0; fi88 < 13; fi88++ {
                _ = len(v21)
                v89 := GenS9{F0: 732}
                if len(v21) > 0 {
                    v21[0] = fi88
                }
                _ = fi88
                _ = v89
            }
            _ = v87
        }
    }
    var v90 uint64 = (v75 % ((889) | 1))
    if v14 {
        v75 -= (v75 ^ v75)
        for fi91 := 0; fi91 < 13; fi91++ {
            if (!(196 != genFn0())) {
                fmt.Println(v3)
            } else {
                v92 := GenS10{F0: 0, F1: 1}
                if len(v21) > 0 {
                    v21[0] = 9223372036854775807
                }
                _ = v92
            }
            if v14 {
                fmt.Println(v90)
                var v93 int64 = 9223372036854775807
                if false {
                    var v94 uint16 = 65535
                    _ = v94
                }
                _ = v93
            }
            _ = fi91
        }
        v95 := GenS11{F0: v2, F1: (v75 - (v75 * (v75 << ((2) & 63))))}
        _ = v95
    } else {
        _ = len(v21)
        v90 = v75
    }
    if (160 <= 0) {
        for fi96 := 0; fi96 < 6; fi96++ {
            v14 = (!((^(v77)) != v77))
            _ = fi96
        }
        var v97 float64 = 1e10
        var v98 float64 = (0.1 / ((-2.25 - (float64((-2.25 / (float32(v97) * 1e10))) * 0.0)) * v97))
        _ = v97
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v74
    _ = v75
    _ = v77
    _ = v90
}
