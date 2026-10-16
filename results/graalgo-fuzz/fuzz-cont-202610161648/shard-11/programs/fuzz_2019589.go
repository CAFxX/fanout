package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
func genFn0() int16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{(fi6 | 1), (fi6 % ((fi6) | 1))}
        _ = fi6
        _ = v7
    }
    var v8 int32 = 101
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = 42
    var v3 int64 = -164
    var v4 string = "\n\t"
    v5 := make(map[string]int)
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 bool = false
            var v12 int32 = -851
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    v13 := GenS1{F0: v3, F1: true, F2: v4}
    if ((^(v3)) < 42) {
        v1 -= (v1 + v1)
        if ((true && false) || false) {
            switch "a" {
            case "foo bar":
                v14 := v5[("hello" + ((v4 + ("a" + ("hello" + v4))) + v4))]
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 int64 = (v3 & (v3 & 42))
                    fmt.Println(v16, v1, v16)
                    v1 = (v1 - (^(v1)))
                    _ = fi15
                    _ = v16
                }
            case "hello":
                v17 := GenS1{F0: v3, F1: (32767 > 127), F2: ""}
                _ = v17
            }
            switch 32767 {
            case 4294967295:
                v1 = (v1 << ((761) & 63))
            default:
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 uint16 = (v1 + (^(v1)))
                    _ = fi18
                    _ = v19
                }
            }
        } else {
            v20 := GenS1{F0: (v3 | v3), F1: ((670 < 2147483647) && (!true)), F2: (v4 + "\n\t")}
            v2 -= (v2 - 0)
            _ = v20
        }
        for fi21 := 0; fi21 < 17; fi21++ {
            v2 += (v2 ^ v2)
            if (true || true) {
                fmt.Println(v4, v3, v3)
            } else {
                fmt.Println(v4)
            }
            _ = fi21
        }
    }
    v22 := GenS1{F0: -1, F1: true, F2: v4}
    fmt.Println(v1, v2)
    var v23 int64 = (-(v3))
    for fi24 := 0; fi24 < 3; fi24++ {
        v3 += -689
        if ((^(v1)) > (v1 / (((v1 % (((v1 % (((v1 + v1)) | 1))) | 1))) | 1))) {
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS1{F0: -9223372036854775808, F1: (158 < 402), F2: "hello"}
                switch fi24 {
                case 128:
                    v27 := GenS2{F0: 2147483647}
                    _ = v27
                case -1:
                    var v28 int8 = v2
                    _ = v28
                case 0:
                    var v29 int64 = v23
                    _ = v29
                default:
                    var v30 float32 = 3.14159
                    _ = v30
                }
                _ = fi25
                _ = v26
            }
        } else {
            switch fi24 {
            case 127:
                switch "a" {
                case "a":
                    fmt.Println(v3, v1)
                default:
                    if false {
                        v31 := []int{fi24, (fi24 % ((fi24) | 1)), fi24, 127}
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v33 := GenS1{F0: 0, F1: ((v4 + "a") < v4), F2: ("a" + (v4 + (("\n\t" + "a") + ("" + v4))))}
                            var v34 int8 = int8(v1)
                            v35 := GenS1{F0: (v3 / ((v23) | 1)), F1: (2 < 126), F2: v4}
                            var v36 int32 = -686
                            _ = fi32
                            _ = v33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        var v37 int = fi24
                        var v38 float32 = 1e10
                        _ = v37
                        _ = v38
                    }
                }
            default:
                if ((((float32((1.5 + 3.14159)) * -2.25) * 1.5) > 0.0) && (-9223372036854775808 == v23)) {
                    v39 := GenS1{F0: (v23 * (-(v3))), F1: (!true), F2: v4}
                    _ = v39
                } else {
                    delete(v5, (((((v4 + ((((v4 + "") + "foo bar") + v4) + v4)) + "foo bar") + v4) + ("\n\t" + "a")) + v4))
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 int8 = v2
                        v42 := GenS1{F0: int64(fi24), F1: (65535 != 1), F2: v4}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v43 uint8 = 42
                    _ = v43
                }
            }
            if false {
                switch (v4 + (("hello" + "a") + (v4 + v4))) {
                case "foo bar":
                    v44 := GenS1{F0: v23, F1: false, F2: v4}
                    _ = v44
                case "hello":
                    if false {
                        var v45 uint16 = 32767
                        if true {
                            v46 := GenS3{F0: 0, F1: v4}
                            _ = v46
                        } else {
                            for fi47 := 0; fi47 < 8; fi47++ {
                                v48 := GenS1{F0: (v23 - (v3 & -9223372036854775808)), F1: (-1 <= -32768), F2: ("a" + "a")}
                                var v49 string = (((v4 + v4) + ("\n\t" + "a")) + "\n\t")
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            var v50 uint32 = 0
                            v4 = v4
                            _ = v50
                        }
                        _ = v45
                    } else {
                        if false {
                            for fi51 := 0; fi51 < 17; fi51++ {
                                v4 = v4
                                _ = fi51
                            }
                            fmt.Println(fi24)
                        } else {
                            if (!(128 <= 255)) {
                                for fi52 := 0; fi52 < 19; fi52++ {
                                    switch "hello" {
                                    case "hello":
                                        if (!((-2147483648 == 1) || false)) {
                                            if (1 == v1) {
                                                var v53 uint64 = 4294967295
                                                _ = v53
                                            } else {
                                                fmt.Println(v23, v23, v1)
                                                fmt.Println(fi52)
                                            }
                                            var v54 int16 = 42
                                            _ = v54
                                        } else {
                                            var v55 uint16 = v1
                                            var v56 string = "a"
                                            _ = v55
                                            _ = v56
                                        }
                                    case "foo bar":
                                        switch (((v4 + v4) + "hello") + ((("a" + "foo bar") + (v4 + (("foo bar" + v4) + "foo bar"))) + "foo bar")) {
                                        case "foo bar":
                                            switch v4 {
                                            case "foo bar":
                                                var v57 uint32 = 4294967295
                                                _ = v57
                                            case "a":
                                                switch (v4 + v4) {
                                                case "":
                                                    var v58 uint16 = v1
                                                    _ = v58
                                                }
                                            }
                                        case "":
                                            var v59 uint = 0
                                            _ = v59
                                        case "\n\t":
                                            v60 := GenS3{F0: 65535, F1: v4}
                                            _ = v60
                                        default:
                                            var v61 uint64 = 252
                                            _ = v61
                                        }
                                    case "":
                                        v62 := GenS2{F0: 32767}
                                        _ = v62
                                    }
                                    v2 = v2
                                    v3 -= 42
                                    _ = fi52
                                }
                                fmt.Println(v2)
                                var v63 bool = (3.14159 > 3.14159)
                                _ = v63
                            }
                            v64 := []int{}
                            _ = v64
                        }
                        for fi65 := 0; fi65 < 6; fi65++ {
                            v66 := GenS4{F0: (v2 & (v2 % (((v2 - v2)) | 1)))}
                            if false {
                                fmt.Println(v3, v3, v1)
                                v5[v4] = fi65
                                var v67 int32 = 0
                                _ = v67
                            } else {
                                switch fi65 {
                                case -2147483648:
                                    var v68 uint8 = 127
                                    _ = v68
                                case -277:
                                    v69 := GenS1{F0: v23, F1: true, F2: v4}
                                    _ = v69
                                case -490:
                                    var v70 float64 = 1.5
                                    _ = v70
                                default:
                                    v71 := GenS5{F0: -1, F1: (v3 + v3)}
                                    _ = v71
                                }
                            }
                            var v72 bool = (!false)
                            _ = fi65
                            _ = v66
                            _ = v72
                        }
                        switch 4294967295 {
                        case 2147483647:
                            for fi73 := 0; fi73 < 2; fi73++ {
                                v74 := []int{0, fi24, (fi24 * (^(fi73))), (fi24 ^ (fi73 & 0))}
                                if false {
                                    v1 = (v1 % (((v1 % ((2) | 1))) | 1))
                                } else {
                                    switch (fi24 + (fi24 - fi73)) {
                                    case 0:
                                        var v75 uint32 = 1
                                        _ = v75
                                    default:
                                        v1 += (v1 ^ 32767)
                                    }
                                    var v76 string = "foo bar"
                                    _ = v76
                                }
                                _ = fi73
                                _ = v74
                            }
                        default:
                            if (v1 <= 32767) {
                                v3 -= -9223372036854775808
                                switch "hello" {
                                case "\n\t":
                                    for fi77 := 0; fi77 < 19; fi77++ {
                                        var v78 int64 = -9223372036854775808
                                        v23 = -1
                                        _ = fi77
                                        _ = v78
                                    }
                                }
                                for fi79 := 0; fi79 < 11; fi79++ {
                                    if true {
                                        switch fi24 {
                                        case 9223372036854775807:
                                            if true {
                                                v23 = -9223372036854775808
                                            } else {
                                                v80 := GenS6{F0: 255}
                                                _ = v80
                                            }
                                            var v81 float32 = 0.1
                                            _ = v81
                                        case 32768:
                                            v82 := GenS5{F0: -2147483648, F1: v3}
                                            var v83 float32 = (1e10 - float32(0.0))
                                            v1 = (^(v1))
                                            _ = v82
                                            _ = v83
                                        case 446:
                                            var v84 float32 = (0.1 * 0.1)
                                            _ = v84
                                        }
                                    } else {
                                        v1 = 321
                                        v2 = -121
                                        v85 := GenS6{F0: 163}
                                        _ = v85
                                    }
                                    _ = fi79
                                }
                            }
                        }
                    }
                default:
                    var v86 uint8 = 7
                    _ = v86
                }
                switch -152 {
                case 9223372036854775807:
                    var v87 float32 = 0.0
                    _ = v87
                case 149:
                    var v88 int64 = (^(v23))
                    _ = v88
                case -9223372036854775808:
                    v2 = (^(v2))
                }
            }
        }
        _ = fi24
    }
    v4 = v4
    var v89 uint32 = 681
    v90 := GenS5{F0: -377, F1: (v23 % ((v23) | 1))}
    switch "" {
    case "":
        v3 = v3
    default:
        if false {
            v5[("\n\t" + "\n\t")] = -1
            var v91 uint64 = 42
            fmt.Println(v89, v3, v91)
            _ = v91
        } else {
            var v92 float32 = 3.14159
            fmt.Println(v3)
            v89 = 245
            _ = v92
        }
    }
    if false {
        v93 := GenS2{F0: 1}
        v2 += (-(v2))
        _ = v93
    } else {
        fmt.Println(v89)
        v94 := GenS7{F0: 65535}
        switch "foo bar" {
        case "foo bar":
            v95 := []int{}
            _ = v95
        default:
            v89 = 1
        }
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v89
    _ = v90
}
