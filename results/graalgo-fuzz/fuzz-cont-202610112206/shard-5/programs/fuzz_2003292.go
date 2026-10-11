package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 string; F1 int32; F2 int32 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint64; F1 int32; F2 int8 }
type GenS8 struct { F0 uint }
func genFn0(p0 uint64) string {
    var v6 int32 = 42
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 1, F1: 131, F2: 128}
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, p0)
    var v9 bool = ((-1 > 32768) && true)
    _ = p0
    _ = v6
    _ = v9
    return ""
}
func genFn1(p0 uint64, p1 uint, p2 int64) string {
    p1 += p1
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: 42, F1: 42, F2: 255}
        var v12 uint8 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return ""
}

func main() {
    v1 := []string{}
    v2 := []string{"a"}
    var v3 uint = 72
    v4 := []int{-460, 275, 42, 127}
    v5 := make(map[int]bool)
    v13 := GenS1{F0: 122, F1: 1, F2: 2}
    v5[832] = true
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(fi14, v3)
        v15 := v5[fi14]
        if true {
            var v16 int = fi14
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    var v17 bool = ((235 >= 127) || (true && ((v3 ^ (v3 - (v3 & (v3 & v3)))) != (v3 / ((v3) | 1)))))
    switch 0 {
    case 310:
        v18 := GenS2{F0: -1, F1: 993}
        _ = v18
    default:
        var v19 int = 2147483647
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        if (!v17) {
            var v21 uint8 = 217
            v17 = (-1 <= 9223372036854775807)
            _ = v21
        }
        for fi22 := 0; fi22 < 3; fi22++ {
            if ((float64(3.14159) > 3.14159) && (v17 || (227 >= 255))) {
                v17 = v17
                v23 := GenS2{F0: 2147483647, F1: 1}
                fmt.Println(fi20, v3)
                _ = v23
            }
            _ = fi22
        }
        _ = fi20
    }
    if ((!(!false)) && v17) {
        var v24 uint64 = 0
        _ = v24
    }
    v4 = append(v4, 919)
    if (1 >= 241) {
        for fi25 := 0; fi25 < 20; fi25++ {
            fmt.Println(fi25)
            v17 = ((0 != 42) || (112 > -32768))
            var v26 uint64 = 2
            _ = fi25
            _ = v26
        }
    } else {
        switch "a" {
        case "\n\t":
            v3 -= v3
        default:
            if (!true) {
                var v27 uint8 = 128
                _ = v27
            } else {
                var v28 uint8 = 128
                var v29 uint = (v3 << ((v3) & 63))
                v17 = (!v17)
                _ = v28
                _ = v29
            }
        }
        fmt.Println(v3, v17)
    }
    var v30 uint = 65535
    v17 = false
    _ = len(v2)
    for fi31 := 0; fi31 < 15; fi31++ {
        switch (fi31 & (fi31 ^ (-(fi31)))) {
        case 963:
            v32 := GenS3{F0: "", F1: v17}
            _ = v32
        case -246:
            var v33 uint8 = 135
            _ = v33
        default:
            switch "hello" {
            case "\n\t":
                var v34 uint = (v30 & (v3 << ((v3) & 63)))
                _ = v34
            case "foo bar":
                if ((!false) || v17) {
                    switch "hello" {
                    case "foo bar":
                        if false {
                            fmt.Println(v30, v3, v17)
                            if (fi31 == fi31) {
                                if (871 <= 388) {
                                    var v35 uint = (^(v30))
                                    var v36 string = ""
                                    v36 = ("hello" + ("" + v36))
                                    _ = v35
                                    _ = v36
                                } else {
                                    _ = len(v5)
                                    if v17 {
                                        var v37 int32 = -742
                                        v38 := GenS2{F0: -450, F1: 2}
                                        for fi39 := 0; fi39 < 20; fi39++ {
                                            v17 = (94 < 1)
                                            v40 := GenS4{F0: 53, F1: (-128 >= -24)}
                                            _ = fi39
                                            _ = v40
                                        }
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v41 := GenS5{F0: "a", F1: 0, F2: 42}
                                        _ = v41
                                    }
                                }
                            } else {
                                if (0 != 529) {
                                    fmt.Println(v3, v17, fi31)
                                    if true {
                                        v3 += v30
                                        if (v17 && (2147483647 >= -2147483648)) {
                                            v17 = (!((v30 << ((65535) & 63)) >= (v3 & (v30 % (((v30 | (v30 % (((v30 >> (((v30 >> ((42) & 63))) & 63))) | 1)))) | 1)))))
                                            for fi42 := 0; fi42 < 14; fi42++ {
                                                v43 := GenS5{F0: "\n\t", F1: 652, F2: -606}
                                                var v44 string = ""
                                                var v45 bool = false
                                                _ = fi42
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                        }
                                        v17 = (-48 <= 0)
                                    }
                                    if (("a" + "\n\t") <= "a") {
                                        v46 := GenS1{F0: -1, F1: uint8(v30), F2: 255}
                                        _ = v46
                                    } else {
                                        if (v17 && (("foo bar" + "foo bar") == "hello")) {
                                            for fi47 := 0; fi47 < 19; fi47++ {
                                                var v48 int64 = -1
                                                _ = fi47
                                                _ = v48
                                            }
                                        }
                                        for fi49 := 0; fi49 < 9; fi49++ {
                                            if (v17 || v17) {
                                                var v50 uint8 = 227
                                                v50 -= 82
                                                v51 := GenS6{F0: 18, F1: fi31}
                                                _ = v50
                                                _ = v51
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v52 float64 = 1e10
                                            _ = fi49
                                            _ = v52
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = (fi31 - (fi31 >> ((v3) & 63)))
                                        }
                                    }
                                } else {
                                    for fi53 := 0; fi53 < 13; fi53++ {
                                        var v54 int64 = 1
                                        _ = fi53
                                        _ = v54
                                    }
                                    v55 := GenS7{F0: 18446744073709551615, F1: 224, F2: -128}
                                    _ = len(v1)
                                    _ = v55
                                }
                                v3 = (v3 << (((v30 - v3)) & 63))
                                fmt.Println(fi31)
                            }
                        }
                    default:
                        var v56 int64 = -301
                        _ = v56
                    }
                }
            case "":
                v3 += v30
            }
        }
        _ = fi31
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        v3 = v30
        _ = fi57
    }
    v58 := GenS8{F0: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v58
}
