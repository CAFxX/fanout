package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 bool }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool; F1 uint8 }
func genFn0(p0 bool, p1 bool) uint8 {
    switch -660 {
    case 32768:
        p1 = p0
    case -9223372036854775808:
        v7 := GenS1{F0: 2}
        _ = v7
    case 678:
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(fi8)
            if p0 {
                v9 := GenS2{F0: 139, F1: 2}
                v10 := GenS2{F0: 247, F1: 741}
                switch ("" + "hello") {
                case "foo bar":
                    switch (fi8 | 128) {
                    case -761:
                        fmt.Println(p0, p0, p0)
                    default:
                        v11 := GenS1{F0: 677}
                        _ = v11
                    }
                }
                _ = v10
                _ = v9
            } else {
                v12 := GenS1{F0: 42}
                _ = v12
            }
            p1 = (!(0 >= 2))
            _ = fi8
        }
    }
    var v13 int16 = 0
    _ = p0
    _ = p1
    _ = v13
    return 2
}
func genFn1() bool {
    var v14 uint = 2
    v14 = (^(v14))
    v14 = v14
    if false {
        switch -1 {
        case -9223372036854775808:
            var v15 uint8 = genFn0((false || (451 > 358)), (874 > 456))
            _ = v15
        case -891:
            var v16 uint16 = 2
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v14 -= 255
            _ = fi17
        }
    } else {
        v18 := []int{-470, 128}
        _ = v18
    }
    v14 += v14
    _ = v14
    return (false || (true && (866 > 42)))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 uint = 1
    var v4 int = 128
    var v5 uint32 = 2
    var v6 int64 = -13
    var v19 float64 = (((0.0 + (((1.5 * 0.0) / float64(float32(float64(3.14159)))) + 3.14159)) + 0.1) - 0.1)
    fmt.Println(v4)
    var v20 int = v4
    v21 := GenS2{F0: 42, F1: (v5 + (^(v5)))}
    var v22 uint8 = 2
    v23 := GenS3{F0: (("foo bar" + "foo bar") + "hello")}
    if false {
        for fi24 := 0; fi24 < 7; fi24++ {
            switch (^(v4)) {
            case 0:
                v25 := GenS2{F0: (v22 & v22), F1: v5}
                _ = v25
            default:
                if (v19 >= 3.14159) {
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v27 := GenS1{F0: 65535}
                        _ = fi26
                        _ = v27
                    }
                }
            }
            _ = fi24
        }
    }
    var v28 int32 = 495
    if false {
        switch v4 {
        case 128:
            if false {
                for fi29 := 0; fi29 < 6; fi29++ {
                    v30 := []int{v4, (v4 - v4)}
                    var v31 int = -2147483648
                    for fi32 := 0; fi32 < 1; fi32++ {
                        if false {
                            if true {
                                fmt.Println(v6, v28)
                                v3 = 18446744073709551615
                                switch (fi32 >> ((v3) & 63)) {
                                case -359:
                                    fmt.Println(v20, v4)
                                case 1:
                                    v31 -= (v31 + v31)
                                default:
                                    v33 := GenS1{F0: 1}
                                    _ = v33
                                }
                                v34 := GenS4{F0: (v22 & (v22 | v22))}
                                _ = v34
                            } else {
                                v6 = (^(v6))
                                fmt.Println(v28, fi32, v28)
                                switch -250 {
                                case 844:
                                    v28 = 42
                                case -9223372036854775808:
                                    if (0.1 <= v19) {
                                        fmt.Println(fi32)
                                    } else {
                                        v20 = -168
                                    }
                                case -1:
                                    var v35 int32 = (v28 - -2147483648)
                                    _ = v35
                                }
                            }
                            switch "hello" {
                            case "a":
                                v3 += 90
                            }
                        } else {
                            v6 -= v6
                            v36 := GenS1{F0: 1}
                            switch int(v6) {
                            case 245:
                                switch (-(v20)) {
                                case -9223372036854775808:
                                    fmt.Println(v20)
                                case 326:
                                    var v37 int16 = 42
                                    _ = v37
                                case 127:
                                    v38 := GenS5{F0: -128, F1: ("a" + ("foo bar" + "a")), F2: false}
                                    _ = v38
                                default:
                                    switch ("a" + ("\n\t" + "a")) {
                                    case "hello":
                                        v20 += fi32
                                    case "":
                                        fmt.Println(fi32)
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            if (false || (true && false)) {
                                                v39 := GenS6{F0: true, F1: 1}
                                                v4 = (^(v31))
                                                if true {
                                                    var v40 bool = (!false)
                                                    v41 := v1[4294967295]
                                                    if ((^(v5)) != v5) {
                                                        fmt.Println(v28)
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    var v42 int32 = v28
                                                    v43 := GenS5{F0: -128, F1: ("a" + "hello"), F2: ((v22 / (((v22 >> (((v3 >> ((v3) & 63))) & 63))) | 1)) <= 127)}
                                                    _ = len(v2)
                                                    _ = v42
                                                    _ = v43
                                                }
                                                fmt.Println(v20, fi32)
                                                _ = v39
                                            } else {
                                                if false {
                                                    v44 := v2["a"]
                                                    v45 := GenS7{F0: (^(v22)), F1: (v22 * (^(v22))), F2: -820}
                                                    v22 = 255
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                        case "\n\t":
                                            v2[(("foo bar" + "a") + "\n\t")] = false
                                        case "":
                                            var v46 string = ("hello" + "foo bar")
                                            _ = v46
                                        default:
                                            fmt.Println(v20, fi32)
                                        }
                                    }
                                }
                                if len(v30) > 0 {
                                    v30[0] = v20
                                }
                                v22 = (v22 * (v22 * v22))
                            case 596:
                                v3 = 255
                            case -101:
                                fmt.Println(v20, v5)
                            }
                            v6 = (v6 ^ (v6 % ((v6) | 1)))
                            _ = v36
                        }
                        v20 = -9223372036854775808
                        fmt.Println(v4, v22, fi29)
                        _ = fi32
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            } else {
                if (false && ((v5 % (((v5 % (((^(v5))) | 1))) | 1)) != v5)) {
                    _ = len(v1)
                    fmt.Println(v20, v20)
                }
                v6 -= 1
            }
        case 736:
            v47 := GenS2{F0: (v22 | v22), F1: (v5 % ((1) | 1))}
            _ = v47
        default:
            var v48 uint32 = 2147483647
            _ = v48
        }
    } else {
        v49 := []int{(v20 | (v4 >> ((v3) & 63))), -517}
        if false {
            switch (v4 / ((v20) | 1)) {
            case 127:
                _ = len(v2)
            case 829:
                if len(v49) > 0 {
                    v49[0] = (v20 + (v4 / ((42) | 1)))
                }
            case -1:
                for fi50 := 0; fi50 < 3; fi50++ {
                    v5 = (v5 * v5)
                    for fi51 := 0; fi51 < 9; fi51++ {
                        v52 := GenS1{F0: 255}
                        fmt.Println(v4, v28)
                        _ = len(v49)
                        _ = fi51
                        _ = v52
                    }
                    if (255 != (v22 >> ((957) & 63))) {
                        if (fi50 < 9223372036854775807) {
                            fmt.Println(fi50, v28)
                            v6 += -194
                            for fi53 := 0; fi53 < 9; fi53++ {
                                v19 -= float64((0.1 / float32(v19)))
                                if (!true) {
                                    var v54 uint8 = (v22 & (^(v22)))
                                    v55 := GenS3{F0: ("a" + ("hello" + "foo bar"))}
                                    switch 857 {
                                    case -421:
                                        v2["\n\t"] = false
                                    case -9223372036854775808:
                                        var v56 int = (fi50 << (((v3 * (v3 % ((4294967295) | 1)))) & 63))
                                        switch ("foo bar" + "\n\t") {
                                        case "foo bar":
                                            fmt.Println(v28, v19, v56)
                                        case "hello":
                                            v4 = (^(v4))
                                        }
                                        switch v4 {
                                        case 42:
                                            switch (-(v20)) {
                                            case 105:
                                                v57 := GenS8{F0: 0, F1: v28}
                                                _ = v57
                                            case 42:
                                                v20 -= v20
                                            case -9223372036854775808:
                                                v6 = (v6 - v6)
                                            }
                                        case 1:
                                            if (((v22 | v22) < genFn0((!(786 >= 88)), (false && true))) && false) {
                                                if ((1.5 != 0.0) || (!false)) {
                                                    v58 := GenS5{F0: -24, F1: "hello", F2: true}
                                                    var v59 uint8 = (^(v22))
                                                    _ = v58
                                                    _ = v59
                                                }
                                            }
                                        }
                                        _ = v56
                                    case 127:
                                        switch "foo bar" {
                                        case "foo bar":
                                            _ = len(v49)
                                        case "\n\t":
                                            v60 := GenS9{F0: 65535}
                                            _ = v60
                                        }
                                    default:
                                        v61 := GenS10{F0: v3}
                                        _ = v61
                                    }
                                    _ = v54
                                    _ = v55
                                } else {
                                    var v62 uint64 = 2
                                    if ((!true) && (!false)) {
                                        v63 := GenS11{F0: 32767}
                                        fmt.Println(fi50)
                                        switch -455 {
                                        case 2147483647:
                                            if true {
                                                v62 -= 18446744073709551615
                                                var v64 bool = false
                                                v65 := GenS7{F0: uint8(v5), F1: (^(v22)), F2: (v28 & int32(v3))}
                                                _ = v64
                                                _ = v65
                                            } else {
                                                var v66 uint = v3
                                                _ = v66
                                            }
                                        }
                                        _ = v63
                                    }
                                    _ = v62
                                }
                                var v67 int64 = v6
                                _ = fi53
                                _ = v67
                            }
                        }
                        if false {
                            fmt.Println(v3, v20, v19)
                            for fi68 := 0; fi68 < 1; fi68++ {
                                v28 += -1
                                _ = fi68
                            }
                        }
                    }
                    _ = fi50
                }
            }
            if (1 == 415) {
                v69 := GenS12{F0: (("a" + "") + "a")}
                if (!(255 > 267)) {
                    if false {
                        if (v5 > v5) {
                            for fi70 := 0; fi70 < 5; fi70++ {
                                var v71 uint16 = 2
                                var v72 bool = (!(true || (!true)))
                                _ = fi70
                                _ = v71
                                _ = v72
                            }
                            var v73 uint8 = (v22 + 218)
                            _ = v73
                        } else {
                            v4 -= (v4 | 42)
                            for fi74 := 0; fi74 < 6; fi74++ {
                                v75 := GenS8{F0: 2, F1: (v28 << ((140) & 63))}
                                _ = fi74
                                _ = v75
                            }
                            var v76 uint32 = (v5 & uint32(v3))
                            _ = v76
                        }
                    }
                    var v77 uint64 = 2
                    var v78 float64 = (v19 * 1.5)
                    _ = v77
                    _ = v78
                }
                _ = v69
            }
        } else {
            var v79 uint16 = uint16(v5)
            _ = v79
        }
        _ = v49
    }
    var v80 int64 = v6
    fmt.Println(v5, v22)
    switch "a" {
    case "\n\t":
        v81 := GenS13{F0: (v80 != (v6 + (v6 ^ (^(v80))))), F1: (v22 % (((v22 << (((v3 % ((797) | 1))) & 63))) | 1))}
        _ = v81
    case "a":
        for fi82 := 0; fi82 < 5; fi82++ {
            for fi83 := 0; fi83 < 14; fi83++ {
                if false {
                    v84 := GenS13{F0: false, F1: genFn0(false, false)}
                    var v85 int64 = (v6 + v80)
                    _ = v84
                    _ = v85
                } else {
                    var v86 int16 = 1
                    v3 += (v3 / ((v3) | 1))
                    switch "\n\t" {
                    case "a":
                        var v87 uint = (v3 & v3)
                        _ = v87
                    }
                    _ = v86
                }
                _ = fi83
            }
            _ = fi82
        }
    default:
        v88 := []int{(v20 + int(v28)), v4, int(v6)}
        _ = v88
    }
    fmt.Println(v80)
    v22 -= (v22 & genFn0(true, (!(0 > (v4 >> ((v3) & 63))))))
    var v89 int = 1
    delete(v2, ("hello" + "hello"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v89
}
