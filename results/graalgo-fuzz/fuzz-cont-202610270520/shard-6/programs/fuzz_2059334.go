package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint8; F2 string }
type GenS6 struct { F0 int32; F1 uint64; F2 bool }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 bool; F1 int8; F2 int32 }
type GenS9 struct { F0 int16; F1 uint16 }
func genFn0() int64 {
    v6 := GenS1{F0: -583}
    switch (((("\n\t" + ("hello" + ("hello" + (("foo bar" + "\n\t") + ((("foo bar" + "a") + "a") + "a"))))) + "foo bar") + "foo bar") + "foo bar") {
    case "foo bar":
        var v7 uint16 = 65535
        _ = v7
    case "":
        var v8 int32 = 999
        _ = v8
    case "\n\t":
        var v9 int = -279
        _ = v9
    default:
        v10 := GenS2{F0: (0 >= 0), F1: 2, F2: 4294967295}
        _ = v10
    }
    v11 := []int{480, 32768, -2147483648}
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint8 = 2
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: (fi14 - 32768)}
            v13 = (v13 + v13)
            v13 += (v13 % ((v13) | 1))
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    _ = v11
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 0.1
    var v3 uint16 = 255
    var v4 int64 = 9223372036854775807
    v5 := []int{1}
    v16 := GenS1{F0: 809}
    fmt.Println(v2, v4, v3)
    v4 += -9223372036854775808
    v2 = v2
    switch "foo bar" {
    case "a":
        v2 = 3.14159
    case "":
        var v17 int = -929
        _ = v17
    default:
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "a":
            var v18 float32 = 1.5
            _ = v18
        case "\n\t":
            var v19 float32 = -2.25
            _ = v19
        case "foo bar":
            v2 += v2
        default:
            fmt.Println(v3)
        }
    }
    v20 := GenS3{F0: 1}
    _ = len(v1)
    var v21 int16 = 42
    v21 = (-(v21))
    var v22 uint8 = 200
    var v23 int8 = -128
    v22 = 42
    if true {
        switch -1 {
        case -2147483648:
            if true {
                v23 = (v23 - int8(v21))
            } else {
                v24 := GenS3{F0: 0}
                var v25 uint64 = 42
                fmt.Println(v4)
                _ = v24
                _ = v25
            }
        case 956:
            fmt.Println(v3)
        default:
            var v26 int = -346
            _ = v26
        }
        switch 1 {
        case 32768:
            for fi27 := 0; fi27 < 17; fi27++ {
                v22 += 205
                _ = fi27
            }
        case -963:
            var v28 uint64 = 4294967295
            _ = v28
        case 0:
            v2 = v2
        default:
            if (4294967295 > 2) {
                fmt.Println(v23, v23)
                var v29 int8 = (v23 | v23)
                switch "a" {
                case "foo bar":
                    if false {
                        v30 := v1[42]
                        v22 += (v22 + (^(v22)))
                        v2 = ((0.0 * (1e10 * 1.5)) * 1.5)
                        _ = v30
                    } else {
                        if true {
                            if ((706 < 183) || false) {
                                for fi31 := 0; fi31 < 9; fi31++ {
                                    if (v2 <= float64(((0.0 * (-2.25 + (1.5 / 1e10))) * 1.5))) {
                                        v23 = 127
                                        var v32 uint32 = 1
                                        _ = v32
                                    } else {
                                        for fi33 := 0; fi33 < 19; fi33++ {
                                            if false {
                                                v5 = append(v5, (fi33 % ((fi31) | 1)))
                                                var v34 int64 = v4
                                                _ = v34
                                            }
                                            var v35 int64 = (-(v4))
                                            var v36 uint32 = 65535
                                            _ = fi33
                                            _ = v35
                                            _ = v36
                                        }
                                        v37 := GenS4{F0: 1}
                                        v38 := GenS3{F0: -1}
                                        _ = v37
                                        _ = v38
                                    }
                                    var v39 float32 = (float32(1.5) + 3.14159)
                                    _ = fi31
                                    _ = v39
                                }
                            } else {
                                v3 -= (v3 % ((v3) | 1))
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v40 int8 = v29
                                _ = v40
                            }
                            v41 := v1[uint32(v22)]
                            var v42 string = "foo bar"
                            _ = v41
                            _ = v42
                        }
                        fmt.Println(v4, v21, v4)
                        v29 = v23
                    }
                default:
                    v43 := GenS5{F0: (v4 - v4), F1: v22, F2: "foo bar"}
                    _ = v43
                }
                _ = v29
            } else {
                var v44 uint = 2
                var v45 uint = v44
                switch 1 {
                case 2147483647:
                    var v46 float32 = 0.0
                    _ = v46
                case 116:
                    if (-2.25 < v2) {
                        v23 = (v23 % ((-128) | 1))
                        var v47 int16 = (v21 - 42)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v47
                    }
                case 700:
                    if (!true) {
                        if false {
                            v48 := GenS1{F0: int(v4)}
                            var v49 uint32 = 2
                            _ = v48
                            _ = v49
                        } else {
                            v21 = v21
                            v45 = v44
                            fmt.Println(v21)
                        }
                    }
                }
                _ = v44
                _ = v45
            }
        }
        fmt.Println(v2, v23, v22)
    }
    fmt.Println(v4, v21)
    fmt.Println(v4, v21)
    var v50 uint = 2
    v4 = v4
    if false {
        if (!("" != "\n\t")) {
            v22 += v22
            var v51 uint = v50
            var v52 uint32 = 2147483647
            _ = v51
            _ = v52
        } else {
            for fi53 := 0; fi53 < 4; fi53++ {
                for fi54 := 0; fi54 < 14; fi54++ {
                    var v55 int8 = -35
                    _ = fi54
                    _ = v55
                }
                var v56 int32 = -129
                _ = fi53
                _ = v56
            }
        }
        v57 := GenS5{F0: v4, F1: v22, F2: "foo bar"}
        if true {
            if true {
                fmt.Println(v4, v23)
                for fi58 := 0; fi58 < 5; fi58++ {
                    for fi59 := 0; fi59 < 12; fi59++ {
                        v50 = v50
                        var v60 uint = (v50 >> ((v50) & 63))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 16; fi61++ {
                        switch fi58 {
                        case -115:
                            _ = len(v1)
                        case -155:
                            fmt.Println(v22)
                        case 4294967295:
                            v62 := GenS6{F0: 1, F1: 18446744073709551615, F2: (!false)}
                            _ = v62
                        }
                        fmt.Println(v2)
                        v63 := GenS5{F0: (v4 / ((-9223372036854775808) | 1)), F1: v22, F2: "foo bar"}
                        _ = fi61
                        _ = v63
                    }
                    fmt.Println(v23, fi58, v23)
                    _ = fi58
                }
            } else {
                v64 := GenS1{F0: 9223372036854775807}
                v22 = v22
                _ = v64
            }
            for fi65 := 0; fi65 < 13; fi65++ {
                v66 := GenS3{F0: 42}
                if true {
                    for fi67 := 0; fi67 < 13; fi67++ {
                        v50 += (^(v50))
                        var v68 int = (fi67 << ((255) & 63))
                        var v69 int32 = 0
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    v70 := GenS7{F0: "hello", F1: (v50 - v50)}
                    if false {
                        v2 += v2
                        _ = len(v1)
                        if (!(v2 != ((0.1 - (v2 + 0.1)) / 1.0))) {
                            v71 := GenS8{F0: false, F1: v23, F2: 492}
                            v21 = (v21 * v21)
                            switch (fi65 + 83) {
                            case 127:
                                v1[65535] = (v4 < -1)
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            _ = v71
                        } else {
                            v4 = (^(v4))
                            for fi72 := 0; fi72 < 2; fi72++ {
                                v21 = 0
                                var v73 float64 = (v2 / v2)
                                v74 := GenS3{F0: 1}
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            _ = len(v5)
                        }
                    } else {
                        v1[851] = (((true && ((("\n\t" + "hello") + ((("hello" + (("a" + "a") + ("\n\t" + ""))) + ((("\n\t" + ("\n\t" + "a")) + ("hello" + "foo bar")) + ("" + ""))) + ((("a" + "") + ("" + "")) + (((("" + "") + "\n\t") + ("foo bar" + "\n\t")) + "foo bar")))) < "a")) && (0.1 < -2.25)) && false)
                    }
                    _ = v70
                } else {
                    fmt.Println(v23)
                    v75 := GenS9{F0: 0, F1: v3}
                    v1[1] = false
                    _ = v75
                }
                var v76 int32 = -1
                _ = fi65
                _ = v66
                _ = v76
            }
        }
        _ = v57
    } else {
        var v77 uint16 = (v3 & v3)
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
