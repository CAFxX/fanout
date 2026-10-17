package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int32; F1 uint; F2 int16 }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 bool; F1 string; F2 uint16 }
type GenS11 struct { F0 int32 }
func genFn0(p0 bool, p1 int64) uint {
    var v4 uint8 = 77
    v4 += 127
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 int = -1
    var v2 int64 = -534
    v3 := make(map[int]uint32)
    for fi5 := 0; fi5 < 19; fi5++ {
        v2 = (v2 / ((v2) | 1))
        v6 := GenS1{F0: 42, F1: 1}
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 1}
    if true {
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                if false {
                    fmt.Println(v2, v1)
                    for fi9 := 0; fi9 < 3; fi9++ {
                        fmt.Println(fi9, v1, fi9)
                        _ = fi9
                    }
                } else {
                    v10 := []int{(fi8 ^ v1), 611, v1, (v1 & v1)}
                    for fi11 := 0; fi11 < 12; fi11++ {
                        v12 := v3[fi8]
                        var v13 bool = false
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                }
                _ = fi8
            }
        default:
            v14 := GenS3{F0: 42, F1: -576}
            _ = v14
        }
        var v15 string = ("hello" + "a")
        _ = v15
    }
    var v16 uint64 = 735
    if true {
        if (!(!false)) {
            v17 := []int{(v1 >> ((65535) & 63)), (v1 | -94), v1}
            v18 := GenS1{F0: 0, F1: 65535}
            switch v1 {
            case 42:
                switch (v1 ^ (v1 / (((v1 << ((255) & 63))) | 1))) {
                case -565:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        fmt.Println(v2)
                        if ((!((v16 * 870) <= v16)) && (42 != -424)) {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                switch "hello" {
                                case "\n\t":
                                    v21 := GenS4{F0: int64(fi20)}
                                    _ = v21
                                default:
                                    if len(v17) > 0 {
                                        v17[0] = -781
                                    }
                                }
                                _ = fi20
                            }
                        } else {
                            switch "hello" {
                            case "":
                                if (1 <= 83) {
                                    v17 = append(v17, 32767)
                                    fmt.Println(fi19, v2)
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                }
                            case "hello":
                                v22 := GenS3{F0: 2, F1: int32(v16)}
                                _ = v22
                            default:
                                v23 := GenS5{F0: (v16 ^ (v16 / (((v16 % ((v16) | 1))) | 1))), F1: 0}
                                _ = v23
                            }
                            v24 := GenS6{F0: 0, F1: 9223372036854775807}
                            _ = v24
                        }
                        _ = fi19
                    }
                }
            }
            _ = v17
            _ = v18
        } else {
            v1 = -817
            if true {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 float32 = 0.1
                    _ = fi25
                    _ = v26
                }
                v27 := GenS3{F0: 227, F1: -2147483648}
                if true {
                    if ((v2 | -1) < 117) {
                        v28 := GenS6{F0: 255, F1: (v2 / ((v2) | 1))}
                        if true {
                            v29 := GenS5{F0: v16, F1: 969}
                            for fi30 := 0; fi30 < 19; fi30++ {
                                v2 = (v2 >> ((0) & 63))
                                _ = fi30
                            }
                            v31 := v3[v1]
                            _ = v29
                            _ = v31
                        }
                        var v32 int64 = v2
                        _ = v28
                        _ = v32
                    }
                }
                _ = v27
            } else {
                switch ("\n\t" + ((("a" + "hello") + "hello") + "")) {
                case "a":
                    var v33 bool = (1 == 42)
                    _ = v33
                case "\n\t":
                    for fi34 := 0; fi34 < 5; fi34++ {
                        v35 := []int{(fi34 << ((898) & 63))}
                        _ = fi34
                        _ = v35
                    }
                case "foo bar":
                    for fi36 := 0; fi36 < 8; fi36++ {
                        for fi37 := 0; fi37 < 13; fi37++ {
                            if true {
                                fmt.Println(fi36, v16)
                                v38 := []int{(-(fi36))}
                                _ = v38
                            } else {
                                if true {
                                    if ((v1 / (((fi37 - 1)) | 1)) >= -2147483648) {
                                        var v39 uint64 = 503
                                        if false {
                                            var v40 int16 = -1
                                            _ = v40
                                        } else {
                                            var v41 uint8 = 83
                                            if true {
                                                if (((1.5 - (3.14159 + 3.14159)) != -2.25) || false) {
                                                    fmt.Println(v41, v1)
                                                } else {
                                                    if true {
                                                        if (true || (genFn0(true, 9223372036854775807) == 190)) {
                                                            if true {
                                                                v42 := []int{fi36, (fi36 - fi36), (v1 % ((153) | 1))}
                                                                _ = v42
                                                            } else {
                                                                var v43 float32 = 1.5
                                                                v44 := GenS7{F0: -1, F1: genFn0(((v39 | v16) >= v39), (v2 / ((v2) | 1))), F2: -1}
                                                                _ = v43
                                                                _ = v44
                                                            }
                                                            var v45 float64 = (-2.25 + 0.1)
                                                            v2 += v2
                                                            _ = v45
                                                        }
                                                        fmt.Println(fi36, fi36)
                                                        v46 := v3[v1]
                                                        _ = v46
                                                    } else {
                                                        var v47 int64 = 42
                                                        v39 += 1
                                                        v41 += 2
                                                        v48 := []int{v1, fi37, v1, 278}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                }
                                                if true {
                                                    v49 := []int{}
                                                    _ = v49
                                                }
                                            }
                                            switch (^(fi37)) {
                                            case 2147483647:
                                                var v50 uint8 = v41
                                                _ = v50
                                            default:
                                                var v51 int8 = 42
                                                _ = v51
                                            }
                                            _ = v41
                                        }
                                        _ = v39
                                    } else {
                                        fmt.Println(v1, v16)
                                        v1 += fi37
                                        var v52 int8 = -36
                                        _ = v52
                                    }
                                    var v53 uint16 = 32767
                                    v3[(-(fi36))] = 272
                                    _ = v53
                                }
                                v54 := GenS2{F0: 0}
                                v2 -= (-(v2))
                                _ = v54
                            }
                            var v55 uint16 = 2
                            _ = fi37
                            _ = v55
                        }
                        var v56 int = -512
                        _ = fi36
                        _ = v56
                    }
                default:
                    v57 := GenS1{F0: -119, F1: genFn0((127 <= -1), (v2 >> ((255) & 63)))}
                    _ = v57
                }
                switch v1 {
                case 1:
                    v58 := GenS3{F0: 2, F1: -543}
                    _ = v58
                case 172:
                    var v59 uint = 4294967295
                    _ = v59
                default:
                    if false {
                        switch ("a" + "foo bar") {
                        case "a":
                            fmt.Println(v2, v16, v1)
                        case "\n\t":
                            if true {
                                var v60 int8 = 42
                                for fi61 := 0; fi61 < 20; fi61++ {
                                    v62 := GenS8{F0: 1, F1: 1}
                                    if false {
                                        if false {
                                            if (!true) {
                                                var v63 int64 = (^(v2))
                                                v3[v1] = 0
                                                var v64 int32 = 0
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        v16 += v16
                                    } else {
                                        var v65 uint32 = 291
                                        fmt.Println(v16)
                                        v66 := GenS9{F0: (v16 + 520), F1: 2}
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = fi61
                                    _ = v62
                                }
                                _ = v60
                            }
                        }
                        if (true && (2 >= 255)) {
                            v67 := GenS6{F0: 189, F1: -1}
                            var v68 uint64 = (v16 ^ v16)
                            fmt.Println(v16, v2, v68)
                            _ = v67
                            _ = v68
                        }
                        v16 = (v16 << ((18446744073709551615) & 63))
                    } else {
                        v1 -= v1
                        v1 = v1
                    }
                }
                if false {
                    v1 = 32768
                    var v69 int64 = (-(v2))
                    for fi70 := 0; fi70 < 2; fi70++ {
                        var v71 uint8 = 128
                        _ = fi70
                        _ = v71
                    }
                    _ = v69
                }
            }
            fmt.Println(v2, v1)
        }
        for fi72 := 0; fi72 < 12; fi72++ {
            v1 = fi72
            var v73 float32 = (float32(1.5) / float32(((3.14159 + 3.14159) + 0.1)))
            _ = fi72
            _ = v73
        }
    }
    fmt.Println(v1)
    for fi74 := 0; fi74 < 15; fi74++ {
        if ((v16 % ((v16) | 1)) == (v16 | (v16 + (v16 * v16)))) {
            delete(v3, (fi74 | v1))
            v75 := []int{(fi74 | v1)}
            if ((-1 <= 1) && true) {
                if len(v75) > 0 {
                    v75[0] = int(v16)
                }
            }
            _ = v75
        }
        v76 := GenS7{F0: -451, F1: genFn0(false, v2), F2: -1}
        _ = fi74
        _ = v76
    }
    var v77 uint32 = 546
    v78 := []int{v1, 0}
    v16 = (v16 + (v16 - v16))
    v2 = (v2 / ((v2) | 1))
    if true {
        v77 -= v77
        v79 := GenS10{F0: (!((!false) || (2 != 18446744073709551615))), F1: "\n\t", F2: 0}
        fmt.Println(v2, v2)
        _ = v79
    } else {
        var v80 float64 = 3.14159
        v2 -= (v2 << ((255) & 63))
        _ = v80
    }
    v2 += (v2 + 42)
    if (9223372036854775807 >= v2) {
        var v81 uint32 = v77
        v82 := GenS11{F0: 295}
        _ = v81
        _ = v82
    } else {
        var v83 int16 = 42
        switch "foo bar" {
        case "foo bar":
            var v84 uint32 = v77
            _ = v84
        }
        _ = len(v3)
        _ = v83
    }
    v85 := GenS8{F0: 1000, F1: -2147483648}
    if len(v78) >= 2 {
        v78 = v78[1:len(v78)]
    }
    v86 := GenS3{F0: 2, F1: 205}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v77:", v77)
    fmt.Println("slicelen v78:", len(v78), v78)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v77
    _ = v78
    _ = v85
    _ = v86
}
