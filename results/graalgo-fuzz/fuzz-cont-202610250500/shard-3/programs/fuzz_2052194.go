package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 bool; F2 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8 }
func genFn0(p0 int, p1 int) int64 {
    p0 = -9223372036854775808
    switch "hello" {
    case "\n\t":
        var v6 uint16 = 1
        _ = v6
    case "hello":
        fmt.Println(p0, p1)
    case "foo bar":
        var v7 string = "foo bar"
        _ = v7
    default:
        v8 := GenS1{F0: 0, F1: 1}
        _ = v8
    }
    p1 += p0
    _ = p0
    _ = p1
    return 125
}
func genFn1() uint32 {
    var v9 int32 = 0
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v9, fi10)
        v11 := []int{fi10, (fi10 + fi10)}
        _ = fi10
        _ = v11
    }
    var v12 bool = false
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v9 -= 1
            var v15 uint32 = 276
            var v16 int8 = -18
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    _ = v12
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 300
    var v3 string = ("\n\t" + "")
    v4 := []int{128}
    v5 := make(map[uint32]bool)
    v4 = append(v4, 345)
    if false {
        switch "foo bar" {
        case "foo bar":
            var v17 int16 = 42
            _ = v17
        default:
            v18 := GenS1{F0: 0, F1: 1}
            _ = v18
        }
        if (314 < 2) {
            fmt.Println(v2, v3)
        }
    } else {
        var v19 bool = (-9223372036854775808 > -816)
        v19 = v19
        _ = v19
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v4 = append(v4, (fi20 * fi20))
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            if (!false) {
                if false {
                    if false {
                        var v21 float64 = -2.25
                        v2 = 0
                        v21 = (v21 - (float64(1.5) - float64(1.5)))
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 11; fi22++ {
                        var v23 string = ""
                        v24 := GenS1{F0: 4294967295, F1: 42}
                        var v25 float32 = (1e10 * (0.1 + -2.25))
                        var v26 float32 = (v25 * (v25 / v25))
                        if true {
                            v25 = v26
                            v3 = v3
                            if false {
                                v27 := GenS1{F0: 42, F1: 2147483647}
                                switch (v23 + (v23 + v3)) {
                                case "foo bar":
                                    fmt.Println(v3, v26, v3)
                                    var v28 uint8 = 1
                                    v4 = append(v4, fi22)
                                    var v29 int16 = -901
                                    v23 = (("hello" + v3) + v3)
                                    v2 -= v2
                                    _ = v28
                                    _ = v29
                                case "":
                                    v2 += 18446744073709551615
                                case "\n\t":
                                    var v30 uint = (^(v2))
                                    _ = v30
                                default:
                                    v2 += 0
                                }
                                var v31 uint8 = 255
                                _ = v27
                                _ = v31
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        } else {
                            var v32 int = (fi20 << ((v2) & 63))
                            fmt.Println(v3, fi20, v3)
                            _ = v32
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    var v33 uint64 = 0
                    _ = v33
                } else {
                    var v34 float64 = float64(0.0)
                    _ = v34
                }
                v35 := GenS2{F0: 42, F1: 631, F2: 1}
                for fi36 := 0; fi36 < 2; fi36++ {
                    if false {
                        var v37 uint32 = genFn1()
                        fmt.Println(fi20, fi36, v37)
                        _ = v37
                    } else {
                        fmt.Println(v3, v2)
                        switch (fi20 >> (((v2 * (^(v2)))) & 63)) {
                        case 557:
                            if false {
                                _ = len(v1)
                                v3 = (((v3 + v3) + "foo bar") + (v3 + v3))
                                var v38 float64 = 0.1
                                _ = v38
                            } else {
                                v2 -= v2
                            }
                        case -914:
                            switch ((v3 + v3) + (v3 + v3)) {
                            case "":
                                if (1e10 <= 1e10) {
                                    if true {
                                        v39 := GenS3{F0: -1, F1: -120}
                                        _ = v39
                                    } else {
                                        v40 := GenS1{F0: 524, F1: 809}
                                        _ = v40
                                    }
                                    v41 := GenS4{F0: 0}
                                    _ = v41
                                } else {
                                    var v42 uint16 = uint16(v2)
                                    v43 := GenS5{F0: 127}
                                    v44 := GenS4{F0: 192}
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                            case "hello":
                                var v45 int32 = 42
                                _ = v45
                            default:
                                v3 = "\n\t"
                            }
                        case 42:
                            if true {
                                v2 = 1
                                if (1 != 127) {
                                    v46 := v5[65535]
                                    v47 := GenS4{F0: 42}
                                    _ = v46
                                    _ = v47
                                }
                                var v48 float64 = 3.14159
                                v2 -= 65535
                                _ = v48
                            }
                            v3 = "a"
                            _ = len(v1)
                        default:
                            switch v3 {
                            case "\n\t":
                                if (1 > 2) {
                                    v49 := GenS5{F0: 0}
                                    _ = v49
                                } else {
                                    v50 := GenS4{F0: 0}
                                    v2 += 42
                                    _ = v50
                                }
                            }
                            switch fi36 {
                            case 42:
                                v4 = append(v4, 0)
                            default:
                                if false {
                                    v2 = 18446744073709551615
                                    var v51 int8 = -1
                                    _ = v51
                                } else {
                                    switch (-(fi36)) {
                                    case -1:
                                        fmt.Println(fi36, fi20, fi20)
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                var v52 uint64 = 18446744073709551615
                                                switch -2147483648 {
                                                case 128:
                                                    if len(v4) > 0 {
                                                        v4[0] = (^(fi20))
                                                    }
                                                case 9223372036854775807:
                                                    v2 += v2
                                                case 0:
                                                    v3 = "foo bar"
                                                default:
                                                    if (!false) {
                                                        _ = len(v1)
                                                        if true {
                                                            if (!(!false)) {
                                                                v52 -= 783
                                                            } else {
                                                                v52 += 265
                                                                if (((v3 + ("" + "foo bar")) == ("\n\t" + "\n\t")) || (128 >= 1)) {
                                                                    delete(v5, 65535)
                                                                    var v53 uint32 = genFn1()
                                                                    _ = v53
                                                                }
                                                                fmt.Println(fi20, v2, fi36)
                                                            }
                                                        } else {
                                                            if (0.1 > 1e10) {
                                                                fmt.Println(fi36, v52, v2)
                                                                var v54 int64 = -1
                                                                var v55 uint8 = 215
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                                if (false && false) {
                                                    if (2147483647 == 42) {
                                                        var v56 uint = (v2 % (((v2 ^ v2)) | 1))
                                                        var v57 int8 = 42
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    v4 = append(v4, fi20)
                                                    var v58 int16 = 32767
                                                    v59 := GenS3{F0: 0, F1: 42}
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v52
                                            } else {
                                                var v60 int64 = 42
                                                if (-366 > 0) {
                                                    var v61 int16 = -32768
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v62 := GenS1{F0: uint64(v2), F1: 935}
                                                    if true {
                                                        if ((v60 >> (((^(v2))) & 63)) != genFn0(fi36, fi36)) {
                                                            var v63 int8 = 0
                                                            switch 0 {
                                                            case -2147483648:
                                                                switch v3 {
                                                                case "":
                                                                    var v64 float32 = (1.5 * -2.25)
                                                                    _ = v64
                                                                }
                                                            case -785:
                                                                v65 := GenS6{F0: (^(v60))}
                                                                _ = v65
                                                            default:
                                                                if true {
                                                                    v61 = v61
                                                                    v66 := GenS7{F0: false}
                                                                    _ = v66
                                                                } else {
                                                                    v67 := GenS2{F0: 340, F1: 977, F2: (v61 * -928)}
                                                                    _ = v67
                                                                }
                                                            }
                                                            _ = v63
                                                        }
                                                        v5[genFn1()] = false
                                                        var v68 bool = ((0.1 + 3.14159) >= 1e10)
                                                        _ = v68
                                                    }
                                                    _ = v61
                                                    _ = v62
                                                }
                                                _ = v60
                                            }
                                        case "foo bar":
                                            v2 = v2
                                        }
                                    }
                                    v69 := GenS2{F0: 343, F1: 184, F2: 1}
                                    fmt.Println(v3, fi36)
                                    v70 := GenS8{F0: -2147483648, F1: (true || (v3 == (("a" + v3) + ((v3 + v3) + "")))), F2: 1}
                                    _ = v69
                                    _ = v70
                                }
                            }
                            if ((!false) && false) {
                                v2 = v2
                            }
                            var v71 uint32 = genFn1()
                            if false {
                                var v72 uint8 = 1
                                _ = v72
                            } else {
                                if true {
                                    if (false && true) {
                                        var v73 uint16 = 65535
                                        var v74 int32 = -1
                                        var v75 uint16 = v73
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    } else {
                                        v71 = (v71 | 72)
                                        v71 += uint32(v2)
                                    }
                                }
                                v76 := GenS7{F0: false}
                                if ((-2147483648 < 42) && false) {
                                    var v77 uint = 255
                                    v78 := GenS5{F0: -128}
                                    v2 = (v2 % ((v2) | 1))
                                    var v79 uint64 = 42
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                } else {
                                    v3 = ""
                                    v1[v71] = "hello"
                                    v80 := v1[(v71 >> (((v2 * v2)) & 63))]
                                    _ = v80
                                }
                                v81 := GenS1{F0: 425, F1: 409}
                                _ = v76
                                _ = v81
                            }
                            _ = v71
                        }
                        switch v3 {
                        case "":
                            v3 = "\n\t"
                        }
                    }
                    var v82 int64 = 42
                    fmt.Println(fi20, fi36, v2)
                    switch ("hello" + v3) {
                    case "hello":
                        fmt.Println(fi20, fi36, v2)
                    case "":
                        if false {
                            if false {
                                if false {
                                    v82 += -1
                                    switch "foo bar" {
                                    case "a":
                                        v2 += v2
                                    case "foo bar":
                                        v82 = genFn0((fi36 >> (((^(v2))) & 63)), fi36)
                                    default:
                                        var v83 int = (-(fi20))
                                        _ = v83
                                    }
                                    switch v3 {
                                    case "foo bar":
                                        var v84 int32 = 254
                                        _ = v84
                                    case "a":
                                        v3 = v3
                                    case "":
                                        var v85 uint8 = 127
                                        _ = v85
                                    default:
                                        if (610 > 4294967295) {
                                            var v86 int32 = 804
                                            _ = v86
                                        } else {
                                            var v87 float64 = (1e10 / (1.5 * ((1e10 + (0.1 * -2.25)) * (float64((0.1 * 1e10)) / (3.14159 - (1.5 + ((1.5 + 3.14159) + (0.1 + -2.25))))))))
                                            _ = v87
                                        }
                                    }
                                }
                                v2 = v2
                                v88 := GenS9{F0: 25}
                                _ = v88
                            }
                            var v89 int = (fi36 / (((fi20 ^ (fi36 + (fi36 - fi20)))) | 1))
                            _ = len(v4)
                            _ = v89
                        }
                    default:
                        _ = len(v4)
                    }
                    _ = fi36
                    _ = v82
                }
                _ = v35
            }
        }
        _ = fi20
    }
    if (!(0.1 < -2.25)) {
        for fi90 := 0; fi90 < 10; fi90++ {
            for fi91 := 0; fi91 < 9; fi91++ {
                fmt.Println(fi90, fi91, v2)
                var v92 float32 = (0.1 / (1e10 / (-2.25 / 1e10)))
                v93 := GenS10{F0: (fi91 * 155)}
                switch (v3 + v3) {
                case "":
                    v92 -= ((3.14159 - ((1e10 / 1.5) - v92)) - (v92 + (1e10 * v92)))
                case "foo bar":
                    fmt.Println(v2)
                case "hello":
                    v94 := GenS8{F0: -1, F1: (!true), F2: 1}
                    _ = v94
                }
                switch (fi90 + (fi91 & (fi90 / ((fi91) | 1)))) {
                case 127:
                    delete(v1, 2)
                case -349:
                    v3 = v3
                case 4294967295:
                    var v95 uint8 = 128
                    _ = v95
                }
                _ = fi91
                _ = v92
                _ = v93
            }
            v96 := GenS6{F0: 1}
            _ = fi90
            _ = v96
        }
        v3 = ("\n\t" + "hello")
    } else {
        _ = len(v4)
        v3 = (v3 + "")
        for fi97 := 0; fi97 < 14; fi97++ {
            for fi98 := 0; fi98 < 20; fi98++ {
                fmt.Println(fi97, v3, fi97)
                if false {
                    var v99 float64 = (-2.25 / (3.14159 - (0.0 / (((0.1 / 3.14159) + 1e10) / 1.5))))
                    v100 := GenS10{F0: 4294967295}
                    var v101 int = (fi98 % (((fi98 - (fi98 | fi98))) | 1))
                    _ = v100
                    _ = v101
                    _ = v99
                }
                var v102 uint16 = 148
                _ = fi98
                _ = v102
            }
            _ = fi97
        }
    }
    fmt.Println(v3, v2)
    var v103 uint64 = 0
    if false {
        if false {
            v104 := GenS11{F0: 127}
            _ = v104
        }
        var v105 bool = (true && false)
        v105 = (!(-501 < -537))
        _ = v105
    } else {
        v106 := GenS6{F0: -9223372036854775808}
        _ = v106
    }
    v103 = v103
    v107 := GenS4{F0: -219}
    for fi108 := 0; fi108 < 18; fi108++ {
        if (!(!(1 <= 65535))) {
            fmt.Println(fi108, v3)
            var v109 float64 = float64(float32(float64(0.0)))
            _ = v109
        } else {
            var v110 float64 = (1e10 / 1e10)
            _ = v110
        }
        switch v3 {
        case "":
            v103 += v103
        default:
            var v111 uint = 108
            _ = v111
        }
        _ = fi108
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v103
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
