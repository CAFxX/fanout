package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 bool; F1 int16; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int16; F2 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int; F1 string; F2 int16 }
func genFn0() uint {
    var v5 string = "a"
    v5 = (v5 + "")
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint16 = 200
        _ = fi6
        _ = v7
    }
    _ = v5
    return 1000
}
func genFn1(p0 bool) uint {
    v8 := GenS1{F0: -710, F1: 0, F2: 477}
    switch "\n\t" {
    case "\n\t":
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := []int{fi9, 4294967295, (fi10 - -9223372036854775808)}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                p0 = p0
                v12 := GenS2{F0: 255}
                if false {
                    v11 = append(v11, 32767)
                    v11 = append(v11, fi10)
                    v13 := GenS2{F0: 629}
                    if (2 >= 127) {
                        v14 := GenS3{F0: 255, F1: true}
                        _ = v14
                    }
                    _ = v13
                }
                fmt.Println(fi9, fi9)
                _ = fi10
                _ = v11
                _ = v12
            }
            if p0 {
                v15 := GenS4{F0: ((1e10 + 0.0) <= 1.5), F1: 501, F2: p0}
                fmt.Println(fi9)
                p0 = true
                _ = v15
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    p0 = (98 < -1)
                    _ = fi16
                }
                p0 = false
                v17 := GenS4{F0: p0, F1: 42, F2: (((9223372036854775807 == 9223372036854775807) || p0) || p0)}
                _ = v17
            }
            _ = fi9
        }
    case "foo bar":
        p0 = p0
    default:
        var v18 uint = 1
        _ = v18
    }
    p0 = p0
    _ = p0
    _ = v8
    return 65535
}

