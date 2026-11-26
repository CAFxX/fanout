package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 uint; F1 uint16 }
func genFn0() int {
    v6 := []int{32768}
    var v7 uint32 = 2147483648
    var v8 uint8 = 115
    v9 := GenS1{F0: -128, F1: -88}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    switch ("\n\t" + p0) {
    case "a":
        v10 := []int{-96, 831}
        _ = v10
    case "":
        fmt.Println(p0)
    case "hello":
        p0 = p0
    default:
        v11 := []int{4294967295, 4294967295, 32767}
        _ = v11
    }
    v12 := GenS2{F0: 4294967295, F1: 31, F2: 895}
    switch "foo bar" {
    case "foo bar":
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{(fi13 ^ (fi13 & (-(fi13)))), (fi13 + genFn0()), (fi13 | fi13), (fi13 ^ fi13)}
            _ = fi13
            _ = v14
        }
    case "hello":
        p0 = p0
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS3{F0: 334}
        for fi17 := 0; fi17 < 7; fi17++ {
            if (1.5 != float32((float64(3.14159) * float64((1.5 * 0.0))))) {
                fmt.Println(p0, fi17, fi17)
            } else {
                v18 := []int{(fi15 / (((fi17 & fi17)) | 1)), fi15, (fi17 | 32767), fi17}
                _ = v18
            }
            var v19 int32 = -2147483648
            _ = fi17
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v12
    return (-1 == -1)
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 128
    var v3 uint8 = v2
    var v4 int16 = 1
    v5 := make(map[uint32]string)
    for fi20 := 0; fi20 < 10; fi20++ {
        v4 -= (v4 + v4)
        _ = fi20
    }
    fmt.Println(v2, v2, v4)
    v21 := []int{1}
    var v22 bool = false
    v3 = v3
    v23 := GenS1{F0: 42, F1: 100}
    if true {
        v1 = 2147483647
        fmt.Println(v4)
    } else {
        if true {
            var v24 uint16 = 42
            _ = v24
        } else {
            switch (("\n\t" + "\n\t") + "foo bar") {
            case "":
                var v25 string = (("hello" + "\n\t") + ((("a" + (("" + ("foo bar" + "")) + "hello")) + ("\n\t" + ("" + ""))) + ("a" + ("" + ("hello" + ("a" + "\n\t"))))))
                _ = v25
            case "foo bar":
                if v22 {
                    if true {
                        v26 := GenS4{F0: uint8(v1)}
                        fmt.Println(v2)
                        _ = v26
                    } else {
                        switch "foo bar" {
                        case "":
                            if (34 >= 42) {
                                if false {
                                    for fi27 := 0; fi27 < 2; fi27++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v28 float64 = 0.0
                                            _ = v28
                                        case "":
                                            var v29 int64 = -1
                                            _ = v29
                                        }
                                        _ = fi27
                                    }
                                    for fi30 := 0; fi30 < 4; fi30++ {
                                        var v31 bool = false
                                        v32 := GenS3{F0: 255}
                                        _ = fi30
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    switch "a" {
                                    case "hello":
                                        v33 := GenS5{F0: 1, F1: v1, F2: 32767}
                                        _ = v33
                                    case "foo bar":
                                        for fi34 := 0; fi34 < 15; fi34++ {
                                            var v35 int32 = -1
                                            for fi36 := 0; fi36 < 8; fi36++ {
                                                v22 = (!v22)
                                                _ = fi36
                                            }
                                            _ = fi34
                                            _ = v35
                                        }
                                    }
                                }
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    if v22 {
                                        if v22 {
                                            for fi38 := 0; fi38 < 4; fi38++ {
                                                if (v1 < v1) {
                                                    v1 += v1
                                                } else {
                                                    if true {
                                                        v39 := v5[(v1 | 4294967295)]
                                                        _ = v39
                                                    } else {
                                                        v5[(^(v1))] = ("a" + (("foo bar" + ("hello" + "")) + "foo bar"))
                                                        if false {
                                                            if ((v22 || (2 != 774)) && v22) {
                                                                fmt.Println(v3)
                                                            }
                                                            var v40 int64 = -9223372036854775808
                                                            _ = v40
                                                        }
                                                        if ((!v22) || false) {
                                                            v41 := GenS6{F0: "\n\t", F1: 1}
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v22 = v22
                                                                var v42 int = (-(fi38))
                                                                _ = v42
                                                            case "hello":
                                                                if true {
                                                                    v1 += uint32(v3)
                                                                }
                                                            case "":
                                                                v22 = ("" == "a")
                                                            }
                                                            _ = v41
                                                        }
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if true {
                                                            v2 = v2
                                                            fmt.Println(fi37, v1)
                                                            fmt.Println(v3, v2)
                                                        } else {
                                                            if (-1 <= -9223372036854775808) {
                                                                if (true && v22) {
                                                                    var v43 uint32 = v1
                                                                    _ = v43
                                                                }
                                                                switch (("\n\t" + "a") + "hello") {
                                                                case "foo bar":
                                                                    if v22 {
                                                                        if false {
                                                                            v44 := GenS7{F0: 18446744073709551615}
                                                                            switch fi38 {
                                                                            case 2147483647:
                                                                                if true {
                                                                                    v2 -= v2
                                                                                    if v22 {
                                                                                        fmt.Println(v22, v22, v2)
                                                                                        v1 = 2
                                                                                        if ((v1 - 65535) <= v1) {
                                                                                            v22 = v22
                                                                                            if len(v21) >= 2 {
                                                                                                v21 = v21[1:len(v21)]
                                                                                            }
                                                                                        }
                                                                                        v45 := GenS6{F0: "", F1: 1}
                                                                                        _ = v45
                                                                                    }
                                                                                } else {
                                                                                    var v46 int64 = -215
                                                                                    _ = v46
                                                                                }
                                                                            default:
                                                                                fmt.Println(v2, v3, v2)
                                                                            }
                                                                            switch 9223372036854775807 {
                                                                            case -1:
                                                                                _ = len(v21)
                                                                            case 22:
                                                                                v47 := v5[(v1 - v1)]
                                                                                _ = v47
                                                                            default:
                                                                                v48 := GenS4{F0: v2}
                                                                                _ = v48
                                                                            }
                                                                            v2 -= (^(v2))
                                                                            _ = v44
                                                                        }
                                                                        v1 -= (v1 << ((801) & 63))
                                                                        if ((!(v22 || ((!(!false)) && ((0.1 == 1e10) || false)))) && true) {
                                                                            if true {
                                                                                v49 := GenS1{F0: 1, F1: -25}
                                                                                fmt.Println(v1, v3, v22)
                                                                                _ = v49
                                                                            } else {
                                                                                if (0 < 255) {
                                                                                    v50 := GenS5{F0: -63, F1: (v1 - v1), F2: 864}
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        var v51 uint64 = 42
                                                                                        _ = v51
                                                                                    case "foo bar":
                                                                                        v4 += (v4 + 0)
                                                                                    case "\n\t":
                                                                                        var v52 int8 = -74
                                                                                        _ = v52
                                                                                    }
                                                                                    _ = v50
                                                                                } else {
                                                                                    v3 = 2
                                                                                }
                                                                                v53 := GenS1{F0: 1, F1: -1}
                                                                                _ = v53
                                                                            }
                                                                            v54 := GenS7{F0: uint(v3)}
                                                                            _ = v54
                                                                        } else {
                                                                            switch 32768 {
                                                                            case 710:
                                                                                var v55 uint8 = (v2 + 2)
                                                                                _ = v55
                                                                            case -9223372036854775808:
                                                                                v56 := GenS6{F0: (((("foo bar" + ("foo bar" + ("a" + "\n\t"))) + "\n\t") + (("foo bar" + "hello") + "foo bar")) + "foo bar"), F1: -1}
                                                                                _ = v56
                                                                            case 626:
                                                                                v4 -= 42
                                                                            }
                                                                        }
                                                                        delete(v5, v1)
                                                                        v2 = v3
                                                                        var v57 uint8 = v3
                                                                        _ = v57
                                                                    } else {
                                                                        v58 := GenS4{F0: 174}
                                                                        if v22 {
                                                                            var v59 bool = ("a" != "hello")
                                                                            v1 += v1
                                                                            if v22 {
                                                                                v60 := GenS8{F0: v4, F1: 128}
                                                                                if len(v21) >= 2 {
                                                                                    v21 = v21[1:len(v21)]
                                                                                }
                                                                                v61 := GenS6{F0: ("hello" + "\n\t"), F1: -128}
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                            _ = v59
                                                                        } else {
                                                                            v22 = (!(88 < 1))
                                                                            if v22 {
                                                                                v62 := GenS9{F0: 2, F1: v3}
                                                                                v63 := GenS10{F0: 379, F1: 415}
                                                                                var v64 int8 = -1
                                                                                _ = v62
                                                                                _ = v63
                                                                                _ = v64
                                                                            }
                                                                            v22 = false
                                                                        }
                                                                        _ = v58
                                                                    }
                                                                case "\n\t":
                                                                    v4 = v4
                                                                case "a":
                                                                    fmt.Println(v22, v2)
                                                                default:
                                                                    var v65 int16 = v4
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                    case "a":
                                                        v4 += (v4 % ((v4) | 1))
                                                    case "\n\t":
                                                        if len(v21) > 0 {
                                                            v21[0] = 127
                                                        }
                                                    }
                                                }
                                                fmt.Println(v2, v2)
                                                v66 := GenS3{F0: 42}
                                                _ = fi38
                                                _ = v66
                                            }
                                        } else {
                                            v67 := GenS9{F0: 2, F1: v3}
                                            _ = v67
                                        }
                                    } else {
                                        for fi68 := 0; fi68 < 7; fi68++ {
                                            var v69 uint8 = 42
                                            _ = fi68
                                            _ = v69
                                        }
                                    }
                                    v2 -= (^(v2))
                                    _ = fi37
                                }
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v70 uint8 = (v3 - v3)
                            _ = v70
                        }
                    }
                    v3 = (^(v2))
                }
            case "hello":
                if (4294967295 >= 588) {
                    for fi71 := 0; fi71 < 5; fi71++ {
                        var v72 uint64 = 147
                        var v73 uint32 = v1
                        for fi74 := 0; fi74 < 9; fi74++ {
                            if (!(!false)) {
                                v3 = 255
                                if (42 <= v4) {
                                    var v75 uint32 = (v73 % (((v73 & v1)) | 1))
                                    _ = v75
                                }
                                var v76 float64 = 3.14159
                                _ = v76
                            }
                            v3 -= 0
                            var v77 uint = 0
                            _ = fi74
                            _ = v77
                        }
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    if (!v22) {
                        var v78 string = "foo bar"
                        _ = v78
                    }
                    fmt.Println(v4, v1, v3)
                } else {
                    v79 := GenS10{F0: 1, F1: 255}
                    fmt.Println(v22, v22, v22)
                    for fi80 := 0; fi80 < 5; fi80++ {
                        var v81 uint32 = (v1 | (v1 << ((757) & 63)))
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        v2 = v3
                        _ = fi80
                        _ = v81
                    }
                    _ = v79
                }
            default:
                v22 = true
            }
            for fi82 := 0; fi82 < 10; fi82++ {
                v4 = v4
                _ = fi82
            }
        }
        v22 = (((v4 > -414) || v22) && v22)
        v21 = append(v21, 554)
    }
    var v83 int32 = 42
    if v22 {
        v1 = (v1 + v1)
    }
    var v84 uint8 = (v2 & (^(v2)))
    if (783 != 32768) {
        v22 = v22
    } else {
        var v85 int64 = 0
        var v86 uint32 = (v1 & (v1 / ((42) | 1)))
        _ = v85
        _ = v86
    }
    if v22 {
        var v87 int16 = v4
        _ = v87
    } else {
        v1 += 127
    }
    var v88 int8 = -128
    var v89 int = 32767
    fmt.Println(v2)
    var v90 int8 = v88
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
    _ = v88
    _ = v89
    _ = v90
}
