package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 int; F2 int }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32; F1 bool; F2 int16 }
func genFn0(p0 bool) int64 {
    var v5 uint64 = 42
    p0 = (p0 && true)
    _ = p0
    _ = v5
    return 1
}
func genFn1() int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        if (1 <= 18446744073709551615) {
            fmt.Println(fi6)
        } else {
            for fi7 := 0; fi7 < 14; fi7++ {
                v8 := []int{}
                var v9 float32 = (3.14159 - 1e10)
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 string = (("a" + "foo bar") + ("\n\t" + ("\n\t" + ("a" + "foo bar"))))
            _ = v10
        }
        var v11 float64 = (0.0 - ((((3.14159 * float64(1.5)) / ((1e10 / float64((float32(float64(3.14159)) + (1e10 / -2.25)))) + 1.5)) - 0.0) / (-2.25 * ((1e10 / 3.14159) - (1e10 * (0.0 - ((float64(-2.25) + ((((0.0 + 3.14159) * (0.0 - -2.25)) / 1.0) + 3.14159)) / (1e10 + ((((3.14159 + float64(3.14159)) / (float64(1.5) + -2.25)) * 0.1) * 1e10)))))))))
        var v12 float32 = 0.0
        _ = fi6
        _ = v11
        _ = v12
    }
    var v13 int16 = -32768
    if true {
        v14 := []int{32767, -2147483648, 608}
        var v15 bool = true
        v13 -= -10
        _ = v14
        _ = v15
    } else {
        if (9223372036854775807 > 42) {
            v13 -= v13
            if (v13 != 1000) {
                switch "a" {
                case "hello":
                    var v16 int16 = (v13 * (v13 + (v13 + v13)))
                    _ = v16
                case "":
                    fmt.Println(v13)
                case "\n\t":
                    var v17 uint16 = 255
                    _ = v17
                }
                var v18 int32 = 42
                _ = v18
            } else {
                var v19 int32 = 807
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := GenS1{F0: -197}
                    var v22 int32 = (v19 | int32(v13))
                    fmt.Println(v19)
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v19
            }
        }
        switch "\n\t" {
        case "a":
            v13 = (v13 * v13)
        case "\n\t":
            fmt.Println(v13)
        case "foo bar":
            v13 = v13
        }
        if (-9223372036854775808 != 42) {
            var v23 uint = 255
            if false {
                var v24 uint64 = 490
                for fi25 := 0; fi25 < 18; fi25++ {
                    fmt.Println(v23, v23, v23)
                    _ = fi25
                }
                _ = v24
            } else {
                for fi26 := 0; fi26 < 16; fi26++ {
                    v23 -= 379
                    if true {
                        v27 := []int{}
                        fmt.Println(v23)
                        fmt.Println(fi26, v13)
                        _ = v27
                    }
                    var v28 uint16 = 42
                    _ = fi26
                    _ = v28
                }
                if (!(976 >= 42)) {
                    v13 = (v13 - 1)
                } else {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v23 = 42
                            v23 = v23
                            fmt.Println(fi30)
                            _ = fi30
                        }
                        v31 := []int{42}
                        _ = fi29
                        _ = v31
                    }
                    v23 = (v23 + v23)
                    switch 303 {
                    case 127:
                        var v32 float32 = 1e10
                        _ = v32
                    case -495:
                        fmt.Println(v13)
                    }
                }
            }
            v23 -= (v23 + (v23 * v23))
            _ = v23
        }
    }
    v33 := []int{-1, -1}
    _ = v13
    _ = v33
    return genFn0(true)
}

