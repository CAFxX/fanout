package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint32, p1 uint) uint8 {
    v4 := GenS1{F0: 286, F1: -664}
    v5 := GenS1{F0: 18446744073709551615, F1: -1}
    v6 := []int{42, 156, 4294967295}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1() string {
    var v7 string = "a"
    switch ("\n\t" + ((v7 + v7) + "a")) {
    case "\n\t":
        v7 = ""
    case "foo bar":
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (-(fi8)) {
            case 0:
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 uint32 = 65535
                    var v11 float64 = 1.5
                    v10 = v10
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            if (!(!false)) {
                if false {
                    v12 := []int{fi8}
                    for fi13 := 0; fi13 < 18; fi13++ {
                        fmt.Println(fi8, fi8)
                        var v14 float32 = 1.5
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v7 = v7
                        v16 := GenS1{F0: 1, F1: 2147483647}
                        var v17 int = fi15
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = v12
                }
                var v18 uint = 42
                _ = v18
            }
            _ = fi8
        }
    case "":
        v7 = v7
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            v7 = v7
            switch v7 {
            case "\n\t":
                v21 := GenS1{F0: 948, F1: -2147483648}
                _ = v21
            case "a":
                var v22 uint = 65535
                _ = v22
            }
            v23 := []int{}
            _ = fi20
            _ = v23
        }
        var v24 float32 = 0.1
        _ = fi19
        _ = v24
    }
    switch 708 {
    case -9223372036854775808:
        for fi25 := 0; fi25 < 9; fi25++ {
            v7 = v7
            _ = fi25
        }
    default:
        switch 9223372036854775807 {
        case 0:
            var v26 int16 = -32768
            _ = v26
        default:
            fmt.Println(v7)
        }
    }
    v27 := []int{-63, 152}
    _ = v27
    _ = v7
    return v7
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    switch "a" {
    case "hello":
        if (!(65535 >= uint32(v1))) {
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int = fi28
                _ = fi28
                _ = v29
            }
            var v30 int16 = (v3 % ((32767) | 1))
            _ = v30
        } else {
            for fi31 := 0; fi31 < 12; fi31++ {
                v3 = (-(v3))
                fmt.Println(v3, v3, v1)
                _ = fi31
            }
        }
    case "\n\t":
        v1 = 558
    default:
        if true {
            v32 := v2[51]
            _ = v32
        } else {
            if (!(!(!false))) {
                v3 -= v3
            }
        }
    }
    if false {
        if true {
            var v33 bool = true
            _ = v33
        }
        var v34 string = "\n\t"
        _ = v34
    } else {
        v1 += v1
        for fi35 := 0; fi35 < 5; fi35++ {
            for fi36 := 0; fi36 < 10; fi36++ {
                v3 = 32767
                v37 := GenS2{F0: 42, F1: -128, F2: 1}
                v1 -= (^(v1))
                v1 -= (^(v1))
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                fmt.Println(v1, fi35, fi35)
                _ = fi38
            }
            if false {
                fmt.Println(fi35)
                v3 = (^(v3))
                switch ("\n\t" + (("" + ("" + "\n\t")) + "\n\t")) {
                case "a":
                    switch "foo bar" {
                    case "hello":
                        v3 += v3
                    case "a":
                        switch "a" {
                        case "":
                            v1 = (v1 * v1)
                        case "foo bar":
                            var v39 bool = true
                            _ = v39
                        case "a":
                            if false {
                                v40 := GenS1{F0: 42, F1: 2147483647}
                                for fi41 := 0; fi41 < 4; fi41++ {
                                    if (0 > v1) {
                                        v3 -= (v3 - v3)
                                        var v42 uint8 = 128
                                        var v43 uint64 = 1
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v1 = (v1 << ((42) & 63))
                                    }
                                    delete(v2, 0)
                                    var v44 int32 = -2147483648
                                    switch fi35 {
                                    case 32768:
                                        v1 = (v1 ^ (^(v1)))
                                    case 42:
                                        fmt.Println(v3)
                                    case 112:
                                        v2[2147483648] = "foo bar"
                                    }
                                    _ = fi41
                                    _ = v44
                                }
                                v2[65535] = ("a" + "a")
                                _ = v40
                            } else {
                                var v45 float32 = -2.25
                                var v46 int32 = -391
                                fmt.Println(fi35, v1)
                                _ = v45
                                _ = v46
                            }
                        default:
                            switch (fi35 / (((fi35 - (-(fi35)))) | 1)) {
                            case 305:
                                v1 = v1
                            default:
                                fmt.Println(v1, v3, v3)
                            }
                        }
                    default:
                        v1 -= (v1 / ((v1) | 1))
                    }
                default:
                    v3 -= (-(v3))
                }
            } else {
                v47 := GenS2{F0: 2, F1: -1, F2: 2147483648}
                var v48 uint8 = 128
                _ = v47
                _ = v48
            }
            _ = fi35
        }
    }
    if false {
        for fi49 := 0; fi49 < 17; fi49++ {
            var v50 uint64 = 1
            var v51 int16 = v3
            _ = fi49
            _ = v50
            _ = v51
        }
        v1 += (v1 % (((v1 * 818)) | 1))
        v3 += (^(v3))
    } else {
        v3 = 510
        var v52 bool = true
        v3 = (v3 ^ (v3 << ((18446744073709551615) & 63)))
        _ = v52
    }
    if ((0.0 <= 0.1) && (32767 < 32767)) {
        if true {
            v53 := []int{}
            v54 := GenS3{F0: 0, F1: 0}
            _ = v53
            _ = v54
        }
    }
    if (!((v1 * (v1 >> ((4294967295) & 63))) != v1)) {
        var v55 uint64 = 302
        var v56 uint8 = 180
        _ = v55
        _ = v56
    }
    v57 := []int{9223372036854775807, 0, -931}
    v58 := GenS1{F0: 2, F1: 2147483647}
    var v59 int16 = 32767
    v60 := GenS3{F0: -1, F1: 4294967295}
    v3 = v3
    v61 := GenS2{F0: (v1 * (v1 | (v1 % ((v1) | 1)))), F1: 98, F2: 2}
    v1 = 2
    var v62 int = -404
    switch "\n\t" {
    case "a":
        if (!(1e10 >= -2.25)) {
            if false {
                switch v62 {
                case 9223372036854775807:
                    switch int(v59) {
                    case 0:
                        var v63 float64 = (float64(1.5) / 1.0)
                        _ = v63
                    case -1:
                        v57 = append(v57, v62)
                    case -536:
                        fmt.Println(v3, v59)
                    }
                case 2147483647:
                    var v64 string = "\n\t"
                    _ = v64
                }
                v65 := GenS4{F0: (v62 & v62), F1: 1}
                switch 2147483647 {
                case 32768:
                    for fi66 := 0; fi66 < 1; fi66++ {
                        if (false && false) {
                            v1 -= (^(v1))
                        }
                        var v67 string = ""
                        var v68 int32 = 2147483647
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
                _ = v65
            } else {
                for fi69 := 0; fi69 < 18; fi69++ {
                    v3 = 0
                    _ = fi69
                }
                v2[0] = ("a" + "hello")
                if (!false) {
                    for fi70 := 0; fi70 < 20; fi70++ {
                        var v71 int = fi70
                        for fi72 := 0; fi72 < 17; fi72++ {
                            v71 = v62
                            _ = fi72
                        }
                        _ = fi70
                        _ = v71
                    }
                    switch ("" + "hello") {
                    case "\n\t":
                        if (194 > 42) {
                            fmt.Println(v3)
                            v73 := GenS5{F0: 0, F1: (v1 + (^(v1))), F2: (v3 | (^(v59)))}
                            if len(v57) >= 2 {
                                v57 = v57[1:len(v57)]
                            }
                            _ = v73
                        }
                    case "hello":
                        if (742 != 1) {
                            v62 = v62
                            if true {
                                if (-719 < 269) {
                                    v3 += 0
                                    switch v62 {
                                    case 508:
                                        var v74 int = (v62 & v62)
                                        _ = v74
                                    default:
                                        v75 := GenS6{F0: int64(v62)}
                                        _ = v75
                                    }
                                } else {
                                    switch (-(v62)) {
                                    case 42:
                                        fmt.Println(v1, v59)
                                    case 0:
                                        var v76 uint8 = 1
                                        _ = v76
                                    case 32768:
                                        v59 = (v59 % (((v59 % ((-32768) | 1))) | 1))
                                    }
                                    var v77 uint64 = 2
                                    v78 := GenS7{F0: 1, F1: 59}
                                    _ = v77
                                    _ = v78
                                }
                                var v79 uint64 = 808
                                var v80 uint16 = (v1 / ((42) | 1))
                                _ = v79
                                _ = v80
                            }
                        } else {
                            v2[916] = ((("\n\t" + "\n\t") + "a") + "a")
                            v81 := GenS1{F0: 1, F1: 0}
                            _ = v81
                        }
                    }
                    var v82 uint16 = 42
                    _ = v82
                } else {
                    if len(v57) >= 2 {
                        v57 = v57[1:len(v57)]
                    }
                }
            }
            v59 += (v59 / ((v3) | 1))
            fmt.Println(v3)
        } else {
            if (false && false) {
                for fi83 := 0; fi83 < 1; fi83++ {
                    var v84 uint16 = 2
                    _ = fi83
                    _ = v84
                }
                v85 := GenS4{F0: v62, F1: -2147483648}
                _ = v85
            }
        }
    case "\n\t":
        if ((!(42 != 0)) || (true && (false && (-1 <= 42)))) {
            var v86 string = "a"
            v3 += 0
            _ = v86
        }
    }
    v87 := v2[268]
    _ = len(v57)
    v88 := GenS5{F0: 2, F1: 716, F2: (v3 - (v59 & (v3 << ((4294967295) & 63))))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v87
    _ = v88
}
