package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string; F1 uint32; F2 int }
type GenS7 struct { F0 int8; F1 bool; F2 string }
type GenS8 struct { F0 string; F1 int64; F2 string }
func genFn0(p0 uint) uint16 {
    var v4 string = (("\n\t" + "foo bar") + "hello")
    fmt.Println(v4, v4)
    p0 += p0
    _ = p0
    _ = v4
    return uint16(p0)
}
func genFn1() string {
    v5 := []int{-12}
    var v6 uint64 = 1
    switch 1 {
    case 9223372036854775807:
        var v7 uint8 = 155
        _ = v7
    case 2147483647:
        v6 += v6
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint32 = 1
        v10 := GenS1{F0: 32767, F1: 2}
        v6 = (^(v6))
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 bool = ((127 != 4294967295) && true)
    _ = v11
    _ = v5
    _ = v6
    return (("\n\t" + ("" + "\n\t")) + "hello")
}

func main() {
    var v1 uint = 657
    var v2 uint64 = 375
    v3 := []int{}
    if (!true) {
        var v12 uint32 = 1
        _ = v12
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            v3 = append(v3, (-(fi13)))
            fmt.Println(v2, v1, v2)
            if ((true && false) || true) {
                var v14 uint32 = 1
                var v15 uint = v1
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    fmt.Println(v2)
                    _ = fi16
                }
                if (!true) {
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v1 = (v1 * (v1 - (^(v1))))
                        var v18 uint8 = 255
                        _ = fi17
                        _ = v18
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = fi13
                    }
                    _ = len(v3)
                    v19 := GenS1{F0: -32768, F1: (v1 * 0)}
                    _ = v19
                }
                v1 += (v1 << ((v1) & 63))
            }
            _ = fi13
        }
        if false {
            if (!true) {
                var v20 float32 = 3.14159
                _ = v20
            }
            if ((-491 >= -1) || false) {
                if true {
                    var v21 float32 = 1.5
                    _ = len(v3)
                    if false {
                        v22 := GenS1{F0: 967, F1: (v1 | v1)}
                        switch "a" {
                        case "\n\t":
                            if false {
                                v21 += -2.25
                                _ = len(v3)
                            } else {
                                var v23 uint8 = 32
                                v24 := GenS2{F0: 0}
                                v23 = v23
                                _ = v23
                                _ = v24
                            }
                        case "":
                            v1 = (v1 % (((v1 >> ((2) & 63))) | 1))
                        default:
                            v25 := GenS2{F0: genFn0(v1)}
                            _ = v25
                        }
                        var v26 float32 = ((v21 / v21) + (v21 - (v21 * v21)))
                        _ = v22
                        _ = v26
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 127
                        }
                    }
                    _ = v21
                } else {
                    if ((!false) && (127 > 42)) {
                        var v27 uint64 = 42
                        var v28 int = -9223372036854775808
                        v29 := GenS2{F0: 2}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v2 = v2
                fmt.Println(v1, v1)
            } else {
                var v30 int32 = 2147483647
                _ = v30
            }
            _ = len(v3)
        }
    }
    if (838 >= v2) {
        v31 := GenS2{F0: 65535}
        for fi32 := 0; fi32 < 12; fi32++ {
            for fi33 := 0; fi33 < 10; fi33++ {
                v1 += v1
                var v34 int8 = -1
                _ = fi33
                _ = v34
            }
            if (!true) {
                fmt.Println(v2)
                fmt.Println(fi32, v1, v1)
            } else {
                var v35 int8 = -128
                _ = v35
            }
            fmt.Println(fi32)
            _ = fi32
        }
        var v36 uint64 = (v2 ^ (v2 - (v2 & 408)))
        _ = v31
        _ = v36
    }
    var v37 int8 = 1
    v2 = (v2 % ((v2) | 1))
    _ = len(v3)
    var v38 uint8 = 42
    var v39 int8 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v37 = v39
    if true {
        var v40 int32 = -763
        _ = v40
    } else {
        fmt.Println(v38, v39)
        v41 := GenS1{F0: 1, F1: v1}
        _ = v41
    }
    var v42 uint = v1
    switch -102 {
    case 4294967295:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -2147483648:
        v43 := GenS3{F0: 909, F1: 42}
        _ = v43
    case 127:
        fmt.Println(v2, v2, v1)
    default:
        var v44 int32 = -613
        _ = v44
    }
    var v45 float32 = -2.25
    switch 128 {
    case 2147483647:
        if ((v1 & 42) <= 255) {
            v46 := GenS4{F0: false, F1: v1, F2: -207}
            var v47 string = (("foo bar" + "") + ((("" + (("foo bar" + ((("a" + ("\n\t" + (("hello" + "a") + "a"))) + "a") + "foo bar")) + ("hello" + ("foo bar" + "hello")))) + "\n\t") + "foo bar"))
            if false {
                for fi48 := 0; fi48 < 1; fi48++ {
                    fmt.Println(v45, fi48, v37)
                    _ = fi48
                }
                var v49 uint8 = 0
                if (0 == int32(v1)) {
                    for fi50 := 0; fi50 < 11; fi50++ {
                        if false {
                            v51 := GenS2{F0: 255}
                            switch v47 {
                            case "\n\t":
                                var v52 bool = false
                                _ = v52
                            case "foo bar":
                                _ = len(v3)
                            case "a":
                                switch "a" {
                                case "hello":
                                    for fi53 := 0; fi53 < 9; fi53++ {
                                        if true {
                                            if true {
                                                if len(v3) > 0 {
                                                    v3[0] = fi50
                                                }
                                                fmt.Println(v37, v45)
                                                v42 -= v1
                                                var v54 float32 = v45
                                                fmt.Println(v37, v1, v38)
                                                _ = v54
                                            }
                                            var v55 string = v47
                                            v37 -= (v39 & v39)
                                            _ = v55
                                        } else {
                                            if (!true) {
                                                var v56 uint = (^(v42))
                                                v56 -= v1
                                                _ = v56
                                            } else {
                                                v57 := GenS2{F0: uint16(v38)}
                                                if true {
                                                    if false {
                                                        var v58 uint64 = v2
                                                        var v59 string = v47
                                                        if len(v3) > 0 {
                                                            v3[0] = fi50
                                                        }
                                                        v42 = (v1 * 65535)
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        fmt.Println(fi50, v37)
                                                        v60 := GenS3{F0: (^(v2)), F1: 1}
                                                        _ = v60
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = -665
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(fi50))
                                                    }
                                                }
                                                var v61 uint64 = (v2 * v2)
                                                _ = v57
                                                _ = v61
                                            }
                                            fmt.Println(v47)
                                            if true {
                                                v62 := GenS3{F0: (v2 | v2), F1: 933}
                                                var v63 int = fi53
                                                switch "hello" {
                                                case "a":
                                                    v42 -= v42
                                                case "":
                                                    var v64 string = v47
                                                    _ = v64
                                                default:
                                                    v42 -= (^(v1))
                                                }
                                                _ = v62
                                                _ = v63
                                            } else {
                                                v45 -= ((0.1 - ((float32((3.14159 / (((0.0 / float64((v45 - ((1e10 / 1.5) + 1e10)))) * 1.5) * ((-2.25 * (3.14159 / 1.5)) + 1.5)))) - ((1e10 * v45) + v45)) * ((v45 / -2.25) + (v45 / ((0.0 / 0.1) - v45))))) / (float32(1e10) / v45))
                                                v1 = 1
                                                var v65 float32 = v45
                                                switch v47 {
                                                case "a":
                                                    switch ((v47 + v47) + v47) {
                                                    case "\n\t":
                                                        var v66 int32 = 1
                                                        _ = v66
                                                    }
                                                default:
                                                    v67 := GenS5{F0: fi50}
                                                    _ = v67
                                                }
                                                _ = v65
                                            }
                                            v45 = float32(-2.25)
                                        }
                                        v68 := GenS2{F0: 65535}
                                        var v69 string = "foo bar"
                                        _ = fi53
                                        _ = v68
                                        _ = v69
                                    }
                                default:
                                    for fi70 := 0; fi70 < 5; fi70++ {
                                        if len(v3) > 0 {
                                            v3[0] = 4294967295
                                        }
                                        var v71 uint64 = (v2 - 50)
                                        fmt.Println(fi70)
                                        _ = fi70
                                        _ = v71
                                    }
                                }
                            }
                            v39 += v39
                            _ = v51
                        }
                        _ = fi50
                    }
                }
                _ = v49
            } else {
                v42 = 1
                v45 += v45
                for fi72 := 0; fi72 < 19; fi72++ {
                    var v73 float32 = v45
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v73 -= 1e10
                    _ = fi72
                    _ = v73
                }
            }
            _ = v46
            _ = v47
        }
    default:
        v3 = append(v3, -581)
    }
    var v74 uint16 = 2
    switch 9223372036854775807 {
    case 1:
        for fi75 := 0; fi75 < 18; fi75++ {
            var v76 uint32 = 65535
            v77 := GenS6{F0: (("\n\t" + ((("foo bar" + ((("\n\t" + "hello") + "a") + "\n\t")) + "") + ((("a" + (("" + "hello") + "")) + "") + ""))) + "a"), F1: uint32(v2), F2: (-(fi75))}
            _ = fi75
            _ = v76
            _ = v77
        }
    case -2147483648:
        if true {
            v42 = v1
        } else {
            var v78 uint8 = 0
            v79 := GenS7{F0: (v39 ^ 127), F1: (0 < 0), F2: "\n\t"}
            for fi80 := 0; fi80 < 3; fi80++ {
                for fi81 := 0; fi81 < 10; fi81++ {
                    if true {
                        v82 := GenS8{F0: "foo bar", F1: 9223372036854775807, F2: ("hello" + ("" + "\n\t"))}
                        _ = v82
                    }
                    _ = fi81
                }
                _ = fi80
            }
            _ = v78
            _ = v79
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v42
    _ = v45
    _ = v74
}
