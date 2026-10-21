package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint }
func genFn0(p0 int32, p1 uint32) int16 {
    fmt.Println(p0, p1)
    v7 := []int{42, -174, -1}
    _ = p0
    _ = p1
    _ = v7
    return -805
}
func genFn1(p0 uint16, p1 uint16) uint64 {
    fmt.Println(p1, p1)
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 0
    var v2 bool = false
    var v3 int32 = 42
    v4 := []int{32767, 32767}
    var v5 int64 = -1
    v6 := make(map[int]string)
    v2 = false
    var v9 string = "hello"
    v10 := GenS1{F0: v3, F1: 1, F2: 32767}
    if ((^(v3)) != -2147483648) {
        v3 -= 1
    }
    v2 = true
    if v2 {
        v3 -= -1
        for fi11 := 0; fi11 < 11; fi11++ {
            if len(v4) > 0 {
                v4[0] = fi11
            }
            _ = fi11
        }
    }
    v12 := GenS2{F0: (^(v3)), F1: 0, F2: 1}
    v2 = true
    var v13 int64 = v5
    if true {
        var v14 float64 = 1.5
        _ = len(v6)
        for fi15 := 0; fi15 < 4; fi15++ {
            v13 += (^(v5))
            fmt.Println(fi15)
            _ = fi15
        }
        _ = v14
    } else {
        switch -722 {
        case 32767:
            v5 = 0
        case 556:
            v4 = append(v4, 4294967295)
        case -208:
            for fi16 := 0; fi16 < 6; fi16++ {
                if v2 {
                    for fi17 := 0; fi17 < 9; fi17++ {
                        v18 := GenS1{F0: v3, F1: v3, F2: (fi16 ^ fi16)}
                        switch "" {
                        case "hello":
                            v19 := v6[fi17]
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                    var v20 uint = 18446744073709551615
                    v1 = (v1 * 174)
                    _ = v20
                }
                if (!false) {
                    if true {
                        switch (fi16 << ((18446744073709551615) & 63)) {
                        case -849:
                            if (0 == 42) {
                                v5 = v5
                            }
                        case 42:
                            for fi21 := 0; fi21 < 11; fi21++ {
                                v5 = -9223372036854775808
                                _ = fi21
                            }
                        case 9223372036854775807:
                            if v2 {
                                var v22 bool = (2 < 42)
                                var v23 bool = v2
                                v2 = (false && (42 <= 1))
                                _ = v22
                                _ = v23
                            } else {
                                for fi24 := 0; fi24 < 7; fi24++ {
                                    v5 = v5
                                    switch (^(fi24)) {
                                    case 240:
                                        var v25 uint16 = 32767
                                        _ = v25
                                    case -2147483648:
                                        v4 = append(v4, fi16)
                                    default:
                                        var v26 float32 = 1.5
                                        _ = v26
                                    }
                                    _ = fi24
                                }
                                fmt.Println(v9)
                            }
                        default:
                            if (!v2) {
                                v1 = 199
                                var v27 uint = 23
                                _ = v27
                            } else {
                                fmt.Println(v5, fi16, v13)
                            }
                        }
                        v5 += v5
                        v5 = (-(v13))
                    } else {
                        for fi28 := 0; fi28 < 12; fi28++ {
                            var v29 float64 = 0.0
                            v13 = v5
                            _ = fi28
                            _ = v29
                        }
                        var v30 uint = 2
                        _ = v30
                    }
                    if (v13 < (v5 | v5)) {
                        for fi31 := 0; fi31 < 16; fi31++ {
                            v2 = false
                            v3 = (v3 / ((v3) | 1))
                            if v2 {
                                switch -501 {
                                case 32767:
                                    fmt.Println(v9)
                                default:
                                    fmt.Println(fi16, v3)
                                }
                                v1 = v1
                            }
                            _ = fi31
                        }
                        var v32 float64 = (0.0 * ((0.0 - -2.25) + (0.0 + 0.0)))
                        _ = v32
                    } else {
                        var v33 uint64 = 1
                        v3 += (-(v3))
                        var v34 int = fi16
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v35 uint8 = v1
                    var v36 bool = (false && v2)
                    var v37 float32 = -2.25
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi16
            }
        default:
            v38 := GenS1{F0: (v3 & v3), F1: v3, F2: 9223372036854775807}
            _ = v38
        }
        if v2 {
            v39 := GenS2{F0: (-(v3)), F1: v1, F2: 4294967295}
            for fi40 := 0; fi40 < 8; fi40++ {
                v2 = (2147483647 <= 4294967295)
                if (false && (255 >= 1)) {
                    switch ("hello" + "foo bar") {
                    case "":
                        fmt.Println(v3, v1, v2)
                    case "foo bar":
                        fmt.Println(v5, v5, v2)
                    case "\n\t":
                        v9 = (v9 + "")
                    default:
                        v9 = "a"
                    }
                    if ((-1 <= (v13 ^ v13)) && false) {
                        if v2 {
                            if (127 == -49) {
                                var v41 uint32 = 2147483647
                                v2 = true
                                v41 = v41
                                _ = v41
                            } else {
                                var v42 uint64 = 2
                                _ = v42
                            }
                        } else {
                            switch fi40 {
                            case 258:
                                var v43 int8 = int8(fi40)
                                _ = v43
                            case 42:
                                switch 42 {
                                case 128:
                                    switch (v9 + (v9 + v9)) {
                                    case "\n\t":
                                        v44 := GenS3{F0: 255, F1: 65535}
                                        _ = v44
                                    case "hello":
                                        v45 := GenS3{F0: 143, F1: 1}
                                        _ = v45
                                    default:
                                        if v2 {
                                            var v46 float32 = 3.14159
                                            _ = v46
                                        } else {
                                            switch (fi40 - (fi40 & (fi40 | (-(fi40))))) {
                                            case 127:
                                                if v2 {
                                                    var v47 uint8 = v1
                                                    _ = v47
                                                } else {
                                                    switch fi40 {
                                                    case -552:
                                                        v48 := GenS3{F0: 255, F1: 160}
                                                        _ = v48
                                                    default:
                                                        fmt.Println(v2, v1, v3)
                                                    }
                                                    v1 = (v1 - 127)
                                                    fmt.Println(v5)
                                                }
                                            case 128:
                                                if (v2 && false) {
                                                    v1 = v1
                                                } else {
                                                    v9 = v9
                                                }
                                            case 32768:
                                                for fi49 := 0; fi49 < 19; fi49++ {
                                                    switch (fi49 / ((fi49) | 1)) {
                                                    case -858:
                                                        switch fi49 {
                                                        case 486:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                var v50 string = "foo bar"
                                                                _ = v50
                                                            case "\n\t":
                                                                var v51 uint = 65535
                                                                _ = v51
                                                            }
                                                        case 32768:
                                                            fmt.Println(v2)
                                                        case -105:
                                                            var v52 int16 = 32767
                                                            _ = v52
                                                        }
                                                    case 1:
                                                        v1 -= (v1 / ((128) | 1))
                                                    case 4294967295:
                                                        var v53 int16 = 0
                                                        _ = len(v6)
                                                        switch v9 {
                                                        case "":
                                                            var v54 uint8 = (v1 - 1)
                                                            _ = v54
                                                        case "hello":
                                                            var v55 float64 = 0.0
                                                            _ = v55
                                                        }
                                                        _ = v53
                                                    }
                                                    fmt.Println(v1)
                                                    _ = fi49
                                                }
                                            }
                                        }
                                    }
                                case -130:
                                    switch fi40 {
                                    case -2147483648:
                                        for fi56 := 0; fi56 < 8; fi56++ {
                                            if v2 {
                                                fmt.Println(fi56, v13)
                                                var v57 int64 = v13
                                                if v2 {
                                                    var v58 uint = 869
                                                    if len(v4) > 0 {
                                                        v4[0] = fi40
                                                    }
                                                    var v59 int = -366
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v57
                                            }
                                            fmt.Println(v3, v1, v2)
                                            v3 += v3
                                            _ = fi56
                                        }
                                    case -1:
                                        var v60 uint = 65535
                                        _ = v60
                                    }
                                case 0:
                                    var v61 uint64 = 1
                                    _ = v61
                                default:
                                    var v62 int16 = -1
                                    _ = v62
                                }
                            default:
                                var v63 int32 = v3
                                _ = v63
                            }
                        }
                    } else {
                        if (v2 || v2) {
                            for fi64 := 0; fi64 < 20; fi64++ {
                                var v65 int = fi64
                                var v66 bool = (2147483648 >= 2)
                                var v67 string = "foo bar"
                                _ = fi64
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                        } else {
                            if (int64(fi40) >= 9223372036854775807) {
                                if ((127 > 100) || (false || false)) {
                                    v9 = ("a" + "")
                                }
                            } else {
                                v2 = (v2 || v2)
                                switch fi40 {
                                case -9223372036854775808:
                                    var v68 string = v9
                                    _ = v68
                                case 128:
                                    v1 = (v1 - v1)
                                }
                                delete(v6, fi40)
                            }
                            var v69 uint8 = 76
                            _ = v69
                        }
                        if ("\n\t" < ("hello" + "foo bar")) {
                            if (1 >= 19) {
                                v70 := GenS1{F0: (v3 + v3), F1: v3, F2: fi40}
                                if v2 {
                                    fmt.Println(fi40, fi40, v3)
                                    v71 := v6[-202]
                                    _ = v71
                                }
                                var v72 int64 = -189
                                _ = v70
                                _ = v72
                            } else {
                                fmt.Println(v5, v1, fi40)
                            }
                            for fi73 := 0; fi73 < 13; fi73++ {
                                fmt.Println(v2, v1)
                                _ = fi73
                            }
                            v74 := GenS4{F0: ((^(v13)) <= (v5 * (v5 - (v13 ^ v5))))}
                            _ = v74
                        }
                        v75 := GenS5{F0: v1, F1: (v1 * (v1 / (((v1 + 42)) | 1))), F2: 100}
                        _ = v75
                    }
                    v76 := GenS5{F0: (v1 >> ((18446744073709551615) & 63)), F1: v1, F2: -11}
                    _ = v76
                } else {
                    v3 = v3
                }
                v77 := GenS4{F0: ((float32(0.0) / 1.5) >= 3.14159)}
                _ = fi40
                _ = v77
            }
            for fi78 := 0; fi78 < 20; fi78++ {
                fmt.Println(v2, v3)
                switch (v9 + ((v9 + v9) + "hello")) {
                case "foo bar":
                    v79 := GenS4{F0: (!(!v2))}
                    _ = v79
                case "":
                    v80 := GenS6{F0: 83}
                    _ = v80
                default:
                    var v81 bool = v2
                    _ = v81
                }
                _ = fi78
            }
            _ = v39
        }
        fmt.Println(v3)
    }
    v82 := GenS3{F0: 916, F1: 2}
    var v83 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v9
}
