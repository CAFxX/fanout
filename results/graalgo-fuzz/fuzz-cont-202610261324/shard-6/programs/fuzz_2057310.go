package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    v3 := make(map[string]string)
    var v4 int32 = -1
    var v5 float64 = (3.14159 + 0.0)
    var v6 float64 = 1e10
    switch 297 {
    case -1:
        var v7 int = -9223372036854775808
        _ = v7
    default:
        v8 := GenS1{F0: 255}
        _ = v8
    }
    switch 9223372036854775807 {
    case 689:
        v6 -= ((3.14159 + v5) + float64(0.1))
    case -2147483648:
        var v9 int16 = -32768
        _ = v9
    case 128:
        fmt.Println(v6, v4, v2)
    }
    v10 := GenS1{F0: 0}
    v11 := GenS1{F0: 83}
    fmt.Println(v6)
    var v12 uint16 = 32767
    fmt.Println(v4, v5)
    if false {
        var v13 int32 = (-(v4))
        v2 = v2
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint = 542
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := []int{127, 830, 128, (fi14 / ((-1) | 1))}
                var v18 uint64 = 42
                _ = fi16
                _ = v17
                _ = v18
            }
            if ("" >= (("" + "hello") + "")) {
                if (-1 > 0) {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if false {
                            fmt.Println(v15, v2, fi19)
                            v20 := GenS1{F0: 213}
                            fmt.Println(fi19)
                            var v21 uint32 = 2147483647
                            _ = v20
                            _ = v21
                        } else {
                            v22 := []int{(fi19 + (fi14 << (((v15 - 18446744073709551615)) & 63))), (fi14 | fi19), (fi19 % (((fi14 + fi14)) | 1))}
                            switch 0 {
                            case 902:
                                v5 = ((v5 * (0.0 * v5)) / float64(1e10))
                            case 127:
                                if false {
                                    v4 = v4
                                    var v23 int32 = (v4 << ((v15) & 63))
                                    _ = v23
                                } else {
                                    var v24 uint = v15
                                    v12 -= v12
                                    v25 := GenS2{F0: 1000, F1: 128, F2: v15}
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                v5 += v6
                            }
                            _ = v22
                        }
                        v26 := GenS1{F0: 135}
                        delete(v3, "a")
                        _ = fi19
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 14; fi27++ {
                        fmt.Println(fi14, v5, v12)
                        if (v2 == 860) {
                            if (-9223372036854775808 == (fi14 / ((845) | 1))) {
                                switch 749 {
                                case 42:
                                    delete(v3, "")
                                case 8:
                                    var v28 int8 = 42
                                    _ = v28
                                case 4294967295:
                                    v29 := GenS2{F0: 0, F1: 69, F2: v15}
                                    _ = v29
                                }
                                var v30 uint = v15
                                var v31 int16 = 32767
                                _ = v30
                                _ = v31
                            } else {
                                switch 2147483647 {
                                case 128:
                                    v12 += 65535
                                case -2147483648:
                                    var v32 int64 = v2
                                    v33 := GenS2{F0: -32768, F1: 199, F2: (v15 + 18446744073709551615)}
                                    if true {
                                        delete(v1, fi27)
                                        var v34 int32 = v4
                                        if (0 <= -469) {
                                            var v35 int32 = (v4 - v4)
                                            v36 := GenS3{F0: 520}
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                    _ = v32
                                    _ = v33
                                }
                            }
                            fmt.Println(v4)
                            v4 += (v4 << ((v15) & 63))
                        } else {
                            var v37 uint64 = 1
                            if false {
                                var v38 string = "\n\t"
                                _ = v38
                            } else {
                                var v39 float64 = -2.25
                                if true {
                                    v40 := GenS3{F0: 2147483647}
                                    v12 -= 42
                                    _ = v40
                                } else {
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        var v41 bool = false
                                        var v42 string = (("foo bar" + (("hello" + "foo bar") + "hello")) + ("foo bar" + "\n\t"))
                                        delete(v1, (fi14 % ((fi14) | 1)))
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 float64 = 1.5
                                    switch 42 {
                                    case -284:
                                        if false {
                                            fmt.Println(v12, v39, v5)
                                            v44 := GenS2{F0: 32767, F1: uint8(v37), F2: v15}
                                            _ = v44
                                        }
                                    default:
                                        if (65535 < 544) {
                                            v45 := GenS4{F0: 0, F1: "a"}
                                            var v46 uint32 = 4294967295
                                            var v47 int32 = (v4 << ((420) & 63))
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        switch fi14 {
                                        case -888:
                                            switch fi14 {
                                            case 32768:
                                                if ((1e10 * 0.1) < -2.25) {
                                                    fmt.Println(v5, v37)
                                                    v2 = v2
                                                }
                                            case 810:
                                                v15 = v15
                                            case 127:
                                                v2 -= v2
                                            default:
                                                v39 += float64(((3.14159 - 1e10) - float32(float64((0.1 / 1e10)))))
                                            }
                                        case 124:
                                            v48 := []int{}
                                            _ = v48
                                        case -187:
                                            _ = len(v1)
                                        }
                                    }
                                    _ = v43
                                }
                                if (!true) {
                                    fmt.Println(v12)
                                    v12 = v12
                                    v49 := GenS1{F0: 121}
                                    _ = v49
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (fi27 <= fi14) {
                                            switch fi27 {
                                            case 32767:
                                                switch 32767 {
                                                case -2147483648:
                                                    var v50 uint8 = 153
                                                    _ = v50
                                                default:
                                                    if true {
                                                        v51 := GenS5{F0: true}
                                                        v2 = -1
                                                        fmt.Println(v6)
                                                        v15 = (v15 - 42)
                                                        _ = v51
                                                    }
                                                }
                                            }
                                            v52 := []int{(fi27 ^ fi14), 32768, (fi14 ^ fi14), fi27}
                                            var v53 int64 = v2
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                }
                                _ = v39
                            }
                            var v54 uint16 = v12
                            _ = v37
                            _ = v54
                        }
                        v55 := GenS2{F0: -812, F1: 0, F2: 598}
                        if true {
                            v4 += v4
                            delete(v3, "\n\t")
                        }
                        _ = fi27
                        _ = v55
                    }
                }
            } else {
                if (true || true) {
                    v5 -= v6
                    var v56 int = fi14
                    _ = v56
                }
                var v57 string = "foo bar"
                _ = v57
            }
            _ = fi14
            _ = v15
        }
    }
    v58 := v1[-584]
    var v59 int64 = v2
    v4 = (v4 * v4)
    var v60 uint64 = 136
    if (false && false) {
        var v61 bool = v58
        var v62 int64 = (v59 + -9223372036854775808)
        v63 := GenS4{F0: 0, F1: "hello"}
        _ = v61
        _ = v62
        _ = v63
    } else {
        switch 899 {
        case -168:
            if v58 {
                if (4294967295 != 2147483647) {
                    var v64 int32 = (v4 / ((v4) | 1))
                    if false {
                        for fi65 := 0; fi65 < 1; fi65++ {
                            v66 := []int{}
                            _ = fi65
                            _ = v66
                        }
                        v67 := []int{}
                        v68 := GenS6{F0: 32768, F1: 867}
                        _ = v67
                        _ = v68
                    }
                    _ = v64
                } else {
                    fmt.Println(v60, v58, v5)
                    v58 = (!(65535 <= uint32(v12)))
                }
                if (65535 >= 2147483647) {
                    delete(v3, ("" + ("hello" + "hello")))
                } else {
                    v3[("foo bar" + "hello")] = ("" + "foo bar")
                    v69 := GenS1{F0: 128}
                    _ = v69
                }
                for fi70 := 0; fi70 < 4; fi70++ {
                    v6 += v6
                    v71 := []int{}
                    _ = fi70
                    _ = v71
                }
            }
        case -9223372036854775808:
            var v72 string = ""
            _ = v72
        default:
            v5 = v5
        }
    }
    if ("hello" <= "\n\t") {
        if (-879 == 1000) {
            v73 := GenS7{F0: -32768, F1: 255}
            var v74 uint8 = 255
            if v58 {
                delete(v1, 128)
                switch ("hello" + "foo bar") {
                case "\n\t":
                    var v75 int64 = (v2 >> ((454) & 63))
                    _ = v75
                case "a":
                    var v76 uint8 = (^(v74))
                    _ = v76
                case "hello":
                    v5 += (v5 * v5)
                }
            } else {
                fmt.Println(v58, v58, v60)
                var v77 float64 = float64(-2.25)
                if v58 {
                    var v78 string = "a"
                    v79 := []int{int(v74)}
                    _ = v78
                    _ = v79
                }
                _ = v77
            }
            _ = v73
            _ = v74
        } else {
            v80 := []int{-870}
            if (v58 || v58) {
                var v81 float32 = (-2.25 / 0.1)
                _ = v81
            } else {
                for fi82 := 0; fi82 < 19; fi82++ {
                    v58 = ((fi82 == 4294967295) && v58)
                    var v83 float64 = (1.5 * 1e10)
                    if (uint16(v60) <= v12) {
                        if (-2.25 > ((float32(1.5) * 0.1) - 3.14159)) {
                            if len(v80) >= 2 {
                                v80 = v80[1:len(v80)]
                            }
                            if (v60 != v60) {
                                var v84 int = fi82
                                if false {
                                    v85 := GenS1{F0: 0}
                                    fmt.Println(v83, v60, v2)
                                    _ = v85
                                } else {
                                    var v86 uint32 = 4294967295
                                    var v87 float32 = float32((-2.25 - (v6 * (v5 / -2.25))))
                                    v88 := GenS4{F0: int16(v60), F1: "foo bar"}
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                                var v89 float32 = 0.1
                                _ = v84
                                _ = v89
                            }
                        } else {
                            var v90 int = fi82
                            _ = v90
                        }
                    } else {
                        v91 := GenS6{F0: fi82, F1: 0}
                        _ = v91
                    }
                    _ = fi82
                    _ = v83
                }
                var v92 uint32 = 578
                v92 = 466
                _ = v92
            }
            _ = v80
        }
        v93 := []int{31, -684, 32767}
        v94 := GenS8{F0: 0, F1: 0, F2: -9223372036854775808}
        _ = v93
        _ = v94
    }
    v3[((("" + ("\n\t" + ("foo bar" + ("foo bar" + "\n\t")))) + (("a" + "hello") + "")) + ("" + "hello"))] = "hello"
    var v95 int32 = (-(v4))
    for fi96 := 0; fi96 < 11; fi96++ {
        v58 = false
        if (!v58) {
            fmt.Println(v59, v95, v60)
            v5 = v6
        }
        v95 += (v95 >> ((258) & 63))
        _ = fi96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v95
}
