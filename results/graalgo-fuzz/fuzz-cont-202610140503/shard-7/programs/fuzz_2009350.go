package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 string }
func genFn0(p0 uint8) uint64 {
    var v4 int8 = -128
    var v5 int32 = -2147483648
    v4 = v4
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int32) int32 {
    var v6 float64 = float64(((-2.25 - 1e10) * 3.14159))
    var v7 int8 = 1
    _ = p0
    _ = v6
    _ = v7
    return (p0 + p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 int16 = 583
    switch 865 {
    case 127:
        for fi8 := 0; fi8 < 11; fi8++ {
            switch ("" + "foo bar") {
            case "\n\t":
                var v9 uint = 443
                _ = v9
            case "foo bar":
                fmt.Println(fi8, fi8)
            case "":
                if true {
                    v3 += (v3 * v3)
                } else {
                    fmt.Println(v3, v3)
                }
            default:
                v10 := GenS1{F0: 750, F1: 0, F2: 128}
                _ = v10
            }
            _ = fi8
        }
    case -9223372036854775808:
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint8 = 255
            delete(v1, 1)
            _ = fi11
            _ = v12
        }
    default:
        var v13 bool = false
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS2{F0: (!(false && false)), F1: v3}
        switch -637 {
        case 4294967295:
            v16 := v1[708]
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    v1[541] = 65535
    if true {
        fmt.Println(v3)
    } else {
        var v17 uint64 = 18446744073709551615
        var v18 uint8 = 255
        _ = v17
        _ = v18
    }
    var v19 int64 = 42
    v3 = (^(v3))
    switch "a" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    default:
        var v20 uint64 = 4294967295
        _ = v20
    }
    var v21 string = "hello"
    if (!false) {
        fmt.Println(v21, v19)
        v22 := GenS3{F0: 2147483647, F1: -409}
        _ = v22
    }
    v23 := GenS1{F0: 4294967295, F1: 853, F2: 128}
    v24 := v1[65535]
    for fi25 := 0; fi25 < 8; fi25++ {
        v19 -= v19
        _ = fi25
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := GenS1{F0: v24, F1: 2147483647, F2: 127}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS4{F0: -77, F1: 2, F2: genFn0(1)}
        _ = fi28
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = v21
    }
    _ = len(v2)
    if true {
        switch (v21 + ((v21 + v21) + v21)) {
        case "\n\t":
            v24 -= v24
        }
    } else {
        var v30 int8 = 127
        if ((18446744073709551615 == uint64(v24)) && true) {
            var v31 float64 = (3.14159 * 0.1)
            v32 := GenS5{F0: "", F1: (v19 & int64(v3))}
            _ = v31
            _ = v32
        } else {
            if (!((855 >= 263) || true)) {
                var v33 uint64 = 18446744073709551615
                for fi34 := 0; fi34 < 19; fi34++ {
                    _ = len(v1)
                    v24 -= (v24 - v24)
                    v30 -= 100
                    _ = fi34
                }
                v35 := GenS2{F0: (0 == 4294967295), F1: (-(v3))}
                _ = v33
                _ = v35
            } else {
                for fi36 := 0; fi36 < 20; fi36++ {
                    if false {
                        switch (fi36 & fi36) {
                        case -576:
                            for fi37 := 0; fi37 < 2; fi37++ {
                                v24 -= 2147483647
                                var v38 string = v21
                                _ = fi37
                                _ = v38
                            }
                        case -843:
                            if false {
                                var v39 uint = 0
                                fmt.Println(v3, v30)
                                _ = v39
                            }
                        case 638:
                            v40 := GenS2{F0: false, F1: v3}
                            _ = v40
                        default:
                            v2 = append(v2, v21)
                        }
                    } else {
                        var v41 int32 = -1
                        _ = v41
                    }
                    _ = fi36
                }
                var v42 uint16 = 921
                for fi43 := 0; fi43 < 13; fi43++ {
                    fmt.Println(v30, v21, v42)
                    var v44 float32 = (((1e10 + 1e10) / 3.14159) / 3.14159)
                    v45 := GenS5{F0: "foo bar", F1: 0}
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = v42
            }
        }
        _ = v30
    }
    if false {
        for fi46 := 0; fi46 < 20; fi46++ {
            var v47 int32 = 42
            fmt.Println(v3)
            var v48 int16 = (v3 >> ((241) & 63))
            _ = fi46
            _ = v47
            _ = v48
        }
    } else {
        v49 := v1[(v24 >> ((18446744073709551615) & 63))]
        switch 2147483647 {
        case 396:
            v24 = v24
        case 127:
            switch ((v21 + v21) + (((("hello" + v21) + v21) + "a") + ((v21 + ((v21 + "foo bar") + "")) + "hello"))) {
            case "a":
                switch -9223372036854775808 {
                case 1:
                    v24 = (v24 ^ v24)
                case 9223372036854775807:
                    delete(v1, 300)
                case 4294967295:
                    for fi50 := 0; fi50 < 5; fi50++ {
                        var v51 int = fi50
                        if true {
                            v51 += (-(fi50))
                            switch 1 {
                            case 1:
                                v49 += v24
                            }
                            var v52 int = fi50
                            _ = v52
                        } else {
                            for fi53 := 0; fi53 < 8; fi53++ {
                                var v54 int8 = 42
                                switch 32767 {
                                case -1:
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                case 9223372036854775807:
                                    v55 := v1[4294967295]
                                    _ = v55
                                default:
                                    fmt.Println(v54, v21)
                                }
                                if true {
                                    switch (-(fi53)) {
                                    case 67:
                                        if (!false) {
                                            v19 -= (v19 + (v19 ^ (v19 >> ((548) & 63))))
                                        }
                                    case 574:
                                        if (!true) {
                                            v56 := GenS2{F0: (!true), F1: -999}
                                            _ = v56
                                        }
                                    case 42:
                                        if (v3 > (v3 >> ((4294967295) & 63))) {
                                            fmt.Println(v54, v49, fi50)
                                            switch fi50 {
                                            case -870:
                                                v57 := GenS4{F0: v54, F1: 18446744073709551615, F2: 4294967295}
                                                _ = v57
                                            case 42:
                                                v19 -= (v19 >> ((112) & 63))
                                            case 32767:
                                                switch "hello" {
                                                case "\n\t":
                                                    if (255 < 255) {
                                                        v24 -= (v24 << ((0) & 63))
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                case "foo bar":
                                                    v58 := v1[(v24 & 42)]
                                                    _ = v58
                                                default:
                                                    v54 = (-(v54))
                                                }
                                            }
                                        } else {
                                            delete(v1, (^(v24)))
                                            switch v21 {
                                            case "":
                                                v54 -= (^(v54))
                                            }
                                        }
                                    default:
                                        if (127 > 164) {
                                            v54 += 127
                                            fmt.Println(v54, v24)
                                            v19 = v19
                                        } else {
                                            v54 = 127
                                            if (!(((false && false) && (!(!(v21 <= (v21 + ((v21 + (((v21 + v21) + v21) + v21)) + v21)))))) && false)) {
                                                var v59 int64 = int64(v51)
                                                switch (v21 + v21) {
                                                case "foo bar":
                                                    v54 = v54
                                                case "":
                                                    v3 += v3
                                                case "a":
                                                    fmt.Println(v21, v3, v24)
                                                default:
                                                    v1[934] = (v24 + 1)
                                                }
                                                switch -482 {
                                                case -962:
                                                    fmt.Println(fi53)
                                                case -9223372036854775808:
                                                    var v60 uint16 = 42
                                                    _ = v60
                                                }
                                                _ = v59
                                            } else {
                                                if false {
                                                    v51 = fi50
                                                    if true {
                                                        v54 = (v54 * v54)
                                                    }
                                                }
                                                v61 := v1[v24]
                                                _ = v61
                                            }
                                            if ((v21 + (v21 + "")) >= "foo bar") {
                                                var v62 int32 = int32(v3)
                                                _ = v62
                                            }
                                            var v63 int64 = v19
                                            _ = v63
                                        }
                                    }
                                } else {
                                    v49 -= v24
                                    v54 += (v54 >> ((2) & 63))
                                    _ = len(v1)
                                    fmt.Println(v3, v54, fi53)
                                }
                                _ = fi53
                                _ = v54
                            }
                            v64 := GenS3{F0: 42, F1: v19}
                            _ = v64
                        }
                        _ = fi50
                        _ = v51
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case "\n\t":
                v21 = (((v21 + ("a" + v21)) + "a") + "")
            case "foo bar":
                fmt.Println(v21, v3, v49)
            default:
                v24 = v49
            }
        case 4294967295:
            if ((v3 ^ v3) <= (^(v3))) {
                for fi65 := 0; fi65 < 5; fi65++ {
                    for fi66 := 0; fi66 < 5; fi66++ {
                        fmt.Println(v19)
                        var v67 uint32 = 656
                        v3 = v3
                        switch fi66 {
                        case 32767:
                            if ((v19 / ((v19) | 1)) >= (v19 % ((v19) | 1))) {
                                switch "a" {
                                case "\n\t":
                                    v2 = append(v2, ("\n\t" + "hello"))
                                }
                            } else {
                                if ((-1 < 42) || false) {
                                    fmt.Println(v24)
                                    _ = len(v2)
                                }
                                if true {
                                    v2 = append(v2, v21)
                                    v24 -= 2147483648
                                } else {
                                    v68 := GenS6{F0: fi66}
                                    v3 += 42
                                    _ = v68
                                }
                                v49 = 4294967295
                                v69 := GenS6{F0: fi65}
                                if (42 != 42) {
                                    var v70 bool = false
                                    var v71 int64 = 0
                                    var v72 int = 295
                                    v73 := GenS5{F0: v21, F1: (v19 ^ (v71 | (v71 << ((255) & 63))))}
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                    _ = v73
                                }
                                _ = v69
                            }
                        case -1:
                            v3 = (v3 >> ((2) & 63))
                        case -960:
                            v21 = (v21 + "hello")
                        default:
                            v74 := GenS3{F0: 2147483647, F1: 352}
                            _ = v74
                        }
                        v75 := GenS7{F0: -1}
                        _ = fi66
                        _ = v67
                        _ = v75
                    }
                    _ = fi65
                }
            } else {
                v76 := GenS2{F0: true, F1: v3}
                if (!(false || ((1.5 + (0.1 * (1.5 / (-2.25 - (1.5 + (0.0 * (3.14159 - (3.14159 - ((0.0 - (1e10 / ((1.5 + 0.0) + (0.0 * (-2.25 / -2.25))))) - (0.1 * (0.1 - 3.14159))))))))))) == 1.5))) {
                    v19 -= v19
                    v77 := GenS2{F0: (1e10 != -2.25), F1: v3}
                    _ = v77
                }
                if true {
                    if (((2 != 727) && (1.5 == 3.14159)) && (1 < 0)) {
                        if false {
                            v49 = (v49 / ((2) | 1))
                            if true {
                                var v78 float32 = 0.0
                                if false {
                                    v49 += (v24 * (v24 % (((^(v24))) | 1)))
                                    v79 := GenS6{F0: 0}
                                    _ = v79
                                }
                                v49 -= (v49 * (^(v24)))
                                _ = v78
                            }
                            switch "" {
                            case "foo bar":
                                v21 = (v21 + ((v21 + (v21 + "\n\t")) + (("foo bar" + "foo bar") + v21)))
                            case "\n\t":
                                if len(v2) > 0 {
                                    v2[0] = v21
                                }
                            case "hello":
                                var v80 uint64 = 2
                                _ = v80
                            }
                        } else {
                            var v81 int8 = -1
                            fmt.Println(v81, v49)
                            v3 = (^(v3))
                            _ = v81
                        }
                        v82 := GenS4{F0: 27, F1: 898, F2: 1}
                        _ = v82
                    }
                    var v83 int64 = 9223372036854775807
                    _ = v83
                }
                _ = v76
            }
        default:
            for fi84 := 0; fi84 < 3; fi84++ {
                for fi85 := 0; fi85 < 15; fi85++ {
                    v86 := GenS8{F0: (^(v24)), F1: v21}
                    _ = fi85
                    _ = v86
                }
                v87 := GenS1{F0: v24, F1: -914, F2: 127}
                v21 = ((v21 + ("hello" + ((v21 + v21) + ""))) + (("a" + v21) + v21))
                _ = fi84
                _ = v87
            }
        }
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
