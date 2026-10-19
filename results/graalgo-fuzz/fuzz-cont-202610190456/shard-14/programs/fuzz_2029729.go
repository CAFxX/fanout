package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 int32 }
type GenS10 struct { F0 bool; F1 string; F2 uint16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 int64; F1 bool; F2 uint32 }
type GenS14 struct { F0 uint32; F1 bool; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 bool = false
    v2 -= (v2 << ((0) & 63))
    v5 := GenS1{F0: v2, F1: 2, F2: (v2 * (v2 << ((0) & 63)))}
    v6 := []int{42, 161, 0, 363}
    v4 = ((1e10 <= 1e10) && true)
    var v7 int64 = 0
    v3[("foo bar" + "a")] = "\n\t"
    switch 1 {
    case 4294967295:
        v4 = v4
    case 1:
        fmt.Println(v2, v4)
    case 953:
        if (!(0.0 == (float32(3.14159) + 3.14159))) {
            if v4 {
                var v8 float32 = (0.0 * float32((1.5 / (((1.5 + (1.5 * (1e10 - (-2.25 / (-2.25 - float64(0.0)))))) * 3.14159) * (1.5 / ((3.14159 + ((((3.14159 + 3.14159) * 3.14159) * 3.14159) - -2.25)) / float64(float32((3.14159 + -2.25)))))))))
                _ = v8
            }
            switch "" {
            case "foo bar":
                v3["hello"] = (("" + "") + ("" + "a"))
            case "a":
                for fi9 := 0; fi9 < 3; fi9++ {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v4 = true
                        _ = fi10
                    }
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 int64 = v7
                        _ = fi11
                        _ = v12
                    }
                    _ = fi9
                }
            case "\n\t":
                v2 = v2
            default:
                fmt.Println(v7, v4, v2)
            }
            if true {
                if (v4 && true) {
                    v7 -= (v7 ^ v7)
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS2{F0: 42, F1: fi13, F2: 0}
                        v15 := GenS2{F0: 0, F1: (fi13 << ((65535) & 63)), F2: 4294967295}
                        if len(v6) > 0 {
                            v6[0] = -2147483648
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v16 uint64 = 824
                    _ = v16
                } else {
                    var v17 int8 = (v2 - v2)
                    for fi18 := 0; fi18 < 6; fi18++ {
                        switch "" {
                        case "":
                            var v19 uint = 0
                            _ = v19
                        }
                        _ = fi18
                    }
                    _ = v17
                }
                for fi20 := 0; fi20 < 4; fi20++ {
                    v6 = append(v6, 708)
                    if v4 {
                        var v21 string = "hello"
                        v22 := GenS2{F0: -294, F1: -340, F2: 673}
                        _ = v21
                        _ = v22
                    } else {
                        v2 += 42
                        _ = len(v6)
                    }
                    var v23 uint64 = 4294967295
                    _ = fi20
                    _ = v23
                }
                if len(v6) > 0 {
                    v6[0] = 128
                }
            }
        } else {
            for fi24 := 0; fi24 < 17; fi24++ {
                if v4 {
                    fmt.Println(fi24)
                    v25 := GenS1{F0: 0, F1: 2, F2: v2}
                    v26 := GenS2{F0: -1, F1: (fi24 - fi24), F2: 937}
                    _ = v25
                    _ = v26
                } else {
                    v2 = v2
                    if v4 {
                        var v27 uint16 = 32767
                        _ = v27
                    } else {
                        v4 = v4
                        if false {
                            v7 = v7
                            v2 += 1
                        } else {
                            if len(v6) > 0 {
                                v6[0] = (^(fi24))
                            }
                            v28 := GenS3{F0: 0, F1: 735}
                            switch "\n\t" {
                            case "":
                                v4 = true
                            case "a":
                                v3["hello"] = "hello"
                            case "foo bar":
                                v29 := GenS4{F0: false}
                                _ = v29
                            default:
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v31 := v3[("foo bar" + "a")]
                                    _ = fi30
                                    _ = v31
                                }
                            }
                            _ = v28
                        }
                    }
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    switch fi24 {
                    case 42:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v33 := GenS5{F0: v2}
                    if false {
                        var v34 uint8 = 8
                        _ = v34
                    } else {
                        var v35 uint = 100
                        _ = v35
                    }
                    _ = fi32
                    _ = v33
                }
                _ = fi24
            }
        }
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 18; fi37++ {
            var v38 int = (-(fi37))
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    v2 -= 0
    fmt.Println(v7, v2)
    if true {
        v4 = (v4 && v4)
        var v39 float32 = 1.5
        _ = v39
    } else {
        v40 := GenS6{F0: 127}
        delete(v1, "\n\t")
        v41 := GenS1{F0: v2, F1: 18446744073709551615, F2: v2}
        _ = v40
        _ = v41
    }
    if false {
        for fi42 := 0; fi42 < 17; fi42++ {
            v3[""] = "foo bar"
            for fi43 := 0; fi43 < 8; fi43++ {
                v44 := GenS7{F0: 0}
                delete(v1, "\n\t")
                _ = fi43
                _ = v44
            }
            if false {
                if (!(!false)) {
                    for fi45 := 0; fi45 < 11; fi45++ {
                        _ = len(v6)
                        switch -9223372036854775808 {
                        case -142:
                            if v4 {
                                if ((v4 || v4) || (true && (false && v4))) {
                                    v46 := GenS8{F0: 1}
                                    fmt.Println(v7)
                                    v2 = v2
                                    _ = v46
                                } else {
                                    var v47 int = (-(fi42))
                                    v48 := v3[(("a" + "foo bar") + "hello")]
                                    _ = v47
                                    _ = v48
                                }
                            } else {
                                v4 = v4
                            }
                        case 0:
                            v4 = (!(!v4))
                        }
                        if ((155 > 4294967295) && true) {
                            if (!((!true) || v4)) {
                                v49 := GenS9{F0: 2, F1: 2147483647}
                                v50 := v3[("\n\t" + "a")]
                                delete(v3, "foo bar")
                                _ = v49
                                _ = v50
                            }
                            v51 := GenS10{F0: v4, F1: "hello", F2: 32767}
                            var v52 int16 = -18
                            v2 = -105
                            var v53 uint8 = 128
                            if (v53 != 0) {
                                var v54 int = fi45
                                switch (v54 | -771) {
                                case 0:
                                    fmt.Println(fi42, v4, v54)
                                case 4294967295:
                                    var v55 int16 = (-(v52))
                                    _ = v55
                                default:
                                    fmt.Println(v52, v52, v7)
                                }
                                _ = v54
                            } else {
                                v56 := GenS9{F0: 158, F1: 42}
                                fmt.Println(v52)
                                _ = v56
                            }
                            var v57 uint32 = 42
                            v58 := GenS4{F0: ((-(v7)) == -113)}
                            switch 752 {
                            case -644:
                                v59 := GenS11{F0: v57}
                                _ = v59
                            }
                            if v4 {
                                v60 := GenS12{F0: ("foo bar" + ("a" + (("\n\t" + "hello") + "\n\t"))), F1: 622}
                                var v61 uint8 = v53
                                var v62 int64 = -1
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            if false {
                                if (0.0 > 0.0) {
                                    var v63 uint32 = (^(v57))
                                    _ = v63
                                } else {
                                    var v64 int8 = v2
                                    var v65 int16 = (v52 << ((65535) & 63))
                                    v66 := GenS13{F0: (v7 << ((18446744073709551615) & 63)), F1: true, F2: v57}
                                    _ = v64
                                    _ = v65
                                    _ = v66
                                }
                                if true {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v3["a"] = (("foo bar" + ("" + "a")) + "hello")
                                } else {
                                    var v67 float32 = ((1.5 - 1.5) / ((1e10 + (3.14159 / 1e10)) / 3.14159))
                                    var v68 uint8 = (v53 / ((12) | 1))
                                    _ = v67
                                    _ = v68
                                }
                            } else {
                                fmt.Println(v52, v7, fi45)
                                delete(v1, "\n\t")
                            }
                            var v69 int = (fi45 % ((0) | 1))
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v57
                            _ = v58
                            _ = v69
                        }
                        v70 := GenS14{F0: 1, F1: (!false), F2: v7}
                        v7 -= (v7 >> ((213) & 63))
                        var v71 float32 = 1.5
                        _ = fi45
                        _ = v70
                        _ = v71
                    }
                } else {
                    var v72 string = (("foo bar" + "") + "foo bar")
                    v73 := GenS5{F0: (v2 % ((v2) | 1))}
                    _ = v72
                    _ = v73
                }
            } else {
                var v74 int8 = v2
                v4 = v4
                var v75 int16 = 32767
                _ = v74
                _ = v75
            }
            _ = fi42
        }
    }
    var v76 uint = 2
    var v77 uint16 = 65535
    v2 += (v2 | -128)
    if (v2 < v2) {
        var v78 int32 = 0
        _ = v78
    } else {
        v2 = -70
    }
    var v79 int32 = 522
    fmt.Println(v76, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v79
}
