package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int8; F1 string; F2 int16 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint16; F1 int64; F2 int32 }
type GenS5 struct { F0 string }
func genFn0(p0 uint64, p1 int64) uint8 {
    v5 := []int{0, -267, -2147483648}
    v6 := GenS1{F0: -1, F1: 96}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 128
}
func genFn1() string {
    var v7 float64 = -2.25
    switch "a" {
    case "hello":
        if (2147483648 == 42) {
            fmt.Println(v7)
            fmt.Println(v7)
        }
    case "foo bar":
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 int32 = 2147483647
                if (true || (true && (1 != -238))) {
                    v11 := GenS1{F0: (v10 << ((65535) & 63)), F1: (fi9 & 0)}
                    _ = v11
                }
                v10 += v10
                _ = fi9
                _ = v10
            }
            var v12 uint16 = 180
            _ = fi8
            _ = v12
        }
    }
    _ = v7
    return "hello"
}

func main() {
    var v1 float32 = ((0.1 - 0.0) / ((1.5 + 0.1) / 1.0))
    var v2 int = 1
    v3 := make(map[int]uint32)
    var v4 uint64 = 4294967295
    switch "\n\t" {
    case "\n\t":
        var v13 float64 = (-2.25 * (((float64((v1 - (v1 - v1))) * 0.1) - (0.0 * 1.5)) + (((0.1 / (1.5 / 1.5)) - float64(v1)) * 1.5)))
        _ = v13
    case "":
        v1 += v1
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            if true {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS2{F0: -12, F1: "", F2: 42}
                    _ = len(v3)
                    v17 := []int{v2, (fi15 / (((fi14 | fi15)) | 1))}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch "a" {
                case "hello":
                    var v18 int8 = -121
                    _ = v18
                default:
                    if (v2 >= fi14) {
                        delete(v3, v2)
                        for fi19 := 0; fi19 < 18; fi19++ {
                            if (1 != genFn0(v4, 42)) {
                                switch ("hello" + "\n\t") {
                                case "a":
                                    v1 -= v1
                                case "\n\t":
                                    switch (fi19 % ((32767) | 1)) {
                                    case 32768:
                                        if true {
                                            fmt.Println(fi19)
                                        } else {
                                            v4 = v4
                                        }
                                    case -2147483648:
                                        var v20 uint64 = v4
                                        _ = v20
                                    case 966:
                                        v1 -= ((v1 + 0.1) * ((-2.25 - ((0.0 + v1) * float32(float64(v1)))) + v1))
                                    }
                                }
                                v4 = v4
                            }
                            _ = fi19
                        }
                        var v21 uint8 = 42
                        _ = v21
                    } else {
                        if true {
                            if true {
                                v3[(fi14 >> ((4294967295) & 63))] = 1
                                var v22 string = ("hello" + "a")
                                _ = v22
                            } else {
                                v2 = (v2 ^ v2)
                                fmt.Println(v2)
                                for fi23 := 0; fi23 < 13; fi23++ {
                                    var v24 uint = 42
                                    v25 := GenS1{F0: -2147483648, F1: (v2 + 4294967295)}
                                    v26 := v3[(v2 * (v2 % ((42) | 1)))]
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            }
                            v1 = v1
                            v4 += 1
                        } else {
                            v27 := GenS1{F0: 0, F1: v2}
                            _ = v27
                        }
                        delete(v3, v2)
                        for fi28 := 0; fi28 < 2; fi28++ {
                            v29 := v3[-2147483648]
                            var v30 float64 = float64(0.0)
                            v31 := GenS3{F0: 127, F1: true}
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    }
                }
            } else {
                var v32 uint = 4294967295
                _ = v32
            }
            var v33 bool = false
            _ = fi14
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 int = (fi34 ^ (fi35 * (fi34 << ((2) & 63))))
            v37 := GenS4{F0: 0, F1: -1, F2: -1}
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = fi34
    }
    v38 := GenS4{F0: 2, F1: -9223372036854775808, F2: 42}
    for fi39 := 0; fi39 < 19; fi39++ {
        switch v2 {
        case 1:
            switch "a" {
            case "\n\t":
                v40 := []int{}
                _ = v40
            default:
                var v41 int32 = -2147483648
                _ = v41
            }
        case -768:
            var v42 int8 = 42
            _ = v42
        }
        var v43 float64 = 0.0
        _ = fi39
        _ = v43
    }
    switch 42 {
    case 9223372036854775807:
        if (!((1 < 18446744073709551615) || ((v2 / (((^(v2))) | 1)) > v2))) {
            v2 -= -968
            v4 = (v4 / ((v4) | 1))
            if (!(976 < 18446744073709551615)) {
                var v44 float64 = 3.14159
                _ = v44
            }
        }
    case 147:
        if ("hello" < "foo bar") {
            if true {
                var v45 bool = (false || true)
                _ = v45
            } else {
                delete(v3, (v2 / (((v2 + 4294967295)) | 1)))
                var v46 uint8 = 42
                if false {
                    fmt.Println(v1)
                }
                _ = v46
            }
        } else {
            for fi47 := 0; fi47 < 11; fi47++ {
                for fi48 := 0; fi48 < 15; fi48++ {
                    fmt.Println(fi48, v1)
                    v1 -= ((float32((0.1 / 3.14159)) / (v1 * -2.25)) + float32(1e10))
                    fmt.Println(v4, v2, v1)
                    _ = fi48
                }
                var v49 float64 = float64(float32(1.5))
                _ = fi47
                _ = v49
            }
            var v50 int32 = 42
            _ = v50
        }
    default:
        var v51 uint = 492
        _ = v51
    }
    v2 -= (v2 | (v2 % ((v2) | 1)))
    v4 += v4
    v2 = (v2 | 1)
    var v52 float32 = v1
    for fi53 := 0; fi53 < 4; fi53++ {
        var v54 int32 = 0
        var v55 bool = (!(genFn0((^(v4)), 0) != 128))
        switch ("\n\t" + "") {
        case "":
            for fi56 := 0; fi56 < 7; fi56++ {
                v4 = v4
                _ = fi56
            }
        case "hello":
            var v57 uint64 = 4294967295
            _ = v57
        }
        _ = fi53
        _ = v54
        _ = v55
    }
    fmt.Println(v52)
    v58 := []int{v2, (v2 ^ (v2 - v2)), -267, (v2 + (-(v2)))}
    v59 := GenS3{F0: 61, F1: (42 > -2147483648)}
    fmt.Println(v2, v52, v52)
    v2 = (v2 + (v2 << ((692) & 63)))
    switch (-(v2)) {
    case 0:
        v52 = ((((v1 / float32((1e10 * 1.5))) * (v52 + v1)) + v1) - v52)
    case 821:
        v2 = -9223372036854775808
    case 42:
        v60 := GenS1{F0: 197, F1: 32767}
        _ = v60
    default:
        if true {
            var v61 float32 = 1.5
            _ = v61
        }
    }
    v4 = 2
    for fi62 := 0; fi62 < 17; fi62++ {
        for fi63 := 0; fi63 < 9; fi63++ {
            v2 -= (^(fi63))
            if (!true) {
                v52 -= 0.0
                if true {
                    var v64 int8 = 0
                    var v65 bool = (v52 >= 0.0)
                    var v66 int64 = 0
                    _ = len(v3)
                    _ = v64
                    _ = v65
                    _ = v66
                }
            } else {
                var v67 uint8 = 128
                _ = v67
            }
            _ = fi63
        }
        var v68 uint16 = 304
        if false {
            var v69 uint = 65535
            var v70 float64 = 3.14159
            _ = v69
            _ = v70
        } else {
            if true {
                switch (fi62 + 873) {
                case 227:
                    v71 := v3[v2]
                    _ = v71
                case 9223372036854775807:
                    if false {
                        var v72 uint16 = (v68 + (v68 / ((v68) | 1)))
                        _ = v72
                    } else {
                        v58 = append(v58, 128)
                        if (-1 > 70) {
                            switch "a" {
                            case "\n\t":
                                for fi73 := 0; fi73 < 20; fi73++ {
                                    if (true || (693 <= 42)) {
                                        v1 = v1
                                        v74 := GenS4{F0: (^(v68)), F1: -333, F2: 30}
                                        _ = v74
                                    } else {
                                        var v75 uint64 = (v4 % (((v4 ^ v4)) | 1))
                                        switch v2 {
                                        case 9223372036854775807:
                                            if len(v58) >= 2 {
                                                v58 = v58[1:len(v58)]
                                            }
                                        default:
                                            v58 = append(v58, (-(v2)))
                                        }
                                        var v76 uint32 = 4294967295
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = fi73
                                }
                            case "foo bar":
                                var v77 int32 = 1
                                _ = v77
                            default:
                                v52 = 0.0
                            }
                            var v78 int16 = 1000
                            _ = v78
                        } else {
                            for fi79 := 0; fi79 < 14; fi79++ {
                                if len(v58) > 0 {
                                    v58[0] = (^(v2))
                                }
                                var v80 bool = true
                                _ = fi79
                                _ = v80
                            }
                            v81 := GenS5{F0: "foo bar"}
                            var v82 uint = 225
                            _ = v81
                            _ = v82
                        }
                    }
                default:
                    for fi83 := 0; fi83 < 17; fi83++ {
                        var v84 int = v2
                        var v85 float32 = v1
                        var v86 bool = (65535 > 0)
                        if true {
                            v87 := GenS2{F0: int8(v68), F1: (("" + "\n\t") + "\n\t"), F2: 44}
                            var v88 string = ("" + "hello")
                            _ = v87
                            _ = v88
                        } else {
                            v52 = 1e10
                            if true {
                                if len(v58) >= 2 {
                                    v58 = v58[1:len(v58)]
                                }
                            }
                            v2 = 32767
                            v52 = ((((float32(float64((1e10 / v1))) + float32((-2.25 - -2.25))) - v1) / float32((-2.25 / (-2.25 / 1e10)))) + ((float32(-2.25) * 0.1) / float32((1e10 / (1e10 / -2.25)))))
                            var v89 uint8 = 117
                            _ = v89
                        }
                        _ = fi83
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                }
            }
        }
        _ = fi62
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v52
    _ = v58
    _ = v59
}
