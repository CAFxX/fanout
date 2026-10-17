package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 int16; F2 uint8 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 bool; F2 int32 }
type GenS11 struct { F0 uint32; F1 int64; F2 int16 }
type GenS12 struct { F0 int64; F1 string; F2 uint16 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint32; F1 bool }
type GenS15 struct { F0 int32 }
type GenS16 struct { F0 int64; F1 int32; F2 bool }
type GenS17 struct { F0 uint16; F1 uint }
type GenS18 struct { F0 string; F1 uint8 }
type GenS19 struct { F0 string; F1 uint64; F2 uint32 }
type GenS20 struct { F0 int16 }
func genFn0(p0 uint, p1 string, p2 uint16) uint {
    p2 = p2
    v6 := GenS1{F0: -857, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (^(p0))
}
func genFn1(p0 uint, p1 bool) string {
    for fi7 := 0; fi7 < 10; fi7++ {
        switch fi7 {
        case 42:
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(p0, p1)
                _ = fi8
            }
        }
        _ = fi7
    }
    var v9 float32 = 0.0
    v10 := []int{}
    switch "hello" {
    case "hello":
        v11 := GenS1{F0: 32767, F1: 0}
        _ = v11
    case "a":
        p0 = genFn0((^(p0)), "\n\t", 775)
    case "\n\t":
        var v12 int8 = -79
        _ = v12
    default:
        if p1 {
            switch "hello" {
            case "hello":
                fmt.Println(p0, p0)
            }
            switch 4294967295 {
            case 2147483647:
                p1 = (-304 != 726)
            }
        }
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return ("foo bar" + "\n\t")
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint8 = 2
    v13 := GenS2{F0: false, F1: 680, F2: v5}
    v14 := GenS1{F0: -32768, F1: 562}
    v15 := GenS2{F0: (!((-(v1)) <= v1)), F1: -1, F2: 1}
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            fmt.Println(v5)
            _ = fi17
        }
        _ = fi16
    }
    v1 += v1
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int32 = 0
        fmt.Println(v5)
        var v20 uint16 = 32767
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 bool = false
    if true {
        v22 := GenS1{F0: -535, F1: 9223372036854775807}
        switch "" {
        case "foo bar":
            switch v1 {
            case 32767:
                if v21 {
                    v1 += v1
                    if v21 {
                        v23 := GenS1{F0: 32767, F1: -9223372036854775808}
                        _ = v23
                    } else {
                        var v24 uint32 = 2
                        if v21 {
                            if v21 {
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    var v26 uint16 = 65535
                                    if (!(false || ((v5 / (((v5 << ((genFn0(2, ("a" + "foo bar"), v26)) & 63))) | 1)) > 235))) {
                                        v27 := []int{32768, (^(fi25))}
                                        v28 := GenS1{F0: 903, F1: 863}
                                        _ = v27
                                        _ = v28
                                    }
                                    v26 -= 164
                                    _ = fi25
                                    _ = v26
                                }
                                var v29 string = "foo bar"
                                if true {
                                    if false {
                                        for fi30 := 0; fi30 < 1; fi30++ {
                                            switch v29 {
                                            case "":
                                                switch v29 {
                                                case "hello":
                                                    v4[v29] = (v24 >> ((0) & 63))
                                                case "a":
                                                    for fi31 := 0; fi31 < 17; fi31++ {
                                                        var v32 float64 = -2.25
                                                        v1 = 1
                                                        if v21 {
                                                            var v33 uint8 = v5
                                                            v1 -= (fi31 << ((741) & 63))
                                                            switch ((v29 + "a") + (v29 + "hello")) {
                                                            case "foo bar":
                                                                v33 += v5
                                                            case "":
                                                                if (!((718 <= 654) && (0.1 < (v32 + (3.14159 - 1.5))))) {
                                                                    switch (^(fi30)) {
                                                                    case -232:
                                                                        v2[197] = "hello"
                                                                    }
                                                                    if v21 {
                                                                        fmt.Println(v32, v32, v24)
                                                                        var v34 string = v29
                                                                        v35 := []int{(fi30 << ((1) & 63)), -405, fi31}
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v1, v1, v33)
                                                                        fmt.Println(v1)
                                                                    } else {
                                                                        if (0.1 != -2.25) {
                                                                            v32 = v32
                                                                            v33 += 42
                                                                            v36 := GenS3{F0: int(v5), F1: 65535}
                                                                            _ = v36
                                                                        }
                                                                        v29 = "foo bar"
                                                                    }
                                                                    if v21 {
                                                                        if true {
                                                                            v5 = (v33 + v33)
                                                                            var v37 int32 = 488
                                                                            switch (fi30 & fi30) {
                                                                            case -9223372036854775808:
                                                                                v38 := GenS3{F0: (fi31 * fi31), F1: 2}
                                                                                _ = v38
                                                                            case -1000:
                                                                                v24 = v24
                                                                            default:
                                                                                v21 = ((v5 & (v5 - v33)) == (v33 << ((genFn0(genFn0(genFn0(0, v29, 0), ((v29 + (v29 + v29)) + "a"), 1), "", 539)) & 63)))
                                                                            }
                                                                            v39 := []int{(v1 << ((1) & 63)), 128, (fi31 ^ -9223372036854775808), fi31}
                                                                            if v21 {
                                                                                if len(v39) > 0 {
                                                                                    v39[0] = -1
                                                                                }
                                                                                var v40 int64 = 1
                                                                                v41 := GenS4{F0: 42}
                                                                                _ = v40
                                                                                _ = v41
                                                                            } else {
                                                                                if v21 {
                                                                                    if false {
                                                                                        if (v21 && v21) {
                                                                                            v32 += 1e10
                                                                                            if ((v21 || ((false || v21) || v21)) && (v29 <= (("hello" + v29) + v29))) {
                                                                                                if v21 {
                                                                                                    v42 := GenS3{F0: 32767, F1: 390}
                                                                                                    _ = v42
                                                                                                } else {
                                                                                                    if true {
                                                                                                        switch (v29 + ("foo bar" + v29)) {
                                                                                                        case "a":
                                                                                                            if len(v39) >= 2 {
                                                                                                                v39 = v39[1:len(v39)]
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            if false {
                                                                                                                v24 += (v24 ^ (^(v24)))
                                                                                                                switch v29 {
                                                                                                                case "\n\t":
                                                                                                                    var v43 uint16 = 39
                                                                                                                    _ = v43
                                                                                                                case "hello":
                                                                                                                    var v44 float32 = (float32((0.0 - v32)) * float32(0.0))
                                                                                                                    _ = v44
                                                                                                                }
                                                                                                                v45 := GenS2{F0: false, F1: 0, F2: v33}
                                                                                                                _ = v45
                                                                                                            }
                                                                                                        default:
                                                                                                            v46 := GenS5{F0: 128, F1: (v37 << ((65535) & 63))}
                                                                                                            _ = v46
                                                                                                        }
                                                                                                        var v47 int32 = v37
                                                                                                        v48 := GenS6{F0: 23, F1: 48}
                                                                                                        _ = v47
                                                                                                        _ = v48
                                                                                                    } else {
                                                                                                        var v49 uint64 = 849
                                                                                                        fmt.Println(v29)
                                                                                                        v50 := GenS5{F0: (fi30 & 1), F1: 42}
                                                                                                        v49 -= 2
                                                                                                        _ = v49
                                                                                                        _ = v50
                                                                                                    }
                                                                                                    v51 := GenS3{F0: (fi30 / (((^(fi31))) | 1)), F1: 480}
                                                                                                    v29 = v29
                                                                                                    _ = v51
                                                                                                }
                                                                                                v52 := GenS7{F0: v37, F1: 4294967295}
                                                                                                _ = v52
                                                                                            } else {
                                                                                                v1 -= fi31
                                                                                                var v53 string = ((v29 + ("" + v29)) + v29)
                                                                                                var v54 uint64 = 404
                                                                                                if (v21 && true) {
                                                                                                    if true {
                                                                                                        var v55 string = ("\n\t" + "\n\t")
                                                                                                        _ = v55
                                                                                                    } else {
                                                                                                        if len(v39) >= 2 {
                                                                                                            v39 = v39[1:len(v39)]
                                                                                                        }
                                                                                                        v1 += v1
                                                                                                    }
                                                                                                    var v56 string = (v29 + "\n\t")
                                                                                                    v57 := GenS2{F0: (((("hello" + v53) + (v29 + v53)) + "foo bar") < v56), F1: 983, F2: (v5 % (((v5 ^ (v5 | (v5 * 255)))) | 1))}
                                                                                                    var v58 uint = 1
                                                                                                    _ = v56
                                                                                                    _ = v57
                                                                                                    _ = v58
                                                                                                } else {
                                                                                                    var v59 int32 = -470
                                                                                                    v54 = v54
                                                                                                    switch fi31 {
                                                                                                    case 128:
                                                                                                        if (!(!("foo bar" == v29))) {
                                                                                                            if false {
                                                                                                                var v60 uint = 18446744073709551615
                                                                                                                v54 += (v54 * (v54 / (((v54 * (v54 - (v54 + (v54 - (v54 - 2)))))) | 1)))
                                                                                                                switch "foo bar" {
                                                                                                                case "hello":
                                                                                                                    if len(v39) >= 2 {
                                                                                                                        v39 = v39[1:len(v39)]
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    var v61 uint32 = v24
                                                                                                                    if ((v61 / ((v61) | 1)) == v61) {
                                                                                                                        if true {
                                                                                                                            var v62 uint = (^(v60))
                                                                                                                            if false {
                                                                                                                                if v21 {
                                                                                                                                    if ((-32768 <= 0) || (-9223372036854775808 < (v1 << (((^(v60))) & 63)))) {
                                                                                                                                        var v63 uint = (^(v62))
                                                                                                                                        fmt.Println(v53, v1)
                                                                                                                                        _ = v63
                                                                                                                                    }
                                                                                                                                    var v64 int64 = 9223372036854775807
                                                                                                                                    v65 := GenS1{F0: 0, F1: int64(v37)}
                                                                                                                                    _ = v64
                                                                                                                                    _ = v65
                                                                                                                                } else {
                                                                                                                                    switch (^(v1)) {
                                                                                                                                    case -647:
                                                                                                                                        v66 := GenS8{F0: 1000, F1: -1}
                                                                                                                                        _ = v66
                                                                                                                                    case -403:
                                                                                                                                        v21 = (!v21)
                                                                                                                                    case -9223372036854775808:
                                                                                                                                        var v67 float32 = 1.5
                                                                                                                                        _ = v67
                                                                                                                                    }
                                                                                                                                    delete(v3, "")
                                                                                                                                }
                                                                                                                                var v68 int8 = -35
                                                                                                                                v37 = v59
                                                                                                                                _ = v68
                                                                                                                            } else {
                                                                                                                                v69 := GenS9{F0: v21}
                                                                                                                                fmt.Println(v21)
                                                                                                                                _ = v69
                                                                                                                            }
                                                                                                                            if len(v39) > 0 {
                                                                                                                                v39[0] = (fi31 / ((-1) | 1))
                                                                                                                            }
                                                                                                                            _ = v62
                                                                                                                        }
                                                                                                                        v61 -= v61
                                                                                                                    }
                                                                                                                    _ = v61
                                                                                                                }
                                                                                                                v70 := GenS7{F0: v37, F1: genFn0((v60 ^ uint(v24)), (v29 + v53), 32767)}
                                                                                                                _ = v60
                                                                                                                _ = v70
                                                                                                            }
                                                                                                        }
                                                                                                    case 4294967295:
                                                                                                        v39 = append(v39, v1)
                                                                                                    }
                                                                                                    v5 = (v5 << ((4294967295) & 63))
                                                                                                    if v21 {
                                                                                                        switch 127 {
                                                                                                        case -1:
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        if false {
                                                                                                            v32 = v32
                                                                                                            if v21 {
                                                                                                                v53 = v53
                                                                                                                v32 += ((v32 * v32) * (float64(0.0) / (3.14159 / (((v32 + v32) + (v32 * float64(3.14159))) - float64((0.1 / 3.14159))))))
                                                                                                                delete(v2, v24)
                                                                                                                if v21 {
                                                                                                                    v71 := v4["a"]
                                                                                                                    var v72 uint32 = v24
                                                                                                                    _ = v71
                                                                                                                    _ = v72
                                                                                                                } else {
                                                                                                                    var v73 float64 = v32
                                                                                                                    _ = len(v39)
                                                                                                                    var v74 uint32 = (v24 - v24)
                                                                                                                    if false {
                                                                                                                        fmt.Println(v33)
                                                                                                                        if v21 {
                                                                                                                            v75 := GenS4{F0: 1000}
                                                                                                                            fmt.Println(v24)
                                                                                                                            v29 = "a"
                                                                                                                            _ = v75
                                                                                                                        } else {
                                                                                                                            var v76 uint32 = (v74 % (((v24 + (v24 ^ (v74 - 511)))) | 1))
                                                                                                                            v29 = ""
                                                                                                                            if v21 {
                                                                                                                                v1 = (fi30 | (^(fi31)))
                                                                                                                                var v77 float32 = float32((float64(1e10) + ((0.0 + v32) - -2.25)))
                                                                                                                                _ = v77
                                                                                                                            } else {
                                                                                                                                if v21 {
                                                                                                                                    _ = len(v2)
                                                                                                                                } else {
                                                                                                                                    if v21 {
                                                                                                                                        v78 := GenS10{F0: -468, F1: v21, F2: (-(v59))}
                                                                                                                                        if false {
                                                                                                                                            v79 := GenS3{F0: (fi30 & -1), F1: 255}
                                                                                                                                            _ = v79
                                                                                                                                        } else {
                                                                                                                                            var v80 uint8 = 0
                                                                                                                                            _ = v80
                                                                                                                                        }
                                                                                                                                        _ = v78
                                                                                                                                    } else {
                                                                                                                                        v29 = v53
                                                                                                                                        v37 += int32(v33)
                                                                                                                                    }
                                                                                                                                    v74 = v76
                                                                                                                                    if (true && false) {
                                                                                                                                        fmt.Println(v5, v53)
                                                                                                                                        v24 -= (v74 | 42)
                                                                                                                                    } else {
                                                                                                                                        v81 := GenS5{F0: fi31, F1: -1}
                                                                                                                                        _ = v81
                                                                                                                                    }
                                                                                                                                    v82 := GenS2{F0: true, F1: 311, F2: 127}
                                                                                                                                    _ = v82
                                                                                                                                }
                                                                                                                                var v83 uint32 = (^(v24))
                                                                                                                                v21 = v21
                                                                                                                                _ = v83
                                                                                                                            }
                                                                                                                            _ = v76
                                                                                                                        }
                                                                                                                        fmt.Println(v53, fi31, v24)
                                                                                                                        if (v24 > v24) {
                                                                                                                            switch (fi31 >> ((255) & 63)) {
                                                                                                                            case -891:
                                                                                                                                fmt.Println(v29)
                                                                                                                            case 128:
                                                                                                                                var v84 int8 = 100
                                                                                                                                _ = v84
                                                                                                                            case 42:
                                                                                                                                var v85 uint32 = (v74 ^ v24)
                                                                                                                                _ = v85
                                                                                                                            default:
                                                                                                                                if (v21 && v21) {
                                                                                                                                    v5 = 42
                                                                                                                                    fmt.Println(v24, fi30, v24)
                                                                                                                                    v29 = "hello"
                                                                                                                                } else {
                                                                                                                                    v73 = v32
                                                                                                                                    var v86 string = "\n\t"
                                                                                                                                    _ = v86
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v87 := v4["a"]
                                                                                                                            v88 := GenS11{F0: (v74 / ((v24) | 1)), F1: 551, F2: 42}
                                                                                                                            _ = v87
                                                                                                                            _ = v88
                                                                                                                        }
                                                                                                                        var v89 string = ("hello" + "\n\t")
                                                                                                                        _ = v89
                                                                                                                    }
                                                                                                                    _ = v73
                                                                                                                    _ = v74
                                                                                                                }
                                                                                                            } else {
                                                                                                                v90 := GenS12{F0: -1, F1: v29, F2: 0}
                                                                                                                v91 := GenS9{F0: v21}
                                                                                                                if false {
                                                                                                                    var v92 uint8 = 91
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                                _ = v90
                                                                                                                _ = v91
                                                                                                            }
                                                                                                        }
                                                                                                        v24 = (v24 / (((^(v24))) | 1))
                                                                                                    } else {
                                                                                                        v2[(v24 + (v24 / ((v24) | 1)))] = v29
                                                                                                        v93 := GenS13{F0: v1}
                                                                                                        _ = v93
                                                                                                    }
                                                                                                    _ = v59
                                                                                                }
                                                                                                switch v29 {
                                                                                                case "hello":
                                                                                                    var v94 bool = false
                                                                                                    _ = v94
                                                                                                }
                                                                                                var v95 int16 = 32767
                                                                                                v96 := GenS2{F0: v21, F1: v95, F2: (v5 & (v5 | v33))}
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                                _ = v95
                                                                                                _ = v96
                                                                                            }
                                                                                        }
                                                                                        if v21 {
                                                                                            v5 = v5
                                                                                            if (132 <= 0) {
                                                                                                if v21 {
                                                                                                    var v97 int16 = -1
                                                                                                    fmt.Println(v21, v33)
                                                                                                    _ = v97
                                                                                                } else {
                                                                                                    v33 -= 254
                                                                                                    fmt.Println(v24, v29)
                                                                                                    if ((("hello" + (v29 + ("a" + v29))) + v29) >= "\n\t") {
                                                                                                        if false {
                                                                                                            v98 := GenS9{F0: (((v24 / ((v24) | 1)) <= v24) || v21)}
                                                                                                            _ = v98
                                                                                                        } else {
                                                                                                            switch v29 {
                                                                                                            case "\n\t":
                                                                                                                _ = len(v2)
                                                                                                            }
                                                                                                            v21 = false
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v32, v37)
                                                                                                _ = len(v39)
                                                                                                v99 := GenS1{F0: -751, F1: -9223372036854775808}
                                                                                                v1 = (fi31 - -175)
                                                                                                _ = v99
                                                                                            } else {
                                                                                                v5 += v33
                                                                                                v33 = v33
                                                                                                v100 := GenS1{F0: 320, F1: 9223372036854775807}
                                                                                                v101 := GenS4{F0: 0}
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v21 = true
                                                                                }
                                                                                var v102 uint = 65535
                                                                                if true {
                                                                                    switch -881 {
                                                                                    case -163:
                                                                                        v4["hello"] = 2147483647
                                                                                    case -672:
                                                                                        var v103 bool = v21
                                                                                        _ = v103
                                                                                    }
                                                                                    fmt.Println(v32, fi31, v21)
                                                                                    v37 = (v37 ^ (v37 - 0))
                                                                                }
                                                                                _ = v102
                                                                            }
                                                                            _ = v37
                                                                            _ = v39
                                                                        } else {
                                                                            v104 := GenS3{F0: (v1 + -9223372036854775808), F1: 0}
                                                                            if (!v21) {
                                                                                v105 := GenS14{F0: (v24 % ((0) | 1)), F1: (42 >= -1)}
                                                                                if true {
                                                                                    if v21 {
                                                                                        v106 := GenS7{F0: -2147483648, F1: 2}
                                                                                        v107 := []int{(fi30 * (v1 | v1)), fi31, (fi31 / (((fi30 | fi31)) | 1))}
                                                                                        var v108 uint64 = 0
                                                                                        _ = v106
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                    v24 -= 4294967295
                                                                                    if (!(!(32767 == 75))) {
                                                                                        v109 := GenS13{F0: -29}
                                                                                        v33 += 42
                                                                                        _ = v109
                                                                                    } else {
                                                                                        if (0.1 <= float32(float64(-2.25))) {
                                                                                            var v110 int64 = 1
                                                                                            _ = v110
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v33 -= v5
                                                                                _ = v105
                                                                            } else {
                                                                                v33 = (v33 + 172)
                                                                            }
                                                                            _ = v104
                                                                        }
                                                                        v1 += v1
                                                                    } else {
                                                                        if v21 {
                                                                            v111 := GenS15{F0: 2147483647}
                                                                            v112 := v2[v24]
                                                                            _ = v111
                                                                            _ = v112
                                                                        } else {
                                                                            if true {
                                                                                delete(v3, v29)
                                                                                v4[v29] = (v24 + (v24 * 2147483648))
                                                                            }
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            v113 := []int{}
                                                                            _ = v113
                                                                        }
                                                                    }
                                                                } else {
                                                                    if v21 {
                                                                        v24 = v24
                                                                    } else {
                                                                        if false {
                                                                            if false {
                                                                                if v21 {
                                                                                    v114 := GenS9{F0: v21}
                                                                                    if false {
                                                                                        v5 = 37
                                                                                        delete(v3, v29)
                                                                                    } else {
                                                                                        var v115 bool = v21
                                                                                        _ = v115
                                                                                    }
                                                                                    var v116 int16 = 361
                                                                                    _ = v114
                                                                                    _ = v116
                                                                                }
                                                                            }
                                                                            v24 = (v24 ^ 351)
                                                                            switch v29 {
                                                                            case "foo bar":
                                                                                var v117 string = v29
                                                                                _ = v117
                                                                            case "":
                                                                                var v118 string = "hello"
                                                                                _ = v118
                                                                            case "a":
                                                                                var v119 uint8 = 128
                                                                                v120 := GenS11{F0: v24, F1: 42, F2: -769}
                                                                                var v121 uint64 = uint64(v24)
                                                                                var v122 int8 = -128
                                                                                v33 = 1
                                                                                _ = v119
                                                                                _ = v120
                                                                                _ = v121
                                                                                _ = v122
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                v32 = ((v32 - v32) + v32)
                                                                if v21 {
                                                                    v24 += v24
                                                                    switch ("" + "") {
                                                                    case "\n\t":
                                                                        var v123 float32 = 0.1
                                                                        _ = v123
                                                                    case "":
                                                                        v24 -= v24
                                                                        if true {
                                                                            fmt.Println(v32, v21, v24)
                                                                            if (!((v29 + ("" + v29)) >= ((v29 + v29) + "\n\t"))) {
                                                                                if v21 {
                                                                                    v24 += v24
                                                                                    v124 := GenS7{F0: -2147483648, F1: 392}
                                                                                    v125 := []int{fi30}
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                } else {
                                                                                    fmt.Println(v33, v21, v24)
                                                                                }
                                                                                var v126 uint32 = 1
                                                                                _ = v126
                                                                            }
                                                                            v33 = (v5 ^ (^(v5)))
                                                                            var v127 float32 = ((float32(v32) / float32(float64((0.1 / (1.5 / 1.0))))) + (3.14159 - 1e10))
                                                                            v128 := GenS6{F0: 127, F1: 0}
                                                                            var v129 uint8 = (v5 & v5)
                                                                            _ = v127
                                                                            _ = v128
                                                                            _ = v129
                                                                        } else {
                                                                            v32 -= v32
                                                                        }
                                                                        var v130 bool = v21
                                                                        v29 = ""
                                                                        v33 = 128
                                                                        _ = v130
                                                                    default:
                                                                        v131 := GenS16{F0: 42, F1: -1, F2: (!false)}
                                                                        _ = v131
                                                                    }
                                                                    var v132 uint32 = v24
                                                                    _ = v132
                                                                } else {
                                                                    var v133 int64 = 101
                                                                    fmt.Println(v21, v24, v5)
                                                                    fmt.Println(v33, v32, v24)
                                                                    _ = v133
                                                                }
                                                                if (v29 == v29) {
                                                                    var v134 int16 = 0
                                                                    var v135 int32 = 14
                                                                    switch (v29 + ((v29 + "a") + v29)) {
                                                                    case "":
                                                                        if (v1 > (fi30 << ((genFn0(255, (v29 + (v29 + v29)), 42)) & 63))) {
                                                                            fmt.Println(v29, v32)
                                                                            v136 := []int{fi31, fi30, fi30, (v1 ^ (v1 & -9223372036854775808))}
                                                                            _ = v136
                                                                        } else {
                                                                            var v137 int = v1
                                                                            if (v134 <= 465) {
                                                                                fmt.Println(v137, v135, v24)
                                                                            } else {
                                                                                fmt.Println(v1, v134, v1)
                                                                            }
                                                                            _ = v137
                                                                        }
                                                                    case "foo bar":
                                                                        var v138 int8 = 80
                                                                        _ = v138
                                                                    }
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            }
                                                            _ = v33
                                                        }
                                                        _ = fi31
                                                        _ = v32
                                                    }
                                                }
                                            default:
                                                var v139 uint8 = v5
                                                _ = v139
                                            }
                                            for fi140 := 0; fi140 < 2; fi140++ {
                                                if (!v21) {
                                                    v1 -= (v1 & fi140)
                                                    if v21 {
                                                        var v141 uint16 = 1
                                                        var v142 int64 = 0
                                                        var v143 uint = 698
                                                        _ = v141
                                                        _ = v142
                                                        _ = v143
                                                    } else {
                                                        fmt.Println(fi30, fi140)
                                                        var v144 int = (fi30 / (((v1 >> ((2) & 63))) | 1))
                                                        _ = v144
                                                    }
                                                    var v145 float64 = (float64(0.1) * -2.25)
                                                    _ = v145
                                                } else {
                                                    switch fi30 {
                                                    case 127:
                                                        var v146 uint64 = 18446744073709551615
                                                        _ = v146
                                                    case 2147483647:
                                                        v147 := GenS17{F0: 2, F1: 255}
                                                        _ = v147
                                                    case 0:
                                                        v148 := []int{(-(fi30))}
                                                        _ = v148
                                                    default:
                                                        var v149 float32 = -2.25
                                                        _ = v149
                                                    }
                                                }
                                                switch fi140 {
                                                case -2147483648:
                                                    if (0 <= 355) {
                                                        if false {
                                                            switch -667 {
                                                            case 128:
                                                                v150 := GenS9{F0: (42 != 0)}
                                                                _ = v150
                                                            case 4294967295:
                                                                switch (fi30 - (fi30 / ((2147483647) | 1))) {
                                                                case -1:
                                                                    var v151 int8 = 42
                                                                    _ = v151
                                                                case 42:
                                                                    v29 = v29
                                                                default:
                                                                    fmt.Println(v1, v1, fi30)
                                                                }
                                                            case -1:
                                                                if (-525 > -178) {
                                                                    v1 = v1
                                                                    v152 := []int{fi30, v1}
                                                                    v153 := GenS2{F0: (!(((!(((fi140 | v1) == fi30) && (0 != 127))) && (!(v21 && v21))) || v21)), F1: -1, F2: 127}
                                                                    _ = v152
                                                                    _ = v153
                                                                } else {
                                                                    v5 = v5
                                                                    v154 := []int{(fi140 ^ (fi30 + -170)), fi30}
                                                                    _ = v154
                                                                }
                                                            default:
                                                                v155 := []int{v1, v1, (v1 << ((571) & 63))}
                                                                _ = v155
                                                            }
                                                            fmt.Println(v5)
                                                        } else {
                                                            v156 := GenS10{F0: 42, F1: (3.14159 >= (-2.25 + ((0.1 / 0.1) - (1.5 - 0.0)))), F2: 0}
                                                            _ = v156
                                                        }
                                                        if v21 {
                                                            switch v29 {
                                                            case "a":
                                                                var v157 int32 = 42
                                                                _ = v157
                                                            case "hello":
                                                                var v158 float32 = 0.0
                                                                _ = v158
                                                            default:
                                                                switch (-(fi140)) {
                                                                case 32767:
                                                                    var v159 int16 = 1000
                                                                    _ = v159
                                                                case 805:
                                                                    if v21 {
                                                                        if v21 {
                                                                            fmt.Println(fi30, v21, v5)
                                                                        }
                                                                        v24 -= (v24 & v24)
                                                                    } else {
                                                                        fmt.Println(v21, v24, v5)
                                                                    }
                                                                default:
                                                                    v5 = 42
                                                                }
                                                            }
                                                            if ((^(v5)) < (v5 - uint8(v24))) {
                                                                _ = len(v3)
                                                                v1 = 1
                                                                if (810 < 42) {
                                                                    v1 = (v1 - 817)
                                                                    v160 := []int{}
                                                                    v21 = (42 < 51)
                                                                    _ = v160
                                                                }
                                                            } else {
                                                                v161 := []int{v1}
                                                                var v162 int8 = 1
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                            v5 += (v5 - (v5 % (((v5 >> ((genFn0(768, (v29 + (("foo bar" + v29) + v29)), 0)) & 63))) | 1)))
                                                        } else {
                                                            if (2 < 668) {
                                                                v5 += 187
                                                                if (-403 == 42) {
                                                                    var v163 float64 = (0.1 + (0.0 - 3.14159))
                                                                    var v164 string = "hello"
                                                                    if (!v21) {
                                                                        v165 := v3[v164]
                                                                        v21 = v21
                                                                        var v166 bool = v21
                                                                        var v167 int64 = -426
                                                                        switch ((((("foo bar" + (((v164 + v29) + "foo bar") + v29)) + (v29 + v164)) + v29) + (v164 + "hello")) + (v164 + (("hello" + ("a" + v29)) + v29))) {
                                                                        case "":
                                                                            v165 += (v24 | v165)
                                                                        case "hello":
                                                                            v168 := GenS9{F0: false}
                                                                            _ = v168
                                                                        default:
                                                                            fmt.Println(fi140, fi140)
                                                                        }
                                                                        _ = v165
                                                                        _ = v166
                                                                        _ = v167
                                                                    } else {
                                                                        switch (v1 << ((688) & 63)) {
                                                                        case 32768:
                                                                            v164 = v29
                                                                        case -151:
                                                                            v29 = ("hello" + v164)
                                                                        case 2147483647:
                                                                            v163 = (float64((3.14159 * -2.25)) - 0.1)
                                                                        default:
                                                                            v169 := GenS16{F0: 9223372036854775807, F1: 2147483647, F2: v21}
                                                                            _ = v169
                                                                        }
                                                                        fmt.Println(v24, v5, v163)
                                                                        if (42 == 0) {
                                                                            v170 := GenS9{F0: (true || v21)}
                                                                            v171 := []int{v1, fi30}
                                                                            v171 = append(v171, v1)
                                                                            _ = v170
                                                                            _ = v171
                                                                        } else {
                                                                            fmt.Println(fi140, v1, v164)
                                                                            v24 += v24
                                                                            v1 = (fi30 * (fi30 / (((fi140 % ((-1) | 1))) | 1)))
                                                                        }
                                                                    }
                                                                    var v172 uint32 = 2
                                                                    var v173 uint16 = 519
                                                                    _ = v163
                                                                    _ = v164
                                                                    _ = v172
                                                                    _ = v173
                                                                } else {
                                                                    var v174 int32 = 930
                                                                    _ = v174
                                                                }
                                                                switch v29 {
                                                                case "foo bar":
                                                                    v175 := GenS18{F0: "", F1: v5}
                                                                    _ = v175
                                                                case "":
                                                                    if (42 < uint64(v24)) {
                                                                        var v176 uint64 = 0
                                                                        _ = v176
                                                                    }
                                                                }
                                                                var v177 bool = (275 > 382)
                                                                v5 += v5
                                                                v178 := []int{(fi30 % (((v1 >> ((255) & 63))) | 1)), 127, fi30}
                                                                if (32767 == -949) {
                                                                    switch v29 {
                                                                    case "foo bar":
                                                                        v179 := GenS11{F0: v24, F1: 9223372036854775807, F2: 42}
                                                                        _ = v179
                                                                    }
                                                                    var v180 float32 = 1.5
                                                                    _ = v180
                                                                } else {
                                                                    v1 += 1
                                                                }
                                                                _ = v177
                                                                _ = v178
                                                            }
                                                            v24 = (^(v24))
                                                            switch fi140 {
                                                            case 127:
                                                                var v181 uint32 = 2147483648
                                                                fmt.Println(fi140)
                                                                switch (fi140 * (fi30 | 0)) {
                                                                case 811:
                                                                    v182 := GenS16{F0: 42, F1: -161, F2: v21}
                                                                    _ = v182
                                                                case 1:
                                                                    var v183 float64 = (3.14159 + -2.25)
                                                                    _ = v183
                                                                }
                                                                _ = v181
                                                            default:
                                                                v29 = v29
                                                                v184 := GenS14{F0: v24, F1: (1 != 257)}
                                                                switch fi30 {
                                                                case 32768:
                                                                    v185 := []int{313, (v1 - (^(fi140))), 9223372036854775807, v1}
                                                                    _ = v185
                                                                case 128:
                                                                    var v186 uint32 = v24
                                                                    _ = v186
                                                                default:
                                                                    fmt.Println(v29, v21, fi30)
                                                                }
                                                                _ = v184
                                                            }
                                                            v29 = (v29 + "foo bar")
                                                        }
                                                    }
                                                case 747:
                                                    v1 = -771
                                                }
                                                v187 := []int{(fi140 - -416), (fi30 | v1), v1}
                                                _ = fi140
                                                _ = v187
                                            }
                                            _ = fi30
                                        }
                                        var v188 bool = false
                                        _ = v188
                                    }
                                    var v189 float32 = 1e10
                                    var v190 float32 = 1.5
                                    _ = v189
                                    _ = v190
                                } else {
                                    if (!v21) {
                                        if true {
                                            v191 := GenS12{F0: -747, F1: v29, F2: 65535}
                                            for fi192 := 0; fi192 < 7; fi192++ {
                                                switch ((("foo bar" + (v29 + v29)) + "foo bar") + ((v29 + "a") + (((v29 + (v29 + (v29 + v29))) + v29) + v29))) {
                                                case "hello":
                                                    var v193 string = v29
                                                    _ = v193
                                                case "":
                                                    for fi194 := 0; fi194 < 13; fi194++ {
                                                        v195 := GenS10{F0: 42, F1: v21, F2: -1}
                                                        v196 := GenS9{F0: true}
                                                        var v197 string = v29
                                                        v198 := []int{}
                                                        var v199 int = (fi192 * fi192)
                                                        _ = len(v4)
                                                        switch (-(fi192)) {
                                                        case -2147483648:
                                                            if (0.0 > 0.0) {
                                                                v200 := GenS19{F0: (v197 + (("a" + ((v197 + "hello") + "\n\t")) + "hello")), F1: 4294967295, F2: 1}
                                                                v201 := GenS16{F0: 9223372036854775807, F1: 924, F2: v21}
                                                                _ = v200
                                                                _ = v201
                                                            }
                                                        default:
                                                            v1 = v1
                                                        }
                                                        _ = fi194
                                                        _ = v195
                                                        _ = v196
                                                        _ = v197
                                                        _ = v198
                                                        _ = v199
                                                    }
                                                default:
                                                    for fi202 := 0; fi202 < 15; fi202++ {
                                                        fmt.Println(v5, v24, v5)
                                                        v1 += v1
                                                        _ = fi202
                                                    }
                                                }
                                                var v203 int64 = 9223372036854775807
                                                if true {
                                                    v204 := v4[""]
                                                    _ = v204
                                                }
                                                _ = fi192
                                                _ = v203
                                            }
                                            for fi205 := 0; fi205 < 9; fi205++ {
                                                switch (fi205 & (fi205 - fi205)) {
                                                case 1:
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v206 uint = 879
                                                        _ = v206
                                                    case "hello":
                                                        fmt.Println(fi205, v5)
                                                    case "":
                                                        for fi207 := 0; fi207 < 20; fi207++ {
                                                            var v208 bool = ((4294967295 == 839) || (213 >= v5))
                                                            _ = fi207
                                                            _ = v208
                                                        }
                                                    default:
                                                        for fi209 := 0; fi209 < 7; fi209++ {
                                                            var v210 uint64 = 2
                                                            v211 := GenS14{F0: (^(v24)), F1: v21}
                                                            v5 = uint8(v210)
                                                            _ = fi209
                                                            _ = v210
                                                            _ = v211
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    switch v29 {
                                                    case "\n\t":
                                                        fmt.Println(v29, v5)
                                                    default:
                                                        var v212 uint64 = 1
                                                        _ = v212
                                                    }
                                                }
                                                var v213 int8 = -128
                                                _ = fi205
                                                _ = v213
                                            }
                                            _ = v191
                                        }
                                        for fi214 := 0; fi214 < 2; fi214++ {
                                            v215 := v4["foo bar"]
                                            _ = fi214
                                            _ = v215
                                        }
                                        switch "\n\t" {
                                        case "\n\t":
                                            v216 := []int{}
                                            _ = v216
                                        case "foo bar":
                                            var v217 string = "hello"
                                            _ = v217
                                        case "":
                                            v218 := GenS7{F0: 2147483647, F1: genFn0(213, v29, 21)}
                                            _ = v218
                                        }
                                    } else {
                                        v29 = "\n\t"
                                        v29 = v29
                                    }
                                }
                                _ = v29
                            }
                        } else {
                            v21 = (int8(v5) > -128)
                        }
                        _ = v24
                    }
                }
            case 4294967295:
                fmt.Println(v1, v1)
            case -9223372036854775808:
                v219 := GenS4{F0: -503}
                _ = v219
            default:
                v5 -= 99
            }
        }
        _ = v22
    } else {
        for fi220 := 0; fi220 < 3; fi220++ {
            for fi221 := 0; fi221 < 6; fi221++ {
                v222 := []int{fi220}
                _ = fi221
                _ = v222
            }
            _ = fi220
        }
        if (("hello" == "a") || false) {
            switch "foo bar" {
            case "foo bar":
                switch "a" {
                case "hello":
                    var v223 int16 = 32767
                    _ = v223
                default:
                    v5 = v5
                }
            case "":
                v224 := GenS20{F0: -1}
                _ = v224
            }
            fmt.Println(v1)
        }
    }
    fmt.Println(v1)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
