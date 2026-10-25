package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64; F1 int8; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 int16; F1 uint8; F2 int16 }
type GenS10 struct { F0 bool; F1 int32; F2 int32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64; F1 int16; F2 uint }
func genFn0(p0 uint, p1 bool, p2 uint32) uint8 {
    v6 := GenS1{F0: p1, F1: (41 > 255)}
    switch -1 {
    case 32767:
        p0 = uint(p2)
    case 42:
        v7 := GenS2{F0: 127, F1: (^(p2)), F2: 0}
        _ = v7
    case 32768:
        switch 32767 {
        case 32768:
            if (!p1) {
                p0 -= 2
                p1 = false
            } else {
                if false {
                    v8 := GenS3{F0: -1, F1: 0, F2: -504}
                    v9 := []int{}
                    v10 := GenS2{F0: 128, F1: 0, F2: 519}
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    if true {
                        var v11 uint32 = (^(p2))
                        if p1 {
                            var v12 int64 = -1
                            var v13 uint32 = (p2 + (p2 | p2))
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: true, F1: (p1 && p1)}
                        _ = v11
                        _ = v14
                    }
                    fmt.Println(p0, p1, p1)
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS4{F0: 0}
                    fmt.Println(fi15, p2, fi15)
                    _ = fi15
                    _ = v16
                }
            }
        }
    default:
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 127
}
func genFn1() bool {
    var v17 uint64 = 4294967295
    if true {
        v17 += v17
        v18 := []int{-9223372036854775808, 1, 127, 1}
        var v19 int32 = 2147483647
        _ = v18
        _ = v19
    } else {
        var v20 int8 = 2
        var v21 uint32 = 2
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if false {
            var v23 string = (("" + "foo bar") + ("foo bar" + ""))
            var v24 uint = 18446744073709551615
            _ = v23
            _ = v24
        } else {
            v25 := []int{2147483647, fi22, 42}
            _ = v25
        }
        var v26 int = 32768
        _ = fi22
        _ = v26
    }
    switch (("foo bar" + ((("hello" + (("hello" + "foo bar") + "hello")) + "a") + "\n\t")) + "hello") {
    case "":
        switch -781 {
        case 2147483647:
            if false {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS2{F0: 0, F1: 880, F2: -1}
                    v29 := []int{}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v30 := []int{128, 128, 0}
                v17 -= v17
                _ = v30
            }
        case -811:
            var v31 int64 = 0
            _ = v31
        default:
            if false {
                if (v17 >= v17) {
                    var v32 uint8 = 1
                    if false {
                        v33 := []int{}
                        _ = v33
                    }
                    _ = v32
                }
                v17 = 4294967295
                var v34 int8 = 0
                _ = v34
            }
        }
    case "hello":
        v35 := GenS1{F0: (!false), F1: (0 <= 820)}
        _ = v35
    default:
        var v36 int = -2147483648
        _ = v36
    }
    _ = v17
    return (!false)
}

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = 1
    var v3 string = ("a" + "")
    v4 := make(map[int]uint32)
    v5 := []string{}
    var v37 uint64 = 959
    var v38 int32 = (v1 + v1)
    var v39 string = (("foo bar" + "") + v3)
    v40 := GenS1{F0: (!true), F1: (false || false)}
    v4[-860] = 0
    v3 = ("foo bar" + v3)
    switch 483 {
    case 42:
        v38 -= v38
    case 304:
        v3 = v3
    default:
        v41 := GenS5{F0: 1, F1: 169}
        _ = v41
    }
    for fi42 := 0; fi42 < 13; fi42++ {
        for fi43 := 0; fi43 < 10; fi43++ {
            switch (fi43 << ((0) & 63)) {
            case 38:
                v44 := GenS3{F0: 1, F1: v2, F2: (v1 >> ((386) & 63))}
                _ = v44
            case 128:
                v37 = v37
            case -694:
                v45 := GenS2{F0: 128, F1: 65535, F2: int32(v37)}
                _ = v45
            default:
                var v46 uint8 = 1
                _ = v46
            }
            _ = fi43
        }
        _ = fi42
    }
    var v47 float32 = 0.0
    delete(v4, 726)
    v47 = -2.25
    if false {
        var v48 int64 = -88
        for fi49 := 0; fi49 < 12; fi49++ {
            var v50 string = v3
            delete(v4, (fi49 / ((fi49) | 1)))
            _ = fi49
            _ = v50
        }
        for fi51 := 0; fi51 < 13; fi51++ {
            if (false || false) {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                switch -694 {
                case -668:
                    v4[fi51] = 0
                default:
                    fmt.Println(v2)
                }
                var v52 int64 = v48
                _ = v52
            } else {
                if (!(true || ((1e10 < (float64(1e10) + 1e10)) && (127 < 255)))) {
                    if ((v37 < (v37 / (((v37 & (v37 | (v37 >> ((4294967295) & 63))))) | 1))) || false) {
                        v53 := GenS6{F0: v48, F1: 65535}
                        if (!false) {
                            var v54 uint32 = 557
                            _ = v54
                        } else {
                            v55 := GenS6{F0: int64(v37), F1: 32767}
                            _ = v55
                        }
                        if (59 < 803) {
                            if true {
                                for fi56 := 0; fi56 < 8; fi56++ {
                                    fmt.Println(v38)
                                    switch (fi51 << ((42) & 63)) {
                                    case 0:
                                        var v57 uint16 = 538
                                        _ = v57
                                    default:
                                        var v58 string = v3
                                        _ = v58
                                    }
                                    v47 -= (float32((((3.14159 * (0.0 * 0.0)) * (3.14159 - -2.25)) + float64((0.0 + (3.14159 - ((v47 + v47) + (v47 * (0.0 / (((float32(-2.25) * 0.0) * v47) / v47))))))))) * 3.14159)
                                    _ = fi56
                                }
                                v38 -= v1
                            } else {
                                var v59 uint32 = 2147483647
                                if false {
                                    if true {
                                        var v60 int64 = 0
                                        v47 -= v47
                                        _ = v60
                                    } else {
                                        var v61 float64 = 0.0
                                        var v62 int16 = 32767
                                        _ = v61
                                        _ = v62
                                    }
                                    v63 := GenS7{F0: 149, F1: 108}
                                    _ = v63
                                } else {
                                    if (57 >= 934) {
                                        v2 = (v2 % (((v2 >> ((65535) & 63))) | 1))
                                    }
                                    if (-2.25 != 1e10) {
                                        v64 := GenS5{F0: 1, F1: 42}
                                        for fi65 := 0; fi65 < 16; fi65++ {
                                            if false {
                                                var v66 int64 = v48
                                                v67 := GenS1{F0: true, F1: false}
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v68 := GenS8{F0: 65535, F1: v38}
                                                _ = v68
                                            }
                                            var v69 uint32 = 2147483648
                                            switch "foo bar" {
                                            case "\n\t":
                                                if (!(true || true)) {
                                                    if (!(!(int8(v48) == v2))) {
                                                        fmt.Println(fi51, v38)
                                                    } else {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        var v70 int16 = 80
                                                        _ = v70
                                                    }
                                                    if true {
                                                        v71 := GenS9{F0: 458, F1: 1, F2: -1}
                                                        var v72 int8 = -53
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        var v73 int32 = (v38 ^ int32(v37))
                                                        if (v3 > v3) {
                                                            if false {
                                                                delete(v4, 9223372036854775807)
                                                                v47 = 0.0
                                                            } else {
                                                                var v74 uint = 857
                                                                _ = v74
                                                            }
                                                            v3 = v3
                                                            v37 = (v37 + (v37 | v37))
                                                        }
                                                        v75 := GenS10{F0: ((uint8(v37) >= 255) && (genFn0(632, (true && false), (^(v59))) < 0)), F1: v73, F2: 2147483647}
                                                        _ = v73
                                                        _ = v75
                                                    }
                                                    v2 = 38
                                                }
                                                v5 = append(v5, (("" + (v3 + ((v39 + ("hello" + v3)) + "foo bar"))) + ("" + v3)))
                                                var v76 uint32 = (v69 + v69)
                                                fmt.Println(fi51, fi65, v37)
                                                v77 := v4[fi51]
                                                _ = v76
                                                _ = v77
                                            case "hello":
                                                v78 := GenS11{F0: 175}
                                                _ = v78
                                            case "foo bar":
                                                var v79 uint = 135
                                                _ = v79
                                            }
                                            _ = fi65
                                            _ = v69
                                        }
                                        v80 := GenS12{F0: (v37 ^ v37), F1: -363, F2: 691}
                                        _ = v64
                                        _ = v80
                                    }
                                }
                                v3 = (v3 + ("" + "hello"))
                                _ = v59
                            }
                            if false {
                                for fi81 := 0; fi81 < 13; fi81++ {
                                    v82 := GenS6{F0: (v48 & (v48 | v48)), F1: 657}
                                    _ = fi81
                                    _ = v82
                                }
                            }
                        } else {
                            var v83 int64 = v48
                            _ = v83
                        }
                        _ = v53
                    } else {
                        if len(v5) > 0 {
                            v5[0] = v39
                        }
                    }
                    var v84 float32 = (0.0 * ((v47 + -2.25) + (1e10 - (((v47 / float32(-2.25)) + ((float32(1e10) * 0.0) + v47)) / v47))))
                    _ = v84
                }
                v47 += (v47 * (0.0 - float32(((1.5 * ((-2.25 * 0.0) - 0.0)) / -2.25))))
                switch v3 {
                case "foo bar":
                    v85 := GenS6{F0: (v48 - v48), F1: 0}
                    _ = v85
                case "":
                    var v86 uint8 = 186
                    _ = v86
                default:
                    var v87 int8 = v2
                    _ = v87
                }
            }
            for fi88 := 0; fi88 < 19; fi88++ {
                var v89 float64 = 1e10
                _ = fi88
                _ = v89
            }
            switch v3 {
            case "hello":
                v1 += (v1 << ((18446744073709551615) & 63))
            case "a":
                for fi90 := 0; fi90 < 13; fi90++ {
                    v5 = append(v5, "")
                    _ = fi90
                }
            case "\n\t":
                for fi91 := 0; fi91 < 17; fi91++ {
                    var v92 int8 = 0
                    _ = fi91
                    _ = v92
                }
            }
            _ = fi51
        }
        _ = v48
    }
    v38 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
}