func main() {
    var v1 bool = true
    var v2 int32 = -1
    var v3 uint16 = 32767
    var v4 int8 = -1
    if false {
        if (474 != -465) {
            v4 = (-(v4))
        } else {
            var v34 uint32 = 0
            switch ("a" + "hello") {
            case "hello":
                v35 := GenS1{F0: -9223372036854775808}
                _ = v35
            case "":
                if v1 {
                    var v36 string = (("" + "a") + "\n\t")
                    var v37 float64 = (-2.25 - 0.0)
                    switch (("hello" + v36) + "") {
                    case "a":
                        fmt.Println(v37, v37, v4)
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            v38 := []int{9223372036854775807, -2147483648, 32767}
                            _ = v38
                        case "a":
                            switch 2147483647 {
                            case -113:
                                v39 := GenS2{F0: 42}
                                _ = v39
                            case -959:
                                v3 = uint16(v34)
                            case 32767:
                                switch "\n\t" {
                                case "\n\t":
                                    v40 := GenS1{F0: 42}
                                    _ = v40
                                }
                            default:
                                var v41 int64 = -1
                                _ = v41
                            }
                        case "foo bar":
                            for fi42 := 0; fi42 < 9; fi42++ {
                                v43 := GenS2{F0: 4294967295}
                                v34 = (v34 * 960)
                                fmt.Println(v34, v1, v36)
                                _ = fi42
                                _ = v43
                            }
                        default:
                            v44 := []int{}
                            _ = v44
                        }
                    }
                    _ = v36
                    _ = v37
                } else {
                    v4 = v4
                    var v45 int64 = -1
                    _ = v45
                }
            }
            _ = v34
        }
        for fi46 := 0; fi46 < 7; fi46++ {
            var v47 float64 = (0.0 * (((0.1 - 3.14159) * ((0.0 * float64((0.1 / 1.5))) - 1.5)) - -2.25))
            v48 := GenS3{F0: fi46, F1: v1, F2: 234}
            _ = fi46
            _ = v47
            _ = v48
        }
        switch 4294967295 {
        case 9223372036854775807:
            switch "foo bar" {
            case "foo bar":
                v49 := GenS1{F0: -385}
                _ = v49
            case "":
                v50 := []int{32768, -2147483648}
                _ = v50
            default:
                if false {
                    v1 = (!v1)
                    switch "a" {
                    case "":
                        for fi51 := 0; fi51 < 20; fi51++ {
                            var v52 float32 = 0.1
                            var v53 uint8 = 127
                            if v1 {
                                var v54 uint32 = 1
                                _ = v54
                            } else {
                                v55 := GenS4{F0: 2}
                                _ = v55
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    case "hello":
                        switch 0 {
                        case 9223372036854775807:
                            fmt.Println(v3)
                        default:
                            var v56 int8 = int8(v2)
                            _ = v56
                        }
                    case "foo bar":
                        v3 = v3
                    }
                    v57 := GenS4{F0: 1}
                    _ = v57
                }
            }
        case -2147483648:
            v58 := []int{-9223372036854775808, 128, -662, 9223372036854775807}
            _ = v58
        case 982:
            var v59 uint32 = 42
            _ = v59
        }
    } else {
        if false {
            v3 -= 65
            v2 = v2
            var v60 bool = ((v3 > (v3 & (v3 + 255))) || v1)
            _ = v60
        } else {
            if ((v1 && v1) || v1) {
                var v61 uint16 = 724
                _ = v61
            } else {
                v62 := []int{-665}
                _ = v62
            }
            for fi63 := 0; fi63 < 16; fi63++ {
                switch fi63 {
                case -304:
                    var v64 uint8 = 1
                    _ = v64
                case -9223372036854775808:
                    if (-409 <= (^(fi63))) {
                        var v65 int32 = v2
                        for fi66 := 0; fi66 < 5; fi66++ {
                            v65 = 1
                            v67 := GenS4{F0: 18446744073709551615}
                            switch -1 {
                            case 2147483647:
                                v68 := []int{}
                                _ = v68
                            case 1:
                                var v69 int = 17
                                switch v69 {
                                case 0:
                                    var v70 float64 = 0.1
                                    _ = v70
                                default:
                                    v4 -= (v4 + v4)
                                }
                                _ = v69
                            case 128:
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v1)
                                }
                            }
                            v1 = (v65 < 42)
                            if true {
                                v1 = (fi63 < -2147483648)
                                if (4294967295 == 920) {
                                    v3 -= v3
                                    if v1 {
                                        v2 = v2
                                        var v71 string = ""
                                        _ = v71
                                    }
                                    v1 = (0.0 < 3.14159)
                                    v3 = (v3 ^ 32767)
                                }
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v72 float32 = 3.14159
                                    _ = v72
                                case "a":
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v73 int8 = v4
                                            _ = v73
                                        case "a":
                                            if true {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v2, v3)
                                                }
                                                v74 := []int{(fi63 | 227), 127, (-(fi63))}
                                                v75 := GenS2{F0: 42}
                                                _ = v74
                                                _ = v75
                                            }
                                        default:
                                            v2 = (v65 - -2147483648)
                                        }
                                    }
                                }
                                v3 = v3
                                var v76 uint64 = 42
                                _ = v76
                            }
                            v65 = (v65 * 966)
                            v77 := []int{}
                            _ = fi66
                            _ = v67
                            _ = v77
                        }
                        var v78 float64 = -2.25
                        _ = v65
                        _ = v78
                    }
                case 32767:
                    for fi79 := 0; fi79 < 14; fi79++ {
                        v80 := GenS1{F0: -642}
                        v81 := GenS1{F0: 1}
                        _ = fi79
                        _ = v80
                        _ = v81
                    }
                }
                _ = fi63
            }
            v82 := []int{}
            _ = v82
        }
        v4 = 6
    }
    var v83 bool = false
    fmt.Println(v3, v3, v4)
    v84 := GenS5{F0: 1000, F1: 4294967295, F2: 32768}
    for fi85 := 0; fi85 < 11; fi85++ {
        var v86 int32 = (v2 >> ((99) & 63))
        switch (^(fi85)) {
        case 42:
            switch "hello" {
            case "\n\t":
                v87 := GenS6{F0: -1, F1: 1}
                _ = v87
            }
        case 9223372036854775807:
            for fi88 := 0; fi88 < 10; fi88++ {
                var v89 string = ""
                v90 := GenS7{F0: v3}
                var v91 bool = false
                _ = fi88
                _ = v89
                _ = v90
                _ = v91
            }
        case -1:
            fmt.Println(v86)
        default:
            var v92 int = fi85
            _ = v92
        }
        for fi93 := 0; fi93 < 17; fi93++ {
            if v1 {
                v94 := []int{}
                var v95 string = "a"
                _ = v94
                _ = v95
            }
            v96 := GenS8{F0: v4}
            var v97 uint64 = 0
            _ = fi93
            _ = v96
            _ = v97
        }
        _ = fi85
        _ = v86
    }
    v4 += (v4 << ((255) & 63))
    v98 := GenS9{F0: -1, F1: (!v83), F2: 454}
    fmt.Println(v1)
    v1 = ("hello" != "a")
    v99 := []int{32767}
    var v100 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v84
    _ = v98
    _ = v99
}
