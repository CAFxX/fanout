package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int; F1 int; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool; F1 int32; F2 int16 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 uint8; F1 uint16; F2 bool }
type GenS12 struct { F0 uint }
func genFn0() bool {
    switch 128 {
    case -2147483648:
        var v7 uint32 = 1
        _ = v7
    case -321:
        switch ("a" + "hello") {
        case "\n\t":
            var v8 int32 = 1
            _ = v8
        default:
            if (false || ((0.0 / 1.5) != 1.5)) {
                var v9 int64 = -1
                v10 := []int{902, 1, 32768, 32768}
                _ = v10
                _ = v9
            } else {
                v11 := []int{}
                var v12 int32 = 42
                var v13 float64 = float64((3.14159 * 3.14159))
                _ = v11
                _ = v12
                _ = v13
            }
        }
    }
    if true {
        v14 := []int{-9223372036854775808, 32767, 0, 128}
        _ = v14
    } else {
        if (!false) {
            for fi15 := 0; fi15 < 15; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 float64 = -2.25
                    v17 -= (1e10 + v17)
                    if false {
                        if (42 == 42) {
                            var v18 bool = (0 > 127)
                            _ = v18
                        } else {
                            v19 := []int{fi16, -2147483648, fi16, (fi15 >> ((0) & 63))}
                            if true {
                                var v20 string = "a"
                                if len(v19) > 0 {
                                    v19[0] = fi16
                                }
                                fmt.Println(fi16, v17)
                                _ = v20
                            }
                            v17 -= ((((float64(((1e10 * -2.25) - ((1e10 / (3.14159 / 1.0)) / (0.1 - 0.0)))) + 1e10) / 1.0) - (float64(((1.5 - 1.5) * (0.0 * 0.0))) / -2.25)) * v17)
                            _ = v19
                        }
                        var v21 float64 = v17
                        v21 = v21
                        _ = v21
                    } else {
                        v22 := []int{fi16}
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        if true {
                            if (1e10 == 3.14159) {
                                var v23 uint16 = 1
                                _ = v23
                            } else {
                                var v24 uint64 = 4294967295
                                _ = v24
                            }
                            v17 -= float64(-2.25)
                            if true {
                                if len(v22) > 0 {
                                    v22[0] = 2147483647
                                }
                            } else {
                                var v25 int = fi15
                                _ = len(v22)
                                _ = v25
                            }
                        } else {
                            if len(v22) > 0 {
                                v22[0] = (fi16 + fi15)
                            }
                            fmt.Println(fi15, v17)
                        }
                        switch fi15 {
                        case -9223372036854775808:
                            if len(v22) > 0 {
                                v22[0] = (^(fi15))
                            }
                        case 42:
                            var v26 float64 = v17
                            _ = v26
                        case -1:
                            fmt.Println(fi16, v17)
                        }
                        if (255 != 100) {
                            var v27 uint64 = 18446744073709551615
                            v28 := GenS1{F0: -1, F1: 65535}
                            var v29 int8 = 100
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        var v30 uint32 = 30
                        v17 += float64(1e10)
                        _ = v22
                        _ = v30
                    }
                    if (("foo bar" + "\n\t") <= "a") {
                        v31 := []int{(-(fi16))}
                        v17 -= v17
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        if false {
                            fmt.Println(fi15, v17, fi16)
                            var v32 float32 = 0.0
                            var v33 uint16 = 65535
                            _ = v32
                            _ = v33
                        }
                        _ = v31
                    }
                    fmt.Println(fi15)
                    v17 += v17
                    _ = fi16
                    _ = v17
                }
                fmt.Println(fi15)
                var v34 int16 = 667
                _ = fi15
                _ = v34
            }
            var v35 uint32 = 65535
            v36 := GenS2{F0: -9223372036854775808, F1: 323, F2: (("foo bar" + ("foo bar" + "\n\t")) + "foo bar")}
            _ = v35
            _ = v36
        }
        v37 := GenS3{F0: ("hello" + (("a" + "") + ("" + "hello")))}
        _ = v37
    }
    var v38 int16 = 32767
    _ = v38
    return false
}

