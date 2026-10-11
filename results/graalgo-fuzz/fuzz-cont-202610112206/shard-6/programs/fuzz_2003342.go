package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string; F1 int16; F2 uint }
type GenS3 struct { F0 int16; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 bool; F2 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int }
func genFn0(p0 uint, p1 uint32) int8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(fi6)
            v7 := []int{fi6, 2147483647, fi5, (fi5 >> (((^(p0))) & 63))}
            var v8 uint16 = 87
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    var v9 int8 = -1
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int32 = 848
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int64 = -740
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    if ((!false) || true) {
        var v14 int = -9223372036854775808
        fmt.Println(v14, v9)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint8 = 202
            v16 = 46
            _ = fi15
            _ = v16
        }
        v9 -= v9
    }
    var v17 int8 = 0
    _ = p0
    _ = p1
    _ = v17
    _ = v9
    return (v17 * 1)
}
func genFn1(p0 uint32, p1 int16, p2 uint32) int32 {
    v18 := GenS1{F0: 65535, F1: 42}
    p1 -= p1
    v19 := []int{-2147483648}
    var v20 uint16 = 139
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return -1
}

func main() {
    v1 := []string{}
    var v2 int8 = -33
    v3 := []string{"\n\t", "hello"}
    var v4 int = 42
    fmt.Println(v2, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 = append(v1, "hello")
        for fi22 := 0; fi22 < 11; fi22++ {
            v2 += (-(v2))
            _ = fi22
        }
        _ = fi21
    }
    var v23 int8 = (v2 >> ((65535) & 63))
    v24 := GenS1{F0: 2, F1: 9223372036854775807}
    switch "foo bar" {
    case "":
        switch (("" + ("foo bar" + ("foo bar" + "foo bar"))) + "foo bar") {
        case "\n\t":
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (((0 == 375) || (128 == (fi25 & -1))) && false) {
                        for fi26 := 0; fi26 < 8; fi26++ {
                            switch fi26 {
                            case 32767:
                                var v27 uint64 = 42
                                _ = v27
                            case 913:
                                v28 := GenS1{F0: 2, F1: (fi25 ^ (fi25 % (((fi25 * (v4 ^ fi25))) | 1)))}
                                _ = v28
                            case -1:
                                var v29 int16 = -1
                                _ = v29
                            default:
                                var v30 int16 = int16(fi25)
                                _ = v30
                            }
                            if true {
                                if (false || (65535 <= 2147483647)) {
                                    switch "a" {
                                    case "a":
                                        v23 = -53
                                    }
                                    var v31 float32 = 1e10
                                    var v32 int64 = 795
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := GenS1{F0: 0, F1: (v4 ^ fi25)}
                                    _ = v33
                                }
                                fmt.Println(fi26, v4)
                                if (("\n\t" + ("hello" + ("\n\t" + ("\n\t" + "\n\t")))) > "a") {
                                    v4 -= (-(fi26))
                                    v34 := GenS2{F0: "\n\t", F1: 32767, F2: 4294967295}
                                    v23 += v23
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: 1, F1: false, F2: ((v23 % ((v2) | 1)) > genFn0(278, 815))}
                                    v36 := GenS4{F0: "a", F1: (!false), F2: (v4 & 13)}
                                    var v37 uint = 712
                                    v4 = (fi25 << ((0) & 63))
                                    v37 = (v37 + v37)
                                    switch (^(fi25)) {
                                    case 1:
                                        if (false && false) {
                                            v3 = append(v3, "a")
                                            switch "\n\t" {
                                            case "a":
                                                var v38 uint8 = 255
                                                _ = v38
                                            }
                                        } else {
                                            v4 -= 705
                                            v3 = append(v3, "a")
                                            var v39 int32 = 1
                                            _ = v39
                                        }
                                    case 32767:
                                        var v40 int32 = -2147483648
                                        v41 := GenS3{F0: 535, F1: (!(((0.0 / -2.25) / 1.5) >= 0.1)), F2: false}
                                        v42 := GenS5{F0: (1 <= 2)}
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            }
                            _ = fi26
                        }
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v2 = 127
                            if false {
                                var v44 int64 = 0
                                _ = v44
                            } else {
                                v23 = (v23 ^ 127)
                                var v45 int16 = 1
                                v46 := GenS3{F0: (v45 / ((1) | 1)), F1: ((507 <= 897) && true), F2: (false && true)}
                                _ = v45
                                _ = v46
                            }
                            _ = fi43
                        }
                    } else {
                        v2 = (v2 ^ (v23 % ((v2) | 1)))
                    }
                    _ = fi25
                }
                v47 := GenS6{F0: 42, F1: 2147483647}
                _ = v47
            }
        case "foo bar":
            if true {
                v4 = (v4 + (v4 << ((42) & 63)))
                v2 = genFn0(18446744073709551615, 0)
                if (0 <= 2147483648) {
                    if (438 >= 1) {
                        switch "foo bar" {
                        case "foo bar":
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        case "":
                            v48 := GenS4{F0: "a", F1: true, F2: v4}
                            _ = v48
                        case "a":
                            for fi49 := 0; fi49 < 7; fi49++ {
                                switch -432 {
                                case 127:
                                    if (2 > 956) {
                                        switch (-(v4)) {
                                        case 127:
                                            _ = len(v3)
                                        case 1:
                                            switch ("hello" + (("a" + "\n\t") + "hello")) {
                                            case "a":
                                                if true {
                                                    var v50 int64 = 134
                                                    for fi51 := 0; fi51 < 15; fi51++ {
                                                        v50 -= 1
                                                        _ = len(v3)
                                                        _ = fi51
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                    _ = v50
                                                }
                                            default:
                                                switch (v4 << ((660) & 63)) {
                                                case -6:
                                                    for fi52 := 0; fi52 < 19; fi52++ {
                                                        v53 := GenS6{F0: -1, F1: 4294967295}
                                                        var v54 string = ""
                                                        v54 = "\n\t"
                                                        _ = fi52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                default:
                                                    var v55 uint = 532
                                                    _ = v55
                                                }
                                            }
                                        }
                                    } else {
                                        if false {
                                            v2 = (v2 ^ v2)
                                            var v56 float64 = 3.14159
                                            for fi57 := 0; fi57 < 5; fi57++ {
                                                v23 += (v2 << ((255) & 63))
                                                _ = fi57
                                            }
                                            _ = v56
                                        } else {
                                            switch (v4 >> ((1) & 63)) {
                                            case 718:
                                                var v58 int8 = (v2 + (v23 * (-(v2))))
                                                _ = v58
                                            case 1:
                                                v2 = v2
                                            }
                                            fmt.Println(v4, v2)
                                        }
                                        var v59 bool = false
                                        _ = v59
                                    }
                                case 9223372036854775807:
                                    var v60 bool = true
                                    _ = v60
                                }
                                if (true || (819 < 272)) {
                                    var v61 uint = 293
                                    var v62 uint = (^(v61))
                                    var v63 int32 = -823
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                v64 := GenS3{F0: 32767, F1: (!(-2147483648 != 42)), F2: (!((!(!(65535 <= 290))) && false))}
                                _ = fi49
                                _ = v64
                            }
                        }
                        v4 -= (v4 >> ((395) & 63))
                        v65 := GenS7{F0: 255}
                        _ = v65
                    }
                }
            }
        case "":
            if true {
                v1 = append(v1, "foo bar")
                v4 = (v4 >> ((65535) & 63))
                for fi66 := 0; fi66 < 18; fi66++ {
                    var v67 float32 = 3.14159
                    if true {
                        for fi68 := 0; fi68 < 16; fi68++ {
                            fmt.Println(fi66)
                            _ = fi68
                        }
                        v69 := GenS1{F0: 65535, F1: (fi66 | fi66)}
                        v4 = 2147483647
                        _ = v69
                    }
                    _ = fi66
                    _ = v67
                }
            }
        }
    }
    if (35 == (v23 >> ((65535) & 63))) {
        var v70 int64 = -59
        _ = v70
    }
    v71 := GenS3{F0: -1, F1: ((1e10 - 0.0) <= 0.0), F2: (-1 > (v2 - -128))}
    if false {
        _ = len(v1)
        var v72 bool = (!(genFn0(2, 0) != (v23 + (v23 - (v2 & (v23 - (v2 >> ((2) & 63))))))))
        if v72 {
            if false {
                var v73 int64 = 1
                for fi74 := 0; fi74 < 1; fi74++ {
                    v72 = true
                    fmt.Println(v23, fi74)
                    _ = fi74
                }
                for fi75 := 0; fi75 < 6; fi75++ {
                    fmt.Println(v2, v72)
                    _ = fi75
                }
                _ = v73
            } else {
                for fi76 := 0; fi76 < 3; fi76++ {
                    v77 := GenS8{F0: (v4 << ((2) & 63))}
                    _ = fi76
                    _ = v77
                }
                if (-546 < 42) {
                    v1 = append(v1, "foo bar")
                    v78 := GenS1{F0: 855, F1: v4}
                    _ = len(v1)
                    _ = v78
                }
                v2 = (v23 & (v23 * (v23 % (((v2 - genFn0(4294967295, 2))) | 1))))
            }
        }
        _ = v72
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v71
}
