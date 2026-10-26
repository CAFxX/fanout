package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int8, p1 bool, p2 uint) uint8 {
    if (((p1 && false) || true) || p1) {
        v7 := []int{32767, -1, 4294967295, -249}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int = (fi8 - 491)
            _ = fi8
            _ = v9
        }
        if p1 {
            fmt.Println(p0)
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int8 = p0
                fmt.Println(p0, p0)
                switch "\n\t" {
                case "hello":
                    var v12 int8 = v11
                    _ = v12
                case "foo bar":
                    if p1 {
                        switch -2147483648 {
                        case -9223372036854775808:
                            switch "\n\t" {
                            case "hello":
                                if ((("hello" + "foo bar") + "\n\t") <= "foo bar") {
                                    fmt.Println(v11, p1)
                                    p0 += (v11 << ((p2) & 63))
                                } else {
                                    v13 := []int{(fi10 >> (((p2 * 1)) & 63)), (fi10 | 42), fi10}
                                    _ = v13
                                }
                            default:
                                for fi14 := 0; fi14 < 8; fi14++ {
                                    var v15 float64 = (1e10 / ((1e10 / (1e10 + -2.25)) * 1e10))
                                    v15 = (1e10 - v15)
                                    var v16 float64 = ((v15 / float64(0.1)) + float64(1e10))
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            }
                        case -441:
                            var v17 int32 = 2147483647
                            _ = v17
                        default:
                            if p1 {
                                var v18 uint8 = 42
                                p0 = v11
                                if (((1 != 255) || ("" <= ("\n\t" + "hello"))) && p1) {
                                    p2 -= (p2 - (^(p2)))
                                } else {
                                    v19 := GenS1{F0: ("a" + "a"), F1: 1, F2: (fi10 + (fi10 & (fi10 >> ((p2) & 63))))}
                                    _ = v19
                                }
                                _ = v18
                            } else {
                                if (!(false || p1)) {
                                    var v20 int16 = 0
                                    _ = v20
                                }
                            }
                        }
                        fmt.Println(p2)
                        v21 := GenS2{F0: 972, F1: 128, F2: -32768}
                        _ = v21
                    } else {
                        p2 = (p2 >> (((p2 ^ 1)) & 63))
                        p0 = -128
                    }
                case "":
                    var v22 uint16 = 255
                    _ = v22
                }
                _ = fi10
                _ = v11
            }
        }
        if p1 {
            v23 := GenS3{F0: (p0 << (((p2 * p2)) & 63)), F1: 0}
            v24 := GenS2{F0: 477, F1: 127, F2: 129}
            _ = v23
            _ = v24
        }
    }
    if (p2 > 65535) {
        v25 := GenS4{F0: 28, F1: 438}
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                v28 := []int{(fi26 * (fi27 - -946)), fi27, -1, fi26}
                _ = fi27
                _ = v28
            }
            var v29 int32 = -2147483648
            _ = fi26
            _ = v29
        }
        _ = v25
    } else {
        var v30 int64 = -1
        p1 = (-1 <= -643)
        _ = v30
    }
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    return 115
}
func genFn1(p0 uint16, p1 int32) uint {
    var v31 string = "\n\t"
    v32 := GenS1{F0: "hello", F1: p0, F2: 128}
    v33 := []int{372, 127}
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    _ = v33
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 42
    var v3 int64 = -9223372036854775808
    v4 := []uint32{42, 42, 1, 924}
    var v5 int8 = 1
    var v6 bool = (!((v5 / ((v5) | 1)) < (v5 * (v5 - (v5 >> ((v2) & 63))))))
    v34 := GenS3{F0: (v5 << (((v2 + v2)) & 63)), F1: -1}
    if v6 {
        var v35 uint16 = 1
        _ = v35
    }
    _ = len(v4)
    if (2147483647 > 613) {
        fmt.Println(v5)
        v36 := GenS4{F0: 127, F1: 235}
        _ = v36
    } else {
        v37 := GenS3{F0: v5, F1: -833}
        var v38 int32 = -1
        if (!v6) {
            v39 := GenS5{F0: (101 >= v5), F1: "a"}
            for fi40 := 0; fi40 < 15; fi40++ {
                v1 -= ((v1 * float64(1.5)) / 1.0)
                _ = fi40
            }
            _ = v39
        } else {
            v4 = append(v4, 0)
            if (127 <= v5) {
                var v41 int16 = 32767
                switch -162 {
                case 756:
                    v1 = v1
                }
                _ = v41
            }
        }
        _ = v37
        _ = v38
    }
    var v42 float64 = (v1 + v1)
    if (42 <= 128) {
        for fi43 := 0; fi43 < 17; fi43++ {
            var v44 float64 = 1.5
            _ = fi43
            _ = v44
        }
        if v6 {
            if (219 < genFn0(v5, (false && true), (v2 << ((v2) & 63)))) {
                fmt.Println(v1, v42)
            }
            if (!(!(165 < 1))) {
                v1 = v42
            }
            for fi45 := 0; fi45 < 17; fi45++ {
                fmt.Println(v2, v6)
                fmt.Println(v5)
                switch fi45 {
                case -807:
                    var v46 uint8 = 1
                    _ = v46
                case 32767:
                    var v47 uint16 = 32767
                    _ = v47
                default:
                    var v48 int64 = (v3 ^ (-(v3)))
                    _ = v48
                }
                _ = fi45
            }
        } else {
            for fi49 := 0; fi49 < 19; fi49++ {
                v5 += v5
                for fi50 := 0; fi50 < 10; fi50++ {
                    v3 = v3
                    v51 := GenS3{F0: (v5 % ((-1) | 1)), F1: -1}
                    if (((484 < 2147483647) || v6) || (127 <= fi50)) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v52 uint32 = uint32(v2)
                        var v53 float64 = ((float64(0.0) / (3.14159 + 1e10)) / v42)
                        if v6 {
                            var v54 int = fi49
                            _ = v54
                        } else {
                            v4 = append(v4, 2147483647)
                        }
                        _ = v52
                        _ = v53
                    } else {
                        v3 += (-(v3))
                        if true {
                            var v55 float32 = ((((float32(0.1) / 1.5) * 1.5) * 0.0) * 0.0)
                            var v56 uint = genFn1(1, -599)
                            _ = v55
                            _ = v56
                        } else {
                            fmt.Println(fi50, v3, fi49)
                            v6 = true
                            _ = len(v4)
                        }
                        if len(v4) > 0 {
                            v4[0] = 770
                        }
                    }
                    _ = fi50
                    _ = v51
                }
                v5 += (v5 + v5)
                _ = fi49
            }
            switch 128 {
            case 9223372036854775807:
                switch ("hello" + "foo bar") {
                case "hello":
                    if (float64(1e10) < ((3.14159 + v1) + v42)) {
                        v57 := GenS6{F0: 936, F1: 4294967295}
                        v4 = append(v4, 0)
                        _ = v57
                    }
                case "":
                    v58 := GenS5{F0: ("\n\t" < ("hello" + "")), F1: "foo bar"}
                    _ = v58
                case "\n\t":
                    var v59 bool = false
                    _ = v59
                }
            }
            var v60 string = ("" + "\n\t")
            _ = v60
        }
        v5 -= (v5 << ((v2) & 63))
    } else {
        if false {
            var v61 int64 = (-(v3))
            if true {
                v2 -= (v2 - v2)
                if (true || false) {
                    var v62 int16 = 32767
                    _ = v62
                } else {
                    if v6 {
                        for fi63 := 0; fi63 < 2; fi63++ {
                            switch "foo bar" {
                            case "a":
                                v3 = (v3 >> (((^(v2))) & 63))
                            case "":
                                var v64 uint = (v2 / ((v2) | 1))
                                _ = v64
                            case "foo bar":
                                var v65 float32 = float32(((((0.0 + float64(1.5)) / 3.14159) + v1) * (v42 / 3.14159)))
                                _ = v65
                            }
                            v2 = (v2 & 65535)
                            _ = fi63
                        }
                        var v66 float32 = 3.14159
                        _ = v66
                    }
                    v4 = append(v4, 2147483648)
                    for fi67 := 0; fi67 < 5; fi67++ {
                        if v6 {
                            switch "hello" {
                            case "hello":
                                v68 := GenS6{F0: 18446744073709551615, F1: 2}
                                _ = v68
                            case "":
                                fmt.Println(v61, v6)
                            case "\n\t":
                                v69 := GenS7{F0: 1}
                                _ = v69
                            }
                            v5 -= -127
                            var v70 bool = v6
                            _ = v70
                        } else {
                            v61 = -9223372036854775808
                            fmt.Println(v2)
                            var v71 uint16 = 2
                            _ = v71
                        }
                        if len(v4) > 0 {
                            v4[0] = 0
                        }
                        _ = fi67
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v61 = v61
                for fi72 := 0; fi72 < 7; fi72++ {
                    fmt.Println(v6)
                    _ = fi72
                }
            }
            _ = v61
        } else {
            var v73 uint = (v2 / (((v2 - v2)) | 1))
            switch 2147483647 {
            case -948:
                v74 := GenS4{F0: 42, F1: 18446744073709551615}
                _ = v74
            case 128:
                v5 += v5
            }
            v5 = -69
            _ = v73
        }
        var v75 int8 = (v5 + v5)
        _ = v75
    }
    switch 2147483647 {
    case 32767:
        if ((!(473 >= 1)) && (465 >= 4294967295)) {
            if len(v4) > 0 {
                v4[0] = 1
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    if (!(-221 != 1000)) {
        fmt.Println(v1)
    }
    if v6 {
        v42 = ((((v42 / v1) / v42) * (-2.25 + v42)) * 0.0)
        v6 = (false || (v3 > v3))
    } else {
        var v76 uint32 = 2147483648
        _ = v76
    }
    var v77 uint32 = 42
    v3 = v3
    var v78 int = -2147483648
    for fi79 := 0; fi79 < 15; fi79++ {
        for fi80 := 0; fi80 < 12; fi80++ {
            v5 += -10
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 -= (v3 << ((v2) & 63))
            _ = fi80
        }
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        var v81 uint16 = 0
        _ = fi79
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v77
    _ = v78
}
