package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 int }
func genFn0(p0 int32, p1 string) uint {
    v5 := []int{-363, 298, 9223372036854775807, -899}
    v6 := GenS1{F0: 0, F1: p1}
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 100, F1: ""}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint, p1 uint16) uint {
    if true {
        var v9 uint64 = 2
        switch -9223372036854775808 {
        case 42:
            p0 = (p0 | p0)
        default:
            if false {
                v10 := []int{32767}
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                v11 := GenS1{F0: -113, F1: (("hello" + (("hello" + "hello") + "a")) + "a")}
                _ = v10
                _ = v11
            } else {
                var v12 bool = (false && true)
                if (1 <= 42) {
                    var v13 bool = v12
                    _ = v13
                }
                _ = v12
            }
        }
        _ = v9
    }
    v14 := []int{0, -422, 127, -439}
    _ = p0
    _ = p1
    _ = v14
    return genFn0(-1, (("a" + "foo bar") + ("a" + ("a" + "hello"))))
}

func main() {
    var v1 bool = (-793 < -1)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 int = 128
    for fi15 := 0; fi15 < 7; fi15++ {
        if false {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 string = ("a" + "foo bar")
                _ = fi16
                _ = v17
            }
            v18 := GenS1{F0: 42, F1: "foo bar"}
            v19 := []int{v4, v4}
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint32 = 2
                v22 := GenS1{F0: -128, F1: "foo bar"}
                delete(v2, fi20)
                var v23 int8 = -128
                switch fi15 {
                case 32767:
                    v24 := GenS1{F0: (v23 % ((v23) | 1)), F1: "hello"}
                    _ = v24
                case 128:
                    if v1 {
                        _ = len(v2)
                        v25 := []int{2147483647, (fi15 << ((594) & 63)), fi20}
                        _ = v25
                    } else {
                        var v26 uint32 = (^(v21))
                        var v27 bool = (!(!((!(v1 && false)) || v1)))
                        _ = v26
                        _ = v27
                    }
                default:
                    v21 += (v21 - 2147483648)
                }
                v1 = ("hello" <= ("hello" + "a"))
                switch fi15 {
                case 549:
                    v28 := GenS2{F0: 65535}
                    _ = v28
                }
                v4 = 174
                var v29 int = (v4 / (((^(fi20))) | 1))
                var v30 float32 = 3.14159
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
                _ = v29
                _ = v30
            }
            delete(v3, ("" + (("a" + "") + "")))
        }
        var v31 bool = (!v1)
        _ = fi15
        _ = v31
    }
    for fi32 := 0; fi32 < 11; fi32++ {
        v4 += fi32
        _ = fi32
    }
    v33 := GenS3{F0: ""}
    for fi34 := 0; fi34 < 6; fi34++ {
        switch "" {
        case "foo bar":
            for fi35 := 0; fi35 < 5; fi35++ {
                v36 := GenS3{F0: ("" + "a")}
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi34, fi34)
                default:
                    v4 += (fi35 ^ (v4 & fi35))
                }
                _ = fi35
                _ = v36
            }
        case "\n\t":
            if v1 {
                for fi37 := 0; fi37 < 3; fi37++ {
                    v2[fi34] = 708
                    v38 := v2[42]
                    _ = fi37
                    _ = v38
                }
            } else {
                v39 := GenS4{F0: 42, F1: 255, F2: 370}
                if (88 == 0) {
                    fmt.Println(fi34, fi34, fi34)
                    v4 += v4
                }
                switch "hello" {
                case "a":
                    if ((!v1) || v1) {
                        v40 := GenS4{F0: 658, F1: 37, F2: 911}
                        fmt.Println(fi34, v4, fi34)
                        _ = v40
                    } else {
                        var v41 uint64 = 1
                        switch 790 {
                        case -1:
                            v3[""] = fi34
                        default:
                            v41 = 2
                        }
                        _ = v41
                    }
                case "":
                    v1 = false
                case "hello":
                    switch fi34 {
                    case 9223372036854775807:
                        v1 = true
                    default:
                        v1 = (42 != 24)
                    }
                }
                _ = v39
            }
        }
        _ = fi34
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        v4 = -820
        _ = fi42
    }
    v43 := GenS5{F0: 65535, F1: 4294967295, F2: genFn1(42, 1)}
    if false {
        v44 := GenS2{F0: 784}
        _ = v44
    } else {
        v1 = (-736 == 32767)
        for fi45 := 0; fi45 < 16; fi45++ {
            if (true && v1) {
                fmt.Println(fi45, fi45, v1)
            }
            v4 = fi45
            _ = fi45
        }
    }
    var v46 uint32 = 0
    v1 = v1
    if (590 < 1) {
        v46 -= v46
        if true {
            for fi47 := 0; fi47 < 3; fi47++ {
                v48 := []int{128}
                if v1 {
                    var v49 uint = 255
                    _ = v49
                } else {
                    var v50 float32 = (float32(((3.14159 / (0.0 + 0.1)) - float64((1e10 + (-2.25 + -2.25))))) / 1.0)
                    for fi51 := 0; fi51 < 6; fi51++ {
                        var v52 int16 = 32767
                        _ = fi51
                        _ = v52
                    }
                    v53 := GenS6{F0: 1, F1: v46}
                    _ = v50
                    _ = v53
                }
                _ = fi47
                _ = v48
            }
        } else {
            _ = len(v2)
            v54 := GenS7{F0: -835, F1: 162}
            _ = v54
        }
        for fi55 := 0; fi55 < 1; fi55++ {
            var v56 bool = v1
            var v57 int8 = 42
            if (("foo bar" + ("a" + "foo bar")) < ((("\n\t" + ("a" + "foo bar")) + "hello") + ("hello" + "\n\t"))) {
                var v58 float32 = (1e10 + 1e10)
                var v59 int64 = -462
                _ = v58
                _ = v59
            } else {
                for fi60 := 0; fi60 < 6; fi60++ {
                    v1 = (!v56)
                    _ = fi60
                }
            }
            _ = fi55
            _ = v56
            _ = v57
        }
    }
    v61 := GenS8{F0: v4}
    if v1 {
        fmt.Println(v46, v1)
        fmt.Println(v1, v46, v1)
    } else {
        var v62 int64 = -273
        _ = v62
    }
    v63 := GenS3{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v46
    _ = v61
    _ = v63
}
