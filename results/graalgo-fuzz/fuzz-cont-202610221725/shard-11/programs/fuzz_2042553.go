package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32; F2 uint }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int64; F1 uint32; F2 uint }
type GenS8 struct { F0 int64; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint; F1 int64; F2 int }
type GenS14 struct { F0 int16; F1 int16 }
func genFn0(p0 uint32, p1 uint32) int32 {
    if (-617 != -976) {
        var v6 int = -9223372036854775808
        if true {
            var v7 int64 = 0
            if (v6 != (v6 + v6)) {
                v8 := []int{v6}
                v9 := GenS1{F0: v6}
                v8 = append(v8, (v6 - (v6 % (((v6 % ((v6) | 1))) | 1))))
                _ = v8
                _ = v9
            }
            fmt.Println(p0, v6, v7)
            _ = v7
        } else {
            fmt.Println(p0, v6)
            switch "hello" {
            case "\n\t":
                v6 = (^(v6))
            case "":
                v6 = 473
            }
        }
        _ = v6
    } else {
        v10 := []int{-1, 127, 2147483647, 1}
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{(^(fi11))}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1(p0 int8, p1 uint8) string {
    switch ("hello" + "foo bar") {
    case "\n\t":
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{(fi13 * (-(fi13))), fi13}
            v15 := GenS1{F0: fi13}
            _ = fi13
            _ = v14
            _ = v15
        }
    case "":
        fmt.Println(p0, p0)
    default:
        if (1 >= 32767) {
            v16 := GenS1{F0: int(p1)}
            _ = v16
        } else {
            v17 := []int{42, 1, 4294967295}
            _ = v17
        }
    }
    var v18 uint64 = 948
    switch ("a" + (("a" + ("a" + "\n\t")) + ("a" + ""))) {
    case "hello":
        v19 := []int{-2147483648, 368}
        _ = v19
    case "a":
        v20 := GenS2{F0: 42, F1: 511, F2: 2}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v18
    return ("hello" + ("\n\t" + ("\n\t" + "hello")))
}

func main() {
    var v1 int32 = 42
    var v2 uint16 = 32767
    v3 := make(map[uint32]int)
    v4 := make(map[string]uint32)
    v5 := []int{-2147483648, -1, 804, 9223372036854775807}
    _ = len(v4)
    var v21 int = 0
    v1 = v1
    if (true && (!false)) {
        fmt.Println(v1)
    } else {
        if false {
            switch ("a" + "a") {
            case "":
                for fi22 := 0; fi22 < 17; fi22++ {
                    v23 := GenS3{F0: ((1.5 + 3.14159) <= 1e10), F1: -1, F2: 2147483647}
                    if (1 < 1) {
                        if (127 < 42) {
                            var v24 bool = (!false)
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v26 := GenS4{F0: 2}
                                var v27 int = fi22
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                        if (1 < 37) {
                            v28 := GenS5{F0: (("foo bar" + "a") + "a"), F1: true, F2: 0}
                            _ = v28
                        }
                    } else {
                        fmt.Println(v1)
                    }
                    var v29 bool = true
                    _ = fi22
                    _ = v23
                    _ = v29
                }
            case "foo bar":
                switch 0 {
                case -39:
                    fmt.Println(v2)
                default:
                    if true {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            delete(v4, "hello")
                            _ = fi30
                        }
                        v21 = v21
                        var v31 int = v21
                        _ = v31
                    }
                }
            }
            v2 -= (v2 % ((v2) | 1))
            if (!(genFn0(65535, 79) < (v1 & (v1 << ((42) & 63))))) {
                if false {
                    v32 := GenS6{F0: 0}
                    if (false && (true || false)) {
                        v4["\n\t"] = uint32(v2)
                        switch 4294967295 {
                        case 32768:
                            fmt.Println(v2, v1)
                        }
                    } else {
                        v33 := v3[2147483647]
                        var v34 int = (v33 % ((v33) | 1))
                        for fi35 := 0; fi35 < 17; fi35++ {
                            var v36 uint8 = 128
                            _ = fi35
                            _ = v36
                        }
                        _ = v33
                        _ = v34
                    }
                    _ = v32
                }
                for fi37 := 0; fi37 < 1; fi37++ {
                    if len(v5) > 0 {
                        v5[0] = v21
                    }
                    if (!false) {
                        var v38 float32 = ((0.1 / 1.0) + 0.1)
                        if true {
                            fmt.Println(v38, v21, v2)
                            var v39 uint16 = (v2 & v2)
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v39 += 255
                                v39 -= 32767
                                _ = fi40
                            }
                            _ = v39
                        } else {
                            v21 = 32767
                            var v41 int8 = 100
                            if (false || false) {
                                v5 = append(v5, fi37)
                            } else {
                                for fi42 := 0; fi42 < 17; fi42++ {
                                    if true {
                                        switch (("foo bar" + "") + (("foo bar" + "foo bar") + ("hello" + "a"))) {
                                        case "\n\t":
                                            fmt.Println(fi37)
                                        case "foo bar":
                                            switch 4294967295 {
                                            case -2147483648:
                                                v2 = 42
                                                v43 := GenS3{F0: (42 != 32767), F1: 42, F2: 2}
                                                _ = v43
                                            case 128:
                                                switch fi37 {
                                                case 4294967295:
                                                    v44 := GenS7{F0: 0, F1: 2147483647, F2: 2}
                                                    _ = v44
                                                default:
                                                    var v45 float32 = v38
                                                    _ = v45
                                                }
                                            case 600:
                                                var v46 uint16 = v2
                                                _ = v46
                                            default:
                                                fmt.Println(v38, v41)
                                                switch (fi42 & (fi37 ^ (-(v21)))) {
                                                case 279:
                                                    var v47 float32 = v38
                                                    _ = v47
                                                case 128:
                                                    if (!(!(3.14159 < v38))) {
                                                        v48 := v3[2147483647]
                                                        v1 = (v1 | genFn0(65535, 737))
                                                        v49 := GenS8{F0: 42, F1: 181, F2: v2}
                                                        fmt.Println(v1)
                                                        v1 += genFn0(2147483648, 2147483648)
                                                        v50 := GenS9{F0: -32768}
                                                        var v51 float32 = ((((v38 * float32(1e10)) * v38) - v38) - v38)
                                                        _ = v48
                                                        _ = v49
                                                        _ = v50
                                                        _ = v51
                                                    } else {
                                                        v52 := v4["a"]
                                                        v53 := GenS4{F0: (v52 << ((0) & 63))}
                                                        var v54 int8 = 127
                                                        switch 9223372036854775807 {
                                                        case 32768:
                                                            v38 += float32(0.0)
                                                        case 2147483647:
                                                            v4[""] = (v52 % ((2147483648) | 1))
                                                        case 456:
                                                            var v55 int64 = 9223372036854775807
                                                            _ = v55
                                                        default:
                                                            v1 += (v1 + v1)
                                                            if false {
                                                                fmt.Println(v41, fi37, v38)
                                                                v54 = v54
                                                                if false {
                                                                    fmt.Println(v38, fi42, v52)
                                                                    var v56 uint = 2
                                                                    v57 := GenS7{F0: 0, F1: v52, F2: v56}
                                                                    _ = v56
                                                                    _ = v57
                                                                } else {
                                                                    var v58 string = ""
                                                                    var v59 uint = 725
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            } else {
                                                                v60 := GenS10{F0: 423}
                                                                if (2 > 128) {
                                                                    v1 -= 2147483647
                                                                }
                                                                fmt.Println(v2, v1)
                                                                _ = v60
                                                            }
                                                        }
                                                        fmt.Println(v41, v21)
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                }
                                            }
                                        case "a":
                                            v2 = (v2 * (v2 & (v2 << ((2) & 63))))
                                        }
                                    } else {
                                        var v61 uint32 = 570
                                        _ = v61
                                    }
                                    _ = fi42
                                }
                                var v62 int32 = v1
                                _ = v62
                            }
                            _ = v41
                        }
                        _ = v38
                    }
                    _ = fi37
                }
            } else {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v21, v21)
                case "hello":
                    var v63 float64 = float64(0.1)
                    _ = v63
                case "":
                    var v64 int8 = -128
                    _ = v64
                default:
                    for fi65 := 0; fi65 < 4; fi65++ {
                        var v66 uint = 255
                        var v67 uint = v66
                        v2 = (v2 >> ((v66) & 63))
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                }
                for fi68 := 0; fi68 < 9; fi68++ {
                    var v69 uint32 = 2147483647
                    switch (-(v21)) {
                    case 1:
                        switch "foo bar" {
                        case "":
                            var v70 uint16 = v2
                            _ = v70
                        case "\n\t":
                            v71 := GenS11{F0: 1}
                            _ = v71
                        case "a":
                            var v72 uint8 = 42
                            _ = v72
                        default:
                            v73 := v4["foo bar"]
                            _ = v73
                        }
                    case 523:
                        v74 := GenS12{F0: 0, F1: 65535, F2: 18446744073709551615}
                        _ = v74
                    }
                    _ = fi68
                    _ = v69
                }
                v21 = 9223372036854775807
            }
        }
        v2 = (v2 - (v2 & (v2 / ((v2) | 1))))
    }
    v75 := GenS7{F0: 974, F1: 2, F2: 705}
    v76 := GenS13{F0: 1, F1: 0, F2: v21}
    if false {
        var v77 string = "hello"
        _ = v77
    }
    if false {
        var v78 float32 = (0.1 / (0.1 - 1e10))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v78
    } else {
        switch (v21 % ((v21) | 1)) {
        case -55:
            v79 := GenS11{F0: 42}
            _ = v79
        default:
            for fi80 := 0; fi80 < 13; fi80++ {
                fmt.Println(v2)
                v81 := GenS14{F0: 1, F1: -1}
                _ = fi80
                _ = v81
            }
        }
        var v82 float32 = float32(-2.25)
        var v83 float32 = (((3.14159 + -2.25) - 0.1) - (v82 / (v82 * -2.25)))
        _ = v82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
}
