package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }
type GenS5 struct { F0 int64; F1 int16; F2 int8 }
type GenS6 struct { F0 uint64; F1 int8; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 uint16 }
func genFn0(p0 int32) int16 {
    var v5 int32 = 251
    var v6 uint32 = 2
    v5 -= -560
    for fi7 := 0; fi7 < 16; fi7++ {
        if (false || ((1.5 != -2.25) || true)) {
            v8 := GenS1{F0: 2, F1: v6}
            v9 := []int{(-(fi7))}
            v10 := GenS2{F0: "a", F1: true}
            _ = v10
            _ = v8
            _ = v9
        } else {
            if true {
                if (42 != 127) {
                    if ((v6 ^ v6) > (v6 ^ (v6 * (v6 - (v6 / ((v6) | 1)))))) {
                        for fi11 := 0; fi11 < 11; fi11++ {
                            v12 := []int{(fi11 & fi11), fi11, (^(fi11)), fi7}
                            var v13 uint = 255
                            p0 -= 931
                            var v14 float64 = (0.1 + float64(1.5))
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    }
                    var v15 int = (fi7 - 0)
                    _ = v15
                }
                v16 := GenS3{F0: (v5 ^ (p0 - v5))}
                _ = v16
            }
            var v17 uint = 1
            for fi18 := 0; fi18 < 3; fi18++ {
                v5 = (p0 >> ((v17) & 63))
                if false {
                    if (uint8(v17) >= 128) {
                        if ((^(v5)) <= (v5 << ((478) & 63))) {
                            if false {
                                v19 := GenS1{F0: 524, F1: 1}
                                fmt.Println(v5, v6, fi7)
                                v20 := GenS4{F0: (fi18 << ((v17) & 63)), F1: -9223372036854775808, F2: 0}
                                _ = v19
                                _ = v20
                            } else {
                                v17 += (v17 * (v17 | v17))
                                v21 := GenS2{F0: "hello", F1: true}
                                _ = v21
                            }
                            var v22 uint16 = 1
                            switch ("" + "hello") {
                            case "":
                                v5 = (p0 & 0)
                            case "\n\t":
                                v23 := GenS1{F0: v22, F1: v6}
                                v24 := GenS1{F0: (v22 % (((v22 & v22)) | 1)), F1: v6}
                                _ = v23
                                _ = v24
                            case "foo bar":
                                if (-1 > -9223372036854775808) {
                                    switch (fi7 % ((fi7) | 1)) {
                                    case 4294967295:
                                        v25 := []int{4294967295, 410, -291, (^(fi18))}
                                        _ = v25
                                    case 32767:
                                        v6 = (v6 << (((v17 | (v17 | (^(v17))))) & 63))
                                    case 32768:
                                        switch 4294967295 {
                                        case 42:
                                            if (!false) {
                                                var v26 bool = (!(-58 <= -1))
                                                _ = v26
                                            }
                                        case -549:
                                            v27 := []int{-1, (fi7 | fi18), (fi18 % ((fi7) | 1))}
                                            _ = v27
                                        default:
                                            var v28 float32 = float32((3.14159 * (3.14159 + 1.5)))
                                            _ = v28
                                        }
                                    default:
                                        fmt.Println(v22)
                                    }
                                    var v29 int32 = (v5 * (p0 & 407))
                                    var v30 float64 = 1e10
                                    var v31 int8 = -128
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            default:
                                fmt.Println(v22)
                            }
                            _ = v22
                        }
                    }
                    if (1 < 42) {
                        if true {
                            var v32 uint64 = 1
                            v32 = (v32 & 1)
                            fmt.Println(v5, p0)
                            _ = v32
                        } else {
                            v33 := GenS5{F0: 42, F1: 0, F2: 96}
                            v34 := []int{128, fi7}
                            _ = v33
                            _ = v34
                        }
                        if ("\n\t" != "hello") {
                            v35 := GenS2{F0: "\n\t", F1: true}
                            var v36 int16 = -32768
                            v37 := []int{}
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            if false {
                                var v38 int = (fi7 + (fi7 << ((v17) & 63)))
                                var v39 int32 = (v5 << ((v17) & 63))
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint32 = v6
                                var v41 uint64 = 169
                                _ = v40
                                _ = v41
                            }
                            var v42 uint64 = 4294967295
                            _ = v42
                        }
                        var v43 int16 = 42
                        _ = v43
                    } else {
                        v5 += (p0 | (v5 + -1))
                        v44 := GenS2{F0: "", F1: true}
                        _ = v44
                    }
                    var v45 int8 = -120
                    _ = v45
                } else {
                    var v46 uint8 = 42
                    v46 -= v46
                    _ = v46
                }
                _ = fi18
            }
            _ = v17
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 680
    v2 := make(map[int]int)
    var v3 uint16 = 42
    v4 := make(map[uint32]bool)
    if (3.14159 < (3.14159 * -2.25)) {
        v3 = v3
    } else {
        var v47 int8 = -1
        _ = v47
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        switch (fi48 ^ int(v3)) {
        case -993:
            var v49 bool = true
            _ = v49
        default:
            if true {
                if true {
                    v1 = 18446744073709551615
                }
            } else {
                var v50 int16 = -1
                fmt.Println(v1)
                _ = v50
            }
        }
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 bool = (-32768 >= genFn0(-172))
            fmt.Println(v3)
            switch "hello" {
            case "hello":
                var v53 bool = false
                _ = v53
            default:
                if (!true) {
                    v54 := []int{fi51, fi48, int(v3), (fi48 % ((fi51) | 1))}
                    if v52 {
                        v55 := GenS3{F0: -2147483648}
                        _ = v55
                    }
                    if (false && v52) {
                        v56 := GenS6{F0: (v1 << ((255) & 63)), F1: 0, F2: (!(false && false))}
                        _ = v56
                    }
                    _ = v54
                } else {
                    if (0 != 127) {
                        var v57 float64 = 1.5
                        switch 1 {
                        case -848:
                            var v58 int8 = -37
                            _ = v58
                        case 2147483647:
                            _ = len(v4)
                        case -1:
                            switch fi48 {
                            case -2147483648:
                                v59 := []int{(fi48 | 4294967295)}
                                _ = v59
                            }
                        default:
                            var v60 uint64 = v1
                            _ = v60
                        }
                        v61 := GenS2{F0: ("a" + (("" + "foo bar") + "a")), F1: v52}
                        _ = v57
                        _ = v61
                    } else {
                        switch "hello" {
                        case "":
                            switch fi48 {
                            case -2147483648:
                                v62 := GenS7{F0: 2}
                                _ = v62
                            case 417:
                                if (!v52) {
                                    var v63 uint = 65535
                                    switch "a" {
                                    case "":
                                        v52 = v52
                                    case "foo bar":
                                        v3 += v3
                                    case "\n\t":
                                        v64 := v2[-2147483648]
                                        _ = v64
                                    default:
                                        v52 = (v52 && v52)
                                        delete(v2, fi51)
                                    }
                                    _ = v63
                                }
                            case 32768:
                                delete(v2, 2147483647)
                            default:
                                v65 := GenS6{F0: v1, F1: -128, F2: v52}
                                _ = v65
                            }
                        default:
                            var v66 uint16 = 65535
                            _ = v66
                        }
                        v1 -= (v1 - (v1 - (v1 * 42)))
                    }
                    v67 := GenS5{F0: -1, F1: -32768, F2: -128}
                    _ = v67
                }
            }
            _ = fi51
            _ = v52
        }
        var v68 uint16 = v3
        _ = fi48
        _ = v68
    }
    fmt.Println(v3)
    v69 := GenS7{F0: 128}
    v3 = (v3 + v3)
    v70 := GenS8{F0: 1000, F1: v3}
    v71 := GenS4{F0: 32768, F1: -763, F2: (v3 + 0)}
    var v72 uint32 = 2147483648
    fmt.Println(v3)
    fmt.Println(v72, v1, v72)
    fmt.Println(v72)
    for fi73 := 0; fi73 < 12; fi73++ {
        v74 := []int{fi73, fi73, (fi73 + (fi73 ^ 32768)), (fi73 + (fi73 + -9223372036854775808))}
        for fi75 := 0; fi75 < 12; fi75++ {
            var v76 float64 = ((-2.25 - 3.14159) / 1.0)
            _ = fi75
            _ = v76
        }
        var v77 int16 = 0
        _ = fi73
        _ = v74
        _ = v77
    }
    for fi78 := 0; fi78 < 8; fi78++ {
        if true {
            for fi79 := 0; fi79 < 1; fi79++ {
                v1 = v1
                _ = fi79
            }
            fmt.Println(v3, v3, v3)
        } else {
            v80 := []int{int(v72)}
            v72 = 2
            var v81 int64 = 9223372036854775807
            _ = v80
            _ = v81
        }
        _ = fi78
    }
    switch 42 {
    case -1:
        for fi82 := 0; fi82 < 16; fi82++ {
            if true {
                if (!(((("foo bar" + ("\n\t" + "foo bar")) + (("" + "foo bar") + "hello")) + ("a" + ("a" + "foo bar"))) >= ("\n\t" + "hello"))) {
                    v1 += v1
                    v3 -= v3
                } else {
                    var v83 bool = ((genFn0(-1) != -281) && true)
                    _ = v83
                }
                v3 -= v3
            }
            var v84 float32 = (1.5 - 0.1)
            if false {
                var v85 int16 = 11
                _ = v85
            } else {
                v1 -= uint64(v3)
                var v86 uint16 = (v3 & v3)
                v87 := []int{}
                _ = v86
                _ = v87
            }
            _ = fi82
            _ = v84
        }
    case -567:
        for fi88 := 0; fi88 < 4; fi88++ {
            v89 := GenS5{F0: -424, F1: -468, F2: 0}
            _ = fi88
            _ = v89
        }
    case 2147483647:
        v3 += 889
    default:
        var v90 int = 42
        _ = v90
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
