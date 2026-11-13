package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 string; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 int; F1 int64; F2 uint64 }
type GenS11 struct { F0 int64; F1 uint16; F2 string }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32; F1 string; F2 int }
type GenS15 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 int16, p2 uint16) int32 {
    p0 += (p0 + p0)
    var v6 uint32 = 114
    v6 += (v6 / (((v6 / ((v6) | 1))) | 1))
    v7 := GenS1{F0: "a"}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int32 = -601
        switch fi8 {
        case 9223372036854775807:
            v10 := []int{}
            _ = v10
        case 127:
            v11 := []int{int(p2), int(p2), fi8, fi8}
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -792
}
func genFn1(p0 int, p1 string) uint8 {
    var v12 int16 = -125
    var v13 int32 = -1
    v13 -= v13
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 201
}

func main() {
    var v1 int16 = -1
    var v2 int8 = 1
    var v3 int8 = (v2 % ((v2) | 1))
    var v4 bool = false
    var v5 string = ((("a" + "\n\t") + "") + "")
    if (true || v4) {
        v14 := GenS2{F0: v3, F1: 9223372036854775807, F2: 65535}
        _ = v14
    }
    var v15 int64 = -1
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := []int{}
            if len(v18) > 0 {
                v18[0] = (-(fi17))
            }
            _ = fi17
            _ = v18
        }
        var v19 int32 = 711
        _ = fi16
        _ = v19
    }
    var v20 float64 = 0.0
    if (!false) {
        fmt.Println(v1)
        var v21 uint8 = 42
        v2 -= (v2 >> ((828) & 63))
        _ = v21
    } else {
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := []int{(fi23 - 32767)}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                v3 = (v2 + v2)
                v20 -= 0.1
                var v26 float32 = ((float32(v20) - 3.14159) / 1e10)
                _ = fi25
                _ = v26
            }
            if ((fi22 / ((-942) | 1)) > 4294967295) {
                fmt.Println(v5, fi22)
                for fi27 := 0; fi27 < 13; fi27++ {
                    v28 := GenS1{F0: "foo bar"}
                    if ((1e10 == ((float32(v20) / 1e10) + ((0.0 * float32((v20 * (v20 - v20)))) + (-2.25 * (float32(v20) * (1e10 - ((0.0 - float32(v20)) / -2.25))))))) && v4) {
                        v29 := GenS2{F0: v3, F1: (fi27 * fi27), F2: 42}
                        if v4 {
                            v30 := []int{(fi27 ^ 4294967295), (fi27 / ((fi27) | 1)), (fi27 - fi22), (fi22 / (((fi27 ^ fi27)) | 1))}
                            _ = v30
                        }
                        if (1 > 255) {
                            fmt.Println(fi22)
                            if v4 {
                                v31 := []int{(fi22 + (^(fi27))), (fi22 % (((-(fi27))) | 1))}
                                _ = v31
                            }
                        } else {
                            var v32 uint8 = 182
                            fmt.Println(fi27)
                            _ = v32
                        }
                        _ = v29
                    }
                    _ = fi27
                    _ = v28
                }
                var v33 int16 = (v1 + v1)
                _ = v33
            } else {
                if v4 {
                    v3 -= v3
                    if (!v4) {
                        var v34 string = (("" + "a") + v5)
                        v15 += v15
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 16; fi35++ {
                            if (v4 || true) {
                                if true {
                                    var v36 float64 = v20
                                    _ = v36
                                }
                                v37 := GenS3{F0: false, F1: 42}
                                var v38 float64 = 3.14159
                                _ = v37
                                _ = v38
                            } else {
                                var v39 uint = 255
                                switch "foo bar" {
                                case "":
                                    v40 := GenS4{F0: 2147483647}
                                    _ = v40
                                case "a":
                                    v41 := []int{int(v15), -967}
                                    _ = v41
                                }
                                _ = v39
                            }
                            var v42 bool = v4
                            _ = fi35
                            _ = v42
                        }
                    }
                    fmt.Println(v1)
                } else {
                    switch (fi22 >> ((611) & 63)) {
                    case 32768:
                        if (v4 || ((268 == 922) || (v4 && ((v15 & 608) == -1)))) {
                            var v43 string = v5
                            for fi44 := 0; fi44 < 8; fi44++ {
                                if false {
                                    switch v5 {
                                    case "\n\t":
                                        if false {
                                            var v45 string = ("" + "\n\t")
                                            v46 := []int{}
                                            _ = v45
                                            _ = v46
                                        }
                                    case "":
                                        v2 -= (^(v3))
                                    case "foo bar":
                                        if false {
                                            v4 = false
                                        } else {
                                            fmt.Println(fi44, v20, v4)
                                            if (255 == 128) {
                                                v1 += -1
                                                fmt.Println(v15, v5, fi22)
                                                if (1 == 128) {
                                                    var v47 uint = 12
                                                    fmt.Println(v5, v2, fi44)
                                                    _ = v47
                                                }
                                            } else {
                                                var v48 uint = 42
                                                if false {
                                                    v2 += v2
                                                } else {
                                                    v49 := []int{42, fi44, (fi44 % ((fi44) | 1))}
                                                    v50 := GenS5{F0: 222, F1: v4}
                                                    var v51 float32 = float32((0.0 * (v20 / v20)))
                                                    if (2147483648 == 2147483647) {
                                                        fmt.Println(v3, fi22, v1)
                                                        if v4 {
                                                            if (!v4) {
                                                                if (true && false) {
                                                                    v52 := GenS6{F0: 144, F1: v4}
                                                                    if ((v4 || v4) || (4294967295 >= 2)) {
                                                                        fmt.Println(v3)
                                                                        var v53 float64 = (float64(((v51 * -2.25) - -2.25)) * v20)
                                                                        var v54 float32 = v51
                                                                        _ = v53
                                                                        _ = v54
                                                                    } else {
                                                                        var v55 bool = (!true)
                                                                        if v55 {
                                                                            if v4 {
                                                                                if true {
                                                                                    v56 := GenS7{F0: ("foo bar" + "hello"), F1: (v15 * v15), F2: 1}
                                                                                    fmt.Println(v3)
                                                                                    _ = v56
                                                                                }
                                                                            } else {
                                                                                v57 := GenS8{F0: 42, F1: genFn0((v3 | (v3 | v2)), 32767, 150)}
                                                                                if (v55 && (v43 <= "a")) {
                                                                                    v1 = (v1 >> ((42) & 63))
                                                                                    switch ((v5 + v43) + (v5 + ("\n\t" + ("a" + "a")))) {
                                                                                    case "hello":
                                                                                        var v58 float64 = (3.14159 / v20)
                                                                                        _ = v58
                                                                                    default:
                                                                                        if (v1 <= 1) {
                                                                                            var v59 uint = (v48 / ((v48) | 1))
                                                                                            v60 := GenS4{F0: int32(v59)}
                                                                                            _ = v59
                                                                                            _ = v60
                                                                                        } else {
                                                                                            if v55 {
                                                                                                if (((fi22 * (fi22 % (((^(fi44))) | 1))) >= (fi22 | fi22)) || (-1 > -1)) {
                                                                                                    if (32768 >= fi44) {
                                                                                                        var v61 uint32 = 300
                                                                                                        if v55 {
                                                                                                            v51 = 1.5
                                                                                                        } else {
                                                                                                            v49 = append(v49, fi22)
                                                                                                            var v62 int64 = (-(v15))
                                                                                                            fmt.Println(v43)
                                                                                                            _ = v62
                                                                                                        }
                                                                                                        _ = v61
                                                                                                    } else {
                                                                                                        v63 := GenS1{F0: v5}
                                                                                                        _ = v63
                                                                                                    }
                                                                                                }
                                                                                                if (!false) {
                                                                                                    v5 = ("foo bar" + (v43 + "a"))
                                                                                                    v43 = ("a" + v5)
                                                                                                } else {
                                                                                                    var v64 uint = (v48 << ((v48) & 63))
                                                                                                    if len(v49) >= 2 {
                                                                                                        v49 = v49[1:len(v49)]
                                                                                                    }
                                                                                                    if len(v49) > 0 {
                                                                                                        v49[0] = fi22
                                                                                                    }
                                                                                                    _ = v64
                                                                                                }
                                                                                                if false {
                                                                                                    fmt.Println(v20, v1, fi22)
                                                                                                    v4 = (((42 > 42) || (326 != 101)) || (127 != 254))
                                                                                                    v43 = "hello"
                                                                                                }
                                                                                                switch v5 {
                                                                                                case "":
                                                                                                    if len(v49) >= 2 {
                                                                                                        v49 = v49[1:len(v49)]
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    var v65 int32 = -2147483648
                                                                                                    _ = v65
                                                                                                default:
                                                                                                    if len(v49) > 0 {
                                                                                                        v49[0] = fi22
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v3 = 1
                                                                                        if true {
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                v49 = append(v49, (fi22 / ((1) | 1)))
                                                                                            case "a":
                                                                                                var v66 int64 = (v15 & (-(v15)))
                                                                                                _ = v66
                                                                                            default:
                                                                                                var v67 int32 = -2147483648
                                                                                                _ = v67
                                                                                            }
                                                                                            if (2147483647 <= -2147483648) {
                                                                                                if ((true || v4) || v55) {
                                                                                                    fmt.Println(v5)
                                                                                                    var v68 uint = 2
                                                                                                    _ = v68
                                                                                                }
                                                                                                if (372 >= 65535) {
                                                                                                    v49 = append(v49, 0)
                                                                                                    var v69 uint64 = 2
                                                                                                    if len(v49) >= 2 {
                                                                                                        v49 = v49[1:len(v49)]
                                                                                                    }
                                                                                                    _ = v69
                                                                                                } else {
                                                                                                    var v70 int32 = genFn0(-69, (v1 - 32767), 399)
                                                                                                    if false {
                                                                                                        var v71 bool = v4
                                                                                                        var v72 int8 = v2
                                                                                                        _ = v71
                                                                                                        _ = v72
                                                                                                    } else {
                                                                                                        if len(v49) > 0 {
                                                                                                            v49[0] = (fi22 | -260)
                                                                                                        }
                                                                                                        var v73 int32 = 1
                                                                                                        var v74 float64 = float64(v51)
                                                                                                        _ = v73
                                                                                                        _ = v74
                                                                                                    }
                                                                                                    _ = v70
                                                                                                }
                                                                                                switch ("\n\t" + "") {
                                                                                                case "a":
                                                                                                    v2 = v2
                                                                                                default:
                                                                                                    fmt.Println(v5, v1)
                                                                                                }
                                                                                                if len(v49) > 0 {
                                                                                                    v49[0] = -9223372036854775808
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v55, v3)
                                                                                    v49 = append(v49, (fi44 + (-(fi44))))
                                                                                }
                                                                                var v75 uint16 = 854
                                                                                v75 = 565
                                                                                v76 := GenS1{F0: (v43 + ((v43 + "a") + (v5 + v43)))}
                                                                                fmt.Println(v48)
                                                                                if (-456 > -1) {
                                                                                    if len(v49) >= 2 {
                                                                                        v49 = v49[1:len(v49)]
                                                                                    }
                                                                                    if true {
                                                                                        if len(v49) > 0 {
                                                                                            v49[0] = fi22
                                                                                        }
                                                                                        var v77 uint = 65535
                                                                                        v48 += (v77 + (v77 + 857))
                                                                                        _ = v77
                                                                                    } else {
                                                                                        fmt.Println(v55, v1, v3)
                                                                                    }
                                                                                } else {
                                                                                    v78 := GenS3{F0: v55, F1: v48}
                                                                                    _ = v78
                                                                                }
                                                                                _ = v57
                                                                                _ = v75
                                                                                _ = v76
                                                                            }
                                                                        } else {
                                                                            if (255 > 65535) {
                                                                                v15 = v15
                                                                                _ = len(v49)
                                                                            } else {
                                                                                fmt.Println(v5)
                                                                            }
                                                                        }
                                                                        v1 = -32768
                                                                        _ = v55
                                                                    }
                                                                    _ = v52
                                                                } else {
                                                                    v79 := GenS9{F0: fi22, F1: v2}
                                                                    var v80 int32 = 0
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                if v4 {
                                                                    v81 := GenS10{F0: 128, F1: v15, F2: 1}
                                                                    _ = v81
                                                                } else {
                                                                    switch ("" + ("hello" + v5)) {
                                                                    case "foo bar":
                                                                        v48 -= (v48 / (((v48 * 393)) | 1))
                                                                    }
                                                                    if len(v49) > 0 {
                                                                        v49[0] = 32767
                                                                    }
                                                                }
                                                            } else {
                                                                if ((((v51 * 3.14159) * (v51 - v51)) + 0.0) > (v51 * ((v51 - (v51 * v51)) / float32(v20)))) {
                                                                    fmt.Println(v48, fi22)
                                                                }
                                                                v49 = append(v49, fi22)
                                                            }
                                                            v51 = float32(1.5)
                                                            v82 := GenS7{F0: "", F1: int64(v2), F2: 693}
                                                            _ = v82
                                                        } else {
                                                            var v83 int = -318
                                                            switch v43 {
                                                            case "\n\t":
                                                                v2 -= (v3 | v2)
                                                            default:
                                                                if (v15 <= 0) {
                                                                    if v4 {
                                                                        var v84 int16 = (v1 / (((v1 & (v1 >> ((v48) & 63)))) | 1))
                                                                        if len(v49) >= 2 {
                                                                            v49 = v49[1:len(v49)]
                                                                        }
                                                                        _ = v84
                                                                    } else {
                                                                        v48 = v48
                                                                    }
                                                                } else {
                                                                    var v85 float64 = v20
                                                                    if ((v15 % (((v15 & (v15 << (((v48 << ((v48) & 63))) & 63)))) | 1)) < (-(v15))) {
                                                                        _ = len(v49)
                                                                        fmt.Println(v3, v85, fi44)
                                                                    }
                                                                    v4 = v4
                                                                    _ = v85
                                                                }
                                                            }
                                                            _ = v83
                                                        }
                                                        v49 = append(v49, (fi22 % (((fi22 >> (((v48 - (^(v48)))) & 63))) | 1)))
                                                    } else {
                                                        v43 = v5
                                                        var v86 int8 = 42
                                                        var v87 float64 = 1.5
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v48
                                            }
                                            fmt.Println(fi22, v2)
                                        }
                                    }
                                }
                                _ = fi44
                            }
                            if true {
                                v88 := []int{(fi22 & fi22), (-(fi22))}
                                switch (fi22 + (-(fi22))) {
                                case -921:
                                    v89 := GenS11{F0: (-(v15)), F1: 631, F2: (v5 + ("" + "hello"))}
                                    _ = v89
                                }
                                for fi90 := 0; fi90 < 16; fi90++ {
                                    if len(v88) >= 2 {
                                        v88 = v88[1:len(v88)]
                                    }
                                    if (!(217 != 162)) {
                                        if ((v4 && (v1 > -32768)) || false) {
                                            fmt.Println(fi90)
                                            var v91 uint16 = 387
                                            if (0 > 42) {
                                                if len(v88) >= 2 {
                                                    v88 = v88[1:len(v88)]
                                                }
                                                if v4 {
                                                    var v92 float64 = (v20 / ((float64((0.0 * 0.1)) + v20) / 3.14159))
                                                    v3 = v3
                                                    v43 = "hello"
                                                    _ = v92
                                                }
                                                v1 += (v1 << ((1) & 63))
                                            } else {
                                                if (2147483647 == 1) {
                                                    if len(v88) >= 2 {
                                                        v88 = v88[1:len(v88)]
                                                    }
                                                    switch (v43 + v43) {
                                                    case "\n\t":
                                                        if v4 {
                                                            if len(v88) >= 2 {
                                                                v88 = v88[1:len(v88)]
                                                            }
                                                        } else {
                                                            v93 := GenS2{F0: 0, F1: (fi90 ^ (fi22 / (((fi90 >> ((255) & 63))) | 1))), F2: (v91 ^ 2)}
                                                            v94 := GenS1{F0: "a"}
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    default:
                                                        _ = len(v88)
                                                    }
                                                    v3 = v2
                                                } else {
                                                    fmt.Println(v15)
                                                    v95 := GenS12{F0: genFn1((fi22 - (-(fi90))), ("a" + "foo bar")), F1: 2}
                                                    _ = v95
                                                }
                                            }
                                            _ = v91
                                        } else {
                                            fmt.Println(fi22, v5)
                                            if (v4 && v4) {
                                                fmt.Println(v5)
                                                if ((fi90 * (fi22 << ((6) & 63))) <= 127) {
                                                    v3 = v3
                                                }
                                            }
                                            v20 -= (((float64((3.14159 * float32(float64((3.14159 * float32(float64(0.1))))))) / (float64(0.0) + (float64(0.0) / v20))) * float64(0.1)) / (v20 + (((v20 + float64((0.0 / 1.5))) / v20) + v20)))
                                            switch (v43 + "hello") {
                                            case "foo bar":
                                                if len(v88) >= 2 {
                                                    v88 = v88[1:len(v88)]
                                                }
                                            case "":
                                                var v96 uint32 = 576
                                                _ = v96
                                            }
                                        }
                                        v97 := GenS5{F0: -684, F1: (genFn0(v3, -1, 0) < 443)}
                                        var v98 int8 = (v3 % (((^(v3))) | 1))
                                        _ = v97
                                        _ = v98
                                    } else {
                                        v4 = true
                                        var v99 string = ((v5 + "hello") + "hello")
                                        v20 -= 1.5
                                        _ = v99
                                    }
                                    if v4 {
                                        fmt.Println(v4, v1)
                                        v5 = v5
                                        v20 = v20
                                    }
                                    var v100 int64 = (v15 >> ((0) & 63))
                                    var v101 float64 = (v20 * float64(float32((v20 + 1.5))))
                                    if false {
                                        fmt.Println(fi90, v2)
                                        v101 = ((v20 * (v101 + v20)) + (float64(-2.25) + (((v101 + ((v101 / v20) + (0.1 + v101))) * v20) / ((v101 * ((3.14159 / v101) + 0.0)) / 3.14159))))
                                    }
                                    switch ("foo bar" + "") {
                                    case "foo bar":
                                        var v102 uint = 4294967295
                                        _ = v102
                                    case "\n\t":
                                        v103 := GenS13{F0: 992}
                                        _ = v103
                                    default:
                                        if ((v4 || (genFn0(v2, 32767, 2) >= 0)) || true) {
                                            if v4 {
                                                v4 = ((v100 + (v100 << ((42) & 63))) != -512)
                                                switch v5 {
                                                case "hello":
                                                    if len(v88) >= 2 {
                                                        v88 = v88[1:len(v88)]
                                                    }
                                                case "a":
                                                    v43 = ("\n\t" + "a")
                                                default:
                                                    var v104 int64 = v15
                                                    _ = v104
                                                }
                                            }
                                            if (123 >= 4294967295) {
                                                v105 := GenS14{F0: 853, F1: v5, F2: int(v100)}
                                                _ = v105
                                            }
                                            var v106 float32 = (1e10 * 0.1)
                                            _ = v106
                                        }
                                        v107 := GenS15{F0: (v4 && (v100 >= (^(v100)))), F1: 585}
                                        var v108 int16 = v1
                                        var v109 string = ""
                                        v3 += v2
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    }
                                    var v110 uint64 = 2
                                    switch "\n\t" {
                                    case "":
                                        var v111 uint = 18446744073709551615
                                        _ = v111
                                    case "foo bar":
                                        v110 = 18446744073709551615
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v112 int32 = int32(v1)
                                        _ = v112
                                    case "":
                                        var v113 uint32 = uint32(v110)
                                        _ = v113
                                    case "a":
                                        switch (fi90 - (fi22 >> ((2) & 63))) {
                                        case -695:
                                            var v114 uint16 = 0
                                            _ = v114
                                        case -995:
                                            if false {
                                                v4 = (true || false)
                                            } else {
                                                v43 = (v43 + v5)
                                            }
                                        }
                                    default:
                                        v115 := GenS14{F0: 2147483647, F1: (v5 + ((v43 + v43) + "")), F2: (fi22 % ((127) | 1))}
                                        _ = v115
                                    }
                                    var v116 uint16 = 276
                                    v116 += v116
                                    v101 -= v101
                                    _ = fi90
                                    _ = v100
                                    _ = v101
                                    _ = v110
                                    _ = v116
                                }
                                _ = v88
                            }
                            _ = v43
                        } else {
                            var v117 uint32 = 42
                            switch "" {
                            case "":
                                for fi118 := 0; fi118 < 4; fi118++ {
                                    var v119 float32 = 1.5
                                    _ = fi118
                                    _ = v119
                                }
                            case "foo bar":
                                var v120 uint8 = 18
                                _ = v120
                            }
                            _ = v117
                        }
                    case 42:
                        for fi121 := 0; fi121 < 2; fi121++ {
                            v20 = (1e10 / v20)
                            var v122 bool = true
                            if ((1 == 653) || true) {
                                var v123 bool = false
                                var v124 float32 = 1e10
                                v125 := []int{fi121}
                                _ = v123
                                _ = v124
                                _ = v125
                            }
                            _ = fi121
                            _ = v122
                        }
                    }
                }
                var v126 uint32 = 0
                _ = v126
            }
            _ = fi22
        }
        var v127 bool = false
        _ = v127
    }
    v128 := GenS6{F0: 0, F1: v4}
    if true {
        if (((v20 - v20) + (v20 + 0.1)) > float64((1.5 / 1.0))) {
            var v129 int16 = v1
            _ = v129
        }
        v5 = v5
        v130 := GenS14{F0: -1, F1: v5, F2: -239}
        _ = v130
    } else {
        v15 = v15
    }
    if v4 {
        var v131 uint32 = 2147483647
        v3 -= (v3 % (((v2 | 42)) | 1))
        _ = v131
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v128
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
