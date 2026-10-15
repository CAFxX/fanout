package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 int32; F2 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int16; F1 uint; F2 int32 }
type GenS13 struct { F0 uint64 }
func genFn0() uint16 {
    switch 127 {
    case 32767:
        switch -980 {
        case 32767:
            v4 := GenS1{F0: 560, F1: 2}
            _ = v4
        case -2147483648:
            for fi5 := 0; fi5 < 14; fi5++ {
                fmt.Println(fi5)
                v6 := GenS2{F0: 42}
                _ = fi5
                _ = v6
            }
        case 993:
            switch "a" {
            case "a":
                for fi7 := 0; fi7 < 2; fi7++ {
                    var v8 uint32 = 0
                    var v9 uint = 255
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = "hello"
        for fi12 := 0; fi12 < 15; fi12++ {
            v13 := GenS1{F0: 4294967295, F1: 2}
            _ = fi12
            _ = v13
        }
        fmt.Println(fi10, v11)
        _ = fi10
        _ = v11
    }
    v14 := GenS1{F0: 594, F1: 2}
    if (!(0.0 < (-2.25 + (((-2.25 / -2.25) + 0.1) + ((-2.25 - -2.25) - 3.14159))))) {
        var v15 uint8 = 42
        if false {
            var v16 uint16 = 0
            _ = v16
        } else {
            if (-103 <= 340) {
                var v17 bool = true
                _ = v17
            }
        }
        v18 := []int{}
        _ = v15
        _ = v18
    }
    _ = v14
    return 42
}
func genFn1() int {
    var v19 bool = ((!false) || true)
    v20 := GenS2{F0: 0}
    v21 := []int{4294967295}
    _ = v19
    _ = v20
    _ = v21
    return -2147483648
}

func main() {
    var v1 uint32 = 1
    var v2 int64 = -9223372036854775808
    var v3 int = 1
    v22 := []int{}
    v23 := GenS3{F0: 715}
    v24 := GenS3{F0: (v1 ^ (v1 + v1))}
    switch (v3 | (v3 + (^(v3)))) {
    case -9223372036854775808:
        v3 = v3
    case 32768:
        v1 = v1
    case 127:
        v25 := GenS4{F0: -125}
        _ = v25
    }
    _ = len(v22)
    v1 += v1
    var v26 uint64 = 42
    v2 -= v2
    switch ("" + "a") {
    case "\n\t":
        fmt.Println(v1)
    }
    if (false || (-420 != int16(v2))) {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 float64 = 1.5
            _ = fi27
            _ = v28
        }
        var v29 bool = (!false)
        _ = v29
    } else {
        var v30 float32 = (-2.25 / ((((-2.25 / (((((float32(0.0) + 0.0) + (3.14159 + 3.14159)) + (3.14159 - float32((-2.25 / 1.0)))) - 1e10) / 1.5)) + float32((-2.25 - -2.25))) + float32(((-2.25 + 1.5) / -2.25))) - 3.14159))
        var v31 uint16 = 762
        if ("hello" == "\n\t") {
            switch "hello" {
            case "\n\t":
                switch 2147483647 {
                case 32767:
                    var v32 string = "hello"
                    _ = v32
                case 182:
                    v1 -= (v1 * v1)
                default:
                    var v33 int = 42
                    _ = v33
                }
            case "foo bar":
                if true {
                    v1 = (^(v1))
                } else {
                    if false {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v35 := GenS5{F0: 329}
                            for fi36 := 0; fi36 < 18; fi36++ {
                                fmt.Println(v26, fi36)
                                var v37 string = ("\n\t" + "a")
                                v22 = append(v22, fi36)
                                _ = fi36
                                _ = v37
                            }
                            v1 += (v1 * (v1 | (v1 ^ (v1 * (^(v1))))))
                            _ = fi34
                            _ = v35
                        }
                        var v38 int16 = 29
                        var v39 uint16 = (v31 | v31)
                        _ = v38
                        _ = v39
                    }
                    fmt.Println(v3, v26)
                }
            }
            var v40 int32 = 42
            _ = v40
        }
        _ = v30
        _ = v31
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := GenS3{F0: uint32(v26)}
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v43 int64 = 218
        _ = fi41
        _ = v42
        _ = v43
    }
    v1 = 565
    v44 := GenS5{F0: -1}
    if (true && false) {
        v3 -= v3
    }
    var v45 string = ("a" + "")
    if true {
        var v46 int16 = 0
        if true {
            if (-2.25 != (0.1 / 3.14159)) {
                for fi47 := 0; fi47 < 11; fi47++ {
                    if true {
                        if true {
                            v48 := GenS6{F0: (false || false), F1: 1, F2: 127}
                            if false {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            if false {
                                v49 := GenS7{F0: 160}
                                for fi50 := 0; fi50 < 8; fi50++ {
                                    switch v45 {
                                    case "a":
                                        var v51 uint32 = v1
                                        if false {
                                            v51 = v1
                                            _ = len(v22)
                                        } else {
                                            var v52 uint8 = 128
                                            v53 := GenS8{F0: ((v46 & 32767) == 0)}
                                            v2 = v2
                                            var v54 uint32 = (v51 >> ((4294967295) & 63))
                                            v52 -= v52
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        if len(v22) > 0 {
                                            v22[0] = int(v26)
                                        }
                                        fmt.Println(v26, v3)
                                        _ = v51
                                    default:
                                        v55 := GenS9{F0: v45}
                                        _ = v55
                                    }
                                    v56 := GenS8{F0: (0.1 != (((0.0 + -2.25) - (-2.25 / 0.1)) * (float64(((-2.25 * (0.0 - 1.5)) + 0.1)) * 0.0)))}
                                    _ = fi50
                                    _ = v56
                                }
                                v45 = v45
                                _ = v49
                            } else {
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    var v58 float32 = 0.0
                                    var v59 string = v45
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                                v60 := GenS10{F0: (v26 | (v26 >> ((255) & 63))), F1: 0, F2: (-(v3))}
                                v61 := GenS1{F0: (v26 | 846), F1: 32767}
                                _ = v60
                                _ = v61
                            }
                            _ = v48
                        } else {
                            v3 += genFn1()
                            if (("" + v45) == ((("" + v45) + v45) + v45)) {
                                if len(v22) > 0 {
                                    v22[0] = (v3 << ((18446744073709551615) & 63))
                                }
                                v46 = (v46 / ((v46) | 1))
                                var v62 int = v3
                                _ = v62
                            }
                        }
                        v46 -= v46
                    } else {
                        var v63 float64 = 0.0
                        _ = len(v22)
                        _ = v63
                    }
                    var v64 uint64 = 1
                    fmt.Println(v45)
                    _ = fi47
                    _ = v64
                }
            } else {
                var v65 string = v45
                _ = v65
            }
            v3 += v3
        }
        v66 := GenS1{F0: v26, F1: 497}
        _ = v46
        _ = v66
    }
    if (3.14159 == -2.25) {
        var v67 int16 = 0
        if true {
            for fi68 := 0; fi68 < 14; fi68++ {
                fmt.Println(v67)
                if (127 <= 128) {
                    if (!(33 > 42)) {
                        var v69 int8 = 100
                        _ = v69
                    } else {
                        var v70 bool = (!true)
                        for fi71 := 0; fi71 < 3; fi71++ {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            _ = fi71
                        }
                        _ = v70
                    }
                    v72 := GenS11{F0: 18446744073709551615}
                    _ = v72
                }
                if ((-1 != 1) && (v2 == (v2 / (((v2 | 152)) | 1)))) {
                    if true {
                        for fi73 := 0; fi73 < 6; fi73++ {
                            v26 = v26
                            v22 = append(v22, (v3 ^ (fi68 - fi73)))
                            _ = fi73
                        }
                        v1 += (v1 / ((v1) | 1))
                        var v74 uint64 = 876
                        _ = v74
                    } else {
                        _ = len(v22)
                    }
                    var v75 float64 = 1.5
                    switch -2147483648 {
                    case 128:
                        v76 := GenS11{F0: v26}
                        _ = v76
                    }
                    _ = v75
                } else {
                    v2 = (v2 << ((888) & 63))
                    var v77 int64 = (v2 % ((v2) | 1))
                    v77 = (v77 * v2)
                    _ = v77
                }
                _ = fi68
            }
            v3 += v3
        } else {
            var v78 bool = (!true)
            _ = v78
        }
        _ = v67
    }
    if ((v3 ^ (^(v3))) < (v3 << ((0) & 63))) {
        if true {
            for fi79 := 0; fi79 < 6; fi79++ {
                for fi80 := 0; fi80 < 8; fi80++ {
                    if true {
                        var v81 float64 = (-2.25 - (0.1 - (((0.1 * 1.5) * 1e10) * ((0.1 + (1e10 * -2.25)) * 0.1))))
                        var v82 string = v45
                        if false {
                            fmt.Println(v81)
                        }
                        _ = v81
                        _ = v82
                    } else {
                        var v83 uint64 = 2
                        v1 = (v1 | 676)
                        if (!false) {
                            v3 -= (-(fi80))
                            switch v3 {
                            case 42:
                                v84 := GenS12{F0: -837, F1: 255, F2: 1}
                                _ = v84
                            case -9223372036854775808:
                                v45 = v45
                            case -2147483648:
                                var v85 bool = false
                                _ = len(v22)
                                if v85 {
                                    v86 := GenS11{F0: v26}
                                    _ = v86
                                }
                                _ = v85
                            }
                            v87 := GenS13{F0: v26}
                            _ = v87
                        }
                        _ = v83
                    }
                    _ = fi80
                }
                var v88 int64 = (v2 & v2)
                v89 := GenS12{F0: 0, F1: 255, F2: 0}
                _ = fi79
                _ = v88
                _ = v89
            }
            v26 += v26
            var v90 int = genFn1()
            _ = v90
        } else {
            v2 += (v2 << ((4294967295) & 63))
            fmt.Println(v45)
            v2 += 375
        }
        var v91 float64 = 1.5
        v91 -= 0.1
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v44
    _ = v45
}
