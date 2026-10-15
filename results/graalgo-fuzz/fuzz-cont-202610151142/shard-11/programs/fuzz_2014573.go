package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int32; F2 int8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 int; F2 uint16 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0() int16 {
    v7 := []int{42, 127, 9223372036854775807}
    var v8 int = 127
    var v9 int8 = 1
    fmt.Println(v8, v8)
    v9 += -57
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := []string{"foo bar", "hello"}
    var v2 uint8 = 1
    var v3 string = ("\n\t" + "a")
    var v4 bool = true
    var v5 int = 625
    v6 := make(map[string]string)
    v2 = 42
    for fi10 := 0; fi10 < 1; fi10++ {
        v3 = (("hello" + ((((v3 + (v3 + v3)) + v3) + v3) + v3)) + (v3 + v3))
        for fi11 := 0; fi11 < 16; fi11++ {
            switch fi11 {
            case 42:
                switch "foo bar" {
                case "foo bar":
                    v12 := GenS1{F0: 1, F1: v4, F2: -1}
                    _ = v12
                case "a":
                    v13 := GenS1{F0: 42, F1: (!v4), F2: 0}
                    _ = v13
                case "\n\t":
                    var v14 float32 = ((0.0 - 1e10) + 0.0)
                    _ = v14
                default:
                    if (v2 > v2) {
                        if v4 {
                            v15 := GenS1{F0: 716, F1: false, F2: 0}
                            v16 := v6["a"]
                            v5 = 9223372036854775807
                            _ = v15
                            _ = v16
                        } else {
                            if (-2147483648 <= 916) {
                                var v17 uint32 = 578
                                var v18 uint = 2
                                var v19 string = (v3 + "hello")
                                _ = len(v6)
                                var v20 float64 = ((1e10 + 0.0) - (float64((1e10 + 0.0)) + 3.14159))
                                v21 := GenS1{F0: 1, F1: v4, F2: int64(v2)}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            v22 := v6[(v3 + ((v3 + (v3 + "")) + (v3 + "a")))]
                            v5 += fi11
                            _ = v22
                        }
                        switch "\n\t" {
                        case "":
                            v23 := GenS2{F0: v4, F1: 0}
                            _ = v23
                        case "foo bar":
                            v2 = v2
                        case "\n\t":
                            switch v3 {
                            case "a":
                                var v24 int64 = 0
                                _ = v24
                            case "foo bar":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "\n\t":
                                var v25 uint32 = 2
                                _ = v25
                            }
                        default:
                            var v26 uint16 = 1
                            _ = v26
                        }
                    }
                }
            case 32767:
                switch (v3 + "a") {
                case "\n\t":
                    if (v4 || (-57 == -11)) {
                        fmt.Println(v5, v5, v5)
                        switch v3 {
                        case "a":
                            var v27 bool = (42 <= -1)
                            _ = v27
                        case "foo bar":
                            switch (v3 + (v3 + (v3 + v3))) {
                            case "foo bar":
                                if v4 {
                                    v28 := GenS3{F0: 42}
                                    v29 := v6[""]
                                    var v30 float32 = (1e10 - (1.5 / (1.5 + float32((0.1 - 3.14159)))))
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                            default:
                                v31 := GenS4{F0: 2, F1: 560, F2: 100}
                                _ = v31
                            }
                        default:
                            if true {
                                var v32 int32 = -113
                                _ = v32
                            } else {
                                switch v5 {
                                case 42:
                                    switch (v3 + (((v3 + ((((v3 + "hello") + v3) + v3) + ((((v3 + "foo bar") + v3) + "a") + v3))) + v3) + ((v3 + (v3 + "a")) + "\n\t"))) {
                                    case "foo bar":
                                        switch (v3 + (v3 + "a")) {
                                        case "foo bar":
                                            v33 := GenS3{F0: -906}
                                            _ = v33
                                        default:
                                            v34 := GenS3{F0: -1}
                                            _ = v34
                                        }
                                    case "hello":
                                        if false {
                                            fmt.Println(v2, v5)
                                        }
                                    case "":
                                        var v35 int64 = -3
                                        _ = v35
                                    }
                                default:
                                    var v36 int64 = -9223372036854775808
                                    v3 = (v3 + v3)
                                    fmt.Println(fi10)
                                    _ = v36
                                }
                                if (166 != 341) {
                                    switch v3 {
                                    case "a":
                                        v2 += (^(v2))
                                    case "hello":
                                        v4 = false
                                    case "":
                                        v4 = false
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if v4 {
                                        v2 = 213
                                        fmt.Println(fi11, v4, v4)
                                    } else {
                                        switch "" {
                                        case "\n\t":
                                            if v4 {
                                                var v37 uint16 = 65535
                                                v38 := GenS5{F0: -143, F1: true}
                                                _ = v37
                                                _ = v38
                                            }
                                        case "":
                                            v39 := GenS5{F0: 778, F1: (-45 == 42)}
                                            _ = v39
                                        default:
                                            v5 = 45
                                        }
                                        v3 = (v3 + (v3 + v3))
                                    }
                                    if v4 {
                                        v5 = 4294967295
                                    } else {
                                        var v40 uint8 = v2
                                        _ = v40
                                    }
                                    switch v3 {
                                    case "a":
                                        var v41 float32 = 0.1
                                        _ = v41
                                    case "foo bar":
                                        v42 := GenS6{F0: 1, F1: v2}
                                        _ = v42
                                    case "":
                                        v3 = v3
                                    }
                                    var v43 int64 = 9223372036854775807
                                    _ = v43
                                } else {
                                    v3 = v3
                                }
                            }
                        }
                        _ = len(v6)
                    } else {
                        var v44 int16 = -1
                        v4 = (-1 > -1)
                        _ = v44
                    }
                case "foo bar":
                    delete(v6, ("\n\t" + "a"))
                case "hello":
                    v4 = ((^(fi11)) == (fi10 | (v5 % (((fi11 ^ -239)) | 1))))
                default:
                    v3 = (v3 + "hello")
                }
            case 0:
                v45 := GenS3{F0: 2147483647}
                _ = v45
            default:
                switch 4294967295 {
                case -1:
                    var v46 int8 = -52
                    _ = v46
                case 9223372036854775807:
                    v47 := GenS5{F0: 2147483647, F1: (((0.0 / float64((-2.25 / (-2.25 / 1.0)))) <= (1.5 * 0.1)) && (!v4))}
                    _ = v47
                case -418:
                    if v4 {
                        v48 := GenS7{F0: 42}
                        fmt.Println(v2, fi10, v5)
                        var v49 int64 = -9223372036854775808
                        _ = v48
                        _ = v49
                    } else {
                        if (0.0 != 0.1) {
                            fmt.Println(fi11, v3, v2)
                        }
                    }
                default:
                    switch v3 {
                    case "a":
                        var v50 float64 = 1.5
                        _ = v50
                    case "":
                        fmt.Println(fi10, fi10, fi10)
                    case "foo bar":
                        switch v5 {
                        case 4294967295:
                            var v51 int8 = -128
                            _ = v51
                        }
                    }
                }
            }
            v52 := GenS4{F0: 18446744073709551615, F1: 42, F2: 100}
            _ = fi11
            _ = v52
        }
        if true {
            var v53 int64 = 362
            v2 -= (^(v2))
            _ = v53
        }
        _ = fi10
    }
    _ = len(v1)
    var v54 string = v3
    if v4 {
        var v55 bool = true
        _ = v55
    } else {
        switch (v5 + -456) {
        case -297:
            var v56 string = v54
            _ = v56
        default:
            if len(v1) > 0 {
                v1[0] = ("\n\t" + (((v3 + ("hello" + v3)) + (v3 + v54)) + ""))
            }
        }
        if (65535 <= 2) {
            for fi57 := 0; fi57 < 6; fi57++ {
                v2 += 144
                v58 := GenS8{F0: 1, F1: v5, F2: 0}
                var v59 uint32 = 2147483647
                _ = fi57
                _ = v58
                _ = v59
            }
            if (42 <= -719) {
                v60 := GenS4{F0: 1, F1: -2147483648, F2: 1}
                _ = v60
            }
        }
    }
    if (v4 || (v5 != v5)) {
        fmt.Println(v54, v5)
        var v61 int32 = 0
        _ = v61
    }
    v6[""] = "foo bar"
    var v62 uint32 = 65535
    for fi63 := 0; fi63 < 12; fi63++ {
        if false {
            v64 := GenS7{F0: v62}
            v5 -= fi63
            v54 = "hello"
            _ = v64
        } else {
            if len(v1) > 0 {
                v1[0] = v54
            }
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 int8 = int8(v2)
                _ = fi65
                _ = v66
            }
        }
        var v67 float64 = 0.1
        _ = fi63
        _ = v67
    }
    if ((!(991 != 2)) || v4) {
        var v68 bool = ((v3 > (v3 + "foo bar")) && (0 <= 18446744073709551615))
        var v69 uint32 = v62
        _ = v68
        _ = v69
    }
    v3 = (v3 + v54)
    var v70 int16 = -625
    var v71 int16 = (v70 << ((255) & 63))
    for fi72 := 0; fi72 < 15; fi72++ {
        for fi73 := 0; fi73 < 1; fi73++ {
            fmt.Println(v70, fi73, v70)
            v6[v54] = "\n\t"
            switch (v3 + v3) {
            case "a":
                var v74 uint16 = 2
                _ = v74
            case "":
                if v4 {
                    var v75 int = 32768
                    v76 := GenS9{F0: 18446744073709551615, F1: 2}
                    v77 := GenS10{F0: (v62 << ((699) & 63)), F1: 0, F2: 42}
                    _ = v75
                    _ = v76
                    _ = v77
                }
            default:
                var v78 int16 = (v70 + (v70 * genFn0()))
                _ = v78
            }
            _ = fi73
        }
        if (72 >= 382) {
            for fi79 := 0; fi79 < 6; fi79++ {
                if v4 {
                    var v80 uint = 4294967295
                    v2 = v2
                    _ = v80
                }
                v81 := GenS4{F0: 1, F1: 2147483647, F2: -1}
                _ = fi79
                _ = v81
            }
        } else {
            v3 = "\n\t"
            for fi82 := 0; fi82 < 12; fi82++ {
                if (0.0 != 1.5) {
                    v5 = v5
                    switch (-(fi72)) {
                    case -440:
                        v83 := GenS6{F0: 1, F1: v2}
                        _ = v83
                    }
                    var v84 int16 = genFn0()
                    _ = v84
                }
                fmt.Println(v54, v4, v71)
                fmt.Println(v54)
                _ = fi82
            }
        }
        var v85 int64 = 42
        _ = fi72
        _ = v85
    }
    v70 = (v70 & (-(v70)))
    var v86 string = ((("hello" + v54) + v3) + "foo bar")
    if v4 {
        var v87 int8 = -128
        _ = v87
    } else {
        v71 -= v71
        fmt.Println(v62, v4)
        var v88 bool = v4
        _ = v88
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v62
    _ = v70
    _ = v71
    _ = v86
}