func main() {
    var v1 uint16 = 255
    v2 := []string{}
    var v3 int16 = 32767
    v4 := []int{-2147483648, 32768}
    v3 -= (v3 % ((-1) | 1))
    fmt.Println(v3, v1)
    fmt.Println(v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "hello" {
    case "a":
        var v19 int8 = -81
        _ = v19
    case "\n\t":
        v1 += (v1 << ((18446744073709551615) & 63))
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            var v21 int = 32768
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        if true {
            v3 = (v3 % ((v3) | 1))
        }
        if (3.14159 <= 3.14159) {
            switch (-(fi22)) {
            case 128:
                switch (fi22 ^ (^(fi22))) {
                case -84:
                    v23 := GenS5{F0: 0}
                    _ = v23
                }
            case -1:
                v24 := GenS5{F0: 18446744073709551615}
                _ = v24
            default:
                if ((true || false) || true) {
                    if (!((!(108 > 746)) || true)) {
                        fmt.Println(v1, v1)
                        v1 = v1
                        for fi25 := 0; fi25 < 4; fi25++ {
                            v3 -= (-(v3))
                            if true {
                                v4 = append(v4, (fi25 ^ fi25))
                                v4 = append(v4, fi22)
                            } else {
                                fmt.Println(v1, v3)
                            }
                            v26 := GenS3{F0: (v1 | (v1 - v1)), F1: true}
                            v27 := GenS6{F0: 1, F1: (v3 * 1000), F2: -9223372036854775808}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    } else {
                        v3 -= (v3 >> ((4294967295) & 63))
                        _ = len(v2)
                    }
                    v1 = (v1 & 0)
                    for fi28 := 0; fi28 < 8; fi28++ {
                        var v29 uint8 = 221
                        if false {
                            var v30 int32 = 2147483647
                            _ = v30
                        } else {
                            var v31 uint64 = 60
                            v32 := GenS4{F0: ((!(((v1 >= (v1 % ((v1) | 1))) || (2 < 2)) && true)) || (!(!(-550 == (v3 << ((uint(v1)) & 63)))))), F1: v3, F2: false}
                            if (v3 != -804) {
                                if (128 >= (v29 << ((18446744073709551615) & 63))) {
                                    fmt.Println(v31, v1)
                                    switch (("\n\t" + "") + "") {
                                    case "foo bar":
                                        var v33 uint = 65535
                                        _ = v33
                                    default:
                                        if false {
                                            v31 -= (v31 - 906)
                                            if false {
                                                if len(v4) > 0 {
                                                    v4[0] = (fi22 | (fi28 << ((1) & 63)))
                                                }
                                                v1 += v1
                                                v34 := GenS3{F0: (v1 + 65535), F1: true}
                                                _ = v34
                                            } else {
                                                v1 = (v1 / ((886) | 1))
                                            }
                                            switch ("a" + "hello") {
                                            case "hello":
                                                fmt.Println(fi28, v1, v1)
                                                fmt.Println(fi28, v1)
                                                var v35 int = (^(fi28))
                                                _ = v35
                                            }
                                        } else {
                                            v29 = (^(v29))
                                        }
                                    }
                                    v36 := GenS7{F0: (fi28 / ((fi22) | 1))}
                                    if (v31 > v31) {
                                        v37 := GenS1{F0: -110, F1: (v29 | 52), F2: v31}
                                        var v38 float32 = 1e10
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = (fi28 | (fi22 * fi28))
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = fi28
                                        }
                                    }
                                    switch 127 {
                                    case 997:
                                        if true {
                                            if ("\n\t" != ("hello" + "\n\t")) {
                                                switch ("hello" + "hello") {
                                                case "a":
                                                    switch "a" {
                                                    case "\n\t":
                                                        v29 = v29
                                                    default:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                case "foo bar":
                                                    v39 := GenS2{F0: (v1 << ((258) & 63))}
                                                    _ = v39
                                                case "hello":
                                                    v40 := GenS1{F0: 1, F1: (v29 + v29), F2: (v31 & v31)}
                                                    _ = v40
                                                default:
                                                    v31 += (v31 - v31)
                                                }
                                            } else {
                                                v1 -= v1
                                                fmt.Println(v3, fi28, v29)
                                            }
                                            fmt.Println(fi28, v1)
                                            var v41 bool = (!(576 == 0))
                                            _ = v41
                                        } else {
                                            v3 += 32767
                                            var v42 float32 = 0.0
                                            v43 := GenS3{F0: (^(v1)), F1: (v42 == 1.5)}
                                            _ = v42
                                            _ = v43
                                        }
                                    default:
                                        v1 = (v1 * (^(v1)))
                                    }
                                    _ = v36
                                }
                            }
                            _ = v31
                            _ = v32
                        }
                        var v44 int16 = (^(v3))
                        v1 += (v1 % (((v1 & v1)) | 1))
                        var v45 int = fi22
                        if true {
                            switch "a" {
                            case "hello":
                                v1 += 2
                            case "":
                                if (v1 != v1) {
                                    v46 := GenS7{F0: v45}
                                    if (42 <= 496) {
                                        var v47 bool = false
                                        fmt.Println(v45, fi28)
                                        _ = v47
                                    } else {
                                        v29 += v29
                                        switch "a" {
                                        case "foo bar":
                                            v3 = 1000
                                        case "\n\t":
                                            var v48 uint8 = v29
                                            _ = v48
                                        default:
                                            var v49 int32 = 42
                                            _ = v49
                                        }
                                        v44 -= v3
                                    }
                                    fmt.Println(v1, fi22, v3)
                                    _ = v46
                                }
                            case "\n\t":
                                v44 += v44
                            default:
                                v50 := GenS8{F0: fi22, F1: ("foo bar" + ""), F2: v3}
                                _ = v50
                            }
                            v3 = -679
                        } else {
                            v51 := GenS2{F0: (v1 + 42)}
                            _ = v51
                        }
                        v4 = append(v4, -371)
                        var v52 bool = true
                        if v52 {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v29 -= (v29 | v29)
                            v45 += fi22
                        }
                        _ = fi28
                        _ = v29
                        _ = v44
                        _ = v45
                        _ = v52
                    }
                } else {
                    if (4294967295 >= 2147483648) {
                        var v53 int32 = 917
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                        _ = v53
                    }
                }
            }
            v54 := GenS3{F0: v1, F1: true}
            _ = v54
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi22
    }
    v2 = append(v2, "a")
    v3 += (^(v3))
    var v55 uint64 = 2
    var v56 int64 = -9223372036854775808
    fmt.Println(v56, v1)
    if len(v2) > 0 {
        v2[0] = ""
    }
    v56 = v56
    if len(v4) > 0 {
        v4[0] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
}
