package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint32; F1 int32; F2 uint }
func genFn0() uint {
    var v6 uint8 = 222
    var v7 int32 = -1
    _ = v6
    _ = v7
    return 511
}
func genFn1(p0 int64, p1 int16, p2 bool) uint8 {
    v8 := []int{}
    var v9 int = 1
    var v10 string = "a"
    for fi11 := 0; fi11 < 5; fi11++ {
        if true {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int = (fi12 / ((v9) | 1))
                v14 := GenS1{F0: 2, F1: 1}
                var v15 int = -558
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        if (fi11 >= (^(fi11))) {
            p1 = p1
        } else {
            var v16 int8 = 100
            v17 := GenS2{F0: p0}
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            v10 = "foo bar"
            var v19 int32 = 42
            switch 127 {
            case -662:
                switch ("" + v10) {
                case "hello":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    if (!(!(1e10 < ((((0.1 * 0.1) * 1.5) + (-2.25 * 0.0)) / 1.5)))) {
                        v8 = append(v8, fi11)
                        fmt.Println(fi11, fi18)
                    } else {
                        var v20 int8 = 127
                        switch (-(fi18)) {
                        case 810:
                            if p2 {
                                fmt.Println(fi18, v20)
                                switch v10 {
                                case "\n\t":
                                    switch v9 {
                                    case -317:
                                        v21 := GenS2{F0: -285}
                                        _ = v21
                                    }
                                }
                            } else {
                                v22 := GenS3{F0: (fi18 % ((42) | 1))}
                                p0 -= -9223372036854775808
                                if len(v8) > 0 {
                                    v8[0] = 42
                                }
                                _ = v22
                            }
                        }
                        _ = v20
                    }
                }
            }
            _ = fi18
            _ = v19
        }
        _ = fi11
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 bool = p2
        var v25 int8 = -128
        _ = fi23
        _ = v24
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]bool)
    var v3 uint = 2
    var v4 int16 = v1
    var v5 int64 = -9223372036854775808
    if false {
        delete(v2, 979)
        delete(v2, 4294967295)
        if (!true) {
            v26 := GenS1{F0: 1, F1: 4294967295}
            v27 := GenS4{F0: 2, F1: (v3 + 0)}
            _ = v26
            _ = v27
        }
    } else {
        v3 += v3
        if true {
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 int8 = 0
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 int64 = -677
                    var v32 string = ((("hello" + ((((("\n\t" + ("hello" + (("\n\t" + ("a" + ("" + "hello"))) + ""))) + "") + "hello") + (((("a" + "foo bar") + "a") + "hello") + "a")) + "a")) + "a") + "foo bar")
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
            for fi33 := 0; fi33 < 13; fi33++ {
                if (-128 <= 42) {
                    fmt.Println(v1, v5)
                    v5 = v5
                    fmt.Println(v4, v4, v1)
                }
                _ = fi33
            }
            var v34 int = 4294967295
            _ = v34
        } else {
            switch "a" {
            case "hello":
                v5 += v5
            case "a":
                v35 := []int{-2147483648, 992, 42}
                _ = v35
            case "\n\t":
                switch int(v5) {
                case 42:
                    fmt.Println(v3, v1)
                case 4294967295:
                    switch 0 {
                    case -1:
                        v3 = 65535
                    case 431:
                        if true {
                            fmt.Println(v1, v3)
                            if (!((!false) || true)) {
                                v5 = v5
                                v36 := v2[1]
                                _ = v36
                            } else {
                                if ((42 <= -1) || false) {
                                    fmt.Println(v1)
                                }
                            }
                            var v37 int64 = v5
                            _ = v37
                        } else {
                            var v38 string = ("a" + "foo bar")
                            for fi39 := 0; fi39 < 9; fi39++ {
                                var v40 int8 = -95
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    var v42 uint16 = 65535
                                    _ = fi41
                                    _ = v42
                                }
                                _ = fi39
                                _ = v40
                            }
                            var v43 string = v38
                            _ = v38
                            _ = v43
                        }
                    default:
                        v1 += v4
                    }
                case 75:
                    if ((1.5 - -2.25) > -2.25) {
                        v44 := []int{}
                        var v45 int16 = 0
                        _ = v44
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 16; fi46++ {
                            fmt.Println(v1)
                            _ = fi46
                        }
                    }
                }
            default:
                var v47 uint = 1
                _ = v47
            }
        }
        v48 := GenS5{F0: false, F1: 4294967295}
        _ = v48
    }
    fmt.Println(v5, v4)
    for fi49 := 0; fi49 < 6; fi49++ {
        v4 = (^(v1))
        for fi50 := 0; fi50 < 11; fi50++ {
            v51 := GenS6{F0: 927, F1: -355, F2: (v3 | genFn0())}
            fmt.Println(v4)
            v3 = v3
            _ = fi50
            _ = v51
        }
        for fi52 := 0; fi52 < 7; fi52++ {
            fmt.Println(v5, fi49, v5)
            if (!(127 >= 128)) {
                var v53 uint16 = 1
                v4 = (v1 << ((465) & 63))
                switch (fi49 / ((int(v5)) | 1)) {
                case -871:
                    if true {
                        var v54 uint32 = 2147483647
                        var v55 int16 = (v1 >> ((genFn0()) & 63))
                        _ = v54
                        _ = v55
                    }
                default:
                    if ((18446744073709551615 <= 0) && (("hello" + "hello") >= "")) {
                        v56 := v2[2147483648]
                        var v57 int8 = 0
                        _ = v56
                        _ = v57
                    }
                    var v58 int8 = 127
                    if true {
                        fmt.Println(v4)
                    }
                    v53 -= 65535
                    delete(v2, 1)
                    if (true || (249 > 408)) {
                        v58 -= 32
                        var v59 uint32 = 2147483648
                        v5 -= 0
                        _ = v59
                    }
                    _ = v58
                }
                _ = v53
            } else {
                if (!(true || false)) {
                    v4 -= v4
                }
            }
            _ = fi52
        }
        _ = fi49
    }
    for fi60 := 0; fi60 < 15; fi60++ {
        var v61 int = -1
        for fi62 := 0; fi62 < 5; fi62++ {
            v1 += 0
            switch -440 {
            case -9223372036854775808:
                fmt.Println(fi62)
            case 9223372036854775807:
                if false {
                    if true {
                        fmt.Println(v61)
                    }
                    var v63 int16 = v1
                    _ = v63
                } else {
                    v64 := GenS3{F0: (fi62 & (^(fi62)))}
                    delete(v2, 2147483648)
                    delete(v2, 65535)
                    _ = v64
                }
            default:
                delete(v2, 65535)
            }
            v65 := v2[42]
            if v65 {
                fmt.Println(v1)
            } else {
                v65 = ("foo bar" < "\n\t")
                switch fi62 {
                case 4294967295:
                    var v66 bool = (!(3.14159 == 1.5))
                    _ = v66
                case -275:
                    v67 := GenS6{F0: 2147483647, F1: 42, F2: (^(v3))}
                    _ = v67
                default:
                    var v68 bool = (!(255 >= 0))
                    _ = v68
                }
            }
            _ = fi62
            _ = v65
        }
        _ = fi60
        _ = v61
    }
    if (!((854 > 526) && (((((1 < 727) && false) && false) && (!(v1 <= (v1 | v1)))) && false))) {
        v4 += 32767
    } else {
        var v69 float64 = -2.25
        v1 = (v4 & v4)
        v3 += (^(v3))
        _ = v69
    }
    switch -651 {
    case 905:
        v3 = genFn0()
    case 4294967295:
        var v70 int16 = v1
        _ = v70
    case -2147483648:
        switch "foo bar" {
        case "hello":
            var v71 uint32 = 1
            _ = v71
        case "foo bar":
            fmt.Println(v5, v3)
        case "\n\t":
            v5 = 0
        default:
            for fi72 := 0; fi72 < 8; fi72++ {
                if (82 > genFn1((v5 | (v5 ^ (-(v5)))), (-(v1)), (1 >= 42))) {
                    v73 := v2[360]
                    for fi74 := 0; fi74 < 8; fi74++ {
                        v3 = 2
                        v2[1] = true
                        v75 := GenS4{F0: 2147483647, F1: 42}
                        _ = fi74
                        _ = v75
                    }
                    delete(v2, uint32(v3))
                    _ = v73
                } else {
                    var v76 uint8 = 42
                    _ = v76
                }
                _ = fi72
            }
        }
    default:
        v4 = v4
    }
    v77 := v2[651]
    v78 := GenS1{F0: 0, F1: 751}
    var v79 string = ("\n\t" + "foo bar")
    fmt.Println(v79, v3)
    v5 -= (v5 * (v5 << ((159) & 63)))
    v5 = v5
    for fi80 := 0; fi80 < 7; fi80++ {
        var v81 uint32 = 918
        _ = fi80
        _ = v81
    }
    switch "a" {
    case "a":
        for fi82 := 0; fi82 < 18; fi82++ {
            if v77 {
                var v83 uint32 = 1
                _ = v83
            }
            v79 = v79
            if false {
                if false {
                    var v84 int8 = 1
                    _ = v84
                }
            }
            _ = fi82
        }
    case "":
        for fi85 := 0; fi85 < 12; fi85++ {
            var v86 int = fi85
            _ = fi85
            _ = v86
        }
    default:
        v5 = v5
    }
    v5 = (-(v5))
    var v87 uint = 65535
    switch 32768 {
    case -9223372036854775808:
        v4 = v1
    case -340:
        for fi88 := 0; fi88 < 20; fi88++ {
            if v77 {
                delete(v2, 231)
            } else {
                if v77 {
                    if (v77 || false) {
                        v89 := []int{300, fi88, (fi88 ^ -350)}
                        fmt.Println(v77, v5)
                        _ = v89
                    }
                    switch (v79 + "\n\t") {
                    case "\n\t":
                        var v90 uint16 = 32767
                        _ = v90
                    }
                }
            }
            _ = fi88
        }
    case 9223372036854775807:
        if (v77 && (4294967295 < 771)) {
            var v91 uint8 = 229
            switch "foo bar" {
            case "":
                var v92 int16 = 0
                _ = v92
            case "foo bar":
                var v93 int16 = v4
                _ = v93
            case "a":
                var v94 bool = (0 < 42)
                _ = v94
            }
            for fi95 := 0; fi95 < 13; fi95++ {
                switch v79 {
                case "a":
                    var v96 uint = (v3 - v3)
                    _ = v96
                }
                v97 := GenS5{F0: (v79 > ((v79 + v79) + "hello")), F1: fi95}
                v98 := []int{-997, (fi95 | -2147483648), fi95, 128}
                _ = fi95
                _ = v97
                _ = v98
            }
            _ = v91
        } else {
            for fi99 := 0; fi99 < 11; fi99++ {
                v79 = "a"
                var v100 int32 = 0
                _ = fi99
                _ = v100
            }
        }
    default:
        v1 += (v4 & 121)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v78
    _ = v79
    _ = v87
}
