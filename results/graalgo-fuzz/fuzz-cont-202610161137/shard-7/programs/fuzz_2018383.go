package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint; F1 string; F2 uint }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 int8; F1 uint16 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 uint) int32 {
    switch "a" {
    case "":
        var v6 uint8 = 137
        _ = v6
    case "hello":
        var v7 int8 = 123
        _ = v7
    case "\n\t":
        v8 := []int{}
        _ = v8
    default:
        v9 := GenS1{F0: 32767, F1: 1, F2: (p1 | (p1 % (((p1 / ((p1) | 1))) | 1)))}
        _ = v9
    }
    var v10 int = -101
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint32 = 690
        var v13 int32 = 40
        if true {
            if ((v13 | v13) == v13) {
                fmt.Println(fi11, p0, v10)
            } else {
                if ((v13 / (((v13 & (-(v13)))) | 1)) < 42) {
                    v14 := GenS1{F0: 1, F1: 0, F2: uint(v12)}
                    _ = v14
                }
                p0 -= int16(v13)
            }
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(p1)
                switch "hello" {
                case "a":
                    v16 := GenS2{F0: 42}
                    _ = v16
                default:
                    v17 := []int{(fi11 - v10)}
                    _ = v17
                }
                v13 = -709
                var v18 int = (-(v10))
                var v19 int = (v10 | v10)
                _ = fi15
                _ = v18
                _ = v19
            }
        } else {
            fmt.Println(v10, fi11)
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v20 uint8 = uint8(p1)
    v21 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v20
    _ = v21
    return 42
}
func genFn1() int8 {
    v22 := []int{-878, -357, 32767}
    if false {
        switch 32768 {
        case -1:
            if (1 > -922) {
                v23 := GenS1{F0: 2, F1: 1, F2: 1}
                _ = v23
            }
        default:
            switch "foo bar" {
            case "a":
                switch (("\n\t" + ("\n\t" + ("\n\t" + "hello"))) + "hello") {
                case "hello":
                    var v24 uint8 = 0
                    _ = v24
                case "\n\t":
                    for fi25 := 0; fi25 < 1; fi25++ {
                        v26 := GenS3{F0: 2, F1: fi25}
                        _ = fi25
                        _ = v26
                    }
                case "foo bar":
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v22 = append(v22, 742)
                        var v28 uint16 = 65535
                        if (false && false) {
                            if true {
                                var v29 uint8 = 198
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    var v31 float32 = (3.14159 * 0.0)
                                    v22 = append(v22, fi27)
                                    v29 += 128
                                    _ = fi30
                                    _ = v31
                                }
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                _ = v29
                            } else {
                                for fi32 := 0; fi32 < 16; fi32++ {
                                    v28 = v28
                                    _ = len(v22)
                                    _ = fi32
                                }
                                v33 := GenS4{F0: 291, F1: "foo bar", F2: 4294967295}
                                var v34 int = fi27
                                _ = v33
                                _ = v34
                            }
                        } else {
                            v28 = (v28 | (v28 * 65535))
                        }
                        _ = fi27
                        _ = v28
                    }
                }
            case "\n\t":
                v35 := GenS5{F0: (!true), F1: 1000}
                _ = v35
            case "foo bar":
                for fi36 := 0; fi36 < 9; fi36++ {
                    var v37 string = "a"
                    _ = fi36
                    _ = v37
                }
            }
        }
        var v38 bool = false
        fmt.Println(v38)
        _ = v38
    } else {
        v39 := GenS2{F0: 2}
        _ = v39
    }
    var v40 bool = false
    if ((!(42 <= 18446744073709551615)) || v40) {
        if false {
            var v41 int8 = 0
            for fi42 := 0; fi42 < 8; fi42++ {
                switch ((("hello" + ("\n\t" + "")) + "hello") + "foo bar") {
                case "\n\t":
                    fmt.Println(v41)
                case "":
                    for fi43 := 0; fi43 < 20; fi43++ {
                        v44 := GenS2{F0: 283}
                        _ = fi43
                        _ = v44
                    }
                }
                var v45 uint32 = 300
                _ = fi42
                _ = v45
            }
            _ = v41
        }
    } else {
        fmt.Println(v40)
        for fi46 := 0; fi46 < 4; fi46++ {
            v40 = v40
            v47 := GenS6{F0: 2, F1: 1000}
            v40 = (("foo bar" + "\n\t") <= "hello")
            _ = fi46
            _ = v47
        }
        v48 := GenS2{F0: 57}
        _ = v48
    }
    if len(v22) > 0 {
        v22[0] = -357
    }
    _ = v22
    _ = v40
    return 36
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[int]string)
    var v3 uint16 = 255
    v4 := []uint32{4294967295, 42, 65535, 151}
    v5 := make(map[string]string)
    fmt.Println(v3)
    fmt.Println(v3)
    delete(v5, "")
    v3 = v3
    for fi49 := 0; fi49 < 6; fi49++ {
        fmt.Println(v3, fi49)
        if true {
            v3 = (v3 >> ((0) & 63))
        }
        if true {
            if (!(true && (1.5 > 0.1))) {
                v50 := GenS7{F0: 885, F1: 100}
                for fi51 := 0; fi51 < 13; fi51++ {
                    v52 := GenS2{F0: 1}
                    _ = len(v5)
                    var v53 uint = 2
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                _ = v50
            }
            for fi54 := 0; fi54 < 11; fi54++ {
                if false {
                    var v55 uint16 = 65535
                    if (601 != 47) {
                        fmt.Println(v3, fi49)
                    }
                    v56 := GenS8{F0: 42, F1: -456}
                    _ = v55
                    _ = v56
                } else {
                    v57 := GenS9{F0: v3, F1: true}
                    fmt.Println(fi49, fi54, fi49)
                    _ = len(v4)
                    _ = v57
                }
                delete(v5, "\n\t")
                switch 0 {
                case 725:
                    v4 = append(v4, 1)
                default:
                    if (-1 < -1) {
                        v58 := GenS5{F0: true, F1: 107}
                        v3 -= 65535
                        _ = v58
                    } else {
                        v3 += (v3 << ((255) & 63))
                    }
                }
                _ = fi54
            }
            for fi59 := 0; fi59 < 10; fi59++ {
                switch "" {
                case "\n\t":
                    v60 := GenS10{F0: 923}
                    _ = v60
                default:
                    v3 = v3
                }
                v1 = append(v1, 459)
                v3 = (^(v3))
                if false {
                    fmt.Println(fi49, fi59)
                } else {
                    v3 -= v3
                    v3 -= (v3 / ((v3) | 1))
                    if ((478 < 988) && (!false)) {
                        var v61 uint8 = 1
                        v62 := GenS3{F0: (v61 + (v61 * (^(v61)))), F1: fi59}
                        _ = v61
                        _ = v62
                    } else {
                        fmt.Println(fi49, fi49)
                        fmt.Println(fi59)
                        var v63 uint64 = 711
                        _ = v63
                    }
                }
                if true {
                    v64 := GenS9{F0: v3, F1: (true || (1.5 > 0.0))}
                    v3 += (v3 ^ (v3 % ((v3) | 1)))
                    _ = v64
                } else {
                    var v65 bool = true
                    v65 = v65
                    _ = v65
                }
                if len(v1) > 0 {
                    v1[0] = 697
                }
                _ = fi59
            }
        }
        _ = fi49
    }
    v66 := v5["hello"]
    var v67 int32 = genFn0(42, 484)
    var v68 bool = (v66 == (("foo bar" + (v66 + "")) + v66))
    for fi69 := 0; fi69 < 16; fi69++ {
        switch ((("a" + v66) + (v66 + "foo bar")) + ("foo bar" + (v66 + v66))) {
        case "\n\t":
            if v68 {
                var v70 uint16 = v3
                if v68 {
                    var v71 int32 = (v67 ^ (v67 % ((-1) | 1)))
                    v3 -= (^(v3))
                    _ = v71
                }
                _ = v70
            } else {
                for fi72 := 0; fi72 < 9; fi72++ {
                    v66 = ("a" + v66)
                    _ = fi72
                }
                if v68 {
                    var v73 int = fi69
                    v68 = (!true)
                    _ = v73
                }
            }
        }
        v68 = (1.5 <= (((1.5 / -2.25) - 1.5) * -2.25))
        _ = fi69
    }
    if v68 {
        _ = len(v1)
        if (-2147483648 > 32767) {
            fmt.Println(v68, v68, v66)
            var v74 int32 = (v67 & genFn0(1, 4294967295))
            if (!(!(v74 < v74))) {
                var v75 uint8 = 129
                switch 32768 {
                case 32767:
                    if false {
                        var v76 int64 = -1
                        for fi77 := 0; fi77 < 1; fi77++ {
                            v78 := GenS6{F0: 0, F1: 1000}
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                            v3 += v3
                            _ = fi77
                            _ = v78
                        }
                        switch v66 {
                        case "hello":
                            switch "foo bar" {
                            case "\n\t":
                                v79 := GenS11{F0: (484 > (v76 % ((-9223372036854775808) | 1))), F1: 1}
                                _ = v79
                            }
                        case "foo bar":
                            v2[2147483647] = v66
                        case "":
                            v80 := GenS12{F0: 100, F1: v3}
                            _ = v80
                        default:
                            var v81 int8 = -72
                            _ = v81
                        }
                        _ = v76
                    }
                case 0:
                    var v82 uint32 = 2
                    _ = v82
                default:
                    var v83 uint = 42
                    _ = v83
                }
                _ = v75
            }
            _ = v74
        }
        var v84 uint32 = 85
        _ = v84
    }
    if (-502 >= 870) {
        v85 := GenS13{F0: (("\n\t" + "") + ("hello" + "a"))}
        if false {
            v67 = (v67 + 1)
            var v86 uint8 = 1
            _ = v86
        } else {
            if (v3 > v3) {
                var v87 uint16 = (v3 - v3)
                _ = v87
            } else {
                var v88 uint16 = (v3 % (((v3 << ((2) & 63))) | 1))
                var v89 int16 = int16(v67)
                for fi90 := 0; fi90 < 12; fi90++ {
                    v66 = v66
                    for fi91 := 0; fi91 < 2; fi91++ {
                        v89 = v89
                        v92 := GenS11{F0: true, F1: (v89 ^ v89)}
                        _ = fi91
                        _ = v92
                    }
                    v93 := GenS14{F0: (^(v67)), F1: (v66 + v66)}
                    _ = fi90
                    _ = v93
                }
                _ = v88
                _ = v89
            }
        }
        _ = v85
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v68
}
