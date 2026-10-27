package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS3 struct { F0 int32; F1 int16; F2 uint64 }
type GenS4 struct { F0 bool; F1 int; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 bool; F1 uint64; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint8; F1 int32; F2 uint32 }
func genFn0(p0 int8, p1 int8) int16 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int32 = 2147483647
        var v8 uint64 = 4294967295
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 248, F1: true, F2: ""}
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 0.1
    var v3 uint32 = 0
    v4 := make(map[string]bool)
    var v5 float32 = -2.25
    switch 849 {
    case 127:
        v5 = float32((v2 / (v2 - v2)))
    case 32768:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: 347, F1: false, F2: "a"}
            if true {
                var v12 int64 = -9223372036854775808
                _ = len(v1)
                v5 = v5
                _ = v12
            } else {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := []int{(-(fi13)), (fi10 ^ fi10), (fi10 - (-(fi10)))}
                    v5 -= v5
                    _ = fi13
                    _ = v14
                }
                var v15 bool = (!true)
                v16 := []int{fi10, (^(fi10))}
                _ = v15
                _ = v16
            }
            _ = fi10
            _ = v11
        }
    case 128:
        fmt.Println(v2, v3, v2)
    default:
        var v17 uint8 = 76
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        v3 += (v3 * 0)
        v3 -= (v3 % ((v3) | 1))
        v19 := GenS1{F0: 65535, F1: (9223372036854775807 != 42), F2: ("" + ("" + "foo bar"))}
        _ = fi18
        _ = v19
    }
    if ("\n\t" >= ("\n\t" + ("foo bar" + ("" + "foo bar")))) {
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(fi20, fi20)
            _ = fi20
        }
        var v21 int32 = 647
        var v22 uint32 = v3
        _ = v21
        _ = v22
    }
    v3 = 1
    for fi23 := 0; fi23 < 16; fi23++ {
        switch "" {
        case "":
            if (false || (("a" + "hello") == "hello")) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    v25 := GenS2{F0: (v3 | 295), F1: 887, F2: 42}
                    _ = fi24
                    _ = v25
                }
            } else {
                var v26 uint8 = 127
                if ((uint8(v3) > (v26 + (v26 + v26))) && false) {
                    switch "foo bar" {
                    case "\n\t":
                        for fi27 := 0; fi27 < 15; fi27++ {
                            fmt.Println(v5, fi23, v2)
                            v28 := GenS3{F0: -713, F1: 613, F2: 563}
                            _ = fi27
                            _ = v28
                        }
                    case "foo bar":
                        switch 4294967295 {
                        case 4294967295:
                            switch (fi23 / ((128) | 1)) {
                            case 9223372036854775807:
                                v3 -= 4294967295
                            default:
                                v29 := v4["hello"]
                                _ = v29
                            }
                        case 9223372036854775807:
                            v5 += (v5 - (((v5 * -2.25) - ((float32(1e10) - float32((v2 + 1.5))) + 3.14159)) + ((v5 / 1e10) - v5)))
                        default:
                            v30 := GenS4{F0: false, F1: (fi23 | fi23), F2: -9223372036854775808}
                            _ = v30
                        }
                    case "a":
                        if ((205 <= 2) && (!(!true))) {
                            for fi31 := 0; fi31 < 7; fi31++ {
                                v3 = v3
                                _ = fi31
                            }
                            v2 += v2
                        }
                    }
                    if ((int64(v3) <= 1) || (!true)) {
                        v32 := []int{4294967295, (fi23 | -88), fi23}
                        for fi33 := 0; fi33 < 20; fi33++ {
                            v3 += (v3 / (((v3 % (((v3 + (v3 / (((v3 & (v3 - 359))) | 1)))) | 1))) | 1))
                            _ = fi33
                        }
                        _ = v32
                    }
                    if true {
                        if true {
                            v34 := GenS5{F0: (^(v26))}
                            var v35 int32 = 0
                            var v36 uint = 818
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS2{F0: v3, F1: 2, F2: -879}
                        v38 := []int{fi23, (fi23 / (((fi23 % ((fi23) | 1))) | 1)), 127}
                        _ = v37
                        _ = v38
                    }
                }
                v2 = v2
                _ = v26
            }
        default:
            if false {
                v2 = 1e10
                v39 := GenS4{F0: (true && false), F1: (fi23 / ((fi23) | 1)), F2: 42}
                _ = v39
            } else {
                v3 -= (^(v3))
                fmt.Println(v5)
            }
        }
        v40 := GenS3{F0: 42, F1: 1000, F2: 2}
        v41 := GenS4{F0: false, F1: fi23, F2: 938}
        _ = fi23
        _ = v40
        _ = v41
    }
    v42 := GenS6{F0: -9, F1: (!true)}
    if false {
        var v43 string = ""
        v44 := GenS2{F0: v3, F1: 813, F2: 9223372036854775807}
        _ = v43
        _ = v44
    }
    _ = len(v4)
    if (0 < 808) {
        v45 := GenS7{F0: 581, F1: 99}
        switch ((("a" + "") + (("a" + ("hello" + "\n\t")) + "hello")) + ("hello" + "a")) {
        case "hello":
            v5 = v5
        default:
            switch (("\n\t" + (((("foo bar" + ("foo bar" + "\n\t")) + "a") + "\n\t") + "")) + "") {
            case "":
                v5 = ((0.0 + v5) / v5)
            }
        }
        if (257 != v3) {
            for fi46 := 0; fi46 < 16; fi46++ {
                v47 := []int{}
                _ = fi46
                _ = v47
            }
            fmt.Println(v5)
            if (true || (!true)) {
                var v48 bool = (((((float32(v2) / 0.1) - v5) * v5) > (float32(v2) - ((float32(v2) / v5) - v5))) || false)
                v1[(v3 & v3)] = v48
                v1[(v3 >> ((1) & 63))] = v48
                _ = v48
            }
        } else {
            var v49 uint = 1
            v3 -= 193
            var v50 string = "hello"
            _ = v49
            _ = v50
        }
        _ = v45
    } else {
        switch "" {
        case "\n\t":
            var v51 uint64 = 2
            _ = v51
        }
        v3 = 2147483648
        if true {
            if true {
                for fi52 := 0; fi52 < 8; fi52++ {
                    var v53 uint16 = 42
                    fmt.Println(v5, v53)
                    v54 := GenS8{F0: (^(fi52)), F1: genFn0(63, -128)}
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            }
            v55 := GenS4{F0: false, F1: 128, F2: 9223372036854775807}
            v56 := []int{-247, 127, 1, 128}
            _ = v55
            _ = v56
        } else {
            if ((((v2 - float64(v5)) + (0.1 - v2)) + (v2 / -2.25)) > 0.1) {
                fmt.Println(v2, v2)
            } else {
                switch "\n\t" {
                case "hello":
                    for fi57 := 0; fi57 < 6; fi57++ {
                        v3 = 0
                        for fi58 := 0; fi58 < 20; fi58++ {
                            switch (("foo bar" + ("\n\t" + ((("" + "a") + "") + "a"))) + ((("\n\t" + ((("a" + "foo bar") + ("\n\t" + "\n\t")) + "hello")) + ("hello" + (("foo bar" + "\n\t") + "a"))) + "")) {
                            case "foo bar":
                                var v59 int8 = int8(fi57)
                                _ = v59
                            case "hello":
                                fmt.Println(v2, fi57, fi58)
                            case "":
                                v60 := v1[(v3 ^ v3)]
                                _ = v60
                            }
                            _ = fi58
                        }
                        _ = fi57
                    }
                case "foo bar":
                    var v61 uint32 = (^(v3))
                    _ = v61
                default:
                    v62 := GenS8{F0: 32767, F1: genFn0(-128, int8(v3))}
                    _ = v62
                }
            }
            v63 := GenS2{F0: v3, F1: 2, F2: -1}
            var v64 uint64 = 2
            _ = v63
            _ = v64
        }
    }
    if (v5 < v5) {
        switch 32767 {
        case 9223372036854775807:
            fmt.Println(v3, v2)
        case -713:
            v65 := GenS3{F0: 1, F1: 0, F2: 18446744073709551615}
            _ = v65
        case 128:
            for fi66 := 0; fi66 < 19; fi66++ {
                var v67 uint64 = 0
                v68 := []int{(-(fi66))}
                if true {
                    for fi69 := 0; fi69 < 20; fi69++ {
                        if true {
                            v70 := GenS9{F0: (!false), F1: v67, F2: int8(v3)}
                            v71 := GenS6{F0: 127, F1: (2147483647 <= 2147483647)}
                            v1[v3] = ((^(v3)) != (^(v3)))
                            _ = v70
                            _ = v71
                        }
                        var v72 float32 = 1.5
                        var v73 int = fi69
                        var v74 string = ""
                        _ = fi69
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                } else {
                    fmt.Println(fi66, fi66)
                }
                _ = fi66
                _ = v67
                _ = v68
            }
        default:
            if true {
                v2 = v2
                var v75 float64 = (v2 - v2)
                v4["\n\t"] = ((-32768 >= 1) && true)
                _ = v75
            }
        }
        switch 9223372036854775807 {
        case 0:
            v76 := GenS10{F0: 0}
            _ = v76
        case 42:
            for fi77 := 0; fi77 < 9; fi77++ {
                if (false && true) {
                    for fi78 := 0; fi78 < 13; fi78++ {
                        var v79 float64 = v2
                        _ = fi78
                        _ = v79
                    }
                    _ = len(v4)
                } else {
                    fmt.Println(v5, v3, fi77)
                    v5 += v5
                }
                v80 := GenS1{F0: 0, F1: false, F2: (("foo bar" + "") + "hello")}
                if false {
                    var v81 int8 = 122
                    _ = v81
                } else {
                    v82 := []int{(fi77 | (^(fi77)))}
                    var v83 uint8 = 51
                    _ = v82
                    _ = v83
                }
                _ = fi77
                _ = v80
            }
        case 1:
            v5 -= (float32(v2) + -2.25)
        default:
            v84 := []int{128, 220, 127, 128}
            _ = v84
        }
        for fi85 := 0; fi85 < 17; fi85++ {
            v86 := []int{-799, (fi85 / ((fi85) | 1)), fi85, (^(fi85))}
            _ = fi85
            _ = v86
        }
    } else {
        var v87 float32 = float32((((1e10 - 1.5) + v2) - 0.0))
        v88 := GenS6{F0: 42, F1: (0 != 1000)}
        for fi89 := 0; fi89 < 7; fi89++ {
            v2 -= (((((v2 / v2) + v2) + v2) * (float64(v5) - v2)) + v2)
            _ = fi89
        }
        _ = v87
        _ = v88
    }
    fmt.Println(v5, v5)
    v90 := GenS11{F0: 128, F1: -2147483648, F2: v3}
    var v91 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v90
    _ = v91
}
