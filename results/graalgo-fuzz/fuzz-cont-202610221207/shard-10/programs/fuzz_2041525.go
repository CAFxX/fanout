package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int16; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint8; F2 int }
type GenS7 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS8 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() string {
    v4 := []int{128, 9223372036854775807, 836, 1}
    var v5 uint8 = 1
    var v6 bool = ((498 == 4294967295) || ((!true) || true))
    v7 := GenS1{F0: 2}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return ("" + "foo bar")
}
func genFn1(p0 bool, p1 int32, p2 int64) int32 {
    fmt.Println(p2, p0, p2)
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            if p0 {
                switch fi9 {
                case -927:
                    var v10 int32 = p1
                    _ = v10
                case 32767:
                    p0 = (p0 || (true && (true && (-2.25 <= 0.0))))
                default:
                    var v11 int16 = 24
                    _ = v11
                }
                var v12 uint16 = 255
                if ((p1 / (((p1 | (p1 ^ (p1 >> ((308) & 63))))) | 1)) <= (-(p1))) {
                    if ((0.0 > (-2.25 + 0.0)) && p0) {
                        if ((((((p1 & 0) > (p1 ^ 0)) || (p2 < -420)) && (false || (!(42 < 239)))) || true) && true) {
                            fmt.Println(v12)
                            var v13 int16 = -477
                            v14 := GenS1{F0: 4294967295}
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(p1, v12, p2)
                        }
                        v15 := GenS1{F0: 4294967295}
                        var v16 float32 = -2.25
                        var v17 float64 = float64(v16)
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v12 -= 831
                } else {
                    switch ("a" + (("a" + (((("\n\t" + "") + "") + "foo bar") + "hello")) + "foo bar")) {
                    case "":
                        v18 := []int{(fi8 - (fi8 % ((-9223372036854775808) | 1)))}
                        _ = v18
                    case "\n\t":
                        if p0 {
                            if (4294967295 >= 0) {
                                v19 := []int{}
                                _ = v19
                            }
                            p0 = (100 > 42)
                            v20 := []int{fi8, 34, (fi9 + 32767), fi9}
                            _ = v20
                        } else {
                            var v21 bool = ("hello" >= "hello")
                            _ = v21
                        }
                    case "foo bar":
                        fmt.Println(v12, fi8, p1)
                    }
                    v22 := []int{-1, fi9, (fi8 ^ (fi9 * fi8)), 32768}
                    _ = v22
                }
                _ = v12
            } else {
                switch (fi9 * 1) {
                case 127:
                    p1 = 0
                case -217:
                    p1 = p1
                default:
                    if p0 {
                        switch 32767 {
                        case -2147483648:
                            var v23 uint8 = 0
                            _ = v23
                        case 32767:
                            if true {
                                p1 += -950
                            }
                        default:
                            p0 = (1 >= 32767)
                        }
                        fmt.Println(fi9)
                        var v24 uint16 = 810
                        if p0 {
                            var v25 int16 = 406
                            if (p0 && p0) {
                                fmt.Println(v24, p0)
                                v26 := GenS2{F0: p0}
                                _ = v26
                            } else {
                                if p0 {
                                    v27 := GenS1{F0: 42}
                                    _ = v27
                                }
                            }
                            _ = v25
                        }
                        var v28 bool = false
                        _ = v24
                        _ = v28
                    }
                }
                fmt.Println(fi8)
            }
            _ = fi9
        }
        _ = fi8
    }
    if (42 < 18446744073709551615) {
        p0 = (1.5 >= (1.5 + 3.14159))
    } else {
        var v29 uint8 = 188
        _ = v29
    }
    switch -939 {
    case 1:
        var v30 uint8 = 42
        _ = v30
    case 605:
        var v31 int8 = 22
        _ = v31
    default:
        v32 := GenS2{F0: p0}
        _ = v32
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 bool = (1 <= 0)
    var v2 int64 = -99
    v3 := []string{"foo bar", ("a" + "foo bar"), ("" + "a"), "a"}
    var v33 bool = v1
    var v34 uint16 = 0
    var v35 int = 134
    v34 -= 65535
    var v36 uint8 = 1
    var v37 uint = 1
    v34 = 0
    var v38 uint = v37
    _ = len(v3)
    var v39 uint64 = 4294967295
    v3 = append(v3, "a")
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v40 uint32 = 42
    if (!v1) {
        if (1000 != 42) {
            for fi41 := 0; fi41 < 11; fi41++ {
                for fi42 := 0; fi42 < 9; fi42++ {
                    v43 := GenS1{F0: 568}
                    v44 := GenS1{F0: uint(v39)}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                _ = fi41
            }
            fmt.Println(v37, v39, v35)
        } else {
            switch (v35 << ((42) & 63)) {
            case 908:
                var v45 int64 = v2
                _ = v45
            case -113:
                for fi46 := 0; fi46 < 5; fi46++ {
                    v3 = append(v3, (("a" + "a") + "\n\t"))
                    switch "a" {
                    case "a":
                        var v47 uint16 = 870
                        _ = v47
                    case "foo bar":
                        if (!((v40 >> ((255) & 63)) != v40)) {
                            var v48 uint32 = v40
                            _ = len(v3)
                            _ = v48
                        }
                    }
                    if v1 {
                        if v33 {
                            if ((float32((0.0 + float64((0.1 + (((1e10 + float32(0.0)) + float32(0.0)) - (1e10 + 0.1)))))) * (-2.25 / 1.5)) >= float32(0.0)) {
                                v49 := GenS1{F0: (v38 % ((v37) | 1))}
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    if v1 {
                                        var v51 uint32 = v40
                                        var v52 int16 = 0
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        v53 := GenS3{F0: v38, F1: 928}
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    _ = fi50
                                }
                                _ = v49
                            } else {
                                var v54 uint64 = uint64(v36)
                                if (!((float32(-2.25) - ((3.14159 / ((0.1 / (-2.25 / 0.1)) * 3.14159)) + 0.0)) <= 1e10)) {
                                    if (!((0.1 > 1.5) || true)) {
                                        v37 += (v38 * (v37 % ((uint(v39)) | 1)))
                                        var v55 int = -20
                                        v56 := GenS1{F0: (^(v38))}
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 uint64 = uint64(v34)
                                    _ = v57
                                }
                                var v58 bool = v1
                                _ = v54
                                _ = v58
                            }
                        }
                        if v33 {
                            switch "a" {
                            case "hello":
                                if (!(!((3.14159 == 0.0) && true))) {
                                    fmt.Println(v2, v40, v36)
                                    var v59 int8 = 42
                                    if (uint32(v36) != uint32(v39)) {
                                        var v60 uint64 = v39
                                        _ = v60
                                    } else {
                                        if ((32768 != 2147483647) && true) {
                                            v36 = 200
                                        }
                                        var v61 string = ("" + ("" + "foo bar"))
                                        v1 = v33
                                        _ = v61
                                    }
                                    _ = v59
                                } else {
                                    v3 = append(v3, "\n\t")
                                    var v62 int = (-(v35))
                                    _ = v62
                                }
                            case "\n\t":
                                _ = len(v3)
                            case "a":
                                v36 -= v36
                            }
                            var v63 float32 = 0.0
                            var v64 int64 = v2
                            _ = v63
                            _ = v64
                        } else {
                            switch int(v38) {
                            case -347:
                                v3 = append(v3, "a")
                            case 32768:
                                var v65 int32 = genFn1(false, genFn1((true && v1), 727, -476), v2)
                                _ = v65
                            case 32767:
                                v3 = append(v3, "\n\t")
                            default:
                                if len(v3) > 0 {
                                    v3[0] = ("foo bar" + ((("\n\t" + "a") + ("a" + ("hello" + "a"))) + "hello"))
                                }
                            }
                            v66 := GenS2{F0: (true && v1)}
                            v39 -= v39
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 9; fi67++ {
                            _ = len(v3)
                            if (int8(v35) <= -1) {
                                fmt.Println(v33, v2, v40)
                                switch v35 {
                                case 1:
                                    v38 = v38
                                case -9223372036854775808:
                                    v34 = v34
                                case -895:
                                    if true {
                                        fmt.Println(fi67)
                                        v39 = (^(v39))
                                        if (false || false) {
                                            _ = len(v3)
                                            v2 = -1
                                        } else {
                                            var v68 uint32 = (v40 >> ((v37) & 63))
                                            var v69 float64 = 1e10
                                            v1 = (42 > (-(v2)))
                                            _ = v68
                                            _ = v69
                                        }
                                        v2 = v2
                                        v70 := GenS3{F0: (^(v37)), F1: (v40 ^ (v40 ^ (v40 ^ v40)))}
                                        var v71 uint16 = v34
                                        _ = v70
                                        _ = v71
                                    } else {
                                        v72 := GenS4{F0: v34, F1: 1000, F2: 1}
                                        var v73 bool = (int32(v37) >= 834)
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                var v74 uint = 2
                                _ = v74
                            } else {
                                v1 = (!true)
                                switch fi67 {
                                case -1:
                                    var v75 int64 = v2
                                    _ = v75
                                case -2147483648:
                                    v35 -= 0
                                default:
                                    v33 = ((v39 & (^(v39))) < 4294967295)
                                }
                            }
                            if (v37 > (v38 | 4294967295)) {
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                var v76 string = "\n\t"
                                if len(v3) > 0 {
                                    v3[0] = v76
                                }
                                _ = v76
                            }
                            _ = fi67
                        }
                    } else {
                        switch (v35 % ((32768) | 1)) {
                        case 9223372036854775807:
                            fmt.Println(v39, v38)
                        case -9223372036854775808:
                            for fi77 := 0; fi77 < 19; fi77++ {
                                if (true || ((^(v36)) <= (v36 ^ (v36 | 119)))) {
                                    switch v35 {
                                    case -2147483648:
                                        v34 = uint16(v37)
                                    case -9223372036854775808:
                                        fmt.Println(v2, v40)
                                        v36 = (v36 >> ((0) & 63))
                                        switch "" {
                                        case "foo bar":
                                            var v78 int8 = -76
                                            _ = v78
                                        default:
                                            v35 -= (fi46 * 857)
                                        }
                                    }
                                    var v79 int16 = 1000
                                    var v80 uint64 = (v39 >> ((902) & 63))
                                    var v81 uint64 = v80
                                    if ((v1 && v33) || (true && true)) {
                                        var v82 int32 = -1
                                        var v83 uint32 = (v40 << ((v37) & 63))
                                        v33 = (int8(fi77) >= -128)
                                        _ = v82
                                        _ = v83
                                    } else {
                                        var v84 int64 = (-(v2))
                                        _ = v84
                                    }
                                    v85 := GenS5{F0: false}
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v85
                                }
                                v86 := GenS6{F0: 9223372036854775807, F1: (^(v36)), F2: 1}
                                _ = fi77
                                _ = v86
                            }
                        }
                    }
                    _ = fi46
                }
            case -2147483648:
                switch "foo bar" {
                case "\n\t":
                    switch (v35 - (v35 ^ 32767)) {
                    case 1:
                        fmt.Println(v2)
                    case -355:
                        v87 := GenS7{F0: (v34 | uint16(v39)), F1: v2, F2: v36}
                        _ = v87
                    default:
                        for fi88 := 0; fi88 < 19; fi88++ {
                            v89 := GenS8{F0: "foo bar", F1: (v2 << ((v37) & 63)), F2: 152}
                            for fi90 := 0; fi90 < 6; fi90++ {
                                if (!(!true)) {
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                    var v91 int8 = 8
                                    _ = v91
                                } else {
                                    var v92 uint64 = 1
                                    v34 -= v34
                                    var v93 uint8 = v36
                                    _ = v92
                                    _ = v93
                                }
                                var v94 int8 = 1
                                _ = fi90
                                _ = v94
                            }
                            _ = fi88
                            _ = v89
                        }
                    }
                default:
                    for fi95 := 0; fi95 < 17; fi95++ {
                        var v96 float32 = 3.14159
                        v38 = (v38 ^ v37)
                        _ = fi95
                        _ = v96
                    }
                }
            default:
                for fi97 := 0; fi97 < 5; fi97++ {
                    v98 := GenS5{F0: (("hello" != ("hello" + "hello")) && ((v38 != v38) && v33))}
                    var v99 uint32 = uint32(v36)
                    _ = fi97
                    _ = v98
                    _ = v99
                }
            }
        }
        if (v40 > v40) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v100 := GenS7{F0: v34, F1: 840, F2: (^(v36))}
            _ = v100
        }
    } else {
        var v101 int8 = 100
        _ = v101
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    for fi102 := 0; fi102 < 1; fi102++ {
        var v103 uint8 = uint8(v40)
        v103 = v36
        _ = fi102
        _ = v103
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
