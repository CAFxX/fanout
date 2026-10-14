package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint64; F1 string; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint16; F1 int16; F2 int32 }
func genFn0(p0 uint8, p1 int16, p2 uint32) int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 127, F1: (p0 >> ((0) & 63)), F2: ((!false) && false)}
        for fi9 := 0; fi9 < 10; fi9++ {
            switch (fi9 & fi9) {
            case 127:
                fmt.Println(p2, fi7)
            case -408:
                v10 := GenS2{F0: 42, F1: -9223372036854775808}
                _ = v10
            default:
                p0 += (p0 >> ((228) & 63))
            }
            if ((p0 >> ((0) & 63)) > 219) {
                v11 := []int{}
                v12 := GenS1{F0: 42, F1: p0, F2: true}
                if (1 <= 2) {
                    switch (fi7 | 9223372036854775807) {
                    case 1:
                        v11 = append(v11, (fi7 % (((^(fi7))) | 1)))
                    case 0:
                        var v13 int64 = 146
                        _ = v13
                    default:
                        var v14 int = -1
                        _ = v14
                    }
                } else {
                    var v15 int32 = 1
                    _ = v15
                }
                _ = v11
                _ = v12
            } else {
                var v16 int64 = -9223372036854775808
                v16 += (v16 << ((864) & 63))
                if (uint(p2) <= 373) {
                    var v17 int32 = -2147483648
                    v18 := []int{(fi9 & 127), 2147483647, fi7}
                    p1 += (p1 ^ (p1 * p1))
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    p0 = 2
    p1 -= 32767
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{"", ("foo bar" + "hello"), "\n\t", ("" + ("" + ""))}
    var v3 int16 = 1000
    var v4 bool = true
    var v5 uint16 = 255
    var v6 int64 = 1
    v19 := GenS1{F0: 119, F1: 42, F2: v4}
    fmt.Println(v4, v3, v1)
    if false {
        switch 32767 {
        case 2147483647:
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 uint16 = v5
                switch fi20 {
                case 0:
                    v22 := GenS1{F0: 100, F1: 42, F2: false}
                    _ = v22
                case 2147483647:
                    var v23 float32 = 1.5
                    _ = v23
                case 32768:
                    if true {
                        switch (fi20 & 9223372036854775807) {
                        case 0:
                            v6 -= 1
                        case -789:
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                        }
                    } else {
                        v5 = 0
                        v3 = v3
                    }
                }
                _ = fi20
                _ = v21
            }
        case 127:
            var v24 float64 = (-2.25 * 0.1)
            _ = v24
        }
        if v4 {
            _ = len(v2)
        }
    }
    switch 32768 {
    case -899:
        if v4 {
            var v25 bool = true
            _ = len(v2)
            v26 := GenS3{F0: (^(v5))}
            _ = v25
            _ = v26
        } else {
            v2 = append(v2, "")
        }
    }
    v5 = (v5 - (v5 | (^(v5))))
    if v4 {
        v27 := GenS3{F0: v5}
        if false {
            if v4 {
                v28 := GenS1{F0: 1, F1: 0, F2: v4}
                _ = v28
            }
            if v4 {
                _ = len(v2)
                var v29 int16 = (v3 - 42)
                _ = v29
            } else {
                var v30 float32 = 1.5
                _ = v30
            }
            v31 := GenS4{F0: (v1 - (^(v1))), F1: v5, F2: (^(v6))}
            _ = v31
        } else {
            v1 = (^(v1))
            if ((v4 && false) && (4294967295 < 42)) {
                var v32 float32 = 3.14159
                switch "foo bar" {
                case "":
                    if (404 > 934) {
                        v5 = v5
                    }
                }
                v3 = (^(v3))
                _ = v32
            }
            v33 := GenS5{F0: 42, F1: "\n\t", F2: -1}
            _ = v33
        }
        fmt.Println(v1, v5, v4)
        _ = v27
    } else {
        switch "" {
        case "hello":
            if v4 {
                switch "\n\t" {
                case "":
                    for fi34 := 0; fi34 < 12; fi34++ {
                        v35 := GenS5{F0: 148, F1: ((("" + "a") + "hello") + ""), F2: 2147483647}
                        if len(v2) > 0 {
                            v2[0] = (("a" + "") + "foo bar")
                        }
                        _ = fi34
                        _ = v35
                    }
                }
            } else {
                for fi36 := 0; fi36 < 11; fi36++ {
                    v3 -= (^(v3))
                    for fi37 := 0; fi37 < 14; fi37++ {
                        v38 := GenS2{F0: (v5 % ((v5) | 1)), F1: (v6 / ((v6) | 1))}
                        if v4 {
                            fmt.Println(v4)
                            fmt.Println(v5)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if ((v3 >> ((309) & 63)) >= v3) {
                                if ((((0.1 / 1.5) / 1.0) * (1e10 + 3.14159)) > 0.1) {
                                    switch fi37 {
                                    case 9223372036854775807:
                                        if false {
                                            v6 = -1
                                            if ("foo bar" == "a") {
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                            } else {
                                                fmt.Println(v5)
                                            }
                                        } else {
                                            v5 = v5
                                            v39 := GenS6{F0: 0}
                                            _ = v39
                                        }
                                    case 42:
                                        var v40 uint = 65535
                                        _ = v40
                                    }
                                }
                                var v41 uint16 = (v5 ^ 620)
                                v41 += v5
                                _ = v41
                            } else {
                                if (!(-2.25 >= 1e10)) {
                                    v5 = v5
                                    if (!(int16(v5) == v3)) {
                                        var v42 uint16 = 625
                                        fmt.Println(fi37, v42)
                                        _ = v42
                                    } else {
                                        v5 -= 2
                                        var v43 uint = 2
                                        v2 = append(v2, "a")
                                        _ = v43
                                    }
                                    var v44 float64 = float64(((0.0 * 0.0) + 0.0))
                                    _ = v44
                                } else {
                                    v1 = v1
                                    switch fi37 {
                                    case -132:
                                        v2 = append(v2, "hello")
                                    case -920:
                                        if v4 {
                                            v4 = v4
                                            switch "\n\t" {
                                            case "hello":
                                                var v45 int16 = (v3 >> ((521) & 63))
                                                _ = v45
                                            default:
                                                var v46 string = ("a" + "a")
                                                _ = v46
                                            }
                                        }
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi36, fi36)
                                    var v47 uint32 = (v1 + (^(v1)))
                                    _ = v47
                                }
                                var v48 float32 = float32(0.1)
                                _ = v48
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if v4 {
                                v5 = 65535
                            }
                            switch 4294967295 {
                            case 4294967295:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 9223372036854775807:
                                v5 = (v5 + uint16(v1))
                            case 308:
                                if v4 {
                                    if v4 {
                                        if len(v2) > 0 {
                                            v2[0] = ("\n\t" + "hello")
                                        }
                                    } else {
                                        switch -360 {
                                        case 4294967295:
                                            fmt.Println(fi37)
                                        }
                                        var v49 string = "\n\t"
                                        v50 := GenS7{F0: (false && true), F1: 899}
                                        if v4 {
                                            if ((fi37 & (fi36 & (-(fi36)))) >= -25) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            v5 -= 102
                                            var v51 float64 = (1.5 + ((1.5 + 3.14159) - ((-2.25 - 1e10) * float64((1.5 + 3.14159)))))
                                            _ = v51
                                        } else {
                                            var v52 uint16 = v5
                                            v1 -= 42
                                            v5 -= (v52 | (v52 & (v52 & v52)))
                                            var v53 int = fi36
                                            v54 := GenS3{F0: v5}
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    v5 -= (^(v5))
                                    var v55 uint32 = (^(v1))
                                    _ = v55
                                } else {
                                    if (((3.14159 - float32(0.0)) - 3.14159) > 3.14159) {
                                        v4 = (((1e10 * 1.5) + float32(1.5)) < (0.0 * 0.0))
                                    }
                                }
                            }
                        }
                        if v4 {
                            if v4 {
                                v5 = (v5 & v5)
                                if ((((0 > v5) || v4) || (((!(v6 <= (v6 % ((-798) | 1)))) || (v4 || (v6 > v6))) || ((349 > 0) || v4))) && (uint16(v1) <= v5)) {
                                    switch "hello" {
                                    case "":
                                        v4 = v4
                                    }
                                    if true {
                                        _ = len(v2)
                                        v1 += v1
                                    }
                                } else {
                                    var v56 bool = true
                                    _ = v56
                                }
                                _ = len(v2)
                            } else {
                                switch fi36 {
                                case 1:
                                    var v57 uint64 = 0
                                    _ = v57
                                case 9223372036854775807:
                                    v2 = append(v2, "")
                                case 0:
                                    fmt.Println(v6)
                                }
                                fmt.Println(fi36)
                                v58 := GenS8{F0: 807, F1: v3, F2: -2147483648}
                                _ = v58
                            }
                            v5 += 42
                            if (true || (!true)) {
                                v6 -= v6
                            }
                        } else {
                            v3 = (v3 % ((v3) | 1))
                        }
                        _ = fi37
                        _ = v38
                    }
                    if v4 {
                        v59 := GenS1{F0: 100, F1: 60, F2: (((0.1 * -2.25) / (-2.25 + (3.14159 + 0.0))) != (1.5 * 1e10))}
                        _ = v59
                    } else {
                        v2 = append(v2, ("hello" + ("a" + "\n\t")))
                    }
                    _ = fi36
                }
            }
        case "":
            var v60 float32 = 1.5
            _ = v60
        case "a":
            v61 := GenS4{F0: (v1 << ((0) & 63)), F1: v5, F2: (-(v6))}
            _ = v61
        default:
            _ = len(v2)
        }
    }
    fmt.Println(v5, v4, v6)
    if (!v4) {
        for fi62 := 0; fi62 < 6; fi62++ {
            var v63 uint64 = 1
            _ = fi62
            _ = v63
        }
    }
    for fi64 := 0; fi64 < 11; fi64++ {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        var v65 uint64 = 49
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 1; fi66++ {
        var v67 float32 = 3.14159
        if (-2.25 != float64((v67 - v67))) {
            if (((0.0 - v67) + v67) == (float32(0.1) + (v67 + 0.0))) {
                v5 = 2
                v3 = (v3 / ((v3) | 1))
                v5 = v5
            }
            var v68 int64 = (v6 << ((18446744073709551615) & 63))
            var v69 uint16 = v5
            _ = v68
            _ = v69
        }
        _ = fi66
        _ = v67
    }
    if (false || (4294967295 != 42)) {
        var v70 int8 = 100
        _ = v70
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v71 int64 = v6
        _ = v71
    }
    v4 = (379 >= 42)
    v6 = (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