func main() {
    var v1 int8 = 63
    v2 := []string{"", "", "hello"}
    var v3 int32 = -569
    var v4 uint = 23
    var v5 bool = true
    var v6 int64 = -404
    for fi39 := 0; fi39 < 4; fi39++ {
        v5 = (32767 > 42)
        for fi40 := 0; fi40 < 1; fi40++ {
            switch fi40 {
            case -9223372036854775808:
                if false {
                    var v41 uint16 = 1
                    var v42 uint32 = 0
                    _ = v41
                    _ = v42
                }
            case 4294967295:
                v5 = (522 == 878)
            default:
                if v5 {
                    v5 = (!v5)
                    if false {
                        v43 := GenS2{F0: 32768, F1: fi40, F2: "a"}
                        _ = len(v2)
                        var v44 int8 = v1
                        _ = v43
                        _ = v44
                    } else {
                        v5 = (v4 != (v4 % (((^(v4))) | 1)))
                        v45 := GenS4{F0: -32768}
                        var v46 uint16 = 255
                        _ = v45
                        _ = v46
                    }
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    v47 := GenS2{F0: (fi39 & 42), F1: 2147483647, F2: "foo bar"}
                    _ = v47
                }
            }
            v6 = v6
            _ = fi40
        }
        v48 := GenS5{F0: (v3 + v3)}
        _ = fi39
        _ = v48
    }
    v3 = (v3 / ((-1) | 1))
    var v49 uint32 = 509
    fmt.Println(v49, v49, v5)
    v4 = v4
    switch 128 {
    case 4294967295:
        if (228 < 42) {
            switch "hello" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                var v50 uint16 = 65535
                _ = v50
            case "":
                var v51 uint8 = 0
                _ = v51
            default:
                for fi52 := 0; fi52 < 14; fi52++ {
                    for fi53 := 0; fi53 < 18; fi53++ {
                        v1 = v1
                        if (true || (v3 == 1)) {
                            v1 = (v1 / ((v1) | 1))
                        } else {
                            v4 -= (v4 << ((v4) & 63))
                            if (v5 || ((v3 + v3) != v3)) {
                                var v54 int = fi52
                                v6 += int64(v1)
                                _ = v54
                            }
                        }
                        var v55 uint32 = v49
                        _ = fi53
                        _ = v55
                    }
                    if (("\n\t" + "foo bar") >= "a") {
                        var v56 int = (fi52 << (((v4 + uint(v49))) & 63))
                        _ = v56
                    } else {
                        var v57 uint = (^(v4))
                        var v58 uint32 = (v49 >> ((v57) & 63))
                        if (!(!v5)) {
                            var v59 int = 127
                            if (42 > 18446744073709551615) {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                switch v59 {
                                case 466:
                                    if v5 {
                                        for fi60 := 0; fi60 < 5; fi60++ {
                                            var v61 int8 = v1
                                            var v62 int64 = (v6 * v6)
                                            var v63 int64 = (v6 >> ((18446744073709551615) & 63))
                                            v3 += v3
                                            var v64 uint16 = 65535
                                            _ = len(v2)
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = len(v2)
                                        fmt.Println(v58, v4)
                                    }
                                }
                            } else {
                                if (false && (1e10 == 0.1)) {
                                    v65 := GenS6{F0: 289}
                                    v3 += (^(v3))
                                    _ = v65
                                }
                                if v5 {
                                    v58 = 1
                                    var v66 uint16 = 65535
                                    for fi67 := 0; fi67 < 7; fi67++ {
                                        if v5 {
                                            v68 := GenS6{F0: 4294967295}
                                            if (!(v66 <= 2)) {
                                                if v5 {
                                                    if v5 {
                                                        v4 = (v57 << ((v57) & 63))
                                                        fmt.Println(v49, v66)
                                                    } else {
                                                        v69 := GenS4{F0: -238}
                                                        _ = v69
                                                    }
                                                    fmt.Println(v58, v6, v6)
                                                } else {
                                                    if (!v5) {
                                                        if false {
                                                            if len(v2) > 0 {
                                                                v2[0] = "hello"
                                                            }
                                                            switch (fi52 >> ((v57) & 63)) {
                                                            case -9223372036854775808:
                                                                switch (v59 | 476) {
                                                                case -344:
                                                                    fmt.Println(fi67, fi67)
                                                                }
                                                            default:
                                                                var v70 uint16 = (v66 / (((^(v66))) | 1))
                                                                _ = v70
                                                            }
                                                        } else {
                                                            v59 -= (fi67 >> (((v4 + v4)) & 63))
                                                            if (!((0.0 > (3.14159 * float64(3.14159))) && v5)) {
                                                                var v71 uint16 = 255
                                                                _ = v71
                                                            } else {
                                                                v58 -= 2147483647
                                                            }
                                                        }
                                                        var v72 bool = false
                                                        v73 := GenS7{F0: v57}
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                }
                                                v74 := GenS8{F0: (((1e10 / ((3.14159 * (1e10 - (1.5 + 1.5))) / 3.14159)) + (3.14159 + (float32(3.14159) / (1e10 / ((0.0 / -2.25) - 1.5))))) != (0.1 * ((float32(3.14159) / 1.5) * (3.14159 + -2.25)))), F1: 379, F2: -140}
                                                v66 = v66
                                                v75 := GenS5{F0: -1}
                                                fmt.Println(v5)
                                                _ = v74
                                                _ = v75
                                            } else {
                                                switch ((("hello" + ("" + "a")) + ("" + "a")) + "hello") {
                                                case "":
                                                    if v5 {
                                                        var v76 int = (fi67 << ((uint(v66)) & 63))
                                                        _ = v76
                                                    }
                                                case "\n\t":
                                                    v77 := GenS3{F0: "\n\t"}
                                                    _ = v77
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                if v5 {
                                                    v2 = append(v2, (("hello" + ("hello" + "foo bar")) + ("foo bar" + ((("\n\t" + "\n\t") + "\n\t") + (((("a" + "foo bar") + "hello") + "foo bar") + "hello")))))
                                                } else {
                                                    fmt.Println(v1, v1, v58)
                                                }
                                                switch (v59 * v59) {
                                                case 0:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v78 := GenS9{F0: 42, F1: (v58 ^ v58)}
                                                if (1 <= v4) {
                                                    v5 = ((!(32767 > 805)) || false)
                                                }
                                                _ = len(v2)
                                                _ = v78
                                            }
                                            _ = v68
                                        } else {
                                            v79 := GenS3{F0: "a"}
                                            v80 := GenS10{F0: v1, F1: v1}
                                            if (!true) {
                                                if v5 {
                                                    v58 = v49
                                                    v3 = -1
                                                }
                                                var v81 uint16 = v66
                                                var v82 float64 = float64((0.0 - 0.0))
                                                _ = v81
                                                _ = v82
                                            } else {
                                                if true {
                                                    v83 := GenS8{F0: ((0 >= v58) && (!(!(true && true)))), F1: v3, F2: 1000}
                                                    _ = v83
                                                } else {
                                                    fmt.Println(v1, v57, v66)
                                                }
                                            }
                                            _ = v79
                                            _ = v80
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch ("hello" + "\n\t") {
                                        case "foo bar":
                                            fmt.Println(v59)
                                        case "hello":
                                            v84 := GenS10{F0: v1, F1: 7}
                                            _ = v84
                                        default:
                                            v85 := GenS3{F0: ""}
                                            _ = v85
                                        }
                                        _ = fi67
                                    }
                                    _ = v66
                                } else {
                                    v3 = -1
                                    switch v59 {
                                    case 4294967295:
                                        v59 = v59
                                    case 874:
                                        var v86 bool = v5
                                        _ = v86
                                    default:
                                        if v5 {
                                            var v87 int32 = v3
                                            if v5 {
                                                _ = len(v2)
                                                v1 = 100
                                                v88 := GenS11{F0: 33, F1: 255, F2: (-2.25 >= (3.14159 + -2.25))}
                                                _ = v88
                                            }
                                            _ = v87
                                        } else {
                                            v4 -= (v4 * (^(v57)))
                                            _ = len(v2)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                }
                            }
                            fmt.Println(v3)
                            _ = v59
                        } else {
                            var v89 float64 = 3.14159
                            v57 = (v57 << ((0) & 63))
                            _ = v89
                        }
                        _ = v57
                        _ = v58
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi52
                }
            }
            var v90 uint16 = 178
            v90 = v90
            _ = v90
        } else {
            for fi91 := 0; fi91 < 3; fi91++ {
                v1 = -128
                v5 = (uint16(v49) == 2)
                switch (("a" + "hello") + "\n\t") {
                case "foo bar":
                    fmt.Println(fi91)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi91
            }
            if v5 {
                v5 = (v5 && true)
                var v92 int64 = v6
                var v93 int64 = v6
                _ = v92
                _ = v93
            }
        }
    }
    if ((false && true) && ((1000 != 42) || (true && ((!v5) && (-108 <= 1))))) {
        for fi94 := 0; fi94 < 1; fi94++ {
            v2 = append(v2, ("hello" + ("" + "hello")))
            for fi95 := 0; fi95 < 6; fi95++ {
                v1 = (v1 - (v1 % ((-73) | 1)))
                var v96 int32 = (v3 >> ((2) & 63))
                _ = fi95
                _ = v96
            }
            var v97 string = ("" + "a")
            _ = fi94
            _ = v97
        }
    }
    for fi98 := 0; fi98 < 7; fi98++ {
        v99 := GenS12{F0: 0}
        v49 += 42
        var v100 int16 = 0
        _ = fi98
        _ = v100
        _ = v99
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v101 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
