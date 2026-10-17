package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 bool }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int16; F1 bool; F2 int32 }
func genFn0() uint64 {
    if (-1 > -1) {
        v6 := GenS1{F0: 255, F1: 1000, F2: true}
        v7 := GenS1{F0: 126, F1: 42, F2: true}
        var v8 int = 42
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 9; fi10++ {
        switch fi10 {
        case 388:
            var v11 uint = 2
            _ = v11
        case 2147483647:
            var v12 int = 1
            _ = v12
        }
        var v13 int8 = 127
        v14 := GenS2{F0: "hello", F1: (v9 / ((769) | 1)), F2: fi10}
        _ = fi10
        _ = v13
        _ = v14
    }
    fmt.Println(v9)
    v15 := []int{-887, -2147483648, 61, -325}
    _ = v15
    _ = v9
    return 527
}
func genFn1(p0 uint8, p1 uint64) int8 {
    v16 := GenS2{F0: "", F1: 914, F2: -222}
    if true {
        p0 += 55
        if (!(p1 == (^(p1)))) {
            p0 = p0
        } else {
            var v17 int8 = 1
            if (!false) {
                switch ("hello" + "hello") {
                case "\n\t":
                    if (42 >= -9223372036854775808) {
                        if true {
                            v18 := []int{0}
                            _ = v18
                        }
                        if (true && (((32767 < 32767) && (((-1 != 100) || (42 > 0)) && false)) || false)) {
                            v19 := GenS2{F0: (("hello" + "a") + (((("" + "hello") + "\n\t") + "a") + "foo bar")), F1: 0, F2: 4294967295}
                            for fi20 := 0; fi20 < 14; fi20++ {
                                p1 = (p1 / (((p1 % ((p1) | 1))) | 1))
                                _ = fi20
                            }
                            v21 := GenS1{F0: 1, F1: -32768, F2: true}
                            _ = v19
                            _ = v21
                        } else {
                            v22 := []int{-633, -2147483648}
                            _ = v22
                        }
                        v23 := []int{42, 32768, -930, 96}
                        _ = v23
                    }
                case "foo bar":
                    if ((!false) && ((18446744073709551615 == 0) || true)) {
                        if true {
                            var v24 int = -9223372036854775808
                            _ = v24
                        } else {
                            p1 = (p1 ^ (p1 >> ((0) & 63)))
                            for fi25 := 0; fi25 < 17; fi25++ {
                                v26 := []int{}
                                _ = fi25
                                _ = v26
                            }
                            if (!false) {
                                v27 := GenS2{F0: "\n\t", F1: 0, F2: 1}
                                var v28 int8 = v17
                                _ = v27
                                _ = v28
                            } else {
                                var v29 float64 = 0.0
                                _ = v29
                            }
                        }
                    }
                }
            }
            _ = v17
        }
        if false {
            if (false || false) {
                p0 = 29
            } else {
                p0 = 20
                for fi30 := 0; fi30 < 20; fi30++ {
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 bool = true
                        var v33 uint8 = (p0 / (((p0 % (((p0 / ((p0) | 1))) | 1))) | 1))
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = fi30
                }
                p1 = genFn0()
            }
            if true {
                for fi34 := 0; fi34 < 18; fi34++ {
                    p1 += p1
                    _ = fi34
                }
            } else {
                if (42 >= 814) {
                    v35 := GenS3{F0: 146}
                    _ = v35
                } else {
                    switch "" {
                    case "\n\t":
                        p0 = p0
                    }
                }
            }
        } else {
            p0 += p0
            var v36 int8 = 1
            for fi37 := 0; fi37 < 6; fi37++ {
                switch fi37 {
                case 4294967295:
                    v38 := []int{}
                    _ = v38
                case 32768:
                    var v39 int64 = -810
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 int32 = -2147483648
                        _ = fi40
                        _ = v41
                    }
                }
                v42 := []int{fi37, 32767}
                p1 += 949
                _ = fi37
                _ = v42
            }
            _ = v36
        }
    }
    var v43 int = 0
    var v44 float32 = ((float32(1e10) / (((3.14159 * 0.1) + 0.1) / 1.0)) / ((0.0 - -2.25) + float32(float64(((((float32(3.14159) - -2.25) + 1.5) / (((float32(1e10) * 1.5) / 1.0) - 0.1)) * float32(float64((float32((((0.1 + 3.14159) - ((-2.25 / 1e10) / ((((0.1 / (1e10 + 0.1)) / (0.1 / 3.14159)) * 3.14159) * 1.5))) / float64(1.5))) + (-2.25 + 3.14159)))))))))
    _ = p0
    _ = p1
    _ = v16
    _ = v43
    _ = v44
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]int)
    var v3 float64 = float64(-2.25)
    v4 := make(map[uint32]string)
    var v5 uint16 = 1
    for fi45 := 0; fi45 < 1; fi45++ {
        if false {
            for fi46 := 0; fi46 < 4; fi46++ {
                if true {
                    if ((false && (v3 > 3.14159)) || (((v5 + (v5 / ((v5) | 1))) != 1) || false)) {
                        fmt.Println(fi46, fi45, v3)
                        var v47 bool = ((v5 * v5) >= v5)
                        v1 = (v1 * (v3 * (v1 + 0.0)))
                        _ = v47
                    } else {
                        _ = len(v4)
                        v5 = (v5 >> ((4294967295) & 63))
                    }
                    if true {
                        if false {
                            var v48 uint = 4294967295
                            switch ("a" + (("foo bar" + "") + "foo bar")) {
                            case "hello":
                                var v49 int16 = 42
                                _ = v49
                            case "foo bar":
                                var v50 uint64 = 2
                                switch fi46 {
                                case 2147483647:
                                    v48 = (v48 - (v48 - v48))
                                }
                                _ = v50
                            default:
                                v51 := GenS1{F0: uint8(v5), F1: -252, F2: true}
                                _ = v51
                            }
                            _ = v48
                        } else {
                            switch "" {
                            case "\n\t":
                                delete(v2, ("hello" + "a"))
                            }
                            v52 := GenS4{F0: 255, F1: (255 < 4294967295)}
                            _ = v52
                        }
                    } else {
                        var v53 int8 = genFn1(0, 1)
                        _ = v53
                    }
                    var v54 uint64 = 4294967295
                    _ = v54
                }
                _ = fi46
            }
            for fi55 := 0; fi55 < 11; fi55++ {
                switch "foo bar" {
                case "a":
                    if false {
                        _ = len(v4)
                    }
                case "hello":
                    if (0 >= 764) {
                        v3 = 0.1
                        if false {
                            switch "\n\t" {
                            case "a":
                                var v56 bool = (1 > 559)
                                _ = v56
                            case "":
                                if (188 > 0) {
                                    v5 -= 47
                                    var v57 uint = 18446744073709551615
                                    _ = v57
                                }
                            default:
                                v5 = v5
                            }
                            var v58 uint32 = 79
                            var v59 int8 = 100
                            _ = v58
                            _ = v59
                        }
                        var v60 float64 = 0.1
                        _ = v60
                    }
                default:
                    switch (^(fi45)) {
                    case -913:
                        delete(v4, 740)
                        delete(v2, "")
                    }
                }
                if true {
                    var v61 uint64 = 2
                    var v62 uint16 = 65535
                    _ = v61
                    _ = v62
                } else {
                    v5 = 0
                    v63 := v2["\n\t"]
                    _ = v63
                }
                if true {
                    v64 := GenS5{F0: 758}
                    _ = v64
                } else {
                    v5 -= v5
                    switch "hello" {
                    case "":
                        v1 = (v3 * (1.5 / 3.14159))
                    default:
                        v65 := []int{}
                        _ = v65
                    }
                }
                _ = fi55
            }
        } else {
            v5 = (v5 - v5)
            v66 := GenS2{F0: ("\n\t" + ""), F1: 9223372036854775807, F2: fi45}
            _ = v66
        }
        v4[0] = ""
        _ = fi45
    }
    v67 := GenS6{F0: 149}
    for fi68 := 0; fi68 < 14; fi68++ {
        fmt.Println(v3, v5, v1)
        _ = fi68
    }
    fmt.Println(v5, v5, v3)
    if true {
        delete(v4, 42)
        delete(v2, "a")
        var v69 int8 = 127
        _ = v69
    } else {
        v1 += (v1 * 1e10)
    }
    _ = len(v2)
    v70 := GenS1{F0: 45, F1: -1, F2: true}
    for fi71 := 0; fi71 < 17; fi71++ {
        if true {
            for fi72 := 0; fi72 < 5; fi72++ {
                switch "foo bar" {
                case "":
                    v73 := GenS7{F0: (^(v5)), F1: 4294967295, F2: 2}
                    _ = v73
                case "hello":
                    fmt.Println(v1, v5)
                default:
                    v4[2147483648] = "foo bar"
                }
                if true {
                    var v74 int32 = 2147483647
                    var v75 int = fi72
                    _ = v74
                    _ = v75
                }
                v76 := GenS8{F0: 0}
                _ = fi72
                _ = v76
            }
            _ = len(v2)
            var v77 int64 = 1
            _ = v77
        } else {
            for fi78 := 0; fi78 < 7; fi78++ {
                fmt.Println(v3, fi71)
                var v79 uint = 18446744073709551615
                _ = fi78
                _ = v79
            }
        }
        fmt.Println(v3, v5, fi71)
        v5 = 0
        _ = fi71
    }
    var v80 string = ("" + "")
    v5 -= (v5 >> ((0) & 63))
    for fi81 := 0; fi81 < 2; fi81++ {
        v82 := []int{}
        v82 = append(v82, -115)
        _ = fi81
        _ = v82
    }
    v83 := GenS1{F0: 255, F1: -543, F2: false}
    for fi84 := 0; fi84 < 16; fi84++ {
        for fi85 := 0; fi85 < 11; fi85++ {
            v86 := GenS9{F0: 255}
            var v87 float32 = float32(v1)
            _ = fi85
            _ = v86
            _ = v87
        }
        var v88 uint = 65535
        _ = fi84
        _ = v88
    }
    v3 = float64(float32(((1.5 - v1) - -2.25)))
    if false {
        for fi89 := 0; fi89 < 15; fi89++ {
            v90 := GenS10{F0: 32767, F1: (9 > 1), F2: -2147483648}
            _ = fi89
            _ = v90
        }
        for fi91 := 0; fi91 < 15; fi91++ {
            if (fi91 == fi91) {
                delete(v2, (v80 + v80))
                var v92 uint64 = 4294967295
                v5 = v5
                _ = v92
            } else {
                var v93 int8 = genFn1(42, 657)
                fmt.Println(fi91, v5, v3)
                _ = v93
            }
            for fi94 := 0; fi94 < 3; fi94++ {
                v95 := []int{}
                v3 = v3
                var v96 int8 = 100
                var v97 string = "hello"
                _ = fi94
                _ = v95
                _ = v96
                _ = v97
            }
            _ = fi91
        }
        var v98 uint32 = 938
        _ = v98
    } else {
        v99 := []int{}
        _ = v99
    }
    v80 = "a"
    v80 = v80
    if ("hello" < "") {
        var v100 int = 195
        var v101 uint8 = 255
        _ = v100
        _ = v101
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v70
    _ = v80
    _ = v83
}
