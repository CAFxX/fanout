package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 int16; F1 int8; F2 int32 }
type GenS8 struct { F0 uint; F1 bool }
func genFn0(p0 bool) int {
    p0 = (!((2 >= 65535) && (p0 || false)))
    p0 = (52 == 2147483648)
    _ = p0
    return 42
}
func genFn1(p0 uint64, p1 uint16, p2 bool) uint64 {
    var v4 string = "a"
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{(fi5 ^ fi5)}
        _ = fi5
        _ = v6
    }
    p2 = false
    v7 := []int{32767, 42, -783}
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            p1 = p1
            if p2 {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if p2 {
                        var v10 int8 = -100
                        if ((v10 - (^(v10))) == (v10 >> ((255) & 63))) {
                            fmt.Println(fi9, v10, p1)
                            v10 = (-(v10))
                        } else {
                            p1 = (^(p1))
                        }
                        fmt.Println(p0, fi9)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v11 uint16 = 388
                        if p2 {
                            var v12 uint16 = (v11 * 494)
                            if true {
                                var v13 int16 = 32767
                                _ = v13
                            }
                            _ = v12
                        }
                        _ = v10
                        _ = v11
                    }
                    v14 := GenS1{F0: (fi8 / (((fi8 * -1)) | 1)), F1: "hello", F2: false}
                    v7 = append(v7, fi9)
                    _ = fi9
                    _ = v14
                }
                fmt.Println(v4)
            }
            _ = fi8
        }
        var v15 int32 = 244
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return (p0 % ((781) | 1))
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 3.14159
    v3 := []int{-83, 128}
    for fi16 := 0; fi16 < 5; fi16++ {
        v2 -= 1e10
        _ = fi16
    }
    var v17 int8 = 127
    var v18 uint8 = 127
    delete(v1, "hello")
    if false {
        var v19 uint32 = 65535
        if ((!false) && (!("hello" < ((("hello" + ("foo bar" + "hello")) + "\n\t") + "\n\t")))) {
            v20 := GenS1{F0: -363, F1: ("" + "\n\t"), F2: false}
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 2; fi22++ {
                    if (true && (!(!false))) {
                        v23 := GenS1{F0: fi21, F1: "hello", F2: (!true)}
                        if len(v3) > 0 {
                            v3[0] = (fi21 + (-(fi22)))
                        }
                        var v24 float64 = v2
                        _ = v23
                        _ = v24
                    } else {
                        v25 := v1["foo bar"]
                        _ = v25
                    }
                    _ = fi22
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi21
            }
            if true {
                if false {
                    var v26 bool = ((0 != v17) && true)
                    v18 = v18
                    _ = v26
                } else {
                    _ = len(v3)
                    if (v18 == (v18 * (v18 << ((42) & 63)))) {
                        var v27 float32 = -2.25
                        _ = len(v1)
                        if false {
                            if true {
                                for fi28 := 0; fi28 < 13; fi28++ {
                                    v17 = v17
                                    v1[("\n\t" + "a")] = "hello"
                                    var v29 uint32 = (v19 & (v19 << ((2) & 63)))
                                    _ = fi28
                                    _ = v29
                                }
                                v18 = v18
                                v30 := GenS1{F0: genFn0(true), F1: "foo bar", F2: (true || true)}
                                _ = v30
                            }
                        }
                        _ = v27
                    }
                }
            }
            _ = v20
        } else {
            var v31 int32 = -1
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 uint = 4294967295
                if (true && false) {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        v18 = v18
                        _ = fi34
                    }
                }
                var v35 float64 = v2
                _ = fi32
                _ = v33
                _ = v35
            }
            v36 := v1[""]
            _ = v31
            _ = v36
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            delete(v1, "foo bar")
            fmt.Println(fi37)
            v17 -= (v17 & v17)
            _ = fi37
        }
        _ = v19
    } else {
        v38 := v1["a"]
        for fi39 := 0; fi39 < 13; fi39++ {
            for fi40 := 0; fi40 < 2; fi40++ {
                fmt.Println(v18, v38, v38)
                _ = fi40
            }
            v1[("foo bar" + v38)] = "a"
            switch ("a" + "\n\t") {
            case "":
                if false {
                    v17 = (v17 + (-(v17)))
                    if (993 < 909) {
                        v1[((v38 + "hello") + v38)] = (v38 + "foo bar")
                        v41 := GenS1{F0: (-(fi39)), F1: (v38 + v38), F2: true}
                        _ = v41
                    } else {
                        v42 := GenS2{F0: false}
                        for fi43 := 0; fi43 < 3; fi43++ {
                            var v44 uint64 = 4294967295
                            if true {
                                if (1 != 1) {
                                    v2 -= 3.14159
                                    var v45 int64 = -9223372036854775808
                                    _ = v45
                                } else {
                                    v38 = (v38 + ("foo bar" + v38))
                                }
                                if (((v38 + v38) + v38) >= (("a" + "foo bar") + v38)) {
                                    fmt.Println(v2)
                                } else {
                                    var v46 int = int(v17)
                                    var v47 int32 = 2147483647
                                    v48 := GenS3{F0: v17, F1: (false && (!(!(false || (v44 > (v44 | v44))))))}
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                v38 = (("a" + v38) + v38)
                            } else {
                                v17 += (v17 >> ((0) & 63))
                                var v49 int16 = -508
                                var v50 uint8 = (v18 + v18)
                                _ = v49
                                _ = v50
                            }
                            v51 := GenS4{F0: false}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi43
                            _ = v44
                            _ = v51
                        }
                        _ = v42
                    }
                }
            case "a":
                v52 := GenS5{F0: v18, F1: -1}
                _ = v52
            case "hello":
                v2 = (1.5 * (1e10 * ((1.5 * v2) - v2)))
            }
            _ = fi39
        }
        _ = v38
    }
    switch "\n\t" {
    case "hello":
        if (!false) {
            v3 = append(v3, genFn0(true))
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 int64 = 0
                var v55 int32 = 335
                v2 -= 0.0
                _ = fi53
                _ = v54
                _ = v55
            }
            for fi56 := 0; fi56 < 13; fi56++ {
                if len(v3) > 0 {
                    v3[0] = 127
                }
                _ = fi56
            }
        }
    case "":
        if true {
            var v57 int = 9223372036854775807
            _ = v57
        }
    }
    v18 = (^(v18))
    v58 := GenS1{F0: 2147483647, F1: "", F2: true}
    if ("hello" <= "foo bar") {
        switch "\n\t" {
        case "hello":
            v2 = v2
        case "":
            switch 0 {
            case 127:
                switch "a" {
                case "":
                    if true {
                        v59 := GenS2{F0: true}
                        v17 = (v17 ^ -128)
                        v18 = (v18 ^ v18)
                        _ = v59
                    } else {
                        delete(v1, "\n\t")
                        var v60 uint16 = 2
                        _ = v60
                    }
                default:
                    v18 -= (^(v18))
                }
            case 4294967295:
                if true {
                    fmt.Println(v17, v17)
                } else {
                    if len(v3) > 0 {
                        v3[0] = 520
                    }
                    fmt.Println(v17, v17)
                }
            }
        default:
            v61 := GenS3{F0: (^(v17)), F1: false}
            _ = v61
        }
        fmt.Println(v2)
    }
    fmt.Println(v17)
    if (!(true && (!false))) {
        v3 = append(v3, -648)
    } else {
        if (65535 <= 4294967295) {
            fmt.Println(v18, v18)
            var v62 uint = 255
            _ = v62
        }
        for fi63 := 0; fi63 < 8; fi63++ {
            if len(v3) > 0 {
                v3[0] = genFn0((0 != 0))
            }
            _ = fi63
        }
        switch ("a" + ("\n\t" + "a")) {
        case "\n\t":
            if ((1.5 / (float64((float32((v2 + v2)) + -2.25)) + v2)) != 1e10) {
                for fi64 := 0; fi64 < 1; fi64++ {
                    switch genFn0(false) {
                    case 2147483647:
                        v65 := GenS6{F0: 42, F1: (fi64 >> ((42) & 63))}
                        _ = v65
                    default:
                        var v66 bool = (v18 <= (v18 & v18))
                        _ = v66
                    }
                    _ = fi64
                }
                for fi67 := 0; fi67 < 19; fi67++ {
                    var v68 uint16 = 589
                    var v69 uint8 = v18
                    v69 -= 241
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            }
        default:
            v18 = 42
        }
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        if (false || (!false)) {
            v18 = v18
            fmt.Println(v18, fi70)
            v71 := GenS2{F0: true}
            _ = v71
        } else {
            fmt.Println(fi70)
            fmt.Println(fi70, v17, v18)
        }
        if (!(!false)) {
            var v72 uint = 652
            _ = v72
        }
        delete(v1, "\n\t")
        _ = fi70
    }
    if (false || true) {
        fmt.Println(v18)
    } else {
        for fi73 := 0; fi73 < 5; fi73++ {
            for fi74 := 0; fi74 < 20; fi74++ {
                var v75 string = "foo bar"
                if false {
                    v76 := GenS7{F0: -804, F1: v17, F2: -2147483648}
                    _ = v76
                }
                _ = fi74
                _ = v75
            }
            for fi77 := 0; fi77 < 7; fi77++ {
                v78 := GenS8{F0: 18446744073709551615, F1: (65535 >= 65535)}
                _ = fi77
                _ = v78
            }
            v18 -= v18
            _ = fi73
        }
        v17 -= (v17 % (((v17 * v17)) | 1))
    }
    if true {
        var v79 float32 = 1.5
        _ = v79
    }
    if false {
        for fi80 := 0; fi80 < 6; fi80++ {
            fmt.Println(fi80, fi80)
            var v81 float64 = v2
            _ = fi80
            _ = v81
        }
        var v82 int8 = v17
        fmt.Println(v17, v2, v17)
        _ = v82
    } else {
        v83 := v1[(("a" + "hello") + "hello")]
        v83 = "\n\t"
        _ = v83
    }
    var v84 int32 = 856
    fmt.Println(v84, v18, v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v58
    _ = v84
}
