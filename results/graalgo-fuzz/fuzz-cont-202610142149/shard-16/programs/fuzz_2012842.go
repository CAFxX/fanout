package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 int; F2 string }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 int8 = 100
    var v2 uint8 = 2
    v3 := make(map[uint32]string)
    var v4 int = 4294967295
    var v5 int8 = (v1 ^ v1)
    v6 := GenS1{F0: -9223372036854775808}
    for fi7 := 0; fi7 < 5; fi7++ {
        v5 = v1
        if true {
            var v8 float64 = 0.0
            v9 := GenS2{F0: 2, F1: 127}
            v10 := GenS1{F0: -9223372036854775808}
            _ = v10
            _ = v8
            _ = v9
        }
        v11 := GenS2{F0: 2147483648, F1: -101}
        _ = fi7
        _ = v11
    }
    var v12 uint8 = v2
    if (true && true) {
        switch ("hello" + "hello") {
        case "\n\t":
            var v13 uint8 = (v12 + (v2 - (^(v2))))
            _ = v13
        case "foo bar":
            var v14 int8 = (v5 << ((2) & 63))
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                if true {
                    var v17 int32 = 2147483647
                    _ = v17
                }
                _ = fi16
            }
            fmt.Println(v1)
            fmt.Println(fi15, v2, v4)
            _ = fi15
        }
        v4 = -1
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v5 -= 0
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 int8 = 100
            v21 := GenS3{F0: 101, F1: v20}
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := GenS2{F0: 65535, F1: 127}
        _ = fi18
        _ = v22
    }
    var v23 uint8 = v2
    fmt.Println(v5)
    fmt.Println(v4, v12)
    if false {
        if (false && (65535 != 42)) {
            switch (v4 * (-(v4))) {
            case 42:
                var v24 int64 = -40
                _ = v24
            case 4294967295:
                fmt.Println(v12, v5)
            case -2147483648:
                for fi25 := 0; fi25 < 20; fi25++ {
                    if ((v2 * (v12 << ((1) & 63))) == v12) {
                        var v26 int64 = 0
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v5 += 1
                            _ = fi27
                        }
                        var v28 int8 = v1
                        _ = v28
                    }
                    _ = fi25
                }
            }
        } else {
            var v29 uint8 = (^(v2))
            var v30 int64 = 42
            var v31 bool = false
            _ = v29
            _ = v30
            _ = v31
        }
    }
    delete(v3, 2147483648)
    if (false && ((0 <= 2) || false)) {
        switch 0 {
        case -788:
            var v32 int16 = 1
            _ = v32
        case 32768:
            v33 := GenS3{F0: 4294967295, F1: (v1 % ((v5) | 1))}
            _ = v33
        case -2147483648:
            var v34 float64 = 0.0
            _ = v34
        default:
            v35 := GenS4{F0: 4294967295, F1: 9223372036854775807}
            _ = v35
        }
        if true {
            v4 -= (v4 % ((-683) | 1))
            if true {
                for fi36 := 0; fi36 < 17; fi36++ {
                    v37 := GenS1{F0: 356}
                    _ = fi36
                    _ = v37
                }
                var v38 int8 = v5
                v23 += 1
                _ = v38
            } else {
                for fi39 := 0; fi39 < 8; fi39++ {
                    v40 := GenS2{F0: 577, F1: -1}
                    for fi41 := 0; fi41 < 5; fi41++ {
                        v42 := GenS5{F0: false, F1: 355, F2: 65535}
                        if ("" >= "\n\t") {
                            v43 := GenS6{F0: 0}
                            _ = v43
                        } else {
                            switch "" {
                            case "\n\t":
                                var v44 int16 = -1
                                _ = v44
                            default:
                                var v45 int = (-(v4))
                                _ = v45
                            }
                            var v46 bool = false
                            v4 -= 42
                            _ = v46
                        }
                        if false {
                            switch "foo bar" {
                            case "a":
                                var v47 int32 = 457
                                _ = v47
                            }
                            v23 = v23
                            if ((0 == 9223372036854775807) || true) {
                                var v48 int32 = -1
                                v5 = v5
                                var v49 uint16 = 2
                                _ = v48
                                _ = v49
                            }
                            var v50 bool = true
                            _ = v50
                        } else {
                            switch v4 {
                            case 32767:
                                v23 = (v12 + 127)
                            case 1:
                                if (false && true) {
                                    v51 := GenS7{F0: v5, F1: 669}
                                    var v52 int = -1
                                    var v53 uint16 = 1
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 uint8 = v23
                                    v55 := GenS8{F0: 42, F1: v5}
                                    v4 = (^(fi41))
                                    fmt.Println(v5, fi41, fi39)
                                    v1 = v5
                                    _ = v54
                                    _ = v55
                                }
                            }
                            if true {
                                v56 := GenS8{F0: v23, F1: v1}
                                v57 := GenS9{F0: fi41}
                                _ = v56
                                _ = v57
                            }
                        }
                        v58 := v3[4294967295]
                        _ = fi41
                        _ = v42
                        _ = v58
                    }
                    _ = fi39
                    _ = v40
                }
            }
        } else {
            switch v4 {
            case -2147483648:
                if true {
                    if (true && (true && (!false))) {
                        var v59 int = (^(v4))
                        if false {
                            v59 -= (-(v4))
                            switch ("\n\t" + "a") {
                            case "":
                                var v60 uint8 = v12
                                _ = v60
                            case "hello":
                                for fi61 := 0; fi61 < 17; fi61++ {
                                    v59 -= (^(fi61))
                                    v5 = v1
                                    var v62 uint16 = 255
                                    _ = fi61
                                    _ = v62
                                }
                            default:
                                switch "a" {
                                case "hello":
                                    v63 := GenS8{F0: 255, F1: (v1 << ((405) & 63))}
                                    _ = v63
                                }
                            }
                        } else {
                            switch (v4 - v4) {
                            case 42:
                                v64 := []int{v59, (v4 | (-(v4))), (-(v4)), int(v23)}
                                _ = v64
                            case -2147483648:
                                var v65 int8 = -128
                                _ = v65
                            case -450:
                                if (!true) {
                                    v4 += v4
                                } else {
                                    if (true || (-57 <= v1)) {
                                        switch (v59 - v4) {
                                        case 128:
                                            if true {
                                                v66 := GenS10{F0: v59, F1: 32768, F2: ""}
                                                for fi67 := 0; fi67 < 13; fi67++ {
                                                    switch fi67 {
                                                    case 1:
                                                        var v68 uint16 = 32767
                                                        _ = v68
                                                    case 9223372036854775807:
                                                        var v69 int32 = 2147483647
                                                        _ = v69
                                                    default:
                                                        v70 := GenS11{F0: v59}
                                                        _ = v70
                                                    }
                                                    _ = fi67
                                                }
                                                if (!((1 <= 417) && ((v23 - (v23 % (((v23 % ((v2) | 1))) | 1))) < (v2 | v12)))) {
                                                    var v71 float64 = -2.25
                                                    if ((v5 % ((v1) | 1)) >= v5) {
                                                        v72 := GenS7{F0: 42, F1: -32768}
                                                        v5 = v1
                                                        var v73 bool = false
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    _ = v71
                                                } else {
                                                    var v74 uint64 = 4294967295
                                                    _ = v74
                                                }
                                                _ = v66
                                            } else {
                                                var v75 float64 = ((((0.1 * 0.1) / (1e10 / 0.1)) - -2.25) / (-2.25 / 1.0))
                                                fmt.Println(v5, v59)
                                                var v76 uint64 = 2
                                                _ = v75
                                                _ = v76
                                            }
                                        case -2147483648:
                                            var v77 uint8 = 6
                                            _ = v77
                                        }
                                        for fi78 := 0; fi78 < 15; fi78++ {
                                            v79 := GenS10{F0: 42, F1: (fi78 ^ -200), F2: "a"}
                                            var v80 uint64 = 42
                                            switch (-(fi78)) {
                                            case -508:
                                                var v81 uint32 = 42
                                                _ = v81
                                            case -796:
                                                v4 = (v4 ^ (^(v4)))
                                            }
                                            _ = fi78
                                            _ = v79
                                            _ = v80
                                        }
                                    } else {
                                        fmt.Println(v12)
                                        v4 += -506
                                        for fi82 := 0; fi82 < 20; fi82++ {
                                            v3[65535] = (((("" + "") + (("a" + ("\n\t" + ("a" + ("\n\t" + ((("\n\t" + "hello") + "") + ("" + "\n\t")))))) + "\n\t")) + ("hello" + ("foo bar" + ((("foo bar" + "") + ("\n\t" + "\n\t")) + "")))) + "")
                                            v83 := GenS10{F0: -2147483648, F1: 1, F2: (((("a" + "a") + "foo bar") + "") + "\n\t")}
                                            _ = fi82
                                            _ = v83
                                        }
                                    }
                                    _ = len(v3)
                                    switch "foo bar" {
                                    case "":
                                        delete(v3, 11)
                                    case "foo bar":
                                        v84 := GenS8{F0: (v2 & (v12 & 1)), F1: v1}
                                        _ = v84
                                    default:
                                        var v85 string = "foo bar"
                                        _ = v85
                                    }
                                }
                            default:
                                v86 := []int{v4}
                                _ = v86
                            }
                            v87 := GenS5{F0: false, F1: 232, F2: 65535}
                            _ = v87
                        }
                        v88 := []int{42, v59, v4}
                        _ = v59
                        _ = v88
                    } else {
                        var v89 int = int(v5)
                        var v90 int8 = -128
                        _ = v89
                        _ = v90
                    }
                    v91 := GenS3{F0: 354, F1: int8(v4)}
                    _ = v91
                }
            case 2147483647:
                var v92 int8 = v5
                _ = v92
            default:
                v93 := GenS12{F0: v23}
                _ = v93
            }
            fmt.Println(v2)
        }
    }
    if (!((^(v4)) >= v4)) {
        for fi94 := 0; fi94 < 2; fi94++ {
            switch (v4 - fi94) {
            case 138:
                var v95 uint8 = (v23 | v23)
                _ = v95
            default:
                fmt.Println(v12)
            }
            var v96 int32 = 42
            _ = fi94
            _ = v96
        }
        var v97 int64 = 480
        _ = v97
    }
    var v98 bool = (65535 > 4294967295)
    if true {
        for fi99 := 0; fi99 < 18; fi99++ {
            v100 := GenS10{F0: (v4 % (((fi99 & (v4 << ((65535) & 63)))) | 1)), F1: (v4 - fi99), F2: "\n\t"}
            _ = fi99
            _ = v100
        }
    } else {
        var v101 int16 = 32767
        _ = v101
    }
    v102 := GenS8{F0: (^(v12)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v102
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
