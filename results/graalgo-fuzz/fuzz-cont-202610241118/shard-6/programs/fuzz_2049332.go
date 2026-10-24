package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int32; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint, p1 int16) uint {
    p0 = p0
    v4 := GenS1{F0: 1, F1: 4294967295, F2: 0}
    _ = p0
    _ = p1
    _ = v4
    return (p0 + 4294967295)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 255
    var v3 uint8 = (^(v2))
    v5 := GenS1{F0: 65535, F1: 4294967295, F2: 896}
    v2 -= (v2 >> ((65535) & 63))
    var v6 int64 = -1
    v3 = 81
    for fi7 := 0; fi7 < 1; fi7++ {
        if ((v6 | v6) >= (v6 - v6)) {
            v8 := GenS1{F0: 0, F1: 2147483647, F2: 576}
            _ = v8
        }
        if true {
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := []int{-986, fi7, (fi9 / (((-(fi7))) | 1)), (fi7 * -2147483648)}
                if (127 == 100) {
                    var v11 float32 = -2.25
                    v3 = (^(v2))
                    var v12 float64 = (-2.25 * 1.5)
                    _ = v11
                    _ = v12
                }
                var v13 int = fi9
                _ = fi9
                _ = v10
                _ = v13
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                v6 += (v6 << ((255) & 63))
                var v15 int16 = 1000
                _ = fi14
                _ = v15
            }
        }
        _ = fi7
    }
    if false {
        v2 = (v2 & (v3 >> ((2) & 63)))
        v3 = (^(v3))
        if true {
            if true {
                switch "a" {
                case "hello":
                    v1["\n\t"] = uint32(v3)
                case "foo bar":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 int32 = 42
                        _ = fi16
                        _ = v17
                    }
                case "":
                    switch "a" {
                    case "a":
                        if (-32768 != 0) {
                            for fi18 := 0; fi18 < 4; fi18++ {
                                v19 := []int{-1, fi18, 2147483647}
                                v2 = 1
                                _ = fi18
                                _ = v19
                            }
                        } else {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                if true {
                                    fmt.Println(v2, v2, v6)
                                    v21 := GenS2{F0: (fi20 * fi20)}
                                    _ = v21
                                } else {
                                    var v22 int64 = 0
                                    fmt.Println(v3, v3, v6)
                                    delete(v1, "foo bar")
                                    _ = v22
                                }
                                _ = fi20
                            }
                            v23 := []int{2147483647, 127, 0}
                            _ = v23
                        }
                    }
                }
                if (!((((((-2.25 / -2.25) / -2.25) - 0.1) / -2.25) - (0.1 / 3.14159)) >= 0.0)) {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        if false {
                            var v25 uint8 = v3
                            if false {
                                var v26 uint8 = 5
                                _ = v26
                            }
                            _ = v25
                        }
                        _ = fi24
                    }
                    var v27 bool = (284 == 65535)
                    var v28 string = "hello"
                    _ = v27
                    _ = v28
                }
            }
            if ((true || (!((("a" + "\n\t") + "hello") < "foo bar"))) && (v2 <= 1)) {
                for fi29 := 0; fi29 < 4; fi29++ {
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v6 = (v6 ^ v6)
                        _ = fi30
                    }
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v32 := GenS1{F0: 650, F1: 437, F2: -168}
                        if false {
                            if ((1 > 627) || false) {
                                if ((v6 | v6) > -909) {
                                    if (2147483647 < 0) {
                                        v33 := v1["\n\t"]
                                        var v34 string = "hello"
                                        var v35 int32 = 0
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    } else {
                                        if false {
                                            v3 -= v2
                                        }
                                        if true {
                                            v36 := GenS1{F0: 269, F1: 226, F2: 2147483647}
                                            fmt.Println(fi29, v3)
                                            _ = v36
                                        }
                                        var v37 uint64 = 0
                                        _ = v37
                                    }
                                    fmt.Println(v3, v2, v2)
                                    v38 := GenS2{F0: fi31}
                                    v39 := GenS2{F0: (^(fi31))}
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 uint64 = 18446744073709551615
                                    var v41 uint16 = 175
                                    var v42 float32 = 0.1
                                    var v43 float64 = 0.1
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            } else {
                                v44 := GenS1{F0: 515, F1: 1, F2: 1}
                                var v45 uint64 = 555
                                fmt.Println(v45)
                                _ = v44
                                _ = v45
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                    delete(v1, "hello")
                    _ = fi29
                }
                var v46 bool = (!true)
                for fi47 := 0; fi47 < 11; fi47++ {
                    v48 := GenS3{F0: 649, F1: v2, F2: 4294967295}
                    for fi49 := 0; fi49 < 1; fi49++ {
                        if false {
                            var v50 int32 = -1
                            var v51 int16 = -24
                            _ = v50
                            _ = v51
                        } else {
                            if (!v46) {
                                v1[("" + ("" + "hello"))] = 1
                            } else {
                                var v52 string = ("a" + "a")
                                var v53 float64 = (1e10 * (0.1 - 0.0))
                                v54 := GenS4{F0: 1, F1: 0}
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        }
                        var v55 int8 = -128
                        v56 := GenS5{F0: 981, F1: -1, F2: v2}
                        _ = fi49
                        _ = v55
                        _ = v56
                    }
                    fmt.Println(v2)
                    _ = fi47
                    _ = v48
                }
                _ = v46
            }
            fmt.Println(v3)
        }
    }
    for fi57 := 0; fi57 < 17; fi57++ {
        var v58 int16 = -1
        _ = fi57
        _ = v58
    }
    v6 = v6
    switch 1 {
    case 810:
        if true {
            v59 := []int{}
            switch ("foo bar" + (("\n\t" + "a") + "foo bar")) {
            case "hello":
                v3 += v2
            case "a":
                for fi60 := 0; fi60 < 5; fi60++ {
                    switch fi60 {
                    case 1:
                        v61 := GenS2{F0: (fi60 * (fi60 % ((fi60) | 1)))}
                        _ = v61
                    case 2147483647:
                        v62 := GenS1{F0: 2, F1: 1, F2: 1}
                        _ = v62
                    }
                    v2 += v2
                    _ = fi60
                }
            case "\n\t":
                if true {
                    v63 := GenS6{F0: 120}
                    _ = v63
                }
            default:
                for fi64 := 0; fi64 < 10; fi64++ {
                    var v65 int8 = 5
                    _ = fi64
                    _ = v65
                }
            }
            _ = v59
        }
    default:
        v2 -= (v2 << ((64) & 63))
    }
    v66 := GenS5{F0: 42, F1: (v6 >> ((2) & 63)), F2: v3}
    v67 := v1[((("" + "foo bar") + "a") + "a")]
    for fi68 := 0; fi68 < 18; fi68++ {
        for fi69 := 0; fi69 < 4; fi69++ {
            var v70 int8 = 1
            delete(v1, "foo bar")
            v70 -= (v70 * (^(v70)))
            _ = fi69
            _ = v70
        }
        var v71 int32 = 0
        _ = fi68
        _ = v71
    }
    if (-736 <= 128) {
        v67 += 0
    } else {
        fmt.Println(v67)
        v3 -= (v2 * 146)
    }
    if true {
        v2 = (^(v3))
        v67 -= v67
    }
    v2 = 1
    fmt.Println(v6, v67)
    switch 4294967295 {
    case 42:
        if true {
            if (-2.25 == 1.5) {
                v2 -= 2
                delete(v1, "a")
                v72 := GenS7{F0: 706, F1: -128}
                _ = v72
            } else {
                v67 = (v67 - (v67 + (^(v67))))
                if true {
                    switch ("a" + "hello") {
                    case "a":
                        for fi73 := 0; fi73 < 7; fi73++ {
                            v2 = v2
                            switch int(v67) {
                            case 90:
                                if (true && (4294967295 == 0)) {
                                    v74 := []int{fi73, fi73}
                                    for fi75 := 0; fi75 < 5; fi75++ {
                                        v3 += (v3 - (v3 - v2))
                                        if (false || true) {
                                            v3 = v3
                                        } else {
                                            var v76 int64 = (v6 % ((964) | 1))
                                            _ = v76
                                        }
                                        _ = fi75
                                    }
                                    _ = v74
                                }
                            }
                            fmt.Println(v6, v3)
                            _ = fi73
                        }
                    case "":
                        fmt.Println(v67, v6, v6)
                    }
                }
                var v77 int = 9223372036854775807
                _ = v77
            }
            v78 := []int{-9223372036854775808, -1, 2147483647, 210}
            _ = v78
        } else {
            v79 := GenS4{F0: 0, F1: -21}
            v3 -= v3
            var v80 int8 = 0
            _ = v79
            _ = v80
        }
    }
    var v81 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v81
}
